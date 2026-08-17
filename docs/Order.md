# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**ResourceVersionId** | **string** |  | [readonly]
**RequestFingerprint** | **string** |  | [readonly]
**PaymentIdentifier** | **string** |  | [readonly]
**BuyerPaymentIdentifier** | **NullableString** |  | [readonly]
**Status** | [**OrderStatusEnum**](OrderStatusEnum.md) |  | [readonly]
**PaidAt** | **NullableTime** |  | [readonly]
**FulfilledAt** | **NullableTime** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]
**UpdatedAt** | **time.Time** |  | [readonly]

## Methods

### NewOrder

`func NewOrder(id string, resourceVersionId string, requestFingerprint string, paymentIdentifier string, buyerPaymentIdentifier NullableString, status OrderStatusEnum, paidAt NullableTime, fulfilledAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.


### GetResourceVersionId

`func (o *Order) GetResourceVersionId() string`

GetResourceVersionId returns the ResourceVersionId field if non-nil, zero value otherwise.

### GetResourceVersionIdOk

`func (o *Order) GetResourceVersionIdOk() (*string, bool)`

GetResourceVersionIdOk returns a tuple with the ResourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersionId

`func (o *Order) SetResourceVersionId(v string)`

SetResourceVersionId sets ResourceVersionId field to given value.


### GetRequestFingerprint

`func (o *Order) GetRequestFingerprint() string`

GetRequestFingerprint returns the RequestFingerprint field if non-nil, zero value otherwise.

### GetRequestFingerprintOk

`func (o *Order) GetRequestFingerprintOk() (*string, bool)`

GetRequestFingerprintOk returns a tuple with the RequestFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFingerprint

`func (o *Order) SetRequestFingerprint(v string)`

SetRequestFingerprint sets RequestFingerprint field to given value.


### GetPaymentIdentifier

`func (o *Order) GetPaymentIdentifier() string`

GetPaymentIdentifier returns the PaymentIdentifier field if non-nil, zero value otherwise.

### GetPaymentIdentifierOk

`func (o *Order) GetPaymentIdentifierOk() (*string, bool)`

GetPaymentIdentifierOk returns a tuple with the PaymentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIdentifier

`func (o *Order) SetPaymentIdentifier(v string)`

SetPaymentIdentifier sets PaymentIdentifier field to given value.


### GetBuyerPaymentIdentifier

`func (o *Order) GetBuyerPaymentIdentifier() string`

GetBuyerPaymentIdentifier returns the BuyerPaymentIdentifier field if non-nil, zero value otherwise.

### GetBuyerPaymentIdentifierOk

`func (o *Order) GetBuyerPaymentIdentifierOk() (*string, bool)`

GetBuyerPaymentIdentifierOk returns a tuple with the BuyerPaymentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPaymentIdentifier

`func (o *Order) SetBuyerPaymentIdentifier(v string)`

SetBuyerPaymentIdentifier sets BuyerPaymentIdentifier field to given value.


### SetBuyerPaymentIdentifierNil

`func (o *Order) SetBuyerPaymentIdentifierNil(b bool)`

 SetBuyerPaymentIdentifierNil sets the value for BuyerPaymentIdentifier to be an explicit nil

### UnsetBuyerPaymentIdentifier
`func (o *Order) UnsetBuyerPaymentIdentifier()`

UnsetBuyerPaymentIdentifier ensures that no value is present for BuyerPaymentIdentifier, not even an explicit nil
### GetStatus

`func (o *Order) GetStatus() OrderStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatusEnum)`

SetStatus sets Status field to given value.


### GetPaidAt

`func (o *Order) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *Order) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *Order) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.


### SetPaidAtNil

`func (o *Order) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *Order) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetFulfilledAt

`func (o *Order) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *Order) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *Order) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.


### SetFulfilledAtNil

`func (o *Order) SetFulfilledAtNil(b bool)`

 SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil

### UnsetFulfilledAt
`func (o *Order) UnsetFulfilledAt()`

UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil
### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
