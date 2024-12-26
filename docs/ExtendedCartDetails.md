# ExtendedCartDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the cart. | [optional] 
**Items** | Pointer to [**[]CartItem**](CartItem.md) |  | [optional] 

## Methods

### NewExtendedCartDetails

`func NewExtendedCartDetails() *ExtendedCartDetails`

NewExtendedCartDetails instantiates a new ExtendedCartDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedCartDetailsWithDefaults

`func NewExtendedCartDetailsWithDefaults() *ExtendedCartDetails`

NewExtendedCartDetailsWithDefaults instantiates a new ExtendedCartDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtendedCartDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedCartDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedCartDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedCartDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *ExtendedCartDetails) GetItems() []CartItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExtendedCartDetails) GetItemsOk() (*[]CartItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExtendedCartDetails) SetItems(v []CartItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ExtendedCartDetails) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


