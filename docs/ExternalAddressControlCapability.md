# ExternalAddressControlCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  | [readonly]
**ProofMethods** | [**[]ExternalAddressProofInputMethodEnum**](ExternalAddressProofInputMethodEnum.md) |  | [readonly]

## Methods

### NewExternalAddressControlCapability

`func NewExternalAddressControlCapability(network string, proofMethods []ExternalAddressProofInputMethodEnum, ) *ExternalAddressControlCapability`

NewExternalAddressControlCapability instantiates a new ExternalAddressControlCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAddressControlCapabilityWithDefaults

`func NewExternalAddressControlCapabilityWithDefaults() *ExternalAddressControlCapability`

NewExternalAddressControlCapabilityWithDefaults instantiates a new ExternalAddressControlCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ExternalAddressControlCapability) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExternalAddressControlCapability) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExternalAddressControlCapability) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetProofMethods

`func (o *ExternalAddressControlCapability) GetProofMethods() []ExternalAddressProofInputMethodEnum`

GetProofMethods returns the ProofMethods field if non-nil, zero value otherwise.

### GetProofMethodsOk

`func (o *ExternalAddressControlCapability) GetProofMethodsOk() (*[]ExternalAddressProofInputMethodEnum, bool)`

GetProofMethodsOk returns a tuple with the ProofMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofMethods

`func (o *ExternalAddressControlCapability) SetProofMethods(v []ExternalAddressProofInputMethodEnum)`

SetProofMethods sets ProofMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
