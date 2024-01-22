# OrderAuthenticatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | **string** | OTP to be submitted | 
**Action** | **string** | The action for this workflow. Could be either SUBMIT_OTP or RESEND_OTP | 

## Methods

### NewOrderAuthenticatePaymentRequest

`func NewOrderAuthenticatePaymentRequest(otp string, action string, ) *OrderAuthenticatePaymentRequest`

NewOrderAuthenticatePaymentRequest instantiates a new OrderAuthenticatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAuthenticatePaymentRequestWithDefaults

`func NewOrderAuthenticatePaymentRequestWithDefaults() *OrderAuthenticatePaymentRequest`

NewOrderAuthenticatePaymentRequestWithDefaults instantiates a new OrderAuthenticatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *OrderAuthenticatePaymentRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *OrderAuthenticatePaymentRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *OrderAuthenticatePaymentRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.


### GetAction

`func (o *OrderAuthenticatePaymentRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderAuthenticatePaymentRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderAuthenticatePaymentRequest) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


