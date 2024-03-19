# ESOrderReconRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ESOrderReconRequestFilters**](ESOrderReconRequestFilters.md) |  | 
**Pagination** | [**ESOrderReconRequestPagination**](ESOrderReconRequestPagination.md) |  | 

## Methods

### NewESOrderReconRequest

`func NewESOrderReconRequest(filters ESOrderReconRequestFilters, pagination ESOrderReconRequestPagination, ) *ESOrderReconRequest`

NewESOrderReconRequest instantiates a new ESOrderReconRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESOrderReconRequestWithDefaults

`func NewESOrderReconRequestWithDefaults() *ESOrderReconRequest`

NewESOrderReconRequestWithDefaults instantiates a new ESOrderReconRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ESOrderReconRequest) GetFilters() ESOrderReconRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ESOrderReconRequest) GetFiltersOk() (*ESOrderReconRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ESOrderReconRequest) SetFilters(v ESOrderReconRequestFilters)`

SetFilters sets Filters field to given value.


### GetPagination

`func (o *ESOrderReconRequest) GetPagination() ESOrderReconRequestPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ESOrderReconRequest) GetPaginationOk() (*ESOrderReconRequestPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ESOrderReconRequest) SetPagination(v ESOrderReconRequestPagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


