# AuthorizationDetailsPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **NullableString** | Can be link. | [optional] 
**UpiId** | Pointer to **string** |  | [optional] 
**UpiInstrument** | Pointer to **string** |  | [optional] 
**UpiInstrumentNumber** | Pointer to **string** |  | [optional] 
**UpiPayerAccountNumber** | Pointer to **string** |  | [optional] 
**UpiPayerIfsc** | Pointer to **string** |  | [optional] 
**AuthMode** | Pointer to **string** | Auth mode. Can be debit_card, aadhaar, or net_banking. | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**AccountIfsc** | Pointer to **string** |  | [optional] 
**AccountHolderName** | Pointer to **string** |  | [optional] 
**AccountBankCode** | Pointer to **string** |  | [optional] 
**MandateCreationDate** | Pointer to **string** |  | [optional] 
**MandateStartDate** | Pointer to **string** |  | [optional] 
**CardNumber** | Pointer to **string** | Card number | [optional] 
**CardNetwork** | Pointer to **string** | Card network | [optional] 
**CardType** | Pointer to **string** | Card type (e.g., credit_card) | [optional] 
**CardSubType** | Pointer to **string** | Card subtype (e.g., R) | [optional] 
**CardCountry** | Pointer to **string** | Country of the card (e.g., IN) | [optional] 
**CardBankName** | Pointer to **string** | Bank name on card | [optional] 
**CardNetworkReferenceId** | Pointer to **NullableString** | Network reference ID | [optional] 
**InstrumentId** | Pointer to **string** | Unique identifier for card instrument | [optional] 

## Methods

### NewAuthorizationDetailsPaymentMethod

`func NewAuthorizationDetailsPaymentMethod() *AuthorizationDetailsPaymentMethod`

NewAuthorizationDetailsPaymentMethod instantiates a new AuthorizationDetailsPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailsPaymentMethodWithDefaults

`func NewAuthorizationDetailsPaymentMethodWithDefaults() *AuthorizationDetailsPaymentMethod`

NewAuthorizationDetailsPaymentMethodWithDefaults instantiates a new AuthorizationDetailsPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *AuthorizationDetailsPaymentMethod) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AuthorizationDetailsPaymentMethod) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AuthorizationDetailsPaymentMethod) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AuthorizationDetailsPaymentMethod) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *AuthorizationDetailsPaymentMethod) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *AuthorizationDetailsPaymentMethod) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetUpiId

`func (o *AuthorizationDetailsPaymentMethod) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *AuthorizationDetailsPaymentMethod) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *AuthorizationDetailsPaymentMethod) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetUpiInstrument

`func (o *AuthorizationDetailsPaymentMethod) GetUpiInstrument() string`

GetUpiInstrument returns the UpiInstrument field if non-nil, zero value otherwise.

### GetUpiInstrumentOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiInstrumentOk() (*string, bool)`

GetUpiInstrumentOk returns a tuple with the UpiInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiInstrument

`func (o *AuthorizationDetailsPaymentMethod) SetUpiInstrument(v string)`

SetUpiInstrument sets UpiInstrument field to given value.

### HasUpiInstrument

`func (o *AuthorizationDetailsPaymentMethod) HasUpiInstrument() bool`

HasUpiInstrument returns a boolean if a field has been set.

### GetUpiInstrumentNumber

`func (o *AuthorizationDetailsPaymentMethod) GetUpiInstrumentNumber() string`

GetUpiInstrumentNumber returns the UpiInstrumentNumber field if non-nil, zero value otherwise.

### GetUpiInstrumentNumberOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiInstrumentNumberOk() (*string, bool)`

GetUpiInstrumentNumberOk returns a tuple with the UpiInstrumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiInstrumentNumber

`func (o *AuthorizationDetailsPaymentMethod) SetUpiInstrumentNumber(v string)`

SetUpiInstrumentNumber sets UpiInstrumentNumber field to given value.

### HasUpiInstrumentNumber

`func (o *AuthorizationDetailsPaymentMethod) HasUpiInstrumentNumber() bool`

HasUpiInstrumentNumber returns a boolean if a field has been set.

### GetUpiPayerAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) GetUpiPayerAccountNumber() string`

