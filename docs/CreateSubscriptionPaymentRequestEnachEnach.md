# CreateSubscriptionPaymentRequestEnachEnach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBankCode** | Pointer to **string** | Account bank code (mandatory). Consists of the first four alphabetic characters of the IFSC. | [optional] 
**AccountHolderName** | Pointer to **string** | Account holder name. | [optional] 
**AccountIfsc** | Pointer to **string** | Account IFSC (optional). Complete 11-character alphanumeric code. | [optional] 
**AccountNumber** | Pointer to **string** | Account number. | [optional] 
**AccountType** | Pointer to **string** | Account type. | [optional] 
**AuthMode** | Pointer to **string** | Authentication mode. can be &#x60;debit_card&#x60;, &#x60;aadhaar&#x60;, or &#x60;net_banking&#x60;. | [optional] 
**Channel** | Pointer to **string** | Channel. can be link. | [optional] 

## Methods

### NewCreateSubscriptionPaymentRequestEnachEnach

`func NewCreateSubscriptionPaymentRequestEnachEnach() *CreateSubscriptionPaymentRequestEnachEnach`

NewCreateSubscriptionPaymentRequestEnachEnach instantiates a new CreateSubscriptionPaymentRequestEnachEnach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentRequestEnachEnachWithDefaults

`func NewCreateSubscriptionPaymentRequestEnachEnachWithDefaults() *CreateSubscriptionPaymentRequestEnachEnach`

NewCreateSubscriptionPaymentRequestEnachEnachWithDefaults instantiates a new CreateSubscriptionPaymentRequestEnachEnach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAuthMode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetChannel

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSubscriptionPaymentRequestEnachEnach) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSubscriptionPaymentRequestEnachEnach) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateSubscriptionPaymentRequestEnachEnach) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


