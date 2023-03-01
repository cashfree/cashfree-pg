# CFCardEMI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for card payments will always be \&quot;link\&quot; | [default to "link"]
**CardNumber** | **string** | Customer card number. | 
**CardHolderName** | Pointer to **string** | Customer name mentioned on the card. | [optional] 
**CardExpiryMm** | **string** | Card expiry month. | 
**CardExpiryYy** | **string** | Card expiry year. | 
**CardCvv** | **string** | CVV mentioned on the card. | 
**CardAlias** | Pointer to **string** | Card alias as returned by Cashfree Vault API | [optional] 
**CardBankName** | **string** | Card bank name, required for EMI payments. This is the bank user has selected for EMI. One of [\&quot;Kotak\&quot;, \&quot;ICICI\&quot;, \&quot;RBL\&quot;, \&quot;BOB\&quot;, \&quot;Standard Chartered\&quot;, \&quot;HDFC\&quot;] | 
**EmiTenure** | **int32** | EMI tenure selected by the user | 

## Methods

### NewCFCardEMI

`func NewCFCardEMI(channel string, cardNumber string, cardExpiryMm string, cardExpiryYy string, cardCvv string, cardBankName string, emiTenure int32, ) *CFCardEMI`

NewCFCardEMI instantiates a new CFCardEMI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCardEMIWithDefaults

`func NewCFCardEMIWithDefaults() *CFCardEMI`

NewCFCardEMIWithDefaults instantiates a new CFCardEMI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFCardEMI) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFCardEMI) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFCardEMI) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCardNumber

`func (o *CFCardEMI) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CFCardEMI) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CFCardEMI) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardHolderName

`func (o *CFCardEMI) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *CFCardEMI) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *CFCardEMI) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *CFCardEMI) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *CFCardEMI) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CFCardEMI) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CFCardEMI) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.


### GetCardExpiryYy

`func (o *CFCardEMI) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CFCardEMI) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CFCardEMI) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.


### GetCardCvv

`func (o *CFCardEMI) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *CFCardEMI) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *CFCardEMI) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.


### GetCardAlias

`func (o *CFCardEMI) GetCardAlias() string`

GetCardAlias returns the CardAlias field if non-nil, zero value otherwise.

### GetCardAliasOk

`func (o *CFCardEMI) GetCardAliasOk() (*string, bool)`

GetCardAliasOk returns a tuple with the CardAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAlias

`func (o *CFCardEMI) SetCardAlias(v string)`

SetCardAlias sets CardAlias field to given value.

### HasCardAlias

`func (o *CFCardEMI) HasCardAlias() bool`

HasCardAlias returns a boolean if a field has been set.

### GetCardBankName

`func (o *CFCardEMI) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *CFCardEMI) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *CFCardEMI) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.


### GetEmiTenure

`func (o *CFCardEMI) GetEmiTenure() int32`

GetEmiTenure returns the EmiTenure field if non-nil, zero value otherwise.

### GetEmiTenureOk

`func (o *CFCardEMI) GetEmiTenureOk() (*int32, bool)`

GetEmiTenureOk returns a tuple with the EmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiTenure

`func (o *CFCardEMI) SetEmiTenure(v int32)`

SetEmiTenure sets EmiTenure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


