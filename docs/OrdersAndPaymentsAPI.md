# \OrdersAndPaymentsAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrdersList**](OrdersAndPaymentsAPI.md#V1OrdersList) | **Get** /v1/orders |
[**V1OrdersRetrieve**](OrdersAndPaymentsAPI.md#V1OrdersRetrieve) | **Get** /v1/orders/{id} |
[**V1PaymentReceiptVerificationKeysRetrieve**](OrdersAndPaymentsAPI.md#V1PaymentReceiptVerificationKeysRetrieve) | **Get** /v1/payment-receipt-verification-keys |
[**V1PaymentsList**](OrdersAndPaymentsAPI.md#V1PaymentsList) | **Get** /v1/payments |
[**V1PaymentsObservationsList**](OrdersAndPaymentsAPI.md#V1PaymentsObservationsList) | **Get** /v1/payments/{id}/observations |
[**V1PaymentsReceiptRetrieve**](OrdersAndPaymentsAPI.md#V1PaymentsReceiptRetrieve) | **Get** /v1/payments/{id}/receipt |
[**V1PaymentsRetrieve**](OrdersAndPaymentsAPI.md#V1PaymentsRetrieve) | **Get** /v1/payments/{id} |



## V1OrdersList

> []Order V1OrdersList(ctx).Cursor(cursor).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1OrdersList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1OrdersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OrdersList`: []Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1OrdersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OrdersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]Order**](Order.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrdersRetrieve

> Order V1OrdersRetrieve(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1OrdersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1OrdersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OrdersRetrieve`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1OrdersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrdersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Order**](Order.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentReceiptVerificationKeysRetrieve

> ReceiptVerificationKeyHistory V1PaymentReceiptVerificationKeysRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1PaymentReceiptVerificationKeysRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1PaymentReceiptVerificationKeysRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentReceiptVerificationKeysRetrieve`: ReceiptVerificationKeyHistory
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1PaymentReceiptVerificationKeysRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentReceiptVerificationKeysRetrieveRequest struct via the builder pattern


### Return type

[**ReceiptVerificationKeyHistory**](ReceiptVerificationKeyHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentsList

> []SettlementJob V1PaymentsList(ctx).Cursor(cursor).PageSize(pageSize).Execute()



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
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1PaymentsList(context.Background()).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1PaymentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentsList`: []SettlementJob
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1PaymentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]SettlementJob**](SettlementJob.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentsObservationsList

> []SettlementChainObservation V1PaymentsObservationsList(ctx, id).Cursor(cursor).PageSize(pageSize).Execute()



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
	cursor := "cursor_example" // string | Opaque pagination cursor from X-X402API-Next-Cursor or rel=next Link. (optional)
	pageSize := int32(56) // int32 | Number of results in the bounded array page (default and maximum 100). (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1PaymentsObservationsList(context.Background(), id).Cursor(cursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1PaymentsObservationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentsObservationsList`: []SettlementChainObservation
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1PaymentsObservationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentsObservationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor from X-X402API-Next-Cursor or rel&#x3D;next Link. |
 **pageSize** | **int32** | Number of results in the bounded array page (default and maximum 100). | [default to 100]

### Return type

[**[]SettlementChainObservation**](SettlementChainObservation.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentsReceiptRetrieve

> PaymentReceipt V1PaymentsReceiptRetrieve(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1PaymentsReceiptRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1PaymentsReceiptRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentsReceiptRetrieve`: PaymentReceipt
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1PaymentsReceiptRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentsReceiptRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentReceipt**](PaymentReceipt.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentsRetrieve

> SettlementJob V1PaymentsRetrieve(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAndPaymentsAPI.V1PaymentsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAndPaymentsAPI.V1PaymentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PaymentsRetrieve`: SettlementJob
	fmt.Fprintf(os.Stdout, "Response from `OrdersAndPaymentsAPI.V1PaymentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiV1PaymentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettlementJob**](SettlementJob.md)

### Authorization

[tenantApiKey](../README.md#tenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
