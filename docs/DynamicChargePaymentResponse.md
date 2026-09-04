# DynamicChargePaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeId** | **string** |  |
**OrderId** | **string** |  |
**PaymentId** | **string** | Durable settlement identifier used by payment and receipt APIs. |
**State** | **string** |  |
**Confirmed** | Pointer to **bool** |  | [optional]
**Finalized** | Pointer to **bool** |  | [optional]
**Payer** | **string** |  |
**Transaction** | **string** |  |
**Network** | **string** |  |
**ErrorReason** | **string** |  |

## Methods

### NewDynamicChargePaymentResponse

`func NewDynamicChargePaymentResponse(chargeId string, orderId string, paymentId string, state string, payer string, transaction string, network string, errorReason string, ) *DynamicChargePaymentResponse`

NewDynamicChargePaymentResponse instantiates a new DynamicChargePaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicChargePaymentResponseWithDefaults

`func NewDynamicChargePaymentResponseWithDefaults() *DynamicChargePaymentResponse`

NewDynamicChargePaymentResponseWithDefaults instantiates a new DynamicChargePaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeId

`func (o *DynamicChargePaymentResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *DynamicChargePaymentResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *DynamicChargePaymentResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetOrderId

`func (o *DynamicChargePaymentResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *DynamicChargePaymentResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *DynamicChargePaymentResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetPaymentId

`func (o *DynamicChargePaymentResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *DynamicChargePaymentResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *DynamicChargePaymentResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetState

`func (o *DynamicChargePaymentResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DynamicChargePaymentResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DynamicChargePaymentResponse) SetState(v string)`

SetState sets State field to given value.


### GetConfirmed

`func (o *DynamicChargePaymentResponse) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *DynamicChargePaymentResponse) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *DynamicChargePaymentResponse) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *DynamicChargePaymentResponse) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetFinalized

`func (o *DynamicChargePaymentResponse) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *DynamicChargePaymentResponse) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *DynamicChargePaymentResponse) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *DynamicChargePaymentResponse) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetPayer

`func (o *DynamicChargePaymentResponse) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *DynamicChargePaymentResponse) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *DynamicChargePaymentResponse) SetPayer(v string)`

SetPayer sets Payer field to given value.


### GetTransaction

`func (o *DynamicChargePaymentResponse) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *DynamicChargePaymentResponse) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *DynamicChargePaymentResponse) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetNetwork

`func (o *DynamicChargePaymentResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DynamicChargePaymentResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DynamicChargePaymentResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetErrorReason

`func (o *DynamicChargePaymentResponse) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *DynamicChargePaymentResponse) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *DynamicChargePaymentResponse) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
