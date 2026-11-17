# OrderSettlementEntityPaymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** | Cashfree Payments unique ID to identify the payment. | [optional] 
**PaymentTime** | Pointer to **string** | Date and time when the transaction was completed. | [optional] 
**PaymentCurrency** | Pointer to **string** | Currency of the payment. | [optional] 
**PaymentAmount** | Pointer to **float32** | Net amount of the transaction that was a part of this settlement. | [optional] [default to 0.0]
**PgServiceCharge** | Pointer to **float32** | PG service charges that were deducted in this settlement. | [optional] [default to 0.0]
**PgServiceTax** | Pointer to **float32** | PG service taxes that were deducted in this settlement. | [optional] [default to 0.0]
**SplitServiceCharge** | Pointer to **float32** | Split service charges that were deducted in this settlement. | [optional] [default to 0.0]
**SplitServiceTax** | Pointer to **float32** | Split service taxes that were deducted in this settlement. | [optional] [default to 0.0]
**Surcharge** | Pointer to **float32** | Surcharge collected from the customer. | [optional] [default to 0.0]
**TaxOnSurcharge** | Pointer to **float32** | Tax on surcharge. Upto 4 decimal points. | [optional] [default to 0.0]
**ChargesCurrency** | Pointer to **string** | Currency in which service charges are applied. | [optional] 
**ForexConversionHandlingCharge** | Pointer to **float32** | Forex conversion handling charge. | [optional] [default to 0.0]
**ForexConversionHandlingTax** | Pointer to **float32** | Forex conversion handling tax. | [optional] [default to 0.0]
**ForexConversionRate** | Pointer to **float32** | Forex conversion rate applied. | [optional] [default to 0.0]
**VendorCommission** | Pointer to **float32** | Vendor commission/split that was deducted in this settlement. | [optional] [default to 0.0]
**SettlementAmount** | Pointer to **float32** | Amount after deducting all charges. | [optional] [default to 0.0]
**ChargeType** | Pointer to **string** | Indicates how Cashfree&#39;s service charges are applied for this payment. | [optional] 
**ChargedTo** | Pointer to **string** | Indicates who bears Cashfree&#39;s service charges. | [optional] 

## Methods

### NewOrderSettlementEntityPaymentDetails

`func NewOrderSettlementEntityPaymentDetails() *OrderSettlementEntityPaymentDetails`

NewOrderSettlementEntityPaymentDetails instantiates a new OrderSettlementEntityPaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSettlementEntityPaymentDetailsWithDefaults

`func NewOrderSettlementEntityPaymentDetailsWithDefaults() *OrderSettlementEntityPaymentDetails`

NewOrderSettlementEntityPaymentDetailsWithDefaults instantiates a new OrderSettlementEntityPaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *OrderSettlementEntityPaymentDetails) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *OrderSettlementEntityPaymentDetails) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *OrderSettlementEntityPaymentDetails) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *OrderSettlementEntityPaymentDetails) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetPaymentTime

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *OrderSettlementEntityPaymentDetails) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *OrderSettlementEntityPaymentDetails) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *OrderSettlementEntityPaymentDetails) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *OrderSettlementEntityPaymentDetails) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *OrderSettlementEntityPaymentDetails) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *OrderSettlementEntityPaymentDetails) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *OrderSettlementEntityPaymentDetails) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPgServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) GetPgServiceCharge() float32`

GetPgServiceCharge returns the PgServiceCharge field if non-nil, zero value otherwise.

### GetPgServiceChargeOk

`func (o *OrderSettlementEntityPaymentDetails) GetPgServiceChargeOk() (*float32, bool)`

GetPgServiceChargeOk returns a tuple with the PgServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) SetPgServiceCharge(v float32)`

SetPgServiceCharge sets PgServiceCharge field to given value.

### HasPgServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) HasPgServiceCharge() bool`

HasPgServiceCharge returns a boolean if a field has been set.

### GetPgServiceTax

`func (o *OrderSettlementEntityPaymentDetails) GetPgServiceTax() float32`

GetPgServiceTax returns the PgServiceTax field if non-nil, zero value otherwise.

### GetPgServiceTaxOk

`func (o *OrderSettlementEntityPaymentDetails) GetPgServiceTaxOk() (*float32, bool)`

GetPgServiceTaxOk returns a tuple with the PgServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceTax

`func (o *OrderSettlementEntityPaymentDetails) SetPgServiceTax(v float32)`

SetPgServiceTax sets PgServiceTax field to given value.

### HasPgServiceTax

`func (o *OrderSettlementEntityPaymentDetails) HasPgServiceTax() bool`

HasPgServiceTax returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *OrderSettlementEntityPaymentDetails) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *OrderSettlementEntityPaymentDetails) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *OrderSettlementEntityPaymentDetails) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *OrderSettlementEntityPaymentDetails) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *OrderSettlementEntityPaymentDetails) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *OrderSettlementEntityPaymentDetails) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetSurcharge

`func (o *OrderSettlementEntityPaymentDetails) GetSurcharge() float32`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *OrderSettlementEntityPaymentDetails) GetSurchargeOk() (*float32, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *OrderSettlementEntityPaymentDetails) SetSurcharge(v float32)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *OrderSettlementEntityPaymentDetails) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetTaxOnSurcharge

