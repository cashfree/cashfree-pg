# GetAllTerminalPaymentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfTerminalId** | Pointer to **int32** | Unique Cashfree terminal identifier. | [optional] 
**TerminalId** | Pointer to **string** | Merchant-defined terminal identifier. | [optional] 
**TerminalVpa** | Pointer to **string** | Virtual payment address (VPA) associated with the terminal. | [optional] 
**CfPaymentId** | Pointer to **int32** | Unique Cashfree payment identifier. | [optional] 
**PaymentAmount** | Pointer to **float32** | Payment transaction amount in the specified currency. | [optional] 
**PaymentMode** | Pointer to **string** | Payment method used for the transaction (for example, UPI_OFFLINE_STATIC). | [optional] 
**PaymentStatus** | Pointer to **string** | Current status of the payment transaction (SUCCESS, FAILED, or PENDING). | [optional] 
**PaymentTime** | Pointer to **string** | Timestamp when the payment was processed in ISO8601 format. | [optional] 
**ErrorDetails** | Pointer to [**ErrorDetailsInPaymentsEntity**](ErrorDetailsInPaymentsEntity.md) |  | [optional] 

## Methods

### NewGetAllTerminalPaymentEntity

`func NewGetAllTerminalPaymentEntity() *GetAllTerminalPaymentEntity`

NewGetAllTerminalPaymentEntity instantiates a new GetAllTerminalPaymentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllTerminalPaymentEntityWithDefaults

`func NewGetAllTerminalPaymentEntityWithDefaults() *GetAllTerminalPaymentEntity`

NewGetAllTerminalPaymentEntityWithDefaults instantiates a new GetAllTerminalPaymentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfTerminalId

`func (o *GetAllTerminalPaymentEntity) GetCfTerminalId() int32`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *GetAllTerminalPaymentEntity) GetCfTerminalIdOk() (*int32, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *GetAllTerminalPaymentEntity) SetCfTerminalId(v int32)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *GetAllTerminalPaymentEntity) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetTerminalId

`func (o *GetAllTerminalPaymentEntity) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *GetAllTerminalPaymentEntity) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *GetAllTerminalPaymentEntity) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *GetAllTerminalPaymentEntity) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalVpa

`func (o *GetAllTerminalPaymentEntity) GetTerminalVpa() string`

GetTerminalVpa returns the TerminalVpa field if non-nil, zero value otherwise.

### GetTerminalVpaOk

`func (o *GetAllTerminalPaymentEntity) GetTerminalVpaOk() (*string, bool)`

GetTerminalVpaOk returns a tuple with the TerminalVpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalVpa

`func (o *GetAllTerminalPaymentEntity) SetTerminalVpa(v string)`

SetTerminalVpa sets TerminalVpa field to given value.

### HasTerminalVpa

`func (o *GetAllTerminalPaymentEntity) HasTerminalVpa() bool`

HasTerminalVpa returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *GetAllTerminalPaymentEntity) GetCfPaymentId() int32`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *GetAllTerminalPaymentEntity) GetCfPaymentIdOk() (*int32, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *GetAllTerminalPaymentEntity) SetCfPaymentId(v int32)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *GetAllTerminalPaymentEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *GetAllTerminalPaymentEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *GetAllTerminalPaymentEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *GetAllTerminalPaymentEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *GetAllTerminalPaymentEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentMode

`func (o *GetAllTerminalPaymentEntity) GetPaymentMode() string`

GetPaymentMode returns the PaymentMode field if non-nil, zero value otherwise.

### GetPaymentModeOk

`func (o *GetAllTerminalPaymentEntity) GetPaymentModeOk() (*string, bool)`

GetPaymentModeOk returns a tuple with the PaymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMode

`func (o *GetAllTerminalPaymentEntity) SetPaymentMode(v string)`

SetPaymentMode sets PaymentMode field to given value.

### HasPaymentMode

`func (o *GetAllTerminalPaymentEntity) HasPaymentMode() bool`

HasPaymentMode returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *GetAllTerminalPaymentEntity) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *GetAllTerminalPaymentEntity) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *GetAllTerminalPaymentEntity) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *GetAllTerminalPaymentEntity) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentTime

`func (o *GetAllTerminalPaymentEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *GetAllTerminalPaymentEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *GetAllTerminalPaymentEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *GetAllTerminalPaymentEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetErrorDetails

`func (o *GetAllTerminalPaymentEntity) GetErrorDetails() ErrorDetailsInPaymentsEntity`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *GetAllTerminalPaymentEntity) GetErrorDetailsOk() (*ErrorDetailsInPaymentsEntity, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *GetAllTerminalPaymentEntity) SetErrorDetails(v ErrorDetailsInPaymentsEntity)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *GetAllTerminalPaymentEntity) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


