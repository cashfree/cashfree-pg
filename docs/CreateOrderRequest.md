# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **NullableString** | Order identifier present in your system. Alphanumeric, &#39;_&#39; and &#39;-&#39; only | [optional] 
**OrderAmount** | **float64** | Bill amount for the order. Provide upto two decimals. 10.15 means Rs 10 and 15 paisa | 
**OrderCurrency** | **string** | Currency for the order. INR if left empty. Contact care@cashfree.com to enable new currencies. | 
**CustomerDetails** | [**CustomerDetails**](CustomerDetails.md) |  | 
**Terminal** | Pointer to [**NullableCreateOrderRequestTerminal**](CreateOrderRequestTerminal.md) |  | [optional] 
**OrderMeta** | Pointer to [**NullableCreateOrderRequestOrderMeta**](CreateOrderRequestOrderMeta.md) |  | [optional] 
**OrderExpiryTime** | Pointer to **string** | Time after which the order expires. Customers will not be able to make the payment beyond the time specified here. We store timestamps in IST, but you can provide them in a valid ISO 8601 time format. Example 2021-07-02T10:20:12+05:30 for IST, 2021-07-02T10:20:12Z for UTC | [optional] 
**OrderNote** | Pointer to **NullableString** | Order note for reference. | [optional] 
**OrderTags** | Pointer to **map[string]string** | Custom Tags in thr form of {\&quot;key\&quot;:\&quot;value\&quot;} which can be passed for an order. A maximum of 10 tags can be added | [optional] 
**OrderSplits** | Pointer to [**[]VendorSplit**](VendorSplit.md) | If you have Easy split enabled in your Cashfree account then you can use this option to split the order amount. | [optional] 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(orderAmount float64, orderCurrency string, customerDetails CustomerDetails, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CreateOrderRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *CreateOrderRequest) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *CreateOrderRequest) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetOrderAmount

`func (o *CreateOrderRequest) GetOrderAmount() float64`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CreateOrderRequest) GetOrderAmountOk() (*float64, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CreateOrderRequest) SetOrderAmount(v float64)`

SetOrderAmount sets OrderAmount field to given value.


### GetOrderCurrency

`func (o *CreateOrderRequest) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *CreateOrderRequest) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *CreateOrderRequest) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.


### GetCustomerDetails

