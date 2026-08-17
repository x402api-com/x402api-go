# ExternalAddressControlChallengeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** |  |
**AssetId** | **string** |  |
**Address** | **string** |  |
**ProofMethod** | [**ExternalAddressProofInputMethodEnum**](ExternalAddressProofInputMethodEnum.md) |  |

## Methods

### NewExternalAddressControlChallengeCreate

`func NewExternalAddressControlChallengeCreate(network string, assetId string, address string, proofMethod ExternalAddressProofInputMethodEnum, ) *ExternalAddressControlChallengeCreate`

NewExternalAddressControlChallengeCreate instantiates a new ExternalAddressControlChallengeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAddressControlChallengeCreateWithDefaults

`func NewExternalAddressControlChallengeCreateWithDefaults() *ExternalAddressControlChallengeCreate`

NewExternalAddressControlChallengeCreateWithDefaults instantiates a new ExternalAddressControlChallengeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ExternalAddressControlChallengeCreate) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExternalAddressControlChallengeCreate) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExternalAddressControlChallengeCreate) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetAssetId

`func (o *ExternalAddressControlChallengeCreate) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ExternalAddressControlChallengeCreate) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ExternalAddressControlChallengeCreate) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetAddress

`func (o *ExternalAddressControlChallengeCreate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExternalAddressControlChallengeCreate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExternalAddressControlChallengeCreate) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProofMethod

`func (o *ExternalAddressControlChallengeCreate) GetProofMethod() ExternalAddressProofInputMethodEnum`

GetProofMethod returns the ProofMethod field if non-nil, zero value otherwise.

### GetProofMethodOk

`func (o *ExternalAddressControlChallengeCreate) GetProofMethodOk() (*ExternalAddressProofInputMethodEnum, bool)`

GetProofMethodOk returns a tuple with the ProofMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofMethod

`func (o *ExternalAddressControlChallengeCreate) SetProofMethod(v ExternalAddressProofInputMethodEnum)`

SetProofMethod sets ProofMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
