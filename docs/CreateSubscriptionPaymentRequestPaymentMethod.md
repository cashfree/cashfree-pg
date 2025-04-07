# CreateSubscriptionPaymentRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel. can be link | [optional] 
**UpiId** | Pointer to **string** |  | [optional] 
**AccountBankCode** | Pointer to **string** | Account bank code | [optional] 
**AccountHolderName** | Pointer to **string** | Account holder name | [optional] 
**AccountIfsc** | Pointer to **string** | Account IFSC | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] 
**AccountType** | Pointer to **string** | Account type | [optional] 
**AuthMode** | Pointer to **string** | Authentication mode. can be debit_card, aadhaar, or net_banking | [optional] 
**MandateCreationDate** | Pointer to **string** | Mandate creation date | [optional] 
**MandateStartDate** | Pointer to **string** | Mandate start date | [optional] 
**CardCvv** | Pointer to **string** | Card CVV | [optional] 
**CardExpiryMm** | Pointer to **string** | Card expiry month | [optional] 
**CardExpiryYy** | Pointer to **string** | Card expiry year | [optional] 
**CardHolderName** | Pointer to **string** | Card holder name | [optional] 
**CardNetwork** | Pointer to **string** | Card network | [optional] 
**CardNumber** | Pointer to **string** | Card number | [optional] 
**CardType** | Pointer to **string** | Card type | [optional] 

## Methods

### NewCreateSubscriptionPaymentRequestPaymentMethod

`func NewCreateSubscriptionPaymentRequestPaymentMethod() *CreateSubscriptionPaymentRequestPaymentMethod`

NewCreateSubscriptionPaymentRequestPaymentMethod instantiates a new CreateSubscriptionPaymentRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentRequestPaymentMethodWithDefaults

`func NewCreateSubscriptionPaymentRequestPaymentMethodWithDefaults() *CreateSubscriptionPaymentRequestPaymentMethod`

NewCreateSubscriptionPaymentRequestPaymentMethodWithDefaults instantiates a new CreateSubscriptionPaymentRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetUpiId

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAuthMode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetMandateCreationDate() string`

GetMandateCreationDate returns the MandateCreationDate field if non-nil, zero value otherwise.

### GetMandateCreationDateOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetMandateCreationDateOk() (*string, bool)`

GetMandateCreationDateOk returns a tuple with the MandateCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetMandateCreationDate(v string)`

SetMandateCreationDate sets MandateCreationDate field to given value.

### HasMandateCreationDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasMandateCreationDate() bool`

HasMandateCreationDate returns a boolean if a field has been set.

### GetMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetMandateStartDate() string`

GetMandateStartDate returns the MandateStartDate field if non-nil, zero value otherwise.

### GetMandateStartDateOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetMandateStartDateOk() (*string, bool)`

GetMandateStartDateOk returns a tuple with the MandateStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetMandateStartDate(v string)`

SetMandateStartDate sets MandateStartDate field to given value.

### HasMandateStartDate

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasMandateStartDate() bool`

HasMandateStartDate returns a boolean if a field has been set.

### GetCardCvv

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.

### HasCardCvv

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardCvv() bool`

HasCardCvv returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.

### HasCardExpiryMm

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardExpiryMm() bool`

HasCardExpiryMm returns a boolean if a field has been set.

### GetCardExpiryYy

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.

### HasCardExpiryYy

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardExpiryYy() bool`

HasCardExpiryYy returns a boolean if a field has been set.

### GetCardHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetCardNetwork

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CreateSubscriptionPaymentRequestPaymentMethod) HasCardType() bool`

HasCardType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


