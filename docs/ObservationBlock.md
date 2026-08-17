# ObservationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** |  |
**Hash** | **string** |  |
**Finality** | **string** |  |

## Methods

### NewObservationBlock

`func NewObservationBlock(number string, hash string, finality string, ) *ObservationBlock`

NewObservationBlock instantiates a new ObservationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationBlockWithDefaults

`func NewObservationBlockWithDefaults() *ObservationBlock`

NewObservationBlockWithDefaults instantiates a new ObservationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ObservationBlock) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ObservationBlock) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ObservationBlock) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetHash

`func (o *ObservationBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ObservationBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ObservationBlock) SetHash(v string)`

SetHash sets Hash field to given value.


### GetFinality

`func (o *ObservationBlock) GetFinality() string`

GetFinality returns the Finality field if non-nil, zero value otherwise.

### GetFinalityOk

`func (o *ObservationBlock) GetFinalityOk() (*string, bool)`

GetFinalityOk returns a tuple with the Finality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinality

`func (o *ObservationBlock) SetFinality(v string)`

SetFinality sets Finality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
