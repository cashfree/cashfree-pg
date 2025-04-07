# SettlementReconEntityDataInnerPaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAmount** | Pointer to **float32** | Payment amount captured. | [optional] 
**PaymentCurrency** | Pointer to **string** | Payment Curreny type - INR. | [optional] 
**BankReference** | Pointer to **string** | Unique transaction reference number of the payment. | [optional] 
**PaymentTime** | Pointer to **string** | Date and time when the payment was initiated. | [optional] 
**PaymentMode** | Pointer to **string** | Mode of the payment. | [optional] 
**PaymentServiceCharge** | Pointer to **float32** | Service charge applicable for the payment. | [optional] 
**PaymentServiceTax** | Pointer to **float32** | Service tax applicable on the payment. | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree Payments unique ID to identify a payment. | [optional] 
**Status** | Pointer to **string** | Status of the Payment. | [optional] 
**ForexConversionHandlingCharge** | Pointer to **string** | Forex Conversion Service Charge. | [optional] 
**ForexConversionHandlingTax** | Pointer to **string** | Forex Conversion Service Tax. | [optional] 
**ChargesCurrency** | Pointer to **string** | Forex Charges Curreny type - INR. | [optional] 

## Methods

### NewSettlementReconEntityDataInnerPaymentDetails

`func NewSettlementReconEntityDataInnerPaymentDetails() *SettlementReconEntityDataInnerPaymentDetails`

NewSettlementReconEntityDataInnerPaymentDetails instantiates a new SettlementReconEntityDataInnerPaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerPaymentDetailsWithDefaults

`func NewSettlementReconEntityDataInnerPaymentDetailsWithDefaults() *SettlementReconEntityDataInnerPaymentDetails`

NewSettlementReconEntityDataInnerPaymentDetailsWithDefaults instantiates a new SettlementReconEntityDataInnerPaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAmount

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetBankReference

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetBankReference() string`

GetBankReference returns the BankReference field if non-nil, zero value otherwise.

### GetBankReferenceOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetBankReferenceOk() (*string, bool)`

GetBankReferenceOk returns a tuple with the BankReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankReference

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetBankReference(v string)`

SetBankReference sets BankReference field to given value.

### HasBankReference

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasBankReference() bool`

HasBankReference returns a boolean if a field has been set.

### GetPaymentTime

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentMode

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentMode() string`

GetPaymentMode returns the PaymentMode field if non-nil, zero value otherwise.

### GetPaymentModeOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentModeOk() (*string, bool)`

GetPaymentModeOk returns a tuple with the PaymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMode

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentMode(v string)`

SetPaymentMode sets PaymentMode field to given value.

### HasPaymentMode

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentMode() bool`

HasPaymentMode returns a boolean if a field has been set.

### GetPaymentServiceCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentServiceCharge() float32`

GetPaymentServiceCharge returns the PaymentServiceCharge field if non-nil, zero value otherwise.

### GetPaymentServiceChargeOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentServiceChargeOk() (*float32, bool)`

GetPaymentServiceChargeOk returns a tuple with the PaymentServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentServiceCharge(v float32)`

SetPaymentServiceCharge sets PaymentServiceCharge field to given value.

### HasPaymentServiceCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentServiceCharge() bool`

HasPaymentServiceCharge returns a boolean if a field has been set.

### GetPaymentServiceTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentServiceTax() float32`

GetPaymentServiceTax returns the PaymentServiceTax field if non-nil, zero value otherwise.

### GetPaymentServiceTaxOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetPaymentServiceTaxOk() (*float32, bool)`

GetPaymentServiceTaxOk returns a tuple with the PaymentServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetPaymentServiceTax(v float32)`

SetPaymentServiceTax sets PaymentServiceTax field to given value.

### HasPaymentServiceTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasPaymentServiceTax() bool`

HasPaymentServiceTax returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetForexConversionHandlingCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetForexConversionHandlingCharge() string`

GetForexConversionHandlingCharge returns the ForexConversionHandlingCharge field if non-nil, zero value otherwise.

### GetForexConversionHandlingChargeOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetForexConversionHandlingChargeOk() (*string, bool)`

GetForexConversionHandlingChargeOk returns a tuple with the ForexConversionHandlingCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetForexConversionHandlingCharge(v string)`

SetForexConversionHandlingCharge sets ForexConversionHandlingCharge field to given value.

### HasForexConversionHandlingCharge

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasForexConversionHandlingCharge() bool`

HasForexConversionHandlingCharge returns a boolean if a field has been set.

### GetForexConversionHandlingTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetForexConversionHandlingTax() string`

GetForexConversionHandlingTax returns the ForexConversionHandlingTax field if non-nil, zero value otherwise.

### GetForexConversionHandlingTaxOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetForexConversionHandlingTaxOk() (*string, bool)`

GetForexConversionHandlingTaxOk returns a tuple with the ForexConversionHandlingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetForexConversionHandlingTax(v string)`

SetForexConversionHandlingTax sets ForexConversionHandlingTax field to given value.

### HasForexConversionHandlingTax

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasForexConversionHandlingTax() bool`

HasForexConversionHandlingTax returns a boolean if a field has been set.

### GetChargesCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetChargesCurrency() string`

GetChargesCurrency returns the ChargesCurrency field if non-nil, zero value otherwise.

### GetChargesCurrencyOk

`func (o *SettlementReconEntityDataInnerPaymentDetails) GetChargesCurrencyOk() (*string, bool)`

GetChargesCurrencyOk returns a tuple with the ChargesCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) SetChargesCurrency(v string)`

SetChargesCurrency sets ChargesCurrency field to given value.

### HasChargesCurrency

`func (o *SettlementReconEntityDataInnerPaymentDetails) HasChargesCurrency() bool`

HasChargesCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


