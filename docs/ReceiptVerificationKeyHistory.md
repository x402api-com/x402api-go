# ReceiptVerificationKeyHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Keys** | [**map[string]ReceiptVerificationKey**](ReceiptVerificationKey.md) |  |

## Methods

### NewReceiptVerificationKeyHistory

`func NewReceiptVerificationKeyHistory(type_ string, keys map[string]ReceiptVerificationKey, ) *ReceiptVerificationKeyHistory`

NewReceiptVerificationKeyHistory instantiates a new ReceiptVerificationKeyHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptVerificationKeyHistoryWithDefaults

`func NewReceiptVerificationKeyHistoryWithDefaults() *ReceiptVerificationKeyHistory`

NewReceiptVerificationKeyHistoryWithDefaults instantiates a new ReceiptVerificationKeyHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReceiptVerificationKeyHistory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReceiptVerificationKeyHistory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReceiptVerificationKeyHistory) SetType(v string)`

SetType sets Type field to given value.


### GetKeys

`func (o *ReceiptVerificationKeyHistory) GetKeys() map[string]ReceiptVerificationKey`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ReceiptVerificationKeyHistory) GetKeysOk() (*map[string]ReceiptVerificationKey, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ReceiptVerificationKeyHistory) SetKeys(v map[string]ReceiptVerificationKey)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
