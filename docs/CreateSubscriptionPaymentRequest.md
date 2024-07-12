# CreateSubscriptionPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | A unique ID passed by merchant for identifying the subscription. | 
**SubscriptionSessionId** | Pointer to **string** | Session ID for the subscription. Required only for Auth. | [optional] 
**PaymentId** | **string** | A unique ID passed by merchant for identifying the subscription payment. | 
**PaymentAmount** | Pointer to **float32** | The charge amount of the payment. Required in case of charge. | [optional] 
**PaymentScheduleDate** | Pointer to **string** | The date on which the payment is scheduled to be processed. Required for UPI and CARD payment modes. | [optional] 
**PaymentRemarks** | Pointer to **string** | Payment remarks. | [optional] 
**PaymentType** | **string** | Payment type. Can be AUTH or CHARGE. | 
**PaymentMethod** | Pointer to **map[string]interface{}** | Payment method. Can be one of [\&quot;upi\&quot;, \&quot;enach\&quot;, \&quot;pnach\&quot;, \&quot;card\&quot;] | [optional] 

## Methods

### NewCreateSubscriptionPaymentRequest

`func NewCreateSubscriptionPaymentRequest(subscriptionId string, paymentId string, paymentType string, ) *CreateSubscriptionPaymentRequest`

NewCreateSubscriptionPaymentRequest instantiates a new CreateSubscriptionPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentRequestWithDefaults

`func NewCreateSubscriptionPaymentRequestWithDefaults() *CreateSubscriptionPaymentRequest`

NewCreateSubscriptionPaymentRequestWithDefaults instantiates a new CreateSubscriptionPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateSubscriptionPaymentRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionPaymentRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionPaymentRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSubscriptionSessionId

`func (o *CreateSubscriptionPaymentRequest) GetSubscriptionSessionId() string`

GetSubscriptionSessionId returns the SubscriptionSessionId field if non-nil, zero value otherwise.

### GetSubscriptionSessionIdOk

`func (o *CreateSubscriptionPaymentRequest) GetSubscriptionSessionIdOk() (*string, bool)`

GetSubscriptionSessionIdOk returns a tuple with the SubscriptionSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionSessionId

`func (o *CreateSubscriptionPaymentRequest) SetSubscriptionSessionId(v string)`

SetSubscriptionSessionId sets SubscriptionSessionId field to given value.

### HasSubscriptionSessionId

`func (o *CreateSubscriptionPaymentRequest) HasSubscriptionSessionId() bool`

HasSubscriptionSessionId returns a boolean if a field has been set.

### GetPaymentId

`func (o *CreateSubscriptionPaymentRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateSubscriptionPaymentRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPaymentAmount

`func (o *CreateSubscriptionPaymentRequest) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *CreateSubscriptionPaymentRequest) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *CreateSubscriptionPaymentRequest) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentScheduleDate

`func (o *CreateSubscriptionPaymentRequest) GetPaymentScheduleDate() string`

GetPaymentScheduleDate returns the PaymentScheduleDate field if non-nil, zero value otherwise.

### GetPaymentScheduleDateOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentScheduleDateOk() (*string, bool)`

GetPaymentScheduleDateOk returns a tuple with the PaymentScheduleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleDate

`func (o *CreateSubscriptionPaymentRequest) SetPaymentScheduleDate(v string)`

SetPaymentScheduleDate sets PaymentScheduleDate field to given value.

### HasPaymentScheduleDate

`func (o *CreateSubscriptionPaymentRequest) HasPaymentScheduleDate() bool`

HasPaymentScheduleDate returns a boolean if a field has been set.

### GetPaymentRemarks

`func (o *CreateSubscriptionPaymentRequest) GetPaymentRemarks() string`

GetPaymentRemarks returns the PaymentRemarks field if non-nil, zero value otherwise.

### GetPaymentRemarksOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentRemarksOk() (*string, bool)`

GetPaymentRemarksOk returns a tuple with the PaymentRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRemarks

`func (o *CreateSubscriptionPaymentRequest) SetPaymentRemarks(v string)`

SetPaymentRemarks sets PaymentRemarks field to given value.

### HasPaymentRemarks

`func (o *CreateSubscriptionPaymentRequest) HasPaymentRemarks() bool`

HasPaymentRemarks returns a boolean if a field has been set.

### GetPaymentType

`func (o *CreateSubscriptionPaymentRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CreateSubscriptionPaymentRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentMethod

`func (o *CreateSubscriptionPaymentRequest) GetPaymentMethod() map[string]interface{}`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateSubscriptionPaymentRequest) GetPaymentMethodOk() (*map[string]interface{}, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateSubscriptionPaymentRequest) SetPaymentMethod(v map[string]interface{})`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreateSubscriptionPaymentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


