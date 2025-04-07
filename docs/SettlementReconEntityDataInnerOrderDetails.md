# SettlementReconEntityDataInnerOrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Unique order ID. Alphanumeric and only &#39;-&#39; and &#39;_&#39; allowed. | [optional] 
**OrderAmount** | Pointer to **float32** | The amount which was passed at the order creation time. | [optional] 
**OrderCurrency** | Pointer to **string** | Order Curreny type - INR. | [optional] 
**OrderTags** | Pointer to **map[string]interface{}** | The order tags provided during order creation | [optional] 

## Methods

### NewSettlementReconEntityDataInnerOrderDetails

`func NewSettlementReconEntityDataInnerOrderDetails() *SettlementReconEntityDataInnerOrderDetails`

NewSettlementReconEntityDataInnerOrderDetails instantiates a new SettlementReconEntityDataInnerOrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerOrderDetailsWithDefaults

`func NewSettlementReconEntityDataInnerOrderDetailsWithDefaults() *SettlementReconEntityDataInnerOrderDetails`

NewSettlementReconEntityDataInnerOrderDetailsWithDefaults instantiates a new SettlementReconEntityDataInnerOrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SettlementReconEntityDataInnerOrderDetails) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SettlementReconEntityDataInnerOrderDetails) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *SettlementReconEntityDataInnerOrderDetails) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *SettlementReconEntityDataInnerOrderDetails) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *SettlementReconEntityDataInnerOrderDetails) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *SettlementReconEntityDataInnerOrderDetails) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetOrderTags

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderTags() map[string]interface{}`

GetOrderTags returns the OrderTags field if non-nil, zero value otherwise.

### GetOrderTagsOk

`func (o *SettlementReconEntityDataInnerOrderDetails) GetOrderTagsOk() (*map[string]interface{}, bool)`

GetOrderTagsOk returns a tuple with the OrderTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTags

`func (o *SettlementReconEntityDataInnerOrderDetails) SetOrderTags(v map[string]interface{})`

SetOrderTags sets OrderTags field to given value.

### HasOrderTags

`func (o *SettlementReconEntityDataInnerOrderDetails) HasOrderTags() bool`

HasOrderTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


