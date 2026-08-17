# ExternalReceivingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**WalletId** | **string** |  | [readonly]
**WalletVersionId** | **string** |  | [readonly]
**Label** | **string** |  | [readonly]
**Network** | **string** |  | [readonly]
**AssetId** | **string** |  | [readonly]
**Address** | **string** |  | [readonly]
**Status** | **string** |  | [readonly]
**ProofMethod** | [**NullableExternalAddressProofInputMethodEnum**](ExternalAddressProofInputMethodEnum.md) |  | [readonly]
**ProofVerifiedAt** | **NullableTime** |  | [readonly]
**ReadinessState** | **string** |  | [readonly]
**ReadinessUsable** | **bool** |  | [readonly]
**ReadinessRefreshEligible** | **bool** |  | [readonly]
**ReadinessStatus** | [**ReadinessStatusEnum**](ReadinessStatusEnum.md) |  | [readonly]
**ActivationEligible** | **bool** |  | [readonly]
**ActivationEligibleAt** | **time.Time** |  | [readonly]
**VerifiedAt** | **time.Time** |  | [readonly]
**ExpiresAt** | **time.Time** |  | [readonly]
**ActivatedAt** | **NullableTime** |  | [readonly]
**ObservedBalanceAtomic** | **NullableString** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]

## Methods

### NewExternalReceivingAddress

`func NewExternalReceivingAddress(id string, walletId string, walletVersionId string, label string, network string, assetId string, address string, status string, proofMethod NullableExternalAddressProofInputMethodEnum, proofVerifiedAt NullableTime, readinessState string, readinessUsable bool, readinessRefreshEligible bool, readinessStatus ReadinessStatusEnum, activationEligible bool, activationEligibleAt time.Time, verifiedAt time.Time, expiresAt time.Time, activatedAt NullableTime, observedBalanceAtomic NullableString, createdAt time.Time, ) *ExternalReceivingAddress`

NewExternalReceivingAddress instantiates a new ExternalReceivingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalReceivingAddressWithDefaults

`func NewExternalReceivingAddressWithDefaults() *ExternalReceivingAddress`

NewExternalReceivingAddressWithDefaults instantiates a new ExternalReceivingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalReceivingAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalReceivingAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalReceivingAddress) SetId(v string)`

SetId sets Id field to given value.


### GetWalletId

`func (o *ExternalReceivingAddress) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExternalReceivingAddress) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExternalReceivingAddress) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletVersionId

`func (o *ExternalReceivingAddress) GetWalletVersionId() string`

GetWalletVersionId returns the WalletVersionId field if non-nil, zero value otherwise.

### GetWalletVersionIdOk

`func (o *ExternalReceivingAddress) GetWalletVersionIdOk() (*string, bool)`

GetWalletVersionIdOk returns a tuple with the WalletVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionId

`func (o *ExternalReceivingAddress) SetWalletVersionId(v string)`

SetWalletVersionId sets WalletVersionId field to given value.


### GetLabel

`func (o *ExternalReceivingAddress) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalReceivingAddress) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalReceivingAddress) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNetwork

`func (o *ExternalReceivingAddress) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExternalReceivingAddress) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExternalReceivingAddress) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *ExternalReceivingAddress) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ExternalReceivingAddress) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ExternalReceivingAddress) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAddress

