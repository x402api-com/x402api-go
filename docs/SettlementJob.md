# SettlementJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**OrderId** | **string** |  | [readonly]
**ReservationId** | **string** |  | [readonly]
**State** | [**SettlementJobStateEnum**](SettlementJobStateEnum.md) |  | [readonly]
**Network** | **string** |  | [readonly]
**TransactionHash** | **string** |  | [readonly]
**OriginalTransactionHash** | **string** |  | [readonly]
**ReplacedByHash** | **string** |  | [readonly]
**GasExecutionState** | **string** |  | [readonly]
**GasExecutionSequence** | **int32** |  | [readonly]
**GasExecutionMaterialDigest** | **string** |  | [readonly]
**GasExecutionObservedAt** | **NullableTime** |  | [readonly]
**Payer** | **string** |  | [readonly]
**LastErrorCode** | **string** |  | [readonly]
**BroadcastAttemptCount** | **int32** |  | [readonly]
**SettlementResult** | **interface{}** |  | [readonly]
**ConfirmedAt** | **NullableTime** |  | [readonly]
**FinalizedAt** | **NullableTime** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]
**UpdatedAt** | **time.Time** |  | [readonly]
**Order** | [**TenantPaymentOrderProjection**](TenantPaymentOrderProjection.md) |  | [readonly]
**Resource** | [**TenantPaymentResourceProjection**](TenantPaymentResourceProjection.md) |  | [readonly]
**Asset** | [**TenantPaymentAssetProjection**](TenantPaymentAssetProjection.md) |  | [readonly]
**Chain** | [**TenantPaymentChainProjection**](TenantPaymentChainProjection.md) |  | [readonly]
**Receipt** | [**TenantPaymentReceiptProjection**](TenantPaymentReceiptProjection.md) |  | [readonly]
**Screening** | [**TenantPaymentScreeningProjection**](TenantPaymentScreeningProjection.md) |  | [readonly]
**Fulfillment** | [**TenantPaymentFulfillmentProjection**](TenantPaymentFulfillmentProjection.md) |  | [readonly]

## Methods

### NewSettlementJob

`func NewSettlementJob(id string, orderId string, reservationId string, state SettlementJobStateEnum, network string, transactionHash string, originalTransactionHash string, replacedByHash string, gasExecutionState string, gasExecutionSequence int32, gasExecutionMaterialDigest string, gasExecutionObservedAt NullableTime, payer string, lastErrorCode string, broadcastAttemptCount int32, settlementResult interface{}, confirmedAt NullableTime, finalizedAt NullableTime, createdAt time.Time, updatedAt time.Time, order TenantPaymentOrderProjection, resource TenantPaymentResourceProjection, asset TenantPaymentAssetProjection, chain TenantPaymentChainProjection, receipt TenantPaymentReceiptProjection, screening TenantPaymentScreeningProjection, fulfillment TenantPaymentFulfillmentProjection, ) *SettlementJob`

NewSettlementJob instantiates a new SettlementJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementJobWithDefaults

`func NewSettlementJobWithDefaults() *SettlementJob`

NewSettlementJobWithDefaults instantiates a new SettlementJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettlementJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettlementJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettlementJob) SetId(v string)`

SetId sets Id field to given value.


### GetOrderId

`func (o *SettlementJob) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SettlementJob) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SettlementJob) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReservationId

`func (o *SettlementJob) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *SettlementJob) GetReservationIdOk() (*string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationId

`func (o *SettlementJob) SetReservationId(v string)`

SetReservationId sets ReservationId field to given value.


### GetState

`func (o *SettlementJob) GetState() SettlementJobStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SettlementJob) GetStateOk() (*SettlementJobStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SettlementJob) SetState(v SettlementJobStateEnum)`

SetState sets State field to given value.


### GetNetwork

`func (o *SettlementJob) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SettlementJob) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SettlementJob) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionHash

