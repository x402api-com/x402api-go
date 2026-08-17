# TenantPaymentOrderProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Status** | **string** |  |
**BuyerPaymentIdentifier** | **NullableString** |  |
**PaidAt** | **NullableTime** |  |
**FulfilledAt** | **NullableTime** |  |

## Methods

### NewTenantPaymentOrderProjection

`func NewTenantPaymentOrderProjection(id string, status string, buyerPaymentIdentifier NullableString, paidAt NullableTime, fulfilledAt NullableTime, ) *TenantPaymentOrderProjection`

NewTenantPaymentOrderProjection instantiates a new TenantPaymentOrderProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPaymentOrderProjectionWithDefaults

`func NewTenantPaymentOrderProjectionWithDefaults() *TenantPaymentOrderProjection`

NewTenantPaymentOrderProjectionWithDefaults instantiates a new TenantPaymentOrderProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantPaymentOrderProjection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantPaymentOrderProjection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantPaymentOrderProjection) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *TenantPaymentOrderProjection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantPaymentOrderProjection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantPaymentOrderProjection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBuyerPaymentIdentifier

`func (o *TenantPaymentOrderProjection) GetBuyerPaymentIdentifier() string`

GetBuyerPaymentIdentifier returns the BuyerPaymentIdentifier field if non-nil, zero value otherwise.

### GetBuyerPaymentIdentifierOk

`func (o *TenantPaymentOrderProjection) GetBuyerPaymentIdentifierOk() (*string, bool)`

GetBuyerPaymentIdentifierOk returns a tuple with the BuyerPaymentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPaymentIdentifier

`func (o *TenantPaymentOrderProjection) SetBuyerPaymentIdentifier(v string)`

SetBuyerPaymentIdentifier sets BuyerPaymentIdentifier field to given value.


### SetBuyerPaymentIdentifierNil

`func (o *TenantPaymentOrderProjection) SetBuyerPaymentIdentifierNil(b bool)`

 SetBuyerPaymentIdentifierNil sets the value for BuyerPaymentIdentifier to be an explicit nil

### UnsetBuyerPaymentIdentifier
`func (o *TenantPaymentOrderProjection) UnsetBuyerPaymentIdentifier()`

UnsetBuyerPaymentIdentifier ensures that no value is present for BuyerPaymentIdentifier, not even an explicit nil
### GetPaidAt

`func (o *TenantPaymentOrderProjection) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *TenantPaymentOrderProjection) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *TenantPaymentOrderProjection) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.


### SetPaidAtNil

`func (o *TenantPaymentOrderProjection) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *TenantPaymentOrderProjection) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetFulfilledAt

`func (o *TenantPaymentOrderProjection) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *TenantPaymentOrderProjection) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *TenantPaymentOrderProjection) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.


### SetFulfilledAtNil

`func (o *TenantPaymentOrderProjection) SetFulfilledAtNil(b bool)`

 SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil

### UnsetFulfilledAt
`func (o *TenantPaymentOrderProjection) UnsetFulfilledAt()`

UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
