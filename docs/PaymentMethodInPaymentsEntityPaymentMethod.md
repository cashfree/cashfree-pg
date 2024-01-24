# PaymentMethodInPaymentsEntityPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**PaymentMethodCardInPaymentsEntityCard**](PaymentMethodCardInPaymentsEntityCard.md) |  | [optional] 
**Netbanking** | Pointer to [**PaymentMethodNetBankingInPaymentsEntityNetbanking**](PaymentMethodNetBankingInPaymentsEntityNetbanking.md) |  | [optional] 
**Upi** | Pointer to [**PaymentMethodUPIInPaymentsEntityUpi**](PaymentMethodUPIInPaymentsEntityUpi.md) |  | [optional] 
**App** | Pointer to [**PaymentMethodAppInPaymentsEntityApp**](PaymentMethodAppInPaymentsEntityApp.md) |  | [optional] 
**CardlessEmi** | Pointer to [**PaymentMethodAppInPaymentsEntityApp**](PaymentMethodAppInPaymentsEntityApp.md) |  | [optional] 
**Paylater** | Pointer to [**PaymentMethodAppInPaymentsEntityApp**](PaymentMethodAppInPaymentsEntityApp.md) |  | [optional] 
**Emi** | Pointer to [**PaymentMethodCardEMIInPaymentsEntityEmi**](PaymentMethodCardEMIInPaymentsEntityEmi.md) |  | [optional] 

## Methods

### NewPaymentMethodInPaymentsEntityPaymentMethod

`func NewPaymentMethodInPaymentsEntityPaymentMethod() *PaymentMethodInPaymentsEntityPaymentMethod`

NewPaymentMethodInPaymentsEntityPaymentMethod instantiates a new PaymentMethodInPaymentsEntityPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults

`func NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults() *PaymentMethodInPaymentsEntityPaymentMethod`

NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults instantiates a new PaymentMethodInPaymentsEntityPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCard() PaymentMethodCardInPaymentsEntityCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardOk() (*PaymentMethodCardInPaymentsEntityCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCard(v PaymentMethodCardInPaymentsEntityCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetNetbanking

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbanking() PaymentMethodNetBankingInPaymentsEntityNetbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbankingOk() (*PaymentMethodNetBankingInPaymentsEntityNetbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetNetbanking(v PaymentMethodNetBankingInPaymentsEntityNetbanking)`

SetNetbanking sets Netbanking field to given value.

### HasNetbanking

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasNetbanking() bool`

HasNetbanking returns a boolean if a field has been set.

### GetUpi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetUpi() PaymentMethodUPIInPaymentsEntityUpi`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetUpiOk() (*PaymentMethodUPIInPaymentsEntityUpi, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetUpi(v PaymentMethodUPIInPaymentsEntityUpi)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetApp

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetApp() PaymentMethodAppInPaymentsEntityApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetAppOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetApp(v PaymentMethodAppInPaymentsEntityApp)`

SetApp sets App field to given value.

### HasApp

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCardlessEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardlessEmi() PaymentMethodAppInPaymentsEntityApp`

GetCardlessEmi returns the CardlessEmi field if non-nil, zero value otherwise.

### GetCardlessEmiOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardlessEmiOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetCardlessEmiOk returns a tuple with the CardlessEmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardlessEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardlessEmi(v PaymentMethodAppInPaymentsEntityApp)`

SetCardlessEmi sets CardlessEmi field to given value.

### HasCardlessEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardlessEmi() bool`

HasCardlessEmi returns a boolean if a field has been set.

### GetPaylater

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetPaylater() PaymentMethodAppInPaymentsEntityApp`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetPaylaterOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetPaylater(v PaymentMethodAppInPaymentsEntityApp)`

SetPaylater sets Paylater field to given value.

### HasPaylater

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasPaylater() bool`

HasPaylater returns a boolean if a field has been set.

### GetEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetEmi() PaymentMethodCardEMIInPaymentsEntityEmi`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetEmiOk() (*PaymentMethodCardEMIInPaymentsEntityEmi, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetEmi(v PaymentMethodCardEMIInPaymentsEntityEmi)`

SetEmi sets Emi field to given value.

### HasEmi

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasEmi() bool`

HasEmi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


