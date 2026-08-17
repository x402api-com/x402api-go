# NativeFeeObservationEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  |
**NativeFeeAtomic** | **string** |  |
**ObservedAt** | **time.Time** |  |

## Methods

### NewNativeFeeObservationEvidence

`func NewNativeFeeObservationEvidence(source string, nativeFeeAtomic string, observedAt time.Time, ) *NativeFeeObservationEvidence`

NewNativeFeeObservationEvidence instantiates a new NativeFeeObservationEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNativeFeeObservationEvidenceWithDefaults

`func NewNativeFeeObservationEvidenceWithDefaults() *NativeFeeObservationEvidence`

NewNativeFeeObservationEvidenceWithDefaults instantiates a new NativeFeeObservationEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *NativeFeeObservationEvidence) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NativeFeeObservationEvidence) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NativeFeeObservationEvidence) SetSource(v string)`

SetSource sets Source field to given value.


### GetNativeFeeAtomic

`func (o *NativeFeeObservationEvidence) GetNativeFeeAtomic() string`

GetNativeFeeAtomic returns the NativeFeeAtomic field if non-nil, zero value otherwise.

### GetNativeFeeAtomicOk

`func (o *NativeFeeObservationEvidence) GetNativeFeeAtomicOk() (*string, bool)`

GetNativeFeeAtomicOk returns a tuple with the NativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFeeAtomic

`func (o *NativeFeeObservationEvidence) SetNativeFeeAtomic(v string)`

SetNativeFeeAtomic sets NativeFeeAtomic field to given value.


### GetObservedAt

`func (o *NativeFeeObservationEvidence) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *NativeFeeObservationEvidence) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *NativeFeeObservationEvidence) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
