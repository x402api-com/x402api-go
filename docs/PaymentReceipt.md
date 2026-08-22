# PaymentReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**OrderId** | **string** |  | [readonly]
**SettlementJobId** | **string** |  | [readonly]
**Receipt** | **interface{}** |  | [readonly]
**ReceiptDigest** | **string** |  | [readonly]
**Signature** | **string** |  | [readonly]
**SigningKeyVersion** | **string** |  | [readonly]
**EligibleAlternatives** | [**[]NetworkFeeAlternative**](NetworkFeeAlternative.md) |  | [readonly]
**FeePolicy** | [**NullableFeePolicyDocument**](FeePolicyDocument.md) |  | [readonly]
**FeeEvidence** | [**NullableNetworkFeeEvidence**](NetworkFeeEvidence.md) |  | [readonly]
**FeeQuoteDigest** | **NullableString** |  | [readonly]
**FeeQuoteExpiresAt** | **NullableTime** |  | [readonly]
**SettlementAmountAtomic** | **string** |  | [readonly]
**GasMode** | **string** |  | [readonly]
**BuyerNativeFeeAtomic** | **NullableString** |  | [readonly]
**SponsoredNativeFeeAtomic** | **NullableString** |  | [readonly]
**SponsoredNativeSymbol** | **NullableString** |  | [readonly]
**TenantGasChargeMicros** | **NullableString** |  | [readonly]
**GasSponsorshipEvidenceDigest** | **NullableString** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]

## Methods

### NewPaymentReceipt

`func NewPaymentReceipt(id string, orderId string, settlementJobId string, receipt interface{}, receiptDigest string, signature string, signingKeyVersion string, eligibleAlternatives []NetworkFeeAlternative, feePolicy NullableFeePolicyDocument, feeEvidence NullableNetworkFeeEvidence, feeQuoteDigest NullableString, feeQuoteExpiresAt NullableTime, settlementAmountAtomic string, gasMode string, buyerNativeFeeAtomic NullableString, sponsoredNativeFeeAtomic NullableString, sponsoredNativeSymbol NullableString, tenantGasChargeMicros NullableString, gasSponsorshipEvidenceDigest NullableString, createdAt time.Time, ) *PaymentReceipt`

NewPaymentReceipt instantiates a new PaymentReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReceiptWithDefaults

`func NewPaymentReceiptWithDefaults() *PaymentReceipt`

NewPaymentReceiptWithDefaults instantiates a new PaymentReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentReceipt) SetId(v string)`

SetId sets Id field to given value.


### GetOrderId

`func (o *PaymentReceipt) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentReceipt) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentReceipt) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetSettlementJobId

`func (o *PaymentReceipt) GetSettlementJobId() string`

GetSettlementJobId returns the SettlementJobId field if non-nil, zero value otherwise.

### GetSettlementJobIdOk

`func (o *PaymentReceipt) GetSettlementJobIdOk() (*string, bool)`

GetSettlementJobIdOk returns a tuple with the SettlementJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementJobId

`func (o *PaymentReceipt) SetSettlementJobId(v string)`

SetSettlementJobId sets SettlementJobId field to given value.


### GetReceipt

`func (o *PaymentReceipt) GetReceipt() interface{}`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *PaymentReceipt) GetReceiptOk() (*interface{}, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *PaymentReceipt) SetReceipt(v interface{})`

SetReceipt sets Receipt field to given value.


### SetReceiptNil

`func (o *PaymentReceipt) SetReceiptNil(b bool)`

 SetReceiptNil sets the value for Receipt to be an explicit nil

### UnsetReceipt
`func (o *PaymentReceipt) UnsetReceipt()`

UnsetReceipt ensures that no value is present for Receipt, not even an explicit nil
### GetReceiptDigest

`func (o *PaymentReceipt) GetReceiptDigest() string`

GetReceiptDigest returns the ReceiptDigest field if non-nil, zero value otherwise.

### GetReceiptDigestOk

`func (o *PaymentReceipt) GetReceiptDigestOk() (*string, bool)`

GetReceiptDigestOk returns a tuple with the ReceiptDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDigest

`func (o *PaymentReceipt) SetReceiptDigest(v string)`

SetReceiptDigest sets ReceiptDigest field to given value.


### GetSignature

`func (o *PaymentReceipt) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentReceipt) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentReceipt) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSigningKeyVersion

