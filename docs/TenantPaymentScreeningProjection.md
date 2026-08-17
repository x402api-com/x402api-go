# TenantPaymentScreeningProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluatedAt** | **time.Time** |  |
**Buyer** | [**TenantPaymentScreeningSubjectProjection**](TenantPaymentScreeningSubjectProjection.md) |  |
**Recipient** | [**TenantPaymentScreeningSubjectProjection**](TenantPaymentScreeningSubjectProjection.md) |  |

## Methods

### NewTenantPaymentScreeningProjection

`func NewTenantPaymentScreeningProjection(evaluatedAt time.Time, buyer TenantPaymentScreeningSubjectProjection, recipient TenantPaymentScreeningSubjectProjection, ) *TenantPaymentScreeningProjection`

NewTenantPaymentScreeningProjection instantiates a new TenantPaymentScreeningProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPaymentScreeningProjectionWithDefaults

`func NewTenantPaymentScreeningProjectionWithDefaults() *TenantPaymentScreeningProjection`

NewTenantPaymentScreeningProjectionWithDefaults instantiates a new TenantPaymentScreeningProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluatedAt

`func (o *TenantPaymentScreeningProjection) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *TenantPaymentScreeningProjection) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *TenantPaymentScreeningProjection) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetBuyer

`func (o *TenantPaymentScreeningProjection) GetBuyer() TenantPaymentScreeningSubjectProjection`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *TenantPaymentScreeningProjection) GetBuyerOk() (*TenantPaymentScreeningSubjectProjection, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *TenantPaymentScreeningProjection) SetBuyer(v TenantPaymentScreeningSubjectProjection)`

SetBuyer sets Buyer field to given value.


### GetRecipient

`func (o *TenantPaymentScreeningProjection) GetRecipient() TenantPaymentScreeningSubjectProjection`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TenantPaymentScreeningProjection) GetRecipientOk() (*TenantPaymentScreeningSubjectProjection, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TenantPaymentScreeningProjection) SetRecipient(v TenantPaymentScreeningSubjectProjection)`

SetRecipient sets Recipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
