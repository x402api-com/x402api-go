# PaymentReadinessBlocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | [readonly]
**Owner** | [**PaymentReadinessBlockerOwnerEnum**](PaymentReadinessBlockerOwnerEnum.md) |  | [readonly]
**Message** | **string** |  | [readonly]
**ActionUrl** | Pointer to **string** |  | [optional]

## Methods

### NewPaymentReadinessBlocker

`func NewPaymentReadinessBlocker(code string, owner PaymentReadinessBlockerOwnerEnum, message string, ) *PaymentReadinessBlocker`

NewPaymentReadinessBlocker instantiates a new PaymentReadinessBlocker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReadinessBlockerWithDefaults

`func NewPaymentReadinessBlockerWithDefaults() *PaymentReadinessBlocker`

NewPaymentReadinessBlockerWithDefaults instantiates a new PaymentReadinessBlocker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PaymentReadinessBlocker) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentReadinessBlocker) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentReadinessBlocker) SetCode(v string)`

SetCode sets Code field to given value.


### GetOwner

`func (o *PaymentReadinessBlocker) GetOwner() PaymentReadinessBlockerOwnerEnum`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PaymentReadinessBlocker) GetOwnerOk() (*PaymentReadinessBlockerOwnerEnum, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PaymentReadinessBlocker) SetOwner(v PaymentReadinessBlockerOwnerEnum)`

SetOwner sets Owner field to given value.


### GetMessage

`func (o *PaymentReadinessBlocker) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentReadinessBlocker) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentReadinessBlocker) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetActionUrl

`func (o *PaymentReadinessBlocker) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *PaymentReadinessBlocker) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *PaymentReadinessBlocker) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *PaymentReadinessBlocker) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