`func (o *PaymentReceipt) GetSigningKeyVersion() string`

GetSigningKeyVersion returns the SigningKeyVersion field if non-nil, zero value otherwise.

### GetSigningKeyVersionOk

`func (o *PaymentReceipt) GetSigningKeyVersionOk() (*string, bool)`

GetSigningKeyVersionOk returns a tuple with the SigningKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyVersion

`func (o *PaymentReceipt) SetSigningKeyVersion(v string)`

SetSigningKeyVersion sets SigningKeyVersion field to given value.


### GetEligibleAlternatives

`func (o *PaymentReceipt) GetEligibleAlternatives() []NetworkFeeAlternative`

GetEligibleAlternatives returns the EligibleAlternatives field if non-nil, zero value otherwise.

### GetEligibleAlternativesOk

`func (o *PaymentReceipt) GetEligibleAlternativesOk() (*[]NetworkFeeAlternative, bool)`

GetEligibleAlternativesOk returns a tuple with the EligibleAlternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleAlternatives

`func (o *PaymentReceipt) SetEligibleAlternatives(v []NetworkFeeAlternative)`

SetEligibleAlternatives sets EligibleAlternatives field to given value.


### GetFeePolicy

`func (o *PaymentReceipt) GetFeePolicy() FeePolicyDocument`

GetFeePolicy returns the FeePolicy field if non-nil, zero value otherwise.

### GetFeePolicyOk

`func (o *PaymentReceipt) GetFeePolicyOk() (*FeePolicyDocument, bool)`

GetFeePolicyOk returns a tuple with the FeePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicy

`func (o *PaymentReceipt) SetFeePolicy(v FeePolicyDocument)`

SetFeePolicy sets FeePolicy field to given value.


### SetFeePolicyNil

`func (o *PaymentReceipt) SetFeePolicyNil(b bool)`

 SetFeePolicyNil sets the value for FeePolicy to be an explicit nil

### UnsetFeePolicy
`func (o *PaymentReceipt) UnsetFeePolicy()`

UnsetFeePolicy ensures that no value is present for FeePolicy, not even an explicit nil
### GetFeeEvidence

`func (o *PaymentReceipt) GetFeeEvidence() NetworkFeeEvidence`

GetFeeEvidence returns the FeeEvidence field if non-nil, zero value otherwise.

### GetFeeEvidenceOk

`func (o *PaymentReceipt) GetFeeEvidenceOk() (*NetworkFeeEvidence, bool)`

GetFeeEvidenceOk returns a tuple with the FeeEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeEvidence

`func (o *PaymentReceipt) SetFeeEvidence(v NetworkFeeEvidence)`

SetFeeEvidence sets FeeEvidence field to given value.


### SetFeeEvidenceNil

`func (o *PaymentReceipt) SetFeeEvidenceNil(b bool)`

 SetFeeEvidenceNil sets the value for FeeEvidence to be an explicit nil

### UnsetFeeEvidence
`func (o *PaymentReceipt) UnsetFeeEvidence()`

UnsetFeeEvidence ensures that no value is present for FeeEvidence, not even an explicit nil
### GetFeeQuoteDigest

`func (o *PaymentReceipt) GetFeeQuoteDigest() string`

GetFeeQuoteDigest returns the FeeQuoteDigest field if non-nil, zero value otherwise.

### GetFeeQuoteDigestOk

`func (o *PaymentReceipt) GetFeeQuoteDigestOk() (*string, bool)`

GetFeeQuoteDigestOk returns a tuple with the FeeQuoteDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeQuoteDigest

`func (o *PaymentReceipt) SetFeeQuoteDigest(v string)`

SetFeeQuoteDigest sets FeeQuoteDigest field to given value.


### SetFeeQuoteDigestNil

`func (o *PaymentReceipt) SetFeeQuoteDigestNil(b bool)`

 SetFeeQuoteDigestNil sets the value for FeeQuoteDigest to be an explicit nil

### UnsetFeeQuoteDigest
`func (o *PaymentReceipt) UnsetFeeQuoteDigest()`

UnsetFeeQuoteDigest ensures that no value is present for FeeQuoteDigest, not even an explicit nil
### GetFeeQuoteExpiresAt

