# FetchReconRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**FetchReconRequestPagination**](FetchReconRequestPagination.md) |  | 
**Filters** | [**FetchReconRequestFilters**](FetchReconRequestFilters.md) |  | 

## Methods

### NewFetchReconRequest

`func NewFetchReconRequest(pagination FetchReconRequestPagination, filters FetchReconRequestFilters, ) *FetchReconRequest`

NewFetchReconRequest instantiates a new FetchReconRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchReconRequestWithDefaults

`func NewFetchReconRequestWithDefaults() *FetchReconRequest`

NewFetchReconRequestWithDefaults instantiates a new FetchReconRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *FetchReconRequest) GetPagination() FetchReconRequestPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FetchReconRequest) GetPaginationOk() (*FetchReconRequestPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FetchReconRequest) SetPagination(v FetchReconRequestPagination)`

SetPagination sets Pagination field to given value.


### GetFilters

`func (o *FetchReconRequest) GetFilters() FetchReconRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FetchReconRequest) GetFiltersOk() (*FetchReconRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FetchReconRequest) SetFilters(v FetchReconRequestFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


