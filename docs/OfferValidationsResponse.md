# OfferValidationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAllowed** | Pointer to **float32** | Maximum Amount for Offer to be Applicable | [optional] 
**MinAmount** | Pointer to **float32** | Minimum Amount for Offer to be Applicable | [optional] 
**PaymentMethod** | Pointer to [**OfferValidationsResponsePaymentMethod**](OfferValidationsResponsePaymentMethod.md) |  | [optional] 

## Methods

### NewOfferValidationsResponse

`func NewOfferValidationsResponse() *OfferValidationsResponse`

NewOfferValidationsResponse instantiates a new OfferValidationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferValidationsResponseWithDefaults

`func NewOfferValidationsResponseWithDefaults() *OfferValidationsResponse`

NewOfferValidationsResponseWithDefaults instantiates a new OfferValidationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAllowed

`func (o *OfferValidationsResponse) GetMaxAllowed() float32`

GetMaxAllowed returns the MaxAllowed field if non-nil, zero value otherwise.

### GetMaxAllowedOk

`func (o *OfferValidationsResponse) GetMaxAllowedOk() (*float32, bool)`

GetMaxAllowedOk returns a tuple with the MaxAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowed

`func (o *OfferValidationsResponse) SetMaxAllowed(v float32)`

SetMaxAllowed sets MaxAllowed field to given value.

### HasMaxAllowed

`func (o *OfferValidationsResponse) HasMaxAllowed() bool`

HasMaxAllowed returns a boolean if a field has been set.

### GetMinAmount

`func (o *OfferValidationsResponse) GetMinAmount() float32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *OfferValidationsResponse) GetMinAmountOk() (*float32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *OfferValidationsResponse) SetMinAmount(v float32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *OfferValidationsResponse) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OfferValidationsResponse) GetPaymentMethod() OfferValidationsResponsePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OfferValidationsResponse) GetPaymentMethodOk() (*OfferValidationsResponsePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OfferValidationsResponse) SetPaymentMethod(v OfferValidationsResponsePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OfferValidationsResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


