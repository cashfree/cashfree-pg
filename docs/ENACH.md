# ENACH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel. can be link | [optional] 
**AuthMode** | Pointer to **string** | Authentication mode. can be debit_card, aadhaar, or net_banking | [optional] 
**AccountHolderName** | Pointer to **string** | Account holder name | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] 
**AccountBankCode** | Pointer to **string** | Account bank code (required without AccountIFSC) | [optional] 
**AccountType** | Pointer to **string** | Account type | [optional] 
**AccountIfsc** | Pointer to **string** | Account IFSC | [optional] 

## Methods

### NewENACH

`func NewENACH() *ENACH`

NewENACH instantiates a new ENACH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewENACHWithDefaults

`func NewENACHWithDefaults() *ENACH`

NewENACHWithDefaults instantiates a new ENACH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ENACH) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ENACH) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ENACH) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ENACH) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetAuthMode

`func (o *ENACH) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ENACH) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ENACH) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *ENACH) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *ENACH) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *ENACH) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *ENACH) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *ENACH) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *ENACH) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ENACH) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ENACH) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *ENACH) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountBankCode

`func (o *ENACH) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *ENACH) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *ENACH) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *ENACH) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetAccountType

`func (o *ENACH) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ENACH) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ENACH) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ENACH) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *ENACH) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *ENACH) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *ENACH) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *ENACH) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


