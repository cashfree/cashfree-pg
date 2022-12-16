# CFSettlementsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** |  | [optional] 
**CfSettlementId** | Pointer to **string** |  | [optional] 
**SettlementCurrency** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**OrderAmount** | Pointer to **float32** |  | [optional] 
**PaymentTime** | Pointer to **string** |  | [optional] 
**ServiceCharge** | Pointer to **float32** |  | [optional] 
**ServiceTax** | Pointer to **float32** |  | [optional] 
**SettlementAmount** | Pointer to **float32** |  | [optional] 
**SettlementId** | Pointer to **int32** |  | [optional] 
**TransferId** | Pointer to **int32** |  | [optional] 
**TransferTime** | Pointer to **string** |  | [optional] 
**TransferUtr** | Pointer to **string** |  | [optional] 

## Methods

### NewCFSettlementsEntity

`func NewCFSettlementsEntity() *CFSettlementsEntity`

NewCFSettlementsEntity instantiates a new CFSettlementsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFSettlementsEntityWithDefaults

`func NewCFSettlementsEntityWithDefaults() *CFSettlementsEntity`

NewCFSettlementsEntityWithDefaults instantiates a new CFSettlementsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CFSettlementsEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CFSettlementsEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CFSettlementsEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CFSettlementsEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSettlementId

`func (o *CFSettlementsEntity) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *CFSettlementsEntity) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *CFSettlementsEntity) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *CFSettlementsEntity) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementCurrency

`func (o *CFSettlementsEntity) GetSettlementCurrency() string`

GetSettlementCurrency returns the SettlementCurrency field if non-nil, zero value otherwise.

### GetSettlementCurrencyOk

`func (o *CFSettlementsEntity) GetSettlementCurrencyOk() (*string, bool)`

GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrency

`func (o *CFSettlementsEntity) SetSettlementCurrency(v string)`

SetSettlementCurrency sets SettlementCurrency field to given value.

### HasSettlementCurrency

`func (o *CFSettlementsEntity) HasSettlementCurrency() bool`

HasSettlementCurrency returns a boolean if a field has been set.

### GetOrderId

`func (o *CFSettlementsEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFSettlementsEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFSettlementsEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFSettlementsEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *CFSettlementsEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CFSettlementsEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CFSettlementsEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CFSettlementsEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CFSettlementsEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CFSettlementsEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CFSettlementsEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *CFSettlementsEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPaymentTime

`func (o *CFSettlementsEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *CFSettlementsEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *CFSettlementsEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *CFSettlementsEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetServiceCharge

`func (o *CFSettlementsEntity) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *CFSettlementsEntity) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *CFSettlementsEntity) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *CFSettlementsEntity) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetServiceTax

`func (o *CFSettlementsEntity) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *CFSettlementsEntity) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *CFSettlementsEntity) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *CFSettlementsEntity) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *CFSettlementsEntity) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *CFSettlementsEntity) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *CFSettlementsEntity) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *CFSettlementsEntity) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetSettlementId

`func (o *CFSettlementsEntity) GetSettlementId() int32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *CFSettlementsEntity) GetSettlementIdOk() (*int32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *CFSettlementsEntity) SetSettlementId(v int32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *CFSettlementsEntity) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetTransferId

`func (o *CFSettlementsEntity) GetTransferId() int32`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *CFSettlementsEntity) GetTransferIdOk() (*int32, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *CFSettlementsEntity) SetTransferId(v int32)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *CFSettlementsEntity) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetTransferTime

`func (o *CFSettlementsEntity) GetTransferTime() string`

GetTransferTime returns the TransferTime field if non-nil, zero value otherwise.

### GetTransferTimeOk

`func (o *CFSettlementsEntity) GetTransferTimeOk() (*string, bool)`

GetTransferTimeOk returns a tuple with the TransferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTime

`func (o *CFSettlementsEntity) SetTransferTime(v string)`

SetTransferTime sets TransferTime field to given value.

### HasTransferTime

`func (o *CFSettlementsEntity) HasTransferTime() bool`

HasTransferTime returns a boolean if a field has been set.

### GetTransferUtr

`func (o *CFSettlementsEntity) GetTransferUtr() string`

GetTransferUtr returns the TransferUtr field if non-nil, zero value otherwise.

### GetTransferUtrOk

`func (o *CFSettlementsEntity) GetTransferUtrOk() (*string, bool)`

GetTransferUtrOk returns a tuple with the TransferUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferUtr

`func (o *CFSettlementsEntity) SetTransferUtr(v string)`

SetTransferUtr sets TransferUtr field to given value.

### HasTransferUtr

`func (o *CFSettlementsEntity) HasTransferUtr() bool`

HasTransferUtr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


