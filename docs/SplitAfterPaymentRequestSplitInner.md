# SplitAfterPaymentRequestSplitInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** | Specify the merchant vendor ID to split the payment. | [optional] 
**Amount** | Pointer to **float32** | Specify the amount to be split to the vendor. | [optional] 
**Percentage** | Pointer to **float32** | Specify the percentage of amount to be split. | [optional] 
**Tags** | Pointer to **map[string]string** | Custom Tags in thr form of {\&quot;key\&quot;:\&quot;value\&quot;} which can be passed for an order. A maximum of 10 tags can be added | [optional] 

## Methods

### NewSplitAfterPaymentRequestSplitInner

`func NewSplitAfterPaymentRequestSplitInner() *SplitAfterPaymentRequestSplitInner`

NewSplitAfterPaymentRequestSplitInner instantiates a new SplitAfterPaymentRequestSplitInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitAfterPaymentRequestSplitInnerWithDefaults

`func NewSplitAfterPaymentRequestSplitInnerWithDefaults() *SplitAfterPaymentRequestSplitInner`

NewSplitAfterPaymentRequestSplitInnerWithDefaults instantiates a new SplitAfterPaymentRequestSplitInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *SplitAfterPaymentRequestSplitInner) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *SplitAfterPaymentRequestSplitInner) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *SplitAfterPaymentRequestSplitInner) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *SplitAfterPaymentRequestSplitInner) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetAmount

`func (o *SplitAfterPaymentRequestSplitInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SplitAfterPaymentRequestSplitInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SplitAfterPaymentRequestSplitInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SplitAfterPaymentRequestSplitInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPercentage

`func (o *SplitAfterPaymentRequestSplitInner) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *SplitAfterPaymentRequestSplitInner) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *SplitAfterPaymentRequestSplitInner) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *SplitAfterPaymentRequestSplitInner) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetTags

`func (o *SplitAfterPaymentRequestSplitInner) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SplitAfterPaymentRequestSplitInner) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SplitAfterPaymentRequestSplitInner) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SplitAfterPaymentRequestSplitInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


