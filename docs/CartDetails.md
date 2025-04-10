# CartDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerNote** | Pointer to **string** |  | [optional] 
**ShippingCharge** | Pointer to **float64** |  | [optional] 
**CartName** | Pointer to **string** | Name of the cart. | [optional] 
**CustomerShippingAddress** | Pointer to [**CartAddress**](CartAddress.md) |  | [optional] 
**CustomerBillingAddress** | Pointer to [**CartAddress**](CartAddress.md) |  | [optional] 
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

### GetCustomerNote

`func (o *CartDetails) GetCustomerNote() string`

GetCustomerNote returns the CustomerNote field if non-nil, zero value otherwise.

### GetCustomerNoteOk

`func (o *CartDetails) GetCustomerNoteOk() (*string, bool)`

GetCustomerNoteOk returns a tuple with the CustomerNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNote

`func (o *CartDetails) SetCustomerNote(v string)`

SetCustomerNote sets CustomerNote field to given value.

### HasCustomerNote

`func (o *CartDetails) HasCustomerNote() bool`

HasCustomerNote returns a boolean if a field has been set.

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

### GetCustomerShippingAddress

`func (o *CartDetails) GetCustomerShippingAddress() CartAddress`

GetCustomerShippingAddress returns the CustomerShippingAddress field if non-nil, zero value otherwise.

### GetCustomerShippingAddressOk

`func (o *CartDetails) GetCustomerShippingAddressOk() (*CartAddress, bool)`

GetCustomerShippingAddressOk returns a tuple with the CustomerShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerShippingAddress

`func (o *CartDetails) SetCustomerShippingAddress(v CartAddress)`

SetCustomerShippingAddress sets CustomerShippingAddress field to given value.

### HasCustomerShippingAddress

`func (o *CartDetails) HasCustomerShippingAddress() bool`

HasCustomerShippingAddress returns a boolean if a field has been set.

### GetCustomerBillingAddress

`func (o *CartDetails) GetCustomerBillingAddress() CartAddress`

GetCustomerBillingAddress returns the CustomerBillingAddress field if non-nil, zero value otherwise.

### GetCustomerBillingAddressOk

`func (o *CartDetails) GetCustomerBillingAddressOk() (*CartAddress, bool)`

GetCustomerBillingAddressOk returns a tuple with the CustomerBillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBillingAddress

`func (o *CartDetails) SetCustomerBillingAddress(v CartAddress)`

SetCustomerBillingAddress sets CustomerBillingAddress field to given value.

### HasCustomerBillingAddress

`func (o *CartDetails) HasCustomerBillingAddress() bool`

HasCustomerBillingAddress returns a boolean if a field has been set.

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


