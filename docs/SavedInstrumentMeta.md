# SavedInstrumentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNetwork** | Pointer to **string** | Card scheme/network of the saved card. Available options are &#x60;visa&#x60;, &#x60;mastercard&#x60;, &#x60;rupay&#x60;, &#x60;amex&#x60;, and &#x60;diners&#x60;. | [optional] 
**CardBankName** | Pointer to **string** | Issuing bank name of the saved card. For example, &#x60;HDFC BANK&#x60;, &#x60;AXIS BANK&#x60;, or &#x60;ICICI BANK&#x60;. | [optional] 
**CardCountry** | Pointer to **string** | Issuing country of the saved card. For example, &#x60;IN&#x60;. | [optional] 
**CardType** | Pointer to **string** | Type of the saved card. Available options are &#x60;credit_card&#x60;, &#x60;debit_card&#x60;, and &#x60;prepaid_card&#x60;. | [optional] 
**CardSubType** | Pointer to **string** | Subtype of the saved card. R indicates retail, P indicates premium, and C indicates corporate. Available options are &#x60;R&#x60;, &#x60;P&#x60;, and &#x60;C&#x60;. | [optional] 
**CardTokenDetails** | Pointer to [**SavedInstrumentMetaCardTokenDetails**](SavedInstrumentMetaCardTokenDetails.md) |  | [optional] 

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

### GetCardSubType

`func (o *SavedInstrumentMeta) GetCardSubType() string`

GetCardSubType returns the CardSubType field if non-nil, zero value otherwise.

### GetCardSubTypeOk

`func (o *SavedInstrumentMeta) GetCardSubTypeOk() (*string, bool)`

GetCardSubTypeOk returns a tuple with the CardSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSubType

`func (o *SavedInstrumentMeta) SetCardSubType(v string)`

SetCardSubType sets CardSubType field to given value.

### HasCardSubType

`func (o *SavedInstrumentMeta) HasCardSubType() bool`

HasCardSubType returns a boolean if a field has been set.

### GetCardTokenDetails

`func (o *SavedInstrumentMeta) GetCardTokenDetails() SavedInstrumentMetaCardTokenDetails`

GetCardTokenDetails returns the CardTokenDetails field if non-nil, zero value otherwise.

### GetCardTokenDetailsOk

`func (o *SavedInstrumentMeta) GetCardTokenDetailsOk() (*SavedInstrumentMetaCardTokenDetails, bool)`

GetCardTokenDetailsOk returns a tuple with the CardTokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTokenDetails

`func (o *SavedInstrumentMeta) SetCardTokenDetails(v SavedInstrumentMetaCardTokenDetails)`

SetCardTokenDetails sets CardTokenDetails field to given value.

### HasCardTokenDetails

`func (o *SavedInstrumentMeta) HasCardTokenDetails() bool`

HasCardTokenDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


