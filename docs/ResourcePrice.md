# ResourcePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | [readonly]
**Network** | **string** |  | [readonly]
**ContractAddress** | **string** |  | [readonly]
**DisplayName** | **string** |  | [readonly]
**Symbol** | **string** |  | [readonly]
**Decimals** | **int32** |  | [readonly]
**WalletId** | **string** |  | [readonly]
**WalletVersionId** | **string** |  | [readonly]
**Recipient** | **string** |  | [readonly]
**AmountAtomic** | **string** |  | [readonly]
**ListedAmountAtomic** | **string** |  | [readonly]
**MaxTimeoutSeconds** | **int32** |  | [readonly]

## Methods

### NewResourcePrice

`func NewResourcePrice(assetId string, network string, contractAddress string, displayName string, symbol string, decimals int32, walletId string, walletVersionId string, recipient string, amountAtomic string, listedAmountAtomic string, maxTimeoutSeconds int32, ) *ResourcePrice`

NewResourcePrice instantiates a new ResourcePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePriceWithDefaults

`func NewResourcePriceWithDefaults() *ResourcePrice`

NewResourcePriceWithDefaults instantiates a new ResourcePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *ResourcePrice) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ResourcePrice) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ResourcePrice) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNetwork

`func (o *ResourcePrice) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ResourcePrice) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ResourcePrice) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetContractAddress

`func (o *ResourcePrice) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ResourcePrice) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ResourcePrice) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetDisplayName

`func (o *ResourcePrice) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResourcePrice) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResourcePrice) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSymbol

`func (o *ResourcePrice) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ResourcePrice) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ResourcePrice) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *ResourcePrice) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *ResourcePrice) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *ResourcePrice) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetWalletId

`func (o *ResourcePrice) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ResourcePrice) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ResourcePrice) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletVersionId

`func (o *ResourcePrice) GetWalletVersionId() string`

GetWalletVersionId returns the WalletVersionId field if non-nil, zero value otherwise.

### GetWalletVersionIdOk

`func (o *ResourcePrice) GetWalletVersionIdOk() (*string, bool)`

GetWalletVersionIdOk returns a tuple with the WalletVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionId

`func (o *ResourcePrice) SetWalletVersionId(v string)`

SetWalletVersionId sets WalletVersionId field to given value.


### GetRecipient

`func (o *ResourcePrice) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ResourcePrice) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ResourcePrice) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAmountAtomic

`func (o *ResourcePrice) GetAmountAtomic() string`

GetAmountAtomic returns the AmountAtomic field if non-nil, zero value otherwise.

### GetAmountAtomicOk

`func (o *ResourcePrice) GetAmountAtomicOk() (*string, bool)`

GetAmountAtomicOk returns a tuple with the AmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomic

`func (o *ResourcePrice) SetAmountAtomic(v string)`

SetAmountAtomic sets AmountAtomic field to given value.


### GetListedAmountAtomic

`func (o *ResourcePrice) GetListedAmountAtomic() string`

GetListedAmountAtomic returns the ListedAmountAtomic field if non-nil, zero value otherwise.

### GetListedAmountAtomicOk

`func (o *ResourcePrice) GetListedAmountAtomicOk() (*string, bool)`

GetListedAmountAtomicOk returns a tuple with the ListedAmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAmountAtomic

`func (o *ResourcePrice) SetListedAmountAtomic(v string)`

SetListedAmountAtomic sets ListedAmountAtomic field to given value.


### GetMaxTimeoutSeconds

`func (o *ResourcePrice) GetMaxTimeoutSeconds() int32`

GetMaxTimeoutSeconds returns the MaxTimeoutSeconds field if non-nil, zero value otherwise.

### GetMaxTimeoutSecondsOk

`func (o *ResourcePrice) GetMaxTimeoutSecondsOk() (*int32, bool)`

GetMaxTimeoutSecondsOk returns a tuple with the MaxTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutSeconds

`func (o *ResourcePrice) SetMaxTimeoutSeconds(v int32)`

SetMaxTimeoutSeconds sets MaxTimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
