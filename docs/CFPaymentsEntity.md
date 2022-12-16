# CFPaymentsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **int32** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**IsCaptured** | Pointer to **bool** |  | [optional] 
**OrderAmount** | Pointer to **float32** | Order amount can be different from payment amount if you collect service fee from the customer | [optional] 
**PaymentGroup** | Pointer to **string** | Type of payment group. One of [&#39;upi&#39;, &#39;card&#39;, &#39;app&#39;, &#39;netbanking&#39;, &#39;paylater&#39;, &#39;cardless_emi&#39;] | [optional] 
**PaymentCurrency** | Pointer to **string** |  | [optional] 
**PaymentAmount** | Pointer to **float32** |  | [optional] 
**PaymentTime** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** | The transaction status can be one of  [\&quot;SUCCESS\&quot;, \&quot;NOT_ATTEMPTED\&quot;, \&quot;FAILED\&quot;, \&quot;USER_DROPPED\&quot;, \&quot;VOID\&quot;, \&quot;CANCELLED\&quot;, \&quot;PENDING\&quot;] | [optional] 
**PaymentMessage** | Pointer to **string** |  | [optional] 
**BankReference** | Pointer to **string** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**Authorization** | Pointer to [**CFAuthorizationInPaymentsEntity**](CFAuthorizationInPaymentsEntity.md) |  | [optional] 
**PaymentMethod** | Pointer to [**CFPaymentsEntityPayment**](CFPaymentsEntityPayment.md) |  | [optional] 

## Methods

### NewCFPaymentsEntity

`func NewCFPaymentsEntity() *CFPaymentsEntity`

NewCFPaymentsEntity instantiates a new CFPaymentsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFPaymentsEntityWithDefaults

`func NewCFPaymentsEntityWithDefaults() *CFPaymentsEntity`

NewCFPaymentsEntityWithDefaults instantiates a new CFPaymentsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CFPaymentsEntity) GetCfPaymentId() int32`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CFPaymentsEntity) GetCfPaymentIdOk() (*int32, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CFPaymentsEntity) SetCfPaymentId(v int32)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CFPaymentsEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetOrderId

`func (o *CFPaymentsEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFPaymentsEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFPaymentsEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFPaymentsEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetEntity

`func (o *CFPaymentsEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CFPaymentsEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CFPaymentsEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CFPaymentsEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetIsCaptured

`func (o *CFPaymentsEntity) GetIsCaptured() bool`

GetIsCaptured returns the IsCaptured field if non-nil, zero value otherwise.

### GetIsCapturedOk

`func (o *CFPaymentsEntity) GetIsCapturedOk() (*bool, bool)`

GetIsCapturedOk returns a tuple with the IsCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaptured

`func (o *CFPaymentsEntity) SetIsCaptured(v bool)`

SetIsCaptured sets IsCaptured field to given value.

### HasIsCaptured

`func (o *CFPaymentsEntity) HasIsCaptured() bool`

HasIsCaptured returns a boolean if a field has been set.

### GetOrderAmount

`func (o *CFPaymentsEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *CFPaymentsEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *CFPaymentsEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *CFPaymentsEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPaymentGroup

`func (o *CFPaymentsEntity) GetPaymentGroup() string`

GetPaymentGroup returns the PaymentGroup field if non-nil, zero value otherwise.

### GetPaymentGroupOk

`func (o *CFPaymentsEntity) GetPaymentGroupOk() (*string, bool)`

GetPaymentGroupOk returns a tuple with the PaymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroup

`func (o *CFPaymentsEntity) SetPaymentGroup(v string)`

SetPaymentGroup sets PaymentGroup field to given value.

### HasPaymentGroup

`func (o *CFPaymentsEntity) HasPaymentGroup() bool`

HasPaymentGroup returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *CFPaymentsEntity) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *CFPaymentsEntity) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *CFPaymentsEntity) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *CFPaymentsEntity) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *CFPaymentsEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *CFPaymentsEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *CFPaymentsEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *CFPaymentsEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentTime

`func (o *CFPaymentsEntity) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *CFPaymentsEntity) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *CFPaymentsEntity) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *CFPaymentsEntity) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *CFPaymentsEntity) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CFPaymentsEntity) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CFPaymentsEntity) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *CFPaymentsEntity) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentMessage

`func (o *CFPaymentsEntity) GetPaymentMessage() string`

GetPaymentMessage returns the PaymentMessage field if non-nil, zero value otherwise.

### GetPaymentMessageOk

`func (o *CFPaymentsEntity) GetPaymentMessageOk() (*string, bool)`

GetPaymentMessageOk returns a tuple with the PaymentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMessage

`func (o *CFPaymentsEntity) SetPaymentMessage(v string)`

SetPaymentMessage sets PaymentMessage field to given value.

### HasPaymentMessage

`func (o *CFPaymentsEntity) HasPaymentMessage() bool`

HasPaymentMessage returns a boolean if a field has been set.

### GetBankReference

`func (o *CFPaymentsEntity) GetBankReference() string`

GetBankReference returns the BankReference field if non-nil, zero value otherwise.

### GetBankReferenceOk

`func (o *CFPaymentsEntity) GetBankReferenceOk() (*string, bool)`

GetBankReferenceOk returns a tuple with the BankReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankReference

`func (o *CFPaymentsEntity) SetBankReference(v string)`

SetBankReference sets BankReference field to given value.

### HasBankReference

`func (o *CFPaymentsEntity) HasBankReference() bool`

HasBankReference returns a boolean if a field has been set.

### GetAuthId

`func (o *CFPaymentsEntity) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *CFPaymentsEntity) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *CFPaymentsEntity) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *CFPaymentsEntity) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetAuthorization

`func (o *CFPaymentsEntity) GetAuthorization() CFAuthorizationInPaymentsEntity`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *CFPaymentsEntity) GetAuthorizationOk() (*CFAuthorizationInPaymentsEntity, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *CFPaymentsEntity) SetAuthorization(v CFAuthorizationInPaymentsEntity)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *CFPaymentsEntity) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CFPaymentsEntity) GetPaymentMethod() CFPaymentsEntityPayment`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CFPaymentsEntity) GetPaymentMethodOk() (*CFPaymentsEntityPayment, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CFPaymentsEntity) SetPaymentMethod(v CFPaymentsEntityPayment)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CFPaymentsEntity) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