`func (o *PaymentReceipt) GetFeeQuoteExpiresAt() time.Time`

GetFeeQuoteExpiresAt returns the FeeQuoteExpiresAt field if non-nil, zero value otherwise.

### GetFeeQuoteExpiresAtOk

`func (o *PaymentReceipt) GetFeeQuoteExpiresAtOk() (*time.Time, bool)`

GetFeeQuoteExpiresAtOk returns a tuple with the FeeQuoteExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeQuoteExpiresAt

`func (o *PaymentReceipt) SetFeeQuoteExpiresAt(v time.Time)`

SetFeeQuoteExpiresAt sets FeeQuoteExpiresAt field to given value.


### SetFeeQuoteExpiresAtNil

`func (o *PaymentReceipt) SetFeeQuoteExpiresAtNil(b bool)`

 SetFeeQuoteExpiresAtNil sets the value for FeeQuoteExpiresAt to be an explicit nil

### UnsetFeeQuoteExpiresAt
`func (o *PaymentReceipt) UnsetFeeQuoteExpiresAt()`

UnsetFeeQuoteExpiresAt ensures that no value is present for FeeQuoteExpiresAt, not even an explicit nil
### GetSettlementAmountAtomic

`func (o *PaymentReceipt) GetSettlementAmountAtomic() string`

GetSettlementAmountAtomic returns the SettlementAmountAtomic field if non-nil, zero value otherwise.

### GetSettlementAmountAtomicOk

`func (o *PaymentReceipt) GetSettlementAmountAtomicOk() (*string, bool)`

GetSettlementAmountAtomicOk returns a tuple with the SettlementAmountAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmountAtomic

`func (o *PaymentReceipt) SetSettlementAmountAtomic(v string)`

SetSettlementAmountAtomic sets SettlementAmountAtomic field to given value.


### GetGasMode

`func (o *PaymentReceipt) GetGasMode() string`

GetGasMode returns the GasMode field if non-nil, zero value otherwise.

### GetGasModeOk

`func (o *PaymentReceipt) GetGasModeOk() (*string, bool)`

GetGasModeOk returns a tuple with the GasMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasMode

`func (o *PaymentReceipt) SetGasMode(v string)`

SetGasMode sets GasMode field to given value.


### GetBuyerNativeFeeAtomic

`func (o *PaymentReceipt) GetBuyerNativeFeeAtomic() string`

GetBuyerNativeFeeAtomic returns the BuyerNativeFeeAtomic field if non-nil, zero value otherwise.

### GetBuyerNativeFeeAtomicOk

`func (o *PaymentReceipt) GetBuyerNativeFeeAtomicOk() (*string, bool)`

GetBuyerNativeFeeAtomicOk returns a tuple with the BuyerNativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerNativeFeeAtomic

`func (o *PaymentReceipt) SetBuyerNativeFeeAtomic(v string)`

SetBuyerNativeFeeAtomic sets BuyerNativeFeeAtomic field to given value.


### SetBuyerNativeFeeAtomicNil

`func (o *PaymentReceipt) SetBuyerNativeFeeAtomicNil(b bool)`

 SetBuyerNativeFeeAtomicNil sets the value for BuyerNativeFeeAtomic to be an explicit nil

### UnsetBuyerNativeFeeAtomic
`func (o *PaymentReceipt) UnsetBuyerNativeFeeAtomic()`

UnsetBuyerNativeFeeAtomic ensures that no value is present for BuyerNativeFeeAtomic, not even an explicit nil
### GetSponsoredNativeFeeAtomic

`func (o *PaymentReceipt) GetSponsoredNativeFeeAtomic() string`

GetSponsoredNativeFeeAtomic returns the SponsoredNativeFeeAtomic field if non-nil, zero value otherwise.

### GetSponsoredNativeFeeAtomicOk

`func (o *PaymentReceipt) GetSponsoredNativeFeeAtomicOk() (*string, bool)`

GetSponsoredNativeFeeAtomicOk returns a tuple with the SponsoredNativeFeeAtomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoredNativeFeeAtomic

`func (o *PaymentReceipt) SetSponsoredNativeFeeAtomic(v string)`

SetSponsoredNativeFeeAtomic sets SponsoredNativeFeeAtomic field to given value.


