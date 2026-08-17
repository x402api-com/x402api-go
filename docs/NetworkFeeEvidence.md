# NetworkFeeEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Version** | **int32** |  |
**Network** | **string** |  |
**AssetId** | **string** |  |
**PayloadProfile** | **string** |  |
**NativeSymbol** | Pointer to **string** |  | [optional]
**NativeDecimals** | Pointer to **int32** |  | [optional]
**NativeFeeObservations** | Pointer to [**[]NativeFeeObservationEvidence**](NativeFeeObservationEvidence.md) |  | [optional]
**NativeUsdObservations** | Pointer to [**[]NativeUsdObservationEvidence**](NativeUsdObservationEvidence.md) |  | [optional]
**ExpiresAt** | Pointer to **time.Time** |  | [optional]

## Methods

### NewNetworkFeeEvidence

`func NewNetworkFeeEvidence(type_ string, version int32, network string, assetId string, payloadProfile string, ) *NetworkFeeEvidence`

NewNetworkFeeEvidence instantiates a new NetworkFeeEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeeEvidenceWithDefaults

`func NewNetworkFeeEvidenceWithDefaults() *NetworkFeeEvidence`

NewNetworkFeeEvidenceWithDefaults instantiates a new NetworkFeeEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkFeeEvidence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkFeeEvidence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkFeeEvidence) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *NetworkFeeEvidence) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkFeeEvidence) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkFeeEvidence) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetNetwork

`func (o *NetworkFeeEvidence) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkFeeEvidence) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkFeeEvidence) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *NetworkFeeEvidence) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *NetworkFeeEvidence) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *NetworkFeeEvidence) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetPayloadProfile

`func (o *NetworkFeeEvidence) GetPayloadProfile() string`

GetPayloadProfile returns the PayloadProfile field if non-nil, zero value otherwise.

### GetPayloadProfileOk

`func (o *NetworkFeeEvidence) GetPayloadProfileOk() (*string, bool)`

GetPayloadProfileOk returns a tuple with the PayloadProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadProfile

`func (o *NetworkFeeEvidence) SetPayloadProfile(v string)`

SetPayloadProfile sets PayloadProfile field to given value.


### GetNativeSymbol

`func (o *NetworkFeeEvidence) GetNativeSymbol() string`

GetNativeSymbol returns the NativeSymbol field if non-nil, zero value otherwise.

### GetNativeSymbolOk

`func (o *NetworkFeeEvidence) GetNativeSymbolOk() (*string, bool)`

GetNativeSymbolOk returns a tuple with the NativeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeSymbol

`func (o *NetworkFeeEvidence) SetNativeSymbol(v string)`

SetNativeSymbol sets NativeSymbol field to given value.

### HasNativeSymbol

`func (o *NetworkFeeEvidence) HasNativeSymbol() bool`

HasNativeSymbol returns a boolean if a field has been set.

### GetNativeDecimals

`func (o *NetworkFeeEvidence) GetNativeDecimals() int32`

GetNativeDecimals returns the NativeDecimals field if non-nil, zero value otherwise.

### GetNativeDecimalsOk

`func (o *NetworkFeeEvidence) GetNativeDecimalsOk() (*int32, bool)`

GetNativeDecimalsOk returns a tuple with the NativeDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeDecimals

`func (o *NetworkFeeEvidence) SetNativeDecimals(v int32)`

SetNativeDecimals sets NativeDecimals field to given value.

### HasNativeDecimals

`func (o *NetworkFeeEvidence) HasNativeDecimals() bool`

HasNativeDecimals returns a boolean if a field has been set.

### GetNativeFeeObservations

`func (o *NetworkFeeEvidence) GetNativeFeeObservations() []NativeFeeObservationEvidence`

GetNativeFeeObservations returns the NativeFeeObservations field if non-nil, zero value otherwise.

### GetNativeFeeObservationsOk

`func (o *NetworkFeeEvidence) GetNativeFeeObservationsOk() (*[]NativeFeeObservationEvidence, bool)`

GetNativeFeeObservationsOk returns a tuple with the NativeFeeObservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFeeObservations

`func (o *NetworkFeeEvidence) SetNativeFeeObservations(v []NativeFeeObservationEvidence)`

SetNativeFeeObservations sets NativeFeeObservations field to given value.

### HasNativeFeeObservations

`func (o *NetworkFeeEvidence) HasNativeFeeObservations() bool`

HasNativeFeeObservations returns a boolean if a field has been set.

### GetNativeUsdObservations

`func (o *NetworkFeeEvidence) GetNativeUsdObservations() []NativeUsdObservationEvidence`

GetNativeUsdObservations returns the NativeUsdObservations field if non-nil, zero value otherwise.

### GetNativeUsdObservationsOk

`func (o *NetworkFeeEvidence) GetNativeUsdObservationsOk() (*[]NativeUsdObservationEvidence, bool)`

GetNativeUsdObservationsOk returns a tuple with the NativeUsdObservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeUsdObservations

`func (o *NetworkFeeEvidence) SetNativeUsdObservations(v []NativeUsdObservationEvidence)`

SetNativeUsdObservations sets NativeUsdObservations field to given value.

### HasNativeUsdObservations

`func (o *NetworkFeeEvidence) HasNativeUsdObservations() bool`

HasNativeUsdObservations returns a boolean if a field has been set.

### GetExpiresAt

`func (o *NetworkFeeEvidence) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NetworkFeeEvidence) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NetworkFeeEvidence) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *NetworkFeeEvidence) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
