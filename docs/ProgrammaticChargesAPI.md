# \ProgrammaticChargesAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChargesCreate**](ProgrammaticChargesAPI.md#ChargesCreate) | **Post** /v1/charges | Create a programmatic charge
[**ChargesRetrieve**](ProgrammaticChargesAPI.md#ChargesRetrieve) | **Get** /v1/charges/{charge_id} | Retrieve a programmatic charge



## ChargesCreate

> DynamicChargeResponse ChargesCreate(ctx).IdempotencyKey(idempotencyKey).DynamicChargeCreate(dynamicChargeCreate).Execute()

Create a programmatic charge



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
	idempotencyKey := "idempotencyKey_example" // string | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome.
	dynamicChargeCreate := *openapiclient.NewDynamicChargeCreate("ResourceVersionId_example", "ResourceUrl_example", []openapiclient.DynamicChargePrice{*openapiclient.NewDynamicChargePrice("AssetId_example", "AmountAtomic_example")}, int32(123)) // DynamicChargeCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgrammaticChargesAPI.ChargesCreate(context.Background()).IdempotencyKey(idempotencyKey).DynamicChargeCreate(dynamicChargeCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticChargesAPI.ChargesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargesCreate`: DynamicChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammaticChargesAPI.ChargesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |
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


## ChargesRetrieve

> DynamicChargeResponse ChargesRetrieve(ctx, chargeId).Execute()

Retrieve a programmatic charge



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
	resp, r, err := apiClient.ProgrammaticChargesAPI.ChargesRetrieve(context.Background(), chargeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgrammaticChargesAPI.ChargesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargesRetrieve`: DynamicChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgrammaticChargesAPI.ChargesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chargeId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiChargesRetrieveRequest struct via the builder pattern


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
