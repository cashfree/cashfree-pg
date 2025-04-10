# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for card payments can be \&quot;link\&quot; or \&quot;post\&quot;. Post is used for seamless OTP payments where merchant captures OTP on their own page. | 
**CardNumber** | Pointer to **string** | Customer card number for plain card transactions. Token pan number for tokenized card transactions. | [optional] 
**CardHolderName** | Pointer to **string** | Customer name mentioned on the card. | [optional] 
**CardExpiryMm** | Pointer to **string** | Card expiry month for plain card transactions. Token expiry month for tokenized card transactions. | [optional] 
**CardExpiryYy** | Pointer to **string** | Card expiry year for plain card transactions. Token expiry year for tokenized card transactions. | [optional] 
**CardCvv** | Pointer to **string** | CVV mentioned on the card. | [optional] 
**InstrumentId** | Pointer to **string** | instrument id of saved card. Required only to make payment using saved instrument. | [optional] 
**Cryptogram** | Pointer to **string** | cryptogram received from card network. Required only for tokenized card transactions. | [optional] 
**TokenRequestorId** | Pointer to **string** | TRID issued by card networks. Required only for tokenized card transactions. | [optional] 
**TokenReferenceId** | Pointer to **string** | Token Reference Id provided by Diners for Guest Checkout Token.  Required only for Diners cards. | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**CardDisplay** | Pointer to **string** | last 4 digits of original card number. Required only for tokenized card transactions. | [optional] 
**CardAlias** | Pointer to **string** | Card alias as returned by Cashfree Vault API. | [optional] 
**CardBankName** | Pointer to **string** | One of [\&quot;Kotak\&quot;, \&quot;ICICI\&quot;, \&quot;RBL\&quot;, \&quot;BOB\&quot;, \&quot;Standard Chartered\&quot;]. Card bank name, required for EMI payments. This is the bank user has selected for EMI | [optional] 
**AddressLineOne** | Pointer to **string** | First line of the address. | [optional] 
**AddressLineTwo** | Pointer to **string** | Second line of the address. | [optional] 
**City** | Pointer to **string** | City Name. | [optional] 
**ZipCode** | Pointer to **string** | Pin Code/Zip Code. | [optional] 
**Country** | Pointer to **string** | Country Name. | [optional] 
**CountryCode** | Pointer to **string** | Country Code. Should be in ISO 2 format (ie. US for United States) | [optional] 
**State** | Pointer to **string** | State Name. | [optional] 
**StateCode** | Pointer to **string** | State Code. Should be in ISO 2 format (ie. FL for Florida) | [optional] 
**EmiTenure** | Pointer to **int32** | EMI tenure selected by the user | [optional] 

## Methods

### NewCard

`func NewCard(channel string, ) *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Card) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Card) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Card) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCardNumber

`func (o *Card) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *Card) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *Card) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *Card) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardHolderName

`func (o *Card) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *Card) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *Card) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *Card) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *Card) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *Card) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *Card) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.

### HasCardExpiryMm

`func (o *Card) HasCardExpiryMm() bool`

HasCardExpiryMm returns a boolean if a field has been set.

### GetCardExpiryYy

`func (o *Card) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *Card) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *Card) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.

### HasCardExpiryYy

`func (o *Card) HasCardExpiryYy() bool`

HasCardExpiryYy returns a boolean if a field has been set.

### GetCardCvv

`func (o *Card) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *Card) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *Card) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.

### HasCardCvv

`func (o *Card) HasCardCvv() bool`

HasCardCvv returns a boolean if a field has been set.

### GetInstrumentId

`func (o *Card) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Card) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Card) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Card) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetCryptogram

`func (o *Card) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *Card) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *Card) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *Card) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### GetTokenRequestorId

`func (o *Card) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *Card) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *Card) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.

### HasTokenRequestorId

`func (o *Card) HasTokenRequestorId() bool`

HasTokenRequestorId returns a boolean if a field has been set.

### GetTokenReferenceId

