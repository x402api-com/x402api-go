# NetworkFeePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeMode** | Pointer to [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  | [optional] [default to FEEPOLICYMODEINPUTENUM_BUYER_PAYS]
**QuoteCurrency** | Pointer to [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  | [optional] [default to FEEPOLICYQUOTECURRENCYINPUTENUM_USD]
**FeeAllowanceCapQuoteMicros** | Pointer to **string** |  | [optional] [default to "0"]
**Prices** | [**[]NetworkFeePreviewPrice**](NetworkFeePreviewPrice.md) |  |

## Methods

### NewNetworkFeePreview

`func NewNetworkFeePreview(prices []NetworkFeePreviewPrice, ) *NetworkFeePreview`

NewNetworkFeePreview instantiates a new NetworkFeePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeePreviewWithDefaults

`func NewNetworkFeePreviewWithDefaults() *NetworkFeePreview`

NewNetworkFeePreviewWithDefaults instantiates a new NetworkFeePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeMode

`func (o *NetworkFeePreview) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *NetworkFeePreview) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *NetworkFeePreview) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.

### HasFeeMode

`func (o *NetworkFeePreview) HasFeeMode() bool`

HasFeeMode returns a boolean if a field has been set.

### GetQuoteCurrency

`func (o *NetworkFeePreview) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *NetworkFeePreview) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *NetworkFeePreview) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.

### HasQuoteCurrency

`func (o *NetworkFeePreview) HasQuoteCurrency() bool`

HasQuoteCurrency returns a boolean if a field has been set.

### GetFeeAllowanceCapQuoteMicros

`func (o *NetworkFeePreview) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *NetworkFeePreview) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *NetworkFeePreview) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.

### HasFeeAllowanceCapQuoteMicros

`func (o *NetworkFeePreview) HasFeeAllowanceCapQuoteMicros() bool`

HasFeeAllowanceCapQuoteMicros returns a boolean if a field has been set.

### GetPrices

`func (o *NetworkFeePreview) GetPrices() []NetworkFeePreviewPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *NetworkFeePreview) GetPricesOk() (*[]NetworkFeePreviewPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *NetworkFeePreview) SetPrices(v []NetworkFeePreviewPrice)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
