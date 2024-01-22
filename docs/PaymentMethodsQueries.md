# PaymentMethodsQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of the order. | [optional] 
**OrderId** | Pointer to **string** | OrderId of the order. Either of &#x60;order_id&#x60; or &#x60;order_amount&#x60; is mandatory. | [optional] 

## Methods

### NewPaymentMethodsQueries

`func NewPaymentMethodsQueries() *PaymentMethodsQueries`

NewPaymentMethodsQueries instantiates a new PaymentMethodsQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsQueriesWithDefaults

`func NewPaymentMethodsQueriesWithDefaults() *PaymentMethodsQueries`

NewPaymentMethodsQueriesWithDefaults instantiates a new PaymentMethodsQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentMethodsQueries) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentMethodsQueries) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentMethodsQueries) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentMethodsQueries) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOrderId

`func (o *PaymentMethodsQueries) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentMethodsQueries) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentMethodsQueries) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentMethodsQueries) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


