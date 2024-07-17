# SubscriptionPaymentEntityFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | Failure reason of the payment if the payment_status is failed. | [optional] 

## Methods

### NewSubscriptionPaymentEntityFailureDetails

`func NewSubscriptionPaymentEntityFailureDetails() *SubscriptionPaymentEntityFailureDetails`

NewSubscriptionPaymentEntityFailureDetails instantiates a new SubscriptionPaymentEntityFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPaymentEntityFailureDetailsWithDefaults

`func NewSubscriptionPaymentEntityFailureDetailsWithDefaults() *SubscriptionPaymentEntityFailureDetails`

NewSubscriptionPaymentEntityFailureDetailsWithDefaults instantiates a new SubscriptionPaymentEntityFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *SubscriptionPaymentEntityFailureDetails) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *SubscriptionPaymentEntityFailureDetails) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *SubscriptionPaymentEntityFailureDetails) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *SubscriptionPaymentEntityFailureDetails) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


