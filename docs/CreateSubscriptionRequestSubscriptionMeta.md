# CreateSubscriptionRequestSubscriptionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **string** | The url to redirect after checkout. | [optional] 
**NotificationChannel** | Pointer to **[]string** | Notification channel for the subscription. SMS, EMAIL are possible values. | [optional] 

## Methods

### NewCreateSubscriptionRequestSubscriptionMeta

`func NewCreateSubscriptionRequestSubscriptionMeta() *CreateSubscriptionRequestSubscriptionMeta`

NewCreateSubscriptionRequestSubscriptionMeta instantiates a new CreateSubscriptionRequestSubscriptionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestSubscriptionMetaWithDefaults

`func NewCreateSubscriptionRequestSubscriptionMetaWithDefaults() *CreateSubscriptionRequestSubscriptionMeta`

NewCreateSubscriptionRequestSubscriptionMetaWithDefaults instantiates a new CreateSubscriptionRequestSubscriptionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *CreateSubscriptionRequestSubscriptionMeta) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateSubscriptionRequestSubscriptionMeta) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateSubscriptionRequestSubscriptionMeta) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreateSubscriptionRequestSubscriptionMeta) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetNotificationChannel

`func (o *CreateSubscriptionRequestSubscriptionMeta) GetNotificationChannel() []string`

GetNotificationChannel returns the NotificationChannel field if non-nil, zero value otherwise.

### GetNotificationChannelOk

`func (o *CreateSubscriptionRequestSubscriptionMeta) GetNotificationChannelOk() (*[]string, bool)`

GetNotificationChannelOk returns a tuple with the NotificationChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannel

`func (o *CreateSubscriptionRequestSubscriptionMeta) SetNotificationChannel(v []string)`

SetNotificationChannel sets NotificationChannel field to given value.

### HasNotificationChannel

`func (o *CreateSubscriptionRequestSubscriptionMeta) HasNotificationChannel() bool`

HasNotificationChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