`func (o *SettlementJob) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *SettlementJob) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *SettlementJob) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetOriginalTransactionHash

`func (o *SettlementJob) GetOriginalTransactionHash() string`

GetOriginalTransactionHash returns the OriginalTransactionHash field if non-nil, zero value otherwise.

### GetOriginalTransactionHashOk

`func (o *SettlementJob) GetOriginalTransactionHashOk() (*string, bool)`

GetOriginalTransactionHashOk returns a tuple with the OriginalTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionHash

`func (o *SettlementJob) SetOriginalTransactionHash(v string)`

SetOriginalTransactionHash sets OriginalTransactionHash field to given value.


### GetReplacedByHash

`func (o *SettlementJob) GetReplacedByHash() string`

GetReplacedByHash returns the ReplacedByHash field if non-nil, zero value otherwise.

### GetReplacedByHashOk

`func (o *SettlementJob) GetReplacedByHashOk() (*string, bool)`

GetReplacedByHashOk returns a tuple with the ReplacedByHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedByHash

`func (o *SettlementJob) SetReplacedByHash(v string)`

SetReplacedByHash sets ReplacedByHash field to given value.


### GetGasExecutionState

`func (o *SettlementJob) GetGasExecutionState() string`

GetGasExecutionState returns the GasExecutionState field if non-nil, zero value otherwise.

### GetGasExecutionStateOk

`func (o *SettlementJob) GetGasExecutionStateOk() (*string, bool)`

GetGasExecutionStateOk returns a tuple with the GasExecutionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasExecutionState

`func (o *SettlementJob) SetGasExecutionState(v string)`

SetGasExecutionState sets GasExecutionState field to given value.


### GetGasExecutionSequence

`func (o *SettlementJob) GetGasExecutionSequence() int32`

GetGasExecutionSequence returns the GasExecutionSequence field if non-nil, zero value otherwise.

### GetGasExecutionSequenceOk

`func (o *SettlementJob) GetGasExecutionSequenceOk() (*int32, bool)`

GetGasExecutionSequenceOk returns a tuple with the GasExecutionSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasExecutionSequence

`func (o *SettlementJob) SetGasExecutionSequence(v int32)`

SetGasExecutionSequence sets GasExecutionSequence field to given value.


### GetGasExecutionMaterialDigest

`func (o *SettlementJob) GetGasExecutionMaterialDigest() string`

GetGasExecutionMaterialDigest returns the GasExecutionMaterialDigest field if non-nil, zero value otherwise.

### GetGasExecutionMaterialDigestOk

`func (o *SettlementJob) GetGasExecutionMaterialDigestOk() (*string, bool)`

GetGasExecutionMaterialDigestOk returns a tuple with the GasExecutionMaterialDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasExecutionMaterialDigest

`func (o *SettlementJob) SetGasExecutionMaterialDigest(v string)`

SetGasExecutionMaterialDigest sets GasExecutionMaterialDigest field to given value.


### GetGasExecutionObservedAt

`func (o *SettlementJob) GetGasExecutionObservedAt() time.Time`

GetGasExecutionObservedAt returns the GasExecutionObservedAt field if non-nil, zero value otherwise.

### GetGasExecutionObservedAtOk

`func (o *SettlementJob) GetGasExecutionObservedAtOk() (*time.Time, bool)`

GetGasExecutionObservedAtOk returns a tuple with the GasExecutionObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasExecutionObservedAt

`func (o *SettlementJob) SetGasExecutionObservedAt(v time.Time)`

SetGasExecutionObservedAt sets GasExecutionObservedAt field to given value.


### SetGasExecutionObservedAtNil

`func (o *SettlementJob) SetGasExecutionObservedAtNil(b bool)`

 SetGasExecutionObservedAtNil sets the value for GasExecutionObservedAt to be an explicit nil

### UnsetGasExecutionObservedAt
`func (o *SettlementJob) UnsetGasExecutionObservedAt()`

UnsetGasExecutionObservedAt ensures that no value is present for GasExecutionObservedAt, not even an explicit nil
### GetPayer

`func (o *SettlementJob) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *SettlementJob) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *SettlementJob) SetPayer(v string)`

SetPayer sets Payer field to given value.


### GetLastErrorCode

`func (o *SettlementJob) GetLastErrorCode() string`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *SettlementJob) GetLastErrorCodeOk() (*string, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *SettlementJob) SetLastErrorCode(v string)`

SetLastErrorCode sets LastErrorCode field to given value.


### GetBroadcastAttemptCount

`func (o *SettlementJob) GetBroadcastAttemptCount() int32`

GetBroadcastAttemptCount returns the BroadcastAttemptCount field if non-nil, zero value otherwise.

### GetBroadcastAttemptCountOk

`func (o *SettlementJob) GetBroadcastAttemptCountOk() (*int32, bool)`

GetBroadcastAttemptCountOk returns a tuple with the BroadcastAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastAttemptCount

`func (o *SettlementJob) SetBroadcastAttemptCount(v int32)`

