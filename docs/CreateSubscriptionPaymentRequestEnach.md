# CreateSubscriptionPaymentRequestEnach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBankCode** | Pointer to **string** | Account bank code (required without AccountIFSC) | [optional] 
**AccountHolderName** | Pointer to **string** | Account holder name | [optional] 
**AccountIfsc** | Pointer to **string** | Account IFSC | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] 
**AccountType** | Pointer to **string** | Account type | [optional] 
**AuthMode** | Pointer to **string** | Authentication mode. can be debit_card, aadhaar, or net_banking | [optional] 
**Channel** | Pointer to **string** | Channel. can be link | [optional] 

## Methods

### NewCreateSubscriptionPaymentRequestEnach

`func NewCreateSubscriptionPaymentRequestEnach() *CreateSubscriptionPaymentRequestEnach`

NewCreateSubscriptionPaymentRequestEnach instantiates a new CreateSubscriptionPaymentRequestEnach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentRequestEnachWithDefaults

`func NewCreateSubscriptionPaymentRequestEnachWithDefaults() *CreateSubscriptionPaymentRequestEnach`

NewCreateSubscriptionPaymentRequestEnachWithDefaults instantiates a new CreateSubscriptionPaymentRequestEnach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnach) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *CreateSubscriptionPaymentRequestEnach) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnach) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *CreateSubscriptionPaymentRequestEnach) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnach) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *CreateSubscriptionPaymentRequestEnach) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnach) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreateSubscriptionPaymentRequestEnach) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreateSubscriptionPaymentRequestEnach) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *CreateSubscriptionPaymentRequestEnach) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAuthMode

`func (o *CreateSubscriptionPaymentRequestEnach) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *CreateSubscriptionPaymentRequestEnach) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *CreateSubscriptionPaymentRequestEnach) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetChannel

`func (o *CreateSubscriptionPaymentRequestEnach) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSubscriptionPaymentRequestEnach) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSubscriptionPaymentRequestEnach) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateSubscriptionPaymentRequestEnach) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


