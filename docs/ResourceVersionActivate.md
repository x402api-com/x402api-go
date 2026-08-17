# ResourceVersionActivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedTargetVersion** | **int32** |  |
**ExpectedActiveVersionId** | **NullableString** |  |

## Methods

### NewResourceVersionActivate

`func NewResourceVersionActivate(expectedTargetVersion int32, expectedActiveVersionId NullableString, ) *ResourceVersionActivate`

NewResourceVersionActivate instantiates a new ResourceVersionActivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionActivateWithDefaults

`func NewResourceVersionActivateWithDefaults() *ResourceVersionActivate`

NewResourceVersionActivateWithDefaults instantiates a new ResourceVersionActivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedTargetVersion

`func (o *ResourceVersionActivate) GetExpectedTargetVersion() int32`

GetExpectedTargetVersion returns the ExpectedTargetVersion field if non-nil, zero value otherwise.

### GetExpectedTargetVersionOk

`func (o *ResourceVersionActivate) GetExpectedTargetVersionOk() (*int32, bool)`

GetExpectedTargetVersionOk returns a tuple with the ExpectedTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTargetVersion

`func (o *ResourceVersionActivate) SetExpectedTargetVersion(v int32)`

SetExpectedTargetVersion sets ExpectedTargetVersion field to given value.


### GetExpectedActiveVersionId

`func (o *ResourceVersionActivate) GetExpectedActiveVersionId() string`

GetExpectedActiveVersionId returns the ExpectedActiveVersionId field if non-nil, zero value otherwise.

### GetExpectedActiveVersionIdOk

`func (o *ResourceVersionActivate) GetExpectedActiveVersionIdOk() (*string, bool)`

GetExpectedActiveVersionIdOk returns a tuple with the ExpectedActiveVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedActiveVersionId

`func (o *ResourceVersionActivate) SetExpectedActiveVersionId(v string)`

SetExpectedActiveVersionId sets ExpectedActiveVersionId field to given value.


### SetExpectedActiveVersionIdNil

`func (o *ResourceVersionActivate) SetExpectedActiveVersionIdNil(b bool)`

 SetExpectedActiveVersionIdNil sets the value for ExpectedActiveVersionId to be an explicit nil

### UnsetExpectedActiveVersionId
`func (o *ResourceVersionActivate) UnsetExpectedActiveVersionId()`

UnsetExpectedActiveVersionId ensures that no value is present for ExpectedActiveVersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
