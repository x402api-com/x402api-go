# PublicFeePolicyDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Version** | **int32** |  |
**FeeMode** | [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  |
**QuoteCurrency** | [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  |

## Methods

### NewPublicFeePolicyDocument

`func NewPublicFeePolicyDocument(type_ string, version int32, feeMode FeePolicyModeInputEnum, quoteCurrency FeePolicyQuoteCurrencyInputEnum, ) *PublicFeePolicyDocument`

NewPublicFeePolicyDocument instantiates a new PublicFeePolicyDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFeePolicyDocumentWithDefaults

`func NewPublicFeePolicyDocumentWithDefaults() *PublicFeePolicyDocument`

NewPublicFeePolicyDocumentWithDefaults instantiates a new PublicFeePolicyDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicFeePolicyDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicFeePolicyDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicFeePolicyDocument) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *PublicFeePolicyDocument) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PublicFeePolicyDocument) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PublicFeePolicyDocument) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFeeMode

`func (o *PublicFeePolicyDocument) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *PublicFeePolicyDocument) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *PublicFeePolicyDocument) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.


### GetQuoteCurrency

`func (o *PublicFeePolicyDocument) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *PublicFeePolicyDocument) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *PublicFeePolicyDocument) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
