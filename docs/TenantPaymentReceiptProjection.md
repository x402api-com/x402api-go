# TenantPaymentReceiptProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TenantPaymentReceiptProjectionStatusEnum**](TenantPaymentReceiptProjectionStatusEnum.md) |  |
**Id** | **NullableString** |  |
**ReceiptDigest** | **string** |  |
**SigningKeyVersion** | **string** |  |
**CreatedAt** | **NullableTime** |  |

## Methods

### NewTenantPaymentReceiptProjection

`func NewTenantPaymentReceiptProjection(status TenantPaymentReceiptProjectionStatusEnum, id NullableString, receiptDigest string, signingKeyVersion string, createdAt NullableTime, ) *TenantPaymentReceiptProjection`

NewTenantPaymentReceiptProjection instantiates a new TenantPaymentReceiptProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantPaymentReceiptProjectionWithDefaults

`func NewTenantPaymentReceiptProjectionWithDefaults() *TenantPaymentReceiptProjection`

NewTenantPaymentReceiptProjectionWithDefaults instantiates a new TenantPaymentReceiptProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TenantPaymentReceiptProjection) GetStatus() TenantPaymentReceiptProjectionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantPaymentReceiptProjection) GetStatusOk() (*TenantPaymentReceiptProjectionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantPaymentReceiptProjection) SetStatus(v TenantPaymentReceiptProjectionStatusEnum)`

SetStatus sets Status field to given value.


### GetId

`func (o *TenantPaymentReceiptProjection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantPaymentReceiptProjection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantPaymentReceiptProjection) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TenantPaymentReceiptProjection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantPaymentReceiptProjection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetReceiptDigest

`func (o *TenantPaymentReceiptProjection) GetReceiptDigest() string`

GetReceiptDigest returns the ReceiptDigest field if non-nil, zero value otherwise.

### GetReceiptDigestOk

`func (o *TenantPaymentReceiptProjection) GetReceiptDigestOk() (*string, bool)`

GetReceiptDigestOk returns a tuple with the ReceiptDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDigest

`func (o *TenantPaymentReceiptProjection) SetReceiptDigest(v string)`

SetReceiptDigest sets ReceiptDigest field to given value.


### GetSigningKeyVersion

`func (o *TenantPaymentReceiptProjection) GetSigningKeyVersion() string`

GetSigningKeyVersion returns the SigningKeyVersion field if non-nil, zero value otherwise.

### GetSigningKeyVersionOk

`func (o *TenantPaymentReceiptProjection) GetSigningKeyVersionOk() (*string, bool)`

GetSigningKeyVersionOk returns a tuple with the SigningKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyVersion

`func (o *TenantPaymentReceiptProjection) SetSigningKeyVersion(v string)`

SetSigningKeyVersion sets SigningKeyVersion field to given value.


### GetCreatedAt

`func (o *TenantPaymentReceiptProjection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantPaymentReceiptProjection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantPaymentReceiptProjection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *TenantPaymentReceiptProjection) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TenantPaymentReceiptProjection) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
