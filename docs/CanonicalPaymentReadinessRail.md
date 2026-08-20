# CanonicalPaymentReadinessRail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | [readonly]
**Network** | **string** |  | [readonly]
**Symbol** | **string** |  | [readonly]
**Selected** | **bool** |  | [readonly]
**WalletReady** | **bool** |  | [readonly]
**PlatformAvailable** | **bool** |  | [readonly]
**AcceptingNewPayments** | **bool** |  | [readonly]
**Status** | [**PaymentReadinessRailStatusEnum**](PaymentReadinessRailStatusEnum.md) |  | [readonly]
**Blockers** | [**[]PaymentReadinessBlocker**](PaymentReadinessBlocker.md) |  | [readonly]

## Methods

### NewCanonicalPaymentReadinessRail

`func NewCanonicalPaymentReadinessRail(assetId string, network string, symbol string, selected bool, walletReady bool, platformAvailable bool, acceptingNewPayments bool, status PaymentReadinessRailStatusEnum, blockers []PaymentReadinessBlocker, ) *CanonicalPaymentReadinessRail`

NewCanonicalPaymentReadinessRail instantiates a new CanonicalPaymentReadinessRail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanonicalPaymentReadinessRailWithDefaults

`func NewCanonicalPaymentReadinessRailWithDefaults() *CanonicalPaymentReadinessRail`

NewCanonicalPaymentReadinessRailWithDefaults instantiates a new CanonicalPaymentReadinessRail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *CanonicalPaymentReadinessRail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CanonicalPaymentReadinessRail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CanonicalPaymentReadinessRail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNetwork

`func (o *CanonicalPaymentReadinessRail) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CanonicalPaymentReadinessRail) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CanonicalPaymentReadinessRail) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSymbol

`func (o *CanonicalPaymentReadinessRail) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CanonicalPaymentReadinessRail) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CanonicalPaymentReadinessRail) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSelected

`func (o *CanonicalPaymentReadinessRail) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *CanonicalPaymentReadinessRail) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *CanonicalPaymentReadinessRail) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetWalletReady

`func (o *CanonicalPaymentReadinessRail) GetWalletReady() bool`

GetWalletReady returns the WalletReady field if non-nil, zero value otherwise.

### GetWalletReadyOk

`func (o *CanonicalPaymentReadinessRail) GetWalletReadyOk() (*bool, bool)`

GetWalletReadyOk returns a tuple with the WalletReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletReady

`func (o *CanonicalPaymentReadinessRail) SetWalletReady(v bool)`

SetWalletReady sets WalletReady field to given value.


### GetPlatformAvailable

`func (o *CanonicalPaymentReadinessRail) GetPlatformAvailable() bool`

GetPlatformAvailable returns the PlatformAvailable field if non-nil, zero value otherwise.

### GetPlatformAvailableOk

`func (o *CanonicalPaymentReadinessRail) GetPlatformAvailableOk() (*bool, bool)`

GetPlatformAvailableOk returns a tuple with the PlatformAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAvailable

`func (o *CanonicalPaymentReadinessRail) SetPlatformAvailable(v bool)`

SetPlatformAvailable sets PlatformAvailable field to given value.


### GetAcceptingNewPayments

`func (o *CanonicalPaymentReadinessRail) GetAcceptingNewPayments() bool`

GetAcceptingNewPayments returns the AcceptingNewPayments field if non-nil, zero value otherwise.

### GetAcceptingNewPaymentsOk

`func (o *CanonicalPaymentReadinessRail) GetAcceptingNewPaymentsOk() (*bool, bool)`

GetAcceptingNewPaymentsOk returns a tuple with the AcceptingNewPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingNewPayments

`func (o *CanonicalPaymentReadinessRail) SetAcceptingNewPayments(v bool)`

SetAcceptingNewPayments sets AcceptingNewPayments field to given value.


### GetStatus

`func (o *CanonicalPaymentReadinessRail) GetStatus() PaymentReadinessRailStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CanonicalPaymentReadinessRail) GetStatusOk() (*PaymentReadinessRailStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CanonicalPaymentReadinessRail) SetStatus(v PaymentReadinessRailStatusEnum)`

SetStatus sets Status field to given value.


### GetBlockers

`func (o *CanonicalPaymentReadinessRail) GetBlockers() []PaymentReadinessBlocker`

GetBlockers returns the Blockers field if non-nil, zero value otherwise.

### GetBlockersOk

`func (o *CanonicalPaymentReadinessRail) GetBlockersOk() (*[]PaymentReadinessBlocker, bool)`

GetBlockersOk returns a tuple with the Blockers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockers

`func (o *CanonicalPaymentReadinessRail) SetBlockers(v []PaymentReadinessBlocker)`

SetBlockers sets Blockers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
