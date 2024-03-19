# ESOrderReconRequestPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Set the minimum/maximum limit for number of filtered data. Min value - 10, Max value - 100. | [optional] 

## Methods

### NewESOrderReconRequestPagination

`func NewESOrderReconRequestPagination() *ESOrderReconRequestPagination`

NewESOrderReconRequestPagination instantiates a new ESOrderReconRequestPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESOrderReconRequestPaginationWithDefaults

`func NewESOrderReconRequestPaginationWithDefaults() *ESOrderReconRequestPagination`

NewESOrderReconRequestPaginationWithDefaults instantiates a new ESOrderReconRequestPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ESOrderReconRequestPagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ESOrderReconRequestPagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ESOrderReconRequestPagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ESOrderReconRequestPagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetLimit

`func (o *ESOrderReconRequestPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ESOrderReconRequestPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ESOrderReconRequestPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ESOrderReconRequestPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


