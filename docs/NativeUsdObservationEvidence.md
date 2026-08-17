# NativeUsdObservationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  |
**NativeUsdQuoteMicros** | **string** |  |
**ObservedAt** | **time.Time** |  |

## Methods

### NewNativeUsdObservationEvidence

`func NewNativeUsdObservationEvidence(source string, nativeUsdQuoteMicros string, observedAt time.Time, ) *NativeUsdObservationEvidence`

NewNativeUsdObservationEvidence instantiates a new NativeUsdObservationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeUsdObservationEvidenceWithDefaults

`func NewNativeUsdObservationEvidenceWithDefaults() *NativeUsdObservationEvidence`

NewNativeUsdObservationEvidenceWithDefaults instantiates a new NativeUsdObservationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *NativeUsdObservationEvidence) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NativeUsdObservationEvidence) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NativeUsdObservationEvidence) SetSource(v string)`

SetSource sets Source field to given value.


### GetNativeUsdQuoteMicros

`func (o *NativeUsdObservationEvidence) GetNativeUsdQuoteMicros() string`

GetNativeUsdQuoteMicros returns the NativeUsdQuoteMicros field if non-nil, zero value otherwise.

### GetNativeUsdQuoteMicrosOk

`func (o *NativeUsdObservationEvidence) GetNativeUsdQuoteMicrosOk() (*string, bool)`

GetNativeUsdQuoteMicrosOk returns a tuple with the NativeUsdQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeUsdQuoteMicros

`func (o *NativeUsdObservationEvidence) SetNativeUsdQuoteMicros(v string)`

SetNativeUsdQuoteMicros sets NativeUsdQuoteMicros field to given value.


### GetObservedAt

`func (o *NativeUsdObservationEvidence) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *NativeUsdObservationEvidence) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *NativeUsdObservationEvidence) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
