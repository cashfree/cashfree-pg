# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | Unique identifier of the item | [optional] 
**ItemName** | Pointer to **string** | Name of the item | [optional] 
**ItemDescription** | Pointer to **string** | Description of the item | [optional] 
**ItemTags** | Pointer to **[]string** | Tags attached to that item | [optional] 
**ItemDetailsUrl** | Pointer to **string** | Item details url | [optional] 
**ItemImageUrl** | Pointer to **string** | Item image url | [optional] 
**ItemOriginalUnitPrice** | Pointer to **float64** | Original price | [optional] 
**ItemDiscountedUnitPrice** | Pointer to **float64** | Discounted Price | [optional] 
**ItemCurrency** | Pointer to **string** | Currency of the item. | [optional] 
**ItemQuantity** | Pointer to **float32** | Quantity if that item | [optional] 

## Methods

### NewCartItem

`func NewCartItem() *CartItem`

NewCartItem instantiates a new CartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartItemWithDefaults

`func NewCartItemWithDefaults() *CartItem`

NewCartItemWithDefaults instantiates a new CartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *CartItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CartItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CartItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CartItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemName

`func (o *CartItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *CartItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *CartItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *CartItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetItemDescription

`func (o *CartItem) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *CartItem) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *CartItem) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *CartItem) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.

### GetItemTags

`func (o *CartItem) GetItemTags() []string`

GetItemTags returns the ItemTags field if non-nil, zero value otherwise.

### GetItemTagsOk

`func (o *CartItem) GetItemTagsOk() (*[]string, bool)`

GetItemTagsOk returns a tuple with the ItemTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTags

`func (o *CartItem) SetItemTags(v []string)`

SetItemTags sets ItemTags field to given value.

### HasItemTags

`func (o *CartItem) HasItemTags() bool`

HasItemTags returns a boolean if a field has been set.

### GetItemDetailsUrl

`func (o *CartItem) GetItemDetailsUrl() string`

GetItemDetailsUrl returns the ItemDetailsUrl field if non-nil, zero value otherwise.

### GetItemDetailsUrlOk

`func (o *CartItem) GetItemDetailsUrlOk() (*string, bool)`

GetItemDetailsUrlOk returns a tuple with the ItemDetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDetailsUrl

`func (o *CartItem) SetItemDetailsUrl(v string)`

SetItemDetailsUrl sets ItemDetailsUrl field to given value.

### HasItemDetailsUrl

`func (o *CartItem) HasItemDetailsUrl() bool`

HasItemDetailsUrl returns a boolean if a field has been set.

### GetItemImageUrl

`func (o *CartItem) GetItemImageUrl() string`

GetItemImageUrl returns the ItemImageUrl field if non-nil, zero value otherwise.

### GetItemImageUrlOk

`func (o *CartItem) GetItemImageUrlOk() (*string, bool)`

GetItemImageUrlOk returns a tuple with the ItemImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemImageUrl

`func (o *CartItem) SetItemImageUrl(v string)`

SetItemImageUrl sets ItemImageUrl field to given value.

### HasItemImageUrl

`func (o *CartItem) HasItemImageUrl() bool`

HasItemImageUrl returns a boolean if a field has been set.

### GetItemOriginalUnitPrice

`func (o *CartItem) GetItemOriginalUnitPrice() float64`

GetItemOriginalUnitPrice returns the ItemOriginalUnitPrice field if non-nil, zero value otherwise.

### GetItemOriginalUnitPriceOk

`func (o *CartItem) GetItemOriginalUnitPriceOk() (*float64, bool)`

GetItemOriginalUnitPriceOk returns a tuple with the ItemOriginalUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemOriginalUnitPrice

`func (o *CartItem) SetItemOriginalUnitPrice(v float64)`

SetItemOriginalUnitPrice sets ItemOriginalUnitPrice field to given value.

### HasItemOriginalUnitPrice

`func (o *CartItem) HasItemOriginalUnitPrice() bool`

HasItemOriginalUnitPrice returns a boolean if a field has been set.

### GetItemDiscountedUnitPrice

`func (o *CartItem) GetItemDiscountedUnitPrice() float64`

GetItemDiscountedUnitPrice returns the ItemDiscountedUnitPrice field if non-nil, zero value otherwise.

### GetItemDiscountedUnitPriceOk

`func (o *CartItem) GetItemDiscountedUnitPriceOk() (*float64, bool)`

GetItemDiscountedUnitPriceOk returns a tuple with the ItemDiscountedUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDiscountedUnitPrice

`func (o *CartItem) SetItemDiscountedUnitPrice(v float64)`

SetItemDiscountedUnitPrice sets ItemDiscountedUnitPrice field to given value.

### HasItemDiscountedUnitPrice

`func (o *CartItem) HasItemDiscountedUnitPrice() bool`

HasItemDiscountedUnitPrice returns a boolean if a field has been set.

### GetItemCurrency

`func (o *CartItem) GetItemCurrency() string`

GetItemCurrency returns the ItemCurrency field if non-nil, zero value otherwise.

### GetItemCurrencyOk

`func (o *CartItem) GetItemCurrencyOk() (*string, bool)`

GetItemCurrencyOk returns a tuple with the ItemCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCurrency

`func (o *CartItem) SetItemCurrency(v string)`

SetItemCurrency sets ItemCurrency field to given value.

### HasItemCurrency

`func (o *CartItem) HasItemCurrency() bool`

HasItemCurrency returns a boolean if a field has been set.

### GetItemQuantity

`func (o *CartItem) GetItemQuantity() float32`

GetItemQuantity returns the ItemQuantity field if non-nil, zero value otherwise.

### GetItemQuantityOk

`func (o *CartItem) GetItemQuantityOk() (*float32, bool)`

GetItemQuantityOk returns a tuple with the ItemQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemQuantity

`func (o *CartItem) SetItemQuantity(v float32)`

SetItemQuantity sets ItemQuantity field to given value.

### HasItemQuantity

`func (o *CartItem) HasItemQuantity() bool`

HasItemQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


