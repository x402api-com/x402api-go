# WalletFencedChainReseedContext

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
**WalletVersionId** | **string** |  |
**WalletVersion** | **int32** |  |
**WalletAddress** | **string** |  |
**WalletVersionState** | [**WalletVersionStateEnum**](WalletVersionStateEnum.md) |  |

## Methods

### NewWalletFencedChainReseedContext

`func NewWalletFencedChainReseedContext(checkpointId string, network string, finality WalletObservationFinalityEnum, manifestDigest string, policyDigest string, expectedGeneration int32, expectedNextBlockNumber string, expectedLastScannedBlockNumber NullableString, expectedLastScannedBlockHash string, expectedReviewRequiredAt time.Time, expectedReviewErrorCode string, observedAt time.Time, walletVersionId string, walletVersion int32, walletAddress string, walletVersionState WalletVersionStateEnum, ) *WalletFencedChainReseedContext`

NewWalletFencedChainReseedContext instantiates a new WalletFencedChainReseedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletFencedChainReseedContextWithDefaults

`func NewWalletFencedChainReseedContextWithDefaults() *WalletFencedChainReseedContext`

NewWalletFencedChainReseedContextWithDefaults instantiates a new WalletFencedChainReseedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpointId

`func (o *WalletFencedChainReseedContext) GetCheckpointId() string`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *WalletFencedChainReseedContext) GetCheckpointIdOk() (*string, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *WalletFencedChainReseedContext) SetCheckpointId(v string)`

SetCheckpointId sets CheckpointId field to given value.


### GetNetwork

`func (o *WalletFencedChainReseedContext) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WalletFencedChainReseedContext) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WalletFencedChainReseedContext) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetFinality

`func (o *WalletFencedChainReseedContext) GetFinality() WalletObservationFinalityEnum`

GetFinality returns the Finality field if non-nil, zero value otherwise.

### GetFinalityOk

`func (o *WalletFencedChainReseedContext) GetFinalityOk() (*WalletObservationFinalityEnum, bool)`

GetFinalityOk returns a tuple with the Finality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinality

`func (o *WalletFencedChainReseedContext) SetFinality(v WalletObservationFinalityEnum)`

SetFinality sets Finality field to given value.


### GetManifestDigest

`func (o *WalletFencedChainReseedContext) GetManifestDigest() string`

GetManifestDigest returns the ManifestDigest field if non-nil, zero value otherwise.

### GetManifestDigestOk

`func (o *WalletFencedChainReseedContext) GetManifestDigestOk() (*string, bool)`

GetManifestDigestOk returns a tuple with the ManifestDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestDigest

`func (o *WalletFencedChainReseedContext) SetManifestDigest(v string)`

SetManifestDigest sets ManifestDigest field to given value.


### GetPolicyDigest

`func (o *WalletFencedChainReseedContext) GetPolicyDigest() string`

GetPolicyDigest returns the PolicyDigest field if non-nil, zero value otherwise.

### GetPolicyDigestOk

`func (o *WalletFencedChainReseedContext) GetPolicyDigestOk() (*string, bool)`

GetPolicyDigestOk returns a tuple with the PolicyDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDigest

`func (o *WalletFencedChainReseedContext) SetPolicyDigest(v string)`

SetPolicyDigest sets PolicyDigest field to given value.


### GetExpectedGeneration

`func (o *WalletFencedChainReseedContext) GetExpectedGeneration() int32`

GetExpectedGeneration returns the ExpectedGeneration field if non-nil, zero value otherwise.

### GetExpectedGenerationOk

`func (o *WalletFencedChainReseedContext) GetExpectedGenerationOk() (*int32, bool)`

GetExpectedGenerationOk returns a tuple with the ExpectedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedGeneration

`func (o *WalletFencedChainReseedContext) SetExpectedGeneration(v int32)`

SetExpectedGeneration sets ExpectedGeneration field to given value.


### GetExpectedNextBlockNumber

`func (o *WalletFencedChainReseedContext) GetExpectedNextBlockNumber() string`

GetExpectedNextBlockNumber returns the ExpectedNextBlockNumber field if non-nil, zero value otherwise.

### GetExpectedNextBlockNumberOk

`func (o *WalletFencedChainReseedContext) GetExpectedNextBlockNumberOk() (*string, bool)`

GetExpectedNextBlockNumberOk returns a tuple with the ExpectedNextBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNextBlockNumber

`func (o *WalletFencedChainReseedContext) SetExpectedNextBlockNumber(v string)`

SetExpectedNextBlockNumber sets ExpectedNextBlockNumber field to given value.


### GetExpectedLastScannedBlockNumber

`func (o *WalletFencedChainReseedContext) GetExpectedLastScannedBlockNumber() string`

GetExpectedLastScannedBlockNumber returns the ExpectedLastScannedBlockNumber field if non-nil, zero value otherwise.

### GetExpectedLastScannedBlockNumberOk

`func (o *WalletFencedChainReseedContext) GetExpectedLastScannedBlockNumberOk() (*string, bool)`

GetExpectedLastScannedBlockNumberOk returns a tuple with the ExpectedLastScannedBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLastScannedBlockNumber

