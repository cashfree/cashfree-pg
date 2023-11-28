# CreateTerminalTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfOrderId** | **int32** | cashfree order ID that was returned while creating an order. | 
**CfTerminalId** | Pointer to **int32** | cashfree terminal id. this is a required parameter when you do not provide the terminal phone number. | [optional] 
**PaymentMethod** | **string** | mention the payment method used for the transaction. possible values - QR_CODE, LINK. | 
**TerminalPhoneNo** | Pointer to **string** | agent mobile number assigned to the terminal. this is a required parameter when you do not provide the cf_terminal_id. | [optional] 

## Methods

### NewCreateTerminalTransactionRequest

`func NewCreateTerminalTransactionRequest(cfOrderId int32, paymentMethod string, ) *CreateTerminalTransactionRequest`

NewCreateTerminalTransactionRequest instantiates a new CreateTerminalTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTerminalTransactionRequestWithDefaults

`func NewCreateTerminalTransactionRequestWithDefaults() *CreateTerminalTransactionRequest`

NewCreateTerminalTransactionRequestWithDefaults instantiates a new CreateTerminalTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfOrderId

`func (o *CreateTerminalTransactionRequest) GetCfOrderId() int32`

GetCfOrderId returns the CfOrderId field if non-nil, zero value otherwise.

### GetCfOrderIdOk

`func (o *CreateTerminalTransactionRequest) GetCfOrderIdOk() (*int32, bool)`

GetCfOrderIdOk returns a tuple with the CfOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfOrderId

`func (o *CreateTerminalTransactionRequest) SetCfOrderId(v int32)`

SetCfOrderId sets CfOrderId field to given value.


### GetCfTerminalId

`func (o *CreateTerminalTransactionRequest) GetCfTerminalId() int32`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *CreateTerminalTransactionRequest) GetCfTerminalIdOk() (*int32, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *CreateTerminalTransactionRequest) SetCfTerminalId(v int32)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *CreateTerminalTransactionRequest) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreateTerminalTransactionRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateTerminalTransactionRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateTerminalTransactionRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetTerminalPhoneNo

`func (o *CreateTerminalTransactionRequest) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *CreateTerminalTransactionRequest) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *CreateTerminalTransactionRequest) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.

### HasTerminalPhoneNo

`func (o *CreateTerminalTransactionRequest) HasTerminalPhoneNo() bool`

HasTerminalPhoneNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


