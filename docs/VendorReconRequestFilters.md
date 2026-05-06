# VendorReconRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementId** | Pointer to **int32** | Specify the Settlement ID for which you want to fetch the order details. Relevant for \&quot;View Split Order Details Using Settlement ID\&quot;. | [optional] 
**MerchantVendorId** | Pointer to **string** | Specify the Vendor ID for which you want to fetch the recon details. Relevant for \&quot;Vendor Recon Using Vendor ID &amp; Time Interval\&quot;. | [optional] 
**StartDate** | Pointer to **string** | Start date for fetching reconciliation details. Relevant for \&quot;Vendor Recon for a Time Period\&quot; and \&quot;Vendor Recon Using Vendor ID &amp; Time Interval\&quot;. | [optional] 
**EndDate** | Pointer to **string** | End date for fetching reconciliation details. Relevant for \&quot;Vendor Recon for a Time Period\&quot; and \&quot;Vendor Recon Using Vendor ID &amp; Time Interval\&quot;. | [optional] 

## Methods

### NewVendorReconRequestFilters

`func NewVendorReconRequestFilters() *VendorReconRequestFilters`

NewVendorReconRequestFilters instantiates a new VendorReconRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorReconRequestFiltersWithDefaults

`func NewVendorReconRequestFiltersWithDefaults() *VendorReconRequestFilters`

NewVendorReconRequestFiltersWithDefaults instantiates a new VendorReconRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementId

`func (o *VendorReconRequestFilters) GetSettlementId() int32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *VendorReconRequestFilters) GetSettlementIdOk() (*int32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *VendorReconRequestFilters) SetSettlementId(v int32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *VendorReconRequestFilters) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetMerchantVendorId

`func (o *VendorReconRequestFilters) GetMerchantVendorId() string`

GetMerchantVendorId returns the MerchantVendorId field if non-nil, zero value otherwise.

### GetMerchantVendorIdOk

`func (o *VendorReconRequestFilters) GetMerchantVendorIdOk() (*string, bool)`

GetMerchantVendorIdOk returns a tuple with the MerchantVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantVendorId

`func (o *VendorReconRequestFilters) SetMerchantVendorId(v string)`

SetMerchantVendorId sets MerchantVendorId field to given value.

### HasMerchantVendorId

`func (o *VendorReconRequestFilters) HasMerchantVendorId() bool`

HasMerchantVendorId returns a boolean if a field has been set.

### GetStartDate

`func (o *VendorReconRequestFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *VendorReconRequestFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *VendorReconRequestFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *VendorReconRequestFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *VendorReconRequestFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *VendorReconRequestFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *VendorReconRequestFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *VendorReconRequestFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


