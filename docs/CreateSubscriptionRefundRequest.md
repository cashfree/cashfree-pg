# CreateSubscriptionRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | A unique ID passed by merchant for identifying the subscription. | 
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree subscription payment reference number. | [optional] 
**RefundId** | **string** | A unique ID passed by merchant for identifying the refund. | 
**RefundAmount** | **float32** | The amount to be refunded. Can be partial or full amount of the payment. | 
**RefundNote** | Pointer to **string** | Refund note. | [optional] 
**RefundSpeed** | Pointer to **string** | Refund speed. Can be INSTANT or STANDARD. UPI supports only STANDARD refunds, Enach and Pnach supports only INSTANT refunds. | [optional] 

## Methods

### NewCreateSubscriptionRefundRequest

`func NewCreateSubscriptionRefundRequest(subscriptionId string, refundId string, refundAmount float32, ) *CreateSubscriptionRefundRequest`

NewCreateSubscriptionRefundRequest instantiates a new CreateSubscriptionRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRefundRequestWithDefaults

`func NewCreateSubscriptionRefundRequestWithDefaults() *CreateSubscriptionRefundRequest`

NewCreateSubscriptionRefundRequestWithDefaults instantiates a new CreateSubscriptionRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateSubscriptionRefundRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionRefundRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionRefundRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPaymentId

`func (o *CreateSubscriptionRefundRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateSubscriptionRefundRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateSubscriptionRefundRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *CreateSubscriptionRefundRequest) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *CreateSubscriptionRefundRequest) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CreateSubscriptionRefundRequest) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CreateSubscriptionRefundRequest) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CreateSubscriptionRefundRequest) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetRefundId

`func (o *CreateSubscriptionRefundRequest) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *CreateSubscriptionRefundRequest) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *CreateSubscriptionRefundRequest) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetRefundAmount

`func (o *CreateSubscriptionRefundRequest) GetRefundAmount() float32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *CreateSubscriptionRefundRequest) GetRefundAmountOk() (*float32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *CreateSubscriptionRefundRequest) SetRefundAmount(v float32)`

SetRefundAmount sets RefundAmount field to given value.


### GetRefundNote

`func (o *CreateSubscriptionRefundRequest) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *CreateSubscriptionRefundRequest) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *CreateSubscriptionRefundRequest) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *CreateSubscriptionRefundRequest) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundSpeed

`func (o *CreateSubscriptionRefundRequest) GetRefundSpeed() string`

GetRefundSpeed returns the RefundSpeed field if non-nil, zero value otherwise.

### GetRefundSpeedOk

`func (o *CreateSubscriptionRefundRequest) GetRefundSpeedOk() (*string, bool)`

GetRefundSpeedOk returns a tuple with the RefundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSpeed

`func (o *CreateSubscriptionRefundRequest) SetRefundSpeed(v string)`

SetRefundSpeed sets RefundSpeed field to given value.

### HasRefundSpeed

`func (o *CreateSubscriptionRefundRequest) HasRefundSpeed() bool`

HasRefundSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


