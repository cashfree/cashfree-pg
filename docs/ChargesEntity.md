# ChargesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCharges** | Pointer to **float32** | Shipping charge of the order | [optional] 
**CodHandlingCharges** | Pointer to **float32** | COD handling fee for order | [optional] 

## Methods

### NewChargesEntity

`func NewChargesEntity() *ChargesEntity`

NewChargesEntity instantiates a new ChargesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargesEntityWithDefaults

`func NewChargesEntityWithDefaults() *ChargesEntity`

NewChargesEntityWithDefaults instantiates a new ChargesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCharges

`func (o *ChargesEntity) GetShippingCharges() float32`

GetShippingCharges returns the ShippingCharges field if non-nil, zero value otherwise.

### GetShippingChargesOk

`func (o *ChargesEntity) GetShippingChargesOk() (*float32, bool)`

GetShippingChargesOk returns a tuple with the ShippingCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCharges

`func (o *ChargesEntity) SetShippingCharges(v float32)`

SetShippingCharges sets ShippingCharges field to given value.

### HasShippingCharges

`func (o *ChargesEntity) HasShippingCharges() bool`

HasShippingCharges returns a boolean if a field has been set.

### GetCodHandlingCharges

`func (o *ChargesEntity) GetCodHandlingCharges() float32`

GetCodHandlingCharges returns the CodHandlingCharges field if non-nil, zero value otherwise.

### GetCodHandlingChargesOk

`func (o *ChargesEntity) GetCodHandlingChargesOk() (*float32, bool)`

GetCodHandlingChargesOk returns a tuple with the CodHandlingCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodHandlingCharges

`func (o *ChargesEntity) SetCodHandlingCharges(v float32)`

SetCodHandlingCharges sets CodHandlingCharges field to given value.

### HasCodHandlingCharges

`func (o *ChargesEntity) HasCodHandlingCharges() bool`

HasCodHandlingCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


