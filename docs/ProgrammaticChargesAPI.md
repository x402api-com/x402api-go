# \ProgrammaticChargesAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicCharge**](ProgrammaticChargesAPI.md#CreateDynamicCharge) | **Post** /v1/charges |
[**RetrieveDynamicCharge**](ProgrammaticChargesAPI.md#RetrieveDynamicCharge) | **Get** /v1/charges/{charge_id} |



## CreateDynamicCharge

> DynamicChargeResponse CreateDynamicCharge(ctx).IdempotencyKey(idempotencyKey).DynamicChargeCreate(dynamicChargeCreate).Execute()





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
	idempotencyKey := "idempotencyKey_example" // string | Unique mutation key; replaying different content returns HTTP 409.
	dynamicChargeCreate := *openapiclient.NewDynamicChargeCreate("ResourceVersionId_example", "ResourceUrl_example", []openapiclient.DynamicChargePrice{*openapiclient.NewDynamicChargePrice("AssetId_example", "AmountAtomic_example")}, int32(123)) // DynamicChargeCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammaticChargesAPI.CreateDynamicCharge(context.Background()).IdempotencyKey(idempotencyKey).DynamicChargeCreate(dynamicChargeCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticChargesAPI.CreateDynamicCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicCharge`: DynamicChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammaticChargesAPI.CreateDynamicCharge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |
 **dynamicChargeCreate** | [**DynamicChargeCreate**](DynamicChargeCreate.md) |  |

### Return type

[**DynamicChargeResponse**](DynamicChargeResponse.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDynamicCharge

> DynamicChargeResponse RetrieveDynamicCharge(ctx, chargeId).Execute()





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
	chargeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammaticChargesAPI.RetrieveDynamicCharge(context.Background(), chargeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticChargesAPI.RetrieveDynamicCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDynamicCharge`: DynamicChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammaticChargesAPI.RetrieveDynamicCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargeId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDynamicChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicChargeResponse**](DynamicChargeResponse.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
