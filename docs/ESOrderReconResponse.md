# ESOrderReconResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]ESOrderReconResponseDataInner**](ESOrderReconResponseDataInner.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewESOrderReconResponse

`func NewESOrderReconResponse() *ESOrderReconResponse`

NewESOrderReconResponse instantiates a new ESOrderReconResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESOrderReconResponseWithDefaults

`func NewESOrderReconResponseWithDefaults() *ESOrderReconResponse`

NewESOrderReconResponseWithDefaults instantiates a new ESOrderReconResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ESOrderReconResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ESOrderReconResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ESOrderReconResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ESOrderReconResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetData

`func (o *ESOrderReconResponse) GetData() []ESOrderReconResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ESOrderReconResponse) GetDataOk() (*[]ESOrderReconResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ESOrderReconResponse) SetData(v []ESOrderReconResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ESOrderReconResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *ESOrderReconResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ESOrderReconResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ESOrderReconResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ESOrderReconResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


