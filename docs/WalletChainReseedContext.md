# WalletChainReseedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckpointId** | **string** |  |
**Network** | **string** |  |
**Finality** | [**WalletObservationFinalityEnum**](WalletObservationFinalityEnum.md) |  |
**ManifestDigest** | **string** |  |
**PolicyDigest** | **string** |  |
**ExpectedGeneration** | **int32** |  |
**ExpectedNextBlockNumber** | **string** |  |
**ExpectedLastScannedBlockNumber** | **NullableString** |  |
**ExpectedLastScannedBlockHash** | **string** |  |
**ExpectedReviewRequiredAt** | **time.Time** |  |
**ExpectedReviewErrorCode** | **string** |  |
**ObservedAt** | **time.Time** |  |

## Methods

### NewWalletChainReseedContext

`func NewWalletChainReseedContext(checkpointId string, network string, finality WalletObservationFinalityEnum, manifestDigest string, policyDigest string, expectedGeneration int32, expectedNextBlockNumber string, expectedLastScannedBlockNumber NullableString, expectedLastScannedBlockHash string, expectedReviewRequiredAt time.Time, expectedReviewErrorCode string, observedAt time.Time, ) *WalletChainReseedContext`

NewWalletChainReseedContext instantiates a new WalletChainReseedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletChainReseedContextWithDefaults

`func NewWalletChainReseedContextWithDefaults() *WalletChainReseedContext`

NewWalletChainReseedContextWithDefaults instantiates a new WalletChainReseedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointId

`func (o *WalletChainReseedContext) GetCheckpointId() string`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *WalletChainReseedContext) GetCheckpointIdOk() (*string, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *WalletChainReseedContext) SetCheckpointId(v string)`

SetCheckpointId sets CheckpointId field to given value.


### GetNetwork

`func (o *WalletChainReseedContext) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WalletChainReseedContext) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WalletChainReseedContext) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFinality

`func (o *WalletChainReseedContext) GetFinality() WalletObservationFinalityEnum`

GetFinality returns the Finality field if non-nil, zero value otherwise.

### GetFinalityOk

`func (o *WalletChainReseedContext) GetFinalityOk() (*WalletObservationFinalityEnum, bool)`

GetFinalityOk returns a tuple with the Finality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinality

`func (o *WalletChainReseedContext) SetFinality(v WalletObservationFinalityEnum)`

SetFinality sets Finality field to given value.


### GetManifestDigest

`func (o *WalletChainReseedContext) GetManifestDigest() string`

GetManifestDigest returns the ManifestDigest field if non-nil, zero value otherwise.

### GetManifestDigestOk

`func (o *WalletChainReseedContext) GetManifestDigestOk() (*string, bool)`

GetManifestDigestOk returns a tuple with the ManifestDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestDigest

`func (o *WalletChainReseedContext) SetManifestDigest(v string)`

SetManifestDigest sets ManifestDigest field to given value.


### GetPolicyDigest

`func (o *WalletChainReseedContext) GetPolicyDigest() string`

GetPolicyDigest returns the PolicyDigest field if non-nil, zero value otherwise.

### GetPolicyDigestOk

`func (o *WalletChainReseedContext) GetPolicyDigestOk() (*string, bool)`

GetPolicyDigestOk returns a tuple with the PolicyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDigest

`func (o *WalletChainReseedContext) SetPolicyDigest(v string)`

SetPolicyDigest sets PolicyDigest field to given value.


### GetExpectedGeneration

`func (o *WalletChainReseedContext) GetExpectedGeneration() int32`

GetExpectedGeneration returns the ExpectedGeneration field if non-nil, zero value otherwise.

### GetExpectedGenerationOk

`func (o *WalletChainReseedContext) GetExpectedGenerationOk() (*int32, bool)`

GetExpectedGenerationOk returns a tuple with the ExpectedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedGeneration

`func (o *WalletChainReseedContext) SetExpectedGeneration(v int32)`

SetExpectedGeneration sets ExpectedGeneration field to given value.


### GetExpectedNextBlockNumber

`func (o *WalletChainReseedContext) GetExpectedNextBlockNumber() string`

