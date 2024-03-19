# OfferValidations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAmount** | Pointer to **float32** | Minimum Amount for Offer to be Applicable | [optional] 
**MaxAllowed** | **float32** | Maximum Amount for Offer to be Applicable | 
**PaymentMethod** | [**OfferValidationsPaymentMethod**](OfferValidationsPaymentMethod.md) |  | 

## Methods

### NewOfferValidations

`func NewOfferValidations(maxAllowed float32, paymentMethod OfferValidationsPaymentMethod, ) *OfferValidations`

NewOfferValidations instantiates a new OfferValidations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferValidationsWithDefaults

`func NewOfferValidationsWithDefaults() *OfferValidations`

NewOfferValidationsWithDefaults instantiates a new OfferValidations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAmount

`func (o *OfferValidations) GetMinAmount() float32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *OfferValidations) GetMinAmountOk() (*float32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *OfferValidations) SetMinAmount(v float32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *OfferValidations) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMaxAllowed

`func (o *OfferValidations) GetMaxAllowed() float32`

GetMaxAllowed returns the MaxAllowed field if non-nil, zero value otherwise.

### GetMaxAllowedOk

`func (o *OfferValidations) GetMaxAllowedOk() (*float32, bool)`

GetMaxAllowedOk returns a tuple with the MaxAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowed

`func (o *OfferValidations) SetMaxAllowed(v float32)`

SetMaxAllowed sets MaxAllowed field to given value.


### GetPaymentMethod

`func (o *OfferValidations) GetPaymentMethod() OfferValidationsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OfferValidations) GetPaymentMethodOk() (*OfferValidationsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OfferValidations) SetPaymentMethod(v OfferValidationsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


