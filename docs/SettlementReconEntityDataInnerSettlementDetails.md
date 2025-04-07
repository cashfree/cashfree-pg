# SettlementReconEntityDataInnerSettlementDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfSettlementId** | Pointer to **string** | Unique ID to identify the settlement. | [optional] 
**SettlementDate** | Pointer to **string** | Date and time when the settlement was processed. | [optional] 
**Utr** | Pointer to **string** | Unique transaction reference number of the settlement. | [optional] 
**SplitServiceCharge** | Pointer to **float32** | Service charge that is applicable for splitting the payment. | [optional] 
**SplitServiceTax** | Pointer to **float32** | Service tax applicable for splitting the amount to vendors. | [optional] 
**VendorCommission** | Pointer to **float32** | Vendor commission applicable for this transaction. | [optional] 
**PaymentFrom** | Pointer to **string** | Date and time from settlement computed. | [optional] 
**PaymentTill** | Pointer to **string** | Date and time till settlement computed. | [optional] 
**Reason** | Pointer to **string** | If any reason for settlement failure. | [optional] 
**Remarks** | Pointer to **string** | Remarks related for settlement. | [optional] 
**ServiceCharge** | Pointer to **float32** | Service charge for the transactions. | [optional] 
**ServiceTax** | Pointer to **float32** | Service tax for the transactions. | [optional] 
**SettlementCharge** | Pointer to **float32** | Settlement Service Charge. | [optional] 
**SettlementInitiatedOn** | Pointer to **string** | Date and time when Settlement initiated. | [optional] 
**SettlementTax** | Pointer to **float32** | Settlement Service Tax. | [optional] 
**SettlementType** | Pointer to **string** | Type of Settlement, Example - Normal Settlement. | [optional] 

## Methods

### NewSettlementReconEntityDataInnerSettlementDetails

`func NewSettlementReconEntityDataInnerSettlementDetails() *SettlementReconEntityDataInnerSettlementDetails`

NewSettlementReconEntityDataInnerSettlementDetails instantiates a new SettlementReconEntityDataInnerSettlementDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerSettlementDetailsWithDefaults

`func NewSettlementReconEntityDataInnerSettlementDetailsWithDefaults() *SettlementReconEntityDataInnerSettlementDetails`

NewSettlementReconEntityDataInnerSettlementDetailsWithDefaults instantiates a new SettlementReconEntityDataInnerSettlementDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementId

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementDate

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetUtr

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetUtr() string`

GetUtr returns the Utr field if non-nil, zero value otherwise.

### GetUtrOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetUtrOk() (*string, bool)`

GetUtrOk returns a tuple with the Utr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtr

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetUtr(v string)`

SetUtr sets Utr field to given value.

### HasUtr

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasUtr() bool`

HasUtr returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetVendorCommission

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetPaymentFrom

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetPaymentFrom() string`

GetPaymentFrom returns the PaymentFrom field if non-nil, zero value otherwise.

### GetPaymentFromOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetPaymentFromOk() (*string, bool)`

GetPaymentFromOk returns a tuple with the PaymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrom

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetPaymentFrom(v string)`

SetPaymentFrom sets PaymentFrom field to given value.

### HasPaymentFrom

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasPaymentFrom() bool`

HasPaymentFrom returns a boolean if a field has been set.

### GetPaymentTill

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetPaymentTill() string`

GetPaymentTill returns the PaymentTill field if non-nil, zero value otherwise.

### GetPaymentTillOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetPaymentTillOk() (*string, bool)`

GetPaymentTillOk returns a tuple with the PaymentTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTill

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetPaymentTill(v string)`

SetPaymentTill sets PaymentTill field to given value.

### HasPaymentTill

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasPaymentTill() bool`

HasPaymentTill returns a boolean if a field has been set.

### GetReason

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRemarks

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetSettlementCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementCharge() float32`

GetSettlementCharge returns the SettlementCharge field if non-nil, zero value otherwise.

### GetSettlementChargeOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementChargeOk() (*float32, bool)`

GetSettlementChargeOk returns a tuple with the SettlementCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSettlementCharge(v float32)`

SetSettlementCharge sets SettlementCharge field to given value.

### HasSettlementCharge

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSettlementCharge() bool`

HasSettlementCharge returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementTax() float32`

GetSettlementTax returns the SettlementTax field if non-nil, zero value otherwise.

### GetSettlementTaxOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementTaxOk() (*float32, bool)`

GetSettlementTaxOk returns a tuple with the SettlementTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSettlementTax(v float32)`

SetSettlementTax sets SettlementTax field to given value.

### HasSettlementTax

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSettlementTax() bool`

HasSettlementTax returns a boolean if a field has been set.

### GetSettlementType

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *SettlementReconEntityDataInnerSettlementDetails) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *SettlementReconEntityDataInnerSettlementDetails) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *SettlementReconEntityDataInnerSettlementDetails) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


