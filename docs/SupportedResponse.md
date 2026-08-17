# SupportedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kinds** | [**[]SupportedKind**](SupportedKind.md) |  |
**Extensions** | **[]string** |  |
**Signers** | **map[string][]string** |  |

## Methods

### NewSupportedResponse

`func NewSupportedResponse(kinds []SupportedKind, extensions []string, signers map[string][]string, ) *SupportedResponse`

NewSupportedResponse instantiates a new SupportedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedResponseWithDefaults

`func NewSupportedResponseWithDefaults() *SupportedResponse`

NewSupportedResponseWithDefaults instantiates a new SupportedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKinds

`func (o *SupportedResponse) GetKinds() []SupportedKind`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *SupportedResponse) GetKindsOk() (*[]SupportedKind, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *SupportedResponse) SetKinds(v []SupportedKind)`

SetKinds sets Kinds field to given value.


### GetExtensions

`func (o *SupportedResponse) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SupportedResponse) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SupportedResponse) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.


### GetSigners

`func (o *SupportedResponse) GetSigners() map[string][]string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *SupportedResponse) GetSignersOk() (*map[string][]string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *SupportedResponse) SetSigners(v map[string][]string)`

SetSigners sets Signers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
