# CartDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCharge** | Pointer to **float64** |  | [optional] 
**CartName** | Pointer to **string** | Name of the cart. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) |  | [optional] 

## Methods

### NewCartDetails

`func NewCartDetails() *CartDetails`

NewCartDetails instantiates a new CartDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartDetailsWithDefaults

`func NewCartDetailsWithDefaults() *CartDetails`

NewCartDetailsWithDefaults instantiates a new CartDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCharge

`func (o *CartDetails) GetShippingCharge() float64`

GetShippingCharge returns the ShippingCharge field if non-nil, zero value otherwise.

### GetShippingChargeOk

`func (o *CartDetails) GetShippingChargeOk() (*float64, bool)`

GetShippingChargeOk returns a tuple with the ShippingCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCharge

`func (o *CartDetails) SetShippingCharge(v float64)`

SetShippingCharge sets ShippingCharge field to given value.

### HasShippingCharge

`func (o *CartDetails) HasShippingCharge() bool`

HasShippingCharge returns a boolean if a field has been set.

### GetCartName

`func (o *CartDetails) GetCartName() string`

GetCartName returns the CartName field if non-nil, zero value otherwise.

### GetCartNameOk

`func (o *CartDetails) GetCartNameOk() (*string, bool)`

GetCartNameOk returns a tuple with the CartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartName

`func (o *CartDetails) SetCartName(v string)`

SetCartName sets CartName field to given value.

### HasCartName

`func (o *CartDetails) HasCartName() bool`

HasCartName returns a boolean if a field has been set.

### GetCartItems

`func (o *CartDetails) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CartDetails) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CartDetails) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *CartDetails) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


