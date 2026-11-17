# GetAllSettlementsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfSettlementId** | Pointer to **string** | Unique ID to identify the settlement. | [optional] 
**Status** | Pointer to **string** | Status of the settlement. | [optional] 
**StatusDescription** | Pointer to **string** | Human-readable description of the settlement status. | [optional] 
**SettlementUtr** | Pointer to **string** | Unique transaction reference number of the settlement. | [optional] 
**PaymentFrom** | Pointer to **string** | Start of the payment period included in this settlement. | [optional] 
**PaymentTill** | Pointer to **string** | End of the payment period included in this settlement. | [optional] 
**SettlementInitiatedOn** | Pointer to **string** | Date and time when the settlement was initiated. | [optional] 
**SettlementProcessedOn** | Pointer to **string** | Date and time when the settlement was processed. | [optional] 
**PaymentAmount** | Pointer to **float32** | Net amount of transactions that was a part of this settlement. | [optional] [default to 0.0]
**PgServiceCharge** | Pointer to **float32** | PG service charges that were deducted in this settlement. | [optional] [default to 0.0]
**PgServiceTax** | Pointer to **float32** | PG service taxes that were deducted in this settlement. | [optional] [default to 0.0]
**Adjustment** | Pointer to **float32** | Adjustments that were a part of this settlement. | [optional] [default to 0.0]
**SettlementCharge** | Pointer to **float32** | Settlement charges that were deducted in this settlement. | [optional] [default to 0.0]
**SettlementTax** | Pointer to **float32** | Settlement taxes that were deducted in this settlement. | [optional] [default to 0.0]
**SplitServiceCharge** | Pointer to **float32** | Split service charges that were deducted in this settlement. | [optional] [default to 0.0]
**SplitServiceTax** | Pointer to **float32** | Split service taxes that were deducted in this settlement. | [optional] [default to 0.0]
**VendorCommission** | Pointer to **float32** | Vendor commission/split that was deducted in this settlement. | [optional] [default to 0.0]
**AmountSettled** | Pointer to **float32** | Amount that was settled to the merchant&#39;s bank account. | [optional] [default to 0.0]
**ChargeType** | Pointer to **string** | Indicates how Cashfree&#39;s service charges are applied for this settlement. | [optional] 
**SettlementType** | Pointer to **string** | Type of settlement. Example - NORMAL_SETTLEMENT. | [optional] 

## Methods

### NewGetAllSettlementsResponseDataInner

`func NewGetAllSettlementsResponseDataInner() *GetAllSettlementsResponseDataInner`

NewGetAllSettlementsResponseDataInner instantiates a new GetAllSettlementsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSettlementsResponseDataInnerWithDefaults

`func NewGetAllSettlementsResponseDataInnerWithDefaults() *GetAllSettlementsResponseDataInner`

NewGetAllSettlementsResponseDataInnerWithDefaults instantiates a new GetAllSettlementsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementId

`func (o *GetAllSettlementsResponseDataInner) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *GetAllSettlementsResponseDataInner) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *GetAllSettlementsResponseDataInner) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *GetAllSettlementsResponseDataInner) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetStatus

`func (o *GetAllSettlementsResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllSettlementsResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllSettlementsResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAllSettlementsResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *GetAllSettlementsResponseDataInner) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *GetAllSettlementsResponseDataInner) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *GetAllSettlementsResponseDataInner) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *GetAllSettlementsResponseDataInner) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetSettlementUtr

`func (o *GetAllSettlementsResponseDataInner) GetSettlementUtr() string`

GetSettlementUtr returns the SettlementUtr field if non-nil, zero value otherwise.

### GetSettlementUtrOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementUtrOk() (*string, bool)`

GetSettlementUtrOk returns a tuple with the SettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtr

`func (o *GetAllSettlementsResponseDataInner) SetSettlementUtr(v string)`

SetSettlementUtr sets SettlementUtr field to given value.

### HasSettlementUtr

`func (o *GetAllSettlementsResponseDataInner) HasSettlementUtr() bool`

HasSettlementUtr returns a boolean if a field has been set.

### GetPaymentFrom

`func (o *GetAllSettlementsResponseDataInner) GetPaymentFrom() string`

GetPaymentFrom returns the PaymentFrom field if non-nil, zero value otherwise.

### GetPaymentFromOk

`func (o *GetAllSettlementsResponseDataInner) GetPaymentFromOk() (*string, bool)`

GetPaymentFromOk returns a tuple with the PaymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrom

`func (o *GetAllSettlementsResponseDataInner) SetPaymentFrom(v string)`

SetPaymentFrom sets PaymentFrom field to given value.

### HasPaymentFrom

`func (o *GetAllSettlementsResponseDataInner) HasPaymentFrom() bool`

HasPaymentFrom returns a boolean if a field has been set.

### GetPaymentTill

`func (o *GetAllSettlementsResponseDataInner) GetPaymentTill() string`

GetPaymentTill returns the PaymentTill field if non-nil, zero value otherwise.

### GetPaymentTillOk

`func (o *GetAllSettlementsResponseDataInner) GetPaymentTillOk() (*string, bool)`

GetPaymentTillOk returns a tuple with the PaymentTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTill

`func (o *GetAllSettlementsResponseDataInner) SetPaymentTill(v string)`

SetPaymentTill sets PaymentTill field to given value.

### HasPaymentTill

`func (o *GetAllSettlementsResponseDataInner) HasPaymentTill() bool`

HasPaymentTill returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *GetAllSettlementsResponseDataInner) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *GetAllSettlementsResponseDataInner) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *GetAllSettlementsResponseDataInner) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementProcessedOn

`func (o *GetAllSettlementsResponseDataInner) GetSettlementProcessedOn() string`