`func (o *OrderSettlementEntityPaymentDetails) GetTaxOnSurcharge() float32`

GetTaxOnSurcharge returns the TaxOnSurcharge field if non-nil, zero value otherwise.

### GetTaxOnSurchargeOk

`func (o *OrderSettlementEntityPaymentDetails) GetTaxOnSurchargeOk() (*float32, bool)`

GetTaxOnSurchargeOk returns a tuple with the TaxOnSurcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxOnSurcharge

`func (o *OrderSettlementEntityPaymentDetails) SetTaxOnSurcharge(v float32)`

SetTaxOnSurcharge sets TaxOnSurcharge field to given value.

### HasTaxOnSurcharge

`func (o *OrderSettlementEntityPaymentDetails) HasTaxOnSurcharge() bool`

HasTaxOnSurcharge returns a boolean if a field has been set.

### GetChargesCurrency

`func (o *OrderSettlementEntityPaymentDetails) GetChargesCurrency() string`

GetChargesCurrency returns the ChargesCurrency field if non-nil, zero value otherwise.

### GetChargesCurrencyOk

`func (o *OrderSettlementEntityPaymentDetails) GetChargesCurrencyOk() (*string, bool)`

GetChargesCurrencyOk returns a tuple with the ChargesCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesCurrency

`func (o *OrderSettlementEntityPaymentDetails) SetChargesCurrency(v string)`

SetChargesCurrency sets ChargesCurrency field to given value.

### HasChargesCurrency

`func (o *OrderSettlementEntityPaymentDetails) HasChargesCurrency() bool`

HasChargesCurrency returns a boolean if a field has been set.

### GetForexConversionHandlingCharge

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionHandlingCharge() float32`

GetForexConversionHandlingCharge returns the ForexConversionHandlingCharge field if non-nil, zero value otherwise.

### GetForexConversionHandlingChargeOk

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionHandlingChargeOk() (*float32, bool)`

GetForexConversionHandlingChargeOk returns a tuple with the ForexConversionHandlingCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingCharge

`func (o *OrderSettlementEntityPaymentDetails) SetForexConversionHandlingCharge(v float32)`

SetForexConversionHandlingCharge sets ForexConversionHandlingCharge field to given value.

### HasForexConversionHandlingCharge

`func (o *OrderSettlementEntityPaymentDetails) HasForexConversionHandlingCharge() bool`

HasForexConversionHandlingCharge returns a boolean if a field has been set.

### GetForexConversionHandlingTax

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionHandlingTax() float32`

GetForexConversionHandlingTax returns the ForexConversionHandlingTax field if non-nil, zero value otherwise.

### GetForexConversionHandlingTaxOk

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionHandlingTaxOk() (*float32, bool)`

GetForexConversionHandlingTaxOk returns a tuple with the ForexConversionHandlingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionHandlingTax

`func (o *OrderSettlementEntityPaymentDetails) SetForexConversionHandlingTax(v float32)`

SetForexConversionHandlingTax sets ForexConversionHandlingTax field to given value.

### HasForexConversionHandlingTax

`func (o *OrderSettlementEntityPaymentDetails) HasForexConversionHandlingTax() bool`

HasForexConversionHandlingTax returns a boolean if a field has been set.

### GetForexConversionRate

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionRate() float32`

GetForexConversionRate returns the ForexConversionRate field if non-nil, zero value otherwise.

### GetForexConversionRateOk

`func (o *OrderSettlementEntityPaymentDetails) GetForexConversionRateOk() (*float32, bool)`

GetForexConversionRateOk returns a tuple with the ForexConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexConversionRate

`func (o *OrderSettlementEntityPaymentDetails) SetForexConversionRate(v float32)`

SetForexConversionRate sets ForexConversionRate field to given value.

### HasForexConversionRate

`func (o *OrderSettlementEntityPaymentDetails) HasForexConversionRate() bool`

HasForexConversionRate returns a boolean if a field has been set.

### GetVendorCommission

`func (o *OrderSettlementEntityPaymentDetails) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *OrderSettlementEntityPaymentDetails) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *OrderSettlementEntityPaymentDetails) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *OrderSettlementEntityPaymentDetails) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *OrderSettlementEntityPaymentDetails) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *OrderSettlementEntityPaymentDetails) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *OrderSettlementEntityPaymentDetails) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *OrderSettlementEntityPaymentDetails) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetChargeType

`func (o *OrderSettlementEntityPaymentDetails) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *OrderSettlementEntityPaymentDetails) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *OrderSettlementEntityPaymentDetails) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *OrderSettlementEntityPaymentDetails) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetChargedTo

`func (o *OrderSettlementEntityPaymentDetails) GetChargedTo() string`

GetChargedTo returns the ChargedTo field if non-nil, zero value otherwise.

### GetChargedToOk

`func (o *OrderSettlementEntityPaymentDetails) GetChargedToOk() (*string, bool)`

GetChargedToOk returns a tuple with the ChargedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedTo

`func (o *OrderSettlementEntityPaymentDetails) SetChargedTo(v string)`

SetChargedTo sets ChargedTo field to given value.

### HasChargedTo

`func (o *OrderSettlementEntityPaymentDetails) HasChargedTo() bool`

HasChargedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


