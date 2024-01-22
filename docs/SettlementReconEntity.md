# SettlementReconEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | Specifies from where the next set of settlement details should be fetched. | [optional] 
**Limit** | Pointer to **int32** | Number of settlements you want to fetch in the next iteration. | [optional] 
**Data** | Pointer to [**[]SettlementReconEntityDataInner**](SettlementReconEntityDataInner.md) |  | [optional] 

## Methods

### NewSettlementReconEntity

`func NewSettlementReconEntity() *SettlementReconEntity`

NewSettlementReconEntity instantiates a new SettlementReconEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityWithDefaults

`func NewSettlementReconEntityWithDefaults() *SettlementReconEntity`

NewSettlementReconEntityWithDefaults instantiates a new SettlementReconEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *SettlementReconEntity) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SettlementReconEntity) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SettlementReconEntity) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SettlementReconEntity) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetLimit

`func (o *SettlementReconEntity) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SettlementReconEntity) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SettlementReconEntity) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SettlementReconEntity) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetData

`func (o *SettlementReconEntity) GetData() []SettlementReconEntityDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SettlementReconEntity) GetDataOk() (*[]SettlementReconEntityDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SettlementReconEntity) SetData(v []SettlementReconEntityDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *SettlementReconEntity) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


