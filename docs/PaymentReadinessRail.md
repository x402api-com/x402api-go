# PaymentReadinessRail

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
**Status** | **string** |  | [readonly]
**Blockers** | [**[]PaymentReadinessBlocker**](PaymentReadinessBlocker.md) |  | [readonly]
**TenantChallengesEnabled** | **bool** |  | [readonly]
**TenantSettlementEnabled** | **bool** |  | [readonly]
**NetworkAssistanceEnabled** | **bool** |  | [readonly]
**ChallengeControlReady** | **bool** |  | [readonly]
**SettlementControlReady** | **bool** |  | [readonly]
**Assets** | [**[]PaymentReadinessAsset**](PaymentReadinessAsset.md) |  | [readonly]

## Methods

### NewPaymentReadinessRail

`func NewPaymentReadinessRail(assetId string, network string, symbol string, selected bool, walletReady bool, platformAvailable bool, acceptingNewPayments bool, status string, blockers []PaymentReadinessBlocker, tenantChallengesEnabled bool, tenantSettlementEnabled bool, networkAssistanceEnabled bool, challengeControlReady bool, settlementControlReady bool, assets []PaymentReadinessAsset, ) *PaymentReadinessRail`

NewPaymentReadinessRail instantiates a new PaymentReadinessRail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReadinessRailWithDefaults

`func NewPaymentReadinessRailWithDefaults() *PaymentReadinessRail`

NewPaymentReadinessRailWithDefaults instantiates a new PaymentReadinessRail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *PaymentReadinessRail) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *PaymentReadinessRail) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *PaymentReadinessRail) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNetwork

`func (o *PaymentReadinessRail) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PaymentReadinessRail) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PaymentReadinessRail) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSymbol

`func (o *PaymentReadinessRail) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PaymentReadinessRail) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PaymentReadinessRail) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSelected

`func (o *PaymentReadinessRail) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *PaymentReadinessRail) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *PaymentReadinessRail) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetWalletReady

`func (o *PaymentReadinessRail) GetWalletReady() bool`

GetWalletReady returns the WalletReady field if non-nil, zero value otherwise.

### GetWalletReadyOk

`func (o *PaymentReadinessRail) GetWalletReadyOk() (*bool, bool)`

GetWalletReadyOk returns a tuple with the WalletReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletReady

`func (o *PaymentReadinessRail) SetWalletReady(v bool)`

SetWalletReady sets WalletReady field to given value.


### GetPlatformAvailable

`func (o *PaymentReadinessRail) GetPlatformAvailable() bool`

GetPlatformAvailable returns the PlatformAvailable field if non-nil, zero value otherwise.

### GetPlatformAvailableOk

`func (o *PaymentReadinessRail) GetPlatformAvailableOk() (*bool, bool)`

GetPlatformAvailableOk returns a tuple with the PlatformAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAvailable

`func (o *PaymentReadinessRail) SetPlatformAvailable(v bool)`

SetPlatformAvailable sets PlatformAvailable field to given value.


### GetAcceptingNewPayments

`func (o *PaymentReadinessRail) GetAcceptingNewPayments() bool`

GetAcceptingNewPayments returns the AcceptingNewPayments field if non-nil, zero value otherwise.

### GetAcceptingNewPaymentsOk

`func (o *PaymentReadinessRail) GetAcceptingNewPaymentsOk() (*bool, bool)`

GetAcceptingNewPaymentsOk returns a tuple with the AcceptingNewPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingNewPayments

`func (o *PaymentReadinessRail) SetAcceptingNewPayments(v bool)`

SetAcceptingNewPayments sets AcceptingNewPayments field to given value.


### GetStatus

`func (o *PaymentReadinessRail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentReadinessRail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentReadinessRail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBlockers

`func (o *PaymentReadinessRail) GetBlockers() []PaymentReadinessBlocker`

GetBlockers returns the Blockers field if non-nil, zero value otherwise.

### GetBlockersOk

`func (o *PaymentReadinessRail) GetBlockersOk() (*[]PaymentReadinessBlocker, bool)`

GetBlockersOk returns a tuple with the Blockers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockers

`func (o *PaymentReadinessRail) SetBlockers(v []PaymentReadinessBlocker)`

SetBlockers sets Blockers field to given value.


### GetTenantChallengesEnabled

`func (o *PaymentReadinessRail) GetTenantChallengesEnabled() bool`

GetTenantChallengesEnabled returns the TenantChallengesEnabled field if non-nil, zero value otherwise.

### GetTenantChallengesEnabledOk

`func (o *PaymentReadinessRail) GetTenantChallengesEnabledOk() (*bool, bool)`

GetTenantChallengesEnabledOk returns a tuple with the TenantChallengesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantChallengesEnabled

`func (o *PaymentReadinessRail) SetTenantChallengesEnabled(v bool)`

SetTenantChallengesEnabled sets TenantChallengesEnabled field to given value.


### GetTenantSettlementEnabled

`func (o *PaymentReadinessRail) GetTenantSettlementEnabled() bool`

GetTenantSettlementEnabled returns the TenantSettlementEnabled field if non-nil, zero value otherwise.

### GetTenantSettlementEnabledOk

`func (o *PaymentReadinessRail) GetTenantSettlementEnabledOk() (*bool, bool)`

GetTenantSettlementEnabledOk returns a tuple with the TenantSettlementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSettlementEnabled

`func (o *PaymentReadinessRail) SetTenantSettlementEnabled(v bool)`

SetTenantSettlementEnabled sets TenantSettlementEnabled field to given value.


### GetNetworkAssistanceEnabled

`func (o *PaymentReadinessRail) GetNetworkAssistanceEnabled() bool`

GetNetworkAssistanceEnabled returns the NetworkAssistanceEnabled field if non-nil, zero value otherwise.

### GetNetworkAssistanceEnabledOk

`func (o *PaymentReadinessRail) GetNetworkAssistanceEnabledOk() (*bool, bool)`

GetNetworkAssistanceEnabledOk returns a tuple with the NetworkAssistanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAssistanceEnabled

`func (o *PaymentReadinessRail) SetNetworkAssistanceEnabled(v bool)`

SetNetworkAssistanceEnabled sets NetworkAssistanceEnabled field to given value.


### GetChallengeControlReady

`func (o *PaymentReadinessRail) GetChallengeControlReady() bool`

GetChallengeControlReady returns the ChallengeControlReady field if non-nil, zero value otherwise.

### GetChallengeControlReadyOk

`func (o *PaymentReadinessRail) GetChallengeControlReadyOk() (*bool, bool)`

GetChallengeControlReadyOk returns a tuple with the ChallengeControlReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeControlReady

`func (o *PaymentReadinessRail) SetChallengeControlReady(v bool)`

SetChallengeControlReady sets ChallengeControlReady field to given value.


### GetSettlementControlReady

`func (o *PaymentReadinessRail) GetSettlementControlReady() bool`

GetSettlementControlReady returns the SettlementControlReady field if non-nil, zero value otherwise.

### GetSettlementControlReadyOk

`func (o *PaymentReadinessRail) GetSettlementControlReadyOk() (*bool, bool)`

GetSettlementControlReadyOk returns a tuple with the SettlementControlReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementControlReady

`func (o *PaymentReadinessRail) SetSettlementControlReady(v bool)`

SetSettlementControlReady sets SettlementControlReady field to given value.


### GetAssets

`func (o *PaymentReadinessRail) GetAssets() []PaymentReadinessAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PaymentReadinessRail) GetAssetsOk() (*[]PaymentReadinessAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PaymentReadinessRail) SetAssets(v []PaymentReadinessAsset)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
