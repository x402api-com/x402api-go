# PaymentReadinessAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | [readonly]
**DisplayName** | **string** |  | [readonly]
**ContractAddress** | **string** |  | [readonly]
**IssuerNative** | **bool** |  | [readonly]
**RegistryEnabled** | **bool** |  | [readonly]
**TenantEnabled** | **bool** |  | [readonly]
**OperatorAssistanceEnabled** | **bool** |  | [readonly]
**BaseReadinessBlockers** | **[]string** |  | [readonly]
**ChallengeControlReady** | **bool** |  | [readonly]
**SettlementControlReady** | **bool** |  | [readonly]

## Methods

### NewPaymentReadinessAsset

`func NewPaymentReadinessAsset(assetId string, displayName string, contractAddress string, issuerNative bool, registryEnabled bool, tenantEnabled bool, operatorAssistanceEnabled bool, baseReadinessBlockers []string, challengeControlReady bool, settlementControlReady bool, ) *PaymentReadinessAsset`

NewPaymentReadinessAsset instantiates a new PaymentReadinessAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReadinessAssetWithDefaults

`func NewPaymentReadinessAssetWithDefaults() *PaymentReadinessAsset`

NewPaymentReadinessAssetWithDefaults instantiates a new PaymentReadinessAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *PaymentReadinessAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *PaymentReadinessAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *PaymentReadinessAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetDisplayName

`func (o *PaymentReadinessAsset) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentReadinessAsset) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentReadinessAsset) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetContractAddress

`func (o *PaymentReadinessAsset) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *PaymentReadinessAsset) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *PaymentReadinessAsset) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetIssuerNative

`func (o *PaymentReadinessAsset) GetIssuerNative() bool`

GetIssuerNative returns the IssuerNative field if non-nil, zero value otherwise.

### GetIssuerNativeOk

`func (o *PaymentReadinessAsset) GetIssuerNativeOk() (*bool, bool)`

GetIssuerNativeOk returns a tuple with the IssuerNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerNative

`func (o *PaymentReadinessAsset) SetIssuerNative(v bool)`

SetIssuerNative sets IssuerNative field to given value.


### GetRegistryEnabled

`func (o *PaymentReadinessAsset) GetRegistryEnabled() bool`

GetRegistryEnabled returns the RegistryEnabled field if non-nil, zero value otherwise.

### GetRegistryEnabledOk

`func (o *PaymentReadinessAsset) GetRegistryEnabledOk() (*bool, bool)`

GetRegistryEnabledOk returns a tuple with the RegistryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryEnabled

`func (o *PaymentReadinessAsset) SetRegistryEnabled(v bool)`

SetRegistryEnabled sets RegistryEnabled field to given value.


### GetTenantEnabled

`func (o *PaymentReadinessAsset) GetTenantEnabled() bool`

GetTenantEnabled returns the TenantEnabled field if non-nil, zero value otherwise.

### GetTenantEnabledOk

`func (o *PaymentReadinessAsset) GetTenantEnabledOk() (*bool, bool)`

GetTenantEnabledOk returns a tuple with the TenantEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantEnabled

`func (o *PaymentReadinessAsset) SetTenantEnabled(v bool)`

SetTenantEnabled sets TenantEnabled field to given value.


### GetOperatorAssistanceEnabled

`func (o *PaymentReadinessAsset) GetOperatorAssistanceEnabled() bool`

GetOperatorAssistanceEnabled returns the OperatorAssistanceEnabled field if non-nil, zero value otherwise.

### GetOperatorAssistanceEnabledOk

`func (o *PaymentReadinessAsset) GetOperatorAssistanceEnabledOk() (*bool, bool)`

GetOperatorAssistanceEnabledOk returns a tuple with the OperatorAssistanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorAssistanceEnabled

`func (o *PaymentReadinessAsset) SetOperatorAssistanceEnabled(v bool)`

SetOperatorAssistanceEnabled sets OperatorAssistanceEnabled field to given value.


### GetBaseReadinessBlockers

`func (o *PaymentReadinessAsset) GetBaseReadinessBlockers() []string`

GetBaseReadinessBlockers returns the BaseReadinessBlockers field if non-nil, zero value otherwise.

### GetBaseReadinessBlockersOk

`func (o *PaymentReadinessAsset) GetBaseReadinessBlockersOk() (*[]string, bool)`

GetBaseReadinessBlockersOk returns a tuple with the BaseReadinessBlockers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseReadinessBlockers

`func (o *PaymentReadinessAsset) SetBaseReadinessBlockers(v []string)`

SetBaseReadinessBlockers sets BaseReadinessBlockers field to given value.


### GetChallengeControlReady

`func (o *PaymentReadinessAsset) GetChallengeControlReady() bool`

GetChallengeControlReady returns the ChallengeControlReady field if non-nil, zero value otherwise.

### GetChallengeControlReadyOk

`func (o *PaymentReadinessAsset) GetChallengeControlReadyOk() (*bool, bool)`

GetChallengeControlReadyOk returns a tuple with the ChallengeControlReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeControlReady

`func (o *PaymentReadinessAsset) SetChallengeControlReady(v bool)`

SetChallengeControlReady sets ChallengeControlReady field to given value.


### GetSettlementControlReady

`func (o *PaymentReadinessAsset) GetSettlementControlReady() bool`

GetSettlementControlReady returns the SettlementControlReady field if non-nil, zero value otherwise.

### GetSettlementControlReadyOk

`func (o *PaymentReadinessAsset) GetSettlementControlReadyOk() (*bool, bool)`

GetSettlementControlReadyOk returns a tuple with the SettlementControlReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementControlReady

`func (o *PaymentReadinessAsset) SetSettlementControlReady(v bool)`

SetSettlementControlReady sets SettlementControlReady field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
