# CFPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**CFCard**](CFCard.md) |  | 
**Upi** | [**CFUPI**](CFUPI.md) |  | 
**Netbanking** | [**CFNetbanking**](CFNetbanking.md) |  | 
**App** | [**CFApp**](CFApp.md) |  | 
**Emi** | Pointer to [**CFCardEMI**](CFCardEMI.md) |  | [optional] 
**CardlessEmi** | [**CFCardlessEMI**](CFCardlessEMI.md) |  | 
**Paylater** | [**CFPaylater**](CFPaylater.md) |  | 

## Methods

### NewCFPaymentMethod

`func NewCFPaymentMethod(card CFCard, upi CFUPI, netbanking CFNetbanking, app CFApp, cardlessEmi CFCardlessEMI, paylater CFPaylater, ) *CFPaymentMethod`

NewCFPaymentMethod instantiates a new CFPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFPaymentMethodWithDefaults

`func NewCFPaymentMethodWithDefaults() *CFPaymentMethod`

NewCFPaymentMethodWithDefaults instantiates a new CFPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *CFPaymentMethod) GetCard() CFCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CFPaymentMethod) GetCardOk() (*CFCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CFPaymentMethod) SetCard(v CFCard)`

SetCard sets Card field to given value.


### GetUpi

`func (o *CFPaymentMethod) GetUpi() CFUPI`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *CFPaymentMethod) GetUpiOk() (*CFUPI, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *CFPaymentMethod) SetUpi(v CFUPI)`

SetUpi sets Upi field to given value.


### GetNetbanking

`func (o *CFPaymentMethod) GetNetbanking() CFNetbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *CFPaymentMethod) GetNetbankingOk() (*CFNetbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *CFPaymentMethod) SetNetbanking(v CFNetbanking)`

SetNetbanking sets Netbanking field to given value.


### GetApp

`func (o *CFPaymentMethod) GetApp() CFApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CFPaymentMethod) GetAppOk() (*CFApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CFPaymentMethod) SetApp(v CFApp)`

SetApp sets App field to given value.


### GetEmi

`func (o *CFPaymentMethod) GetEmi() CFCardEMI`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *CFPaymentMethod) GetEmiOk() (*CFCardEMI, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *CFPaymentMethod) SetEmi(v CFCardEMI)`

SetEmi sets Emi field to given value.

### HasEmi

`func (o *CFPaymentMethod) HasEmi() bool`

HasEmi returns a boolean if a field has been set.

### GetCardlessEmi

`func (o *CFPaymentMethod) GetCardlessEmi() CFCardlessEMI`

GetCardlessEmi returns the CardlessEmi field if non-nil, zero value otherwise.

### GetCardlessEmiOk

`func (o *CFPaymentMethod) GetCardlessEmiOk() (*CFCardlessEMI, bool)`

GetCardlessEmiOk returns a tuple with the CardlessEmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardlessEmi

`func (o *CFPaymentMethod) SetCardlessEmi(v CFCardlessEMI)`

SetCardlessEmi sets CardlessEmi field to given value.


### GetPaylater

`func (o *CFPaymentMethod) GetPaylater() CFPaylater`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *CFPaymentMethod) GetPaylaterOk() (*CFPaylater, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *CFPaymentMethod) SetPaylater(v CFPaylater)`

SetPaylater sets Paylater field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


