# CFOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Order identifier present in your system. Alphanumeric and only - and _ allowed. | [optional] 
**OrderAmount** | **float64** | Bill amount for the order. Provide upto two decimals. 10.15 means Rs 10 and 15 paisa | 
**OrderCurrency** | **string** | Currency for the order. INR if left empty. Contact care@cashfree.com to enable new currencies.  | 
**CustomerDetails** | [**CFCustomerDetails**](CFCustomerDetails.md) |  | 
**OrderMeta** | Pointer to [**CFOrderMeta**](CFOrderMeta.md) |  | [optional] 
**OrderExpiryTime** | Pointer to **string** | Time after which the order expires. Customers will not be able to make the payment beyond the time specified here. We store timestamps in IST, but you can provide them in a valid ISO 8601 time format. | [optional] 
**OrderNote** | Pointer to **string** | Order note for reference. | [optional] 
**OrderTags** | Pointer to **map[string]string** | Custom Tags which can be passed for an order. A maximum of 6 tags can be added | [optional] 
**OrderSplits** | Pointer to [**[]CFVendorSplit**](CFVendorSplit.md) |  | [optional] 

## Methods

### NewCFOrderRequest

`func NewCFOrderRequest(orderAmount float64, orderCurrency string, customerDetails CFCustomerDetails, ) *CFOrderRequest`

NewCFOrderRequest instantiates a new CFOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFOrderRequestWithDefaults

`func NewCFOrderRequestWithDefaults() *CFOrderRequest`

NewCFOrderRequestWithDefaults instantiates a new CFOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CFOrderRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFOrderRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFOrderRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFOrderRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CFOrderRequest) GetOrderAmount() float64`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CFOrderRequest) GetOrderAmountOk() (*float64, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CFOrderRequest) SetOrderAmount(v float64)`

SetOrderAmount sets OrderAmount field to given value.


### GetOrderCurrency

`func (o *CFOrderRequest) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *CFOrderRequest) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *CFOrderRequest) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.


### GetCustomerDetails

`func (o *CFOrderRequest) GetCustomerDetails() CFCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CFOrderRequest) GetCustomerDetailsOk() (*CFCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CFOrderRequest) SetCustomerDetails(v CFCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.


### GetOrderMeta

`func (o *CFOrderRequest) GetOrderMeta() CFOrderMeta`

GetOrderMeta returns the OrderMeta field if non-nil, zero value otherwise.

### GetOrderMetaOk

`func (o *CFOrderRequest) GetOrderMetaOk() (*CFOrderMeta, bool)`

GetOrderMetaOk returns a tuple with the OrderMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderMeta

`func (o *CFOrderRequest) SetOrderMeta(v CFOrderMeta)`

SetOrderMeta sets OrderMeta field to given value.

### HasOrderMeta

`func (o *CFOrderRequest) HasOrderMeta() bool`

HasOrderMeta returns a boolean if a field has been set.

### GetOrderExpiryTime

`func (o *CFOrderRequest) GetOrderExpiryTime() string`

GetOrderExpiryTime returns the OrderExpiryTime field if non-nil, zero value otherwise.

### GetOrderExpiryTimeOk

`func (o *CFOrderRequest) GetOrderExpiryTimeOk() (*string, bool)`

GetOrderExpiryTimeOk returns a tuple with the OrderExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderExpiryTime

`func (o *CFOrderRequest) SetOrderExpiryTime(v string)`

SetOrderExpiryTime sets OrderExpiryTime field to given value.

### HasOrderExpiryTime

`func (o *CFOrderRequest) HasOrderExpiryTime() bool`

HasOrderExpiryTime returns a boolean if a field has been set.

### GetOrderNote

`func (o *CFOrderRequest) GetOrderNote() string`

GetOrderNote returns the OrderNote field if non-nil, zero value otherwise.

### GetOrderNoteOk

`func (o *CFOrderRequest) GetOrderNoteOk() (*string, bool)`

GetOrderNoteOk returns a tuple with the OrderNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNote

`func (o *CFOrderRequest) SetOrderNote(v string)`

SetOrderNote sets OrderNote field to given value.

### HasOrderNote

`func (o *CFOrderRequest) HasOrderNote() bool`

HasOrderNote returns a boolean if a field has been set.

### GetOrderTags

`func (o *CFOrderRequest) GetOrderTags() map[string]string`

GetOrderTags returns the OrderTags field if non-nil, zero value otherwise.

### GetOrderTagsOk

`func (o *CFOrderRequest) GetOrderTagsOk() (*map[string]string, bool)`

GetOrderTagsOk returns a tuple with the OrderTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTags

`func (o *CFOrderRequest) SetOrderTags(v map[string]string)`

SetOrderTags sets OrderTags field to given value.

### HasOrderTags

`func (o *CFOrderRequest) HasOrderTags() bool`

HasOrderTags returns a boolean if a field has been set.

### GetOrderSplits

`func (o *CFOrderRequest) GetOrderSplits() []CFVendorSplit`

GetOrderSplits returns the OrderSplits field if non-nil, zero value otherwise.

### GetOrderSplitsOk

`func (o *CFOrderRequest) GetOrderSplitsOk() (*[]CFVendorSplit, bool)`

GetOrderSplitsOk returns a tuple with the OrderSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplits

`func (o *CFOrderRequest) SetOrderSplits(v []CFVendorSplit)`

SetOrderSplits sets OrderSplits field to given value.

### HasOrderSplits

`func (o *CFOrderRequest) HasOrderSplits() bool`

HasOrderSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


