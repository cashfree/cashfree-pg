# SubscriptionPaymentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree subscription payment reference number | [optional] 
**CfSubscriptionId** | Pointer to **string** | Cashfree subscription reference number | [optional] 
**CfTxnId** | Pointer to **string** | Cashfree subscription payment transaction ID | [optional] 
**CfOrderId** | Pointer to **string** | Cashfree subscription payment order ID | [optional] 
**FailureDetails** | Pointer to [**CreateSubscriptionPaymentAuthResponseFailureDetails**](CreateSubscriptionPaymentAuthResponseFailureDetails.md) |  | [optional] 
**PaymentAmount** | Pointer to **float32** | The charge amount of the payment. | [optional] 
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**PaymentInitiatedDate** | Pointer to **string** | The date on which the payment was initiated. | [optional] 
**PaymentRemarks** | Pointer to **string** | Payment remarks. | [optional] 
**PaymentScheduleDate** | Pointer to **string** | The date on which the payment is scheduled to be processed. | [optional] 
**PaymentStatus** | Pointer to **string** | Status of the payment. | [optional] 
**PaymentType** | Pointer to **string** | Payment type. Can be AUTH or CHARGE. | [optional] 
**RetryAttempts** | Pointer to **int32** | Retry attempts. | [optional] 
**SubscriptionId** | Pointer to **string** | A unique ID passed by merchant for identifying the subscription. | [optional] 

## Methods

### NewSubscriptionPaymentEntity

`func NewSubscriptionPaymentEntity() *SubscriptionPaymentEntity`

NewSubscriptionPaymentEntity instantiates a new SubscriptionPaymentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPaymentEntityWithDefaults

`func NewSubscriptionPaymentEntityWithDefaults() *SubscriptionPaymentEntity`

NewSubscriptionPaymentEntityWithDefaults instantiates a new SubscriptionPaymentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationDetails

`func (o *SubscriptionPaymentEntity) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *SubscriptionPaymentEntity) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *SubscriptionPaymentEntity) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *SubscriptionPaymentEntity) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *SubscriptionPaymentEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SubscriptionPaymentEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SubscriptionPaymentEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SubscriptionPaymentEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSubscriptionId

`func (o *SubscriptionPaymentEntity) GetCfSubscriptionId() string`

GetCfSubscriptionId returns the CfSubscriptionId field if non-nil, zero value otherwise.

### GetCfSubscriptionIdOk

`func (o *SubscriptionPaymentEntity) GetCfSubscriptionIdOk() (*string, bool)`

GetCfSubscriptionIdOk returns a tuple with the CfSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSubscriptionId

`func (o *SubscriptionPaymentEntity) SetCfSubscriptionId(v string)`

SetCfSubscriptionId sets CfSubscriptionId field to given value.

### HasCfSubscriptionId

`func (o *SubscriptionPaymentEntity) HasCfSubscriptionId() bool`

HasCfSubscriptionId returns a boolean if a field has been set.

### GetCfTxnId

`func (o *SubscriptionPaymentEntity) GetCfTxnId() string`

GetCfTxnId returns the CfTxnId field if non-nil, zero value otherwise.

### GetCfTxnIdOk

`func (o *SubscriptionPaymentEntity) GetCfTxnIdOk() (*string, bool)`

GetCfTxnIdOk returns a tuple with the CfTxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTxnId

`func (o *SubscriptionPaymentEntity) SetCfTxnId(v string)`

SetCfTxnId sets CfTxnId field to given value.

### HasCfTxnId

`func (o *SubscriptionPaymentEntity) HasCfTxnId() bool`

HasCfTxnId returns a boolean if a field has been set.

### GetCfOrderId

`func (o *SubscriptionPaymentEntity) GetCfOrderId() string`

GetCfOrderId returns the CfOrderId field if non-nil, zero value otherwise.

### GetCfOrderIdOk

`func (o *SubscriptionPaymentEntity) GetCfOrderIdOk() (*string, bool)`

GetCfOrderIdOk returns a tuple with the CfOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfOrderId

`func (o *SubscriptionPaymentEntity) SetCfOrderId(v string)`

SetCfOrderId sets CfOrderId field to given value.

### HasCfOrderId

`func (o *SubscriptionPaymentEntity) HasCfOrderId() bool`

HasCfOrderId returns a boolean if a field has been set.

### GetFailureDetails

`func (o *SubscriptionPaymentEntity) GetFailureDetails() CreateSubscriptionPaymentAuthResponseFailureDetails`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *SubscriptionPaymentEntity) GetFailureDetailsOk() (*CreateSubscriptionPaymentAuthResponseFailureDetails, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *SubscriptionPaymentEntity) SetFailureDetails(v CreateSubscriptionPaymentAuthResponseFailureDetails)`