SetBroadcastAttemptCount sets BroadcastAttemptCount field to given value.


### GetSettlementResult

`func (o *SettlementJob) GetSettlementResult() interface{}`

GetSettlementResult returns the SettlementResult field if non-nil, zero value otherwise.

### GetSettlementResultOk

`func (o *SettlementJob) GetSettlementResultOk() (*interface{}, bool)`

GetSettlementResultOk returns a tuple with the SettlementResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementResult

`func (o *SettlementJob) SetSettlementResult(v interface{})`

SetSettlementResult sets SettlementResult field to given value.


### SetSettlementResultNil

`func (o *SettlementJob) SetSettlementResultNil(b bool)`

 SetSettlementResultNil sets the value for SettlementResult to be an explicit nil

### UnsetSettlementResult
`func (o *SettlementJob) UnsetSettlementResult()`

UnsetSettlementResult ensures that no value is present for SettlementResult, not even an explicit nil
### GetConfirmedAt

`func (o *SettlementJob) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *SettlementJob) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *SettlementJob) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.


### SetConfirmedAtNil

`func (o *SettlementJob) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *SettlementJob) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetFinalizedAt

`func (o *SettlementJob) GetFinalizedAt() time.Time`

GetFinalizedAt returns the FinalizedAt field if non-nil, zero value otherwise.

### GetFinalizedAtOk

`func (o *SettlementJob) GetFinalizedAtOk() (*time.Time, bool)`

GetFinalizedAtOk returns a tuple with the FinalizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedAt

`func (o *SettlementJob) SetFinalizedAt(v time.Time)`

SetFinalizedAt sets FinalizedAt field to given value.


### SetFinalizedAtNil

`func (o *SettlementJob) SetFinalizedAtNil(b bool)`

 SetFinalizedAtNil sets the value for FinalizedAt to be an explicit nil

### UnsetFinalizedAt
`func (o *SettlementJob) UnsetFinalizedAt()`

UnsetFinalizedAt ensures that no value is present for FinalizedAt, not even an explicit nil
### GetCreatedAt

`func (o *SettlementJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettlementJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettlementJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SettlementJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettlementJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettlementJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrder

`func (o *SettlementJob) GetOrder() TenantPaymentOrderProjection`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettlementJob) GetOrderOk() (*TenantPaymentOrderProjection, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettlementJob) SetOrder(v TenantPaymentOrderProjection)`

SetOrder sets Order field to given value.


### GetResource

`func (o *SettlementJob) GetResource() TenantPaymentResourceProjection`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SettlementJob) GetResourceOk() (*TenantPaymentResourceProjection, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SettlementJob) SetResource(v TenantPaymentResourceProjection)`

SetResource sets Resource field to given value.


### GetAsset

`func (o *SettlementJob) GetAsset() TenantPaymentAssetProjection`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SettlementJob) GetAssetOk() (*TenantPaymentAssetProjection, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SettlementJob) SetAsset(v TenantPaymentAssetProjection)`

SetAsset sets Asset field to given value.


### GetChain

`func (o *SettlementJob) GetChain() TenantPaymentChainProjection`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *SettlementJob) GetChainOk() (*TenantPaymentChainProjection, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *SettlementJob) SetChain(v TenantPaymentChainProjection)`

SetChain sets Chain field to given value.


### GetReceipt

`func (o *SettlementJob) GetReceipt() TenantPaymentReceiptProjection`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *SettlementJob) GetReceiptOk() (*TenantPaymentReceiptProjection, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *SettlementJob) SetReceipt(v TenantPaymentReceiptProjection)`

SetReceipt sets Receipt field to given value.


### GetScreening

`func (o *SettlementJob) GetScreening() TenantPaymentScreeningProjection`

GetScreening returns the Screening field if non-nil, zero value otherwise.

### GetScreeningOk

`func (o *SettlementJob) GetScreeningOk() (*TenantPaymentScreeningProjection, bool)`

GetScreeningOk returns a tuple with the Screening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreening

`func (o *SettlementJob) SetScreening(v TenantPaymentScreeningProjection)`

SetScreening sets Screening field to given value.


### GetFulfillment

`func (o *SettlementJob) GetFulfillment() TenantPaymentFulfillmentProjection`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *SettlementJob) GetFulfillmentOk() (*TenantPaymentFulfillmentProjection, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *SettlementJob) SetFulfillment(v TenantPaymentFulfillmentProjection)`

SetFulfillment sets Fulfillment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
