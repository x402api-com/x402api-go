# ExternalAddressControlProofInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**ExternalAddressProofInputMethodEnum**](ExternalAddressProofInputMethodEnum.md) |  |
**Signature** | Pointer to **string** |  | [optional]
**TransactionHash** | Pointer to **string** |  | [optional]

## Methods

### NewExternalAddressControlProofInput

`func NewExternalAddressControlProofInput(method ExternalAddressProofInputMethodEnum, ) *ExternalAddressControlProofInput`

NewExternalAddressControlProofInput instantiates a new ExternalAddressControlProofInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAddressControlProofInputWithDefaults

`func NewExternalAddressControlProofInputWithDefaults() *ExternalAddressControlProofInput`

NewExternalAddressControlProofInputWithDefaults instantiates a new ExternalAddressControlProofInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ExternalAddressControlProofInput) GetMethod() ExternalAddressProofInputMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExternalAddressControlProofInput) GetMethodOk() (*ExternalAddressProofInputMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExternalAddressControlProofInput) SetMethod(v ExternalAddressProofInputMethodEnum)`

SetMethod sets Method field to given value.


### GetSignature

`func (o *ExternalAddressControlProofInput) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ExternalAddressControlProofInput) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ExternalAddressControlProofInput) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ExternalAddressControlProofInput) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ExternalAddressControlProofInput) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ExternalAddressControlProofInput) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ExternalAddressControlProofInput) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ExternalAddressControlProofInput) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