GetUpiPayerAccountNumber returns the UpiPayerAccountNumber field if non-nil, zero value otherwise.

### GetUpiPayerAccountNumberOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiPayerAccountNumberOk() (*string, bool)`

GetUpiPayerAccountNumberOk returns a tuple with the UpiPayerAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiPayerAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) SetUpiPayerAccountNumber(v string)`

SetUpiPayerAccountNumber sets UpiPayerAccountNumber field to given value.

### HasUpiPayerAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) HasUpiPayerAccountNumber() bool`

HasUpiPayerAccountNumber returns a boolean if a field has been set.

### GetUpiPayerIfsc

`func (o *AuthorizationDetailsPaymentMethod) GetUpiPayerIfsc() string`

GetUpiPayerIfsc returns the UpiPayerIfsc field if non-nil, zero value otherwise.

### GetUpiPayerIfscOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiPayerIfscOk() (*string, bool)`

GetUpiPayerIfscOk returns a tuple with the UpiPayerIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiPayerIfsc

`func (o *AuthorizationDetailsPaymentMethod) SetUpiPayerIfsc(v string)`

SetUpiPayerIfsc sets UpiPayerIfsc field to given value.

### HasUpiPayerIfsc

`func (o *AuthorizationDetailsPaymentMethod) HasUpiPayerIfsc() bool`

HasUpiPayerIfsc returns a boolean if a field has been set.

### GetAuthMode

`func (o *AuthorizationDetailsPaymentMethod) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *AuthorizationDetailsPaymentMethod) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *AuthorizationDetailsPaymentMethod) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *AuthorizationDetailsPaymentMethod) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAccountType

`func (o *AuthorizationDetailsPaymentMethod) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AuthorizationDetailsPaymentMethod) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AuthorizationDetailsPaymentMethod) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AuthorizationDetailsPaymentMethod) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AuthorizationDetailsPaymentMethod) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AuthorizationDetailsPaymentMethod) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountIfsc

`func (o *AuthorizationDetailsPaymentMethod) GetAccountIfsc() string`

GetAccountIfsc returns the AccountIfsc field if non-nil, zero value otherwise.

### GetAccountIfscOk

`func (o *AuthorizationDetailsPaymentMethod) GetAccountIfscOk() (*string, bool)`

GetAccountIfscOk returns a tuple with the AccountIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIfsc

`func (o *AuthorizationDetailsPaymentMethod) SetAccountIfsc(v string)`

SetAccountIfsc sets AccountIfsc field to given value.

### HasAccountIfsc

`func (o *AuthorizationDetailsPaymentMethod) HasAccountIfsc() bool`

HasAccountIfsc returns a boolean if a field has been set.

### GetAccountHolderName

`func (o *AuthorizationDetailsPaymentMethod) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *AuthorizationDetailsPaymentMethod) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *AuthorizationDetailsPaymentMethod) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.

### HasAccountHolderName

`func (o *AuthorizationDetailsPaymentMethod) HasAccountHolderName() bool`

HasAccountHolderName returns a boolean if a field has been set.

### GetAccountBankCode

`func (o *AuthorizationDetailsPaymentMethod) GetAccountBankCode() string`

GetAccountBankCode returns the AccountBankCode field if non-nil, zero value otherwise.

### GetAccountBankCodeOk

`func (o *AuthorizationDetailsPaymentMethod) GetAccountBankCodeOk() (*string, bool)`

GetAccountBankCodeOk returns a tuple with the AccountBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBankCode

`func (o *AuthorizationDetailsPaymentMethod) SetAccountBankCode(v string)`

SetAccountBankCode sets AccountBankCode field to given value.

### HasAccountBankCode

`func (o *AuthorizationDetailsPaymentMethod) HasAccountBankCode() bool`

HasAccountBankCode returns a boolean if a field has been set.

### GetMandateCreationDate

`func (o *AuthorizationDetailsPaymentMethod) GetMandateCreationDate() string`

GetMandateCreationDate returns the MandateCreationDate field if non-nil, zero value otherwise.

### GetMandateCreationDateOk

`func (o *AuthorizationDetailsPaymentMethod) GetMandateCreationDateOk() (*string, bool)`

GetMandateCreationDateOk returns a tuple with the MandateCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateCreationDate

`func (o *AuthorizationDetailsPaymentMethod) SetMandateCreationDate(v string)`

SetMandateCreationDate sets MandateCreationDate field to given value.

### HasMandateCreationDate

`func (o *AuthorizationDetailsPaymentMethod) HasMandateCreationDate() bool`

HasMandateCreationDate returns a boolean if a field has been set.

### GetMandateStartDate

`func (o *AuthorizationDetailsPaymentMethod) GetMandateStartDate() string`

GetMandateStartDate returns the MandateStartDate field if non-nil, zero value otherwise.

### GetMandateStartDateOk

`func (o *AuthorizationDetailsPaymentMethod) GetMandateStartDateOk() (*string, bool)`

GetMandateStartDateOk returns a tuple with the MandateStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStartDate

`func (o *AuthorizationDetailsPaymentMethod) SetMandateStartDate(v string)`

SetMandateStartDate sets MandateStartDate field to given value.

### HasMandateStartDate

`func (o *AuthorizationDetailsPaymentMethod) HasMandateStartDate() bool`

HasMandateStartDate returns a boolean if a field has been set.

### GetCardNumber

`func (o *AuthorizationDetailsPaymentMethod) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *AuthorizationDetailsPaymentMethod) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *AuthorizationDetailsPaymentMethod) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNetwork

