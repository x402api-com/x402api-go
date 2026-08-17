# x402api Go SDK

Official server-side Go client for the [x402api public API](https://api.x402api.com/openapi/openapi.json). It provides typed request and response models for programmatic x402 charges, resources, receiving addresses, payments, receipts, and wallet balances.

The module is `github.com/x402api-com/x402api-go`, targets Go 1.23+, and uses the standard `net/http` client. The production base URL is `https://api.x402api.com`.

## Installation

```bash
go get github.com/x402api-com/x402api-go@latest
```

## Authentication

Create a scoped tenant API key and provide it as a bearer token through the request context. Keep it in a server-side secret store; do not ship tenant credentials in browser, mobile, or desktop applications.

```go
ctx := context.WithValue(
    context.Background(),
    x402api.ContextAccessToken,
    os.Getenv("X402API_TENANT_API_KEY"),
)

cfg := x402api.NewConfiguration()
client := x402api.NewAPIClient(cfg)
```

`FacilitatorGetSupported` and `ReceiptVerificationKeysRetrieve` are public and may use `context.Background()` without a token. All other operations use tenant bearer authentication.

## Quick start: create a charge

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    x402api "github.com/x402api-com/x402api-go"
)

func main() {
    cfg := x402api.NewConfiguration()
    client := x402api.NewAPIClient(cfg)
    ctx := context.WithValue(
        context.Background(),
        x402api.ContextAccessToken,
        os.Getenv("X402API_TENANT_API_KEY"),
    )

    request := x402api.NewDynamicChargeCreate(
        "00000000-0000-4000-8000-000000000001",
        "https://merchant.example.com/premium-report",
        []x402api.DynamicChargePrice{
            *x402api.NewDynamicChargePrice("base_usdc", "1000000"),
        },
        900,
    )

    charge, response, err := client.ProgrammaticChargesAPI.
        ChargesCreate(ctx).
        IdempotencyKey("charge-example-001").
        DynamicChargeCreate(*request).
        Execute()
    if err != nil {
        if response != nil {
            log.Printf("x402api status: %s", response.Status)
        }
        log.Fatal(err)
    }

    fmt.Printf("charge: %+v\n", charge)
}
```

`IdempotencyKey` is required for mutations. Use a new key for each intended mutation. If the outcome is uncertain, retry the identical payload with the same key.

## Response metadata and pagination

Every `Execute()` call returns the decoded value, `*http.Response`, and `error`. Read pagination headers from the HTTP response and always close its body:

```go
payments, response, err := client.OrdersAndPaymentsAPI.
    PaymentsList(ctx).
    PageSize(25).
    Execute()
if response != nil {
    defer response.Body.Close()
}
if err != nil {
    log.Fatal(err)
}

for _, payment := range payments {
    fmt.Printf("%+v\n", payment)
}

