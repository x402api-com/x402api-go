# BalanceAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  |
**DisplayName** | **string** |  |
**Symbol** | **string** |  |
**ContractAddress** | **string** |  |
**Decimals** | **int32** |  |
**AmountAtomic** | **string** |  |
**Amount** | **string** |  |
**IssuerNative** | **bool** |  |
**ObservedAt** | **time.Time** |  |
**NodeSource** | **string** |  |
**SourceConsensus** | **string** |  |
**Block** | [**ObservationBlock**](ObservationBlock.md) |  |

## Methods

### NewBalanceAsset

`func NewBalanceAsset(assetId string, displayName string, symbol string, contractAddress string, decimals int32, amountAtomic string, amount string, issuerNative bool, observedAt time.Time, nodeSource string, sourceConsensus string, block ObservationBlock, ) *BalanceAsset`

NewBalanceAsset instantiates a new BalanceAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceAssetWithDefaults

`func NewBalanceAssetWithDefaults() *BalanceAsset`

NewBalanceAssetWithDefaults instantiates a new BalanceAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *BalanceAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *BalanceAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *BalanceAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetDisplayName

`func (o *BalanceAsset) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BalanceAsset) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BalanceAsset) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSymbol

`func (o *BalanceAsset) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BalanceAsset) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BalanceAsset) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetContractAddress

`func (o *BalanceAsset) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *BalanceAsset) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *BalanceAsset) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetDecimals

`func (o *BalanceAsset) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *BalanceAsset) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *BalanceAsset) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetAmountAtomic

`func (o *BalanceAsset) GetAmountAtomic() string`

GetAmountAtomic returns the AmountAtomic field if non-nil, zero value otherwise.

### GetAmountAtomicOk

`func (o *BalanceAsset) GetAmountAtomicOk() (*string, bool)`

GetAmountAtomicOk returns a tuple with the AmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAtomic

`func (o *BalanceAsset) SetAmountAtomic(v string)`

SetAmountAtomic sets AmountAtomic field to given value.


### GetAmount

`func (o *BalanceAsset) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BalanceAsset) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BalanceAsset) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIssuerNative

`func (o *BalanceAsset) GetIssuerNative() bool`

GetIssuerNative returns the IssuerNative field if non-nil, zero value otherwise.

### GetIssuerNativeOk

`func (o *BalanceAsset) GetIssuerNativeOk() (*bool, bool)`

GetIssuerNativeOk returns a tuple with the IssuerNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerNative

`func (o *BalanceAsset) SetIssuerNative(v bool)`

SetIssuerNative sets IssuerNative field to given value.


### GetObservedAt

`func (o *BalanceAsset) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *BalanceAsset) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *BalanceAsset) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetNodeSource

`func (o *BalanceAsset) GetNodeSource() string`

GetNodeSource returns the NodeSource field if non-nil, zero value otherwise.

### GetNodeSourceOk

`func (o *BalanceAsset) GetNodeSourceOk() (*string, bool)`

GetNodeSourceOk returns a tuple with the NodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSource

`func (o *BalanceAsset) SetNodeSource(v string)`

SetNodeSource sets NodeSource field to given value.


### GetSourceConsensus

`func (o *BalanceAsset) GetSourceConsensus() string`

GetSourceConsensus returns the SourceConsensus field if non-nil, zero value otherwise.

### GetSourceConsensusOk

`func (o *BalanceAsset) GetSourceConsensusOk() (*string, bool)`

GetSourceConsensusOk returns a tuple with the SourceConsensus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConsensus

`func (o *BalanceAsset) SetSourceConsensus(v string)`

SetSourceConsensus sets SourceConsensus field to given value.


### GetBlock

`func (o *BalanceAsset) GetBlock() ObservationBlock`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *BalanceAsset) GetBlockOk() (*ObservationBlock, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *BalanceAsset) SetBlock(v ObservationBlock)`

SetBlock sets Block field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
