# CreateVendorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | Specify the unique Vendor ID to identify the beneficiary. Alphanumeric and underscore (_) allowed. | 
**Status** | **string** | Specify the status of vendor that should be updated. Possible values: ACTIVE,BLOCKED, DELETED | 
**Name** | **string** | Specify the name of the vendor to be updated. Name should not have any special character except . / - &amp; | 
**Email** | **string** | Specify the vendor email ID that should be updated. String in email ID format (Ex:johndoe_1@cashfree.com) should contain @ and dot (.) | 
**Phone** | **string** | Specify the beneficiaries phone number to be updated. Phone number registered in India (only digits, 8 - 12 characters after excluding +91). | 
**VerifyAccount** | Pointer to **bool** | Specify if the vendor bank account details should be verified. Possible values: true or false | [optional] 
**DashboardAccess** | Pointer to **bool** | Update if the vendor will have dashboard access or not. Possible values are: true or false | [optional] 
**ScheduleOption** | Pointer to **float32** | Specify the settlement cycle to be updated. View the settlement cycle details from the \&quot;Settlement Cycles Supported\&quot; table. If no schedule option is configured, the settlement cycle ID \&quot;1\&quot; will be in effect. Select \&quot;8\&quot; or \&quot;9\&quot; if you want to schedule instant vendor settlements. | [optional] 
**Bank** | Pointer to [**[]BankDetails**](BankDetails.md) | Specify the vendor bank account details to be updated. | [optional] 
**Upi** | Pointer to [**[]UpiDetails**](UpiDetails.md) | Updated beneficiary upi vpa. Alphanumeric, dot (.), hyphen (-), at sign (@), and underscore allowed (100 character limit). Note: underscore and dot (.) gets accepted before and after @, but hyphen (-) is only accepted before @ sign. | [optional] 
**KycDetails** | [**[]KycDetails**](KycDetails.md) | Specify the kyc details that should be updated. | 

## Methods

### NewCreateVendorRequest

`func NewCreateVendorRequest(vendorId string, status string, name string, email string, phone string, kycDetails []KycDetails, ) *CreateVendorRequest`

NewCreateVendorRequest instantiates a new CreateVendorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVendorRequestWithDefaults

`func NewCreateVendorRequestWithDefaults() *CreateVendorRequest`

NewCreateVendorRequestWithDefaults instantiates a new CreateVendorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *CreateVendorRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *CreateVendorRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *CreateVendorRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetStatus

`func (o *CreateVendorRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateVendorRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateVendorRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetName

`func (o *CreateVendorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVendorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVendorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateVendorRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateVendorRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateVendorRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *CreateVendorRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateVendorRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateVendorRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetVerifyAccount

`func (o *CreateVendorRequest) GetVerifyAccount() bool`

GetVerifyAccount returns the VerifyAccount field if non-nil, zero value otherwise.

### GetVerifyAccountOk

`func (o *CreateVendorRequest) GetVerifyAccountOk() (*bool, bool)`

GetVerifyAccountOk returns a tuple with the VerifyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAccount

`func (o *CreateVendorRequest) SetVerifyAccount(v bool)`

SetVerifyAccount sets VerifyAccount field to given value.

### HasVerifyAccount

`func (o *CreateVendorRequest) HasVerifyAccount() bool`

HasVerifyAccount returns a boolean if a field has been set.

### GetDashboardAccess

`func (o *CreateVendorRequest) GetDashboardAccess() bool`

GetDashboardAccess returns the DashboardAccess field if non-nil, zero value otherwise.

### GetDashboardAccessOk

`func (o *CreateVendorRequest) GetDashboardAccessOk() (*bool, bool)`

GetDashboardAccessOk returns a tuple with the DashboardAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardAccess

`func (o *CreateVendorRequest) SetDashboardAccess(v bool)`

SetDashboardAccess sets DashboardAccess field to given value.

### HasDashboardAccess

`func (o *CreateVendorRequest) HasDashboardAccess() bool`

HasDashboardAccess returns a boolean if a field has been set.

### GetScheduleOption

`func (o *CreateVendorRequest) GetScheduleOption() float32`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *CreateVendorRequest) GetScheduleOptionOk() (*float32, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *CreateVendorRequest) SetScheduleOption(v float32)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *CreateVendorRequest) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetBank

`func (o *CreateVendorRequest) GetBank() []BankDetails`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *CreateVendorRequest) GetBankOk() (*[]BankDetails, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *CreateVendorRequest) SetBank(v []BankDetails)`

SetBank sets Bank field to given value.

### HasBank

`func (o *CreateVendorRequest) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetUpi

`func (o *CreateVendorRequest) GetUpi() []UpiDetails`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *CreateVendorRequest) GetUpiOk() (*[]UpiDetails, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *CreateVendorRequest) SetUpi(v []UpiDetails)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *CreateVendorRequest) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetKycDetails

`func (o *CreateVendorRequest) GetKycDetails() []KycDetails`

GetKycDetails returns the KycDetails field if non-nil, zero value otherwise.

### GetKycDetailsOk

`func (o *CreateVendorRequest) GetKycDetailsOk() (*[]KycDetails, bool)`

GetKycDetailsOk returns a tuple with the KycDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycDetails

`func (o *CreateVendorRequest) SetKycDetails(v []KycDetails)`

SetKycDetails sets KycDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


