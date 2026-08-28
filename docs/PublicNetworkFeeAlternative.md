# PublicNetworkFeeAlternative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Version** | **int32** |  |
**Network** | **string** |  |
**AssetId** | **string** |  |
**ContractAddress** | **string** |  |
**ListedAmountAtomic** | **string** |  |
**GasMode** | [**GasModeEnum**](GasModeEnum.md) |  |
**BuyerNativeFeeAtomic** | **NullableString** |  |
**BuyerPaymentAtomic** | **string** |  |
**TenantProceedsAtomic** | **string** |  |
**QuoteExpiresAt** | **NullableTime** |  |
**Eligible** | **bool** |  |
**ExclusionReason** | **NullableString** |  |

## Methods

### NewPublicNetworkFeeAlternative

`func NewPublicNetworkFeeAlternative(type_ string, version int32, network string, assetId string, contractAddress string, listedAmountAtomic string, gasMode GasModeEnum, buyerNativeFeeAtomic NullableString, buyerPaymentAtomic string, tenantProceedsAtomic string, quoteExpiresAt NullableTime, eligible bool, exclusionReason NullableString, ) *PublicNetworkFeeAlternative`

NewPublicNetworkFeeAlternative instantiates a new PublicNetworkFeeAlternative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkFeeAlternativeWithDefaults

`func NewPublicNetworkFeeAlternativeWithDefaults() *PublicNetworkFeeAlternative`

NewPublicNetworkFeeAlternativeWithDefaults instantiates a new PublicNetworkFeeAlternative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicNetworkFeeAlternative) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicNetworkFeeAlternative) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicNetworkFeeAlternative) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *PublicNetworkFeeAlternative) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PublicNetworkFeeAlternative) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PublicNetworkFeeAlternative) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetNetwork

`func (o *PublicNetworkFeeAlternative) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PublicNetworkFeeAlternative) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PublicNetworkFeeAlternative) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *PublicNetworkFeeAlternative) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *PublicNetworkFeeAlternative) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *PublicNetworkFeeAlternative) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetContractAddress

`func (o *PublicNetworkFeeAlternative) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *PublicNetworkFeeAlternative) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *PublicNetworkFeeAlternative) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetListedAmountAtomic

`func (o *PublicNetworkFeeAlternative) GetListedAmountAtomic() string`

GetListedAmountAtomic returns the ListedAmountAtomic field if non-nil, zero value otherwise.

### GetListedAmountAtomicOk

`func (o *PublicNetworkFeeAlternative) GetListedAmountAtomicOk() (*string, bool)`

GetListedAmountAtomicOk returns a tuple with the ListedAmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAmountAtomic

`func (o *PublicNetworkFeeAlternative) SetListedAmountAtomic(v string)`

SetListedAmountAtomic sets ListedAmountAtomic field to given value.


### GetGasMode

`func (o *PublicNetworkFeeAlternative) GetGasMode() GasModeEnum`

GetGasMode returns the GasMode field if non-nil, zero value otherwise.

### GetGasModeOk

`func (o *PublicNetworkFeeAlternative) GetGasModeOk() (*GasModeEnum, bool)`

GetGasModeOk returns a tuple with the GasMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasMode

`func (o *PublicNetworkFeeAlternative) SetGasMode(v GasModeEnum)`

SetGasMode sets GasMode field to given value.


### GetBuyerNativeFeeAtomic

`func (o *PublicNetworkFeeAlternative) GetBuyerNativeFeeAtomic() string`

GetBuyerNativeFeeAtomic returns the BuyerNativeFeeAtomic field if non-nil, zero value otherwise.

### GetBuyerNativeFeeAtomicOk

`func (o *PublicNetworkFeeAlternative) GetBuyerNativeFeeAtomicOk() (*string, bool)`

GetBuyerNativeFeeAtomicOk returns a tuple with the BuyerNativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerNativeFeeAtomic

`func (o *PublicNetworkFeeAlternative) SetBuyerNativeFeeAtomic(v string)`

SetBuyerNativeFeeAtomic sets BuyerNativeFeeAtomic field to given value.


### SetBuyerNativeFeeAtomicNil

