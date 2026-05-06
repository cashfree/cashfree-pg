# GetCardBinResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | Pointer to **string** | Issuing bank name of the card. For example &#x60;hdfc bank&#x60;, &#x60;icici bank&#x60;, &#x60;axis bank&#x60;. | [optional] 
**CountryCode** | Pointer to **string** | Issuing country of the card. For example &#x60;in&#x60;, &#x60;us&#x60;. | [optional] 
**Scheme** | Pointer to **string** | Card scheme/network of the card. For example &#x60;visa&#x60;, &#x60;mastercard&#x60;, &#x60;rupay&#x60;, &#x60;amex&#x60;, &#x60;diners&#x60;. | [optional] 
**SubType** | Pointer to **string** | Sub-type of card. Available options are &#x60;retail&#x60;, &#x60;premium&#x60; and &#x60;corporate&#x60;. | [optional] 
**Type** | Pointer to **string** | Type of card. Available options are &#x60;credit&#x60;, &#x60;debit&#x60; and &#x60;prepaid&#x60;. | [optional] 

## Methods

### NewGetCardBinResponseSchema

`func NewGetCardBinResponseSchema() *GetCardBinResponseSchema`

NewGetCardBinResponseSchema instantiates a new GetCardBinResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCardBinResponseSchemaWithDefaults

`func NewGetCardBinResponseSchemaWithDefaults() *GetCardBinResponseSchema`

NewGetCardBinResponseSchemaWithDefaults instantiates a new GetCardBinResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *GetCardBinResponseSchema) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *GetCardBinResponseSchema) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *GetCardBinResponseSchema) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *GetCardBinResponseSchema) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetCountryCode

`func (o *GetCardBinResponseSchema) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GetCardBinResponseSchema) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GetCardBinResponseSchema) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GetCardBinResponseSchema) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetScheme

`func (o *GetCardBinResponseSchema) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *GetCardBinResponseSchema) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *GetCardBinResponseSchema) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *GetCardBinResponseSchema) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSubType

`func (o *GetCardBinResponseSchema) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetCardBinResponseSchema) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetCardBinResponseSchema) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetCardBinResponseSchema) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *GetCardBinResponseSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCardBinResponseSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCardBinResponseSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCardBinResponseSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


