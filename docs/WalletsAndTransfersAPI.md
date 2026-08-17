# \WalletsAndTransfersAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1WalletsBalancesRetrieve**](WalletsAndTransfersAPI.md#V1WalletsBalancesRetrieve) | **Get** /v1/wallets/{id}/balances |



## V1WalletsBalancesRetrieve

> WalletBalanceResponse V1WalletsBalancesRetrieve(ctx, id).Finality(finality).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	finality := "finality_example" // string |  (optional) (default to "finalized")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAndTransfersAPI.V1WalletsBalancesRetrieve(context.Background(), id).Finality(finality).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAndTransfersAPI.V1WalletsBalancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1WalletsBalancesRetrieve`: WalletBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAndTransfersAPI.V1WalletsBalancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1WalletsBalancesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **finality** | **string** |  | [default to &quot;finalized&quot;]

### Return type

[**WalletBalanceResponse**](WalletBalanceResponse.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
