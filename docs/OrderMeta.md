# OrderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **string** | The URL to which user will be redirected to after the payment on bank OTP page. Maximum length: 250. We suggest to keep context of order_id in your return_url so that you can identify the order when customer lands on your page. Example of return_url format could be https://www.cashfree.com/devstudio/thankyou | [optional] 
**NotifyUrl** | Pointer to **string** | Notification URL for server-server communication. Useful when user&#39;s connection drops while re-directing. NotifyUrl should be an https URL. Maximum length: 250. | [optional] 
**PaymentMethods** | Pointer to **interface{}** | Allowed payment modes for this order. Pass comma-separated values among following options - \&quot;cc\&quot;, \&quot;dc\&quot;, \&quot;ccc\&quot;, \&quot;ppc\&quot;,\&quot;nb\&quot;,\&quot;upi\&quot;,\&quot;paypal\&quot;,\&quot;app\&quot;,\&quot;paylater\&quot;,\&quot;cardlessemi\&quot;,\&quot;dcemi\&quot;,\&quot;ccemi\&quot;,\&quot;banktransfer\&quot;. Leave it blank to show all available payment methods | [optional] 
**PaymentMethodsFilters** | Pointer to [**OrderMetaPaymentMethodsFilters**](OrderMetaPaymentMethodsFilters.md) |  | [optional] 

## Methods

### NewOrderMeta

`func NewOrderMeta() *OrderMeta`

NewOrderMeta instantiates a new OrderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderMetaWithDefaults

`func NewOrderMetaWithDefaults() *OrderMeta`

NewOrderMetaWithDefaults instantiates a new OrderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *OrderMeta) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OrderMeta) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OrderMeta) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OrderMeta) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetNotifyUrl

`func (o *OrderMeta) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *OrderMeta) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *OrderMeta) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *OrderMeta) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *OrderMeta) GetPaymentMethods() interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *OrderMeta) GetPaymentMethodsOk() (*interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *OrderMeta) SetPaymentMethods(v interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *OrderMeta) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### SetPaymentMethodsNil

`func (o *OrderMeta) SetPaymentMethodsNil(b bool)`

 SetPaymentMethodsNil sets the value for PaymentMethods to be an explicit nil

### UnsetPaymentMethods
`func (o *OrderMeta) UnsetPaymentMethods()`

UnsetPaymentMethods ensures that no value is present for PaymentMethods, not even an explicit nil
### GetPaymentMethodsFilters

`func (o *OrderMeta) GetPaymentMethodsFilters() OrderMetaPaymentMethodsFilters`

GetPaymentMethodsFilters returns the PaymentMethodsFilters field if non-nil, zero value otherwise.

### GetPaymentMethodsFiltersOk

`func (o *OrderMeta) GetPaymentMethodsFiltersOk() (*OrderMetaPaymentMethodsFilters, bool)`

GetPaymentMethodsFiltersOk returns a tuple with the PaymentMethodsFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodsFilters

`func (o *OrderMeta) SetPaymentMethodsFilters(v OrderMetaPaymentMethodsFilters)`

SetPaymentMethodsFilters sets PaymentMethodsFilters field to given value.

### HasPaymentMethodsFilters

`func (o *OrderMeta) HasPaymentMethodsFilters() bool`

HasPaymentMethodsFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