`func (o *WalletFencedChainReseedContext) SetExpectedLastScannedBlockNumber(v string)`

SetExpectedLastScannedBlockNumber sets ExpectedLastScannedBlockNumber field to given value.


### SetExpectedLastScannedBlockNumberNil

`func (o *WalletFencedChainReseedContext) SetExpectedLastScannedBlockNumberNil(b bool)`

 SetExpectedLastScannedBlockNumberNil sets the value for ExpectedLastScannedBlockNumber to be an explicit nil

### UnsetExpectedLastScannedBlockNumber
`func (o *WalletFencedChainReseedContext) UnsetExpectedLastScannedBlockNumber()`

UnsetExpectedLastScannedBlockNumber ensures that no value is present for ExpectedLastScannedBlockNumber, not even an explicit nil
### GetExpectedLastScannedBlockHash

`func (o *WalletFencedChainReseedContext) GetExpectedLastScannedBlockHash() string`

GetExpectedLastScannedBlockHash returns the ExpectedLastScannedBlockHash field if non-nil, zero value otherwise.

### GetExpectedLastScannedBlockHashOk

`func (o *WalletFencedChainReseedContext) GetExpectedLastScannedBlockHashOk() (*string, bool)`

GetExpectedLastScannedBlockHashOk returns a tuple with the ExpectedLastScannedBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLastScannedBlockHash

`func (o *WalletFencedChainReseedContext) SetExpectedLastScannedBlockHash(v string)`

SetExpectedLastScannedBlockHash sets ExpectedLastScannedBlockHash field to given value.


### GetExpectedReviewRequiredAt

`func (o *WalletFencedChainReseedContext) GetExpectedReviewRequiredAt() time.Time`

GetExpectedReviewRequiredAt returns the ExpectedReviewRequiredAt field if non-nil, zero value otherwise.

### GetExpectedReviewRequiredAtOk

`func (o *WalletFencedChainReseedContext) GetExpectedReviewRequiredAtOk() (*time.Time, bool)`

GetExpectedReviewRequiredAtOk returns a tuple with the ExpectedReviewRequiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReviewRequiredAt

`func (o *WalletFencedChainReseedContext) SetExpectedReviewRequiredAt(v time.Time)`

SetExpectedReviewRequiredAt sets ExpectedReviewRequiredAt field to given value.


### GetExpectedReviewErrorCode

`func (o *WalletFencedChainReseedContext) GetExpectedReviewErrorCode() string`

GetExpectedReviewErrorCode returns the ExpectedReviewErrorCode field if non-nil, zero value otherwise.

### GetExpectedReviewErrorCodeOk

`func (o *WalletFencedChainReseedContext) GetExpectedReviewErrorCodeOk() (*string, bool)`

GetExpectedReviewErrorCodeOk returns a tuple with the ExpectedReviewErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReviewErrorCode

`func (o *WalletFencedChainReseedContext) SetExpectedReviewErrorCode(v string)`

SetExpectedReviewErrorCode sets ExpectedReviewErrorCode field to given value.


### GetObservedAt

`func (o *WalletFencedChainReseedContext) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *WalletFencedChainReseedContext) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *WalletFencedChainReseedContext) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.


### GetWalletVersionId

`func (o *WalletFencedChainReseedContext) GetWalletVersionId() string`

GetWalletVersionId returns the WalletVersionId field if non-nil, zero value otherwise.

### GetWalletVersionIdOk

`func (o *WalletFencedChainReseedContext) GetWalletVersionIdOk() (*string, bool)`

GetWalletVersionIdOk returns a tuple with the WalletVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionId

`func (o *WalletFencedChainReseedContext) SetWalletVersionId(v string)`

SetWalletVersionId sets WalletVersionId field to given value.


### GetWalletVersion

`func (o *WalletFencedChainReseedContext) GetWalletVersion() int32`

GetWalletVersion returns the WalletVersion field if non-nil, zero value otherwise.

### GetWalletVersionOk

`func (o *WalletFencedChainReseedContext) GetWalletVersionOk() (*int32, bool)`

GetWalletVersionOk returns a tuple with the WalletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersion

`func (o *WalletFencedChainReseedContext) SetWalletVersion(v int32)`

SetWalletVersion sets WalletVersion field to given value.


### GetWalletAddress

`func (o *WalletFencedChainReseedContext) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *WalletFencedChainReseedContext) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *WalletFencedChainReseedContext) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetWalletVersionState

`func (o *WalletFencedChainReseedContext) GetWalletVersionState() WalletVersionStateEnum`

GetWalletVersionState returns the WalletVersionState field if non-nil, zero value otherwise.

### GetWalletVersionStateOk

`func (o *WalletFencedChainReseedContext) GetWalletVersionStateOk() (*WalletVersionStateEnum, bool)`

GetWalletVersionStateOk returns a tuple with the WalletVersionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletVersionState

`func (o *WalletFencedChainReseedContext) SetWalletVersionState(v WalletVersionStateEnum)`

SetWalletVersionState sets WalletVersionState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