GetSettlementProcessedOn returns the SettlementProcessedOn field if non-nil, zero value otherwise.

### GetSettlementProcessedOnOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementProcessedOnOk() (*string, bool)`

GetSettlementProcessedOnOk returns a tuple with the SettlementProcessedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementProcessedOn

`func (o *GetAllSettlementsResponseDataInner) SetSettlementProcessedOn(v string)`

SetSettlementProcessedOn sets SettlementProcessedOn field to given value.

### HasSettlementProcessedOn

`func (o *GetAllSettlementsResponseDataInner) HasSettlementProcessedOn() bool`

HasSettlementProcessedOn returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *GetAllSettlementsResponseDataInner) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *GetAllSettlementsResponseDataInner) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *GetAllSettlementsResponseDataInner) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *GetAllSettlementsResponseDataInner) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPgServiceCharge

`func (o *GetAllSettlementsResponseDataInner) GetPgServiceCharge() float32`

GetPgServiceCharge returns the PgServiceCharge field if non-nil, zero value otherwise.

### GetPgServiceChargeOk

`func (o *GetAllSettlementsResponseDataInner) GetPgServiceChargeOk() (*float32, bool)`

GetPgServiceChargeOk returns a tuple with the PgServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceCharge

`func (o *GetAllSettlementsResponseDataInner) SetPgServiceCharge(v float32)`

SetPgServiceCharge sets PgServiceCharge field to given value.

### HasPgServiceCharge

`func (o *GetAllSettlementsResponseDataInner) HasPgServiceCharge() bool`

HasPgServiceCharge returns a boolean if a field has been set.

### GetPgServiceTax

`func (o *GetAllSettlementsResponseDataInner) GetPgServiceTax() float32`

GetPgServiceTax returns the PgServiceTax field if non-nil, zero value otherwise.

### GetPgServiceTaxOk

`func (o *GetAllSettlementsResponseDataInner) GetPgServiceTaxOk() (*float32, bool)`

GetPgServiceTaxOk returns a tuple with the PgServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceTax

`func (o *GetAllSettlementsResponseDataInner) SetPgServiceTax(v float32)`

SetPgServiceTax sets PgServiceTax field to given value.

### HasPgServiceTax

`func (o *GetAllSettlementsResponseDataInner) HasPgServiceTax() bool`

HasPgServiceTax returns a boolean if a field has been set.

### GetAdjustment

`func (o *GetAllSettlementsResponseDataInner) GetAdjustment() float32`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *GetAllSettlementsResponseDataInner) GetAdjustmentOk() (*float32, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *GetAllSettlementsResponseDataInner) SetAdjustment(v float32)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *GetAllSettlementsResponseDataInner) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetSettlementCharge

`func (o *GetAllSettlementsResponseDataInner) GetSettlementCharge() float32`

GetSettlementCharge returns the SettlementCharge field if non-nil, zero value otherwise.

### GetSettlementChargeOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementChargeOk() (*float32, bool)`

GetSettlementChargeOk returns a tuple with the SettlementCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCharge

`func (o *GetAllSettlementsResponseDataInner) SetSettlementCharge(v float32)`

SetSettlementCharge sets SettlementCharge field to given value.

### HasSettlementCharge

`func (o *GetAllSettlementsResponseDataInner) HasSettlementCharge() bool`

HasSettlementCharge returns a boolean if a field has been set.

### GetSettlementTax

`func (o *GetAllSettlementsResponseDataInner) GetSettlementTax() float32`

GetSettlementTax returns the SettlementTax field if non-nil, zero value otherwise.

### GetSettlementTaxOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementTaxOk() (*float32, bool)`

GetSettlementTaxOk returns a tuple with the SettlementTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTax

`func (o *GetAllSettlementsResponseDataInner) SetSettlementTax(v float32)`

SetSettlementTax sets SettlementTax field to given value.

### HasSettlementTax

`func (o *GetAllSettlementsResponseDataInner) HasSettlementTax() bool`

HasSettlementTax returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *GetAllSettlementsResponseDataInner) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *GetAllSettlementsResponseDataInner) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *GetAllSettlementsResponseDataInner) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *GetAllSettlementsResponseDataInner) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *GetAllSettlementsResponseDataInner) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *GetAllSettlementsResponseDataInner) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *GetAllSettlementsResponseDataInner) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *GetAllSettlementsResponseDataInner) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetVendorCommission

`func (o *GetAllSettlementsResponseDataInner) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *GetAllSettlementsResponseDataInner) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *GetAllSettlementsResponseDataInner) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *GetAllSettlementsResponseDataInner) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetAmountSettled

`func (o *GetAllSettlementsResponseDataInner) GetAmountSettled() float32`

GetAmountSettled returns the AmountSettled field if non-nil, zero value otherwise.

### GetAmountSettledOk

`func (o *GetAllSettlementsResponseDataInner) GetAmountSettledOk() (*float32, bool)`

GetAmountSettledOk returns a tuple with the AmountSettled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSettled

`func (o *GetAllSettlementsResponseDataInner) SetAmountSettled(v float32)`

SetAmountSettled sets AmountSettled field to given value.

### HasAmountSettled

`func (o *GetAllSettlementsResponseDataInner) HasAmountSettled() bool`

HasAmountSettled returns a boolean if a field has been set.

### GetChargeType

`func (o *GetAllSettlementsResponseDataInner) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *GetAllSettlementsResponseDataInner) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *GetAllSettlementsResponseDataInner) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *GetAllSettlementsResponseDataInner) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetSettlementType

`func (o *GetAllSettlementsResponseDataInner) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *GetAllSettlementsResponseDataInner) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *GetAllSettlementsResponseDataInner) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *GetAllSettlementsResponseDataInner) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