### SetSponsoredNativeFeeAtomicNil

`func (o *PaymentReceipt) SetSponsoredNativeFeeAtomicNil(b bool)`

 SetSponsoredNativeFeeAtomicNil sets the value for SponsoredNativeFeeAtomic to be an explicit nil

### UnsetSponsoredNativeFeeAtomic
`func (o *PaymentReceipt) UnsetSponsoredNativeFeeAtomic()`

UnsetSponsoredNativeFeeAtomic ensures that no value is present for SponsoredNativeFeeAtomic, not even an explicit nil
### GetSponsoredNativeSymbol

`func (o *PaymentReceipt) GetSponsoredNativeSymbol() string`

GetSponsoredNativeSymbol returns the SponsoredNativeSymbol field if non-nil, zero value otherwise.

### GetSponsoredNativeSymbolOk

`func (o *PaymentReceipt) GetSponsoredNativeSymbolOk() (*string, bool)`

GetSponsoredNativeSymbolOk returns a tuple with the SponsoredNativeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoredNativeSymbol

`func (o *PaymentReceipt) SetSponsoredNativeSymbol(v string)`

SetSponsoredNativeSymbol sets SponsoredNativeSymbol field to given value.


### SetSponsoredNativeSymbolNil

`func (o *PaymentReceipt) SetSponsoredNativeSymbolNil(b bool)`

 SetSponsoredNativeSymbolNil sets the value for SponsoredNativeSymbol to be an explicit nil

### UnsetSponsoredNativeSymbol
`func (o *PaymentReceipt) UnsetSponsoredNativeSymbol()`

UnsetSponsoredNativeSymbol ensures that no value is present for SponsoredNativeSymbol, not even an explicit nil
### GetTenantGasChargeMicros

`func (o *PaymentReceipt) GetTenantGasChargeMicros() string`

GetTenantGasChargeMicros returns the TenantGasChargeMicros field if non-nil, zero value otherwise.

### GetTenantGasChargeMicrosOk

`func (o *PaymentReceipt) GetTenantGasChargeMicrosOk() (*string, bool)`

GetTenantGasChargeMicrosOk returns a tuple with the TenantGasChargeMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGasChargeMicros

`func (o *PaymentReceipt) SetTenantGasChargeMicros(v string)`

SetTenantGasChargeMicros sets TenantGasChargeMicros field to given value.


### SetTenantGasChargeMicrosNil

`func (o *PaymentReceipt) SetTenantGasChargeMicrosNil(b bool)`

 SetTenantGasChargeMicrosNil sets the value for TenantGasChargeMicros to be an explicit nil

### UnsetTenantGasChargeMicros
`func (o *PaymentReceipt) UnsetTenantGasChargeMicros()`

UnsetTenantGasChargeMicros ensures that no value is present for TenantGasChargeMicros, not even an explicit nil
### GetGasSponsorshipEvidenceDigest

`func (o *PaymentReceipt) GetGasSponsorshipEvidenceDigest() string`

GetGasSponsorshipEvidenceDigest returns the GasSponsorshipEvidenceDigest field if non-nil, zero value otherwise.

### GetGasSponsorshipEvidenceDigestOk

`func (o *PaymentReceipt) GetGasSponsorshipEvidenceDigestOk() (*string, bool)`

GetGasSponsorshipEvidenceDigestOk returns a tuple with the GasSponsorshipEvidenceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasSponsorshipEvidenceDigest

`func (o *PaymentReceipt) SetGasSponsorshipEvidenceDigest(v string)`

SetGasSponsorshipEvidenceDigest sets GasSponsorshipEvidenceDigest field to given value.


### SetGasSponsorshipEvidenceDigestNil

`func (o *PaymentReceipt) SetGasSponsorshipEvidenceDigestNil(b bool)`

 SetGasSponsorshipEvidenceDigestNil sets the value for GasSponsorshipEvidenceDigest to be an explicit nil

### UnsetGasSponsorshipEvidenceDigest
`func (o *PaymentReceipt) UnsetGasSponsorshipEvidenceDigest()`

UnsetGasSponsorshipEvidenceDigest ensures that no value is present for GasSponsorshipEvidenceDigest, not even an explicit nil
### GetCreatedAt

`func (o *PaymentReceipt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentReceipt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentReceipt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
