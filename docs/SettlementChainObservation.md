# SettlementChainObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**SettlementJobId** | **string** |  | [readonly]
**Network** | **string** |  | [readonly]
**TransactionHash** | **string** |  | [readonly]
**State** | [**SettlementChainObservationStateEnum**](SettlementChainObservationStateEnum.md) |  | [readonly]
**ObservationDigest** | **string** |  | [readonly]
**LogIndex** | **NullableInt32** |  | [readonly]
**BlockNumber** | **NullableString** |  | [readonly]
**BlockHash** | **string** |  | [readonly]
**AssetContract** | **string** |  | [readonly]
**Payer** | **string** |  | [readonly]
**Recipient** | **string** |  | [readonly]
**AmountAtomic** | **NullableString** |  | [readonly]
**ExecutionSuccess** | **NullableBool** |  | [readonly]
**ObservedAt** | **time.Time** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]

## Methods

### NewSettlementChainObservation

`func NewSettlementChainObservation(id string, settlementJobId string, network string, transactionHash string, state SettlementChainObservationStateEnum, observationDigest string, logIndex NullableInt32, blockNumber NullableString, blockHash string, assetContract string, payer string, recipient string, amountAtomic NullableString, executionSuccess NullableBool, observedAt time.Time, createdAt time.Time, ) *SettlementChainObservation`

NewSettlementChainObservation instantiates a new SettlementChainObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementChainObservationWithDefaults

`func NewSettlementChainObservationWithDefaults() *SettlementChainObservation`

NewSettlementChainObservationWithDefaults instantiates a new SettlementChainObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettlementChainObservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettlementChainObservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettlementChainObservation) SetId(v string)`

SetId sets Id field to given value.


### GetSettlementJobId

`func (o *SettlementChainObservation) GetSettlementJobId() string`

GetSettlementJobId returns the SettlementJobId field if non-nil, zero value otherwise.

### GetSettlementJobIdOk

`func (o *SettlementChainObservation) GetSettlementJobIdOk() (*string, bool)`

GetSettlementJobIdOk returns a tuple with the SettlementJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementJobId

`func (o *SettlementChainObservation) SetSettlementJobId(v string)`

SetSettlementJobId sets SettlementJobId field to given value.


### GetNetwork

`func (o *SettlementChainObservation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SettlementChainObservation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SettlementChainObservation) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTransactionHash

`func (o *SettlementChainObservation) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *SettlementChainObservation) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *SettlementChainObservation) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetState

`func (o *SettlementChainObservation) GetState() SettlementChainObservationStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SettlementChainObservation) GetStateOk() (*SettlementChainObservationStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SettlementChainObservation) SetState(v SettlementChainObservationStateEnum)`

SetState sets State field to given value.


### GetObservationDigest

`func (o *SettlementChainObservation) GetObservationDigest() string`

GetObservationDigest returns the ObservationDigest field if non-nil, zero value otherwise.

### GetObservationDigestOk

`func (o *SettlementChainObservation) GetObservationDigestOk() (*string, bool)`

GetObservationDigestOk returns a tuple with the ObservationDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationDigest

`func (o *SettlementChainObservation) SetObservationDigest(v string)`

SetObservationDigest sets ObservationDigest field to given value.


### GetLogIndex

`func (o *SettlementChainObservation) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *SettlementChainObservation) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *SettlementChainObservation) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.


### SetLogIndexNil

`func (o *SettlementChainObservation) SetLogIndexNil(b bool)`

 SetLogIndexNil sets the value for LogIndex to be an explicit nil

### UnsetLogIndex
`func (o *SettlementChainObservation) UnsetLogIndex()`

UnsetLogIndex ensures that no value is present for LogIndex, not even an explicit nil
### GetBlockNumber

`func (o *SettlementChainObservation) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *SettlementChainObservation) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *SettlementChainObservation) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### SetBlockNumberNil

`func (o *SettlementChainObservation) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *SettlementChainObservation) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
### GetBlockHash

`func (o *SettlementChainObservation) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *SettlementChainObservation) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *SettlementChainObservation) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetAssetContract

`func (o *SettlementChainObservation) GetAssetContract() string`

GetAssetContract returns the AssetContract field if non-nil, zero value otherwise.

### GetAssetContractOk

`func (o *SettlementChainObservation) GetAssetContractOk() (*string, bool)`

GetAssetContractOk returns a tuple with the AssetContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetContract

`func (o *SettlementChainObservation) SetAssetContract(v string)`

SetAssetContract sets AssetContract field to given value.


### GetPayer

`func (o *SettlementChainObservation) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *SettlementChainObservation) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *SettlementChainObservation) SetPayer(v string)`

SetPayer sets Payer field to given value.


### GetRecipient

`func (o *SettlementChainObservation) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SettlementChainObservation) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SettlementChainObservation) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAmountAtomic

`func (o *SettlementChainObservation) GetAmountAtomic() string`

GetAmountAtomic returns the AmountAtomic field if non-nil, zero value otherwise.

### GetAmountAtomicOk

`func (o *SettlementChainObservation) GetAmountAtomicOk() (*string, bool)`

GetAmountAtomicOk returns a tuple with the AmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomic

`func (o *SettlementChainObservation) SetAmountAtomic(v string)`

SetAmountAtomic sets AmountAtomic field to given value.


### SetAmountAtomicNil

`func (o *SettlementChainObservation) SetAmountAtomicNil(b bool)`

 SetAmountAtomicNil sets the value for AmountAtomic to be an explicit nil

### UnsetAmountAtomic
`func (o *SettlementChainObservation) UnsetAmountAtomic()`

UnsetAmountAtomic ensures that no value is present for AmountAtomic, not even an explicit nil
### GetExecutionSuccess

`func (o *SettlementChainObservation) GetExecutionSuccess() bool`

GetExecutionSuccess returns the ExecutionSuccess field if non-nil, zero value otherwise.

### GetExecutionSuccessOk

`func (o *SettlementChainObservation) GetExecutionSuccessOk() (*bool, bool)`

GetExecutionSuccessOk returns a tuple with the ExecutionSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionSuccess

`func (o *SettlementChainObservation) SetExecutionSuccess(v bool)`

SetExecutionSuccess sets ExecutionSuccess field to given value.


### SetExecutionSuccessNil

`func (o *SettlementChainObservation) SetExecutionSuccessNil(b bool)`

 SetExecutionSuccessNil sets the value for ExecutionSuccess to be an explicit nil

### UnsetExecutionSuccess
`func (o *SettlementChainObservation) UnsetExecutionSuccess()`

UnsetExecutionSuccess ensures that no value is present for ExecutionSuccess, not even an explicit nil
### GetObservedAt

`func (o *SettlementChainObservation) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *SettlementChainObservation) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *SettlementChainObservation) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetCreatedAt

`func (o *SettlementChainObservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettlementChainObservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettlementChainObservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
