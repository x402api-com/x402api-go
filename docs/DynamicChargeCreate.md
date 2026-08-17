# DynamicChargeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersionId** | **string** |  |
**Method** | Pointer to [**HTTPMethodEnum**](HTTPMethodEnum.md) |  | [optional] [default to HTTPMETHODENUM_POST]
**ResourceUrl** | **string** |  |
**BodyBase64** | Pointer to **string** |  | [optional] [default to ""]
**ContentType** | Pointer to **NullableString** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Prices** | [**[]DynamicChargePrice**](DynamicChargePrice.md) |  |
**FeeMode** | Pointer to [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  | [optional]
**QuoteCurrency** | Pointer to [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  | [optional]
**FeeAllowanceCapQuoteMicros** | Pointer to **string** |  | [optional]
**ExpiresInSeconds** | **int32** |  |
**Metadata** | Pointer to **map[string]interface{}** | Tenant application metadata frozen into the charge digest. Maximum canonical size is 16 KiB; floating-point numbers are not accepted. | [optional]

## Methods

### NewDynamicChargeCreate

`func NewDynamicChargeCreate(resourceVersionId string, resourceUrl string, prices []DynamicChargePrice, expiresInSeconds int32, ) *DynamicChargeCreate`

NewDynamicChargeCreate instantiates a new DynamicChargeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicChargeCreateWithDefaults

`func NewDynamicChargeCreateWithDefaults() *DynamicChargeCreate`

NewDynamicChargeCreateWithDefaults instantiates a new DynamicChargeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVersionId

`func (o *DynamicChargeCreate) GetResourceVersionId() string`

GetResourceVersionId returns the ResourceVersionId field if non-nil, zero value otherwise.

### GetResourceVersionIdOk

`func (o *DynamicChargeCreate) GetResourceVersionIdOk() (*string, bool)`

GetResourceVersionIdOk returns a tuple with the ResourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersionId

`func (o *DynamicChargeCreate) SetResourceVersionId(v string)`

SetResourceVersionId sets ResourceVersionId field to given value.


### GetMethod

`func (o *DynamicChargeCreate) GetMethod() HTTPMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *DynamicChargeCreate) GetMethodOk() (*HTTPMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *DynamicChargeCreate) SetMethod(v HTTPMethodEnum)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *DynamicChargeCreate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetResourceUrl

`func (o *DynamicChargeCreate) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *DynamicChargeCreate) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *DynamicChargeCreate) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.


### GetBodyBase64

`func (o *DynamicChargeCreate) GetBodyBase64() string`

GetBodyBase64 returns the BodyBase64 field if non-nil, zero value otherwise.

### GetBodyBase64Ok

`func (o *DynamicChargeCreate) GetBodyBase64Ok() (*string, bool)`

GetBodyBase64Ok returns a tuple with the BodyBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyBase64

`func (o *DynamicChargeCreate) SetBodyBase64(v string)`

SetBodyBase64 sets BodyBase64 field to given value.

### HasBodyBase64

`func (o *DynamicChargeCreate) HasBodyBase64() bool`

HasBodyBase64 returns a boolean if a field has been set.

### GetContentType

`func (o *DynamicChargeCreate) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DynamicChargeCreate) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DynamicChargeCreate) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DynamicChargeCreate) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *DynamicChargeCreate) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *DynamicChargeCreate) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDescription

`func (o *DynamicChargeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicChargeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicChargeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicChargeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrices

`func (o *DynamicChargeCreate) GetPrices() []DynamicChargePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DynamicChargeCreate) GetPricesOk() (*[]DynamicChargePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DynamicChargeCreate) SetPrices(v []DynamicChargePrice)`

SetPrices sets Prices field to given value.


### GetFeeMode

`func (o *DynamicChargeCreate) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *DynamicChargeCreate) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *DynamicChargeCreate) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.

### HasFeeMode

`func (o *DynamicChargeCreate) HasFeeMode() bool`

HasFeeMode returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *DynamicChargeCreate) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *DynamicChargeCreate) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *DynamicChargeCreate) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *DynamicChargeCreate) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetFeeAllowanceCapQuoteMicros

`func (o *DynamicChargeCreate) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *DynamicChargeCreate) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *DynamicChargeCreate) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.

### HasFeeAllowanceCapQuoteMicros

`func (o *DynamicChargeCreate) HasFeeAllowanceCapQuoteMicros() bool`

HasFeeAllowanceCapQuoteMicros returns a boolean if a field has been set.

### GetExpiresInSeconds

`func (o *DynamicChargeCreate) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *DynamicChargeCreate) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *DynamicChargeCreate) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.


### GetMetadata

`func (o *DynamicChargeCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DynamicChargeCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DynamicChargeCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DynamicChargeCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
