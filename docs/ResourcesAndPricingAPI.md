# \ResourcesAndPricingAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkFeesCreateQuote**](ResourcesAndPricingAPI.md#NetworkFeesCreateQuote) | **Post** /v1/network-fee-quotes | Create a network-fee quote
[**ResourcesActivateVersion**](ResourcesAndPricingAPI.md#ResourcesActivateVersion) | **Post** /v1/resources/{resource_id}/versions/{version_id}/activate | Activate a resource version
[**ResourcesCreate**](ResourcesAndPricingAPI.md#ResourcesCreate) | **Post** /v1/resources | Create a resource
[**ResourcesCreateVersion**](ResourcesAndPricingAPI.md#ResourcesCreateVersion) | **Post** /v1/resources/{resource_id}/versions | Create a resource version
[**ResourcesList**](ResourcesAndPricingAPI.md#ResourcesList) | **Get** /v1/resources | List resources
[**ResourcesListVersions**](ResourcesAndPricingAPI.md#ResourcesListVersions) | **Get** /v1/resources/{resource_id}/versions | List resource versions
[**ResourcesRetireVersion**](ResourcesAndPricingAPI.md#ResourcesRetireVersion) | **Post** /v1/resources/{resource_id}/versions/{version_id}/retire | Retire a resource version



## NetworkFeesCreateQuote

> NetworkFeePreviewResponse NetworkFeesCreateQuote(ctx).NetworkFeePreview(networkFeePreview).Execute()

Create a network-fee quote



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
	resp, r, err := apiClient.ResourcesAndPricingAPI.NetworkFeesCreateQuote(context.Background()).NetworkFeePreview(networkFeePreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.NetworkFeesCreateQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkFeesCreateQuote`: NetworkFeePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.NetworkFeesCreateQuote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkFeesCreateQuoteRequest struct via the builder pattern


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


## ResourcesActivateVersion

> ResourceVersion ResourcesActivateVersion(ctx, resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionActivate(resourceVersionActivate).Execute()

Activate a resource version



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionActivate := *openapiclient.NewResourceVersionActivate(int32(123), "ExpectedActiveVersionId_example") // ResourceVersionActivate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesActivateVersion(context.Background(), resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionActivate(resourceVersionActivate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesActivateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesActivateVersion`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesActivateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |
**versionId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesActivateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |


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


## ResourcesCreate

> Resource ResourcesCreate(ctx).IdempotencyKey(idempotencyKey).ResourceCreate(resourceCreate).Execute()

Create a resource



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
	resourceCreate := *openapiclient.NewResourceCreate("Key_example", "Name_example", openapiclient.HTTPMethodEnum("GET"), "Path_example", "Description_example", openapiclient.ResourceInputFulfillmentModeEnum("webhook"), []openapiclient.PriceInput{*openapiclient.NewPriceInput("AssetId_example", "WalletVersionId_example", "AmountAtomic_example", int32(123))}) // ResourceCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesCreate(context.Background()).IdempotencyKey(idempotencyKey).ResourceCreate(resourceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesCreate`: Resource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourcesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |
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


## ResourcesCreateVersion

> ResourceVersion ResourcesCreateVersion(ctx, resourceId).IdempotencyKey(idempotencyKey).ResourceVersionCreate(resourceVersionCreate).Execute()

Create a resource version



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionCreate := *openapiclient.NewResourceVersionCreate(int32(123), openapiclient.HTTPMethodEnum("GET"), "Path_example", "Description_example", openapiclient.ResourceInputFulfillmentModeEnum("webhook"), []openapiclient.PriceInput{*openapiclient.NewPriceInput("AssetId_example", "WalletVersionId_example", "AmountAtomic_example", int32(123))}) // ResourceVersionCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesCreateVersion(context.Background(), resourceId).IdempotencyKey(idempotencyKey).ResourceVersionCreate(resourceVersionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesCreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesCreateVersion`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesCreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |

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


## ResourcesList

> []Resource ResourcesList(ctx).Cursor(cursor).PageSize(pageSize).Execute()

List resources



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
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesList`: []Resource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourcesListRequest struct via the builder pattern


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


## ResourcesListVersions

> []ResourceVersion ResourcesListVersions(ctx, resourceId).Cursor(cursor).PageSize(pageSize).Execute()

List resource versions



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
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesListVersions(context.Background(), resourceId).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesListVersions`: []ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesListVersionsRequest struct via the builder pattern


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


## ResourcesRetireVersion

> ResourceVersion ResourcesRetireVersion(ctx, resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionRetire(resourceVersionRetire).Execute()

Retire a resource version



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
	resourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	versionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	resourceVersionRetire := *openapiclient.NewResourceVersionRetire(int32(123), openapiclient.ResourceVersionRetireExpectedStateEnum("draft")) // ResourceVersionRetire |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAndPricingAPI.ResourcesRetireVersion(context.Background(), resourceId, versionId).IdempotencyKey(idempotencyKey).ResourceVersionRetire(resourceVersionRetire).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAndPricingAPI.ResourcesRetireVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesRetireVersion`: ResourceVersion
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAndPricingAPI.ResourcesRetireVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  |
**versionId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesRetireVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |


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