GetExpectedNextBlockNumber returns the ExpectedNextBlockNumber field if non-nil, zero value otherwise.

### GetExpectedNextBlockNumberOk

`func (o *WalletChainReseedContext) GetExpectedNextBlockNumberOk() (*string, bool)`

GetExpectedNextBlockNumberOk returns a tuple with the ExpectedNextBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextBlockNumber

`func (o *WalletChainReseedContext) SetExpectedNextBlockNumber(v string)`

SetExpectedNextBlockNumber sets ExpectedNextBlockNumber field to given value.


### GetExpectedLastScannedBlockNumber

`func (o *WalletChainReseedContext) GetExpectedLastScannedBlockNumber() string`

GetExpectedLastScannedBlockNumber returns the ExpectedLastScannedBlockNumber field if non-nil, zero value otherwise.

### GetExpectedLastScannedBlockNumberOk

`func (o *WalletChainReseedContext) GetExpectedLastScannedBlockNumberOk() (*string, bool)`

GetExpectedLastScannedBlockNumberOk returns a tuple with the ExpectedLastScannedBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLastScannedBlockNumber

`func (o *WalletChainReseedContext) SetExpectedLastScannedBlockNumber(v string)`

SetExpectedLastScannedBlockNumber sets ExpectedLastScannedBlockNumber field to given value.


### SetExpectedLastScannedBlockNumberNil

`func (o *WalletChainReseedContext) SetExpectedLastScannedBlockNumberNil(b bool)`

 SetExpectedLastScannedBlockNumberNil sets the value for ExpectedLastScannedBlockNumber to be an explicit nil

### UnsetExpectedLastScannedBlockNumber
`func (o *WalletChainReseedContext) UnsetExpectedLastScannedBlockNumber()`

UnsetExpectedLastScannedBlockNumber ensures that no value is present for ExpectedLastScannedBlockNumber, not even an explicit nil
### GetExpectedLastScannedBlockHash

`func (o *WalletChainReseedContext) GetExpectedLastScannedBlockHash() string`

GetExpectedLastScannedBlockHash returns the ExpectedLastScannedBlockHash field if non-nil, zero value otherwise.

### GetExpectedLastScannedBlockHashOk

`func (o *WalletChainReseedContext) GetExpectedLastScannedBlockHashOk() (*string, bool)`

GetExpectedLastScannedBlockHashOk returns a tuple with the ExpectedLastScannedBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLastScannedBlockHash

`func (o *WalletChainReseedContext) SetExpectedLastScannedBlockHash(v string)`

SetExpectedLastScannedBlockHash sets ExpectedLastScannedBlockHash field to given value.


### GetExpectedReviewRequiredAt

`func (o *WalletChainReseedContext) GetExpectedReviewRequiredAt() time.Time`

GetExpectedReviewRequiredAt returns the ExpectedReviewRequiredAt field if non-nil, zero value otherwise.

### GetExpectedReviewRequiredAtOk

`func (o *WalletChainReseedContext) GetExpectedReviewRequiredAtOk() (*time.Time, bool)`

GetExpectedReviewRequiredAtOk returns a tuple with the ExpectedReviewRequiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReviewRequiredAt

`func (o *WalletChainReseedContext) SetExpectedReviewRequiredAt(v time.Time)`

SetExpectedReviewRequiredAt sets ExpectedReviewRequiredAt field to given value.


### GetExpectedReviewErrorCode

`func (o *WalletChainReseedContext) GetExpectedReviewErrorCode() string`

GetExpectedReviewErrorCode returns the ExpectedReviewErrorCode field if non-nil, zero value otherwise.

### GetExpectedReviewErrorCodeOk

`func (o *WalletChainReseedContext) GetExpectedReviewErrorCodeOk() (*string, bool)`

GetExpectedReviewErrorCodeOk returns a tuple with the ExpectedReviewErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReviewErrorCode

`func (o *WalletChainReseedContext) SetExpectedReviewErrorCode(v string)`

SetExpectedReviewErrorCode sets ExpectedReviewErrorCode field to given value.


### GetObservedAt

`func (o *WalletChainReseedContext) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *WalletChainReseedContext) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *WalletChainReseedContext) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
