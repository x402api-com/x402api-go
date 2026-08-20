# PaymentReadiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**PaymentReadinessStateEnum**](PaymentReadinessStateEnum.md) |  | [readonly]
**AcceptingNewPayments** | **bool** |  | [readonly]
**PausedByTenant** | **bool** |  | [readonly]
**PlatformAvailable** | **bool** |  | [readonly]
**HealthValidUntil** | **time.Time** |  | [readonly]
**ObservedAt** | **time.Time** |  | [readonly]
**TenantStatus** | **string** |  | [readonly]
**TenantAcceptingNewChallenges** | **bool** |  | [readonly]
**GlobalChallengesEnabled** | **bool** |  | [readonly]
**GlobalSettlementEnabled** | **bool** |  | [readonly]
**ControlPlaneReadyForNewChallenges** | **bool** |  | [readonly]
**ControlPlaneReadyForSettlement** | **bool** |  | [readonly]
**ExternalOnboarding** | **interface{}** |  | [readonly]
**Rails** | [**[]PaymentReadinessRail**](PaymentReadinessRail.md) |  | [readonly]
**CanonicalRails** | [**[]CanonicalPaymentReadinessRail**](CanonicalPaymentReadinessRail.md) |  | [readonly]

## Methods

### NewPaymentReadiness

`func NewPaymentReadiness(state PaymentReadinessStateEnum, acceptingNewPayments bool, pausedByTenant bool, platformAvailable bool, healthValidUntil time.Time, observedAt time.Time, tenantStatus string, tenantAcceptingNewChallenges bool, globalChallengesEnabled bool, globalSettlementEnabled bool, controlPlaneReadyForNewChallenges bool, controlPlaneReadyForSettlement bool, externalOnboarding interface{}, rails []PaymentReadinessRail, canonicalRails []CanonicalPaymentReadinessRail, ) *PaymentReadiness`

NewPaymentReadiness instantiates a new PaymentReadiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReadinessWithDefaults

`func NewPaymentReadinessWithDefaults() *PaymentReadiness`

NewPaymentReadinessWithDefaults instantiates a new PaymentReadiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PaymentReadiness) GetState() PaymentReadinessStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaymentReadiness) GetStateOk() (*PaymentReadinessStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaymentReadiness) SetState(v PaymentReadinessStateEnum)`

SetState sets State field to given value.


### GetAcceptingNewPayments

`func (o *PaymentReadiness) GetAcceptingNewPayments() bool`

GetAcceptingNewPayments returns the AcceptingNewPayments field if non-nil, zero value otherwise.

### GetAcceptingNewPaymentsOk

`func (o *PaymentReadiness) GetAcceptingNewPaymentsOk() (*bool, bool)`

GetAcceptingNewPaymentsOk returns a tuple with the AcceptingNewPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptingNewPayments

`func (o *PaymentReadiness) SetAcceptingNewPayments(v bool)`

SetAcceptingNewPayments sets AcceptingNewPayments field to given value.


### GetPausedByTenant

`func (o *PaymentReadiness) GetPausedByTenant() bool`

GetPausedByTenant returns the PausedByTenant field if non-nil, zero value otherwise.

### GetPausedByTenantOk

`func (o *PaymentReadiness) GetPausedByTenantOk() (*bool, bool)`

GetPausedByTenantOk returns a tuple with the PausedByTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedByTenant

`func (o *PaymentReadiness) SetPausedByTenant(v bool)`

SetPausedByTenant sets PausedByTenant field to given value.


### GetPlatformAvailable

`func (o *PaymentReadiness) GetPlatformAvailable() bool`

GetPlatformAvailable returns the PlatformAvailable field if non-nil, zero value otherwise.

### GetPlatformAvailableOk

`func (o *PaymentReadiness) GetPlatformAvailableOk() (*bool, bool)`

GetPlatformAvailableOk returns a tuple with the PlatformAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAvailable

`func (o *PaymentReadiness) SetPlatformAvailable(v bool)`

SetPlatformAvailable sets PlatformAvailable field to given value.


### GetHealthValidUntil

`func (o *PaymentReadiness) GetHealthValidUntil() time.Time`

GetHealthValidUntil returns the HealthValidUntil field if non-nil, zero value otherwise.

### GetHealthValidUntilOk

`func (o *PaymentReadiness) GetHealthValidUntilOk() (*time.Time, bool)`

GetHealthValidUntilOk returns a tuple with the HealthValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthValidUntil

`func (o *PaymentReadiness) SetHealthValidUntil(v time.Time)`

SetHealthValidUntil sets HealthValidUntil field to given value.


### GetObservedAt

`func (o *PaymentReadiness) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *PaymentReadiness) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *PaymentReadiness) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetTenantStatus

`func (o *PaymentReadiness) GetTenantStatus() string`

GetTenantStatus returns the TenantStatus field if non-nil, zero value otherwise.

### GetTenantStatusOk

`func (o *PaymentReadiness) GetTenantStatusOk() (*string, bool)`

GetTenantStatusOk returns a tuple with the TenantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStatus

`func (o *PaymentReadiness) SetTenantStatus(v string)`

SetTenantStatus sets TenantStatus field to given value.


### GetTenantAcceptingNewChallenges

`func (o *PaymentReadiness) GetTenantAcceptingNewChallenges() bool`

GetTenantAcceptingNewChallenges returns the TenantAcceptingNewChallenges field if non-nil, zero value otherwise.

