# OfferValidationsResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **map[string]interface{}** | All offers applicable | 
**Card** | [**CardOffer**](CardOffer.md) |  | 
**Netbanking** | [**OfferNBNetbanking**](OfferNBNetbanking.md) |  | 
**App** | [**WalletOffer**](WalletOffer.md) |  | 
**Upi** | **map[string]interface{}** |  | 
**Paylater** | [**PaylaterOffer**](PaylaterOffer.md) |  | 
**Emi** | [**EMIOffer**](EMIOffer.md) |  | 

## Methods

### NewOfferValidationsResponsePaymentMethod

`func NewOfferValidationsResponsePaymentMethod(all map[string]interface{}, card CardOffer, netbanking OfferNBNetbanking, app WalletOffer, upi map[string]interface{}, paylater PaylaterOffer, emi EMIOffer, ) *OfferValidationsResponsePaymentMethod`

NewOfferValidationsResponsePaymentMethod instantiates a new OfferValidationsResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferValidationsResponsePaymentMethodWithDefaults

`func NewOfferValidationsResponsePaymentMethodWithDefaults() *OfferValidationsResponsePaymentMethod`

NewOfferValidationsResponsePaymentMethodWithDefaults instantiates a new OfferValidationsResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *OfferValidationsResponsePaymentMethod) GetAll() map[string]interface{}`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *OfferValidationsResponsePaymentMethod) GetAllOk() (*map[string]interface{}, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *OfferValidationsResponsePaymentMethod) SetAll(v map[string]interface{})`

SetAll sets All field to given value.


### GetCard

`func (o *OfferValidationsResponsePaymentMethod) GetCard() CardOffer`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *OfferValidationsResponsePaymentMethod) GetCardOk() (*CardOffer, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *OfferValidationsResponsePaymentMethod) SetCard(v CardOffer)`

SetCard sets Card field to given value.


### GetNetbanking

`func (o *OfferValidationsResponsePaymentMethod) GetNetbanking() OfferNBNetbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *OfferValidationsResponsePaymentMethod) GetNetbankingOk() (*OfferNBNetbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *OfferValidationsResponsePaymentMethod) SetNetbanking(v OfferNBNetbanking)`

SetNetbanking sets Netbanking field to given value.


### GetApp

`func (o *OfferValidationsResponsePaymentMethod) GetApp() WalletOffer`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OfferValidationsResponsePaymentMethod) GetAppOk() (*WalletOffer, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OfferValidationsResponsePaymentMethod) SetApp(v WalletOffer)`

SetApp sets App field to given value.


### GetUpi

`func (o *OfferValidationsResponsePaymentMethod) GetUpi() map[string]interface{}`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *OfferValidationsResponsePaymentMethod) GetUpiOk() (*map[string]interface{}, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *OfferValidationsResponsePaymentMethod) SetUpi(v map[string]interface{})`

SetUpi sets Upi field to given value.


### GetPaylater

`func (o *OfferValidationsResponsePaymentMethod) GetPaylater() PaylaterOffer`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *OfferValidationsResponsePaymentMethod) GetPaylaterOk() (*PaylaterOffer, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *OfferValidationsResponsePaymentMethod) SetPaylater(v PaylaterOffer)`

SetPaylater sets Paylater field to given value.


### GetEmi

`func (o *OfferValidationsResponsePaymentMethod) GetEmi() EMIOffer`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *OfferValidationsResponsePaymentMethod) GetEmiOk() (*EMIOffer, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *OfferValidationsResponsePaymentMethod) SetEmi(v EMIOffer)`

SetEmi sets Emi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


