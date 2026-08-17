# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**PublicPaymentId** | **string** |  | [readonly]
**Key** | **string** |  | [readonly]
**Name** | **string** |  | [readonly]
**ActiveVersion** | [**NullableResourceVersion**](ResourceVersion.md) |  | [readonly]
**Versions** | [**[]ResourceVersion**](ResourceVersion.md) |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]
**UpdatedAt** | **time.Time** |  | [readonly]

## Methods

### NewResource

`func NewResource(id string, publicPaymentId string, key string, name string, activeVersion NullableResourceVersion, versions []ResourceVersion, createdAt time.Time, updatedAt time.Time, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### GetPublicPaymentId

`func (o *Resource) GetPublicPaymentId() string`

GetPublicPaymentId returns the PublicPaymentId field if non-nil, zero value otherwise.

### GetPublicPaymentIdOk

`func (o *Resource) GetPublicPaymentIdOk() (*string, bool)`

GetPublicPaymentIdOk returns a tuple with the PublicPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPaymentId

`func (o *Resource) SetPublicPaymentId(v string)`

SetPublicPaymentId sets PublicPaymentId field to given value.


### GetKey

`func (o *Resource) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Resource) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Resource) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.


### GetActiveVersion

`func (o *Resource) GetActiveVersion() ResourceVersion`

GetActiveVersion returns the ActiveVersion field if non-nil, zero value otherwise.

### GetActiveVersionOk

`func (o *Resource) GetActiveVersionOk() (*ResourceVersion, bool)`

GetActiveVersionOk returns a tuple with the ActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVersion

`func (o *Resource) SetActiveVersion(v ResourceVersion)`

SetActiveVersion sets ActiveVersion field to given value.


### SetActiveVersionNil

`func (o *Resource) SetActiveVersionNil(b bool)`

 SetActiveVersionNil sets the value for ActiveVersion to be an explicit nil

### UnsetActiveVersion
`func (o *Resource) UnsetActiveVersion()`

UnsetActiveVersion ensures that no value is present for ActiveVersion, not even an explicit nil
### GetVersions

`func (o *Resource) GetVersions() []ResourceVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Resource) GetVersionsOk() (*[]ResourceVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Resource) SetVersions(v []ResourceVersion)`

SetVersions sets Versions field to given value.


### GetCreatedAt

`func (o *Resource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Resource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Resource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Resource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Resource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Resource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
