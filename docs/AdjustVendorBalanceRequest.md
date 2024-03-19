# AdjustVendorBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferFrom** | **string** | Mention to whom you want to transfer the on demand balance. Possible values - MERCHANT, VENDOR. | 
**TransferType** | **string** | Mention the type of transfer. Possible values: ON_DEMAND. | 
**TransferAmount** | **float32** | Mention the on demand transfer amount. | 
**Remark** | Pointer to **string** | Mention remarks if any for the on demand transfer. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Provide additional data fields using tags. | [optional] 

## Methods

### NewAdjustVendorBalanceRequest

`func NewAdjustVendorBalanceRequest(transferFrom string, transferType string, transferAmount float32, ) *AdjustVendorBalanceRequest`

NewAdjustVendorBalanceRequest instantiates a new AdjustVendorBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustVendorBalanceRequestWithDefaults

`func NewAdjustVendorBalanceRequestWithDefaults() *AdjustVendorBalanceRequest`

NewAdjustVendorBalanceRequestWithDefaults instantiates a new AdjustVendorBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferFrom

`func (o *AdjustVendorBalanceRequest) GetTransferFrom() string`

GetTransferFrom returns the TransferFrom field if non-nil, zero value otherwise.

### GetTransferFromOk

`func (o *AdjustVendorBalanceRequest) GetTransferFromOk() (*string, bool)`

GetTransferFromOk returns a tuple with the TransferFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFrom

`func (o *AdjustVendorBalanceRequest) SetTransferFrom(v string)`

SetTransferFrom sets TransferFrom field to given value.


### GetTransferType

`func (o *AdjustVendorBalanceRequest) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *AdjustVendorBalanceRequest) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *AdjustVendorBalanceRequest) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.


### GetTransferAmount

`func (o *AdjustVendorBalanceRequest) GetTransferAmount() float32`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *AdjustVendorBalanceRequest) GetTransferAmountOk() (*float32, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *AdjustVendorBalanceRequest) SetTransferAmount(v float32)`

SetTransferAmount sets TransferAmount field to given value.


### GetRemark

`func (o *AdjustVendorBalanceRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *AdjustVendorBalanceRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *AdjustVendorBalanceRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *AdjustVendorBalanceRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetTags

`func (o *AdjustVendorBalanceRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdjustVendorBalanceRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdjustVendorBalanceRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdjustVendorBalanceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


