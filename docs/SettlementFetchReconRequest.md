# SettlementFetchReconRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**FetchSettlementsRequestPagination**](FetchSettlementsRequestPagination.md) |  | 
**Filters** | [**FetchSettlementsRequestFilters**](FetchSettlementsRequestFilters.md) |  | 

## Methods

### NewSettlementFetchReconRequest

`func NewSettlementFetchReconRequest(pagination FetchSettlementsRequestPagination, filters FetchSettlementsRequestFilters, ) *SettlementFetchReconRequest`

NewSettlementFetchReconRequest instantiates a new SettlementFetchReconRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementFetchReconRequestWithDefaults

`func NewSettlementFetchReconRequestWithDefaults() *SettlementFetchReconRequest`

NewSettlementFetchReconRequestWithDefaults instantiates a new SettlementFetchReconRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SettlementFetchReconRequest) GetPagination() FetchSettlementsRequestPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SettlementFetchReconRequest) GetPaginationOk() (*FetchSettlementsRequestPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SettlementFetchReconRequest) SetPagination(v FetchSettlementsRequestPagination)`

SetPagination sets Pagination field to given value.


### GetFilters

`func (o *SettlementFetchReconRequest) GetFilters() FetchSettlementsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SettlementFetchReconRequest) GetFiltersOk() (*FetchSettlementsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SettlementFetchReconRequest) SetFilters(v FetchSettlementsRequestFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


