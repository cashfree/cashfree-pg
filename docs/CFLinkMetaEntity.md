# CFLinkMetaEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyUrl** | Pointer to **string** | Notification URL for server-server communication. It should be an https URL. | [optional] 
**UpiIntent** | Pointer to **bool** | If \&quot;true\&quot;, link will directly open UPI Intent flow on mobile, and normal link flow elsewhere | [optional] 
**ReturnUrl** | Pointer to **string** | The URL to which user will be redirected to after the payment is done on the link. Maximum length: 250. | [optional] 
**PaymentMethods** | Pointer to **string** | Allowed payment modes for this link. Pass comma-separated values among following options - \&quot;cc\&quot;, \&quot;dc\&quot;, \&quot;ccc\&quot;, \&quot;ppc\&quot;, \&quot;nb\&quot;, \&quot;upi\&quot;, \&quot;paypal\&quot;, \&quot;app\&quot;. Leave it blank to show all available payment methods | [optional] 

## Methods

### NewCFLinkMetaEntity

`func NewCFLinkMetaEntity() *CFLinkMetaEntity`

NewCFLinkMetaEntity instantiates a new CFLinkMetaEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkMetaEntityWithDefaults

`func NewCFLinkMetaEntityWithDefaults() *CFLinkMetaEntity`

NewCFLinkMetaEntityWithDefaults instantiates a new CFLinkMetaEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyUrl

`func (o *CFLinkMetaEntity) GetNotifyUrl() string`

GetNotifyUrl returns the NotifyUrl field if non-nil, zero value otherwise.

### GetNotifyUrlOk

`func (o *CFLinkMetaEntity) GetNotifyUrlOk() (*string, bool)`

GetNotifyUrlOk returns a tuple with the NotifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUrl

`func (o *CFLinkMetaEntity) SetNotifyUrl(v string)`

SetNotifyUrl sets NotifyUrl field to given value.

### HasNotifyUrl

`func (o *CFLinkMetaEntity) HasNotifyUrl() bool`

HasNotifyUrl returns a boolean if a field has been set.

### GetUpiIntent

`func (o *CFLinkMetaEntity) GetUpiIntent() bool`

GetUpiIntent returns the UpiIntent field if non-nil, zero value otherwise.

### GetUpiIntentOk

`func (o *CFLinkMetaEntity) GetUpiIntentOk() (*bool, bool)`

GetUpiIntentOk returns a tuple with the UpiIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiIntent

`func (o *CFLinkMetaEntity) SetUpiIntent(v bool)`

SetUpiIntent sets UpiIntent field to given value.

### HasUpiIntent

`func (o *CFLinkMetaEntity) HasUpiIntent() bool`

HasUpiIntent returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CFLinkMetaEntity) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CFLinkMetaEntity) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CFLinkMetaEntity) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CFLinkMetaEntity) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *CFLinkMetaEntity) GetPaymentMethods() string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *CFLinkMetaEntity) GetPaymentMethodsOk() (*string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *CFLinkMetaEntity) SetPaymentMethods(v string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *CFLinkMetaEntity) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