`func (o *ExternalReceivingAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExternalReceivingAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExternalReceivingAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetStatus

`func (o *ExternalReceivingAddress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalReceivingAddress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalReceivingAddress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProofMethod

`func (o *ExternalReceivingAddress) GetProofMethod() ExternalAddressProofInputMethodEnum`

GetProofMethod returns the ProofMethod field if non-nil, zero value otherwise.

### GetProofMethodOk

`func (o *ExternalReceivingAddress) GetProofMethodOk() (*ExternalAddressProofInputMethodEnum, bool)`

GetProofMethodOk returns a tuple with the ProofMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofMethod

`func (o *ExternalReceivingAddress) SetProofMethod(v ExternalAddressProofInputMethodEnum)`

SetProofMethod sets ProofMethod field to given value.


### SetProofMethodNil

`func (o *ExternalReceivingAddress) SetProofMethodNil(b bool)`

 SetProofMethodNil sets the value for ProofMethod to be an explicit nil

### UnsetProofMethod
`func (o *ExternalReceivingAddress) UnsetProofMethod()`

UnsetProofMethod ensures that no value is present for ProofMethod, not even an explicit nil
### GetProofVerifiedAt

`func (o *ExternalReceivingAddress) GetProofVerifiedAt() time.Time`

GetProofVerifiedAt returns the ProofVerifiedAt field if non-nil, zero value otherwise.

### GetProofVerifiedAtOk

`func (o *ExternalReceivingAddress) GetProofVerifiedAtOk() (*time.Time, bool)`

GetProofVerifiedAtOk returns a tuple with the ProofVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofVerifiedAt

`func (o *ExternalReceivingAddress) SetProofVerifiedAt(v time.Time)`

SetProofVerifiedAt sets ProofVerifiedAt field to given value.


### SetProofVerifiedAtNil

`func (o *ExternalReceivingAddress) SetProofVerifiedAtNil(b bool)`

 SetProofVerifiedAtNil sets the value for ProofVerifiedAt to be an explicit nil

### UnsetProofVerifiedAt
`func (o *ExternalReceivingAddress) UnsetProofVerifiedAt()`

UnsetProofVerifiedAt ensures that no value is present for ProofVerifiedAt, not even an explicit nil
### GetReadinessState

`func (o *ExternalReceivingAddress) GetReadinessState() string`

GetReadinessState returns the ReadinessState field if non-nil, zero value otherwise.

### GetReadinessStateOk

`func (o *ExternalReceivingAddress) GetReadinessStateOk() (*string, bool)`

GetReadinessStateOk returns a tuple with the ReadinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessState

`func (o *ExternalReceivingAddress) SetReadinessState(v string)`

SetReadinessState sets ReadinessState field to given value.


### GetReadinessUsable

`func (o *ExternalReceivingAddress) GetReadinessUsable() bool`

GetReadinessUsable returns the ReadinessUsable field if non-nil, zero value otherwise.

### GetReadinessUsableOk

`func (o *ExternalReceivingAddress) GetReadinessUsableOk() (*bool, bool)`

GetReadinessUsableOk returns a tuple with the ReadinessUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessUsable

`func (o *ExternalReceivingAddress) SetReadinessUsable(v bool)`

SetReadinessUsable sets ReadinessUsable field to given value.


### GetReadinessRefreshEligible

`func (o *ExternalReceivingAddress) GetReadinessRefreshEligible() bool`

GetReadinessRefreshEligible returns the ReadinessRefreshEligible field if non-nil, zero value otherwise.

### GetReadinessRefreshEligibleOk

`func (o *ExternalReceivingAddress) GetReadinessRefreshEligibleOk() (*bool, bool)`

GetReadinessRefreshEligibleOk returns a tuple with the ReadinessRefreshEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessRefreshEligible

`func (o *ExternalReceivingAddress) SetReadinessRefreshEligible(v bool)`

SetReadinessRefreshEligible sets ReadinessRefreshEligible field to given value.


### GetReadinessStatus

`func (o *ExternalReceivingAddress) GetReadinessStatus() ReadinessStatusEnum`

GetReadinessStatus returns the ReadinessStatus field if non-nil, zero value otherwise.

### GetReadinessStatusOk

`func (o *ExternalReceivingAddress) GetReadinessStatusOk() (*ReadinessStatusEnum, bool)`

GetReadinessStatusOk returns a tuple with the ReadinessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessStatus

`func (o *ExternalReceivingAddress) SetReadinessStatus(v ReadinessStatusEnum)`

SetReadinessStatus sets ReadinessStatus field to given value.


### GetActivationEligible

`func (o *ExternalReceivingAddress) GetActivationEligible() bool`

GetActivationEligible returns the ActivationEligible field if non-nil, zero value otherwise.

### GetActivationEligibleOk

`func (o *ExternalReceivingAddress) GetActivationEligibleOk() (*bool, bool)`

GetActivationEligibleOk returns a tuple with the ActivationEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationEligible

`func (o *ExternalReceivingAddress) SetActivationEligible(v bool)`

SetActivationEligible sets ActivationEligible field to given value.


### GetActivationEligibleAt

`func (o *ExternalReceivingAddress) GetActivationEligibleAt() time.Time`

GetActivationEligibleAt returns the ActivationEligibleAt field if non-nil, zero value otherwise.

### GetActivationEligibleAtOk

`func (o *ExternalReceivingAddress) GetActivationEligibleAtOk() (*time.Time, bool)`

GetActivationEligibleAtOk returns a tuple with the ActivationEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationEligibleAt

`func (o *ExternalReceivingAddress) SetActivationEligibleAt(v time.Time)`

SetActivationEligibleAt sets ActivationEligibleAt field to given value.


### GetVerifiedAt

`func (o *ExternalReceivingAddress) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ExternalReceivingAddress) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ExternalReceivingAddress) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.


### GetExpiresAt

`func (o *ExternalReceivingAddress) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExternalReceivingAddress) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExternalReceivingAddress) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetActivatedAt

`func (o *ExternalReceivingAddress) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *ExternalReceivingAddress) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *ExternalReceivingAddress) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.


### SetActivatedAtNil

`func (o *ExternalReceivingAddress) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *ExternalReceivingAddress) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetObservedBalanceAtomic

`func (o *ExternalReceivingAddress) GetObservedBalanceAtomic() string`

GetObservedBalanceAtomic returns the ObservedBalanceAtomic field if non-nil, zero value otherwise.

### GetObservedBalanceAtomicOk

`func (o *ExternalReceivingAddress) GetObservedBalanceAtomicOk() (*string, bool)`

GetObservedBalanceAtomicOk returns a tuple with the ObservedBalanceAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedBalanceAtomic

`func (o *ExternalReceivingAddress) SetObservedBalanceAtomic(v string)`

SetObservedBalanceAtomic sets ObservedBalanceAtomic field to given value.


### SetObservedBalanceAtomicNil

`func (o *ExternalReceivingAddress) SetObservedBalanceAtomicNil(b bool)`

 SetObservedBalanceAtomicNil sets the value for ObservedBalanceAtomic to be an explicit nil

### UnsetObservedBalanceAtomic
`func (o *ExternalReceivingAddress) UnsetObservedBalanceAtomic()`

UnsetObservedBalanceAtomic ensures that no value is present for ObservedBalanceAtomic, not even an explicit nil
### GetCreatedAt

`func (o *ExternalReceivingAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalReceivingAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalReceivingAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
