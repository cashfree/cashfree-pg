# CFLinkOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfOrderId** | Pointer to **int32** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**OrderCurrency** | Pointer to **string** |  | [optional] 
**OrderAmount** | Pointer to **float32** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**OrderToken** | Pointer to **string** |  | [optional] 
**OrderExpiryTime** | Pointer to **string** |  | [optional] 
**OrderNote** | Pointer to **string** |  | [optional] 
**CustomerDetails** | Pointer to [**CFLinkCustomerDetailsEntity**](CFLinkCustomerDetailsEntity.md) |  | [optional] 
**Payments** | Pointer to [**CFPaymentURLObject**](CFPaymentURLObject.md) |  | [optional] 
**Settlements** | Pointer to [**CFSettlementURLObject**](CFSettlementURLObject.md) |  | [optional] 
**Refunds** | Pointer to [**CFRefundURLObject**](CFRefundURLObject.md) |  | [optional] 

## Methods

### NewCFLinkOrders

`func NewCFLinkOrders() *CFLinkOrders`

NewCFLinkOrders instantiates a new CFLinkOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkOrdersWithDefaults

`func NewCFLinkOrdersWithDefaults() *CFLinkOrders`

NewCFLinkOrdersWithDefaults instantiates a new CFLinkOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfOrderId

`func (o *CFLinkOrders) GetCfOrderId() int32`

GetCfOrderId returns the CfOrderId field if non-nil, zero value otherwise.

### GetCfOrderIdOk

`func (o *CFLinkOrders) GetCfOrderIdOk() (*int32, bool)`

GetCfOrderIdOk returns a tuple with the CfOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfOrderId

`func (o *CFLinkOrders) SetCfOrderId(v int32)`

SetCfOrderId sets CfOrderId field to given value.

### HasCfOrderId

`func (o *CFLinkOrders) HasCfOrderId() bool`

HasCfOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *CFLinkOrders) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFLinkOrders) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFLinkOrders) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFLinkOrders) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *CFLinkOrders) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CFLinkOrders) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CFLinkOrders) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CFLinkOrders) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *CFLinkOrders) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *CFLinkOrders) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *CFLinkOrders) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *CFLinkOrders) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CFLinkOrders) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CFLinkOrders) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CFLinkOrders) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *CFLinkOrders) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetOrderStatus

`func (o *CFLinkOrders) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *CFLinkOrders) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *CFLinkOrders) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *CFLinkOrders) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderToken

`func (o *CFLinkOrders) GetOrderToken() string`

GetOrderToken returns the OrderToken field if non-nil, zero value otherwise.

### GetOrderTokenOk

`func (o *CFLinkOrders) GetOrderTokenOk() (*string, bool)`

GetOrderTokenOk returns a tuple with the OrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderToken

`func (o *CFLinkOrders) SetOrderToken(v string)`

SetOrderToken sets OrderToken field to given value.

### HasOrderToken

`func (o *CFLinkOrders) HasOrderToken() bool`

HasOrderToken returns a boolean if a field has been set.

### GetOrderExpiryTime

`func (o *CFLinkOrders) GetOrderExpiryTime() string`

GetOrderExpiryTime returns the OrderExpiryTime field if non-nil, zero value otherwise.

### GetOrderExpiryTimeOk

`func (o *CFLinkOrders) GetOrderExpiryTimeOk() (*string, bool)`

GetOrderExpiryTimeOk returns a tuple with the OrderExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderExpiryTime

`func (o *CFLinkOrders) SetOrderExpiryTime(v string)`

SetOrderExpiryTime sets OrderExpiryTime field to given value.

### HasOrderExpiryTime

`func (o *CFLinkOrders) HasOrderExpiryTime() bool`

HasOrderExpiryTime returns a boolean if a field has been set.

### GetOrderNote

`func (o *CFLinkOrders) GetOrderNote() string`

GetOrderNote returns the OrderNote field if non-nil, zero value otherwise.

### GetOrderNoteOk

`func (o *CFLinkOrders) GetOrderNoteOk() (*string, bool)`

GetOrderNoteOk returns a tuple with the OrderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNote

`func (o *CFLinkOrders) SetOrderNote(v string)`

SetOrderNote sets OrderNote field to given value.

### HasOrderNote

`func (o *CFLinkOrders) HasOrderNote() bool`

HasOrderNote returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *CFLinkOrders) GetCustomerDetails() CFLinkCustomerDetailsEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CFLinkOrders) GetCustomerDetailsOk() (*CFLinkCustomerDetailsEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CFLinkOrders) SetCustomerDetails(v CFLinkCustomerDetailsEntity)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *CFLinkOrders) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetPayments

`func (o *CFLinkOrders) GetPayments() CFPaymentURLObject`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *CFLinkOrders) GetPaymentsOk() (*CFPaymentURLObject, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *CFLinkOrders) SetPayments(v CFPaymentURLObject)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *CFLinkOrders) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetSettlements

`func (o *CFLinkOrders) GetSettlements() CFSettlementURLObject`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *CFLinkOrders) GetSettlementsOk() (*CFSettlementURLObject, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *CFLinkOrders) SetSettlements(v CFSettlementURLObject)`

SetSettlements sets Settlements field to given value.

### HasSettlements

`func (o *CFLinkOrders) HasSettlements() bool`

HasSettlements returns a boolean if a field has been set.

### GetRefunds

`func (o *CFLinkOrders) GetRefunds() CFRefundURLObject`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *CFLinkOrders) GetRefundsOk() (*CFRefundURLObject, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *CFLinkOrders) SetRefunds(v CFRefundURLObject)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *CFLinkOrders) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


