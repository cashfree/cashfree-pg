# CryptogramEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentId** | Pointer to **string** | instrument_id of saved instrument | [optional] 
**TokenRequestorId** | Pointer to **string** | TRID issued by card networks | [optional] 
**CardNumber** | Pointer to **string** | token pan number | [optional] 
**CardExpiryMm** | Pointer to **string** | token pan expiry month | [optional] 
**CardExpiryYy** | Pointer to **string** | token pan expiry year | [optional] 
**Cryptogram** | Pointer to **string** | cryptogram | [optional] 
**CardDisplay** | Pointer to **string** | last 4 digits of original card number | [optional] 

## Methods

### NewCryptogramEntity

`func NewCryptogramEntity() *CryptogramEntity`

NewCryptogramEntity instantiates a new CryptogramEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptogramEntityWithDefaults

`func NewCryptogramEntityWithDefaults() *CryptogramEntity`

NewCryptogramEntityWithDefaults instantiates a new CryptogramEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *CryptogramEntity) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *CryptogramEntity) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *CryptogramEntity) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *CryptogramEntity) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetTokenRequestorId

`func (o *CryptogramEntity) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *CryptogramEntity) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *CryptogramEntity) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.

### HasTokenRequestorId

`func (o *CryptogramEntity) HasTokenRequestorId() bool`

HasTokenRequestorId returns a boolean if a field has been set.

### GetCardNumber

`func (o *CryptogramEntity) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CryptogramEntity) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CryptogramEntity) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CryptogramEntity) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *CryptogramEntity) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CryptogramEntity) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CryptogramEntity) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.

### HasCardExpiryMm

`func (o *CryptogramEntity) HasCardExpiryMm() bool`

HasCardExpiryMm returns a boolean if a field has been set.

### GetCardExpiryYy

`func (o *CryptogramEntity) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CryptogramEntity) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CryptogramEntity) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.

### HasCardExpiryYy

`func (o *CryptogramEntity) HasCardExpiryYy() bool`

HasCardExpiryYy returns a boolean if a field has been set.

### GetCryptogram

`func (o *CryptogramEntity) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *CryptogramEntity) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *CryptogramEntity) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *CryptogramEntity) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### GetCardDisplay

`func (o *CryptogramEntity) GetCardDisplay() string`

GetCardDisplay returns the CardDisplay field if non-nil, zero value otherwise.

### GetCardDisplayOk

`func (o *CryptogramEntity) GetCardDisplayOk() (*string, bool)`

GetCardDisplayOk returns a tuple with the CardDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDisplay

`func (o *CryptogramEntity) SetCardDisplay(v string)`

SetCardDisplay sets CardDisplay field to given value.

### HasCardDisplay

`func (o *CryptogramEntity) HasCardDisplay() bool`

HasCardDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


