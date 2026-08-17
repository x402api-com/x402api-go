# \AssetsAndPaymentControlsAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PaymentReadinessRetrieve**](AssetsAndPaymentControlsAPI.md#V1PaymentReadinessRetrieve) | **Get** /v1/payment-readiness |



## V1PaymentReadinessRetrieve

> PaymentReadiness V1PaymentReadinessRetrieve(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAndPaymentControlsAPI.V1PaymentReadinessRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAndPaymentControlsAPI.V1PaymentReadinessRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentReadinessRetrieve`: PaymentReadiness
	fmt.Fprintf(os.Stdout, "Response from `AssetsAndPaymentControlsAPI.V1PaymentReadinessRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentReadinessRetrieveRequest struct via the builder pattern


### Return type

[**PaymentReadiness**](PaymentReadiness.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
