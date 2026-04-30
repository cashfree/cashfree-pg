# CreateSubscriptionPaymentRequestPnachPnach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBankCode** | Pointer to **string** | Account bank code. Consists of the first four alphabetic characters of the IFSC. | [optional] 
**AccountHolderName** | Pointer to **string** | Account holder name. | [optional] 
**AccountIfsc** | Pointer to **string** | Account IFSC. Complete 11-character alphanumeric code. | [optional] 
**AccountNumber** | Pointer to **string** | Account number. | [optional] 
**AccountType** | Pointer to **string** | Account type. | [optional] 
**Channel** | Pointer to **string** | Channel. can be post. | [optional] 
**MandateCreationDate** | Pointer to **string** | Mandate creation date. | [optional] 
**MandateStartDate** | Pointer to **string** | Mandate start date. | [optional] 

## Methods

### NewCreateSubscriptionPaymentRequestPnachPnach

`func NewCreateSubscriptionPaymentRequestPnachPnach() *CreateSubscriptionPaymentRequestPnachPnach`

NewCreateSubscriptionPaymentRequestPnachPnach instantiates a new CreateSubscriptionPaymentRequestPnachPnach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentRequestPnachPnachWithDefaults

`func NewCreateSubscriptionPaymentRequestPnachPnachWithDefaults() *CreateSubscriptionPaymentRequestPnachPnach`

NewCreateSubscriptionPaymentRequestPnachPnachWithDefaults instantiates a new CreateSubscriptionPaymentRequestPnachPnach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetChannel

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetMandateCreationDate() string`

GetMandateCreationDate returns the MandateCreationDate field if non-nil, zero value otherwise.

### GetMandateCreationDateOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetMandateCreationDateOk() (*string, bool)`

GetMandateCreationDateOk returns a tuple with the MandateCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetMandateCreationDate(v string)`

SetMandateCreationDate sets MandateCreationDate field to given value.

### HasMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasMandateCreationDate() bool`

HasMandateCreationDate returns a boolean if a field has been set.

### GetMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetMandateStartDate() string`

GetMandateStartDate returns the MandateStartDate field if non-nil, zero value otherwise.

### GetMandateStartDateOk

`func (o *CreateSubscriptionPaymentRequestPnachPnach) GetMandateStartDateOk() (*string, bool)`

GetMandateStartDateOk returns a tuple with the MandateStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) SetMandateStartDate(v string)`

SetMandateStartDate sets MandateStartDate field to given value.

### HasMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPnachPnach) HasMandateStartDate() bool`

HasMandateStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


