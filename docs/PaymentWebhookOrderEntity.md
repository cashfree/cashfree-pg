# PaymentWebhookOrderEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**OrderAmount** | Pointer to **float64** |  | [optional] 
**OrderCurrency** | Pointer to **string** |  | [optional] 
**OrderTags** | Pointer to **map[string]string** | Custom Tags in thr form of {\&quot;key\&quot;:\&quot;value\&quot;} which can be passed for an order. A maximum of 10 tags can be added | [optional] 

## Methods

### NewPaymentWebhookOrderEntity

`func NewPaymentWebhookOrderEntity() *PaymentWebhookOrderEntity`

NewPaymentWebhookOrderEntity instantiates a new PaymentWebhookOrderEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWebhookOrderEntityWithDefaults

`func NewPaymentWebhookOrderEntityWithDefaults() *PaymentWebhookOrderEntity`

NewPaymentWebhookOrderEntityWithDefaults instantiates a new PaymentWebhookOrderEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *PaymentWebhookOrderEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentWebhookOrderEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentWebhookOrderEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentWebhookOrderEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *PaymentWebhookOrderEntity) GetOrderAmount() float64`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *PaymentWebhookOrderEntity) GetOrderAmountOk() (*float64, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *PaymentWebhookOrderEntity) SetOrderAmount(v float64)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *PaymentWebhookOrderEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *PaymentWebhookOrderEntity) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *PaymentWebhookOrderEntity) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *PaymentWebhookOrderEntity) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *PaymentWebhookOrderEntity) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetOrderTags

`func (o *PaymentWebhookOrderEntity) GetOrderTags() map[string]string`

GetOrderTags returns the OrderTags field if non-nil, zero value otherwise.

### GetOrderTagsOk

`func (o *PaymentWebhookOrderEntity) GetOrderTagsOk() (*map[string]string, bool)`

GetOrderTagsOk returns a tuple with the OrderTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTags

`func (o *PaymentWebhookOrderEntity) SetOrderTags(v map[string]string)`

SetOrderTags sets OrderTags field to given value.

### HasOrderTags

`func (o *PaymentWebhookOrderEntity) HasOrderTags() bool`

HasOrderTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


