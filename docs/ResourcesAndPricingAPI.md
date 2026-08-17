# \ResourcesAndPricingAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NetworkFeeQuotesCreate**](ResourcesAndPricingAPI.md#V1NetworkFeeQuotesCreate) | **Post** /v1/network-fee-quotes |
[**V1ResourcesCreate**](ResourcesAndPricingAPI.md#V1ResourcesCreate) | **Post** /v1/resources |
[**V1ResourcesList**](ResourcesAndPricingAPI.md#V1ResourcesList) | **Get** /v1/resources |
[**V1ResourcesVersionsActivateCreate**](ResourcesAndPricingAPI.md#V1ResourcesVersionsActivateCreate) | **Post** /v1/resources/{resource_id}/versions/{version_id}/activate |
[**V1ResourcesVersionsCreate**](ResourcesAndPricingAPI.md#V1ResourcesVersionsCreate) | **Post** /v1/resources/{resource_id}/versions |
[**V1ResourcesVersionsList**](ResourcesAndPricingAPI.md#V1ResourcesVersionsList) | **Get** /v1/resources/{resource_id}/versions |
[**V1ResourcesVersionsRetireCreate**](ResourcesAndPricingAPI.md#V1ResourcesVersionsRetireCreate) | **Post** /v1/resources/{resource_id}/versions/{version_id}/retire |



## V1NetworkFeeQuotesCreate

> NetworkFeePreviewResponse V1NetworkFeeQuotesCreate(ctx).NetworkFeePreview(networkFeePreview).Execute()



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
	networkFeePreview := *openapiclient.NewNetworkFeePreview([]openapiclient.NetworkFeePreviewPrice{*openapiclient.NewNetworkFeePreviewPrice("AssetId_example", "ListedAmountAtomic_example")}) // NetworkFeePreview |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1NetworkFeeQuotesCreate(context.Background()).NetworkFeePreview(networkFeePreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1NetworkFeeQuotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1NetworkFeeQuotesCreate`: NetworkFeePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1NetworkFeeQuotesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1NetworkFeeQuotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkFeePreview** | [**NetworkFeePreview**](NetworkFeePreview.md) |  |

### Return type

[**NetworkFeePreviewResponse**](NetworkFeePreviewResponse.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesCreate

> Resource V1ResourcesCreate(ctx).IdempotencyKey(idempotencyKey).ResourceCreate(resourceCreate).Execute()



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
	resourceCreate := *openapiclient.NewResourceCreate("Key_example", "Name_example", openapiclient.HTTPMethodEnum("GET"), "Path_example", "Description_example", openapiclient.ResourceInputFulfillmentModeEnum("webhook"), []openapiclient.PriceInput{*openapiclient.NewPriceInput("AssetId_example", "WalletVersionId_example", "AmountAtomic_example", int32(123))}) // ResourceCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesCreate(context.Background()).IdempotencyKey(idempotencyKey).ResourceCreate(resourceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesCreate`: Resource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |
 **resourceCreate** | [**ResourceCreate**](ResourceCreate.md) |  |

### Return type

[**Resource**](Resource.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesList

> []Resource V1ResourcesList(ctx).Cursor(cursor).PageSize(pageSize).Execute()



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
	cursor := "cursor_example" // string | Opaque pagination cursor from X-X402API-Next-Cursor or rel=next Link. (optional)
	pageSize := int32(56) // int32 | Number of results in the bounded array page (default and maximum 100). (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesList`: []Resource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]Resource**](Resource.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesVersionsActivateCreate

> ResourceVersion V1ResourcesVersionsActivateCreate(ctx, resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionActivate(resourceVersionActivate).Execute()



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionActivate := *openapiclient.NewResourceVersionActivate(int32(123), "ExpectedActiveVersionId_example") // ResourceVersionActivate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesVersionsActivateCreate(context.Background(), resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionActivate(resourceVersionActivate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesVersionsActivateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesVersionsActivateCreate`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesVersionsActivateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |
**versionId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesVersionsActivateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |


 **resourceVersionActivate** | [**ResourceVersionActivate**](ResourceVersionActivate.md) |  |

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesVersionsCreate

> ResourceVersion V1ResourcesVersionsCreate(ctx, resourceId).IdempotencyKey(idempotencyKey).ResourceVersionCreate(resourceVersionCreate).Execute()



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionCreate := *openapiclient.NewResourceVersionCreate(int32(123), openapiclient.HTTPMethodEnum("GET"), "Path_example", "Description_example", openapiclient.ResourceInputFulfillmentModeEnum("webhook"), []openapiclient.PriceInput{*openapiclient.NewPriceInput("AssetId_example", "WalletVersionId_example", "AmountAtomic_example", int32(123))}) // ResourceVersionCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesVersionsCreate(context.Background(), resourceId).IdempotencyKey(idempotencyKey).ResourceVersionCreate(resourceVersionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesVersionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesVersionsCreate`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesVersionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesVersionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |

 **resourceVersionCreate** | [**ResourceVersionCreate**](ResourceVersionCreate.md) |  |

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesVersionsList

> []ResourceVersion V1ResourcesVersionsList(ctx, resourceId).Cursor(cursor).PageSize(pageSize).Execute()



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	cursor := "cursor_example" // string | Opaque pagination cursor from X-X402API-Next-Cursor or rel=next Link. (optional)
	pageSize := int32(56) // int32 | Number of results in the bounded array page (default and maximum 100). (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesVersionsList(context.Background(), resourceId).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesVersionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesVersionsList`: []ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesVersionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesVersionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]ResourceVersion**](ResourceVersion.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ResourcesVersionsRetireCreate

> ResourceVersion V1ResourcesVersionsRetireCreate(ctx, resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionRetire(resourceVersionRetire).Execute()



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionRetire := *openapiclient.NewResourceVersionRetire(int32(123), openapiclient.ResourceVersionRetireExpectedStateEnum("draft")) // ResourceVersionRetire |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.V1ResourcesVersionsRetireCreate(context.Background(), resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionRetire(resourceVersionRetire).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.V1ResourcesVersionsRetireCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ResourcesVersionsRetireCreate`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.V1ResourcesVersionsRetireCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |
**versionId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ResourcesVersionsRetireCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |


 **resourceVersionRetire** | [**ResourceVersionRetire**](ResourceVersionRetire.md) |  |

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
