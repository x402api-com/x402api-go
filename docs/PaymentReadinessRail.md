# PaymentReadinessRail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  | [readonly]
**TenantChallengesEnabled** | **bool** |  | [readonly]
**TenantSettlementEnabled** | **bool** |  | [readonly]
**NetworkAssistanceEnabled** | **bool** |  | [readonly]
**ChallengeControlReady** | **bool** |  | [readonly]
**SettlementControlReady** | **bool** |  | [readonly]
**Assets** | [**[]PaymentReadinessAsset**](PaymentReadinessAsset.md) |  | [readonly]

## Methods

### NewPaymentReadinessRail

`func NewPaymentReadinessRail(network string, tenantChallengesEnabled bool, tenantSettlementEnabled bool, networkAssistanceEnabled bool, challengeControlReady bool, settlementControlReady bool, assets []PaymentReadinessAsset, ) *PaymentReadinessRail`

NewPaymentReadinessRail instantiates a new PaymentReadinessRail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReadinessRailWithDefaults

`func NewPaymentReadinessRailWithDefaults() *PaymentReadinessRail`

NewPaymentReadinessRailWithDefaults instantiates a new PaymentReadinessRail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
