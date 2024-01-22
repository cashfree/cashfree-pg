# PaymentWebhookDataEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**PaymentWebhookOrderEntity**](PaymentWebhookOrderEntity.md) |  | [optional] 
**Payment** | Pointer to [**PaymentEntity**](PaymentEntity.md) |  | [optional] 
**CustomerDetails** | Pointer to [**PaymentWebhookCustomerEntity**](PaymentWebhookCustomerEntity.md) |  | [optional] 
**ErrorDetails** | Pointer to [**PaymentWebhookErrorEntity**](PaymentWebhookErrorEntity.md) |  | [optional] 
**PaymentGatewayDetails** | Pointer to [**PaymentWebhookGatewayDetailsEntity**](PaymentWebhookGatewayDetailsEntity.md) |  | [optional] 
**PaymentOffers** | Pointer to [**[]OfferEntity**](OfferEntity.md) |  | [optional] 

## Methods

### NewPaymentWebhookDataEntity

`func NewPaymentWebhookDataEntity() *PaymentWebhookDataEntity`

NewPaymentWebhookDataEntity instantiates a new PaymentWebhookDataEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWebhookDataEntityWithDefaults

`func NewPaymentWebhookDataEntityWithDefaults() *PaymentWebhookDataEntity`

NewPaymentWebhookDataEntityWithDefaults instantiates a new PaymentWebhookDataEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *PaymentWebhookDataEntity) GetOrder() PaymentWebhookOrderEntity`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentWebhookDataEntity) GetOrderOk() (*PaymentWebhookOrderEntity, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentWebhookDataEntity) SetOrder(v PaymentWebhookOrderEntity)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentWebhookDataEntity) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPayment

`func (o *PaymentWebhookDataEntity) GetPayment() PaymentEntity`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *PaymentWebhookDataEntity) GetPaymentOk() (*PaymentEntity, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *PaymentWebhookDataEntity) SetPayment(v PaymentEntity)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *PaymentWebhookDataEntity) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *PaymentWebhookDataEntity) GetCustomerDetails() PaymentWebhookCustomerEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *PaymentWebhookDataEntity) GetCustomerDetailsOk() (*PaymentWebhookCustomerEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *PaymentWebhookDataEntity) SetCustomerDetails(v PaymentWebhookCustomerEntity)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *PaymentWebhookDataEntity) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetErrorDetails

`func (o *PaymentWebhookDataEntity) GetErrorDetails() PaymentWebhookErrorEntity`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *PaymentWebhookDataEntity) GetErrorDetailsOk() (*PaymentWebhookErrorEntity, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *PaymentWebhookDataEntity) SetErrorDetails(v PaymentWebhookErrorEntity)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *PaymentWebhookDataEntity) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetPaymentGatewayDetails

`func (o *PaymentWebhookDataEntity) GetPaymentGatewayDetails() PaymentWebhookGatewayDetailsEntity`

GetPaymentGatewayDetails returns the PaymentGatewayDetails field if non-nil, zero value otherwise.

### GetPaymentGatewayDetailsOk

`func (o *PaymentWebhookDataEntity) GetPaymentGatewayDetailsOk() (*PaymentWebhookGatewayDetailsEntity, bool)`

GetPaymentGatewayDetailsOk returns a tuple with the PaymentGatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayDetails

`func (o *PaymentWebhookDataEntity) SetPaymentGatewayDetails(v PaymentWebhookGatewayDetailsEntity)`

SetPaymentGatewayDetails sets PaymentGatewayDetails field to given value.

### HasPaymentGatewayDetails

`func (o *PaymentWebhookDataEntity) HasPaymentGatewayDetails() bool`

HasPaymentGatewayDetails returns a boolean if a field has been set.

### GetPaymentOffers

`func (o *PaymentWebhookDataEntity) GetPaymentOffers() []OfferEntity`

GetPaymentOffers returns the PaymentOffers field if non-nil, zero value otherwise.

### GetPaymentOffersOk

`func (o *PaymentWebhookDataEntity) GetPaymentOffersOk() (*[]OfferEntity, bool)`

GetPaymentOffersOk returns a tuple with the PaymentOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOffers

`func (o *PaymentWebhookDataEntity) SetPaymentOffers(v []OfferEntity)`

SetPaymentOffers sets PaymentOffers field to given value.

### HasPaymentOffers

`func (o *PaymentWebhookDataEntity) HasPaymentOffers() bool`

HasPaymentOffers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


