# NetworkFeePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeePolicy** | [**PublicFeePolicyDocument**](PublicFeePolicyDocument.md) |  |
**Alternatives** | [**[]PublicNetworkFeeAlternative**](PublicNetworkFeeAlternative.md) |  |
**FeeQuoteDigest** | **string** |  |

## Methods

### NewNetworkFeePreviewResponse

`func NewNetworkFeePreviewResponse(feePolicy PublicFeePolicyDocument, alternatives []PublicNetworkFeeAlternative, feeQuoteDigest string, ) *NetworkFeePreviewResponse`

NewNetworkFeePreviewResponse instantiates a new NetworkFeePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeePreviewResponseWithDefaults

`func NewNetworkFeePreviewResponseWithDefaults() *NetworkFeePreviewResponse`

NewNetworkFeePreviewResponseWithDefaults instantiates a new NetworkFeePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeePolicy

`func (o *NetworkFeePreviewResponse) GetFeePolicy() PublicFeePolicyDocument`

GetFeePolicy returns the FeePolicy field if non-nil, zero value otherwise.

### GetFeePolicyOk

`func (o *NetworkFeePreviewResponse) GetFeePolicyOk() (*PublicFeePolicyDocument, bool)`

GetFeePolicyOk returns a tuple with the FeePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePolicy

`func (o *NetworkFeePreviewResponse) SetFeePolicy(v PublicFeePolicyDocument)`

SetFeePolicy sets FeePolicy field to given value.


### GetAlternatives

`func (o *NetworkFeePreviewResponse) GetAlternatives() []PublicNetworkFeeAlternative`

GetAlternatives returns the Alternatives field if non-nil, zero value otherwise.

### GetAlternativesOk

`func (o *NetworkFeePreviewResponse) GetAlternativesOk() (*[]PublicNetworkFeeAlternative, bool)`

GetAlternativesOk returns a tuple with the Alternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatives

`func (o *NetworkFeePreviewResponse) SetAlternatives(v []PublicNetworkFeeAlternative)`

SetAlternatives sets Alternatives field to given value.


### GetFeeQuoteDigest

`func (o *NetworkFeePreviewResponse) GetFeeQuoteDigest() string`

GetFeeQuoteDigest returns the FeeQuoteDigest field if non-nil, zero value otherwise.

### GetFeeQuoteDigestOk

`func (o *NetworkFeePreviewResponse) GetFeeQuoteDigestOk() (*string, bool)`

GetFeeQuoteDigestOk returns a tuple with the FeeQuoteDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeQuoteDigest

`func (o *NetworkFeePreviewResponse) SetFeeQuoteDigest(v string)`

SetFeeQuoteDigest sets FeeQuoteDigest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
