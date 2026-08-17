# ResourceVersionRetire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedVersion** | **int32** |  |
**ExpectedState** | [**ResourceVersionRetireExpectedStateEnum**](ResourceVersionRetireExpectedStateEnum.md) |  |

## Methods

### NewResourceVersionRetire

`func NewResourceVersionRetire(expectedVersion int32, expectedState ResourceVersionRetireExpectedStateEnum, ) *ResourceVersionRetire`

NewResourceVersionRetire instantiates a new ResourceVersionRetire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionRetireWithDefaults

`func NewResourceVersionRetireWithDefaults() *ResourceVersionRetire`

NewResourceVersionRetireWithDefaults instantiates a new ResourceVersionRetire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedVersion

`func (o *ResourceVersionRetire) GetExpectedVersion() int32`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *ResourceVersionRetire) GetExpectedVersionOk() (*int32, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *ResourceVersionRetire) SetExpectedVersion(v int32)`

SetExpectedVersion sets ExpectedVersion field to given value.


### GetExpectedState

`func (o *ResourceVersionRetire) GetExpectedState() ResourceVersionRetireExpectedStateEnum`

GetExpectedState returns the ExpectedState field if non-nil, zero value otherwise.

### GetExpectedStateOk

`func (o *ResourceVersionRetire) GetExpectedStateOk() (*ResourceVersionRetireExpectedStateEnum, bool)`

GetExpectedStateOk returns a tuple with the ExpectedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedState

`func (o *ResourceVersionRetire) SetExpectedState(v ResourceVersionRetireExpectedStateEnum)`

SetExpectedState sets ExpectedState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
