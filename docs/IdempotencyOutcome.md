# IdempotencyOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**IdempotencyOutcomeStateEnum**](IdempotencyOutcomeStateEnum.md) |  | [readonly]

## Methods

### NewIdempotencyOutcome

`func NewIdempotencyOutcome(state IdempotencyOutcomeStateEnum, ) *IdempotencyOutcome`

NewIdempotencyOutcome instantiates a new IdempotencyOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdempotencyOutcomeWithDefaults

`func NewIdempotencyOutcomeWithDefaults() *IdempotencyOutcome`

NewIdempotencyOutcomeWithDefaults instantiates a new IdempotencyOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *IdempotencyOutcome) GetState() IdempotencyOutcomeStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdempotencyOutcome) GetStateOk() (*IdempotencyOutcomeStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdempotencyOutcome) SetState(v IdempotencyOutcomeStateEnum)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
