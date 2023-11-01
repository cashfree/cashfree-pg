# FetchSettlementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**FetchSettlementsRequestPagination**](FetchSettlementsRequestPagination.md) |  | 
**Filters** | [**FetchSettlementsRequestFilters**](FetchSettlementsRequestFilters.md) |  | 

## Methods

### NewFetchSettlementsRequest

`func NewFetchSettlementsRequest(pagination FetchSettlementsRequestPagination, filters FetchSettlementsRequestFilters, ) *FetchSettlementsRequest`

NewFetchSettlementsRequest instantiates a new FetchSettlementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSettlementsRequestWithDefaults

`func NewFetchSettlementsRequestWithDefaults() *FetchSettlementsRequest`

NewFetchSettlementsRequestWithDefaults instantiates a new FetchSettlementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *FetchSettlementsRequest) GetPagination() FetchSettlementsRequestPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FetchSettlementsRequest) GetPaginationOk() (*FetchSettlementsRequestPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FetchSettlementsRequest) SetPagination(v FetchSettlementsRequestPagination)`

SetPagination sets Pagination field to given value.


### GetFilters

`func (o *FetchSettlementsRequest) GetFilters() FetchSettlementsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FetchSettlementsRequest) GetFiltersOk() (*FetchSettlementsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FetchSettlementsRequest) SetFilters(v FetchSettlementsRequestFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


