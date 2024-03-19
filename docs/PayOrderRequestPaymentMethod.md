# PayOrderRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**Card**](Card.md) |  | 
**Upi** | [**Upi**](Upi.md) |  | 
**Netbanking** | [**Netbanking**](Netbanking.md) |  | 
**App** | [**App**](App.md) |  | 
**Emi** | [**CardEMI**](CardEMI.md) |  | 
**CardlessEmi** | [**CardlessEMI**](CardlessEMI.md) |  | 
**Paylater** | [**Paylater**](Paylater.md) |  | 
**Banktransfer** | [**Banktransfer**](Banktransfer.md) |  | 

## Methods

### NewPayOrderRequestPaymentMethod

`func NewPayOrderRequestPaymentMethod(card Card, upi Upi, netbanking Netbanking, app App, emi CardEMI, cardlessEmi CardlessEMI, paylater Paylater, banktransfer Banktransfer, ) *PayOrderRequestPaymentMethod`

NewPayOrderRequestPaymentMethod instantiates a new PayOrderRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderRequestPaymentMethodWithDefaults

`func NewPayOrderRequestPaymentMethodWithDefaults() *PayOrderRequestPaymentMethod`

NewPayOrderRequestPaymentMethodWithDefaults instantiates a new PayOrderRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *PayOrderRequestPaymentMethod) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PayOrderRequestPaymentMethod) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PayOrderRequestPaymentMethod) SetCard(v Card)`

SetCard sets Card field to given value.


### GetUpi

`func (o *PayOrderRequestPaymentMethod) GetUpi() Upi`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *PayOrderRequestPaymentMethod) GetUpiOk() (*Upi, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *PayOrderRequestPaymentMethod) SetUpi(v Upi)`

SetUpi sets Upi field to given value.


### GetNetbanking

`func (o *PayOrderRequestPaymentMethod) GetNetbanking() Netbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *PayOrderRequestPaymentMethod) GetNetbankingOk() (*Netbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *PayOrderRequestPaymentMethod) SetNetbanking(v Netbanking)`

SetNetbanking sets Netbanking field to given value.


### GetApp

`func (o *PayOrderRequestPaymentMethod) GetApp() App`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PayOrderRequestPaymentMethod) GetAppOk() (*App, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PayOrderRequestPaymentMethod) SetApp(v App)`

SetApp sets App field to given value.


### GetEmi

`func (o *PayOrderRequestPaymentMethod) GetEmi() CardEMI`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *PayOrderRequestPaymentMethod) GetEmiOk() (*CardEMI, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *PayOrderRequestPaymentMethod) SetEmi(v CardEMI)`

SetEmi sets Emi field to given value.


### GetCardlessEmi

`func (o *PayOrderRequestPaymentMethod) GetCardlessEmi() CardlessEMI`

GetCardlessEmi returns the CardlessEmi field if non-nil, zero value otherwise.

### GetCardlessEmiOk

`func (o *PayOrderRequestPaymentMethod) GetCardlessEmiOk() (*CardlessEMI, bool)`

GetCardlessEmiOk returns a tuple with the CardlessEmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardlessEmi

`func (o *PayOrderRequestPaymentMethod) SetCardlessEmi(v CardlessEMI)`

SetCardlessEmi sets CardlessEmi field to given value.


### GetPaylater

`func (o *PayOrderRequestPaymentMethod) GetPaylater() Paylater`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *PayOrderRequestPaymentMethod) GetPaylaterOk() (*Paylater, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *PayOrderRequestPaymentMethod) SetPaylater(v Paylater)`

SetPaylater sets Paylater field to given value.


### GetBanktransfer

`func (o *PayOrderRequestPaymentMethod) GetBanktransfer() Banktransfer`

GetBanktransfer returns the Banktransfer field if non-nil, zero value otherwise.

### GetBanktransferOk

`func (o *PayOrderRequestPaymentMethod) GetBanktransferOk() (*Banktransfer, bool)`

GetBanktransferOk returns a tuple with the Banktransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanktransfer

`func (o *PayOrderRequestPaymentMethod) SetBanktransfer(v Banktransfer)`

SetBanktransfer sets Banktransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


