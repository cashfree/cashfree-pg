# CreateSubscriptionRequestAuthorizationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationAmount** | Pointer to **float32** | Authorization amount for the auth payment. | [optional] 
**AuthorizationAmountRefund** | Pointer to **bool** | Indicates whether the authorization amount should be refunded to the customer automatically. Merchants can use this field to specify if the authorized funds should be returned to the customer after authorization of the subscription. | [optional] 
**PaymentMethods** | Pointer to **[]string** | Payment methods for the subscription. enach, pnach, upi, card are possible values. | [optional] 

## Methods

### NewCreateSubscriptionRequestAuthorizationDetails

`func NewCreateSubscriptionRequestAuthorizationDetails() *CreateSubscriptionRequestAuthorizationDetails`

NewCreateSubscriptionRequestAuthorizationDetails instantiates a new CreateSubscriptionRequestAuthorizationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestAuthorizationDetailsWithDefaults

`func NewCreateSubscriptionRequestAuthorizationDetailsWithDefaults() *CreateSubscriptionRequestAuthorizationDetails`

NewCreateSubscriptionRequestAuthorizationDetailsWithDefaults instantiates a new CreateSubscriptionRequestAuthorizationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationAmount

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetAuthorizationAmount() float32`

GetAuthorizationAmount returns the AuthorizationAmount field if non-nil, zero value otherwise.

### GetAuthorizationAmountOk

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetAuthorizationAmountOk() (*float32, bool)`

GetAuthorizationAmountOk returns a tuple with the AuthorizationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmount

`func (o *CreateSubscriptionRequestAuthorizationDetails) SetAuthorizationAmount(v float32)`

SetAuthorizationAmount sets AuthorizationAmount field to given value.

### HasAuthorizationAmount

`func (o *CreateSubscriptionRequestAuthorizationDetails) HasAuthorizationAmount() bool`

HasAuthorizationAmount returns a boolean if a field has been set.

### GetAuthorizationAmountRefund

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetAuthorizationAmountRefund() bool`

GetAuthorizationAmountRefund returns the AuthorizationAmountRefund field if non-nil, zero value otherwise.

### GetAuthorizationAmountRefundOk

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetAuthorizationAmountRefundOk() (*bool, bool)`

GetAuthorizationAmountRefundOk returns a tuple with the AuthorizationAmountRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmountRefund

`func (o *CreateSubscriptionRequestAuthorizationDetails) SetAuthorizationAmountRefund(v bool)`

SetAuthorizationAmountRefund sets AuthorizationAmountRefund field to given value.

### HasAuthorizationAmountRefund

`func (o *CreateSubscriptionRequestAuthorizationDetails) HasAuthorizationAmountRefund() bool`

HasAuthorizationAmountRefund returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *CreateSubscriptionRequestAuthorizationDetails) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *CreateSubscriptionRequestAuthorizationDetails) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *CreateSubscriptionRequestAuthorizationDetails) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


