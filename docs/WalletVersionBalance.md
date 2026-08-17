# WalletVersionBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletVersionId** | **string** |  |
**Version** | **int32** |  |
**WalletAddress** | **string** |  |
**State** | [**WalletVersionBalanceStateEnum**](WalletVersionBalanceStateEnum.md) |  |
**ObservationState** | [**ObservationStateEnum**](ObservationStateEnum.md) |  |
**ObservedAt** | **NullableTime** |  |
**Assets** | [**[]BalanceAsset**](BalanceAsset.md) |  |
**ReseedContext** | [**NullableWalletChainReseedContext**](WalletChainReseedContext.md) |  |

## Methods

### NewWalletVersionBalance

`func NewWalletVersionBalance(walletVersionId string, version int32, walletAddress string, state WalletVersionBalanceStateEnum, observationState ObservationStateEnum, observedAt NullableTime, assets []BalanceAsset, reseedContext NullableWalletChainReseedContext, ) *WalletVersionBalance`

NewWalletVersionBalance instantiates a new WalletVersionBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletVersionBalanceWithDefaults

`func NewWalletVersionBalanceWithDefaults() *WalletVersionBalance`

NewWalletVersionBalanceWithDefaults instantiates a new WalletVersionBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletVersionId

`func (o *WalletVersionBalance) GetWalletVersionId() string`

GetWalletVersionId returns the WalletVersionId field if non-nil, zero value otherwise.

### GetWalletVersionIdOk

`func (o *WalletVersionBalance) GetWalletVersionIdOk() (*string, bool)`

GetWalletVersionIdOk returns a tuple with the WalletVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionId

`func (o *WalletVersionBalance) SetWalletVersionId(v string)`

SetWalletVersionId sets WalletVersionId field to given value.


### GetVersion

`func (o *WalletVersionBalance) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WalletVersionBalance) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WalletVersionBalance) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWalletAddress

`func (o *WalletVersionBalance) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *WalletVersionBalance) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *WalletVersionBalance) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetState

`func (o *WalletVersionBalance) GetState() WalletVersionBalanceStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WalletVersionBalance) GetStateOk() (*WalletVersionBalanceStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WalletVersionBalance) SetState(v WalletVersionBalanceStateEnum)`

SetState sets State field to given value.


### GetObservationState

`func (o *WalletVersionBalance) GetObservationState() ObservationStateEnum`

GetObservationState returns the ObservationState field if non-nil, zero value otherwise.

### GetObservationStateOk

`func (o *WalletVersionBalance) GetObservationStateOk() (*ObservationStateEnum, bool)`

GetObservationStateOk returns a tuple with the ObservationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationState

`func (o *WalletVersionBalance) SetObservationState(v ObservationStateEnum)`

SetObservationState sets ObservationState field to given value.


### GetObservedAt

`func (o *WalletVersionBalance) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *WalletVersionBalance) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *WalletVersionBalance) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *WalletVersionBalance) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *WalletVersionBalance) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetAssets

`func (o *WalletVersionBalance) GetAssets() []BalanceAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *WalletVersionBalance) GetAssetsOk() (*[]BalanceAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *WalletVersionBalance) SetAssets(v []BalanceAsset)`

SetAssets sets Assets field to given value.


### GetReseedContext

`func (o *WalletVersionBalance) GetReseedContext() WalletChainReseedContext`

GetReseedContext returns the ReseedContext field if non-nil, zero value otherwise.

### GetReseedContextOk

`func (o *WalletVersionBalance) GetReseedContextOk() (*WalletChainReseedContext, bool)`

GetReseedContextOk returns a tuple with the ReseedContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReseedContext

`func (o *WalletVersionBalance) SetReseedContext(v WalletChainReseedContext)`

SetReseedContext sets ReseedContext field to given value.


### SetReseedContextNil

`func (o *WalletVersionBalance) SetReseedContextNil(b bool)`

 SetReseedContextNil sets the value for ReseedContext to be an explicit nil

### UnsetReseedContext
`func (o *WalletVersionBalance) UnsetReseedContext()`

UnsetReseedContext ensures that no value is present for ReseedContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
