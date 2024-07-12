# VendorEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**VendorId** | Pointer to **string** |  | [optional] 
**AddedOn** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to [**[]BankDetails**](BankDetails.md) |  | [optional] 
**Upi** | Pointer to **string** |  | [optional] 
**ScheduleOption** | Pointer to [**[]ScheduleOption**](ScheduleOption.md) |  | [optional] 
**VendorType** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**BusinessType** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RelatedDocs** | Pointer to [**[]UpdateVendorResponseRelatedDocsInner**](UpdateVendorResponseRelatedDocsInner.md) |  | [optional] 

## Methods

### NewVendorEntity

`func NewVendorEntity() *VendorEntity`

NewVendorEntity instantiates a new VendorEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorEntityWithDefaults

`func NewVendorEntityWithDefaults() *VendorEntity`

NewVendorEntityWithDefaults instantiates a new VendorEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *VendorEntity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *VendorEntity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *VendorEntity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *VendorEntity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *VendorEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VendorEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VendorEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VendorEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhone

`func (o *VendorEntity) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VendorEntity) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VendorEntity) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *VendorEntity) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetName

`func (o *VendorEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VendorEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VendorEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VendorEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendorId

`func (o *VendorEntity) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *VendorEntity) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *VendorEntity) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *VendorEntity) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetAddedOn

`func (o *VendorEntity) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *VendorEntity) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *VendorEntity) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *VendorEntity) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *VendorEntity) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *VendorEntity) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *VendorEntity) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *VendorEntity) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetBank

`func (o *VendorEntity) GetBank() []BankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *VendorEntity) GetBankOk() (*[]BankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *VendorEntity) SetBank(v []BankDetails)`

SetBank sets Bank field to given value.

### HasBank

`func (o *VendorEntity) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetUpi

`func (o *VendorEntity) GetUpi() string`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *VendorEntity) GetUpiOk() (*string, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *VendorEntity) SetUpi(v string)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *VendorEntity) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetScheduleOption

`func (o *VendorEntity) GetScheduleOption() []ScheduleOption`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *VendorEntity) GetScheduleOptionOk() (*[]ScheduleOption, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *VendorEntity) SetScheduleOption(v []ScheduleOption)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *VendorEntity) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetVendorType

`func (o *VendorEntity) GetVendorType() string`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *VendorEntity) GetVendorTypeOk() (*string, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *VendorEntity) SetVendorType(v string)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *VendorEntity) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.

### GetAccountType

`func (o *VendorEntity) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *VendorEntity) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *VendorEntity) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *VendorEntity) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBusinessType

`func (o *VendorEntity) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *VendorEntity) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *VendorEntity) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *VendorEntity) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetRemarks

`func (o *VendorEntity) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *VendorEntity) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *VendorEntity) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *VendorEntity) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRelatedDocs

`func (o *VendorEntity) GetRelatedDocs() []UpdateVendorResponseRelatedDocsInner`

GetRelatedDocs returns the RelatedDocs field if non-nil, zero value otherwise.

### GetRelatedDocsOk

`func (o *VendorEntity) GetRelatedDocsOk() (*[]UpdateVendorResponseRelatedDocsInner, bool)`

GetRelatedDocsOk returns a tuple with the RelatedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocs

`func (o *VendorEntity) SetRelatedDocs(v []UpdateVendorResponseRelatedDocsInner)`

SetRelatedDocs sets RelatedDocs field to given value.

### HasRelatedDocs

`func (o *VendorEntity) HasRelatedDocs() bool`

HasRelatedDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


