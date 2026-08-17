# \FacilitatorDiscoveryAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FacilitatorSupportedRetrieve**](FacilitatorDiscoveryAPI.md#FacilitatorSupportedRetrieve) | **Get** /v1/facilitator/supported |



## FacilitatorSupportedRetrieve

> SupportedResponse FacilitatorSupportedRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.FacilitatorDiscoveryAPI.FacilitatorSupportedRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitatorDiscoveryAPI.FacilitatorSupportedRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilitatorSupportedRetrieve`: SupportedResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitatorDiscoveryAPI.FacilitatorSupportedRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFacilitatorSupportedRetrieveRequest struct via the builder pattern


### Return type

[**SupportedResponse**](SupportedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
