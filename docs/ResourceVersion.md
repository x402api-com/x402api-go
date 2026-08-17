# ResourceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**Version** | **int32** |  | [readonly]
**Method** | **string** |  | [readonly]
**Path** | **string** |  | [readonly]
**Description** | **string** |  | [readonly]
**MimeType** | **string** |  | [readonly]
**FulfillmentMode** | [**ResourceVersionFulfillmentModeEnum**](ResourceVersionFulfillmentModeEnum.md) |  | [readonly]
**FulfillmentConfig** | [**ResourceVersionFulfillmentConfig**](ResourceVersionFulfillmentConfig.md) |  |
**FeeMode** | [**ResourceFeeModeEnum**](ResourceFeeModeEnum.md) |  | [readonly]
**QuoteCurrency** | [**ResourceQuoteCurrencyEnum**](ResourceQuoteCurrencyEnum.md) |  | [readonly]
**FeeAllowanceCapQuoteMicros** | **string** |  | [readonly]
**State** | [**ResourceVersionStateEnum**](ResourceVersionStateEnum.md) |  | [readonly]
**Prices** | [**[]ResourcePrice**](ResourcePrice.md) |  | [readonly]

## Methods

### NewResourceVersion

`func NewResourceVersion(id string, version int32, method string, path string, description string, mimeType string, fulfillmentMode ResourceVersionFulfillmentModeEnum, fulfillmentConfig ResourceVersionFulfillmentConfig, feeMode ResourceFeeModeEnum, quoteCurrency ResourceQuoteCurrencyEnum, feeAllowanceCapQuoteMicros string, state ResourceVersionStateEnum, prices []ResourcePrice, ) *ResourceVersion`

NewResourceVersion instantiates a new ResourceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionWithDefaults

`func NewResourceVersionWithDefaults() *ResourceVersion`

NewResourceVersionWithDefaults instantiates a new ResourceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceVersion) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ResourceVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResourceVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResourceVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetMethod

`func (o *ResourceVersion) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ResourceVersion) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ResourceVersion) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *ResourceVersion) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceVersion) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceVersion) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *ResourceVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceVersion) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMimeType

`func (o *ResourceVersion) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResourceVersion) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResourceVersion) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetFulfillmentMode

`func (o *ResourceVersion) GetFulfillmentMode() ResourceVersionFulfillmentModeEnum`

GetFulfillmentMode returns the FulfillmentMode field if non-nil, zero value otherwise.

### GetFulfillmentModeOk

`func (o *ResourceVersion) GetFulfillmentModeOk() (*ResourceVersionFulfillmentModeEnum, bool)`

GetFulfillmentModeOk returns a tuple with the FulfillmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMode

`func (o *ResourceVersion) SetFulfillmentMode(v ResourceVersionFulfillmentModeEnum)`

SetFulfillmentMode sets FulfillmentMode field to given value.


### GetFulfillmentConfig

`func (o *ResourceVersion) GetFulfillmentConfig() ResourceVersionFulfillmentConfig`

GetFulfillmentConfig returns the FulfillmentConfig field if non-nil, zero value otherwise.

### GetFulfillmentConfigOk

`func (o *ResourceVersion) GetFulfillmentConfigOk() (*ResourceVersionFulfillmentConfig, bool)`

GetFulfillmentConfigOk returns a tuple with the FulfillmentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentConfig

`func (o *ResourceVersion) SetFulfillmentConfig(v ResourceVersionFulfillmentConfig)`

SetFulfillmentConfig sets FulfillmentConfig field to given value.


### GetFeeMode

`func (o *ResourceVersion) GetFeeMode() ResourceFeeModeEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *ResourceVersion) GetFeeModeOk() (*ResourceFeeModeEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *ResourceVersion) SetFeeMode(v ResourceFeeModeEnum)`

SetFeeMode sets FeeMode field to given value.


### GetQuoteCurrency

`func (o *ResourceVersion) GetQuoteCurrency() ResourceQuoteCurrencyEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *ResourceVersion) GetQuoteCurrencyOk() (*ResourceQuoteCurrencyEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *ResourceVersion) SetQuoteCurrency(v ResourceQuoteCurrencyEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetFeeAllowanceCapQuoteMicros

`func (o *ResourceVersion) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *ResourceVersion) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *ResourceVersion) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.


### GetState

`func (o *ResourceVersion) GetState() ResourceVersionStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceVersion) GetStateOk() (*ResourceVersionStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceVersion) SetState(v ResourceVersionStateEnum)`

SetState sets State field to given value.


### GetPrices

`func (o *ResourceVersion) GetPrices() []ResourcePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ResourceVersion) GetPricesOk() (*[]ResourcePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ResourceVersion) SetPrices(v []ResourcePrice)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
