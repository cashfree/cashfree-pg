# PaymentEntityPaymentMethod

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

### NewPaymentEntityPaymentMethod

`func NewPaymentEntityPaymentMethod() *PaymentEntityPaymentMethod`

NewPaymentEntityPaymentMethod instantiates a new PaymentEntityPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEntityPaymentMethodWithDefaults

`func NewPaymentEntityPaymentMethodWithDefaults() *PaymentEntityPaymentMethod`

NewPaymentEntityPaymentMethodWithDefaults instantiates a new PaymentEntityPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *PaymentEntityPaymentMethod) GetCard() PaymentMethodCardInPaymentsEntityCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentEntityPaymentMethod) GetCardOk() (*PaymentMethodCardInPaymentsEntityCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentEntityPaymentMethod) SetCard(v PaymentMethodCardInPaymentsEntityCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentEntityPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetNetbanking

`func (o *PaymentEntityPaymentMethod) GetNetbanking() PaymentMethodNetBankingInPaymentsEntityNetbanking`

GetNetbanking returns the Netbanking field if non-nil, zero value otherwise.

### GetNetbankingOk

`func (o *PaymentEntityPaymentMethod) GetNetbankingOk() (*PaymentMethodNetBankingInPaymentsEntityNetbanking, bool)`

GetNetbankingOk returns a tuple with the Netbanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbanking

`func (o *PaymentEntityPaymentMethod) SetNetbanking(v PaymentMethodNetBankingInPaymentsEntityNetbanking)`

SetNetbanking sets Netbanking field to given value.

### HasNetbanking

`func (o *PaymentEntityPaymentMethod) HasNetbanking() bool`

HasNetbanking returns a boolean if a field has been set.

### GetUpi

`func (o *PaymentEntityPaymentMethod) GetUpi() PaymentMethodUPIInPaymentsEntityUpi`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *PaymentEntityPaymentMethod) GetUpiOk() (*PaymentMethodUPIInPaymentsEntityUpi, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *PaymentEntityPaymentMethod) SetUpi(v PaymentMethodUPIInPaymentsEntityUpi)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *PaymentEntityPaymentMethod) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetApp

`func (o *PaymentEntityPaymentMethod) GetApp() PaymentMethodAppInPaymentsEntityApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PaymentEntityPaymentMethod) GetAppOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PaymentEntityPaymentMethod) SetApp(v PaymentMethodAppInPaymentsEntityApp)`

SetApp sets App field to given value.

### HasApp

`func (o *PaymentEntityPaymentMethod) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCardlessEmi

`func (o *PaymentEntityPaymentMethod) GetCardlessEmi() PaymentMethodAppInPaymentsEntityApp`

GetCardlessEmi returns the CardlessEmi field if non-nil, zero value otherwise.

### GetCardlessEmiOk

`func (o *PaymentEntityPaymentMethod) GetCardlessEmiOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetCardlessEmiOk returns a tuple with the CardlessEmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardlessEmi

`func (o *PaymentEntityPaymentMethod) SetCardlessEmi(v PaymentMethodAppInPaymentsEntityApp)`

SetCardlessEmi sets CardlessEmi field to given value.

### HasCardlessEmi

`func (o *PaymentEntityPaymentMethod) HasCardlessEmi() bool`

HasCardlessEmi returns a boolean if a field has been set.

### GetPaylater

`func (o *PaymentEntityPaymentMethod) GetPaylater() PaymentMethodAppInPaymentsEntityApp`

GetPaylater returns the Paylater field if non-nil, zero value otherwise.

### GetPaylaterOk

`func (o *PaymentEntityPaymentMethod) GetPaylaterOk() (*PaymentMethodAppInPaymentsEntityApp, bool)`

GetPaylaterOk returns a tuple with the Paylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaylater

`func (o *PaymentEntityPaymentMethod) SetPaylater(v PaymentMethodAppInPaymentsEntityApp)`

SetPaylater sets Paylater field to given value.

### HasPaylater

`func (o *PaymentEntityPaymentMethod) HasPaylater() bool`

HasPaylater returns a boolean if a field has been set.

### GetEmi

`func (o *PaymentEntityPaymentMethod) GetEmi() PaymentMethodCardEMIInPaymentsEntityEmi`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *PaymentEntityPaymentMethod) GetEmiOk() (*PaymentMethodCardEMIInPaymentsEntityEmi, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *PaymentEntityPaymentMethod) SetEmi(v PaymentMethodCardEMIInPaymentsEntityEmi)`

SetEmi sets Emi field to given value.

### HasEmi

`func (o *PaymentEntityPaymentMethod) HasEmi() bool`

HasEmi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


