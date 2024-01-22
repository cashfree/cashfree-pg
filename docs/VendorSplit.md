# VendorSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** | Vendor id created in Cashfree system | [optional] 
**Amount** | Pointer to **float32** | Amount which will be associated with this vendor | [optional] 
**Percentage** | Pointer to **float32** | Percentage of order amount which shall get added to vendor account | [optional] 

## Methods

### NewVendorSplit

`func NewVendorSplit() *VendorSplit`

NewVendorSplit instantiates a new VendorSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorSplitWithDefaults

`func NewVendorSplitWithDefaults() *VendorSplit`

NewVendorSplitWithDefaults instantiates a new VendorSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *VendorSplit) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *VendorSplit) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *VendorSplit) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *VendorSplit) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetAmount

`func (o *VendorSplit) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VendorSplit) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VendorSplit) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VendorSplit) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPercentage

`func (o *VendorSplit) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *VendorSplit) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *VendorSplit) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *VendorSplit) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


