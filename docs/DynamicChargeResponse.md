# DynamicChargeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeId** | **string** | Immutable challenge UUID created for this charge. |
**ChargeDigest** | **string** |  |
**OrderId** | **string** |  |
**Status** | **string** | Current projected order status; payment terms remain immutable. |
**ResourceVersionId** | **string** |  |
**PaymentIdentifier** | **string** | Opaque server challenge handle. Return it to the buyer as X-X402API-Challenge-Handle; it is not the buyer payment identifier. |
**ExpiresAt** | **time.Time** |  |
**CreatedAt** | **time.Time** |  |
**Prices** | [**[]DynamicChargePrice**](DynamicChargePrice.md) |  |
**RequestedExpiresInSeconds** | **int32** |  |
**Metadata** | **map[string]interface{}** | Tenant application metadata frozen into the charge digest. Maximum canonical size is 16 KiB; floating-point numbers are not accepted. |
**MetadataDigest** | **string** |  |
**PaymentRequired** | **interface{}** | Complete immutable x402 v2 PAYMENT-REQUIRED document. |
**PaymentRequiredHeader** | **string** | Canonical base64-encoded value to return in the buyer-facing PAYMENT-REQUIRED header. |
**EligibleAlternatives** | [**[]PublicNetworkFeeAlternative**](PublicNetworkFeeAlternative.md) |  |
**FeePolicy** | [**PublicFeePolicyDocument**](PublicFeePolicyDocument.md) |  |
**FeeQuoteDigest** | **string** |  |

## Methods

### NewDynamicChargeResponse

`func NewDynamicChargeResponse(chargeId string, chargeDigest string, orderId string, status string, resourceVersionId string, paymentIdentifier string, expiresAt time.Time, createdAt time.Time, prices []DynamicChargePrice, requestedExpiresInSeconds int32, metadata map[string]interface{}, metadataDigest string, paymentRequired interface{}, paymentRequiredHeader string, eligibleAlternatives []PublicNetworkFeeAlternative, feePolicy PublicFeePolicyDocument, feeQuoteDigest string, ) *DynamicChargeResponse`

NewDynamicChargeResponse instantiates a new DynamicChargeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicChargeResponseWithDefaults

`func NewDynamicChargeResponseWithDefaults() *DynamicChargeResponse`

NewDynamicChargeResponseWithDefaults instantiates a new DynamicChargeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeId

