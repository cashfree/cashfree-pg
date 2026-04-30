# ReconEntityDataInnerSettlementDetails

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

### NewReconEntityDataInnerSettlementDetails

`func NewReconEntityDataInnerSettlementDetails() *ReconEntityDataInnerSettlementDetails`

NewReconEntityDataInnerSettlementDetails instantiates a new ReconEntityDataInnerSettlementDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityDataInnerSettlementDetailsWithDefaults

`func NewReconEntityDataInnerSettlementDetailsWithDefaults() *ReconEntityDataInnerSettlementDetails`

NewReconEntityDataInnerSettlementDetailsWithDefaults instantiates a new ReconEntityDataInnerSettlementDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementId

`func (o *ReconEntityDataInnerSettlementDetails) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *ReconEntityDataInnerSettlementDetails) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *ReconEntityDataInnerSettlementDetails) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *ReconEntityDataInnerSettlementDetails) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementDate

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *ReconEntityDataInnerSettlementDetails) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *ReconEntityDataInnerSettlementDetails) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetUtr

`func (o *ReconEntityDataInnerSettlementDetails) GetUtr() string`

GetUtr returns the Utr field if non-nil, zero value otherwise.

### GetUtrOk

`func (o *ReconEntityDataInnerSettlementDetails) GetUtrOk() (*string, bool)`

GetUtrOk returns a tuple with the Utr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtr

`func (o *ReconEntityDataInnerSettlementDetails) SetUtr(v string)`

SetUtr sets Utr field to given value.

### HasUtr

`func (o *ReconEntityDataInnerSettlementDetails) HasUtr() bool`

HasUtr returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetVendorCommission

`func (o *ReconEntityDataInnerSettlementDetails) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *ReconEntityDataInnerSettlementDetails) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *ReconEntityDataInnerSettlementDetails) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *ReconEntityDataInnerSettlementDetails) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetPaymentFrom

`func (o *ReconEntityDataInnerSettlementDetails) GetPaymentFrom() string`

GetPaymentFrom returns the PaymentFrom field if non-nil, zero value otherwise.

### GetPaymentFromOk

`func (o *ReconEntityDataInnerSettlementDetails) GetPaymentFromOk() (*string, bool)`

GetPaymentFromOk returns a tuple with the PaymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrom

`func (o *ReconEntityDataInnerSettlementDetails) SetPaymentFrom(v string)`

SetPaymentFrom sets PaymentFrom field to given value.

### HasPaymentFrom

`func (o *ReconEntityDataInnerSettlementDetails) HasPaymentFrom() bool`

HasPaymentFrom returns a boolean if a field has been set.

### GetPaymentTill

`func (o *ReconEntityDataInnerSettlementDetails) GetPaymentTill() string`

GetPaymentTill returns the PaymentTill field if non-nil, zero value otherwise.

### GetPaymentTillOk

`func (o *ReconEntityDataInnerSettlementDetails) GetPaymentTillOk() (*string, bool)`

GetPaymentTillOk returns a tuple with the PaymentTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTill

`func (o *ReconEntityDataInnerSettlementDetails) SetPaymentTill(v string)`

SetPaymentTill sets PaymentTill field to given value.

### HasPaymentTill

`func (o *ReconEntityDataInnerSettlementDetails) HasPaymentTill() bool`

HasPaymentTill returns a boolean if a field has been set.

### GetReason

`func (o *ReconEntityDataInnerSettlementDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReconEntityDataInnerSettlementDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReconEntityDataInnerSettlementDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReconEntityDataInnerSettlementDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRemarks

`func (o *ReconEntityDataInnerSettlementDetails) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ReconEntityDataInnerSettlementDetails) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ReconEntityDataInnerSettlementDetails) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ReconEntityDataInnerSettlementDetails) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *ReconEntityDataInnerSettlementDetails) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *ReconEntityDataInnerSettlementDetails) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *ReconEntityDataInnerSettlementDetails) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *ReconEntityDataInnerSettlementDetails) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetSettlementCharge

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementCharge() float32`

GetSettlementCharge returns the SettlementCharge field if non-nil, zero value otherwise.

### GetSettlementChargeOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementChargeOk() (*float32, bool)`

GetSettlementChargeOk returns a tuple with the SettlementCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCharge

`func (o *ReconEntityDataInnerSettlementDetails) SetSettlementCharge(v float32)`

SetSettlementCharge sets SettlementCharge field to given value.

### HasSettlementCharge

`func (o *ReconEntityDataInnerSettlementDetails) HasSettlementCharge() bool`

HasSettlementCharge returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *ReconEntityDataInnerSettlementDetails) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *ReconEntityDataInnerSettlementDetails) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementTax

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementTax() float32`

GetSettlementTax returns the SettlementTax field if non-nil, zero value otherwise.

### GetSettlementTaxOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementTaxOk() (*float32, bool)`

GetSettlementTaxOk returns a tuple with the SettlementTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTax

`func (o *ReconEntityDataInnerSettlementDetails) SetSettlementTax(v float32)`

SetSettlementTax sets SettlementTax field to given value.

### HasSettlementTax

`func (o *ReconEntityDataInnerSettlementDetails) HasSettlementTax() bool`

HasSettlementTax returns a boolean if a field has been set.

### GetSettlementType

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *ReconEntityDataInnerSettlementDetails) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *ReconEntityDataInnerSettlementDetails) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *ReconEntityDataInnerSettlementDetails) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


