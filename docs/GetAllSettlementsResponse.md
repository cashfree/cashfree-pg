# GetAllSettlementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **NullableString** | Cursor for the next set of settlements. Pass this value in the next API call to fetch the next page. Null if there are no more settlements. | [optional] 
**Data** | Pointer to [**[]GetAllSettlementsResponseDataInner**](GetAllSettlementsResponseDataInner.md) |  | [optional] 
**Limit** | Pointer to **int32** | Number of settlement records returned in this response. | [optional] 

## Methods

### NewGetAllSettlementsResponse

`func NewGetAllSettlementsResponse() *GetAllSettlementsResponse`

NewGetAllSettlementsResponse instantiates a new GetAllSettlementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSettlementsResponseWithDefaults

`func NewGetAllSettlementsResponseWithDefaults() *GetAllSettlementsResponse`

NewGetAllSettlementsResponseWithDefaults instantiates a new GetAllSettlementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetAllSettlementsResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetAllSettlementsResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetAllSettlementsResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetAllSettlementsResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *GetAllSettlementsResponse) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *GetAllSettlementsResponse) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetData

`func (o *GetAllSettlementsResponse) GetData() []GetAllSettlementsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllSettlementsResponse) GetDataOk() (*[]GetAllSettlementsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllSettlementsResponse) SetData(v []GetAllSettlementsResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllSettlementsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *GetAllSettlementsResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetAllSettlementsResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetAllSettlementsResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetAllSettlementsResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


