# PaymentReceiptStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  |
**State** | **string** |  |
**Confirmed** | **bool** |  |
**Finalized** | **bool** |  |
**ConfirmedAt** | **NullableTime** |  |
**FinalizedAt** | **NullableTime** |  |
**Transaction** | **string** |  |
**Network** | **string** |  |
**ReceiptStatus** | [**ReceiptStatusEnum**](ReceiptStatusEnum.md) |  |

## Methods

### NewPaymentReceiptStatus

`func NewPaymentReceiptStatus(paymentId string, state string, confirmed bool, finalized bool, confirmedAt NullableTime, finalizedAt NullableTime, transaction string, network string, receiptStatus ReceiptStatusEnum, ) *PaymentReceiptStatus`

NewPaymentReceiptStatus instantiates a new PaymentReceiptStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReceiptStatusWithDefaults

`func NewPaymentReceiptStatusWithDefaults() *PaymentReceiptStatus`

NewPaymentReceiptStatusWithDefaults instantiates a new PaymentReceiptStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentReceiptStatus) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentReceiptStatus) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentReceiptStatus) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetState

`func (o *PaymentReceiptStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaymentReceiptStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaymentReceiptStatus) SetState(v string)`

SetState sets State field to given value.


### GetConfirmed

`func (o *PaymentReceiptStatus) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *PaymentReceiptStatus) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *PaymentReceiptStatus) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetFinalized

`func (o *PaymentReceiptStatus) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *PaymentReceiptStatus) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *PaymentReceiptStatus) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.


### GetConfirmedAt

`func (o *PaymentReceiptStatus) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *PaymentReceiptStatus) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *PaymentReceiptStatus) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.


### SetConfirmedAtNil

`func (o *PaymentReceiptStatus) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *PaymentReceiptStatus) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetFinalizedAt

`func (o *PaymentReceiptStatus) GetFinalizedAt() time.Time`

GetFinalizedAt returns the FinalizedAt field if non-nil, zero value otherwise.

### GetFinalizedAtOk

`func (o *PaymentReceiptStatus) GetFinalizedAtOk() (*time.Time, bool)`

GetFinalizedAtOk returns a tuple with the FinalizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedAt

`func (o *PaymentReceiptStatus) SetFinalizedAt(v time.Time)`

SetFinalizedAt sets FinalizedAt field to given value.


### SetFinalizedAtNil

`func (o *PaymentReceiptStatus) SetFinalizedAtNil(b bool)`

 SetFinalizedAtNil sets the value for FinalizedAt to be an explicit nil

### UnsetFinalizedAt
`func (o *PaymentReceiptStatus) UnsetFinalizedAt()`

UnsetFinalizedAt ensures that no value is present for FinalizedAt, not even an explicit nil
### GetTransaction

`func (o *PaymentReceiptStatus) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *PaymentReceiptStatus) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *PaymentReceiptStatus) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetNetwork

`func (o *PaymentReceiptStatus) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PaymentReceiptStatus) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PaymentReceiptStatus) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetReceiptStatus

`func (o *PaymentReceiptStatus) GetReceiptStatus() ReceiptStatusEnum`

GetReceiptStatus returns the ReceiptStatus field if non-nil, zero value otherwise.

### GetReceiptStatusOk

`func (o *PaymentReceiptStatus) GetReceiptStatusOk() (*ReceiptStatusEnum, bool)`

GetReceiptStatusOk returns a tuple with the ReceiptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptStatus

`func (o *PaymentReceiptStatus) SetReceiptStatus(v ReceiptStatusEnum)`

SetReceiptStatus sets ReceiptStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
