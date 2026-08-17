# ResourceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeMode** | Pointer to [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  | [optional] [default to FEEPOLICYMODEINPUTENUM_BUYER_PAYS]
**QuoteCurrency** | Pointer to [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  | [optional] [default to FEEPOLICYQUOTECURRENCYINPUTENUM_USD]
**FeeAllowanceCapQuoteMicros** | Pointer to **string** |  | [optional] [default to "0"]
**Key** | **string** |  |
**Name** | **string** |  |
**Method** | [**HTTPMethodEnum**](HTTPMethodEnum.md) |  |
**Path** | **string** |  |
**Description** | **string** |  |
**MimeType** | Pointer to **string** |  | [optional] [default to "application/json"]
**FulfillmentMode** | [**ResourceInputFulfillmentModeEnum**](ResourceInputFulfillmentModeEnum.md) |  |
**FulfillmentConfig** | Pointer to [**ResourceCreateFulfillmentConfig**](ResourceCreateFulfillmentConfig.md) |  | [optional]
**Prices** | [**[]PriceInput**](PriceInput.md) |  |

## Methods

### NewResourceCreate

`func NewResourceCreate(key string, name string, method HTTPMethodEnum, path string, description string, fulfillmentMode ResourceInputFulfillmentModeEnum, prices []PriceInput, ) *ResourceCreate`

NewResourceCreate instantiates a new ResourceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCreateWithDefaults

`func NewResourceCreateWithDefaults() *ResourceCreate`

NewResourceCreateWithDefaults instantiates a new ResourceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeMode

`func (o *ResourceCreate) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *ResourceCreate) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *ResourceCreate) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.

### HasFeeMode

`func (o *ResourceCreate) HasFeeMode() bool`

HasFeeMode returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *ResourceCreate) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *ResourceCreate) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *ResourceCreate) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *ResourceCreate) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetFeeAllowanceCapQuoteMicros

`func (o *ResourceCreate) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *ResourceCreate) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *ResourceCreate) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.

### HasFeeAllowanceCapQuoteMicros

`func (o *ResourceCreate) HasFeeAllowanceCapQuoteMicros() bool`

HasFeeAllowanceCapQuoteMicros returns a boolean if a field has been set.

### GetKey

`func (o *ResourceCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResourceCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResourceCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ResourceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetMethod

`func (o *ResourceCreate) GetMethod() HTTPMethodEnum`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ResourceCreate) GetMethodOk() (*HTTPMethodEnum, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ResourceCreate) SetMethod(v HTTPMethodEnum)`

SetMethod sets Method field to given value.


### GetPath

`func (o *ResourceCreate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceCreate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceCreate) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *ResourceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMimeType

`func (o *ResourceCreate) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResourceCreate) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResourceCreate) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResourceCreate) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFulfillmentMode

`func (o *ResourceCreate) GetFulfillmentMode() ResourceInputFulfillmentModeEnum`

GetFulfillmentMode returns the FulfillmentMode field if non-nil, zero value otherwise.

### GetFulfillmentModeOk

`func (o *ResourceCreate) GetFulfillmentModeOk() (*ResourceInputFulfillmentModeEnum, bool)`

GetFulfillmentModeOk returns a tuple with the FulfillmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMode

`func (o *ResourceCreate) SetFulfillmentMode(v ResourceInputFulfillmentModeEnum)`

SetFulfillmentMode sets FulfillmentMode field to given value.


### GetFulfillmentConfig

`func (o *ResourceCreate) GetFulfillmentConfig() ResourceCreateFulfillmentConfig`

GetFulfillmentConfig returns the FulfillmentConfig field if non-nil, zero value otherwise.

### GetFulfillmentConfigOk

`func (o *ResourceCreate) GetFulfillmentConfigOk() (*ResourceCreateFulfillmentConfig, bool)`

GetFulfillmentConfigOk returns a tuple with the FulfillmentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentConfig

`func (o *ResourceCreate) SetFulfillmentConfig(v ResourceCreateFulfillmentConfig)`

SetFulfillmentConfig sets FulfillmentConfig field to given value.

### HasFulfillmentConfig

`func (o *ResourceCreate) HasFulfillmentConfig() bool`

HasFulfillmentConfig returns a boolean if a field has been set.

### GetPrices

`func (o *ResourceCreate) GetPrices() []PriceInput`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ResourceCreate) GetPricesOk() (*[]PriceInput, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ResourceCreate) SetPrices(v []PriceInput)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
