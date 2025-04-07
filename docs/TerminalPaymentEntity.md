# TerminalPaymentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**ErrorDetails** | Pointer to [**ErrorDetailsInPaymentsEntity**](ErrorDetailsInPaymentsEntity.md) |  | [optional] 
**IsCaptured** | Pointer to **bool** |  | [optional] 
**OrderAmount** | Pointer to **float32** | Order amount can be different from payment amount if you collect service fee from the customer | [optional] 
**PaymentGroup** | Pointer to **string** | Type of payment group. One of [&#39;prepaid_card&#39;, &#39;upi_ppi_offline&#39;, &#39;cash&#39;, &#39;upi_credit_card&#39;, &#39;paypal&#39;, &#39;net_banking&#39;, &#39;cardless_emi&#39;, &#39;credit_card&#39;, &#39;bank_transfer&#39;, &#39;pay_later&#39;, &#39;debit_card_emi&#39;, &#39;debit_card&#39;, &#39;wallet&#39;, &#39;upi_ppi&#39;, &#39;upi&#39;, &#39;credit_card_emi&#39;] | [optional] 
**PaymentCurrency** | Pointer to **string** |  | [optional] 
**PaymentAmount** | Pointer to **float32** |  | [optional] 
**PaymentTime** | Pointer to **string** | This is the time when the payment was initiated | [optional] 
**PaymentCompletionTime** | Pointer to **string** | This is the time when the payment reaches its terminal state | [optional] 
**PaymentStatus** | Pointer to **string** | The transaction status can be one of  [\&quot;SUCCESS\&quot;, \&quot;NOT_ATTEMPTED\&quot;, \&quot;FAILED\&quot;, \&quot;USER_DROPPED\&quot;, \&quot;VOID\&quot;, \&quot;CANCELLED\&quot;, \&quot;PENDING\&quot;] | [optional] 
**PaymentMessage** | Pointer to **string** |  | [optional] 
**BankReference** | Pointer to **string** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**Authorization** | Pointer to [**AuthorizationInPaymentsEntity**](AuthorizationInPaymentsEntity.md) |  | [optional] 
**CustomerDetails** | Pointer to [**CustomerDetails**](CustomerDetails.md) |  | [optional] 
**PaymentMethod** | Pointer to [**PaymentEntityPaymentMethod**](PaymentEntityPaymentMethod.md) |  | [optional] 
**PaymentGatewayDetails** | Pointer to [**PaymentGatewayDetails**](PaymentGatewayDetails.md) |  | [optional] 

## Methods

### NewTerminalPaymentEntity

`func NewTerminalPaymentEntity() *TerminalPaymentEntity`

NewTerminalPaymentEntity instantiates a new TerminalPaymentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalPaymentEntityWithDefaults

`func NewTerminalPaymentEntityWithDefaults() *TerminalPaymentEntity`

NewTerminalPaymentEntityWithDefaults instantiates a new TerminalPaymentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *TerminalPaymentEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *TerminalPaymentEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *TerminalPaymentEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *TerminalPaymentEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetOrderId

`func (o *TerminalPaymentEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TerminalPaymentEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TerminalPaymentEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *TerminalPaymentEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *TerminalPaymentEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TerminalPaymentEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *TerminalPaymentEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *TerminalPaymentEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetErrorDetails

`func (o *TerminalPaymentEntity) GetErrorDetails() ErrorDetailsInPaymentsEntity`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *TerminalPaymentEntity) GetErrorDetailsOk() (*ErrorDetailsInPaymentsEntity, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *TerminalPaymentEntity) SetErrorDetails(v ErrorDetailsInPaymentsEntity)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *TerminalPaymentEntity) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetIsCaptured

`func (o *TerminalPaymentEntity) GetIsCaptured() bool`

GetIsCaptured returns the IsCaptured field if non-nil, zero value otherwise.

### GetIsCapturedOk

`func (o *TerminalPaymentEntity) GetIsCapturedOk() (*bool, bool)`

GetIsCapturedOk returns a tuple with the IsCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaptured

`func (o *TerminalPaymentEntity) SetIsCaptured(v bool)`

SetIsCaptured sets IsCaptured field to given value.

### HasIsCaptured

`func (o *TerminalPaymentEntity) HasIsCaptured() bool`

HasIsCaptured returns a boolean if a field has been set.

### GetOrderAmount

`func (o *TerminalPaymentEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *TerminalPaymentEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *TerminalPaymentEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *TerminalPaymentEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPaymentGroup

`func (o *TerminalPaymentEntity) GetPaymentGroup() string`

GetPaymentGroup returns the PaymentGroup field if non-nil, zero value otherwise.

### GetPaymentGroupOk

`func (o *TerminalPaymentEntity) GetPaymentGroupOk() (*string, bool)`

GetPaymentGroupOk returns a tuple with the PaymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroup

`func (o *TerminalPaymentEntity) SetPaymentGroup(v string)`

SetPaymentGroup sets PaymentGroup field to given value.

### HasPaymentGroup

`func (o *TerminalPaymentEntity) HasPaymentGroup() bool`

HasPaymentGroup returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *TerminalPaymentEntity) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *TerminalPaymentEntity) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *TerminalPaymentEntity) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *TerminalPaymentEntity) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *TerminalPaymentEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *TerminalPaymentEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *TerminalPaymentEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *TerminalPaymentEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentTime

