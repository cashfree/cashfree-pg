# OrderEntityProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneClickCheckout** | Pointer to [**ProductDetailsEntity**](ProductDetailsEntity.md) |  | [optional] 
**VerifyPay** | Pointer to [**ProductDetailsEntity**](ProductDetailsEntity.md) |  | [optional] 

## Methods

### NewOrderEntityProducts

`func NewOrderEntityProducts() *OrderEntityProducts`

NewOrderEntityProducts instantiates a new OrderEntityProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderEntityProductsWithDefaults

`func NewOrderEntityProductsWithDefaults() *OrderEntityProducts`

NewOrderEntityProductsWithDefaults instantiates a new OrderEntityProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneClickCheckout

`func (o *OrderEntityProducts) GetOneClickCheckout() ProductDetailsEntity`

GetOneClickCheckout returns the OneClickCheckout field if non-nil, zero value otherwise.

### GetOneClickCheckoutOk

`func (o *OrderEntityProducts) GetOneClickCheckoutOk() (*ProductDetailsEntity, bool)`

GetOneClickCheckoutOk returns a tuple with the OneClickCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneClickCheckout

`func (o *OrderEntityProducts) SetOneClickCheckout(v ProductDetailsEntity)`

SetOneClickCheckout sets OneClickCheckout field to given value.

### HasOneClickCheckout

`func (o *OrderEntityProducts) HasOneClickCheckout() bool`

HasOneClickCheckout returns a boolean if a field has been set.

### GetVerifyPay

`func (o *OrderEntityProducts) GetVerifyPay() ProductDetailsEntity`

GetVerifyPay returns the VerifyPay field if non-nil, zero value otherwise.

### GetVerifyPayOk

`func (o *OrderEntityProducts) GetVerifyPayOk() (*ProductDetailsEntity, bool)`

GetVerifyPayOk returns a tuple with the VerifyPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPay

`func (o *OrderEntityProducts) SetVerifyPay(v ProductDetailsEntity)`

SetVerifyPay sets VerifyPay field to given value.

### HasVerifyPay

`func (o *OrderEntityProducts) HasVerifyPay() bool`

HasVerifyPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


