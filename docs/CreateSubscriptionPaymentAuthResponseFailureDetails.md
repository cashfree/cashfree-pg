# CreateSubscriptionPaymentAuthResponseFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | Failure reason of the payment if the payment_status is failed. | [optional] 

## Methods

### NewCreateSubscriptionPaymentAuthResponseFailureDetails

`func NewCreateSubscriptionPaymentAuthResponseFailureDetails() *CreateSubscriptionPaymentAuthResponseFailureDetails`

NewCreateSubscriptionPaymentAuthResponseFailureDetails instantiates a new CreateSubscriptionPaymentAuthResponseFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentAuthResponseFailureDetailsWithDefaults

`func NewCreateSubscriptionPaymentAuthResponseFailureDetailsWithDefaults() *CreateSubscriptionPaymentAuthResponseFailureDetails`

NewCreateSubscriptionPaymentAuthResponseFailureDetailsWithDefaults instantiates a new CreateSubscriptionPaymentAuthResponseFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *CreateSubscriptionPaymentAuthResponseFailureDetails) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *CreateSubscriptionPaymentAuthResponseFailureDetails) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *CreateSubscriptionPaymentAuthResponseFailureDetails) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *CreateSubscriptionPaymentAuthResponseFailureDetails) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


