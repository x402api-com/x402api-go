# NetworkFeeAlternative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Version** | **int32** |  |
**Network** | **string** |  |
**AssetId** | **string** |  |
**ContractAddress** | **string** |  |
**FeeMode** | [**FeePolicyModeInputEnum**](FeePolicyModeInputEnum.md) |  |
**QuoteCurrency** | [**FeePolicyQuoteCurrencyInputEnum**](FeePolicyQuoteCurrencyInputEnum.md) |  |
**ListedAmountAtomic** | **string** |  |
**FeeAllowanceCapQuoteMicros** | **string** |  |
**EstimatedNativeFeeAtomic** | **NullableString** |  |
**NativeSymbol** | **NullableString** |  |
**NativeDecimals** | **NullableInt32** |  |
**NativeUsdQuoteMicros** | **NullableString** |  |
**EstimatedFeeQuoteMicros** | **NullableString** |  |
**GasMode** | [**GasModeEnum**](GasModeEnum.md) |  |
**BuyerNativeFeeAtomic** | **NullableString** |  |
**MaximumTenantGasReservationMicros** | **string** |  |
**ProviderDisagreementBps** | **NullableInt32** |  |
**FeeAllowanceQuoteMicros** | **string** |  |
**FeeAllowanceAtomic** | **string** |  |
**BuyerPaymentAtomic** | **string** |  |
**TenantProceedsAtomic** | **string** |  |
**QuoteExpiresAt** | **NullableTime** |  |
**FeeEvidence** | [**NetworkFeeEvidence**](NetworkFeeEvidence.md) |  |
**FeeEvidenceDigest** | **string** |  |
**Eligible** | **bool** |  |
**ExclusionReason** | **NullableString** |  |

## Methods

### NewNetworkFeeAlternative

`func NewNetworkFeeAlternative(type_ string, version int32, network string, assetId string, contractAddress string, feeMode FeePolicyModeInputEnum, quoteCurrency FeePolicyQuoteCurrencyInputEnum, listedAmountAtomic string, feeAllowanceCapQuoteMicros string, estimatedNativeFeeAtomic NullableString, nativeSymbol NullableString, nativeDecimals NullableInt32, nativeUsdQuoteMicros NullableString, estimatedFeeQuoteMicros NullableString, gasMode GasModeEnum, buyerNativeFeeAtomic NullableString, maximumTenantGasReservationMicros string, providerDisagreementBps NullableInt32, feeAllowanceQuoteMicros string, feeAllowanceAtomic string, buyerPaymentAtomic string, tenantProceedsAtomic string, quoteExpiresAt NullableTime, feeEvidence NetworkFeeEvidence, feeEvidenceDigest string, eligible bool, exclusionReason NullableString, ) *NetworkFeeAlternative`

NewNetworkFeeAlternative instantiates a new NetworkFeeAlternative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeeAlternativeWithDefaults

`func NewNetworkFeeAlternativeWithDefaults() *NetworkFeeAlternative`

NewNetworkFeeAlternativeWithDefaults instantiates a new NetworkFeeAlternative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkFeeAlternative) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkFeeAlternative) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkFeeAlternative) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *NetworkFeeAlternative) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkFeeAlternative) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkFeeAlternative) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetNetwork

`func (o *NetworkFeeAlternative) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkFeeAlternative) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkFeeAlternative) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *NetworkFeeAlternative) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *NetworkFeeAlternative) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *NetworkFeeAlternative) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetContractAddress

`func (o *NetworkFeeAlternative) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *NetworkFeeAlternative) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *NetworkFeeAlternative) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetFeeMode

`func (o *NetworkFeeAlternative) GetFeeMode() FeePolicyModeInputEnum`

GetFeeMode returns the FeeMode field if non-nil, zero value otherwise.

### GetFeeModeOk

`func (o *NetworkFeeAlternative) GetFeeModeOk() (*FeePolicyModeInputEnum, bool)`

GetFeeModeOk returns a tuple with the FeeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeMode

`func (o *NetworkFeeAlternative) SetFeeMode(v FeePolicyModeInputEnum)`

SetFeeMode sets FeeMode field to given value.


### GetQuoteCurrency

