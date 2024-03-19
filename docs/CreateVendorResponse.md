# CreateVendorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to [**[]BankDetails**](BankDetails.md) |  | [optional] 
**Upi** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**VendorId** | Pointer to **string** |  | [optional] 
**ScheduleOption** | Pointer to [**[]ScheduleOption**](ScheduleOption.md) |  | [optional] 
**KycDetails** | Pointer to [**[]KycDetails**](KycDetails.md) |  | [optional] 
**DashboardAccess** | Pointer to **bool** |  | [optional] 
**BankDetails** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateVendorResponse

`func NewCreateVendorResponse() *CreateVendorResponse`

NewCreateVendorResponse instantiates a new CreateVendorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVendorResponseWithDefaults

`func NewCreateVendorResponseWithDefaults() *CreateVendorResponse`

NewCreateVendorResponseWithDefaults instantiates a new CreateVendorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateVendorResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateVendorResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateVendorResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateVendorResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *CreateVendorResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateVendorResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateVendorResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateVendorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBank

`func (o *CreateVendorResponse) GetBank() []BankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *CreateVendorResponse) GetBankOk() (*[]BankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *CreateVendorResponse) SetBank(v []BankDetails)`

SetBank sets Bank field to given value.

### HasBank

`func (o *CreateVendorResponse) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetUpi

`func (o *CreateVendorResponse) GetUpi() string`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *CreateVendorResponse) GetUpiOk() (*string, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *CreateVendorResponse) SetUpi(v string)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *CreateVendorResponse) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetPhone

`func (o *CreateVendorResponse) GetPhone() float32`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateVendorResponse) GetPhoneOk() (*float32, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateVendorResponse) SetPhone(v float32)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateVendorResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetName

`func (o *CreateVendorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVendorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVendorResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVendorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendorId

`func (o *CreateVendorResponse) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *CreateVendorResponse) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *CreateVendorResponse) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *CreateVendorResponse) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetScheduleOption

`func (o *CreateVendorResponse) GetScheduleOption() []ScheduleOption`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *CreateVendorResponse) GetScheduleOptionOk() (*[]ScheduleOption, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *CreateVendorResponse) SetScheduleOption(v []ScheduleOption)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *CreateVendorResponse) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetKycDetails

`func (o *CreateVendorResponse) GetKycDetails() []KycDetails`

GetKycDetails returns the KycDetails field if non-nil, zero value otherwise.

### GetKycDetailsOk

`func (o *CreateVendorResponse) GetKycDetailsOk() (*[]KycDetails, bool)`

GetKycDetailsOk returns a tuple with the KycDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDetails

`func (o *CreateVendorResponse) SetKycDetails(v []KycDetails)`

SetKycDetails sets KycDetails field to given value.

### HasKycDetails

`func (o *CreateVendorResponse) HasKycDetails() bool`

HasKycDetails returns a boolean if a field has been set.

### GetDashboardAccess

`func (o *CreateVendorResponse) GetDashboardAccess() bool`

GetDashboardAccess returns the DashboardAccess field if non-nil, zero value otherwise.

### GetDashboardAccessOk

`func (o *CreateVendorResponse) GetDashboardAccessOk() (*bool, bool)`

GetDashboardAccessOk returns a tuple with the DashboardAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardAccess

`func (o *CreateVendorResponse) SetDashboardAccess(v bool)`

SetDashboardAccess sets DashboardAccess field to given value.

### HasDashboardAccess

`func (o *CreateVendorResponse) HasDashboardAccess() bool`

HasDashboardAccess returns a boolean if a field has been set.

### GetBankDetails

`func (o *CreateVendorResponse) GetBankDetails() string`

GetBankDetails returns the BankDetails field if non-nil, zero value otherwise.

### GetBankDetailsOk

`func (o *CreateVendorResponse) GetBankDetailsOk() (*string, bool)`

GetBankDetailsOk returns a tuple with the BankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankDetails

`func (o *CreateVendorResponse) SetBankDetails(v string)`

SetBankDetails sets BankDetails field to given value.

### HasBankDetails

`func (o *CreateVendorResponse) HasBankDetails() bool`

HasBankDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


