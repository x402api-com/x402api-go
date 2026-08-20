# WalletBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** |  |
**Network** | **string** |  |
**WalletAddress** | **NullableString** |  |
**RequestedFinality** | [**WalletObservationFinalityEnum**](WalletObservationFinalityEnum.md) |  |
**ObservationState** | [**ObservationStateEnum**](ObservationStateEnum.md) |  |
**TrackingStatus** | [**TrackingStatusEnum**](TrackingStatusEnum.md) |  |
**ObservedAt** | **NullableTime** |  |
**Assets** | [**[]BalanceAsset**](BalanceAsset.md) |  |
**WalletVersions** | [**[]WalletVersionBalance**](WalletVersionBalance.md) |  |
**ReseedContexts** | [**[]WalletFencedChainReseedContext**](WalletFencedChainReseedContext.md) |  |

## Methods

### NewWalletBalanceResponse

`func NewWalletBalanceResponse(walletId string, network string, walletAddress NullableString, requestedFinality WalletObservationFinalityEnum, observationState ObservationStateEnum, trackingStatus TrackingStatusEnum, observedAt NullableTime, assets []BalanceAsset, walletVersions []WalletVersionBalance, reseedContexts []WalletFencedChainReseedContext, ) *WalletBalanceResponse`

NewWalletBalanceResponse instantiates a new WalletBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceResponseWithDefaults

`func NewWalletBalanceResponseWithDefaults() *WalletBalanceResponse`

NewWalletBalanceResponseWithDefaults instantiates a new WalletBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletBalanceResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletBalanceResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletBalanceResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetNetwork

`func (o *WalletBalanceResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WalletBalanceResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WalletBalanceResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetWalletAddress

`func (o *WalletBalanceResponse) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *WalletBalanceResponse) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *WalletBalanceResponse) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### SetWalletAddressNil

`func (o *WalletBalanceResponse) SetWalletAddressNil(b bool)`

 SetWalletAddressNil sets the value for WalletAddress to be an explicit nil

### UnsetWalletAddress
`func (o *WalletBalanceResponse) UnsetWalletAddress()`

UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
### GetRequestedFinality

`func (o *WalletBalanceResponse) GetRequestedFinality() WalletObservationFinalityEnum`

GetRequestedFinality returns the RequestedFinality field if non-nil, zero value otherwise.

### GetRequestedFinalityOk

`func (o *WalletBalanceResponse) GetRequestedFinalityOk() (*WalletObservationFinalityEnum, bool)`

GetRequestedFinalityOk returns a tuple with the RequestedFinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFinality

`func (o *WalletBalanceResponse) SetRequestedFinality(v WalletObservationFinalityEnum)`

SetRequestedFinality sets RequestedFinality field to given value.


### GetObservationState

`func (o *WalletBalanceResponse) GetObservationState() ObservationStateEnum`

GetObservationState returns the ObservationState field if non-nil, zero value otherwise.

### GetObservationStateOk

`func (o *WalletBalanceResponse) GetObservationStateOk() (*ObservationStateEnum, bool)`

GetObservationStateOk returns a tuple with the ObservationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationState

`func (o *WalletBalanceResponse) SetObservationState(v ObservationStateEnum)`

SetObservationState sets ObservationState field to given value.


### GetTrackingStatus

`func (o *WalletBalanceResponse) GetTrackingStatus() TrackingStatusEnum`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *WalletBalanceResponse) GetTrackingStatusOk() (*TrackingStatusEnum, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *WalletBalanceResponse) SetTrackingStatus(v TrackingStatusEnum)`

SetTrackingStatus sets TrackingStatus field to given value.


### GetObservedAt

`func (o *WalletBalanceResponse) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *WalletBalanceResponse) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *WalletBalanceResponse) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *WalletBalanceResponse) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *WalletBalanceResponse) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil
### GetAssets

`func (o *WalletBalanceResponse) GetAssets() []BalanceAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *WalletBalanceResponse) GetAssetsOk() (*[]BalanceAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *WalletBalanceResponse) SetAssets(v []BalanceAsset)`

SetAssets sets Assets field to given value.


### GetWalletVersions

`func (o *WalletBalanceResponse) GetWalletVersions() []WalletVersionBalance`

GetWalletVersions returns the WalletVersions field if non-nil, zero value otherwise.

### GetWalletVersionsOk

`func (o *WalletBalanceResponse) GetWalletVersionsOk() (*[]WalletVersionBalance, bool)`

GetWalletVersionsOk returns a tuple with the WalletVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersions

`func (o *WalletBalanceResponse) SetWalletVersions(v []WalletVersionBalance)`

SetWalletVersions sets WalletVersions field to given value.


### GetReseedContexts

`func (o *WalletBalanceResponse) GetReseedContexts() []WalletFencedChainReseedContext`

GetReseedContexts returns the ReseedContexts field if non-nil, zero value otherwise.

### GetReseedContextsOk

`func (o *WalletBalanceResponse) GetReseedContextsOk() (*[]WalletFencedChainReseedContext, bool)`

GetReseedContextsOk returns a tuple with the ReseedContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReseedContexts

`func (o *WalletBalanceResponse) SetReseedContexts(v []WalletFencedChainReseedContext)`

SetReseedContexts sets ReseedContexts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