`func (o *DynamicChargeResponse) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *DynamicChargeResponse) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *DynamicChargeResponse) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.


### GetChargeDigest

`func (o *DynamicChargeResponse) GetChargeDigest() string`

GetChargeDigest returns the ChargeDigest field if non-nil, zero value otherwise.

### GetChargeDigestOk

`func (o *DynamicChargeResponse) GetChargeDigestOk() (*string, bool)`

GetChargeDigestOk returns a tuple with the ChargeDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDigest

`func (o *DynamicChargeResponse) SetChargeDigest(v string)`

SetChargeDigest sets ChargeDigest field to given value.


### GetOrderId

`func (o *DynamicChargeResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *DynamicChargeResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *DynamicChargeResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetStatus

`func (o *DynamicChargeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicChargeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicChargeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResourceVersionId

`func (o *DynamicChargeResponse) GetResourceVersionId() string`

GetResourceVersionId returns the ResourceVersionId field if non-nil, zero value otherwise.

### GetResourceVersionIdOk

`func (o *DynamicChargeResponse) GetResourceVersionIdOk() (*string, bool)`

GetResourceVersionIdOk returns a tuple with the ResourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersionId

`func (o *DynamicChargeResponse) SetResourceVersionId(v string)`

SetResourceVersionId sets ResourceVersionId field to given value.


### GetPaymentIdentifier

`func (o *DynamicChargeResponse) GetPaymentIdentifier() string`

GetPaymentIdentifier returns the PaymentIdentifier field if non-nil, zero value otherwise.

### GetPaymentIdentifierOk

`func (o *DynamicChargeResponse) GetPaymentIdentifierOk() (*string, bool)`

GetPaymentIdentifierOk returns a tuple with the PaymentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIdentifier

`func (o *DynamicChargeResponse) SetPaymentIdentifier(v string)`

SetPaymentIdentifier sets PaymentIdentifier field to given value.


### GetExpiresAt

`func (o *DynamicChargeResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DynamicChargeResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DynamicChargeResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *DynamicChargeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DynamicChargeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DynamicChargeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPrices

`func (o *DynamicChargeResponse) GetPrices() []DynamicChargePrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DynamicChargeResponse) GetPricesOk() (*[]DynamicChargePrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DynamicChargeResponse) SetPrices(v []DynamicChargePrice)`

SetPrices sets Prices field to given value.


### GetRequestedExpiresInSeconds

`func (o *DynamicChargeResponse) GetRequestedExpiresInSeconds() int32`

GetRequestedExpiresInSeconds returns the RequestedExpiresInSeconds field if non-nil, zero value otherwise.

### GetRequestedExpiresInSecondsOk

`func (o *DynamicChargeResponse) GetRequestedExpiresInSecondsOk() (*int32, bool)`

GetRequestedExpiresInSecondsOk returns a tuple with the RequestedExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedExpiresInSeconds

`func (o *DynamicChargeResponse) SetRequestedExpiresInSeconds(v int32)`

SetRequestedExpiresInSeconds sets RequestedExpiresInSeconds field to given value.


### GetMetadata

`func (o *DynamicChargeResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DynamicChargeResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DynamicChargeResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMetadataDigest

`func (o *DynamicChargeResponse) GetMetadataDigest() string`

GetMetadataDigest returns the MetadataDigest field if non-nil, zero value otherwise.

### GetMetadataDigestOk

`func (o *DynamicChargeResponse) GetMetadataDigestOk() (*string, bool)`

GetMetadataDigestOk returns a tuple with the MetadataDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataDigest

`func (o *DynamicChargeResponse) SetMetadataDigest(v string)`

SetMetadataDigest sets MetadataDigest field to given value.


### GetPaymentRequired

`func (o *DynamicChargeResponse) GetPaymentRequired() interface{}`

GetPaymentRequired returns the PaymentRequired field if non-nil, zero value otherwise.

### GetPaymentRequiredOk

`func (o *DynamicChargeResponse) GetPaymentRequiredOk() (*interface{}, bool)`

GetPaymentRequiredOk returns a tuple with the PaymentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequired

`func (o *DynamicChargeResponse) SetPaymentRequired(v interface{})`

SetPaymentRequired sets PaymentRequired field to given value.


### SetPaymentRequiredNil

`func (o *DynamicChargeResponse) SetPaymentRequiredNil(b bool)`

 SetPaymentRequiredNil sets the value for PaymentRequired to be an explicit nil

### UnsetPaymentRequired
`func (o *DynamicChargeResponse) UnsetPaymentRequired()`

UnsetPaymentRequired ensures that no value is present for PaymentRequired, not even an explicit nil
### GetPaymentRequiredHeader

`func (o *DynamicChargeResponse) GetPaymentRequiredHeader() string`

GetPaymentRequiredHeader returns the PaymentRequiredHeader field if non-nil, zero value otherwise.

### GetPaymentRequiredHeaderOk

`func (o *DynamicChargeResponse) GetPaymentRequiredHeaderOk() (*string, bool)`

GetPaymentRequiredHeaderOk returns a tuple with the PaymentRequiredHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequiredHeader

`func (o *DynamicChargeResponse) SetPaymentRequiredHeader(v string)`

SetPaymentRequiredHeader sets PaymentRequiredHeader field to given value.


### GetEligibleAlternatives

`func (o *DynamicChargeResponse) GetEligibleAlternatives() []PublicNetworkFeeAlternative`

GetEligibleAlternatives returns the EligibleAlternatives field if non-nil, zero value otherwise.

### GetEligibleAlternativesOk

`func (o *DynamicChargeResponse) GetEligibleAlternativesOk() (*[]PublicNetworkFeeAlternative, bool)`

GetEligibleAlternativesOk returns a tuple with the EligibleAlternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleAlternatives

`func (o *DynamicChargeResponse) SetEligibleAlternatives(v []PublicNetworkFeeAlternative)`

SetEligibleAlternatives sets EligibleAlternatives field to given value.


### GetFeePolicy

`func (o *DynamicChargeResponse) GetFeePolicy() PublicFeePolicyDocument`

GetFeePolicy returns the FeePolicy field if non-nil, zero value otherwise.

### GetFeePolicyOk

`func (o *DynamicChargeResponse) GetFeePolicyOk() (*PublicFeePolicyDocument, bool)`

GetFeePolicyOk returns a tuple with the FeePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicy

`func (o *DynamicChargeResponse) SetFeePolicy(v PublicFeePolicyDocument)`

SetFeePolicy sets FeePolicy field to given value.


### GetFeeQuoteDigest

`func (o *DynamicChargeResponse) GetFeeQuoteDigest() string`

GetFeeQuoteDigest returns the FeeQuoteDigest field if non-nil, zero value otherwise.

### GetFeeQuoteDigestOk

`func (o *DynamicChargeResponse) GetFeeQuoteDigestOk() (*string, bool)`

GetFeeQuoteDigestOk returns a tuple with the FeeQuoteDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeQuoteDigest

`func (o *DynamicChargeResponse) SetFeeQuoteDigest(v string)`

SetFeeQuoteDigest sets FeeQuoteDigest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
