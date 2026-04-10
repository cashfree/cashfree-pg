# ReconEntityDataInnerPaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAmount** | Pointer to **float32** | Payment amount captured. | [optional] 
**PaymentCurrency** | Pointer to **string** | Payment Curreny type - INR. | [optional] 
**BankReference** | Pointer to **string** | Unique transaction reference number of the payment. | [optional] 
**PaymentTime** | Pointer to **string** | Date and time when the payment was initiated. | [optional] 
**PaymentGroup** | Pointer to **string** | Mode of the payment. | [optional] 
**PaymentServiceCharge** | Pointer to **float32** | Service charge applicable for the payment. | [optional] 
**PaymentServiceTax** | Pointer to **float32** | Service tax applicable on the payment. | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree Payments unique ID to identify a payment. | [optional] 
**Status** | Pointer to **string** | Status of the Payment. | [optional] 
**ForexConversionHandlingCharge** | Pointer to **string** | Forex Conversion Service Charge. | [optional] 
**ForexConversionHandlingTax** | Pointer to **string** | Forex Conversion Service Tax. | [optional] 
**ChargesCurrency** | Pointer to **string** | Forex Charges Curreny type - INR. | [optional] 

## Methods

### NewReconEntityDataInnerPaymentDetails

`func NewReconEntityDataInnerPaymentDetails() *ReconEntityDataInnerPaymentDetails`

NewReconEntityDataInnerPaymentDetails instantiates a new ReconEntityDataInnerPaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityDataInnerPaymentDetailsWithDefaults

`func NewReconEntityDataInnerPaymentDetailsWithDefaults() *ReconEntityDataInnerPaymentDetails`

NewReconEntityDataInnerPaymentDetailsWithDefaults instantiates a new ReconEntityDataInnerPaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAmount

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetBankReference

`func (o *ReconEntityDataInnerPaymentDetails) GetBankReference() string`

GetBankReference returns the BankReference field if non-nil, zero value otherwise.

### GetBankReferenceOk

`func (o *ReconEntityDataInnerPaymentDetails) GetBankReferenceOk() (*string, bool)`

GetBankReferenceOk returns a tuple with the BankReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankReference

`func (o *ReconEntityDataInnerPaymentDetails) SetBankReference(v string)`

SetBankReference sets BankReference field to given value.

### HasBankReference

`func (o *ReconEntityDataInnerPaymentDetails) HasBankReference() bool`

HasBankReference returns a boolean if a field has been set.

### GetPaymentTime

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentGroup

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentGroup() string`

GetPaymentGroup returns the PaymentGroup field if non-nil, zero value otherwise.

### GetPaymentGroupOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentGroupOk() (*string, bool)`

GetPaymentGroupOk returns a tuple with the PaymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroup

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentGroup(v string)`

SetPaymentGroup sets PaymentGroup field to given value.

### HasPaymentGroup

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentGroup() bool`

HasPaymentGroup returns a boolean if a field has been set.

### GetPaymentServiceCharge

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentServiceCharge() float32`

GetPaymentServiceCharge returns the PaymentServiceCharge field if non-nil, zero value otherwise.

### GetPaymentServiceChargeOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentServiceChargeOk() (*float32, bool)`

GetPaymentServiceChargeOk returns a tuple with the PaymentServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceCharge

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentServiceCharge(v float32)`

SetPaymentServiceCharge sets PaymentServiceCharge field to given value.

### HasPaymentServiceCharge

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentServiceCharge() bool`

HasPaymentServiceCharge returns a boolean if a field has been set.

### GetPaymentServiceTax

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentServiceTax() float32`

GetPaymentServiceTax returns the PaymentServiceTax field if non-nil, zero value otherwise.

### GetPaymentServiceTaxOk

`func (o *ReconEntityDataInnerPaymentDetails) GetPaymentServiceTaxOk() (*float32, bool)`

GetPaymentServiceTaxOk returns a tuple with the PaymentServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTax

`func (o *ReconEntityDataInnerPaymentDetails) SetPaymentServiceTax(v float32)`

SetPaymentServiceTax sets PaymentServiceTax field to given value.

### HasPaymentServiceTax

`func (o *ReconEntityDataInnerPaymentDetails) HasPaymentServiceTax() bool`

HasPaymentServiceTax returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *ReconEntityDataInnerPaymentDetails) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *ReconEntityDataInnerPaymentDetails) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *ReconEntityDataInnerPaymentDetails) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *ReconEntityDataInnerPaymentDetails) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *ReconEntityDataInnerPaymentDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReconEntityDataInnerPaymentDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReconEntityDataInnerPaymentDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReconEntityDataInnerPaymentDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetForexConversionHandlingCharge

`func (o *ReconEntityDataInnerPaymentDetails) GetForexConversionHandlingCharge() string`

GetForexConversionHandlingCharge returns the ForexConversionHandlingCharge field if non-nil, zero value otherwise.

### GetForexConversionHandlingChargeOk

`func (o *ReconEntityDataInnerPaymentDetails) GetForexConversionHandlingChargeOk() (*string, bool)`

GetForexConversionHandlingChargeOk returns a tuple with the ForexConversionHandlingCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingCharge

`func (o *ReconEntityDataInnerPaymentDetails) SetForexConversionHandlingCharge(v string)`

SetForexConversionHandlingCharge sets ForexConversionHandlingCharge field to given value.

### HasForexConversionHandlingCharge

`func (o *ReconEntityDataInnerPaymentDetails) HasForexConversionHandlingCharge() bool`

HasForexConversionHandlingCharge returns a boolean if a field has been set.

### GetForexConversionHandlingTax

`func (o *ReconEntityDataInnerPaymentDetails) GetForexConversionHandlingTax() string`

GetForexConversionHandlingTax returns the ForexConversionHandlingTax field if non-nil, zero value otherwise.

### GetForexConversionHandlingTaxOk

`func (o *ReconEntityDataInnerPaymentDetails) GetForexConversionHandlingTaxOk() (*string, bool)`

GetForexConversionHandlingTaxOk returns a tuple with the ForexConversionHandlingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingTax

`func (o *ReconEntityDataInnerPaymentDetails) SetForexConversionHandlingTax(v string)`

SetForexConversionHandlingTax sets ForexConversionHandlingTax field to given value.

### HasForexConversionHandlingTax

`func (o *ReconEntityDataInnerPaymentDetails) HasForexConversionHandlingTax() bool`

HasForexConversionHandlingTax returns a boolean if a field has been set.

### GetChargesCurrency

`func (o *ReconEntityDataInnerPaymentDetails) GetChargesCurrency() string`

GetChargesCurrency returns the ChargesCurrency field if non-nil, zero value otherwise.

### GetChargesCurrencyOk

`func (o *ReconEntityDataInnerPaymentDetails) GetChargesCurrencyOk() (*string, bool)`

GetChargesCurrencyOk returns a tuple with the ChargesCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesCurrency

`func (o *ReconEntityDataInnerPaymentDetails) SetChargesCurrency(v string)`

SetChargesCurrency sets ChargesCurrency field to given value.

### HasChargesCurrency

`func (o *ReconEntityDataInnerPaymentDetails) HasChargesCurrency() bool`

HasChargesCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


