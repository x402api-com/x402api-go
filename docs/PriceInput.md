# PriceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  |
**WalletVersionId** | **string** |  |
**AmountAtomic** | **string** |  |
**MaxTimeoutSeconds** | **int32** |  |

## Methods

### NewPriceInput

`func NewPriceInput(assetId string, walletVersionId string, amountAtomic string, maxTimeoutSeconds int32, ) *PriceInput`

NewPriceInput instantiates a new PriceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceInputWithDefaults

`func NewPriceInputWithDefaults() *PriceInput`

NewPriceInputWithDefaults instantiates a new PriceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *PriceInput) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *PriceInput) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *PriceInput) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetWalletVersionId

`func (o *PriceInput) GetWalletVersionId() string`

GetWalletVersionId returns the WalletVersionId field if non-nil, zero value otherwise.

### GetWalletVersionIdOk

`func (o *PriceInput) GetWalletVersionIdOk() (*string, bool)`

GetWalletVersionIdOk returns a tuple with the WalletVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionId

`func (o *PriceInput) SetWalletVersionId(v string)`

SetWalletVersionId sets WalletVersionId field to given value.


### GetAmountAtomic

`func (o *PriceInput) GetAmountAtomic() string`

GetAmountAtomic returns the AmountAtomic field if non-nil, zero value otherwise.

### GetAmountAtomicOk

`func (o *PriceInput) GetAmountAtomicOk() (*string, bool)`

GetAmountAtomicOk returns a tuple with the AmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomic

`func (o *PriceInput) SetAmountAtomic(v string)`

SetAmountAtomic sets AmountAtomic field to given value.


### GetMaxTimeoutSeconds

`func (o *PriceInput) GetMaxTimeoutSeconds() int32`

GetMaxTimeoutSeconds returns the MaxTimeoutSeconds field if non-nil, zero value otherwise.

### GetMaxTimeoutSecondsOk

`func (o *PriceInput) GetMaxTimeoutSecondsOk() (*int32, bool)`

GetMaxTimeoutSecondsOk returns a tuple with the MaxTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutSeconds

`func (o *PriceInput) SetMaxTimeoutSeconds(v int32)`

SetMaxTimeoutSeconds sets MaxTimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
