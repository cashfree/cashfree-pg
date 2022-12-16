# CFOrderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | **string** | The URL to which user will be redirected to after the payment on bank OTP page. Maximum length: 250. | 
**NotifyUrl** | **string** | Notification URL for server-server communication. Useful when user&#39;s connection drops while re-directing. NotifyUrl should be an https URL. Maximum length: 250. | 
**PaymentMethods** | **string** | Allowed payment modes for this order. Pass comma-separated values among following options - \&quot;cc\&quot;, \&quot;dc\&quot;, \&quot;ccc\&quot;, \&quot;ppc\&quot;, \&quot;nb\&quot;, \&quot;upi\&quot;, \&quot;paypal\&quot;, \&quot;app\&quot;. Leave it blank to show all available payment methods | 

## Methods

### NewCFOrderMeta

`func NewCFOrderMeta(returnUrl string, notifyUrl string, paymentMethods string, ) *CFOrderMeta`

NewCFOrderMeta instantiates a new CFOrderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFOrderMetaWithDefaults

`func NewCFOrderMetaWithDefaults() *CFOrderMeta`

NewCFOrderMetaWithDefaults instantiates a new CFOrderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *CFOrderMeta) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CFOrderMeta) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CFOrderMeta) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetNotifyUrl

`func (o *CFOrderMeta) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *CFOrderMeta) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *CFOrderMeta) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.


### GetPaymentMethods

`func (o *CFOrderMeta) GetPaymentMethods() string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *CFOrderMeta) GetPaymentMethodsOk() (*string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *CFOrderMeta) SetPaymentMethods(v string)`

SetPaymentMethods sets PaymentMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