`func (o *CreateOrderRequest) GetCustomerDetails() CustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CreateOrderRequest) GetCustomerDetailsOk() (*CustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CreateOrderRequest) SetCustomerDetails(v CustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.


### GetTerminal

`func (o *CreateOrderRequest) GetTerminal() CreateOrderRequestTerminal`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *CreateOrderRequest) GetTerminalOk() (*CreateOrderRequestTerminal, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *CreateOrderRequest) SetTerminal(v CreateOrderRequestTerminal)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *CreateOrderRequest) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### SetTerminalNil

`func (o *CreateOrderRequest) SetTerminalNil(b bool)`

 SetTerminalNil sets the value for Terminal to be an explicit nil

### UnsetTerminal
`func (o *CreateOrderRequest) UnsetTerminal()`

UnsetTerminal ensures that no value is present for Terminal, not even an explicit nil
### GetOrderMeta

`func (o *CreateOrderRequest) GetOrderMeta() CreateOrderRequestOrderMeta`

GetOrderMeta returns the OrderMeta field if non-nil, zero value otherwise.

### GetOrderMetaOk

`func (o *CreateOrderRequest) GetOrderMetaOk() (*CreateOrderRequestOrderMeta, bool)`

GetOrderMetaOk returns a tuple with the OrderMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderMeta

`func (o *CreateOrderRequest) SetOrderMeta(v CreateOrderRequestOrderMeta)`

SetOrderMeta sets OrderMeta field to given value.

### HasOrderMeta

`func (o *CreateOrderRequest) HasOrderMeta() bool`

HasOrderMeta returns a boolean if a field has been set.

### SetOrderMetaNil

`func (o *CreateOrderRequest) SetOrderMetaNil(b bool)`

 SetOrderMetaNil sets the value for OrderMeta to be an explicit nil

### UnsetOrderMeta
`func (o *CreateOrderRequest) UnsetOrderMeta()`

UnsetOrderMeta ensures that no value is present for OrderMeta, not even an explicit nil
### GetOrderExpiryTime

`func (o *CreateOrderRequest) GetOrderExpiryTime() string`

GetOrderExpiryTime returns the OrderExpiryTime field if non-nil, zero value otherwise.

### GetOrderExpiryTimeOk

`func (o *CreateOrderRequest) GetOrderExpiryTimeOk() (*string, bool)`

GetOrderExpiryTimeOk returns a tuple with the OrderExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderExpiryTime

`func (o *CreateOrderRequest) SetOrderExpiryTime(v string)`

SetOrderExpiryTime sets OrderExpiryTime field to given value.

### HasOrderExpiryTime

`func (o *CreateOrderRequest) HasOrderExpiryTime() bool`

HasOrderExpiryTime returns a boolean if a field has been set.

### GetOrderNote

`func (o *CreateOrderRequest) GetOrderNote() string`

GetOrderNote returns the OrderNote field if non-nil, zero value otherwise.

### GetOrderNoteOk

`func (o *CreateOrderRequest) GetOrderNoteOk() (*string, bool)`

GetOrderNoteOk returns a tuple with the OrderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNote

`func (o *CreateOrderRequest) SetOrderNote(v string)`

SetOrderNote sets OrderNote field to given value.

### HasOrderNote

`func (o *CreateOrderRequest) HasOrderNote() bool`

HasOrderNote returns a boolean if a field has been set.

### SetOrderNoteNil

`func (o *CreateOrderRequest) SetOrderNoteNil(b bool)`

 SetOrderNoteNil sets the value for OrderNote to be an explicit nil

### UnsetOrderNote
`func (o *CreateOrderRequest) UnsetOrderNote()`

UnsetOrderNote ensures that no value is present for OrderNote, not even an explicit nil
### GetOrderTags

`func (o *CreateOrderRequest) GetOrderTags() map[string]string`

GetOrderTags returns the OrderTags field if non-nil, zero value otherwise.

### GetOrderTagsOk

`func (o *CreateOrderRequest) GetOrderTagsOk() (*map[string]string, bool)`

GetOrderTagsOk returns a tuple with the OrderTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTags

`func (o *CreateOrderRequest) SetOrderTags(v map[string]string)`

SetOrderTags sets OrderTags field to given value.

### HasOrderTags

`func (o *CreateOrderRequest) HasOrderTags() bool`

HasOrderTags returns a boolean if a field has been set.

### SetOrderTagsNil

`func (o *CreateOrderRequest) SetOrderTagsNil(b bool)`

 SetOrderTagsNil sets the value for OrderTags to be an explicit nil

### UnsetOrderTags
`func (o *CreateOrderRequest) UnsetOrderTags()`

UnsetOrderTags ensures that no value is present for OrderTags, not even an explicit nil
### GetOrderSplits

`func (o *CreateOrderRequest) GetOrderSplits() []VendorSplit`

GetOrderSplits returns the OrderSplits field if non-nil, zero value otherwise.

### GetOrderSplitsOk

`func (o *CreateOrderRequest) GetOrderSplitsOk() (*[]VendorSplit, bool)`

GetOrderSplitsOk returns a tuple with the OrderSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplits

`func (o *CreateOrderRequest) SetOrderSplits(v []VendorSplit)`

SetOrderSplits sets OrderSplits field to given value.

### HasOrderSplits

`func (o *CreateOrderRequest) HasOrderSplits() bool`

HasOrderSplits returns a boolean if a field has been set.

### SetOrderSplitsNil

`func (o *CreateOrderRequest) SetOrderSplitsNil(b bool)`

 SetOrderSplitsNil sets the value for OrderSplits to be an explicit nil

### UnsetOrderSplits
`func (o *CreateOrderRequest) UnsetOrderSplits()`

UnsetOrderSplits ensures that no value is present for OrderSplits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


