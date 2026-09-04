package x402api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// PaymentReceiptPendingError is returned by Execute when the receipt endpoint
// returns HTTP 202 with a valid PaymentReceiptStatus. It is a polling outcome,
// not proof that the request or payment failed.
type PaymentReceiptPendingError struct {
	Status     *PaymentReceiptStatus
	RetryAfter string
	body       []byte
}

// Error implements error.
func (e *PaymentReceiptPendingError) Error() string {
	if e == nil || e.Status == nil {
		return "payment receipt is pending"
	}
	return fmt.Sprintf(
		"payment receipt for %s is %s; retry after %s",
		e.Status.PaymentId,
		e.Status.ReceiptStatus,
		e.RetryAfter,
	)
}

// Body returns a copy of the raw HTTP 202 response body.
func (e *PaymentReceiptPendingError) Body() []byte {
	if e == nil {
		return nil
	}
	return append([]byte(nil), e.body...)
}

// PaymentReceiptResult contains exactly one of Receipt or PendingStatus.
type PaymentReceiptResult struct {
	StatusCode    int
	Receipt       *PaymentReceipt
	PendingStatus *PaymentReceiptStatus
	RetryAfter    string
}

// IsFinalized reports whether the result contains a signed finalized receipt.
func (r *PaymentReceiptResult) IsFinalized() bool {
	return r != nil && r.StatusCode == http.StatusOK && r.Receipt != nil && r.PendingStatus == nil
}

// IsPending reports whether the result contains a valid HTTP 202 status.
func (r *PaymentReceiptResult) IsPending() bool {
	return r != nil && r.StatusCode == http.StatusAccepted && r.Receipt == nil && r.PendingStatus != nil
}

// IsConfirmed reports whether chain confirmation has been reached. A signed
// receipt necessarily represents a finalized, and therefore confirmed, payment.
func (r *PaymentReceiptResult) IsConfirmed() bool {
	return r != nil && (r.IsFinalized() || (r.IsPending() && r.PendingStatus.Confirmed))
}

// ExecuteStatusAware performs one receipt request and returns a typed outcome
// for either HTTP 200 or HTTP 202. Other responses and malformed success
// payloads remain errors.
func (r ApiPaymentsRetrieveReceiptRequest) ExecuteStatusAware() (*PaymentReceiptResult, *http.Response, error) {
	receipt, response, err := r.Execute()
	if err != nil {
		var pending *PaymentReceiptPendingError
		if !errors.As(err, &pending) {
			return nil, response, err
		}
		if response == nil || response.StatusCode != http.StatusAccepted || pending.Status == nil {
			return nil, response, paymentReceiptProtocolError(
				pending.Body(),
				"pending result did not originate from HTTP 202",
			)
		}
		return &PaymentReceiptResult{
			StatusCode:    http.StatusAccepted,
			PendingStatus: pending.Status,
			RetryAfter:    pending.RetryAfter,
		}, response, nil
	}

	if response == nil {
		return nil, nil, paymentReceiptProtocolError(nil, "missing HTTP response")
	}
	if response.StatusCode != http.StatusOK {
		return nil, response, paymentReceiptProtocolError(
			copyResponseBody(response),
			"unexpected successful HTTP status %d",
			response.StatusCode,
		)
	}
	if validationErr := validateFinalizedPaymentReceipt(r.id, receipt); validationErr != nil {
		return nil, response, paymentReceiptProtocolError(
			copyResponseBody(response),
			validationErr.Error(),
		)
	}

	return &PaymentReceiptResult{
		StatusCode: http.StatusOK,
		Receipt:    receipt,
	}, response, nil
}

func newPaymentReceiptPendingError(
	client *APIClient,
	expectedPaymentID string,
	response *http.Response,
	body []byte,
) (*PaymentReceiptPendingError, error) {
	if client == nil || response == nil || response.StatusCode != http.StatusAccepted {
		return nil, paymentReceiptProtocolError(body, "invalid HTTP 202 decoder invocation")
	}
	if err := requirePendingStatusFields(body); err != nil {
		return nil, paymentReceiptProtocolError(body, err.Error())
	}

	var status PaymentReceiptStatus
	if err := client.decode(&status, body, response.Header.Get("Content-Type")); err != nil {
		return nil, paymentReceiptProtocolError(body, "cannot decode HTTP 202 status: %v", err)
	}
	if err := validatePendingPaymentReceiptStatus(expectedPaymentID, &status); err != nil {
		return nil, paymentReceiptProtocolError(body, err.Error())
	}

	retryAfter := strings.TrimSpace(response.Header.Get("Retry-After"))
	if !isValidRetryAfter(retryAfter) {
		return nil, paymentReceiptProtocolError(body, "missing or invalid Retry-After header")
	}

	return &PaymentReceiptPendingError{
		Status:     &status,
		RetryAfter: retryAfter,
		body:       append([]byte(nil), body...),
	}, nil
}