`func (o *TerminalPaymentEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *TerminalPaymentEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *TerminalPaymentEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *TerminalPaymentEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentCompletionTime

`func (o *TerminalPaymentEntity) GetPaymentCompletionTime() string`

GetPaymentCompletionTime returns the PaymentCompletionTime field if non-nil, zero value otherwise.

### GetPaymentCompletionTimeOk

`func (o *TerminalPaymentEntity) GetPaymentCompletionTimeOk() (*string, bool)`

GetPaymentCompletionTimeOk returns a tuple with the PaymentCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCompletionTime

`func (o *TerminalPaymentEntity) SetPaymentCompletionTime(v string)`

SetPaymentCompletionTime sets PaymentCompletionTime field to given value.

### HasPaymentCompletionTime

`func (o *TerminalPaymentEntity) HasPaymentCompletionTime() bool`

HasPaymentCompletionTime returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *TerminalPaymentEntity) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *TerminalPaymentEntity) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *TerminalPaymentEntity) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *TerminalPaymentEntity) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentMessage

`func (o *TerminalPaymentEntity) GetPaymentMessage() string`

GetPaymentMessage returns the PaymentMessage field if non-nil, zero value otherwise.

### GetPaymentMessageOk

`func (o *TerminalPaymentEntity) GetPaymentMessageOk() (*string, bool)`

GetPaymentMessageOk returns a tuple with the PaymentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMessage

`func (o *TerminalPaymentEntity) SetPaymentMessage(v string)`

SetPaymentMessage sets PaymentMessage field to given value.

### HasPaymentMessage

`func (o *TerminalPaymentEntity) HasPaymentMessage() bool`

HasPaymentMessage returns a boolean if a field has been set.

### GetBankReference

`func (o *TerminalPaymentEntity) GetBankReference() string`

GetBankReference returns the BankReference field if non-nil, zero value otherwise.

### GetBankReferenceOk

`func (o *TerminalPaymentEntity) GetBankReferenceOk() (*string, bool)`

GetBankReferenceOk returns a tuple with the BankReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankReference

`func (o *TerminalPaymentEntity) SetBankReference(v string)`

SetBankReference sets BankReference field to given value.

### HasBankReference

`func (o *TerminalPaymentEntity) HasBankReference() bool`

HasBankReference returns a boolean if a field has been set.

### GetAuthId

`func (o *TerminalPaymentEntity) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *TerminalPaymentEntity) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *TerminalPaymentEntity) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *TerminalPaymentEntity) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetAuthorization

`func (o *TerminalPaymentEntity) GetAuthorization() AuthorizationInPaymentsEntity`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *TerminalPaymentEntity) GetAuthorizationOk() (*AuthorizationInPaymentsEntity, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *TerminalPaymentEntity) SetAuthorization(v AuthorizationInPaymentsEntity)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *TerminalPaymentEntity) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *TerminalPaymentEntity) GetCustomerDetails() CustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *TerminalPaymentEntity) GetCustomerDetailsOk() (*CustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *TerminalPaymentEntity) SetCustomerDetails(v CustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *TerminalPaymentEntity) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *TerminalPaymentEntity) GetPaymentMethod() PaymentEntityPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TerminalPaymentEntity) GetPaymentMethodOk() (*PaymentEntityPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TerminalPaymentEntity) SetPaymentMethod(v PaymentEntityPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *TerminalPaymentEntity) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentGatewayDetails

`func (o *TerminalPaymentEntity) GetPaymentGatewayDetails() PaymentGatewayDetails`

GetPaymentGatewayDetails returns the PaymentGatewayDetails field if non-nil, zero value otherwise.

### GetPaymentGatewayDetailsOk

`func (o *TerminalPaymentEntity) GetPaymentGatewayDetailsOk() (*PaymentGatewayDetails, bool)`

GetPaymentGatewayDetailsOk returns a tuple with the PaymentGatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayDetails

`func (o *TerminalPaymentEntity) SetPaymentGatewayDetails(v PaymentGatewayDetails)`

SetPaymentGatewayDetails sets PaymentGatewayDetails field to given value.

### HasPaymentGatewayDetails

`func (o *TerminalPaymentEntity) HasPaymentGatewayDetails() bool`

HasPaymentGatewayDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