`func (o *AuthorizationDetailsPaymentMethod) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *AuthorizationDetailsPaymentMethod) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *AuthorizationDetailsPaymentMethod) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardType

`func (o *AuthorizationDetailsPaymentMethod) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *AuthorizationDetailsPaymentMethod) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *AuthorizationDetailsPaymentMethod) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardSubType

`func (o *AuthorizationDetailsPaymentMethod) GetCardSubType() string`

GetCardSubType returns the CardSubType field if non-nil, zero value otherwise.

### GetCardSubTypeOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardSubTypeOk() (*string, bool)`

GetCardSubTypeOk returns a tuple with the CardSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSubType

`func (o *AuthorizationDetailsPaymentMethod) SetCardSubType(v string)`

SetCardSubType sets CardSubType field to given value.

### HasCardSubType

`func (o *AuthorizationDetailsPaymentMethod) HasCardSubType() bool`

HasCardSubType returns a boolean if a field has been set.

### GetCardCountry

`func (o *AuthorizationDetailsPaymentMethod) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *AuthorizationDetailsPaymentMethod) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *AuthorizationDetailsPaymentMethod) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardBankName

`func (o *AuthorizationDetailsPaymentMethod) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *AuthorizationDetailsPaymentMethod) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *AuthorizationDetailsPaymentMethod) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetCardNetworkReferenceId

`func (o *AuthorizationDetailsPaymentMethod) GetCardNetworkReferenceId() string`

GetCardNetworkReferenceId returns the CardNetworkReferenceId field if non-nil, zero value otherwise.

### GetCardNetworkReferenceIdOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardNetworkReferenceIdOk() (*string, bool)`

GetCardNetworkReferenceIdOk returns a tuple with the CardNetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetworkReferenceId

`func (o *AuthorizationDetailsPaymentMethod) SetCardNetworkReferenceId(v string)`

SetCardNetworkReferenceId sets CardNetworkReferenceId field to given value.

### HasCardNetworkReferenceId

`func (o *AuthorizationDetailsPaymentMethod) HasCardNetworkReferenceId() bool`

HasCardNetworkReferenceId returns a boolean if a field has been set.

### SetCardNetworkReferenceIdNil

`func (o *AuthorizationDetailsPaymentMethod) SetCardNetworkReferenceIdNil(b bool)`

 SetCardNetworkReferenceIdNil sets the value for CardNetworkReferenceId to be an explicit nil

### UnsetCardNetworkReferenceId
`func (o *AuthorizationDetailsPaymentMethod) UnsetCardNetworkReferenceId()`

UnsetCardNetworkReferenceId ensures that no value is present for CardNetworkReferenceId, not even an explicit nil
### GetInstrumentId

`func (o *AuthorizationDetailsPaymentMethod) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *AuthorizationDetailsPaymentMethod) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *AuthorizationDetailsPaymentMethod) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *AuthorizationDetailsPaymentMethod) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