func requirePendingStatusFields(body []byte) error {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(body, &object); err != nil {
		return fmt.Errorf("cannot decode HTTP 202 status: %w", err)
	}
	if object == nil {
		return errors.New("HTTP 202 status must be a JSON object")
	}

	required := []string{
		"payment_id",
		"state",
		"confirmed",
		"finalized",
		"confirmed_at",
		"finalized_at",
		"transaction",
		"network",
		"receipt_status",
	}
	nullable := map[string]bool{"confirmed_at": true, "finalized_at": true}
	for _, name := range required {
		value, ok := object[name]
		if !ok {
			return fmt.Errorf("HTTP 202 status is missing required field %q", name)
		}
		if !nullable[name] && bytes.Equal(bytes.TrimSpace(value), []byte("null")) {
			return fmt.Errorf("HTTP 202 status field %q cannot be null", name)
		}
	}
	return nil
}

func validatePendingPaymentReceiptStatus(expectedPaymentID string, status *PaymentReceiptStatus) error {
	if status == nil {
		return errors.New("HTTP 202 status is missing")
	}
	if status.PaymentId != expectedPaymentID {
		return errors.New("HTTP 202 payment_id does not match the request")
	}
	if strings.TrimSpace(status.State) == "" || strings.TrimSpace(status.Network) == "" {
		return errors.New("HTTP 202 state and network must be non-empty")
	}
	if status.Finalized || status.FinalizedAt.Get() != nil {
		return errors.New("HTTP 202 status cannot be finalized")
	}
	if !status.ConfirmedAt.IsSet() || !status.FinalizedAt.IsSet() {
		return errors.New("HTTP 202 confirmation timestamps must be present, including null")
	}

	switch status.ReceiptStatus {
	case RECEIPTSTATUSENUM_PENDING_CONFIRMATION:
		if status.Confirmed {
			return errors.New("pending_confirmation status cannot be confirmed")
		}
	case RECEIPTSTATUSENUM_PENDING_FINALITY:
		if !status.Confirmed || status.ConfirmedAt.Get() == nil {
			return errors.New("pending_finality status requires confirmation evidence")
		}
		if strings.TrimSpace(status.Transaction) == "" {
			return errors.New("pending_finality status requires a transaction")
		}
	default:
		return errors.New("HTTP 202 receipt_status is unknown")
	}
	return nil
}

func validateFinalizedPaymentReceipt(expectedPaymentID string, receipt *PaymentReceipt) error {
	if receipt == nil {
		return errors.New("HTTP 200 receipt is missing")
	}
	if receipt.SettlementJobId != expectedPaymentID {
		return errors.New("HTTP 200 settlement_job_id does not match the request")
	}
	if strings.TrimSpace(receipt.Id) == "" || strings.TrimSpace(receipt.OrderId) == "" {
		return errors.New("HTTP 200 receipt and order identifiers must be non-empty")
	}
	if receipt.Receipt == nil || strings.TrimSpace(receipt.ReceiptDigest) == "" ||
		strings.TrimSpace(receipt.Signature) == "" || strings.TrimSpace(receipt.SigningKeyVersion) == "" {
		return errors.New("HTTP 200 signed receipt evidence is incomplete")
	}
	if strings.TrimSpace(receipt.SettlementAmountAtomic) == "" ||
		strings.TrimSpace(receipt.GasMode) == "" || receipt.CreatedAt.IsZero() {
		return errors.New("HTTP 200 receipt settlement fields are incomplete")
	}
	return nil
}

func isValidRetryAfter(value string) bool {
	if value == "" {
		return false
	}
	if _, err := strconv.ParseUint(value, 10, 64); err == nil {
		return true
	}
	_, err := http.ParseTime(value)
	return err == nil
}

func paymentReceiptProtocolError(body []byte, format string, args ...interface{}) error {
	return &GenericOpenAPIError{
		body:  append([]byte(nil), body...),
		error: "invalid payment receipt response: " + fmt.Sprintf(format, args...),
	}
}

func copyResponseBody(response *http.Response) []byte {
	if response == nil || response.Body == nil {
		return nil
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewReader(body))
	return body
}
