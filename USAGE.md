# Go usage guide

The [README](README.md) contains installation instructions and the complete function index. This guide focuses on safe production patterns.

## Create and reuse a client

Create one configuration and API client for a service. Pass credentials, deadlines, tracing, and cancellation through `context.Context`.

```go
cfg := x402api.NewConfiguration()
cfg.HTTPClient = &http.Client{Timeout: 15 * time.Second}
client := x402api.NewAPIClient(cfg)

ctx := context.WithValue(
    context.Background(),
    x402api.ContextAccessToken,
    os.Getenv("X402API_TENANT_API_KEY"),
)
```

Do not store per-request tokens by mutating a shared global. Derive a context for each request or tenant.

## Create and retrieve a charge

```go
request := x402api.NewDynamicChargeCreate(
    "00000000-0000-4000-8000-000000000001",
    "https://merchant.example.com/premium-report",
    []x402api.DynamicChargePrice{
        *x402api.NewDynamicChargePrice("base_usdc", "1000000"),
    },
    900,
)
request.SetMetadata(map[string]interface{}{
    "order_id": "order-123",
})

idempotencyKey := "charge-order-123-v1"
charge, response, err := client.ProgrammaticChargesAPI.
    ChargesCreate(ctx).
    IdempotencyKey(idempotencyKey).
    DynamicChargeCreate(*request).
    Execute()
if response != nil {
    defer response.Body.Close()
}
if err != nil {
    handleError(response, err)
}

sameCharge, response, err := client.ProgrammaticChargesAPI.
    ChargesRetrieve(ctx, charge.GetChargeId()).
    Execute()
```

Prices use atomic-unit strings, not floating point. For example, `"1000000"` represents one token for an asset with six decimals.

## Pagination and HTTP headers

```go
cursor := ""

for {
    request := client.OrdersAndPaymentsAPI.
        PaymentsList(ctx).
        PageSize(100)
    if cursor != "" {
        request = request.Cursor(cursor)
    }

    payments, response, err := request.Execute()
    if err != nil {
        handleError(response, err)
        break
    }

    for _, payment := range payments {
        process(payment)
    }

    cursor = response.Header.Get("X-X402API-Next-Cursor")
    response.Body.Close()
    if cursor == "" {
        break
    }
}
```

Treat the cursor as opaque and pass it back unchanged. The same pattern applies to orders, payment observations, receiving addresses, resources, and resource versions.

## Error handling

```go
func handleError(response *http.Response, err error) {
    if response != nil {
        requestID := response.Header.Get("X-Request-ID")
        retryAfter := response.Header.Get("Retry-After")
        log.Printf(
            "x402api status=%d request_id=%s retry_after=%s",
            response.StatusCode,
            requestID,
            retryAfter,
        )
    }

    var apiError *x402api.GenericOpenAPIError
    if errors.As(err, &apiError) {
        log.Printf("x402api response body: %s", apiError.Body())
    }

    log.Printf("x402api request failed: %v", err)
}
```

Always inspect both `error` and `*http.Response`; an error can occur before a response exists. Close non-nil response bodies.

## Idempotency and retries

Mutations require keys of 8-160 characters matching `[A-Za-z0-9._:-]+`. Persist the key with the intent you are executing.

- New intended mutation: generate a new key.
- Timeout or connection reset after sending: retry the identical body with the same key.
- Known validation failure: fix the request and use a new key.
- Uncertain durable outcome: call `client.IdempotencyAPI.IdempotencyGetOutcome(ctx, key).Execute()`.

The SDK does not retry automatically. Bound application retries, use exponential backoff with jitter, respect `Retry-After`, and normally retry only connection failures plus HTTP `408`, `429`, `500`, `502`, `503`, and `504`.

## Public endpoints

These endpoints do not need a tenant key:

```go
publicContext := context.Background()
supported, response, err := client.FacilitatorDiscoveryAPI.
    FacilitatorGetSupported(publicContext).
    Execute()

keys, response, err := client.OrdersAndPaymentsAPI.
    ReceiptVerificationKeysRetrieve(publicContext).
    Execute()
```

Do not edit generated `api_*.go`, `model_*.go`, or `docs/` files; update the OpenAPI contract or generator configuration instead.
