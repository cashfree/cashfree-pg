# SavedInstrumentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNetwork** | Pointer to **string** | card scheme/network of the saved card. Example visa, mastercard | [optional] 
**CardBankName** | Pointer to **string** | Issuing bank name of saved card | [optional] 
**CardCountry** | Pointer to **string** | Issuing country of saved card | [optional] 
**CardType** | Pointer to **string** | Type of saved card | [optional] 
**CardTokenDetails** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSavedInstrumentMeta

`func NewSavedInstrumentMeta() *SavedInstrumentMeta`

NewSavedInstrumentMeta instantiates a new SavedInstrumentMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedInstrumentMetaWithDefaults

`func NewSavedInstrumentMetaWithDefaults() *SavedInstrumentMeta`

NewSavedInstrumentMetaWithDefaults instantiates a new SavedInstrumentMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNetwork

`func (o *SavedInstrumentMeta) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *SavedInstrumentMeta) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *SavedInstrumentMeta) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *SavedInstrumentMeta) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardBankName

`func (o *SavedInstrumentMeta) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *SavedInstrumentMeta) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *SavedInstrumentMeta) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *SavedInstrumentMeta) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetCardCountry

`func (o *SavedInstrumentMeta) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *SavedInstrumentMeta) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *SavedInstrumentMeta) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *SavedInstrumentMeta) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardType

`func (o *SavedInstrumentMeta) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *SavedInstrumentMeta) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *SavedInstrumentMeta) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *SavedInstrumentMeta) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardTokenDetails

`func (o *SavedInstrumentMeta) GetCardTokenDetails() map[string]interface{}`

GetCardTokenDetails returns the CardTokenDetails field if non-nil, zero value otherwise.

### GetCardTokenDetailsOk

`func (o *SavedInstrumentMeta) GetCardTokenDetailsOk() (*map[string]interface{}, bool)`

GetCardTokenDetailsOk returns a tuple with the CardTokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTokenDetails

`func (o *SavedInstrumentMeta) SetCardTokenDetails(v map[string]interface{})`

SetCardTokenDetails sets CardTokenDetails field to given value.

### HasCardTokenDetails

`func (o *SavedInstrumentMeta) HasCardTokenDetails() bool`

HasCardTokenDetails returns a boolean if a field has been set.

### SetCardTokenDetailsNil

`func (o *SavedInstrumentMeta) SetCardTokenDetailsNil(b bool)`

 SetCardTokenDetailsNil sets the value for CardTokenDetails to be an explicit nil

### UnsetCardTokenDetails
`func (o *SavedInstrumentMeta) UnsetCardTokenDetails()`

UnsetCardTokenDetails ensures that no value is present for CardTokenDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


