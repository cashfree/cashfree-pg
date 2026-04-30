# AuthResponseCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Can be link. | [optional] 
**CardNumber** | Pointer to **string** | Card number. | [optional] 
**CardNetwork** | Pointer to **string** | Card network (e.g., VISA, Mastercard). | [optional] 
**CardType** | Pointer to **string** | Card type (e.g., credit_card). | [optional] 
**CardSubType** | Pointer to **string** | Card subtype (e.g., R). | [optional] 
**CardCountry** | Pointer to **string** | Country of the card (e.g., IN). | [optional] 
**CardBankName** | Pointer to **string** | Bank name on card. | [optional] 
**CardNetworkReferenceId** | Pointer to **string** | Network reference ID. | [optional] 
**InstrumentId** | Pointer to **string** | Unique identifier for card instrument. | [optional] 

## Methods

### NewAuthResponseCard

`func NewAuthResponseCard() *AuthResponseCard`

NewAuthResponseCard instantiates a new AuthResponseCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponseCardWithDefaults

`func NewAuthResponseCardWithDefaults() *AuthResponseCard`

NewAuthResponseCardWithDefaults instantiates a new AuthResponseCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *AuthResponseCard) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AuthResponseCard) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AuthResponseCard) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AuthResponseCard) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCardNumber

`func (o *AuthResponseCard) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *AuthResponseCard) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *AuthResponseCard) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *AuthResponseCard) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNetwork

`func (o *AuthResponseCard) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *AuthResponseCard) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *AuthResponseCard) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *AuthResponseCard) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardType

`func (o *AuthResponseCard) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *AuthResponseCard) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *AuthResponseCard) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *AuthResponseCard) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardSubType

`func (o *AuthResponseCard) GetCardSubType() string`

GetCardSubType returns the CardSubType field if non-nil, zero value otherwise.

### GetCardSubTypeOk

`func (o *AuthResponseCard) GetCardSubTypeOk() (*string, bool)`

GetCardSubTypeOk returns a tuple with the CardSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSubType

`func (o *AuthResponseCard) SetCardSubType(v string)`

SetCardSubType sets CardSubType field to given value.

### HasCardSubType

`func (o *AuthResponseCard) HasCardSubType() bool`

HasCardSubType returns a boolean if a field has been set.

### GetCardCountry

`func (o *AuthResponseCard) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *AuthResponseCard) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *AuthResponseCard) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *AuthResponseCard) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardBankName

`func (o *AuthResponseCard) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *AuthResponseCard) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *AuthResponseCard) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *AuthResponseCard) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetCardNetworkReferenceId

`func (o *AuthResponseCard) GetCardNetworkReferenceId() string`

GetCardNetworkReferenceId returns the CardNetworkReferenceId field if non-nil, zero value otherwise.

### GetCardNetworkReferenceIdOk

`func (o *AuthResponseCard) GetCardNetworkReferenceIdOk() (*string, bool)`

GetCardNetworkReferenceIdOk returns a tuple with the CardNetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetworkReferenceId

`func (o *AuthResponseCard) SetCardNetworkReferenceId(v string)`

SetCardNetworkReferenceId sets CardNetworkReferenceId field to given value.

### HasCardNetworkReferenceId

`func (o *AuthResponseCard) HasCardNetworkReferenceId() bool`

HasCardNetworkReferenceId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *AuthResponseCard) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *AuthResponseCard) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *AuthResponseCard) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *AuthResponseCard) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


