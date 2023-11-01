# PaymentSuccessWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WHdata**](WHdata.md) |  | [optional] 
**EventTime** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentSuccessWebhook

`func NewPaymentSuccessWebhook() *PaymentSuccessWebhook`

NewPaymentSuccessWebhook instantiates a new PaymentSuccessWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSuccessWebhookWithDefaults

`func NewPaymentSuccessWebhookWithDefaults() *PaymentSuccessWebhook`

NewPaymentSuccessWebhookWithDefaults instantiates a new PaymentSuccessWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentSuccessWebhook) GetData() WHdata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentSuccessWebhook) GetDataOk() (*WHdata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentSuccessWebhook) SetData(v WHdata)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentSuccessWebhook) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventTime

`func (o *PaymentSuccessWebhook) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *PaymentSuccessWebhook) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *PaymentSuccessWebhook) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *PaymentSuccessWebhook) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetType

`func (o *PaymentSuccessWebhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentSuccessWebhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentSuccessWebhook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentSuccessWebhook) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


