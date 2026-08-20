# ApiErrorEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ApiError**](ApiError.md) |  |

## Methods

### NewApiErrorEnvelope

`func NewApiErrorEnvelope(error_ ApiError, ) *ApiErrorEnvelope`

NewApiErrorEnvelope instantiates a new ApiErrorEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorEnvelopeWithDefaults

`func NewApiErrorEnvelopeWithDefaults() *ApiErrorEnvelope`

NewApiErrorEnvelopeWithDefaults instantiates a new ApiErrorEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiErrorEnvelope) GetError() ApiError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiErrorEnvelope) GetErrorOk() (*ApiError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiErrorEnvelope) SetError(v ApiError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
