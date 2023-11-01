# OfferValidationsPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **map[string]interface{}** | All offers applicable | 
**Card** | [**CardOffer**](CardOffer.md) |  | 
**Netbanking** | [**OfferNBNetbanking**](OfferNBNetbanking.md) |  | 
**App** | Pointer to [**WalletOffer**](WalletOffer.md) |  | [optional] 
**Upi** | **map[string]interface{}** |  | 
**Paylater** | [**PaylaterOffer**](PaylaterOffer.md) |  | 
**Emi** | [**EMIOffer**](EMIOffer.md) |  | 

## Methods

### NewOfferValidationsPaymentMethod

`func NewOfferValidationsPaymentMethod(all map[string]interface{}, card CardOffer, netbanking OfferNBNetbanking, upi map[string]interface{}, paylater PaylaterOffer, emi EMIOffer, ) *OfferValidationsPaymentMethod`

NewOfferValidationsPaymentMethod instantiates a new OfferValidationsPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferValidationsPaymentMethodWithDefaults

`func NewOfferValidationsPaymentMethodWithDefaults() *OfferValidationsPaymentMethod`

NewOfferValidationsPaymentMethodWithDefaults instantiates a new OfferValidationsPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *OfferValidationsPaymentMethod) GetAll() map[string]interface{}`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *OfferValidationsPaymentMethod) GetAllOk() (*map[string]interface{}, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *OfferValidationsPaymentMethod) SetAll(v map[string]interface{})`

SetAll sets All field to given value.


### GetCard

`func (o *OfferValidationsPaymentMethod) GetCard() CardOffer`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *OfferValidationsPaymentMethod) GetCardOk() (*CardOffer, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *OfferValidationsPaymentMethod) SetCard(v CardOffer)`

SetCard sets Card field to given value.


### GetNetbanking

`func (o *OfferValidationsPaymentMethod) GetNetbanking() OfferNBNetbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *OfferValidationsPaymentMethod) GetNetbankingOk() (*OfferNBNetbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *OfferValidationsPaymentMethod) SetNetbanking(v OfferNBNetbanking)`

SetNetbanking sets Netbanking field to given value.


### GetApp

`func (o *OfferValidationsPaymentMethod) GetApp() WalletOffer`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OfferValidationsPaymentMethod) GetAppOk() (*WalletOffer, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OfferValidationsPaymentMethod) SetApp(v WalletOffer)`

SetApp sets App field to given value.

### HasApp

`func (o *OfferValidationsPaymentMethod) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetUpi

`func (o *OfferValidationsPaymentMethod) GetUpi() map[string]interface{}`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *OfferValidationsPaymentMethod) GetUpiOk() (*map[string]interface{}, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *OfferValidationsPaymentMethod) SetUpi(v map[string]interface{})`

SetUpi sets Upi field to given value.


### GetPaylater

`func (o *OfferValidationsPaymentMethod) GetPaylater() PaylaterOffer`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *OfferValidationsPaymentMethod) GetPaylaterOk() (*PaylaterOffer, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *OfferValidationsPaymentMethod) SetPaylater(v PaylaterOffer)`

SetPaylater sets Paylater field to given value.


### GetEmi

`func (o *OfferValidationsPaymentMethod) GetEmi() EMIOffer`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *OfferValidationsPaymentMethod) GetEmiOk() (*EMIOffer, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *OfferValidationsPaymentMethod) SetEmi(v EMIOffer)`

SetEmi sets Emi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