nextCursor := response.Header.Get("X-X402API-Next-Cursor")
if nextCursor != "" {
    nextPage, nextResponse, err := client.OrdersAndPaymentsAPI.
        PaymentsList(ctx).
        Cursor(nextCursor).
        PageSize(25).
        Execute()
    _ = nextPage
    _ = nextResponse
    _ = err
}
```

Cursors are opaque. Pass them back unchanged; do not decode or construct them. Configure timeouts by assigning an `http.Client` to `cfg.HTTPClient` before calling `NewAPIClient`.

The client does not retry automatically. For connection failures and HTTP `408`, `429`, `500`, `502`, `503`, or `504`, add bounded exponential backoff in your application. Respect `Retry-After`, and preserve the same idempotency key and body when retrying a mutation. HTTP failures are returned as errors, commonly `*x402api.GenericOpenAPIError`; the accompanying `*http.Response` contains status and headers.

## API services and functions

Methods with optional query or body values use a request builder: call the service method, set optional values such as `.Cursor(...)` or `.PageSize(...)`, then call `.Execute()`. Links lead to generated parameter and model documentation.

| API service | Function | HTTP endpoint |
| --- | --- | --- |
| [`ProgrammaticChargesAPI`](docs/ProgrammaticChargesAPI.md) | `ChargesCreate(ctx).IdempotencyKey(...).DynamicChargeCreate(...).Execute()` | `POST /v1/charges` |
| [`ProgrammaticChargesAPI`](docs/ProgrammaticChargesAPI.md) | `ChargesRetrieve(ctx, chargeID).Execute()` | `GET /v1/charges/{charge_id}` |
| [`FacilitatorDiscoveryAPI`](docs/FacilitatorDiscoveryAPI.md) | `FacilitatorGetSupported(ctx).Execute()` | `GET /v1/facilitator/supported` |
| [`IdempotencyAPI`](docs/IdempotencyAPI.md) | `IdempotencyGetOutcome(ctx, idempotencyKey).Execute()` | `GET /v1/idempotency-outcomes/{idempotency_key}` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `NetworkFeesCreateQuote(ctx).NetworkFeePreview(...).Execute()` | `POST /v1/network-fee-quotes` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `OrdersList(ctx).Cursor(...).PageSize(...).Execute()` | `GET /v1/orders` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `OrdersRetrieve(ctx, id).Execute()` | `GET /v1/orders/{id}` |
| [`AssetsAndPaymentControlsAPI`](docs/AssetsAndPaymentControlsAPI.md) | `PaymentReadinessRetrieve(ctx).Execute()` | `GET /v1/payment-readiness` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `PaymentsList(ctx).Cursor(...).PageSize(...).Execute()` | `GET /v1/payments` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `PaymentsRetrieve(ctx, id).Execute()` | `GET /v1/payments/{id}` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `PaymentsListObservations(ctx, id).Cursor(...).PageSize(...).Execute()` | `GET /v1/payments/{id}/observations` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `PaymentsRetrieveReceipt(ctx, id).Execute()` | `GET /v1/payments/{id}/receipt` |
| [`OrdersAndPaymentsAPI`](docs/OrdersAndPaymentsAPI.md) | `ReceiptVerificationKeysRetrieve(ctx).Execute()` | `GET /v1/payment-receipt-verification-keys` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesGetControlCapabilities(ctx).Execute()` | `GET /v1/receiving-address-control-capabilities` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesCreateControlChallenge(ctx).IdempotencyKey(...).ExternalAddressControlChallengeCreate(...).Execute()` | `POST /v1/receiving-address-control-challenges` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesList(ctx).Cursor(...).PageSize(...).Execute()` | `GET /v1/receiving-addresses` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesRegister(ctx).IdempotencyKey(...).ExternalReceivingAddressCreate(...).Execute()` | `POST /v1/receiving-addresses` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesActivate(ctx, readinessID).IdempotencyKey(...).Execute()` | `POST /v1/receiving-addresses/{readiness_id}/activate` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesRefreshReadiness(ctx, readinessID).IdempotencyKey(...).Execute()` | `POST /v1/receiving-addresses/{readiness_id}/readiness-refreshes` |
| [`ReceivingAddressesAPI`](docs/ReceivingAddressesAPI.md) | `ReceivingAddressesRotate(ctx, readinessID).IdempotencyKey(...).ExternalReceivingAddressRotation(...).Execute()` | `POST /v1/receiving-addresses/{readiness_id}/rotations` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesList(ctx).Cursor(...).PageSize(...).Execute()` | `GET /v1/resources` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesCreate(ctx).IdempotencyKey(...).ResourceCreate(...).Execute()` | `POST /v1/resources` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesListVersions(ctx, resourceID).Cursor(...).PageSize(...).Execute()` | `GET /v1/resources/{resource_id}/versions` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesCreateVersion(ctx, resourceID).IdempotencyKey(...).ResourceVersionCreate(...).Execute()` | `POST /v1/resources/{resource_id}/versions` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesActivateVersion(ctx, resourceID, versionID).IdempotencyKey(...).ResourceVersionActivate(...).Execute()` | `POST /v1/resources/{resource_id}/versions/{version_id}/activate` |
| [`ResourcesAndPricingAPI`](docs/ResourcesAndPricingAPI.md) | `ResourcesRetireVersion(ctx, resourceID, versionID).IdempotencyKey(...).ResourceVersionRetire(...).Execute()` | `POST /v1/resources/{resource_id}/versions/{version_id}/retire` |
| [`WalletsAndTransfersAPI`](docs/WalletsAndTransfersAPI.md) | `WalletsRetrieveBalance(ctx, id).Finality(...).Execute()` | `GET /v1/wallets/{id}/balances` |

All request and response model documentation is in [`docs/`](docs/). See [`USAGE.md`](USAGE.md) for more complete patterns.

## Automatic generation

This repository uses OpenAPI Generator 7.24.0, pinned by Docker image and digest in [`scripts/generate-sdk.sh`](scripts/generate-sdk.sh). The [`SDK generation workflow`](.github/workflows/sdk_generation.yaml) checks the live OpenAPI document hourly and on manual or repository dispatch. When its normalized contract changes, GitHub Actions regenerates, validates, and commits the SDK to `main`.

To regenerate locally with Docker:

```bash
./scripts/generate-sdk.sh
go test ./...
```

Persistent files such as this README, `USAGE.md`, workflow configuration, and generator scripts are protected by [`.openapi-generator-ignore`](.openapi-generator-ignore). Generated client and model files should not be edited manually.

Licensed under the [MIT License](LICENSE).
