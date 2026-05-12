# AuthorizationDetailsPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upi** | Pointer to [**AuthResponseUpi**](AuthResponseUpi.md) |  | [optional] 
**Enach** | Pointer to [**AuthResponseEnach**](AuthResponseEnach.md) |  | [optional] 
**Pnach** | Pointer to [**AuthResponsePnach**](AuthResponsePnach.md) |  | [optional] 
**Card** | Pointer to [**AuthResponseCard**](AuthResponseCard.md) |  | [optional] 

## Methods

### NewAuthorizationDetailsPaymentMethod

`func NewAuthorizationDetailsPaymentMethod() *AuthorizationDetailsPaymentMethod`

NewAuthorizationDetailsPaymentMethod instantiates a new AuthorizationDetailsPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailsPaymentMethodWithDefaults

`func NewAuthorizationDetailsPaymentMethodWithDefaults() *AuthorizationDetailsPaymentMethod`

NewAuthorizationDetailsPaymentMethodWithDefaults instantiates a new AuthorizationDetailsPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpi

`func (o *AuthorizationDetailsPaymentMethod) GetUpi() AuthResponseUpi`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *AuthorizationDetailsPaymentMethod) GetUpiOk() (*AuthResponseUpi, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *AuthorizationDetailsPaymentMethod) SetUpi(v AuthResponseUpi)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *AuthorizationDetailsPaymentMethod) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetEnach

`func (o *AuthorizationDetailsPaymentMethod) GetEnach() AuthResponseEnach`

GetEnach returns the Enach field if non-nil, zero value otherwise.

### GetEnachOk

`func (o *AuthorizationDetailsPaymentMethod) GetEnachOk() (*AuthResponseEnach, bool)`

GetEnachOk returns a tuple with the Enach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnach

`func (o *AuthorizationDetailsPaymentMethod) SetEnach(v AuthResponseEnach)`

SetEnach sets Enach field to given value.

### HasEnach

`func (o *AuthorizationDetailsPaymentMethod) HasEnach() bool`

HasEnach returns a boolean if a field has been set.

### GetPnach

`func (o *AuthorizationDetailsPaymentMethod) GetPnach() AuthResponsePnach`

GetPnach returns the Pnach field if non-nil, zero value otherwise.

### GetPnachOk

`func (o *AuthorizationDetailsPaymentMethod) GetPnachOk() (*AuthResponsePnach, bool)`

GetPnachOk returns a tuple with the Pnach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnach

`func (o *AuthorizationDetailsPaymentMethod) SetPnach(v AuthResponsePnach)`

SetPnach sets Pnach field to given value.

### HasPnach

`func (o *AuthorizationDetailsPaymentMethod) HasPnach() bool`

HasPnach returns a boolean if a field has been set.

### GetCard

`func (o *AuthorizationDetailsPaymentMethod) GetCard() AuthResponseCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *AuthorizationDetailsPaymentMethod) GetCardOk() (*AuthResponseCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *AuthorizationDetailsPaymentMethod) SetCard(v AuthResponseCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *AuthorizationDetailsPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


