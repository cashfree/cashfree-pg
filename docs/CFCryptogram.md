# CFCryptogram

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

### NewCFCryptogram

`func NewCFCryptogram() *CFCryptogram`

NewCFCryptogram instantiates a new CFCryptogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCryptogramWithDefaults

`func NewCFCryptogramWithDefaults() *CFCryptogram`

NewCFCryptogramWithDefaults instantiates a new CFCryptogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentId

`func (o *CFCryptogram) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *CFCryptogram) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *CFCryptogram) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *CFCryptogram) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetTokenRequestorId

`func (o *CFCryptogram) GetTokenRequestorId() string`

GetTokenRequestorId returns the TokenRequestorId field if non-nil, zero value otherwise.

### GetTokenRequestorIdOk

`func (o *CFCryptogram) GetTokenRequestorIdOk() (*string, bool)`

GetTokenRequestorIdOk returns a tuple with the TokenRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestorId

`func (o *CFCryptogram) SetTokenRequestorId(v string)`

SetTokenRequestorId sets TokenRequestorId field to given value.

### HasTokenRequestorId

`func (o *CFCryptogram) HasTokenRequestorId() bool`

HasTokenRequestorId returns a boolean if a field has been set.

### GetCardNumber

`func (o *CFCryptogram) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CFCryptogram) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CFCryptogram) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CFCryptogram) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardExpiryMm

`func (o *CFCryptogram) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *CFCryptogram) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *CFCryptogram) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.

### HasCardExpiryMm

`func (o *CFCryptogram) HasCardExpiryMm() bool`

HasCardExpiryMm returns a boolean if a field has been set.

### GetCardExpiryYy

`func (o *CFCryptogram) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *CFCryptogram) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *CFCryptogram) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.

### HasCardExpiryYy

`func (o *CFCryptogram) HasCardExpiryYy() bool`

HasCardExpiryYy returns a boolean if a field has been set.

### GetCryptogram

`func (o *CFCryptogram) GetCryptogram() string`

GetCryptogram returns the Cryptogram field if non-nil, zero value otherwise.

### GetCryptogramOk

`func (o *CFCryptogram) GetCryptogramOk() (*string, bool)`

GetCryptogramOk returns a tuple with the Cryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptogram

`func (o *CFCryptogram) SetCryptogram(v string)`

SetCryptogram sets Cryptogram field to given value.

### HasCryptogram

`func (o *CFCryptogram) HasCryptogram() bool`

HasCryptogram returns a boolean if a field has been set.

### GetCardDisplay

`func (o *CFCryptogram) GetCardDisplay() string`

GetCardDisplay returns the CardDisplay field if non-nil, zero value otherwise.

### GetCardDisplayOk

`func (o *CFCryptogram) GetCardDisplayOk() (*string, bool)`

GetCardDisplayOk returns a tuple with the CardDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDisplay

`func (o *CFCryptogram) SetCardDisplay(v string)`

SetCardDisplay sets CardDisplay field to given value.

### HasCardDisplay

`func (o *CFCryptogram) HasCardDisplay() bool`

HasCardDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


