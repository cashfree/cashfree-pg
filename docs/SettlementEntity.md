# SettlementEntity

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

### NewSettlementEntity

`func NewSettlementEntity() *SettlementEntity`

NewSettlementEntity instantiates a new SettlementEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementEntityWithDefaults

`func NewSettlementEntityWithDefaults() *SettlementEntity`

NewSettlementEntityWithDefaults instantiates a new SettlementEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *SettlementEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SettlementEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SettlementEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SettlementEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSettlementId

`func (o *SettlementEntity) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *SettlementEntity) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *SettlementEntity) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *SettlementEntity) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementCurrency

`func (o *SettlementEntity) GetSettlementCurrency() string`

GetSettlementCurrency returns the SettlementCurrency field if non-nil, zero value otherwise.

### GetSettlementCurrencyOk

`func (o *SettlementEntity) GetSettlementCurrencyOk() (*string, bool)`

GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrency

`func (o *SettlementEntity) SetSettlementCurrency(v string)`

SetSettlementCurrency sets SettlementCurrency field to given value.

### HasSettlementCurrency

`func (o *SettlementEntity) HasSettlementCurrency() bool`

HasSettlementCurrency returns a boolean if a field has been set.

### GetOrderId

`func (o *SettlementEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SettlementEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SettlementEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SettlementEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *SettlementEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SettlementEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SettlementEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SettlementEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetOrderAmount

`func (o *SettlementEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *SettlementEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *SettlementEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *SettlementEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPaymentTime

`func (o *SettlementEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *SettlementEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *SettlementEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *SettlementEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetServiceCharge

`func (o *SettlementEntity) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *SettlementEntity) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *SettlementEntity) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *SettlementEntity) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetServiceTax

`func (o *SettlementEntity) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *SettlementEntity) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *SettlementEntity) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *SettlementEntity) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *SettlementEntity) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *SettlementEntity) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *SettlementEntity) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *SettlementEntity) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetSettlementId

`func (o *SettlementEntity) GetSettlementId() int32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *SettlementEntity) GetSettlementIdOk() (*int32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *SettlementEntity) SetSettlementId(v int32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *SettlementEntity) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetTransferId

`func (o *SettlementEntity) GetTransferId() int32`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *SettlementEntity) GetTransferIdOk() (*int32, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *SettlementEntity) SetTransferId(v int32)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *SettlementEntity) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetTransferTime

`func (o *SettlementEntity) GetTransferTime() string`

GetTransferTime returns the TransferTime field if non-nil, zero value otherwise.

### GetTransferTimeOk

`func (o *SettlementEntity) GetTransferTimeOk() (*string, bool)`

GetTransferTimeOk returns a tuple with the TransferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTime

`func (o *SettlementEntity) SetTransferTime(v string)`

SetTransferTime sets TransferTime field to given value.

### HasTransferTime

`func (o *SettlementEntity) HasTransferTime() bool`

HasTransferTime returns a boolean if a field has been set.

### GetTransferUtr

`func (o *SettlementEntity) GetTransferUtr() string`

GetTransferUtr returns the TransferUtr field if non-nil, zero value otherwise.

### GetTransferUtrOk

`func (o *SettlementEntity) GetTransferUtrOk() (*string, bool)`

GetTransferUtrOk returns a tuple with the TransferUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferUtr

`func (o *SettlementEntity) SetTransferUtr(v string)`

SetTransferUtr sets TransferUtr field to given value.

### HasTransferUtr

`func (o *SettlementEntity) HasTransferUtr() bool`

HasTransferUtr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


