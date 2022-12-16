# CFOrder

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
**PaymentLink** | Pointer to **string** |  | [optional] 
**CustomerDetails** | Pointer to [**CFCustomerDetails**](CFCustomerDetails.md) |  | [optional] 
**OrderMeta** | Pointer to [**CFOrderMeta**](CFOrderMeta.md) |  | [optional] 
**Payments** | Pointer to [**CFPaymentURLObject**](CFPaymentURLObject.md) |  | [optional] 
**Settlements** | Pointer to [**CFSettlementURLObject**](CFSettlementURLObject.md) |  | [optional] 
**Refunds** | Pointer to [**CFRefundURLObject**](CFRefundURLObject.md) |  | [optional] 

## Methods

### NewCFOrder

`func NewCFOrder() *CFOrder`

NewCFOrder instantiates a new CFOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFOrderWithDefaults

`func NewCFOrderWithDefaults() *CFOrder`

NewCFOrderWithDefaults instantiates a new CFOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfOrderId

`func (o *CFOrder) GetCfOrderId() int32`

GetCfOrderId returns the CfOrderId field if non-nil, zero value otherwise.

### GetCfOrderIdOk

`func (o *CFOrder) GetCfOrderIdOk() (*int32, bool)`

GetCfOrderIdOk returns a tuple with the CfOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfOrderId

`func (o *CFOrder) SetCfOrderId(v int32)`

SetCfOrderId sets CfOrderId field to given value.

### HasCfOrderId

`func (o *CFOrder) HasCfOrderId() bool`

HasCfOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *CFOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *CFOrder) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CFOrder) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CFOrder) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CFOrder) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *CFOrder) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *CFOrder) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *CFOrder) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *CFOrder) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CFOrder) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CFOrder) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CFOrder) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *CFOrder) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetOrderStatus

`func (o *CFOrder) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *CFOrder) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *CFOrder) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *CFOrder) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderToken

`func (o *CFOrder) GetOrderToken() string`

GetOrderToken returns the OrderToken field if non-nil, zero value otherwise.

### GetOrderTokenOk

`func (o *CFOrder) GetOrderTokenOk() (*string, bool)`

GetOrderTokenOk returns a tuple with the OrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderToken

`func (o *CFOrder) SetOrderToken(v string)`

SetOrderToken sets OrderToken field to given value.

### HasOrderToken

`func (o *CFOrder) HasOrderToken() bool`

HasOrderToken returns a boolean if a field has been set.

### GetOrderExpiryTime

`func (o *CFOrder) GetOrderExpiryTime() string`

GetOrderExpiryTime returns the OrderExpiryTime field if non-nil, zero value otherwise.

### GetOrderExpiryTimeOk

`func (o *CFOrder) GetOrderExpiryTimeOk() (*string, bool)`

GetOrderExpiryTimeOk returns a tuple with the OrderExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderExpiryTime

`func (o *CFOrder) SetOrderExpiryTime(v string)`

SetOrderExpiryTime sets OrderExpiryTime field to given value.

### HasOrderExpiryTime

`func (o *CFOrder) HasOrderExpiryTime() bool`

HasOrderExpiryTime returns a boolean if a field has been set.

### GetOrderNote

`func (o *CFOrder) GetOrderNote() string`

GetOrderNote returns the OrderNote field if non-nil, zero value otherwise.

### GetOrderNoteOk

`func (o *CFOrder) GetOrderNoteOk() (*string, bool)`

GetOrderNoteOk returns a tuple with the OrderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNote

`func (o *CFOrder) SetOrderNote(v string)`

SetOrderNote sets OrderNote field to given value.

### HasOrderNote

`func (o *CFOrder) HasOrderNote() bool`

HasOrderNote returns a boolean if a field has been set.

### GetPaymentLink

`func (o *CFOrder) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *CFOrder) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *CFOrder) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *CFOrder) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *CFOrder) GetCustomerDetails() CFCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CFOrder) GetCustomerDetailsOk() (*CFCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CFOrder) SetCustomerDetails(v CFCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *CFOrder) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetOrderMeta

`func (o *CFOrder) GetOrderMeta() CFOrderMeta`

GetOrderMeta returns the OrderMeta field if non-nil, zero value otherwise.

### GetOrderMetaOk

`func (o *CFOrder) GetOrderMetaOk() (*CFOrderMeta, bool)`

GetOrderMetaOk returns a tuple with the OrderMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderMeta

`func (o *CFOrder) SetOrderMeta(v CFOrderMeta)`

SetOrderMeta sets OrderMeta field to given value.

### HasOrderMeta

`func (o *CFOrder) HasOrderMeta() bool`

HasOrderMeta returns a boolean if a field has been set.

### GetPayments

`func (o *CFOrder) GetPayments() CFPaymentURLObject`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *CFOrder) GetPaymentsOk() (*CFPaymentURLObject, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *CFOrder) SetPayments(v CFPaymentURLObject)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *CFOrder) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetSettlements

`func (o *CFOrder) GetSettlements() CFSettlementURLObject`

GetSettlements returns the Settlements field if non-nil, zero value otherwise.

### GetSettlementsOk

`func (o *CFOrder) GetSettlementsOk() (*CFSettlementURLObject, bool)`

GetSettlementsOk returns a tuple with the Settlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlements

`func (o *CFOrder) SetSettlements(v CFSettlementURLObject)`

SetSettlements sets Settlements field to given value.

### HasSettlements

`func (o *CFOrder) HasSettlements() bool`

HasSettlements returns a boolean if a field has been set.

### GetRefunds

`func (o *CFOrder) GetRefunds() CFRefundURLObject`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *CFOrder) GetRefundsOk() (*CFRefundURLObject, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *CFOrder) SetRefunds(v CFRefundURLObject)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *CFOrder) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


