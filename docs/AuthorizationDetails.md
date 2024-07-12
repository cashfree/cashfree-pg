# AuthorizationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationAmount** | Pointer to **float32** | Authorization amount for the auth payment. | [optional] 
**AuthorizationAmountRefund** | Pointer to **bool** | Indicates whether the authorization amount should be refunded to the customer automatically. Merchants can use this field to specify if the authorized funds should be returned to the customer after authorization of the subscription. | [optional] 
**AuthorizationReference** | Pointer to **string** | Authorization reference. UMN for UPI, UMRN for EMandate/Physical Mandate and Enrollment ID for cards. | [optional] 
**AuthorizationTime** | Pointer to **string** | Authorization time. | [optional] 
**AuthorizationStatus** | Pointer to **string** | Status of the authorization. | [optional] 
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**PaymentMethod** | Pointer to **string** | Payment method used for the authorization. | [optional] 

## Methods

### NewAuthorizationDetails

`func NewAuthorizationDetails() *AuthorizationDetails`

NewAuthorizationDetails instantiates a new AuthorizationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailsWithDefaults

`func NewAuthorizationDetailsWithDefaults() *AuthorizationDetails`

NewAuthorizationDetailsWithDefaults instantiates a new AuthorizationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationAmount

`func (o *AuthorizationDetails) GetAuthorizationAmount() float32`

GetAuthorizationAmount returns the AuthorizationAmount field if non-nil, zero value otherwise.

### GetAuthorizationAmountOk

`func (o *AuthorizationDetails) GetAuthorizationAmountOk() (*float32, bool)`

GetAuthorizationAmountOk returns a tuple with the AuthorizationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmount

`func (o *AuthorizationDetails) SetAuthorizationAmount(v float32)`

SetAuthorizationAmount sets AuthorizationAmount field to given value.

### HasAuthorizationAmount

`func (o *AuthorizationDetails) HasAuthorizationAmount() bool`

HasAuthorizationAmount returns a boolean if a field has been set.

### GetAuthorizationAmountRefund

`func (o *AuthorizationDetails) GetAuthorizationAmountRefund() bool`

GetAuthorizationAmountRefund returns the AuthorizationAmountRefund field if non-nil, zero value otherwise.

### GetAuthorizationAmountRefundOk

`func (o *AuthorizationDetails) GetAuthorizationAmountRefundOk() (*bool, bool)`

GetAuthorizationAmountRefundOk returns a tuple with the AuthorizationAmountRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmountRefund

`func (o *AuthorizationDetails) SetAuthorizationAmountRefund(v bool)`

SetAuthorizationAmountRefund sets AuthorizationAmountRefund field to given value.

### HasAuthorizationAmountRefund

`func (o *AuthorizationDetails) HasAuthorizationAmountRefund() bool`

HasAuthorizationAmountRefund returns a boolean if a field has been set.

### GetAuthorizationReference

`func (o *AuthorizationDetails) GetAuthorizationReference() string`

GetAuthorizationReference returns the AuthorizationReference field if non-nil, zero value otherwise.

### GetAuthorizationReferenceOk

`func (o *AuthorizationDetails) GetAuthorizationReferenceOk() (*string, bool)`

GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationReference

`func (o *AuthorizationDetails) SetAuthorizationReference(v string)`

SetAuthorizationReference sets AuthorizationReference field to given value.

### HasAuthorizationReference

`func (o *AuthorizationDetails) HasAuthorizationReference() bool`

HasAuthorizationReference returns a boolean if a field has been set.

### GetAuthorizationTime

`func (o *AuthorizationDetails) GetAuthorizationTime() string`

GetAuthorizationTime returns the AuthorizationTime field if non-nil, zero value otherwise.

### GetAuthorizationTimeOk

`func (o *AuthorizationDetails) GetAuthorizationTimeOk() (*string, bool)`

GetAuthorizationTimeOk returns a tuple with the AuthorizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationTime

`func (o *AuthorizationDetails) SetAuthorizationTime(v string)`

SetAuthorizationTime sets AuthorizationTime field to given value.

### HasAuthorizationTime

`func (o *AuthorizationDetails) HasAuthorizationTime() bool`

HasAuthorizationTime returns a boolean if a field has been set.

### GetAuthorizationStatus

`func (o *AuthorizationDetails) GetAuthorizationStatus() string`

GetAuthorizationStatus returns the AuthorizationStatus field if non-nil, zero value otherwise.

### GetAuthorizationStatusOk

`func (o *AuthorizationDetails) GetAuthorizationStatusOk() (*string, bool)`

GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatus

`func (o *AuthorizationDetails) SetAuthorizationStatus(v string)`

SetAuthorizationStatus sets AuthorizationStatus field to given value.

### HasAuthorizationStatus

`func (o *AuthorizationDetails) HasAuthorizationStatus() bool`

HasAuthorizationStatus returns a boolean if a field has been set.

### GetPaymentId

`func (o *AuthorizationDetails) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *AuthorizationDetails) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *AuthorizationDetails) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *AuthorizationDetails) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *AuthorizationDetails) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AuthorizationDetails) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AuthorizationDetails) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *AuthorizationDetails) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


