# OrderSettlementEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderDetails** | Pointer to [**OrderSettlementEntityOrderDetails**](OrderSettlementEntityOrderDetails.md) |  | [optional] 
**PaymentDetails** | Pointer to [**OrderSettlementEntityPaymentDetails**](OrderSettlementEntityPaymentDetails.md) |  | [optional] 
**CustomerDetails** | Pointer to [**OrderSettlementEntityCustomerDetails**](OrderSettlementEntityCustomerDetails.md) |  | [optional] 
**SettlementDetails** | Pointer to [**OrderSettlementEntitySettlementDetails**](OrderSettlementEntitySettlementDetails.md) |  | [optional] 

## Methods

### NewOrderSettlementEntity

`func NewOrderSettlementEntity() *OrderSettlementEntity`

NewOrderSettlementEntity instantiates a new OrderSettlementEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSettlementEntityWithDefaults

`func NewOrderSettlementEntityWithDefaults() *OrderSettlementEntity`

NewOrderSettlementEntityWithDefaults instantiates a new OrderSettlementEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderDetails

`func (o *OrderSettlementEntity) GetOrderDetails() OrderSettlementEntityOrderDetails`

GetOrderDetails returns the OrderDetails field if non-nil, zero value otherwise.

### GetOrderDetailsOk

`func (o *OrderSettlementEntity) GetOrderDetailsOk() (*OrderSettlementEntityOrderDetails, bool)`

GetOrderDetailsOk returns a tuple with the OrderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetails

`func (o *OrderSettlementEntity) SetOrderDetails(v OrderSettlementEntityOrderDetails)`

SetOrderDetails sets OrderDetails field to given value.

### HasOrderDetails

`func (o *OrderSettlementEntity) HasOrderDetails() bool`

HasOrderDetails returns a boolean if a field has been set.

### GetPaymentDetails

`func (o *OrderSettlementEntity) GetPaymentDetails() OrderSettlementEntityPaymentDetails`

GetPaymentDetails returns the PaymentDetails field if non-nil, zero value otherwise.

### GetPaymentDetailsOk

`func (o *OrderSettlementEntity) GetPaymentDetailsOk() (*OrderSettlementEntityPaymentDetails, bool)`

GetPaymentDetailsOk returns a tuple with the PaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetails

`func (o *OrderSettlementEntity) SetPaymentDetails(v OrderSettlementEntityPaymentDetails)`

SetPaymentDetails sets PaymentDetails field to given value.

### HasPaymentDetails

`func (o *OrderSettlementEntity) HasPaymentDetails() bool`

HasPaymentDetails returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *OrderSettlementEntity) GetCustomerDetails() OrderSettlementEntityCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *OrderSettlementEntity) GetCustomerDetailsOk() (*OrderSettlementEntityCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *OrderSettlementEntity) SetCustomerDetails(v OrderSettlementEntityCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *OrderSettlementEntity) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetSettlementDetails

`func (o *OrderSettlementEntity) GetSettlementDetails() OrderSettlementEntitySettlementDetails`

GetSettlementDetails returns the SettlementDetails field if non-nil, zero value otherwise.

### GetSettlementDetailsOk

`func (o *OrderSettlementEntity) GetSettlementDetailsOk() (*OrderSettlementEntitySettlementDetails, bool)`

GetSettlementDetailsOk returns a tuple with the SettlementDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDetails

`func (o *OrderSettlementEntity) SetSettlementDetails(v OrderSettlementEntitySettlementDetails)`

SetSettlementDetails sets SettlementDetails field to given value.

### HasSettlementDetails

`func (o *OrderSettlementEntity) HasSettlementDetails() bool`

HasSettlementDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


