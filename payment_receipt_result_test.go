package x402api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

const receiptTestPaymentID = "00000000-0000-4000-8000-000000000001"

func TestPaymentsRetrieveReceiptExecuteReturnsTypedPendingError(t *testing.T) {
	body := validPendingFinalityBody(receiptTestPaymentID)
	client, requests := newReceiptTestClient(t, http.StatusAccepted, body, "2")

	receipt, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		Execute()
	closeResponse(t, response)

	if receipt != nil {
		t.Fatalf("Execute() receipt = %#v, want nil", receipt)
	}
	var pending *PaymentReceiptPendingError
	if !errors.As(err, &pending) {
		t.Fatalf("Execute() error = %T %v, want *PaymentReceiptPendingError", err, err)
	}
	if response == nil || response.StatusCode != http.StatusAccepted {
		t.Fatalf("Execute() response = %#v, want HTTP 202", response)
	}
	if pending.Status.PaymentId != receiptTestPaymentID || !pending.Status.Confirmed {
		t.Fatalf("pending status = %#v, want confirmed matching payment", pending.Status)
	}
	if pending.RetryAfter != "2" {
		t.Fatalf("RetryAfter = %q, want 2", pending.RetryAfter)
	}
	if string(pending.Body()) != body {
		t.Fatalf("Body() = %q, want %q", pending.Body(), body)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwarePending(t *testing.T) {
	client, requests := newReceiptTestClient(
		t,
		http.StatusAccepted,
		validPendingFinalityBody(receiptTestPaymentID),
		"2",
	)

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if err != nil {
		t.Fatalf("ExecuteStatusAware() error = %v", err)
	}
	if result == nil || !result.IsPending() || !result.IsConfirmed() || result.IsFinalized() {
		t.Fatalf("ExecuteStatusAware() result = %#v, want confirmed pending result", result)
	}
	if result.PendingStatus.PaymentId != receiptTestPaymentID || result.RetryAfter != "2" {
		t.Fatalf("pending result = %#v, want matching payment and Retry-After", result)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwarePendingConfirmation(t *testing.T) {
	body := fmt.Sprintf(`{
  "payment_id": %q,
  "state": "reserved",
  "confirmed": false,
  "finalized": false,
  "confirmed_at": null,
  "finalized_at": null,
  "transaction": "",
  "network": "eip155:8453",
  "receipt_status": "pending_confirmation"
}`, receiptTestPaymentID)
	client, requests := newReceiptTestClient(t, http.StatusAccepted, body, "2")

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if err != nil {
		t.Fatalf("ExecuteStatusAware() error = %v", err)
	}
	if result == nil || !result.IsPending() || result.IsConfirmed() {
		t.Fatalf("ExecuteStatusAware() result = %#v, want unconfirmed pending result", result)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwareFinalized(t *testing.T) {
	body := fmt.Sprintf(`{
  "id": "00000000-0000-4000-8000-000000000002",
  "order_id": "00000000-0000-4000-8000-000000000003",
  "settlement_job_id": %q,
  "receipt": {"receipt_version": 2},
  "receipt_digest": "sha256:receipt",
  "signature": "signature",
  "signing_key_version": "v1",
  "eligible_alternatives": [],
  "fee_policy": null,
  "fee_evidence": null,
  "fee_quote_digest": null,
  "fee_quote_expires_at": null,
  "settlement_amount_atomic": "25000000",
  "gas_mode": "buyer_pays",
  "buyer_native_fee_atomic": "0",
  "sponsored_native_fee_atomic": null,
  "sponsored_native_symbol": null,
  "tenant_gas_charge_micros": null,
  "gas_sponsorship_evidence_digest": null,
  "created_at": "2026-09-04T00:00:01Z"
}`, receiptTestPaymentID)
	client, requests := newReceiptTestClient(t, http.StatusOK, body, "")

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if err != nil {
		t.Fatalf("ExecuteStatusAware() error = %v", err)
	}
	if result == nil || !result.IsFinalized() || !result.IsConfirmed() || result.IsPending() {
		t.Fatalf("ExecuteStatusAware() result = %#v, want finalized result", result)
	}
	if result.Receipt.SettlementJobId != receiptTestPaymentID || result.PendingStatus != nil {
		t.Fatalf("finalized result = %#v, want matching receipt only", result)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptRejectsMalformedOrMismatched202(t *testing.T) {
	otherPaymentID := "00000000-0000-4000-8000-000000000099"
	tests := []struct {
		name       string
		body       string
		retryAfter string
	}{
		{name: "malformed JSON", body: `{"payment_id":`, retryAfter: "2"},
		{name: "mismatched payment", body: validPendingFinalityBody(otherPaymentID), retryAfter: "2"},
		{name: "missing required field", body: strings.Replace(validPendingFinalityBody(receiptTestPaymentID), "  \"finalized_at\": null,\n", "", 1), retryAfter: "2"},
		{name: "unknown receipt status", body: strings.Replace(validPendingFinalityBody(receiptTestPaymentID), "pending_finality", "future_state", 1), retryAfter: "2"},
		{name: "inconsistent confirmation", body: strings.Replace(validPendingFinalityBody(receiptTestPaymentID), `"confirmed": true`, `"confirmed": false`, 1), retryAfter: "2"},
		{name: "finalized 202", body: strings.Replace(validPendingFinalityBody(receiptTestPaymentID), `"finalized": false`, `"finalized": true`, 1), retryAfter: "2"},
		{name: "missing Retry-After", body: validPendingFinalityBody(receiptTestPaymentID), retryAfter: ""},
		{name: "invalid Retry-After", body: validPendingFinalityBody(receiptTestPaymentID), retryAfter: "soon"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, requests := newReceiptTestClient(t, http.StatusAccepted, test.body, test.retryAfter)

			receipt, response, err := client.OrdersAndPaymentsAPI.
				PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
				Execute()
			closeResponse(t, response)

			if receipt != nil {
				t.Fatalf("Execute() receipt = %#v, want nil", receipt)
			}
			var pending *PaymentReceiptPendingError
			if errors.As(err, &pending) {
				t.Fatalf("Execute() error = %#v, must not expose malformed 202 as pending", pending)
			}
			var apiErr *GenericOpenAPIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("Execute() error = %T %v, want *GenericOpenAPIError", err, err)
			}
			if got := requests.Load(); got != 1 {
				t.Fatalf("request count = %d, want 1", got)
			}
		})
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwareRejectsMismatched200(t *testing.T) {
	body := `{
  "id": "00000000-0000-4000-8000-000000000002",
  "order_id": "00000000-0000-4000-8000-000000000003",
  "settlement_job_id": "00000000-0000-4000-8000-000000000099",
  "receipt": {"receipt_version": 2},
  "receipt_digest": "sha256:receipt",
  "signature": "signature",
  "signing_key_version": "v1",
  "eligible_alternatives": [],
  "fee_policy": null,
  "fee_evidence": null,
  "fee_quote_digest": null,
  "fee_quote_expires_at": null,
  "settlement_amount_atomic": "25000000",
  "gas_mode": "buyer_pays",
  "buyer_native_fee_atomic": "0",
  "sponsored_native_fee_atomic": null,
  "sponsored_native_symbol": null,
  "tenant_gas_charge_micros": null,
  "gas_sponsorship_evidence_digest": null,
  "created_at": "2026-09-04T00:00:01Z"
}`
	client, requests := newReceiptTestClient(t, http.StatusOK, body, "")

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if result != nil {
		t.Fatalf("ExecuteStatusAware() result = %#v, want nil", result)
	}
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) || !strings.Contains(err.Error(), "settlement_job_id") {
		t.Fatalf("ExecuteStatusAware() error = %T %v, want ID mismatch error", err, err)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwareRejectsIncomplete200(t *testing.T) {
	client, requests := newReceiptTestClient(t, http.StatusOK, `{}`, "")

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if result != nil {
		t.Fatalf("ExecuteStatusAware() result = %#v, want nil", result)
	}
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("ExecuteStatusAware() error = %T %v, want *GenericOpenAPIError", err, err)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func TestPaymentsRetrieveReceiptExecuteStatusAwarePreservesAPIError(t *testing.T) {
	body := `{"error":{"code":"receipt_temporarily_unavailable","detail":"retry later"}}`
	client, requests := newReceiptTestClient(t, http.StatusServiceUnavailable, body, "2")

	result, response, err := client.OrdersAndPaymentsAPI.
		PaymentsRetrieveReceipt(receiptTestContext(), receiptTestPaymentID).
		ExecuteStatusAware()
	closeResponse(t, response)

	if result != nil {
		t.Fatalf("ExecuteStatusAware() result = %#v, want nil", result)
	}
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("ExecuteStatusAware() error = %T %v, want *GenericOpenAPIError", err, err)
	}
	if response == nil || response.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("ExecuteStatusAware() response = %#v, want HTTP 503", response)
	}
	if got := requests.Load(); got != 1 {
		t.Fatalf("request count = %d, want 1", got)
	}
}

func newReceiptTestClient(
	t *testing.T,
	statusCode int,
	body string,
	retryAfter string,
) (*APIClient, *atomic.Int32) {
	t.Helper()
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		requests.Add(1)
		if request.Method != http.MethodGet {
			t.Errorf("request method = %s, want GET", request.Method)
		}
		if request.URL.Path != "/v1/payments/"+receiptTestPaymentID+"/receipt" {
			t.Errorf("request path = %q, want receipt path", request.URL.Path)
		}
		if authorization := request.Header.Get("Authorization"); authorization != "Bearer test-token" {
			t.Errorf("Authorization = %q, want bearer token", authorization)
		}
		response.Header().Set("Content-Type", "application/json")
		if retryAfter != "" {
			response.Header().Set("Retry-After", retryAfter)
		}
		response.WriteHeader(statusCode)
		if _, err := io.WriteString(response, body); err != nil {
			t.Errorf("write response: %v", err)
		}
	}))
	t.Cleanup(server.Close)

	configuration := NewConfiguration()
	configuration.Servers = ServerConfigurations{{URL: server.URL}}
	configuration.HTTPClient = server.Client()
	return NewAPIClient(configuration), &requests
}

func receiptTestContext() context.Context {
	return context.WithValue(context.Background(), ContextAccessToken, "test-token")
}

func validPendingFinalityBody(paymentID string) string {
	return fmt.Sprintf(`{
  "payment_id": %q,
  "state": "confirmed",
  "confirmed": true,
  "finalized": false,
  "confirmed_at": "2026-09-04T00:00:00Z",
  "finalized_at": null,
  "transaction": "0xcac79c1e",
  "network": "eip155:8453",
  "receipt_status": "pending_finality"
}`, paymentID)
}

func closeResponse(t *testing.T, response *http.Response) {
	t.Helper()
	if response != nil && response.Body != nil {
		if err := response.Body.Close(); err != nil {
			t.Errorf("close response: %v", err)
		}
	}
}