### GetTenantAcceptingNewChallengesOk

`func (o *PaymentReadiness) GetTenantAcceptingNewChallengesOk() (*bool, bool)`

GetTenantAcceptingNewChallengesOk returns a tuple with the TenantAcceptingNewChallenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAcceptingNewChallenges

`func (o *PaymentReadiness) SetTenantAcceptingNewChallenges(v bool)`

SetTenantAcceptingNewChallenges sets TenantAcceptingNewChallenges field to given value.


### GetGlobalChallengesEnabled

`func (o *PaymentReadiness) GetGlobalChallengesEnabled() bool`

GetGlobalChallengesEnabled returns the GlobalChallengesEnabled field if non-nil, zero value otherwise.

### GetGlobalChallengesEnabledOk

`func (o *PaymentReadiness) GetGlobalChallengesEnabledOk() (*bool, bool)`

GetGlobalChallengesEnabledOk returns a tuple with the GlobalChallengesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalChallengesEnabled

`func (o *PaymentReadiness) SetGlobalChallengesEnabled(v bool)`

SetGlobalChallengesEnabled sets GlobalChallengesEnabled field to given value.


### GetGlobalSettlementEnabled

`func (o *PaymentReadiness) GetGlobalSettlementEnabled() bool`

GetGlobalSettlementEnabled returns the GlobalSettlementEnabled field if non-nil, zero value otherwise.

### GetGlobalSettlementEnabledOk

`func (o *PaymentReadiness) GetGlobalSettlementEnabledOk() (*bool, bool)`

GetGlobalSettlementEnabledOk returns a tuple with the GlobalSettlementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSettlementEnabled

`func (o *PaymentReadiness) SetGlobalSettlementEnabled(v bool)`

SetGlobalSettlementEnabled sets GlobalSettlementEnabled field to given value.


### GetControlPlaneReadyForNewChallenges

`func (o *PaymentReadiness) GetControlPlaneReadyForNewChallenges() bool`

GetControlPlaneReadyForNewChallenges returns the ControlPlaneReadyForNewChallenges field if non-nil, zero value otherwise.

### GetControlPlaneReadyForNewChallengesOk

`func (o *PaymentReadiness) GetControlPlaneReadyForNewChallengesOk() (*bool, bool)`

GetControlPlaneReadyForNewChallengesOk returns a tuple with the ControlPlaneReadyForNewChallenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneReadyForNewChallenges

`func (o *PaymentReadiness) SetControlPlaneReadyForNewChallenges(v bool)`

SetControlPlaneReadyForNewChallenges sets ControlPlaneReadyForNewChallenges field to given value.


### GetControlPlaneReadyForSettlement

`func (o *PaymentReadiness) GetControlPlaneReadyForSettlement() bool`

GetControlPlaneReadyForSettlement returns the ControlPlaneReadyForSettlement field if non-nil, zero value otherwise.

### GetControlPlaneReadyForSettlementOk

`func (o *PaymentReadiness) GetControlPlaneReadyForSettlementOk() (*bool, bool)`

GetControlPlaneReadyForSettlementOk returns a tuple with the ControlPlaneReadyForSettlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneReadyForSettlement

`func (o *PaymentReadiness) SetControlPlaneReadyForSettlement(v bool)`

SetControlPlaneReadyForSettlement sets ControlPlaneReadyForSettlement field to given value.


### GetExternalOnboarding

`func (o *PaymentReadiness) GetExternalOnboarding() interface{}`

GetExternalOnboarding returns the ExternalOnboarding field if non-nil, zero value otherwise.

### GetExternalOnboardingOk

`func (o *PaymentReadiness) GetExternalOnboardingOk() (*interface{}, bool)`

GetExternalOnboardingOk returns a tuple with the ExternalOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOnboarding

`func (o *PaymentReadiness) SetExternalOnboarding(v interface{})`

SetExternalOnboarding sets ExternalOnboarding field to given value.


### SetExternalOnboardingNil

`func (o *PaymentReadiness) SetExternalOnboardingNil(b bool)`

 SetExternalOnboardingNil sets the value for ExternalOnboarding to be an explicit nil

### UnsetExternalOnboarding
`func (o *PaymentReadiness) UnsetExternalOnboarding()`

UnsetExternalOnboarding ensures that no value is present for ExternalOnboarding, not even an explicit nil
### GetRails

`func (o *PaymentReadiness) GetRails() []PaymentReadinessRail`

GetRails returns the Rails field if non-nil, zero value otherwise.

### GetRailsOk

`func (o *PaymentReadiness) GetRailsOk() (*[]PaymentReadinessRail, bool)`

GetRailsOk returns a tuple with the Rails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRails

`func (o *PaymentReadiness) SetRails(v []PaymentReadinessRail)`

SetRails sets Rails field to given value.


### GetCanonicalRails

`func (o *PaymentReadiness) GetCanonicalRails() []CanonicalPaymentReadinessRail`

GetCanonicalRails returns the CanonicalRails field if non-nil, zero value otherwise.

### GetCanonicalRailsOk

`func (o *PaymentReadiness) GetCanonicalRailsOk() (*[]CanonicalPaymentReadinessRail, bool)`

GetCanonicalRailsOk returns a tuple with the CanonicalRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalRails

`func (o *PaymentReadiness) SetCanonicalRails(v []CanonicalPaymentReadinessRail)`

SetCanonicalRails sets CanonicalRails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
