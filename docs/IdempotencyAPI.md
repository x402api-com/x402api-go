# \IdempotencyAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1IdempotencyOutcomesRetrieve**](IdempotencyAPI.md#V1IdempotencyOutcomesRetrieve) | **Get** /v1/idempotency-outcomes/{idempotency_key} |



## V1IdempotencyOutcomesRetrieve

> IdempotencyOutcome V1IdempotencyOutcomesRetrieve(ctx, idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/x402api-com/x402api-go"
)

func main() {
	idempotencyKey := "idempotencyKey_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdempotencyAPI.V1IdempotencyOutcomesRetrieve(context.Background(), idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdempotencyAPI.V1IdempotencyOutcomesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IdempotencyOutcomesRetrieve`: IdempotencyOutcome
	fmt.Fprintf(os.Stdout, "Response from `IdempotencyAPI.V1IdempotencyOutcomesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idempotencyKey** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1IdempotencyOutcomesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdempotencyOutcome**](IdempotencyOutcome.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
