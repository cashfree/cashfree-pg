# FetchSettlementsRequestPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The number of settlements you want to fetch. Maximum limit is 1000, default value is 10. | 
**Cursor** | Pointer to **string** | Specifies from where the next set of settlement details should be fetched. | [optional] 

## Methods

### NewFetchSettlementsRequestPagination

`func NewFetchSettlementsRequestPagination(limit int32, ) *FetchSettlementsRequestPagination`

NewFetchSettlementsRequestPagination instantiates a new FetchSettlementsRequestPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSettlementsRequestPaginationWithDefaults

`func NewFetchSettlementsRequestPaginationWithDefaults() *FetchSettlementsRequestPagination`

NewFetchSettlementsRequestPaginationWithDefaults instantiates a new FetchSettlementsRequestPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *FetchSettlementsRequestPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FetchSettlementsRequestPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FetchSettlementsRequestPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetCursor

`func (o *FetchSettlementsRequestPagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FetchSettlementsRequestPagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FetchSettlementsRequestPagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FetchSettlementsRequestPagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


