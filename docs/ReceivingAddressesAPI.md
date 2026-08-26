# \ReceivingAddressesAPI

All URIs are relative to *https://api.x402api.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceivingAddressesGetControlCapabilities**](ReceivingAddressesAPI.md#ReceivingAddressesGetControlCapabilities) | **Get** /v1/receiving-address-control-capabilities | Get receiving-address control capabilities
[**ReceivingAddressesList**](ReceivingAddressesAPI.md#ReceivingAddressesList) | **Get** /v1/receiving-addresses | List receiving addresses



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
