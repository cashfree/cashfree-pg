# LinkNotifyEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendSms** | Pointer to **bool** | If \&quot;true\&quot;, Cashfree will send sms on customer_phone | [optional] 
**SendEmail** | Pointer to **bool** | If \&quot;true\&quot;, Cashfree will send email on customer_email | [optional] 

## Methods

### NewLinkNotifyEntity

`func NewLinkNotifyEntity() *LinkNotifyEntity`

NewLinkNotifyEntity instantiates a new LinkNotifyEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkNotifyEntityWithDefaults

`func NewLinkNotifyEntityWithDefaults() *LinkNotifyEntity`

NewLinkNotifyEntityWithDefaults instantiates a new LinkNotifyEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendSms

`func (o *LinkNotifyEntity) GetSendSms() bool`

GetSendSms returns the SendSms field if non-nil, zero value otherwise.

### GetSendSmsOk

`func (o *LinkNotifyEntity) GetSendSmsOk() (*bool, bool)`

GetSendSmsOk returns a tuple with the SendSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSms

`func (o *LinkNotifyEntity) SetSendSms(v bool)`

SetSendSms sets SendSms field to given value.

### HasSendSms

`func (o *LinkNotifyEntity) HasSendSms() bool`

HasSendSms returns a boolean if a field has been set.

### GetSendEmail

`func (o *LinkNotifyEntity) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *LinkNotifyEntity) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *LinkNotifyEntity) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *LinkNotifyEntity) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


