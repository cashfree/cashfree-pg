# SettlementWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**SettlementWebhookDataEntity**](SettlementWebhookDataEntity.md) |  | [optional] 
**EventTime** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSettlementWebhook

`func NewSettlementWebhook() *SettlementWebhook`

NewSettlementWebhook instantiates a new SettlementWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWebhookWithDefaults

`func NewSettlementWebhookWithDefaults() *SettlementWebhook`

NewSettlementWebhookWithDefaults instantiates a new SettlementWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SettlementWebhook) GetData() SettlementWebhookDataEntity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SettlementWebhook) GetDataOk() (*SettlementWebhookDataEntity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SettlementWebhook) SetData(v SettlementWebhookDataEntity)`

SetData sets Data field to given value.

### HasData

`func (o *SettlementWebhook) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventTime

`func (o *SettlementWebhook) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SettlementWebhook) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SettlementWebhook) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *SettlementWebhook) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetType

`func (o *SettlementWebhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SettlementWebhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SettlementWebhook) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SettlementWebhook) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


