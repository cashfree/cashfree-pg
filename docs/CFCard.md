# CFCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for card payments will always be \&quot;link\&quot; | 
**CardNumber** | **string** | Customer card number for plain card transactions. Token pan number for tokenized card transactions. | 
**CardHolderName** | **string** | Customer name mentioned on the card. | 
**CardExpiryMm** | **string** | Card expiry month for plain card transactions. Token expiry month for tokenized card transactions. | 
**CardExpiryYy** | **string** | Card expiry year for plain card transactions. Token expiry year for tokenized card transactions. | 
**CardCvv** | **string** | CVV mentioned on the card. | 
**InstrumentId** | **string** | instrument id of saved card. Required only to make payment using saved instrument. | 
**Cryptogram** | **string** | cryptogram received from card network. Required only for tokenized card transactions. | 
**TokenRequestorId** | **string** | TRID issued by card networks. Required only for tokenized card transactions. | 
**CardDisplay** | **string** | last 4 digits of original card number. Required only for tokenized card transactions. | 
**CardAlias** | **string** | Card alias as returned by Cashfree Vault API. | 
**CardBankName** | **string** | One of [\&quot;Kotak\&quot;, \&quot;ICICI\&quot;, \&quot;RBL\&quot;, \&quot;BOB\&quot;, \&quot;Standard Chartered\&quot;]. Card bank name, required for EMI payments. This is the bank user has selected for EMI | 
**EmiTenure** | **int32** | EMI tenure selected by the user | 

## Methods

### NewCFCard

`func NewCFCard(channel string, cardNumber string, cardHolderName string, cardExpiryMm string, cardExpiryYy string, cardCvv string, instrumentId string, cryptogram string, tokenRequestorId string, cardDisplay string, cardAlias string, cardBankName string, emiTenure int32, ) *CFCard`

NewCFCard instantiates a new CFCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCardWithDefaults

`func NewCFCardWithDefaults() *CFCard`

NewCFCardWithDefaults instantiates a new CFCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFCard) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFCard) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFCard) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCardNumber

`func (o *CFCard) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CFCard) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CFCard) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardHolderName

`func (o *CFCard) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *CFCard) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *CFCard) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.


### GetCardExpiryMm

`func (o *CFCard) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CFCard) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CFCard) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.


### GetCardExpiryYy

`func (o *CFCard) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CFCard) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CFCard) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.


### GetCardCvv

`func (o *CFCard) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *CFCard) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *CFCard) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.


### GetInstrumentId

`func (o *CFCard) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *CFCard) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *CFCard) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetCryptogram

`func (o *CFCard) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *CFCard) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *CFCard) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.


### GetTokenRequestorId

`func (o *CFCard) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *CFCard) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *CFCard) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.


### GetCardDisplay

`func (o *CFCard) GetCardDisplay() string`

GetCardDisplay returns the CardDisplay field if non-nil, zero value otherwise.

### GetCardDisplayOk

`func (o *CFCard) GetCardDisplayOk() (*string, bool)`

GetCardDisplayOk returns a tuple with the CardDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDisplay

`func (o *CFCard) SetCardDisplay(v string)`

SetCardDisplay sets CardDisplay field to given value.


### GetCardAlias

`func (o *CFCard) GetCardAlias() string`

GetCardAlias returns the CardAlias field if non-nil, zero value otherwise.

### GetCardAliasOk

`func (o *CFCard) GetCardAliasOk() (*string, bool)`

GetCardAliasOk returns a tuple with the CardAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAlias

`func (o *CFCard) SetCardAlias(v string)`

SetCardAlias sets CardAlias field to given value.


### GetCardBankName

`func (o *CFCard) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *CFCard) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *CFCard) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.


### GetEmiTenure

`func (o *CFCard) GetEmiTenure() int32`

GetEmiTenure returns the EmiTenure field if non-nil, zero value otherwise.

### GetEmiTenureOk

`func (o *CFCard) GetEmiTenureOk() (*int32, bool)`

GetEmiTenureOk returns a tuple with the EmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiTenure

`func (o *CFCard) SetEmiTenure(v int32)`

SetEmiTenure sets EmiTenure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


