# AuthResponsePnach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel. Can be post. | [optional] 
**MandateCreationDate** | Pointer to **string** | Date on which the mandate was created. Cashfree stores timestamps in IST. | [optional] 
**MandateStartDate** | Pointer to **string** | Start date of the mandate. Cashfree stores timestamps in IST. | [optional] 
**AccountType** | Pointer to **string** | Type of the bank account. | [optional] 
**AccountNumber** | Pointer to **string** | Bank account number. | [optional] 
**AccountIfsc** | Pointer to **string** | IFSC code of the bank account. | [optional] 
**AccountHolderName** | Pointer to **string** | Name of the account holder. | [optional] 
**AccountBankCode** | Pointer to **string** | Bank code of the account-holding bank. | [optional] 

## Methods

### NewAuthResponsePnach

`func NewAuthResponsePnach() *AuthResponsePnach`

NewAuthResponsePnach instantiates a new AuthResponsePnach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponsePnachWithDefaults

`func NewAuthResponsePnachWithDefaults() *AuthResponsePnach`

NewAuthResponsePnachWithDefaults instantiates a new AuthResponsePnach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *AuthResponsePnach) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AuthResponsePnach) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AuthResponsePnach) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AuthResponsePnach) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMandateCreationDate

`func (o *AuthResponsePnach) GetMandateCreationDate() string`

GetMandateCreationDate returns the MandateCreationDate field if non-nil, zero value otherwise.

### GetMandateCreationDateOk

`func (o *AuthResponsePnach) GetMandateCreationDateOk() (*string, bool)`

GetMandateCreationDateOk returns a tuple with the MandateCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateCreationDate

`func (o *AuthResponsePnach) SetMandateCreationDate(v string)`

SetMandateCreationDate sets MandateCreationDate field to given value.

### HasMandateCreationDate

`func (o *AuthResponsePnach) HasMandateCreationDate() bool`

HasMandateCreationDate returns a boolean if a field has been set.

### GetMandateStartDate

`func (o *AuthResponsePnach) GetMandateStartDate() string`

GetMandateStartDate returns the MandateStartDate field if non-nil, zero value otherwise.

### GetMandateStartDateOk

`func (o *AuthResponsePnach) GetMandateStartDateOk() (*string, bool)`

GetMandateStartDateOk returns a tuple with the MandateStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStartDate

`func (o *AuthResponsePnach) SetMandateStartDate(v string)`

SetMandateStartDate sets MandateStartDate field to given value.

### HasMandateStartDate

`func (o *AuthResponsePnach) HasMandateStartDate() bool`

HasMandateStartDate returns a boolean if a field has been set.

### GetAccountType

`func (o *AuthResponsePnach) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AuthResponsePnach) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AuthResponsePnach) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AuthResponsePnach) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AuthResponsePnach) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AuthResponsePnach) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AuthResponsePnach) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AuthResponsePnach) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *AuthResponsePnach) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *AuthResponsePnach) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *AuthResponsePnach) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *AuthResponsePnach) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *AuthResponsePnach) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *AuthResponsePnach) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *AuthResponsePnach) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *AuthResponsePnach) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountBankCode

`func (o *AuthResponsePnach) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *AuthResponsePnach) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *AuthResponsePnach) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *AuthResponsePnach) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