`func (o *PublicNetworkFeeAlternative) SetBuyerNativeFeeAtomicNil(b bool)`

 SetBuyerNativeFeeAtomicNil sets the value for BuyerNativeFeeAtomic to be an explicit nil

### UnsetBuyerNativeFeeAtomic
`func (o *PublicNetworkFeeAlternative) UnsetBuyerNativeFeeAtomic()`

UnsetBuyerNativeFeeAtomic ensures that no value is present for BuyerNativeFeeAtomic, not even an explicit nil
### GetBuyerPaymentAtomic

`func (o *PublicNetworkFeeAlternative) GetBuyerPaymentAtomic() string`

GetBuyerPaymentAtomic returns the BuyerPaymentAtomic field if non-nil, zero value otherwise.

### GetBuyerPaymentAtomicOk

`func (o *PublicNetworkFeeAlternative) GetBuyerPaymentAtomicOk() (*string, bool)`

GetBuyerPaymentAtomicOk returns a tuple with the BuyerPaymentAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPaymentAtomic

`func (o *PublicNetworkFeeAlternative) SetBuyerPaymentAtomic(v string)`

SetBuyerPaymentAtomic sets BuyerPaymentAtomic field to given value.


### GetTenantProceedsAtomic

`func (o *PublicNetworkFeeAlternative) GetTenantProceedsAtomic() string`

GetTenantProceedsAtomic returns the TenantProceedsAtomic field if non-nil, zero value otherwise.

### GetTenantProceedsAtomicOk

`func (o *PublicNetworkFeeAlternative) GetTenantProceedsAtomicOk() (*string, bool)`

GetTenantProceedsAtomicOk returns a tuple with the TenantProceedsAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantProceedsAtomic

`func (o *PublicNetworkFeeAlternative) SetTenantProceedsAtomic(v string)`

SetTenantProceedsAtomic sets TenantProceedsAtomic field to given value.


### GetQuoteExpiresAt

`func (o *PublicNetworkFeeAlternative) GetQuoteExpiresAt() time.Time`

GetQuoteExpiresAt returns the QuoteExpiresAt field if non-nil, zero value otherwise.

### GetQuoteExpiresAtOk

`func (o *PublicNetworkFeeAlternative) GetQuoteExpiresAtOk() (*time.Time, bool)`

GetQuoteExpiresAtOk returns a tuple with the QuoteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiresAt

`func (o *PublicNetworkFeeAlternative) SetQuoteExpiresAt(v time.Time)`

SetQuoteExpiresAt sets QuoteExpiresAt field to given value.


### SetQuoteExpiresAtNil

`func (o *PublicNetworkFeeAlternative) SetQuoteExpiresAtNil(b bool)`

 SetQuoteExpiresAtNil sets the value for QuoteExpiresAt to be an explicit nil

### UnsetQuoteExpiresAt
`func (o *PublicNetworkFeeAlternative) UnsetQuoteExpiresAt()`

UnsetQuoteExpiresAt ensures that no value is present for QuoteExpiresAt, not even an explicit nil
### GetEligible

`func (o *PublicNetworkFeeAlternative) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *PublicNetworkFeeAlternative) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *PublicNetworkFeeAlternative) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetExclusionReason

`func (o *PublicNetworkFeeAlternative) GetExclusionReason() string`

GetExclusionReason returns the ExclusionReason field if non-nil, zero value otherwise.

### GetExclusionReasonOk

`func (o *PublicNetworkFeeAlternative) GetExclusionReasonOk() (*string, bool)`

GetExclusionReasonOk returns a tuple with the ExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionReason

`func (o *PublicNetworkFeeAlternative) SetExclusionReason(v string)`

SetExclusionReason sets ExclusionReason field to given value.


### SetExclusionReasonNil

`func (o *PublicNetworkFeeAlternative) SetExclusionReasonNil(b bool)`

 SetExclusionReasonNil sets the value for ExclusionReason to be an explicit nil

### UnsetExclusionReason
`func (o *PublicNetworkFeeAlternative) UnsetExclusionReason()`

UnsetExclusionReason ensures that no value is present for ExclusionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
