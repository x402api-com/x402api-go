# ExternalReceivingAddressRotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeId** | **string** |  |
**Proof** | [**ExternalAddressControlProofInput**](ExternalAddressControlProofInput.md) |  |
**Reason** | **string** |  |

## Methods

### NewExternalReceivingAddressRotation

`func NewExternalReceivingAddressRotation(challengeId string, proof ExternalAddressControlProofInput, reason string, ) *ExternalReceivingAddressRotation`

NewExternalReceivingAddressRotation instantiates a new ExternalReceivingAddressRotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalReceivingAddressRotationWithDefaults

`func NewExternalReceivingAddressRotationWithDefaults() *ExternalReceivingAddressRotation`

NewExternalReceivingAddressRotationWithDefaults instantiates a new ExternalReceivingAddressRotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeId

`func (o *ExternalReceivingAddressRotation) GetChallengeId() string`

GetChallengeId returns the ChallengeId field if non-nil, zero value otherwise.

### GetChallengeIdOk

`func (o *ExternalReceivingAddressRotation) GetChallengeIdOk() (*string, bool)`

GetChallengeIdOk returns a tuple with the ChallengeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeId

`func (o *ExternalReceivingAddressRotation) SetChallengeId(v string)`

SetChallengeId sets ChallengeId field to given value.


### GetProof

`func (o *ExternalReceivingAddressRotation) GetProof() ExternalAddressControlProofInput`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ExternalReceivingAddressRotation) GetProofOk() (*ExternalAddressControlProofInput, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ExternalReceivingAddressRotation) SetProof(v ExternalAddressControlProofInput)`

SetProof sets Proof field to given value.


### GetReason

`func (o *ExternalReceivingAddressRotation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ExternalReceivingAddressRotation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ExternalReceivingAddressRotation) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