`func (o *Card) GetTokenReferenceId() string`

GetTokenReferenceId returns the TokenReferenceId field if non-nil, zero value otherwise.

### GetTokenReferenceIdOk

`func (o *Card) GetTokenReferenceIdOk() (*string, bool)`

GetTokenReferenceIdOk returns a tuple with the TokenReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReferenceId

`func (o *Card) SetTokenReferenceId(v string)`

SetTokenReferenceId sets TokenReferenceId field to given value.

### HasTokenReferenceId

`func (o *Card) HasTokenReferenceId() bool`

HasTokenReferenceId returns a boolean if a field has been set.

### GetTokenType

`func (o *Card) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Card) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Card) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Card) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetCardDisplay

`func (o *Card) GetCardDisplay() string`

GetCardDisplay returns the CardDisplay field if non-nil, zero value otherwise.

### GetCardDisplayOk

`func (o *Card) GetCardDisplayOk() (*string, bool)`

GetCardDisplayOk returns a tuple with the CardDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDisplay

`func (o *Card) SetCardDisplay(v string)`

SetCardDisplay sets CardDisplay field to given value.

### HasCardDisplay

`func (o *Card) HasCardDisplay() bool`

HasCardDisplay returns a boolean if a field has been set.

### GetCardAlias

`func (o *Card) GetCardAlias() string`

GetCardAlias returns the CardAlias field if non-nil, zero value otherwise.

### GetCardAliasOk

`func (o *Card) GetCardAliasOk() (*string, bool)`

GetCardAliasOk returns a tuple with the CardAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAlias

`func (o *Card) SetCardAlias(v string)`

SetCardAlias sets CardAlias field to given value.

### HasCardAlias

`func (o *Card) HasCardAlias() bool`

HasCardAlias returns a boolean if a field has been set.

### GetCardBankName

`func (o *Card) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *Card) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *Card) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *Card) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetAddressLineOne

`func (o *Card) GetAddressLineOne() string`

GetAddressLineOne returns the AddressLineOne field if non-nil, zero value otherwise.

### GetAddressLineOneOk

`func (o *Card) GetAddressLineOneOk() (*string, bool)`

GetAddressLineOneOk returns a tuple with the AddressLineOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLineOne

`func (o *Card) SetAddressLineOne(v string)`

SetAddressLineOne sets AddressLineOne field to given value.

### HasAddressLineOne

`func (o *Card) HasAddressLineOne() bool`

HasAddressLineOne returns a boolean if a field has been set.

### GetAddressLineTwo

`func (o *Card) GetAddressLineTwo() string`

GetAddressLineTwo returns the AddressLineTwo field if non-nil, zero value otherwise.

### GetAddressLineTwoOk

`func (o *Card) GetAddressLineTwoOk() (*string, bool)`

GetAddressLineTwoOk returns a tuple with the AddressLineTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLineTwo

`func (o *Card) SetAddressLineTwo(v string)`

SetAddressLineTwo sets AddressLineTwo field to given value.

### HasAddressLineTwo

`func (o *Card) HasAddressLineTwo() bool`

HasAddressLineTwo returns a boolean if a field has been set.

### GetCity

`func (o *Card) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Card) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Card) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Card) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *Card) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Card) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Card) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Card) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *Card) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Card) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Card) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Card) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *Card) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Card) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Card) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Card) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetState

`func (o *Card) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Card) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Card) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Card) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateCode

`func (o *Card) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *Card) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *Card) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *Card) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetEmiTenure

`func (o *Card) GetEmiTenure() int32`

GetEmiTenure returns the EmiTenure field if non-nil, zero value otherwise.

### GetEmiTenureOk

`func (o *Card) GetEmiTenureOk() (*int32, bool)`

GetEmiTenureOk returns a tuple with the EmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiTenure

`func (o *Card) SetEmiTenure(v int32)`

SetEmiTenure sets EmiTenure field to given value.

### HasEmiTenure

`func (o *Card) HasEmiTenure() bool`

HasEmiTenure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


