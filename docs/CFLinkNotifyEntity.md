# CFLinkNotifyEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendSms** | Pointer to **bool** | If \&quot;true\&quot;, Cashfree will send sms on customer_phone | [optional] 
**SendEmail** | Pointer to **bool** | If \&quot;true\&quot;, Cashfree will send email on customer_email | [optional] 

## Methods

### NewCFLinkNotifyEntity

`func NewCFLinkNotifyEntity() *CFLinkNotifyEntity`

NewCFLinkNotifyEntity instantiates a new CFLinkNotifyEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkNotifyEntityWithDefaults

`func NewCFLinkNotifyEntityWithDefaults() *CFLinkNotifyEntity`

NewCFLinkNotifyEntityWithDefaults instantiates a new CFLinkNotifyEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendSms

`func (o *CFLinkNotifyEntity) GetSendSms() bool`

GetSendSms returns the SendSms field if non-nil, zero value otherwise.

### GetSendSmsOk

`func (o *CFLinkNotifyEntity) GetSendSmsOk() (*bool, bool)`

GetSendSmsOk returns a tuple with the SendSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSms

`func (o *CFLinkNotifyEntity) SetSendSms(v bool)`

SetSendSms sets SendSms field to given value.

### HasSendSms

`func (o *CFLinkNotifyEntity) HasSendSms() bool`

HasSendSms returns a boolean if a field has been set.

### GetSendEmail

`func (o *CFLinkNotifyEntity) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *CFLinkNotifyEntity) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *CFLinkNotifyEntity) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *CFLinkNotifyEntity) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


