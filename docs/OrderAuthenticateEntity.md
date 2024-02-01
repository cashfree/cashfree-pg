# OrderAuthenticateEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** | The payment id for which this request was sent | [optional] 
**Action** | Pointer to **string** | The action that was invoked for this request. | [optional] 
**AuthenticateStatus** | Pointer to **string** | Status of the is action. Will be either failed or successful. If the action is successful, you should still call the authorization status to verify the final payment status. | [optional] 
**PaymentMessage** | Pointer to **string** | Human readable message which describes the status in more detail | [optional] 

## Methods

### NewOrderAuthenticateEntity

`func NewOrderAuthenticateEntity() *OrderAuthenticateEntity`

NewOrderAuthenticateEntity instantiates a new OrderAuthenticateEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAuthenticateEntityWithDefaults

`func NewOrderAuthenticateEntityWithDefaults() *OrderAuthenticateEntity`

NewOrderAuthenticateEntityWithDefaults instantiates a new OrderAuthenticateEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *OrderAuthenticateEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *OrderAuthenticateEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *OrderAuthenticateEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *OrderAuthenticateEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetAction

`func (o *OrderAuthenticateEntity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderAuthenticateEntity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderAuthenticateEntity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderAuthenticateEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuthenticateStatus

`func (o *OrderAuthenticateEntity) GetAuthenticateStatus() string`

GetAuthenticateStatus returns the AuthenticateStatus field if non-nil, zero value otherwise.

### GetAuthenticateStatusOk

`func (o *OrderAuthenticateEntity) GetAuthenticateStatusOk() (*string, bool)`

GetAuthenticateStatusOk returns a tuple with the AuthenticateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticateStatus

`func (o *OrderAuthenticateEntity) SetAuthenticateStatus(v string)`

SetAuthenticateStatus sets AuthenticateStatus field to given value.

### HasAuthenticateStatus

`func (o *OrderAuthenticateEntity) HasAuthenticateStatus() bool`

HasAuthenticateStatus returns a boolean if a field has been set.

### GetPaymentMessage

`func (o *OrderAuthenticateEntity) GetPaymentMessage() string`

GetPaymentMessage returns the PaymentMessage field if non-nil, zero value otherwise.

### GetPaymentMessageOk

`func (o *OrderAuthenticateEntity) GetPaymentMessageOk() (*string, bool)`

GetPaymentMessageOk returns a tuple with the PaymentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMessage

`func (o *OrderAuthenticateEntity) SetPaymentMessage(v string)`

SetPaymentMessage sets PaymentMessage field to given value.

### HasPaymentMessage

`func (o *OrderAuthenticateEntity) HasPaymentMessage() bool`

HasPaymentMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