SetFailureDetails sets FailureDetails field to given value.

### HasFailureDetails

`func (o *SubscriptionPaymentEntity) HasFailureDetails() bool`

HasFailureDetails returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *SubscriptionPaymentEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *SubscriptionPaymentEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *SubscriptionPaymentEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *SubscriptionPaymentEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *SubscriptionPaymentEntity) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *SubscriptionPaymentEntity) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *SubscriptionPaymentEntity) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *SubscriptionPaymentEntity) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentInitiatedDate

`func (o *SubscriptionPaymentEntity) GetPaymentInitiatedDate() string`

GetPaymentInitiatedDate returns the PaymentInitiatedDate field if non-nil, zero value otherwise.

### GetPaymentInitiatedDateOk

`func (o *SubscriptionPaymentEntity) GetPaymentInitiatedDateOk() (*string, bool)`

GetPaymentInitiatedDateOk returns a tuple with the PaymentInitiatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInitiatedDate

`func (o *SubscriptionPaymentEntity) SetPaymentInitiatedDate(v string)`

SetPaymentInitiatedDate sets PaymentInitiatedDate field to given value.

### HasPaymentInitiatedDate

`func (o *SubscriptionPaymentEntity) HasPaymentInitiatedDate() bool`

HasPaymentInitiatedDate returns a boolean if a field has been set.

### GetPaymentRemarks

`func (o *SubscriptionPaymentEntity) GetPaymentRemarks() string`

GetPaymentRemarks returns the PaymentRemarks field if non-nil, zero value otherwise.

### GetPaymentRemarksOk

`func (o *SubscriptionPaymentEntity) GetPaymentRemarksOk() (*string, bool)`

GetPaymentRemarksOk returns a tuple with the PaymentRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRemarks

`func (o *SubscriptionPaymentEntity) SetPaymentRemarks(v string)`

SetPaymentRemarks sets PaymentRemarks field to given value.

### HasPaymentRemarks

`func (o *SubscriptionPaymentEntity) HasPaymentRemarks() bool`

HasPaymentRemarks returns a boolean if a field has been set.

### GetPaymentScheduleDate

`func (o *SubscriptionPaymentEntity) GetPaymentScheduleDate() string`

GetPaymentScheduleDate returns the PaymentScheduleDate field if non-nil, zero value otherwise.

### GetPaymentScheduleDateOk

`func (o *SubscriptionPaymentEntity) GetPaymentScheduleDateOk() (*string, bool)`

GetPaymentScheduleDateOk returns a tuple with the PaymentScheduleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleDate

`func (o *SubscriptionPaymentEntity) SetPaymentScheduleDate(v string)`

SetPaymentScheduleDate sets PaymentScheduleDate field to given value.

### HasPaymentScheduleDate

`func (o *SubscriptionPaymentEntity) HasPaymentScheduleDate() bool`

HasPaymentScheduleDate returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *SubscriptionPaymentEntity) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *SubscriptionPaymentEntity) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *SubscriptionPaymentEntity) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *SubscriptionPaymentEntity) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentType

`func (o *SubscriptionPaymentEntity) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *SubscriptionPaymentEntity) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *SubscriptionPaymentEntity) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *SubscriptionPaymentEntity) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetRetryAttempts

`func (o *SubscriptionPaymentEntity) GetRetryAttempts() int32`

GetRetryAttempts returns the RetryAttempts field if non-nil, zero value otherwise.

### GetRetryAttemptsOk

`func (o *SubscriptionPaymentEntity) GetRetryAttemptsOk() (*int32, bool)`

GetRetryAttemptsOk returns a tuple with the RetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttempts

`func (o *SubscriptionPaymentEntity) SetRetryAttempts(v int32)`

SetRetryAttempts sets RetryAttempts field to given value.

### HasRetryAttempts

`func (o *SubscriptionPaymentEntity) HasRetryAttempts() bool`

HasRetryAttempts returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionPaymentEntity) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionPaymentEntity) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionPaymentEntity) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionPaymentEntity) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


