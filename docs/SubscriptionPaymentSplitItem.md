# SubscriptionPaymentSplitItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** | Vendor ID | [optional] 
**Percentage** | Pointer to **float32** | Percentage of the payment to be split to vendor | [optional] 

## Methods

### NewSubscriptionPaymentSplitItem

`func NewSubscriptionPaymentSplitItem() *SubscriptionPaymentSplitItem`

NewSubscriptionPaymentSplitItem instantiates a new SubscriptionPaymentSplitItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPaymentSplitItemWithDefaults

`func NewSubscriptionPaymentSplitItemWithDefaults() *SubscriptionPaymentSplitItem`

NewSubscriptionPaymentSplitItemWithDefaults instantiates a new SubscriptionPaymentSplitItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *SubscriptionPaymentSplitItem) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *SubscriptionPaymentSplitItem) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *SubscriptionPaymentSplitItem) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *SubscriptionPaymentSplitItem) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetPercentage

`func (o *SubscriptionPaymentSplitItem) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *SubscriptionPaymentSplitItem) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *SubscriptionPaymentSplitItem) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *SubscriptionPaymentSplitItem) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


