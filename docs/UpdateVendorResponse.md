# UpdateVendorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to [**[]BankDetails**](BankDetails.md) |  | [optional] 
**Upi** | Pointer to **string** |  | [optional] 
**AddedOn** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**VendorType** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**BusinessType** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**VendorId** | Pointer to **string** |  | [optional] 
**ScheduleOption** | Pointer to [**[]ScheduleOption**](ScheduleOption.md) |  | [optional] 
**KycDetails** | Pointer to [**[]KycDetails**](KycDetails.md) |  | [optional] 
**DashboardAccess** | Pointer to **bool** |  | [optional] 
**BankDetails** | Pointer to **string** |  | [optional] 
**RelatedDocs** | Pointer to [**[]UpdateVendorResponseRelatedDocsInner**](UpdateVendorResponseRelatedDocsInner.md) |  | [optional] 

## Methods

### NewUpdateVendorResponse

`func NewUpdateVendorResponse() *UpdateVendorResponse`

NewUpdateVendorResponse instantiates a new UpdateVendorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVendorResponseWithDefaults

`func NewUpdateVendorResponseWithDefaults() *UpdateVendorResponse`

NewUpdateVendorResponseWithDefaults instantiates a new UpdateVendorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateVendorResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateVendorResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateVendorResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateVendorResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateVendorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateVendorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateVendorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateVendorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBank

`func (o *UpdateVendorResponse) GetBank() []BankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UpdateVendorResponse) GetBankOk() (*[]BankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UpdateVendorResponse) SetBank(v []BankDetails)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UpdateVendorResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetUpi

`func (o *UpdateVendorResponse) GetUpi() string`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *UpdateVendorResponse) GetUpiOk() (*string, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *UpdateVendorResponse) SetUpi(v string)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *UpdateVendorResponse) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetAddedOn

`func (o *UpdateVendorResponse) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *UpdateVendorResponse) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *UpdateVendorResponse) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *UpdateVendorResponse) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *UpdateVendorResponse) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *UpdateVendorResponse) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *UpdateVendorResponse) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *UpdateVendorResponse) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetVendorType

`func (o *UpdateVendorResponse) GetVendorType() string`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *UpdateVendorResponse) GetVendorTypeOk() (*string, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *UpdateVendorResponse) SetVendorType(v string)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *UpdateVendorResponse) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.

### GetAccountType

`func (o *UpdateVendorResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UpdateVendorResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UpdateVendorResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UpdateVendorResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBusinessType

`func (o *UpdateVendorResponse) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *UpdateVendorResponse) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *UpdateVendorResponse) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.

### HasBusinessType

`func (o *UpdateVendorResponse) HasBusinessType() bool`

HasBusinessType returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateVendorResponse) GetPhone() float32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateVendorResponse) GetPhoneOk() (*float32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateVendorResponse) SetPhone(v float32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateVendorResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetName

`func (o *UpdateVendorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVendorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVendorResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVendorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendorId

`func (o *UpdateVendorResponse) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *UpdateVendorResponse) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *UpdateVendorResponse) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *UpdateVendorResponse) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetScheduleOption

`func (o *UpdateVendorResponse) GetScheduleOption() []ScheduleOption`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *UpdateVendorResponse) GetScheduleOptionOk() (*[]ScheduleOption, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *UpdateVendorResponse) SetScheduleOption(v []ScheduleOption)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *UpdateVendorResponse) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetKycDetails

`func (o *UpdateVendorResponse) GetKycDetails() []KycDetails`

GetKycDetails returns the KycDetails field if non-nil, zero value otherwise.

### GetKycDetailsOk

`func (o *UpdateVendorResponse) GetKycDetailsOk() (*[]KycDetails, bool)`

GetKycDetailsOk returns a tuple with the KycDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDetails

`func (o *UpdateVendorResponse) SetKycDetails(v []KycDetails)`

SetKycDetails sets KycDetails field to given value.

### HasKycDetails

`func (o *UpdateVendorResponse) HasKycDetails() bool`

HasKycDetails returns a boolean if a field has been set.

### GetDashboardAccess

`func (o *UpdateVendorResponse) GetDashboardAccess() bool`

GetDashboardAccess returns the DashboardAccess field if non-nil, zero value otherwise.

### GetDashboardAccessOk

`func (o *UpdateVendorResponse) GetDashboardAccessOk() (*bool, bool)`

GetDashboardAccessOk returns a tuple with the DashboardAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardAccess

`func (o *UpdateVendorResponse) SetDashboardAccess(v bool)`

SetDashboardAccess sets DashboardAccess field to given value.

### HasDashboardAccess

`func (o *UpdateVendorResponse) HasDashboardAccess() bool`

HasDashboardAccess returns a boolean if a field has been set.

### GetBankDetails

`func (o *UpdateVendorResponse) GetBankDetails() string`

GetBankDetails returns the BankDetails field if non-nil, zero value otherwise.

### GetBankDetailsOk

`func (o *UpdateVendorResponse) GetBankDetailsOk() (*string, bool)`

GetBankDetailsOk returns a tuple with the BankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankDetails

`func (o *UpdateVendorResponse) SetBankDetails(v string)`

SetBankDetails sets BankDetails field to given value.

### HasBankDetails

`func (o *UpdateVendorResponse) HasBankDetails() bool`

HasBankDetails returns a boolean if a field has been set.

### GetRelatedDocs

`func (o *UpdateVendorResponse) GetRelatedDocs() []UpdateVendorResponseRelatedDocsInner`

GetRelatedDocs returns the RelatedDocs field if non-nil, zero value otherwise.

### GetRelatedDocsOk

`func (o *UpdateVendorResponse) GetRelatedDocsOk() (*[]UpdateVendorResponseRelatedDocsInner, bool)`

GetRelatedDocsOk returns a tuple with the RelatedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocs

`func (o *UpdateVendorResponse) SetRelatedDocs(v []UpdateVendorResponseRelatedDocsInner)`

SetRelatedDocs sets RelatedDocs field to given value.

### HasRelatedDocs

`func (o *UpdateVendorResponse) HasRelatedDocs() bool`

HasRelatedDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


