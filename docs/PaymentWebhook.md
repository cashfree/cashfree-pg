# PaymentWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaymentWebhookDataEntity**](PaymentWebhookDataEntity.md) |  | [optional] 
**EventTime** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentWebhook

`func NewPaymentWebhook() *PaymentWebhook`

NewPaymentWebhook instantiates a new PaymentWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWebhookWithDefaults

`func NewPaymentWebhookWithDefaults() *PaymentWebhook`

NewPaymentWebhookWithDefaults instantiates a new PaymentWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentWebhook) GetData() PaymentWebhookDataEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentWebhook) GetDataOk() (*PaymentWebhookDataEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentWebhook) SetData(v PaymentWebhookDataEntity)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentWebhook) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventTime

`func (o *PaymentWebhook) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *PaymentWebhook) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *PaymentWebhook) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *PaymentWebhook) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetType

`func (o *PaymentWebhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentWebhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentWebhook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentWebhook) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


