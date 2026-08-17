# TenantPaymentChainProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  |
**TransactionHash** | **string** |  |
**BlockNumber** | **NullableString** |  |
**BlockHash** | **string** |  |
**Confirmations** | **int32** |  |
**ConfirmationsRequired** | **int32** |  |
**ObservedAt** | **NullableTime** |  |

## Methods

### NewTenantPaymentChainProjection

`func NewTenantPaymentChainProjection(state string, transactionHash string, blockNumber NullableString, blockHash string, confirmations int32, confirmationsRequired int32, observedAt NullableTime, ) *TenantPaymentChainProjection`

NewTenantPaymentChainProjection instantiates a new TenantPaymentChainProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPaymentChainProjectionWithDefaults

`func NewTenantPaymentChainProjectionWithDefaults() *TenantPaymentChainProjection`

NewTenantPaymentChainProjectionWithDefaults instantiates a new TenantPaymentChainProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *TenantPaymentChainProjection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TenantPaymentChainProjection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TenantPaymentChainProjection) SetState(v string)`

SetState sets State field to given value.


### GetTransactionHash

`func (o *TenantPaymentChainProjection) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TenantPaymentChainProjection) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TenantPaymentChainProjection) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetBlockNumber

`func (o *TenantPaymentChainProjection) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TenantPaymentChainProjection) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TenantPaymentChainProjection) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### SetBlockNumberNil

`func (o *TenantPaymentChainProjection) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *TenantPaymentChainProjection) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
### GetBlockHash

`func (o *TenantPaymentChainProjection) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TenantPaymentChainProjection) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TenantPaymentChainProjection) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetConfirmations

`func (o *TenantPaymentChainProjection) GetConfirmations() int32`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *TenantPaymentChainProjection) GetConfirmationsOk() (*int32, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *TenantPaymentChainProjection) SetConfirmations(v int32)`

SetConfirmations sets Confirmations field to given value.


### GetConfirmationsRequired

`func (o *TenantPaymentChainProjection) GetConfirmationsRequired() int32`

GetConfirmationsRequired returns the ConfirmationsRequired field if non-nil, zero value otherwise.

### GetConfirmationsRequiredOk

`func (o *TenantPaymentChainProjection) GetConfirmationsRequiredOk() (*int32, bool)`

GetConfirmationsRequiredOk returns a tuple with the ConfirmationsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsRequired

`func (o *TenantPaymentChainProjection) SetConfirmationsRequired(v int32)`

SetConfirmationsRequired sets ConfirmationsRequired field to given value.


### GetObservedAt

`func (o *TenantPaymentChainProjection) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *TenantPaymentChainProjection) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *TenantPaymentChainProjection) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *TenantPaymentChainProjection) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *TenantPaymentChainProjection) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
