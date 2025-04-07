# Products

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneClickCheckout** | Pointer to [**ProductDetails**](ProductDetails.md) |  | [optional] 
**VerifyPay** | Pointer to [**ProductDetails**](ProductDetails.md) |  | [optional] 

## Methods

### NewProducts

`func NewProducts() *Products`

NewProducts instantiates a new Products object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsWithDefaults

`func NewProductsWithDefaults() *Products`

NewProductsWithDefaults instantiates a new Products object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneClickCheckout

`func (o *Products) GetOneClickCheckout() ProductDetails`

GetOneClickCheckout returns the OneClickCheckout field if non-nil, zero value otherwise.

### GetOneClickCheckoutOk

`func (o *Products) GetOneClickCheckoutOk() (*ProductDetails, bool)`

GetOneClickCheckoutOk returns a tuple with the OneClickCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneClickCheckout

`func (o *Products) SetOneClickCheckout(v ProductDetails)`

SetOneClickCheckout sets OneClickCheckout field to given value.

### HasOneClickCheckout

`func (o *Products) HasOneClickCheckout() bool`

HasOneClickCheckout returns a boolean if a field has been set.

### GetVerifyPay

`func (o *Products) GetVerifyPay() ProductDetails`

GetVerifyPay returns the VerifyPay field if non-nil, zero value otherwise.

### GetVerifyPayOk

`func (o *Products) GetVerifyPayOk() (*ProductDetails, bool)`

GetVerifyPayOk returns a tuple with the VerifyPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPay

`func (o *Products) SetVerifyPay(v ProductDetails)`

SetVerifyPay sets VerifyPay field to given value.

### HasVerifyPay

`func (o *Products) HasVerifyPay() bool`

HasVerifyPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


