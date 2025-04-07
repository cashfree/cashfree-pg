# PARRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | Card number, between 15 and 19 digits. | 
**CardCvv** | **string** | Card CVV, 3 or 4 digits. | 
**CardExpiryMm** | **string** | Two-digit card expiry month (01-12). | 
**CardExpiryYy** | **string** | Two-digit card expiry year. | 
**CardType** | **string** | Card type; allowed value is PLAIN_CARD. | 

## Methods

### NewPARRequest

`func NewPARRequest(cardNumber string, cardCvv string, cardExpiryMm string, cardExpiryYy string, cardType string, ) *PARRequest`

NewPARRequest instantiates a new PARRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPARRequestWithDefaults

`func NewPARRequestWithDefaults() *PARRequest`

NewPARRequestWithDefaults instantiates a new PARRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *PARRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PARRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PARRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCardCvv

`func (o *PARRequest) GetCardCvv() string`

GetCardCvv returns the CardCvv field if non-nil, zero value otherwise.

### GetCardCvvOk

`func (o *PARRequest) GetCardCvvOk() (*string, bool)`

GetCardCvvOk returns a tuple with the CardCvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCvv

`func (o *PARRequest) SetCardCvv(v string)`

SetCardCvv sets CardCvv field to given value.


### GetCardExpiryMm

`func (o *PARRequest) GetCardExpiryMm() string`

GetCardExpiryMm returns the CardExpiryMm field if non-nil, zero value otherwise.

### GetCardExpiryMmOk

`func (o *PARRequest) GetCardExpiryMmOk() (*string, bool)`

GetCardExpiryMmOk returns a tuple with the CardExpiryMm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryMm

`func (o *PARRequest) SetCardExpiryMm(v string)`

SetCardExpiryMm sets CardExpiryMm field to given value.


### GetCardExpiryYy

`func (o *PARRequest) GetCardExpiryYy() string`

GetCardExpiryYy returns the CardExpiryYy field if non-nil, zero value otherwise.

### GetCardExpiryYyOk

`func (o *PARRequest) GetCardExpiryYyOk() (*string, bool)`

GetCardExpiryYyOk returns a tuple with the CardExpiryYy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpiryYy

`func (o *PARRequest) SetCardExpiryYy(v string)`

SetCardExpiryYy sets CardExpiryYy field to given value.


### GetCardType

`func (o *PARRequest) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PARRequest) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PARRequest) SetCardType(v string)`

SetCardType sets CardType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


