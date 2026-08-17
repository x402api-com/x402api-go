# ResourceVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeMode** | Pointer to [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  | [optional] [default to FEEPOLICYMODEINPUTENUM_BUYER_PAYS]
**QuoteCurrency** | Pointer to [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  | [optional] [default to FEEPOLICYQUOTECURRENCYINPUTENUM_USD]
**FeeAllowanceCapQuoteMicros** | Pointer to **string** |  | [optional] [default to "0"]
**ExpectedLatestVersion** | **int32** |  |
**Method** | [**HTTPMethodEnum**](HTTPMethodEnum.md) |  |
**Path** | **string** |  |
**Description** | **string** |  |
**MimeType** | Pointer to **string** |  | [optional] [default to "application/json"]
**FulfillmentMode** | [**ResourceInputFulfillmentModeEnum**](ResourceInputFulfillmentModeEnum.md) |  |
**FulfillmentConfig** | Pointer to [**ResourceCreateFulfillmentConfig**](ResourceCreateFulfillmentConfig.md) |  | [optional]
**Prices** | [**[]PriceInput**](PriceInput.md) |  |

## Methods

### NewResourceVersionCreate

`func NewResourceVersionCreate(expectedLatestVersion int32, method HTTPMethodEnum, path string, description string, fulfillmentMode ResourceInputFulfillmentModeEnum, prices []PriceInput, ) *ResourceVersionCreate`

NewResourceVersionCreate instantiates a new ResourceVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionCreateWithDefaults

`func NewResourceVersionCreateWithDefaults() *ResourceVersionCreate`

NewResourceVersionCreateWithDefaults instantiates a new ResourceVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeMode

`func (o *ResourceVersionCreate) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *ResourceVersionCreate) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *ResourceVersionCreate) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.

### HasFeeMode

`func (o *ResourceVersionCreate) HasFeeMode() bool`

HasFeeMode returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *ResourceVersionCreate) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *ResourceVersionCreate) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *ResourceVersionCreate) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *ResourceVersionCreate) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetFeeAllowanceCapQuoteMicros

`func (o *ResourceVersionCreate) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *ResourceVersionCreate) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *ResourceVersionCreate) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.

### HasFeeAllowanceCapQuoteMicros

`func (o *ResourceVersionCreate) HasFeeAllowanceCapQuoteMicros() bool`

HasFeeAllowanceCapQuoteMicros returns a boolean if a field has been set.

### GetExpectedLatestVersion

`func (o *ResourceVersionCreate) GetExpectedLatestVersion() int32`

GetExpectedLatestVersion returns the ExpectedLatestVersion field if non-nil, zero value otherwise.

### GetExpectedLatestVersionOk

`func (o *ResourceVersionCreate) GetExpectedLatestVersionOk() (*int32, bool)`

GetExpectedLatestVersionOk returns a tuple with the ExpectedLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLatestVersion

`func (o *ResourceVersionCreate) SetExpectedLatestVersion(v int32)`

SetExpectedLatestVersion sets ExpectedLatestVersion field to given value.


### GetMethod

`func (o *ResourceVersionCreate) GetMethod() HTTPMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ResourceVersionCreate) GetMethodOk() (*HTTPMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ResourceVersionCreate) SetMethod(v HTTPMethodEnum)`

SetMethod sets Method field to given value.


### GetPath

`func (o *ResourceVersionCreate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceVersionCreate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceVersionCreate) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *ResourceVersionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceVersionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceVersionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMimeType

`func (o *ResourceVersionCreate) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResourceVersionCreate) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResourceVersionCreate) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResourceVersionCreate) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFulfillmentMode

`func (o *ResourceVersionCreate) GetFulfillmentMode() ResourceInputFulfillmentModeEnum`

GetFulfillmentMode returns the FulfillmentMode field if non-nil, zero value otherwise.

### GetFulfillmentModeOk

`func (o *ResourceVersionCreate) GetFulfillmentModeOk() (*ResourceInputFulfillmentModeEnum, bool)`

GetFulfillmentModeOk returns a tuple with the FulfillmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMode

`func (o *ResourceVersionCreate) SetFulfillmentMode(v ResourceInputFulfillmentModeEnum)`

SetFulfillmentMode sets FulfillmentMode field to given value.


### GetFulfillmentConfig

`func (o *ResourceVersionCreate) GetFulfillmentConfig() ResourceCreateFulfillmentConfig`

GetFulfillmentConfig returns the FulfillmentConfig field if non-nil, zero value otherwise.

### GetFulfillmentConfigOk

`func (o *ResourceVersionCreate) GetFulfillmentConfigOk() (*ResourceCreateFulfillmentConfig, bool)`

GetFulfillmentConfigOk returns a tuple with the FulfillmentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentConfig

`func (o *ResourceVersionCreate) SetFulfillmentConfig(v ResourceCreateFulfillmentConfig)`

SetFulfillmentConfig sets FulfillmentConfig field to given value.

### HasFulfillmentConfig

`func (o *ResourceVersionCreate) HasFulfillmentConfig() bool`

HasFulfillmentConfig returns a boolean if a field has been set.

### GetPrices

`func (o *ResourceVersionCreate) GetPrices() []PriceInput`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ResourceVersionCreate) GetPricesOk() (*[]PriceInput, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ResourceVersionCreate) SetPrices(v []PriceInput)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
