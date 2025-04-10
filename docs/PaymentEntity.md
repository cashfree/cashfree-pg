# PaymentEntity

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
**OrderCurrency** | Pointer to **string** |  | [optional] 
**Authorization** | Pointer to [**AuthorizationInPaymentsEntity**](AuthorizationInPaymentsEntity.md) |  | [optional] 
**PaymentMethod** | Pointer to [**PaymentEntityPaymentMethod**](PaymentEntityPaymentMethod.md) |  | [optional] 
**InternationalPayment** | Pointer to [**InternationalPaymentEntity**](InternationalPaymentEntity.md) |  | [optional] 
**PaymentGatewayDetails** | Pointer to [**PaymentGatewayDetails**](PaymentGatewayDetails.md) |  | [optional] 
**PaymentSurcharge** | Pointer to [**PaymentEntityPaymentSurcharge**](PaymentEntityPaymentSurcharge.md) |  | [optional] 

## Methods

### NewPaymentEntity

`func NewPaymentEntity() *PaymentEntity`

NewPaymentEntity instantiates a new PaymentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEntityWithDefaults

`func NewPaymentEntityWithDefaults() *PaymentEntity`

NewPaymentEntityWithDefaults instantiates a new PaymentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *PaymentEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *PaymentEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *PaymentEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *PaymentEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetOrderId

`func (o *PaymentEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *PaymentEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *PaymentEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *PaymentEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *PaymentEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetErrorDetails

`func (o *PaymentEntity) GetErrorDetails() ErrorDetailsInPaymentsEntity`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *PaymentEntity) GetErrorDetailsOk() (*ErrorDetailsInPaymentsEntity, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *PaymentEntity) SetErrorDetails(v ErrorDetailsInPaymentsEntity)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *PaymentEntity) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetIsCaptured

`func (o *PaymentEntity) GetIsCaptured() bool`

GetIsCaptured returns the IsCaptured field if non-nil, zero value otherwise.

### GetIsCapturedOk

`func (o *PaymentEntity) GetIsCapturedOk() (*bool, bool)`

GetIsCapturedOk returns a tuple with the IsCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaptured

`func (o *PaymentEntity) SetIsCaptured(v bool)`

SetIsCaptured sets IsCaptured field to given value.

### HasIsCaptured

`func (o *PaymentEntity) HasIsCaptured() bool`

HasIsCaptured returns a boolean if a field has been set.

### GetOrderAmount

`func (o *PaymentEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *PaymentEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *PaymentEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *PaymentEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPaymentGroup

`func (o *PaymentEntity) GetPaymentGroup() string`

GetPaymentGroup returns the PaymentGroup field if non-nil, zero value otherwise.

### GetPaymentGroupOk

`func (o *PaymentEntity) GetPaymentGroupOk() (*string, bool)`

GetPaymentGroupOk returns a tuple with the PaymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroup

`func (o *PaymentEntity) SetPaymentGroup(v string)`

SetPaymentGroup sets PaymentGroup field to given value.

### HasPaymentGroup

`func (o *PaymentEntity) HasPaymentGroup() bool`

HasPaymentGroup returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *PaymentEntity) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentEntity) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentEntity) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentEntity) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentTime

`func (o *PaymentEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *PaymentEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *PaymentEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *PaymentEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentCompletionTime

`func (o *PaymentEntity) GetPaymentCompletionTime() string`

GetPaymentCompletionTime returns the PaymentCompletionTime field if non-nil, zero value otherwise.

### GetPaymentCompletionTimeOk

`func (o *PaymentEntity) GetPaymentCompletionTimeOk() (*string, bool)`

GetPaymentCompletionTimeOk returns a tuple with the PaymentCompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCompletionTime

`func (o *PaymentEntity) SetPaymentCompletionTime(v string)`

SetPaymentCompletionTime sets PaymentCompletionTime field to given value.

### HasPaymentCompletionTime

`func (o *PaymentEntity) HasPaymentCompletionTime() bool`

HasPaymentCompletionTime returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *PaymentEntity) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentEntity) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentEntity) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *PaymentEntity) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentMessage

`func (o *PaymentEntity) GetPaymentMessage() string`

GetPaymentMessage returns the PaymentMessage field if non-nil, zero value otherwise.

### GetPaymentMessageOk

`func (o *PaymentEntity) GetPaymentMessageOk() (*string, bool)`

GetPaymentMessageOk returns a tuple with the PaymentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMessage

`func (o *PaymentEntity) SetPaymentMessage(v string)`

SetPaymentMessage sets PaymentMessage field to given value.

### HasPaymentMessage

`func (o *PaymentEntity) HasPaymentMessage() bool`

HasPaymentMessage returns a boolean if a field has been set.

### GetBankReference

`func (o *PaymentEntity) GetBankReference() string`

GetBankReference returns the BankReference field if non-nil, zero value otherwise.

### GetBankReferenceOk

`func (o *PaymentEntity) GetBankReferenceOk() (*string, bool)`

GetBankReferenceOk returns a tuple with the BankReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankReference

`func (o *PaymentEntity) SetBankReference(v string)`

SetBankReference sets BankReference field to given value.

### HasBankReference

`func (o *PaymentEntity) HasBankReference() bool`

HasBankReference returns a boolean if a field has been set.

### GetAuthId

`func (o *PaymentEntity) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *PaymentEntity) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *PaymentEntity) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *PaymentEntity) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetOrderCurrency

`func (o *PaymentEntity) GetOrderCurrency() string`

GetOrderCurrency returns the OrderCurrency field if non-nil, zero value otherwise.

### GetOrderCurrencyOk

`func (o *PaymentEntity) GetOrderCurrencyOk() (*string, bool)`

GetOrderCurrencyOk returns a tuple with the OrderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCurrency

`func (o *PaymentEntity) SetOrderCurrency(v string)`

SetOrderCurrency sets OrderCurrency field to given value.

### HasOrderCurrency

`func (o *PaymentEntity) HasOrderCurrency() bool`

HasOrderCurrency returns a boolean if a field has been set.

### GetAuthorization

`func (o *PaymentEntity) GetAuthorization() AuthorizationInPaymentsEntity`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *PaymentEntity) GetAuthorizationOk() (*AuthorizationInPaymentsEntity, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *PaymentEntity) SetAuthorization(v AuthorizationInPaymentsEntity)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *PaymentEntity) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentEntity) GetPaymentMethod() PaymentEntityPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentEntity) GetPaymentMethodOk() (*PaymentEntityPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentEntity) SetPaymentMethod(v PaymentEntityPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentEntity) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetInternationalPayment

`func (o *PaymentEntity) GetInternationalPayment() InternationalPaymentEntity`

GetInternationalPayment returns the InternationalPayment field if non-nil, zero value otherwise.

### GetInternationalPaymentOk

`func (o *PaymentEntity) GetInternationalPaymentOk() (*InternationalPaymentEntity, bool)`

GetInternationalPaymentOk returns a tuple with the InternationalPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalPayment

`func (o *PaymentEntity) SetInternationalPayment(v InternationalPaymentEntity)`

SetInternationalPayment sets InternationalPayment field to given value.

### HasInternationalPayment

`func (o *PaymentEntity) HasInternationalPayment() bool`

HasInternationalPayment returns a boolean if a field has been set.

### GetPaymentGatewayDetails

`func (o *PaymentEntity) GetPaymentGatewayDetails() PaymentGatewayDetails`

GetPaymentGatewayDetails returns the PaymentGatewayDetails field if non-nil, zero value otherwise.

### GetPaymentGatewayDetailsOk

`func (o *PaymentEntity) GetPaymentGatewayDetailsOk() (*PaymentGatewayDetails, bool)`

GetPaymentGatewayDetailsOk returns a tuple with the PaymentGatewayDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayDetails

`func (o *PaymentEntity) SetPaymentGatewayDetails(v PaymentGatewayDetails)`

SetPaymentGatewayDetails sets PaymentGatewayDetails field to given value.

### HasPaymentGatewayDetails

`func (o *PaymentEntity) HasPaymentGatewayDetails() bool`

HasPaymentGatewayDetails returns a boolean if a field has been set.

### GetPaymentSurcharge

`func (o *PaymentEntity) GetPaymentSurcharge() PaymentEntityPaymentSurcharge`

GetPaymentSurcharge returns the PaymentSurcharge field if non-nil, zero value otherwise.

### GetPaymentSurchargeOk

`func (o *PaymentEntity) GetPaymentSurchargeOk() (*PaymentEntityPaymentSurcharge, bool)`

GetPaymentSurchargeOk returns a tuple with the PaymentSurcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSurcharge

`func (o *PaymentEntity) SetPaymentSurcharge(v PaymentEntityPaymentSurcharge)`

SetPaymentSurcharge sets PaymentSurcharge field to given value.

### HasPaymentSurcharge

`func (o *PaymentEntity) HasPaymentSurcharge() bool`

HasPaymentSurcharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


