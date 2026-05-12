# OrderCreateRefundRequestRefundSplitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | Vendor id created in Cashfree system. | 
**Amount** | Pointer to **float32** | Amount which will be associated with this vendor. | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** | Custom Tags in the form of {\&quot;key\&quot;:\&quot;value\&quot;} which can be passed for an order. A maximum of 10 tags can be added. | [optional] 

## Methods

### NewOrderCreateRefundRequestRefundSplitsInner

`func NewOrderCreateRefundRequestRefundSplitsInner(vendorId string, ) *OrderCreateRefundRequestRefundSplitsInner`

NewOrderCreateRefundRequestRefundSplitsInner instantiates a new OrderCreateRefundRequestRefundSplitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRefundRequestRefundSplitsInnerWithDefaults

`func NewOrderCreateRefundRequestRefundSplitsInnerWithDefaults() *OrderCreateRefundRequestRefundSplitsInner`

NewOrderCreateRefundRequestRefundSplitsInnerWithDefaults instantiates a new OrderCreateRefundRequestRefundSplitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *OrderCreateRefundRequestRefundSplitsInner) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetAmount

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderCreateRefundRequestRefundSplitsInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderCreateRefundRequestRefundSplitsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTags

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderCreateRefundRequestRefundSplitsInner) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderCreateRefundRequestRefundSplitsInner) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderCreateRefundRequestRefundSplitsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


