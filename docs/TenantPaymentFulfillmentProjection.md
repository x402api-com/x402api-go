# TenantPaymentFulfillmentProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TenantPaymentFulfillmentProjectionStatusEnum**](TenantPaymentFulfillmentProjectionStatusEnum.md) |  |
**Id** | **NullableString** |  |
**Mode** | **string** |  |
**State** | **string** |  |
**AttemptCount** | **int32** |  |
**LastErrorCode** | **string** |  |
**CompletedAt** | **NullableTime** |  |

## Methods

### NewTenantPaymentFulfillmentProjection

`func NewTenantPaymentFulfillmentProjection(status TenantPaymentFulfillmentProjectionStatusEnum, id NullableString, mode string, state string, attemptCount int32, lastErrorCode string, completedAt NullableTime, ) *TenantPaymentFulfillmentProjection`

NewTenantPaymentFulfillmentProjection instantiates a new TenantPaymentFulfillmentProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPaymentFulfillmentProjectionWithDefaults

`func NewTenantPaymentFulfillmentProjectionWithDefaults() *TenantPaymentFulfillmentProjection`

NewTenantPaymentFulfillmentProjectionWithDefaults instantiates a new TenantPaymentFulfillmentProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TenantPaymentFulfillmentProjection) GetStatus() TenantPaymentFulfillmentProjectionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantPaymentFulfillmentProjection) GetStatusOk() (*TenantPaymentFulfillmentProjectionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantPaymentFulfillmentProjection) SetStatus(v TenantPaymentFulfillmentProjectionStatusEnum)`

SetStatus sets Status field to given value.


### GetId

`func (o *TenantPaymentFulfillmentProjection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantPaymentFulfillmentProjection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantPaymentFulfillmentProjection) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TenantPaymentFulfillmentProjection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantPaymentFulfillmentProjection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMode

`func (o *TenantPaymentFulfillmentProjection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TenantPaymentFulfillmentProjection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TenantPaymentFulfillmentProjection) SetMode(v string)`

SetMode sets Mode field to given value.


### GetState

`func (o *TenantPaymentFulfillmentProjection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TenantPaymentFulfillmentProjection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TenantPaymentFulfillmentProjection) SetState(v string)`

SetState sets State field to given value.


### GetAttemptCount

`func (o *TenantPaymentFulfillmentProjection) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *TenantPaymentFulfillmentProjection) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *TenantPaymentFulfillmentProjection) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.


### GetLastErrorCode

`func (o *TenantPaymentFulfillmentProjection) GetLastErrorCode() string`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *TenantPaymentFulfillmentProjection) GetLastErrorCodeOk() (*string, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *TenantPaymentFulfillmentProjection) SetLastErrorCode(v string)`

SetLastErrorCode sets LastErrorCode field to given value.


### GetCompletedAt

`func (o *TenantPaymentFulfillmentProjection) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TenantPaymentFulfillmentProjection) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TenantPaymentFulfillmentProjection) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *TenantPaymentFulfillmentProjection) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *TenantPaymentFulfillmentProjection) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
