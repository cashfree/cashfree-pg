# CardOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **[]string** |  | 
**BankName** | **string** | Bank Name of Card. | 
**SchemeName** | **[]string** |  | 

## Methods

### NewCardOffer

`func NewCardOffer(type_ []string, bankName string, schemeName []string, ) *CardOffer`

NewCardOffer instantiates a new CardOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOfferWithDefaults

`func NewCardOfferWithDefaults() *CardOffer`

NewCardOfferWithDefaults instantiates a new CardOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CardOffer) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardOffer) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardOffer) SetType(v []string)`

SetType sets Type field to given value.


### GetBankName

`func (o *CardOffer) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CardOffer) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CardOffer) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetSchemeName

`func (o *CardOffer) GetSchemeName() []string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *CardOffer) GetSchemeNameOk() (*[]string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *CardOffer) SetSchemeName(v []string)`

SetSchemeName sets SchemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


