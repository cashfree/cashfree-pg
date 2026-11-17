# CardDowntimeCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardIssuer** | Pointer to **[]string** |  | [optional] 
**CardType** | Pointer to [**[]CardDowntimeCardTypeEnum**](CardDowntimeCardTypeEnum.md) |  | [optional] 
**CardNetwork** | Pointer to [**[]CardDowntimeCardNetworkEnum**](CardDowntimeCardNetworkEnum.md) |  | [optional] 

## Methods

### NewCardDowntimeCard

`func NewCardDowntimeCard() *CardDowntimeCard`

NewCardDowntimeCard instantiates a new CardDowntimeCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDowntimeCardWithDefaults

`func NewCardDowntimeCardWithDefaults() *CardDowntimeCard`

NewCardDowntimeCardWithDefaults instantiates a new CardDowntimeCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIssuer

`func (o *CardDowntimeCard) GetCardIssuer() []string`

GetCardIssuer returns the CardIssuer field if non-nil, zero value otherwise.

### GetCardIssuerOk

`func (o *CardDowntimeCard) GetCardIssuerOk() (*[]string, bool)`

GetCardIssuerOk returns a tuple with the CardIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuer

`func (o *CardDowntimeCard) SetCardIssuer(v []string)`

SetCardIssuer sets CardIssuer field to given value.

### HasCardIssuer

`func (o *CardDowntimeCard) HasCardIssuer() bool`

HasCardIssuer returns a boolean if a field has been set.

### GetCardType

`func (o *CardDowntimeCard) GetCardType() []CardDowntimeCardTypeEnum`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardDowntimeCard) GetCardTypeOk() (*[]CardDowntimeCardTypeEnum, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardDowntimeCard) SetCardType(v []CardDowntimeCardTypeEnum)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CardDowntimeCard) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardNetwork

`func (o *CardDowntimeCard) GetCardNetwork() []CardDowntimeCardNetworkEnum`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *CardDowntimeCard) GetCardNetworkOk() (*[]CardDowntimeCardNetworkEnum, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *CardDowntimeCard) SetCardNetwork(v []CardDowntimeCardNetworkEnum)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *CardDowntimeCard) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


