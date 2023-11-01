# CardEMI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel for card payments will always be \&quot;link\&quot; | [optional] 
**CardNumber** | Pointer to **string** | Customer card number. | [optional] 
**CardHolderName** | Pointer to **string** | Customer name mentioned on the card. | [optional] 
**CardExpiryMm** | Pointer to **string** | Card expiry month. | [optional] 
**CardExpiryYy** | Pointer to **string** | Card expiry year. | [optional] 
**CardCvv** | Pointer to **string** | CVV mentioned on the card. | [optional] 
**CardAlias** | Pointer to **string** | Card alias as returned by Cashfree Vault API | [optional] 
**CardBankName** | Pointer to **string** | Card bank name, required for EMI payments. This is the bank user has selected for EMI. One of [\&quot;hdfc, \&quot;kotak\&quot;, \&quot;icici\&quot;, \&quot;rbl\&quot;, \&quot;bob\&quot;, \&quot;standard chartered\&quot;, \&quot;axis\&quot;, \&quot;au\&quot;, \&quot;yes\&quot;, \&quot;sbi\&quot;, \&quot;fed\&quot;, \&quot;hsbc\&quot;, \&quot;citi\&quot;, \&quot;amex\&quot;] | [optional] 
**EmiTenure** | Pointer to **int32** | EMI tenure selected by the user | [optional] 

## Methods

### NewCardEMI

`func NewCardEMI() *CardEMI`

NewCardEMI instantiates a new CardEMI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardEMIWithDefaults

`func NewCardEMIWithDefaults() *CardEMI`

NewCardEMIWithDefaults instantiates a new CardEMI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CardEMI) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CardEMI) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CardEMI) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CardEMI) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCardNumber

`func (o *CardEMI) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CardEMI) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CardEMI) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CardEMI) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardHolderName

`func (o *CardEMI) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *CardEMI) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *CardEMI) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *CardEMI) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *CardEMI) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CardEMI) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CardEMI) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.

### HasCardExpiryMm

`func (o *CardEMI) HasCardExpiryMm() bool`

HasCardExpiryMm returns a boolean if a field has been set.

### GetCardExpiryYy

`func (o *CardEMI) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CardEMI) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CardEMI) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.

### HasCardExpiryYy

`func (o *CardEMI) HasCardExpiryYy() bool`

HasCardExpiryYy returns a boolean if a field has been set.

### GetCardCvv

`func (o *CardEMI) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *CardEMI) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *CardEMI) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.

### HasCardCvv

`func (o *CardEMI) HasCardCvv() bool`

HasCardCvv returns a boolean if a field has been set.

### GetCardAlias

`func (o *CardEMI) GetCardAlias() string`

GetCardAlias returns the CardAlias field if non-nil, zero value otherwise.

### GetCardAliasOk

`func (o *CardEMI) GetCardAliasOk() (*string, bool)`

GetCardAliasOk returns a tuple with the CardAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAlias

`func (o *CardEMI) SetCardAlias(v string)`

SetCardAlias sets CardAlias field to given value.

### HasCardAlias

`func (o *CardEMI) HasCardAlias() bool`

HasCardAlias returns a boolean if a field has been set.

### GetCardBankName

`func (o *CardEMI) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *CardEMI) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *CardEMI) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *CardEMI) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetEmiTenure

`func (o *CardEMI) GetEmiTenure() int32`

GetEmiTenure returns the EmiTenure field if non-nil, zero value otherwise.

### GetEmiTenureOk

`func (o *CardEMI) GetEmiTenureOk() (*int32, bool)`

GetEmiTenureOk returns a tuple with the EmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiTenure

`func (o *CardEMI) SetEmiTenure(v int32)`

SetEmiTenure sets EmiTenure field to given value.

### HasEmiTenure

`func (o *CardEMI) HasEmiTenure() bool`

HasEmiTenure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


