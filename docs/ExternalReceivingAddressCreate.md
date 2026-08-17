# ExternalReceivingAddressCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  |
**ChallengeId** | **string** |  |
**Proof** | [**ExternalAddressControlProofInput**](ExternalAddressControlProofInput.md) |  |

## Methods

### NewExternalReceivingAddressCreate

`func NewExternalReceivingAddressCreate(label string, challengeId string, proof ExternalAddressControlProofInput, ) *ExternalReceivingAddressCreate`

NewExternalReceivingAddressCreate instantiates a new ExternalReceivingAddressCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalReceivingAddressCreateWithDefaults

`func NewExternalReceivingAddressCreateWithDefaults() *ExternalReceivingAddressCreate`

NewExternalReceivingAddressCreateWithDefaults instantiates a new ExternalReceivingAddressCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExternalReceivingAddressCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalReceivingAddressCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalReceivingAddressCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetChallengeId

`func (o *ExternalReceivingAddressCreate) GetChallengeId() string`

GetChallengeId returns the ChallengeId field if non-nil, zero value otherwise.

### GetChallengeIdOk

`func (o *ExternalReceivingAddressCreate) GetChallengeIdOk() (*string, bool)`

GetChallengeIdOk returns a tuple with the ChallengeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeId

`func (o *ExternalReceivingAddressCreate) SetChallengeId(v string)`

SetChallengeId sets ChallengeId field to given value.


### GetProof

`func (o *ExternalReceivingAddressCreate) GetProof() ExternalAddressControlProofInput`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *ExternalReceivingAddressCreate) GetProofOk() (*ExternalAddressControlProofInput, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *ExternalReceivingAddressCreate) SetProof(v ExternalAddressControlProofInput)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
