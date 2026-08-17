# SupportedKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X402Version** | **int32** |  |
**Scheme** | **string** |  |
**Network** | **string** |  |
**Extra** | Pointer to **interface{}** |  | [optional]

## Methods

### NewSupportedKind

`func NewSupportedKind(x402Version int32, scheme string, network string, ) *SupportedKind`

NewSupportedKind instantiates a new SupportedKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedKindWithDefaults

`func NewSupportedKindWithDefaults() *SupportedKind`

NewSupportedKindWithDefaults instantiates a new SupportedKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX402Version

`func (o *SupportedKind) GetX402Version() int32`

GetX402Version returns the X402Version field if non-nil, zero value otherwise.

### GetX402VersionOk

`func (o *SupportedKind) GetX402VersionOk() (*int32, bool)`

GetX402VersionOk returns a tuple with the X402Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX402Version

`func (o *SupportedKind) SetX402Version(v int32)`

SetX402Version sets X402Version field to given value.


### GetScheme

`func (o *SupportedKind) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SupportedKind) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SupportedKind) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetNetwork

`func (o *SupportedKind) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SupportedKind) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SupportedKind) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetExtra

`func (o *SupportedKind) GetExtra() interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SupportedKind) GetExtraOk() (*interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SupportedKind) SetExtra(v interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *SupportedKind) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *SupportedKind) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *SupportedKind) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
