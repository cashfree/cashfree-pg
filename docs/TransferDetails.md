# TransferDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** |  | [optional] 
**TransferFrom** | Pointer to **string** |  | [optional] 
**TransferType** | Pointer to **string** |  | [optional] 
**TransferAmount** | Pointer to **float32** |  | [optional] 
**Remark** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TransferDetailsTagsInner**](TransferDetailsTagsInner.md) |  | [optional] 

## Methods

### NewTransferDetails

`func NewTransferDetails() *TransferDetails`

NewTransferDetails instantiates a new TransferDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDetailsWithDefaults

`func NewTransferDetailsWithDefaults() *TransferDetails`

NewTransferDetailsWithDefaults instantiates a new TransferDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *TransferDetails) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *TransferDetails) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *TransferDetails) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *TransferDetails) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetTransferFrom

`func (o *TransferDetails) GetTransferFrom() string`

GetTransferFrom returns the TransferFrom field if non-nil, zero value otherwise.

### GetTransferFromOk

`func (o *TransferDetails) GetTransferFromOk() (*string, bool)`

GetTransferFromOk returns a tuple with the TransferFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFrom

`func (o *TransferDetails) SetTransferFrom(v string)`

SetTransferFrom sets TransferFrom field to given value.

### HasTransferFrom

`func (o *TransferDetails) HasTransferFrom() bool`

HasTransferFrom returns a boolean if a field has been set.

### GetTransferType

`func (o *TransferDetails) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *TransferDetails) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *TransferDetails) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *TransferDetails) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetTransferAmount

`func (o *TransferDetails) GetTransferAmount() float32`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *TransferDetails) GetTransferAmountOk() (*float32, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *TransferDetails) SetTransferAmount(v float32)`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *TransferDetails) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.

### GetRemark

`func (o *TransferDetails) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *TransferDetails) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *TransferDetails) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *TransferDetails) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetTags

`func (o *TransferDetails) GetTags() []TransferDetailsTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransferDetails) GetTagsOk() (*[]TransferDetailsTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransferDetails) SetTags(v []TransferDetailsTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TransferDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


