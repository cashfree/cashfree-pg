# SubscriptionPaymentRefundEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree subscription payment reference number. | [optional] 
**RefundId** | Pointer to **string** | A unique ID passed by merchant for identifying the refund. | [optional] 
**CfRefundId** | Pointer to **string** | Cashfree subscription payment refund reference number. | [optional] 
**RefundAmount** | Pointer to **float32** | The refund amount. | [optional] 
**RefundNote** | Pointer to **string** | Refund note. | [optional] 
**RefundSpeed** | Pointer to **string** | Refund speed. Can be INSTANT or NORMAL. | [optional] 
**RefundStatus** | Pointer to **string** | Status of the refund. | [optional] 

## Methods

### NewSubscriptionPaymentRefundEntity

`func NewSubscriptionPaymentRefundEntity() *SubscriptionPaymentRefundEntity`

NewSubscriptionPaymentRefundEntity instantiates a new SubscriptionPaymentRefundEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPaymentRefundEntityWithDefaults

`func NewSubscriptionPaymentRefundEntityWithDefaults() *SubscriptionPaymentRefundEntity`

NewSubscriptionPaymentRefundEntityWithDefaults instantiates a new SubscriptionPaymentRefundEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *SubscriptionPaymentRefundEntity) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *SubscriptionPaymentRefundEntity) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *SubscriptionPaymentRefundEntity) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *SubscriptionPaymentRefundEntity) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *SubscriptionPaymentRefundEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SubscriptionPaymentRefundEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SubscriptionPaymentRefundEntity) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SubscriptionPaymentRefundEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetRefundId

`func (o *SubscriptionPaymentRefundEntity) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *SubscriptionPaymentRefundEntity) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *SubscriptionPaymentRefundEntity) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *SubscriptionPaymentRefundEntity) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetCfRefundId

`func (o *SubscriptionPaymentRefundEntity) GetCfRefundId() string`

GetCfRefundId returns the CfRefundId field if non-nil, zero value otherwise.

### GetCfRefundIdOk

`func (o *SubscriptionPaymentRefundEntity) GetCfRefundIdOk() (*string, bool)`

GetCfRefundIdOk returns a tuple with the CfRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfRefundId

`func (o *SubscriptionPaymentRefundEntity) SetCfRefundId(v string)`

SetCfRefundId sets CfRefundId field to given value.

### HasCfRefundId

`func (o *SubscriptionPaymentRefundEntity) HasCfRefundId() bool`

HasCfRefundId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *SubscriptionPaymentRefundEntity) GetRefundAmount() float32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *SubscriptionPaymentRefundEntity) GetRefundAmountOk() (*float32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *SubscriptionPaymentRefundEntity) SetRefundAmount(v float32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *SubscriptionPaymentRefundEntity) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundNote

`func (o *SubscriptionPaymentRefundEntity) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *SubscriptionPaymentRefundEntity) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *SubscriptionPaymentRefundEntity) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *SubscriptionPaymentRefundEntity) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundSpeed

`func (o *SubscriptionPaymentRefundEntity) GetRefundSpeed() string`

GetRefundSpeed returns the RefundSpeed field if non-nil, zero value otherwise.

### GetRefundSpeedOk

`func (o *SubscriptionPaymentRefundEntity) GetRefundSpeedOk() (*string, bool)`

GetRefundSpeedOk returns a tuple with the RefundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSpeed

`func (o *SubscriptionPaymentRefundEntity) SetRefundSpeed(v string)`

SetRefundSpeed sets RefundSpeed field to given value.

### HasRefundSpeed

`func (o *SubscriptionPaymentRefundEntity) HasRefundSpeed() bool`

HasRefundSpeed returns a boolean if a field has been set.

### GetRefundStatus

`func (o *SubscriptionPaymentRefundEntity) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *SubscriptionPaymentRefundEntity) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *SubscriptionPaymentRefundEntity) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.

### HasRefundStatus

`func (o *SubscriptionPaymentRefundEntity) HasRefundStatus() bool`

HasRefundStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


