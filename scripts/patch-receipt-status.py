#!/usr/bin/env python3
"""Patch the generated Go receipt operation to type HTTP 202 safely."""

from __future__ import annotations

from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
API_FILE = ROOT / "api_orders_and_payments.go"

FUNCTION_START = (
    "func (a *OrdersAndPaymentsAPIService) PaymentsRetrieveReceiptExecute("
    "r ApiPaymentsRetrieveReceiptRequest) (*PaymentReceipt, *http.Response, error) {"
)
FUNCTION_END = "\ntype ApiReceiptVerificationKeysRetrieveRequest struct {"
ERROR_BRANCH = "\n\tif localVarHTTPResponse.StatusCode >= 300 {\n"
PATCH_MARKER = "// x402api: decode HTTP 202 as PaymentReceiptStatus, never PaymentReceipt."
PATCH = f"""
\t{PATCH_MARKER}
\tif localVarHTTPResponse.StatusCode == http.StatusAccepted {{
\t\tpendingErr, decodeErr := newPaymentReceiptPendingError(
\t\t\ta.client,
\t\t\tr.id,
\t\t\tlocalVarHTTPResponse,
\t\t\tlocalVarBody,
\t\t)
\t\tif decodeErr != nil {{
\t\t\treturn localVarReturnValue, localVarHTTPResponse, decodeErr
\t\t}}
\t\treturn localVarReturnValue, localVarHTTPResponse, pendingErr
\t}}
"""


def main() -> None:
    source = API_FILE.read_text(encoding="utf-8")
    if source.count(FUNCTION_START) != 1 or source.count(FUNCTION_END) != 1:
        raise SystemExit("generated receipt function markers changed; refusing to patch")

    start = source.index(FUNCTION_START)
    end = source.index(FUNCTION_END, start)
    function = source[start:end]

    marker_count = function.count(PATCH_MARKER)
    if marker_count:
        if marker_count != 1 or function.count(PATCH) != 1:
            raise SystemExit("generated receipt patch is incomplete or duplicated")
        return

    if function.count(ERROR_BRANCH) != 1:
        raise SystemExit("generated receipt error branch changed; refusing to patch")

    patched_function = function.replace(ERROR_BRANCH, PATCH + ERROR_BRANCH, 1)
    API_FILE.write_text(source[:start] + patched_function + source[end:], encoding="utf-8")


if __name__ == "__main__":
    main()
