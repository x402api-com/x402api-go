# ExternalAddressControlChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly]
**Network** | **string** |  | [readonly]
**AssetId** | **string** |  | [readonly]
**AddressDisplay** | **string** |  | [readonly]
**ProofMethod** | [**ExternalAddressProofMethodEnum**](ExternalAddressProofMethodEnum.md) |  | [readonly]
**Message** | **string** |  | [readonly]
**ChallengeDigest** | **string** |  | [readonly]
**CanaryInstructions** | **interface{}** |  | [readonly]
**ExpiresAt** | **time.Time** |  | [readonly]
**ConsumedAt** | **NullableTime** |  | [readonly]
**CreatedAt** | **time.Time** |  | [readonly]

## Methods

### NewExternalAddressControlChallenge

`func NewExternalAddressControlChallenge(id string, network string, assetId string, addressDisplay string, proofMethod ExternalAddressProofMethodEnum, message string, challengeDigest string, canaryInstructions interface{}, expiresAt time.Time, consumedAt NullableTime, createdAt time.Time, ) *ExternalAddressControlChallenge`

NewExternalAddressControlChallenge instantiates a new ExternalAddressControlChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAddressControlChallengeWithDefaults

`func NewExternalAddressControlChallengeWithDefaults() *ExternalAddressControlChallenge`

NewExternalAddressControlChallengeWithDefaults instantiates a new ExternalAddressControlChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalAddressControlChallenge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalAddressControlChallenge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalAddressControlChallenge) SetId(v string)`

SetId sets Id field to given value.


### GetNetwork

`func (o *ExternalAddressControlChallenge) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExternalAddressControlChallenge) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExternalAddressControlChallenge) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *ExternalAddressControlChallenge) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ExternalAddressControlChallenge) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ExternalAddressControlChallenge) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAddressDisplay

`func (o *ExternalAddressControlChallenge) GetAddressDisplay() string`

GetAddressDisplay returns the AddressDisplay field if non-nil, zero value otherwise.

### GetAddressDisplayOk

`func (o *ExternalAddressControlChallenge) GetAddressDisplayOk() (*string, bool)`

GetAddressDisplayOk returns a tuple with the AddressDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressDisplay

`func (o *ExternalAddressControlChallenge) SetAddressDisplay(v string)`

SetAddressDisplay sets AddressDisplay field to given value.


### GetProofMethod

`func (o *ExternalAddressControlChallenge) GetProofMethod() ExternalAddressProofMethodEnum`

GetProofMethod returns the ProofMethod field if non-nil, zero value otherwise.

### GetProofMethodOk

`func (o *ExternalAddressControlChallenge) GetProofMethodOk() (*ExternalAddressProofMethodEnum, bool)`

GetProofMethodOk returns a tuple with the ProofMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofMethod

`func (o *ExternalAddressControlChallenge) SetProofMethod(v ExternalAddressProofMethodEnum)`

SetProofMethod sets ProofMethod field to given value.


### GetMessage

`func (o *ExternalAddressControlChallenge) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExternalAddressControlChallenge) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExternalAddressControlChallenge) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetChallengeDigest

`func (o *ExternalAddressControlChallenge) GetChallengeDigest() string`

GetChallengeDigest returns the ChallengeDigest field if non-nil, zero value otherwise.

### GetChallengeDigestOk

`func (o *ExternalAddressControlChallenge) GetChallengeDigestOk() (*string, bool)`

GetChallengeDigestOk returns a tuple with the ChallengeDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeDigest

`func (o *ExternalAddressControlChallenge) SetChallengeDigest(v string)`

SetChallengeDigest sets ChallengeDigest field to given value.


### GetCanaryInstructions

`func (o *ExternalAddressControlChallenge) GetCanaryInstructions() interface{}`

GetCanaryInstructions returns the CanaryInstructions field if non-nil, zero value otherwise.

### GetCanaryInstructionsOk

`func (o *ExternalAddressControlChallenge) GetCanaryInstructionsOk() (*interface{}, bool)`

GetCanaryInstructionsOk returns a tuple with the CanaryInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryInstructions

`func (o *ExternalAddressControlChallenge) SetCanaryInstructions(v interface{})`

SetCanaryInstructions sets CanaryInstructions field to given value.


### SetCanaryInstructionsNil

`func (o *ExternalAddressControlChallenge) SetCanaryInstructionsNil(b bool)`

 SetCanaryInstructionsNil sets the value for CanaryInstructions to be an explicit nil

### UnsetCanaryInstructions
`func (o *ExternalAddressControlChallenge) UnsetCanaryInstructions()`

UnsetCanaryInstructions ensures that no value is present for CanaryInstructions, not even an explicit nil
### GetExpiresAt

`func (o *ExternalAddressControlChallenge) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExternalAddressControlChallenge) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExternalAddressControlChallenge) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetConsumedAt

`func (o *ExternalAddressControlChallenge) GetConsumedAt() time.Time`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *ExternalAddressControlChallenge) GetConsumedAtOk() (*time.Time, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *ExternalAddressControlChallenge) SetConsumedAt(v time.Time)`

SetConsumedAt sets ConsumedAt field to given value.


### SetConsumedAtNil

`func (o *ExternalAddressControlChallenge) SetConsumedAtNil(b bool)`

 SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil

### UnsetConsumedAt
`func (o *ExternalAddressControlChallenge) UnsetConsumedAt()`

UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
### GetCreatedAt

`func (o *ExternalAddressControlChallenge) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalAddressControlChallenge) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalAddressControlChallenge) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
