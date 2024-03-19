# StaticSplitRequestSchemeInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantVendorId** | Pointer to **string** | Specify the merchant vendor ID to create the split scheme for the payment. | [optional] 
**Percentage** | Pointer to **string** | Specify the percentage of amount to be split. | [optional] 

## Methods

### NewStaticSplitRequestSchemeInner

`func NewStaticSplitRequestSchemeInner() *StaticSplitRequestSchemeInner`

NewStaticSplitRequestSchemeInner instantiates a new StaticSplitRequestSchemeInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSplitRequestSchemeInnerWithDefaults

`func NewStaticSplitRequestSchemeInnerWithDefaults() *StaticSplitRequestSchemeInner`

NewStaticSplitRequestSchemeInnerWithDefaults instantiates a new StaticSplitRequestSchemeInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantVendorId

`func (o *StaticSplitRequestSchemeInner) GetMerchantVendorId() string`

GetMerchantVendorId returns the MerchantVendorId field if non-nil, zero value otherwise.

### GetMerchantVendorIdOk

`func (o *StaticSplitRequestSchemeInner) GetMerchantVendorIdOk() (*string, bool)`

GetMerchantVendorIdOk returns a tuple with the MerchantVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantVendorId

`func (o *StaticSplitRequestSchemeInner) SetMerchantVendorId(v string)`

SetMerchantVendorId sets MerchantVendorId field to given value.

### HasMerchantVendorId

`func (o *StaticSplitRequestSchemeInner) HasMerchantVendorId() bool`

HasMerchantVendorId returns a boolean if a field has been set.

### GetPercentage

`func (o *StaticSplitRequestSchemeInner) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *StaticSplitRequestSchemeInner) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *StaticSplitRequestSchemeInner) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *StaticSplitRequestSchemeInner) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