`func (o *NetworkFeeAlternative) GetQuoteCurrency() FeePolicyQuoteCurrencyInputEnum`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *NetworkFeeAlternative) GetQuoteCurrencyOk() (*FeePolicyQuoteCurrencyInputEnum, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *NetworkFeeAlternative) SetQuoteCurrency(v FeePolicyQuoteCurrencyInputEnum)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetListedAmountAtomic

`func (o *NetworkFeeAlternative) GetListedAmountAtomic() string`

GetListedAmountAtomic returns the ListedAmountAtomic field if non-nil, zero value otherwise.

### GetListedAmountAtomicOk

`func (o *NetworkFeeAlternative) GetListedAmountAtomicOk() (*string, bool)`

GetListedAmountAtomicOk returns a tuple with the ListedAmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAmountAtomic

`func (o *NetworkFeeAlternative) SetListedAmountAtomic(v string)`

SetListedAmountAtomic sets ListedAmountAtomic field to given value.


### GetFeeAllowanceCapQuoteMicros

`func (o *NetworkFeeAlternative) GetFeeAllowanceCapQuoteMicros() string`

GetFeeAllowanceCapQuoteMicros returns the FeeAllowanceCapQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceCapQuoteMicrosOk

`func (o *NetworkFeeAlternative) GetFeeAllowanceCapQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceCapQuoteMicrosOk returns a tuple with the FeeAllowanceCapQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceCapQuoteMicros

`func (o *NetworkFeeAlternative) SetFeeAllowanceCapQuoteMicros(v string)`

SetFeeAllowanceCapQuoteMicros sets FeeAllowanceCapQuoteMicros field to given value.


### GetEstimatedNativeFeeAtomic

`func (o *NetworkFeeAlternative) GetEstimatedNativeFeeAtomic() string`

GetEstimatedNativeFeeAtomic returns the EstimatedNativeFeeAtomic field if non-nil, zero value otherwise.

### GetEstimatedNativeFeeAtomicOk

`func (o *NetworkFeeAlternative) GetEstimatedNativeFeeAtomicOk() (*string, bool)`

GetEstimatedNativeFeeAtomicOk returns a tuple with the EstimatedNativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNativeFeeAtomic

`func (o *NetworkFeeAlternative) SetEstimatedNativeFeeAtomic(v string)`

SetEstimatedNativeFeeAtomic sets EstimatedNativeFeeAtomic field to given value.


### SetEstimatedNativeFeeAtomicNil

`func (o *NetworkFeeAlternative) SetEstimatedNativeFeeAtomicNil(b bool)`

 SetEstimatedNativeFeeAtomicNil sets the value for EstimatedNativeFeeAtomic to be an explicit nil

### UnsetEstimatedNativeFeeAtomic
`func (o *NetworkFeeAlternative) UnsetEstimatedNativeFeeAtomic()`

UnsetEstimatedNativeFeeAtomic ensures that no value is present for EstimatedNativeFeeAtomic, not even an explicit nil
### GetNativeSymbol

`func (o *NetworkFeeAlternative) GetNativeSymbol() string`

GetNativeSymbol returns the NativeSymbol field if non-nil, zero value otherwise.

### GetNativeSymbolOk

`func (o *NetworkFeeAlternative) GetNativeSymbolOk() (*string, bool)`

GetNativeSymbolOk returns a tuple with the NativeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeSymbol

`func (o *NetworkFeeAlternative) SetNativeSymbol(v string)`

SetNativeSymbol sets NativeSymbol field to given value.


### SetNativeSymbolNil

`func (o *NetworkFeeAlternative) SetNativeSymbolNil(b bool)`

 SetNativeSymbolNil sets the value for NativeSymbol to be an explicit nil

### UnsetNativeSymbol
`func (o *NetworkFeeAlternative) UnsetNativeSymbol()`

UnsetNativeSymbol ensures that no value is present for NativeSymbol, not even an explicit nil
### GetNativeDecimals

`func (o *NetworkFeeAlternative) GetNativeDecimals() int32`

GetNativeDecimals returns the NativeDecimals field if non-nil, zero value otherwise.

### GetNativeDecimalsOk

`func (o *NetworkFeeAlternative) GetNativeDecimalsOk() (*int32, bool)`

GetNativeDecimalsOk returns a tuple with the NativeDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeDecimals

`func (o *NetworkFeeAlternative) SetNativeDecimals(v int32)`

SetNativeDecimals sets NativeDecimals field to given value.


### SetNativeDecimalsNil

`func (o *NetworkFeeAlternative) SetNativeDecimalsNil(b bool)`

 SetNativeDecimalsNil sets the value for NativeDecimals to be an explicit nil

### UnsetNativeDecimals
`func (o *NetworkFeeAlternative) UnsetNativeDecimals()`

UnsetNativeDecimals ensures that no value is present for NativeDecimals, not even an explicit nil
### GetNativeUsdQuoteMicros

`func (o *NetworkFeeAlternative) GetNativeUsdQuoteMicros() string`

GetNativeUsdQuoteMicros returns the NativeUsdQuoteMicros field if non-nil, zero value otherwise.

### GetNativeUsdQuoteMicrosOk

`func (o *NetworkFeeAlternative) GetNativeUsdQuoteMicrosOk() (*string, bool)`

GetNativeUsdQuoteMicrosOk returns a tuple with the NativeUsdQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeUsdQuoteMicros

`func (o *NetworkFeeAlternative) SetNativeUsdQuoteMicros(v string)`

SetNativeUsdQuoteMicros sets NativeUsdQuoteMicros field to given value.


### SetNativeUsdQuoteMicrosNil

`func (o *NetworkFeeAlternative) SetNativeUsdQuoteMicrosNil(b bool)`

 SetNativeUsdQuoteMicrosNil sets the value for NativeUsdQuoteMicros to be an explicit nil

### UnsetNativeUsdQuoteMicros
`func (o *NetworkFeeAlternative) UnsetNativeUsdQuoteMicros()`

UnsetNativeUsdQuoteMicros ensures that no value is present for NativeUsdQuoteMicros, not even an explicit nil
### GetEstimatedFeeQuoteMicros

`func (o *NetworkFeeAlternative) GetEstimatedFeeQuoteMicros() string`

GetEstimatedFeeQuoteMicros returns the EstimatedFeeQuoteMicros field if non-nil, zero value otherwise.

### GetEstimatedFeeQuoteMicrosOk

`func (o *NetworkFeeAlternative) GetEstimatedFeeQuoteMicrosOk() (*string, bool)`

GetEstimatedFeeQuoteMicrosOk returns a tuple with the EstimatedFeeQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeQuoteMicros

`func (o *NetworkFeeAlternative) SetEstimatedFeeQuoteMicros(v string)`

SetEstimatedFeeQuoteMicros sets EstimatedFeeQuoteMicros field to given value.


### SetEstimatedFeeQuoteMicrosNil

`func (o *NetworkFeeAlternative) SetEstimatedFeeQuoteMicrosNil(b bool)`

 SetEstimatedFeeQuoteMicrosNil sets the value for EstimatedFeeQuoteMicros to be an explicit nil

### UnsetEstimatedFeeQuoteMicros
`func (o *NetworkFeeAlternative) UnsetEstimatedFeeQuoteMicros()`

UnsetEstimatedFeeQuoteMicros ensures that no value is present for EstimatedFeeQuoteMicros, not even an explicit nil
### GetGasMode

`func (o *NetworkFeeAlternative) GetGasMode() GasModeEnum`

GetGasMode returns the GasMode field if non-nil, zero value otherwise.

### GetGasModeOk

`func (o *NetworkFeeAlternative) GetGasModeOk() (*GasModeEnum, bool)`

GetGasModeOk returns a tuple with the GasMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasMode

`func (o *NetworkFeeAlternative) SetGasMode(v GasModeEnum)`

SetGasMode sets GasMode field to given value.


### GetBuyerNativeFeeAtomic

`func (o *NetworkFeeAlternative) GetBuyerNativeFeeAtomic() string`

GetBuyerNativeFeeAtomic returns the BuyerNativeFeeAtomic field if non-nil, zero value otherwise.

### GetBuyerNativeFeeAtomicOk

`func (o *NetworkFeeAlternative) GetBuyerNativeFeeAtomicOk() (*string, bool)`

GetBuyerNativeFeeAtomicOk returns a tuple with the BuyerNativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerNativeFeeAtomic

`func (o *NetworkFeeAlternative) SetBuyerNativeFeeAtomic(v string)`

SetBuyerNativeFeeAtomic sets BuyerNativeFeeAtomic field to given value.


### SetBuyerNativeFeeAtomicNil

`func (o *NetworkFeeAlternative) SetBuyerNativeFeeAtomicNil(b bool)`

 SetBuyerNativeFeeAtomicNil sets the value for BuyerNativeFeeAtomic to be an explicit nil

### UnsetBuyerNativeFeeAtomic
`func (o *NetworkFeeAlternative) UnsetBuyerNativeFeeAtomic()`

UnsetBuyerNativeFeeAtomic ensures that no value is present for BuyerNativeFeeAtomic, not even an explicit nil
### GetMaximumTenantGasReservationMicros

`func (o *NetworkFeeAlternative) GetMaximumTenantGasReservationMicros() string`

GetMaximumTenantGasReservationMicros returns the MaximumTenantGasReservationMicros field if non-nil, zero value otherwise.

### GetMaximumTenantGasReservationMicrosOk

`func (o *NetworkFeeAlternative) GetMaximumTenantGasReservationMicrosOk() (*string, bool)`

GetMaximumTenantGasReservationMicrosOk returns a tuple with the MaximumTenantGasReservationMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTenantGasReservationMicros

`func (o *NetworkFeeAlternative) SetMaximumTenantGasReservationMicros(v string)`

SetMaximumTenantGasReservationMicros sets MaximumTenantGasReservationMicros field to given value.


### GetProviderDisagreementBps

`func (o *NetworkFeeAlternative) GetProviderDisagreementBps() int32`

GetProviderDisagreementBps returns the ProviderDisagreementBps field if non-nil, zero value otherwise.

### GetProviderDisagreementBpsOk

`func (o *NetworkFeeAlternative) GetProviderDisagreementBpsOk() (*int32, bool)`

GetProviderDisagreementBpsOk returns a tuple with the ProviderDisagreementBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDisagreementBps

`func (o *NetworkFeeAlternative) SetProviderDisagreementBps(v int32)`

SetProviderDisagreementBps sets ProviderDisagreementBps field to given value.


### SetProviderDisagreementBpsNil

`func (o *NetworkFeeAlternative) SetProviderDisagreementBpsNil(b bool)`

 SetProviderDisagreementBpsNil sets the value for ProviderDisagreementBps to be an explicit nil

### UnsetProviderDisagreementBps
`func (o *NetworkFeeAlternative) UnsetProviderDisagreementBps()`

UnsetProviderDisagreementBps ensures that no value is present for ProviderDisagreementBps, not even an explicit nil
### GetFeeAllowanceQuoteMicros

`func (o *NetworkFeeAlternative) GetFeeAllowanceQuoteMicros() string`

GetFeeAllowanceQuoteMicros returns the FeeAllowanceQuoteMicros field if non-nil, zero value otherwise.

### GetFeeAllowanceQuoteMicrosOk

`func (o *NetworkFeeAlternative) GetFeeAllowanceQuoteMicrosOk() (*string, bool)`

GetFeeAllowanceQuoteMicrosOk returns a tuple with the FeeAllowanceQuoteMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceQuoteMicros

`func (o *NetworkFeeAlternative) SetFeeAllowanceQuoteMicros(v string)`

SetFeeAllowanceQuoteMicros sets FeeAllowanceQuoteMicros field to given value.


### GetFeeAllowanceAtomic

`func (o *NetworkFeeAlternative) GetFeeAllowanceAtomic() string`

GetFeeAllowanceAtomic returns the FeeAllowanceAtomic field if non-nil, zero value otherwise.

### GetFeeAllowanceAtomicOk

`func (o *NetworkFeeAlternative) GetFeeAllowanceAtomicOk() (*string, bool)`

GetFeeAllowanceAtomicOk returns a tuple with the FeeAllowanceAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAllowanceAtomic

`func (o *NetworkFeeAlternative) SetFeeAllowanceAtomic(v string)`

SetFeeAllowanceAtomic sets FeeAllowanceAtomic field to given value.


### GetBuyerPaymentAtomic

`func (o *NetworkFeeAlternative) GetBuyerPaymentAtomic() string`

GetBuyerPaymentAtomic returns the BuyerPaymentAtomic field if non-nil, zero value otherwise.

### GetBuyerPaymentAtomicOk

`func (o *NetworkFeeAlternative) GetBuyerPaymentAtomicOk() (*string, bool)`

GetBuyerPaymentAtomicOk returns a tuple with the BuyerPaymentAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPaymentAtomic

`func (o *NetworkFeeAlternative) SetBuyerPaymentAtomic(v string)`

SetBuyerPaymentAtomic sets BuyerPaymentAtomic field to given value.


### GetTenantProceedsAtomic

`func (o *NetworkFeeAlternative) GetTenantProceedsAtomic() string`

GetTenantProceedsAtomic returns the TenantProceedsAtomic field if non-nil, zero value otherwise.

### GetTenantProceedsAtomicOk

`func (o *NetworkFeeAlternative) GetTenantProceedsAtomicOk() (*string, bool)`

GetTenantProceedsAtomicOk returns a tuple with the TenantProceedsAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantProceedsAtomic

`func (o *NetworkFeeAlternative) SetTenantProceedsAtomic(v string)`

SetTenantProceedsAtomic sets TenantProceedsAtomic field to given value.


### GetQuoteExpiresAt

`func (o *NetworkFeeAlternative) GetQuoteExpiresAt() time.Time`

GetQuoteExpiresAt returns the QuoteExpiresAt field if non-nil, zero value otherwise.

### GetQuoteExpiresAtOk

`func (o *NetworkFeeAlternative) GetQuoteExpiresAtOk() (*time.Time, bool)`

GetQuoteExpiresAtOk returns a tuple with the QuoteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExpiresAt

`func (o *NetworkFeeAlternative) SetQuoteExpiresAt(v time.Time)`

SetQuoteExpiresAt sets QuoteExpiresAt field to given value.


### SetQuoteExpiresAtNil

`func (o *NetworkFeeAlternative) SetQuoteExpiresAtNil(b bool)`

 SetQuoteExpiresAtNil sets the value for QuoteExpiresAt to be an explicit nil

### UnsetQuoteExpiresAt
`func (o *NetworkFeeAlternative) UnsetQuoteExpiresAt()`

UnsetQuoteExpiresAt ensures that no value is present for QuoteExpiresAt, not even an explicit nil
### GetFeeEvidence

`func (o *NetworkFeeAlternative) GetFeeEvidence() NetworkFeeEvidence`

GetFeeEvidence returns the FeeEvidence field if non-nil, zero value otherwise.

### GetFeeEvidenceOk

`func (o *NetworkFeeAlternative) GetFeeEvidenceOk() (*NetworkFeeEvidence, bool)`

GetFeeEvidenceOk returns a tuple with the FeeEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeEvidence

`func (o *NetworkFeeAlternative) SetFeeEvidence(v NetworkFeeEvidence)`

SetFeeEvidence sets FeeEvidence field to given value.


### GetFeeEvidenceDigest

`func (o *NetworkFeeAlternative) GetFeeEvidenceDigest() string`

GetFeeEvidenceDigest returns the FeeEvidenceDigest field if non-nil, zero value otherwise.

### GetFeeEvidenceDigestOk

`func (o *NetworkFeeAlternative) GetFeeEvidenceDigestOk() (*string, bool)`

GetFeeEvidenceDigestOk returns a tuple with the FeeEvidenceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeEvidenceDigest

`func (o *NetworkFeeAlternative) SetFeeEvidenceDigest(v string)`

SetFeeEvidenceDigest sets FeeEvidenceDigest field to given value.


### GetEligible

`func (o *NetworkFeeAlternative) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *NetworkFeeAlternative) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *NetworkFeeAlternative) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetExclusionReason

`func (o *NetworkFeeAlternative) GetExclusionReason() string`

GetExclusionReason returns the ExclusionReason field if non-nil, zero value otherwise.

### GetExclusionReasonOk

`func (o *NetworkFeeAlternative) GetExclusionReasonOk() (*string, bool)`

GetExclusionReasonOk returns a tuple with the ExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionReason

`func (o *NetworkFeeAlternative) SetExclusionReason(v string)`

SetExclusionReason sets ExclusionReason field to given value.


### SetExclusionReasonNil

`func (o *NetworkFeeAlternative) SetExclusionReasonNil(b bool)`

 SetExclusionReasonNil sets the value for ExclusionReason to be an explicit nil

### UnsetExclusionReason
`func (o *NetworkFeeAlternative) UnsetExclusionReason()`

UnsetExclusionReason ensures that no value is present for ExclusionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
