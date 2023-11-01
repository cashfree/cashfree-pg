# ReconEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **NullableString** | Specifies from where the next set of settlement details should be fetched. | [optional] 
**Limit** | Pointer to **int32** | Number of settlements you want to fetch in the next iteration. | [optional] 
**Data** | Pointer to [**[]ReconEntityDataInner**](ReconEntityDataInner.md) |  | [optional] 

## Methods

### NewReconEntity

`func NewReconEntity() *ReconEntity`

NewReconEntity instantiates a new ReconEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityWithDefaults

`func NewReconEntityWithDefaults() *ReconEntity`

NewReconEntityWithDefaults instantiates a new ReconEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ReconEntity) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ReconEntity) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ReconEntity) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ReconEntity) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *ReconEntity) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *ReconEntity) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetLimit

`func (o *ReconEntity) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReconEntity) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReconEntity) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReconEntity) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetData

`func (o *ReconEntity) GetData() []ReconEntityDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReconEntity) GetDataOk() (*[]ReconEntityDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReconEntity) SetData(v []ReconEntityDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ReconEntity) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


