# SplitOrderReconSuccessResponseSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | Type of entity. Example: \&quot;settlement\&quot;. | [optional] 
**CfSettlementId** | Pointer to **int64** | Unique Cashfree settlement ID. | [optional] 
**CfPaymentId** | Pointer to **int64** | Unique Cashfree payment ID associated with the order. | [optional] 
**OrderId** | Pointer to **string** | Unique identifier for the order. | [optional] 
**OrderCurrency** | Pointer to **string** | Currency of the order. Example: \&quot;INR\&quot;. | [optional] 
**TransferId** | Pointer to **NullableString** | Unique transfer ID if available, otherwise null. | [optional] 
**OrderAmount** | Pointer to **float32** | Total amount of the order. | [optional] 
**ServiceCharge** | Pointer to **float32** | Service charge for the order. | [optional] 
**ServiceTax** | Pointer to **float32** | Service tax for the order. | [optional] 
**SettlementAmount** | Pointer to **float32** | Amount to be settled after charges and tax. | [optional] 
**SettlementCurrency** | Pointer to **string** | Currency of the settlement. Example: \&quot;INR\&quot;. | [optional] 
**TransferUtr** | Pointer to **NullableString** | UTR (Unique Transaction Reference) for the transfer if available, otherwise null. | [optional] 
**TransferTime** | Pointer to **NullableString** | Time of transfer if available, otherwise null. | [optional] 
**PaymentTime** | Pointer to **string** | Timestamp when payment was made. | [optional] 

## Methods

### NewSplitOrderReconSuccessResponseSettlement

`func NewSplitOrderReconSuccessResponseSettlement() *SplitOrderReconSuccessResponseSettlement`

NewSplitOrderReconSuccessResponseSettlement instantiates a new SplitOrderReconSuccessResponseSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitOrderReconSuccessResponseSettlementWithDefaults

`func NewSplitOrderReconSuccessResponseSettlementWithDefaults() *SplitOrderReconSuccessResponseSettlement`

NewSplitOrderReconSuccessResponseSettlementWithDefaults instantiates a new SplitOrderReconSuccessResponseSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *SplitOrderReconSuccessResponseSettlement) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SplitOrderReconSuccessResponseSettlement) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SplitOrderReconSuccessResponseSettlement) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetCfSettlementId

`func (o *SplitOrderReconSuccessResponseSettlement) GetCfSettlementId() int64`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetCfSettlementIdOk() (*int64, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *SplitOrderReconSuccessResponseSettlement) SetCfSettlementId(v int64)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *SplitOrderReconSuccessResponseSettlement) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *SplitOrderReconSuccessResponseSettlement) GetCfPaymentId() int64`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetCfPaymentIdOk() (*int64, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SplitOrderReconSuccessResponseSettlement) SetCfPaymentId(v int64)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SplitOrderReconSuccessResponseSettlement) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetOrderId

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SplitOrderReconSuccessResponseSettlement) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SplitOrderReconSuccessResponseSettlement) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetTransferId

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *SplitOrderReconSuccessResponseSettlement) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### SetTransferIdNil

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferIdNil(b bool)`

 SetTransferIdNil sets the value for TransferId to be an explicit nil

### UnsetTransferId
`func (o *SplitOrderReconSuccessResponseSettlement) UnsetTransferId()`

UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
### GetOrderAmount

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *SplitOrderReconSuccessResponseSettlement) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *SplitOrderReconSuccessResponseSettlement) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetServiceCharge

`func (o *SplitOrderReconSuccessResponseSettlement) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *SplitOrderReconSuccessResponseSettlement) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *SplitOrderReconSuccessResponseSettlement) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetServiceTax

`func (o *SplitOrderReconSuccessResponseSettlement) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *SplitOrderReconSuccessResponseSettlement) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *SplitOrderReconSuccessResponseSettlement) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *SplitOrderReconSuccessResponseSettlement) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *SplitOrderReconSuccessResponseSettlement) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *SplitOrderReconSuccessResponseSettlement) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetSettlementCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) GetSettlementCurrency() string`

GetSettlementCurrency returns the SettlementCurrency field if non-nil, zero value otherwise.

### GetSettlementCurrencyOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetSettlementCurrencyOk() (*string, bool)`

GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) SetSettlementCurrency(v string)`

SetSettlementCurrency sets SettlementCurrency field to given value.

### HasSettlementCurrency

`func (o *SplitOrderReconSuccessResponseSettlement) HasSettlementCurrency() bool`

HasSettlementCurrency returns a boolean if a field has been set.

### GetTransferUtr

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferUtr() string`

GetTransferUtr returns the TransferUtr field if non-nil, zero value otherwise.

### GetTransferUtrOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferUtrOk() (*string, bool)`

GetTransferUtrOk returns a tuple with the TransferUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferUtr

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferUtr(v string)`

SetTransferUtr sets TransferUtr field to given value.

### HasTransferUtr

`func (o *SplitOrderReconSuccessResponseSettlement) HasTransferUtr() bool`

HasTransferUtr returns a boolean if a field has been set.

### SetTransferUtrNil

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferUtrNil(b bool)`

 SetTransferUtrNil sets the value for TransferUtr to be an explicit nil

### UnsetTransferUtr
`func (o *SplitOrderReconSuccessResponseSettlement) UnsetTransferUtr()`

UnsetTransferUtr ensures that no value is present for TransferUtr, not even an explicit nil
### GetTransferTime

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferTime() string`

GetTransferTime returns the TransferTime field if non-nil, zero value otherwise.

### GetTransferTimeOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetTransferTimeOk() (*string, bool)`

GetTransferTimeOk returns a tuple with the TransferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTime

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferTime(v string)`

SetTransferTime sets TransferTime field to given value.

### HasTransferTime

`func (o *SplitOrderReconSuccessResponseSettlement) HasTransferTime() bool`

HasTransferTime returns a boolean if a field has been set.

### SetTransferTimeNil

`func (o *SplitOrderReconSuccessResponseSettlement) SetTransferTimeNil(b bool)`

 SetTransferTimeNil sets the value for TransferTime to be an explicit nil

### UnsetTransferTime
`func (o *SplitOrderReconSuccessResponseSettlement) UnsetTransferTime()`

UnsetTransferTime ensures that no value is present for TransferTime, not even an explicit nil
### GetPaymentTime

`func (o *SplitOrderReconSuccessResponseSettlement) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *SplitOrderReconSuccessResponseSettlement) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *SplitOrderReconSuccessResponseSettlement) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *SplitOrderReconSuccessResponseSettlement) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


