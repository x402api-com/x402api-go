# \ReceivingAddressesAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReceivingAddressControlCapabilitiesRetrieve**](ReceivingAddressesAPI.md#V1ReceivingAddressControlCapabilitiesRetrieve) | **Get** /v1/receiving-address-control-capabilities |
[**V1ReceivingAddressControlChallengesCreate**](ReceivingAddressesAPI.md#V1ReceivingAddressControlChallengesCreate) | **Post** /v1/receiving-address-control-challenges |
[**V1ReceivingAddressesActivateCreate**](ReceivingAddressesAPI.md#V1ReceivingAddressesActivateCreate) | **Post** /v1/receiving-addresses/{readiness_id}/activate |
[**V1ReceivingAddressesCreate**](ReceivingAddressesAPI.md#V1ReceivingAddressesCreate) | **Post** /v1/receiving-addresses |
[**V1ReceivingAddressesList**](ReceivingAddressesAPI.md#V1ReceivingAddressesList) | **Get** /v1/receiving-addresses |
[**V1ReceivingAddressesReadinessRefreshesCreate**](ReceivingAddressesAPI.md#V1ReceivingAddressesReadinessRefreshesCreate) | **Post** /v1/receiving-addresses/{readiness_id}/readiness-refreshes |
[**V1ReceivingAddressesRotationsCreate**](ReceivingAddressesAPI.md#V1ReceivingAddressesRotationsCreate) | **Post** /v1/receiving-addresses/{readiness_id}/rotations |



## V1ReceivingAddressControlCapabilitiesRetrieve

> ExternalAddressControlCapabilities V1ReceivingAddressControlCapabilitiesRetrieve(ctx).Execute()



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
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressControlCapabilitiesRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressControlCapabilitiesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressControlCapabilitiesRetrieve`: ExternalAddressControlCapabilities
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressControlCapabilitiesRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressControlCapabilitiesRetrieveRequest struct via the builder pattern


### Return type

[**ExternalAddressControlCapabilities**](ExternalAddressControlCapabilities.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressControlChallengesCreate

> ExternalAddressControlChallenge V1ReceivingAddressControlChallengesCreate(ctx).IdempotencyKey(idempotencyKey).ExternalAddressControlChallengeCreate(externalAddressControlChallengeCreate).Execute()



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
	externalAddressControlChallengeCreate := *openapiclient.NewExternalAddressControlChallengeCreate("Network_example", "AssetId_example", "Address_example", openapiclient.ExternalAddressProofInputMethodEnum("signed_message")) // ExternalAddressControlChallengeCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressControlChallengesCreate(context.Background()).IdempotencyKey(idempotencyKey).ExternalAddressControlChallengeCreate(externalAddressControlChallengeCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressControlChallengesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressControlChallengesCreate`: ExternalAddressControlChallenge
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressControlChallengesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressControlChallengesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |
 **externalAddressControlChallengeCreate** | [**ExternalAddressControlChallengeCreate**](ExternalAddressControlChallengeCreate.md) |  |

### Return type

[**ExternalAddressControlChallenge**](ExternalAddressControlChallenge.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressesActivateCreate

> ExternalReceivingAddress V1ReceivingAddressesActivateCreate(ctx, readinessId).IdempotencyKey(idempotencyKey).Execute()



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressesActivateCreate(context.Background(), readinessId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressesActivateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressesActivateCreate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressesActivateCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressesActivateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |


### Return type

[**ExternalReceivingAddress**](ExternalReceivingAddress.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressesCreate

> ExternalReceivingAddress V1ReceivingAddressesCreate(ctx).IdempotencyKey(idempotencyKey).ExternalReceivingAddressCreate(externalReceivingAddressCreate).Execute()



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
	externalReceivingAddressCreate := *openapiclient.NewExternalReceivingAddressCreate("Label_example", "ChallengeId_example", *openapiclient.NewExternalAddressControlProofInput(openapiclient.ExternalAddressProofInputMethodEnum("signed_message"))) // ExternalReceivingAddressCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressesCreate(context.Background()).IdempotencyKey(idempotencyKey).ExternalReceivingAddressCreate(externalReceivingAddressCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressesCreate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |
 **externalReceivingAddressCreate** | [**ExternalReceivingAddressCreate**](ExternalReceivingAddressCreate.md) |  |

### Return type

[**ExternalReceivingAddress**](ExternalReceivingAddress.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressesList

> []ExternalReceivingAddress V1ReceivingAddressesList(ctx).Cursor(cursor).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressesList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressesList`: []ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]ExternalReceivingAddress**](ExternalReceivingAddress.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressesReadinessRefreshesCreate

> ExternalReceivingAddress V1ReceivingAddressesReadinessRefreshesCreate(ctx, readinessId).IdempotencyKey(idempotencyKey).Execute()



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressesReadinessRefreshesCreate(context.Background(), readinessId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressesReadinessRefreshesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressesReadinessRefreshesCreate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressesReadinessRefreshesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressesReadinessRefreshesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |


### Return type

[**ExternalReceivingAddress**](ExternalReceivingAddress.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReceivingAddressesRotationsCreate

> ExternalReceivingAddress V1ReceivingAddressesRotationsCreate(ctx, readinessId).IdempotencyKey(idempotencyKey).ExternalReceivingAddressRotation(externalReceivingAddressRotation).Execute()



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	externalReceivingAddressRotation := *openapiclient.NewExternalReceivingAddressRotation("ChallengeId_example", *openapiclient.NewExternalAddressControlProofInput(openapiclient.ExternalAddressProofInputMethodEnum("signed_message")), "Reason_example") // ExternalReceivingAddressRotation |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.V1ReceivingAddressesRotationsCreate(context.Background(), readinessId).IdempotencyKey(idempotencyKey).ExternalReceivingAddressRotation(externalReceivingAddressRotation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.V1ReceivingAddressesRotationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReceivingAddressesRotationsCreate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.V1ReceivingAddressesRotationsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReceivingAddressesRotationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Unique mutation key; replaying different content returns HTTP 409. |

 **externalReceivingAddressRotation** | [**ExternalReceivingAddressRotation**](ExternalReceivingAddressRotation.md) |  |

### Return type

[**ExternalReceivingAddress**](ExternalReceivingAddress.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
