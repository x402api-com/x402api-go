# \ReceivingAddressesAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceivingAddressesActivate**](ReceivingAddressesAPI.md#ReceivingAddressesActivate) | **Post** /v1/receiving-addresses/{readiness_id}/activate | Activate a receiving address
[**ReceivingAddressesCreateControlChallenge**](ReceivingAddressesAPI.md#ReceivingAddressesCreateControlChallenge) | **Post** /v1/receiving-address-control-challenges | Create a receiving-address control challenge
[**ReceivingAddressesGetControlCapabilities**](ReceivingAddressesAPI.md#ReceivingAddressesGetControlCapabilities) | **Get** /v1/receiving-address-control-capabilities | Get receiving-address control capabilities
[**ReceivingAddressesList**](ReceivingAddressesAPI.md#ReceivingAddressesList) | **Get** /v1/receiving-addresses | List receiving addresses
[**ReceivingAddressesRefreshReadiness**](ReceivingAddressesAPI.md#ReceivingAddressesRefreshReadiness) | **Post** /v1/receiving-addresses/{readiness_id}/readiness-refreshes | Refresh receiving-address readiness
[**ReceivingAddressesRegister**](ReceivingAddressesAPI.md#ReceivingAddressesRegister) | **Post** /v1/receiving-addresses | Register a receiving address
[**ReceivingAddressesRotate**](ReceivingAddressesAPI.md#ReceivingAddressesRotate) | **Post** /v1/receiving-addresses/{readiness_id}/rotations | Rotate a receiving address



## ReceivingAddressesActivate

> ExternalReceivingAddress ReceivingAddressesActivate(ctx, readinessId).IdempotencyKey(idempotencyKey).Execute()

Activate a receiving address



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesActivate(context.Background(), readinessId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesActivate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |


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


## ReceivingAddressesCreateControlChallenge

> ExternalAddressControlChallenge ReceivingAddressesCreateControlChallenge(ctx).IdempotencyKey(idempotencyKey).ExternalAddressControlChallengeCreate(externalAddressControlChallengeCreate).Execute()

Create a receiving-address control challenge



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
	externalAddressControlChallengeCreate := *openapiclient.NewExternalAddressControlChallengeCreate("Network_example", "AssetId_example", "Address_example", openapiclient.ExternalAddressProofInputMethodEnum("signed_message")) // ExternalAddressControlChallengeCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesCreateControlChallenge(context.Background()).IdempotencyKey(idempotencyKey).ExternalAddressControlChallengeCreate(externalAddressControlChallengeCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesCreateControlChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesCreateControlChallenge`: ExternalAddressControlChallenge
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesCreateControlChallenge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesCreateControlChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |
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


## ReceivingAddressesGetControlCapabilities

> ExternalAddressControlCapabilities ReceivingAddressesGetControlCapabilities(ctx).Execute()

Get receiving-address control capabilities



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
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesGetControlCapabilities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesGetControlCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesGetControlCapabilities`: ExternalAddressControlCapabilities
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesGetControlCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesGetControlCapabilitiesRequest struct via the builder pattern


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


## ReceivingAddressesList

> []ExternalReceivingAddress ReceivingAddressesList(ctx).Cursor(cursor).PageSize(pageSize).Execute()

List receiving addresses



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
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesList`: []ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesListRequest struct via the builder pattern


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


## ReceivingAddressesRefreshReadiness

> ExternalReceivingAddress ReceivingAddressesRefreshReadiness(ctx, readinessId).IdempotencyKey(idempotencyKey).Execute()

Refresh receiving-address readiness



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesRefreshReadiness(context.Background(), readinessId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesRefreshReadiness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesRefreshReadiness`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesRefreshReadiness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesRefreshReadinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |


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


## ReceivingAddressesRegister

> ExternalReceivingAddress ReceivingAddressesRegister(ctx).IdempotencyKey(idempotencyKey).ExternalReceivingAddressCreate(externalReceivingAddressCreate).Execute()

Register a receiving address



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
	externalReceivingAddressCreate := *openapiclient.NewExternalReceivingAddressCreate("Label_example", "ChallengeId_example", *openapiclient.NewExternalAddressControlProofInput(openapiclient.ExternalAddressProofInputMethodEnum("signed_message"))) // ExternalReceivingAddressCreate |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesRegister(context.Background()).IdempotencyKey(idempotencyKey).ExternalReceivingAddressCreate(externalReceivingAddressCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesRegister`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |
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


## ReceivingAddressesRotate

> ExternalReceivingAddress ReceivingAddressesRotate(ctx, readinessId).IdempotencyKey(idempotencyKey).ExternalReceivingAddressRotation(externalReceivingAddressRotation).Execute()

Rotate a receiving address



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
	readinessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	externalReceivingAddressRotation := *openapiclient.NewExternalReceivingAddressRotation("ChallengeId_example", *openapiclient.NewExternalAddressControlProofInput(openapiclient.ExternalAddressProofInputMethodEnum("signed_message")), "Reason_example") // ExternalReceivingAddressRotation |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceivingAddressesAPI.ReceivingAddressesRotate(context.Background(), readinessId).IdempotencyKey(idempotencyKey).ExternalReceivingAddressRotation(externalReceivingAddressRotation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceivingAddressesAPI.ReceivingAddressesRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceivingAddressesRotate`: ExternalReceivingAddress
	fmt.Fprintf(os.Stdout, "Response from `ReceivingAddressesAPI.ReceivingAddressesRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**readinessId** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiReceivingAddressesRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Caller-persisted mutation key containing 8 to 160 safe ASCII characters. Replay the exact key and body after an uncertain outcome. |

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
