# SubscriptionCustomerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | Pointer to **string** | Name of the customer. | [optional] 
**CustomerEmail** | **string** | Email of the customer. | 
**CustomerPhone** | **string** | Phone number of the customer. | 
**CustomerBankAccountHolderName** | Pointer to **string** | Bank holder name of the customer. | [optional] 
**CustomerBankAccountNumber** | Pointer to **string** | Bank account number of the customer. | [optional] 
**CustomerBankIfsc** | Pointer to **string** | IFSC code of the customer. | [optional] 
**CustomerBankCode** | Pointer to **string** | Bank code of the customer. Refer to https://www.npci.org.in/PDF/nach/live-members-e-mandates/Live-Banks-in-API-E-Mandate.pdf | [optional] 
**CustomerBankAccountType** | Pointer to **string** | Bank account type of the customer. | [optional] 

## Methods

### NewSubscriptionCustomerDetails

`func NewSubscriptionCustomerDetails(customerEmail string, customerPhone string, ) *SubscriptionCustomerDetails`

NewSubscriptionCustomerDetails instantiates a new SubscriptionCustomerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCustomerDetailsWithDefaults

`func NewSubscriptionCustomerDetailsWithDefaults() *SubscriptionCustomerDetails`

NewSubscriptionCustomerDetailsWithDefaults instantiates a new SubscriptionCustomerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *SubscriptionCustomerDetails) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *SubscriptionCustomerDetails) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *SubscriptionCustomerDetails) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *SubscriptionCustomerDetails) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *SubscriptionCustomerDetails) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *SubscriptionCustomerDetails) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *SubscriptionCustomerDetails) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### GetCustomerPhone

`func (o *SubscriptionCustomerDetails) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *SubscriptionCustomerDetails) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *SubscriptionCustomerDetails) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerBankAccountHolderName

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountHolderName() string`

GetCustomerBankAccountHolderName returns the CustomerBankAccountHolderName field if non-nil, zero value otherwise.

### GetCustomerBankAccountHolderNameOk

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountHolderNameOk() (*string, bool)`

GetCustomerBankAccountHolderNameOk returns a tuple with the CustomerBankAccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountHolderName

`func (o *SubscriptionCustomerDetails) SetCustomerBankAccountHolderName(v string)`

SetCustomerBankAccountHolderName sets CustomerBankAccountHolderName field to given value.

### HasCustomerBankAccountHolderName

`func (o *SubscriptionCustomerDetails) HasCustomerBankAccountHolderName() bool`

HasCustomerBankAccountHolderName returns a boolean if a field has been set.

### GetCustomerBankAccountNumber

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountNumber() string`

GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field if non-nil, zero value otherwise.

### GetCustomerBankAccountNumberOk

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountNumberOk() (*string, bool)`

GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountNumber

`func (o *SubscriptionCustomerDetails) SetCustomerBankAccountNumber(v string)`

SetCustomerBankAccountNumber sets CustomerBankAccountNumber field to given value.

### HasCustomerBankAccountNumber

`func (o *SubscriptionCustomerDetails) HasCustomerBankAccountNumber() bool`

HasCustomerBankAccountNumber returns a boolean if a field has been set.

### GetCustomerBankIfsc

`func (o *SubscriptionCustomerDetails) GetCustomerBankIfsc() string`

GetCustomerBankIfsc returns the CustomerBankIfsc field if non-nil, zero value otherwise.

### GetCustomerBankIfscOk

`func (o *SubscriptionCustomerDetails) GetCustomerBankIfscOk() (*string, bool)`

GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankIfsc

`func (o *SubscriptionCustomerDetails) SetCustomerBankIfsc(v string)`

SetCustomerBankIfsc sets CustomerBankIfsc field to given value.

### HasCustomerBankIfsc

`func (o *SubscriptionCustomerDetails) HasCustomerBankIfsc() bool`

HasCustomerBankIfsc returns a boolean if a field has been set.

### GetCustomerBankCode

`func (o *SubscriptionCustomerDetails) GetCustomerBankCode() string`

GetCustomerBankCode returns the CustomerBankCode field if non-nil, zero value otherwise.

### GetCustomerBankCodeOk

`func (o *SubscriptionCustomerDetails) GetCustomerBankCodeOk() (*string, bool)`

GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankCode

`func (o *SubscriptionCustomerDetails) SetCustomerBankCode(v string)`

SetCustomerBankCode sets CustomerBankCode field to given value.

### HasCustomerBankCode

`func (o *SubscriptionCustomerDetails) HasCustomerBankCode() bool`

HasCustomerBankCode returns a boolean if a field has been set.

### GetCustomerBankAccountType

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountType() string`

GetCustomerBankAccountType returns the CustomerBankAccountType field if non-nil, zero value otherwise.

### GetCustomerBankAccountTypeOk

`func (o *SubscriptionCustomerDetails) GetCustomerBankAccountTypeOk() (*string, bool)`

GetCustomerBankAccountTypeOk returns a tuple with the CustomerBankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountType

`func (o *SubscriptionCustomerDetails) SetCustomerBankAccountType(v string)`

SetCustomerBankAccountType sets CustomerBankAccountType field to given value.

### HasCustomerBankAccountType

`func (o *SubscriptionCustomerDetails) HasCustomerBankAccountType() bool`

HasCustomerBankAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


