# OrderSettlementEntityOrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Unique order ID. | [optional] 
**OrderAmount** | Pointer to **float32** | The amount which was passed at the order creation time. | [optional] 
**OrderCurrency** | Pointer to **string** | Currency of the order. | [optional] 
**OrderTags** | Pointer to **map[string]interface{}** | The order tags provided during order creation. | [optional] 

## Methods

### NewOrderSettlementEntityOrderDetails

`func NewOrderSettlementEntityOrderDetails() *OrderSettlementEntityOrderDetails`

NewOrderSettlementEntityOrderDetails instantiates a new OrderSettlementEntityOrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSettlementEntityOrderDetailsWithDefaults

`func NewOrderSettlementEntityOrderDetailsWithDefaults() *OrderSettlementEntityOrderDetails`

NewOrderSettlementEntityOrderDetailsWithDefaults instantiates a new OrderSettlementEntityOrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderSettlementEntityOrderDetails) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderSettlementEntityOrderDetails) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderSettlementEntityOrderDetails) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderSettlementEntityOrderDetails) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *OrderSettlementEntityOrderDetails) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *OrderSettlementEntityOrderDetails) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *OrderSettlementEntityOrderDetails) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *OrderSettlementEntityOrderDetails) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *OrderSettlementEntityOrderDetails) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *OrderSettlementEntityOrderDetails) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *OrderSettlementEntityOrderDetails) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *OrderSettlementEntityOrderDetails) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetOrderTags

`func (o *OrderSettlementEntityOrderDetails) GetOrderTags() map[string]interface{}`

GetOrderTags returns the OrderTags field if non-nil, zero value otherwise.

### GetOrderTagsOk

`func (o *OrderSettlementEntityOrderDetails) GetOrderTagsOk() (*map[string]interface{}, bool)`

GetOrderTagsOk returns a tuple with the OrderTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTags

`func (o *OrderSettlementEntityOrderDetails) SetOrderTags(v map[string]interface{})`

SetOrderTags sets OrderTags field to given value.

### HasOrderTags

`func (o *OrderSettlementEntityOrderDetails) HasOrderTags() bool`

HasOrderTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


