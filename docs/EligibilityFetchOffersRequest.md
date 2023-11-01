# EligibilityFetchOffersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**OfferQueries**](OfferQueries.md) |  | 
**Filters** | Pointer to [**OfferFilters**](OfferFilters.md) |  | [optional] 

## Methods

### NewEligibilityFetchOffersRequest

`func NewEligibilityFetchOffersRequest(queries OfferQueries, ) *EligibilityFetchOffersRequest`

NewEligibilityFetchOffersRequest instantiates a new EligibilityFetchOffersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityFetchOffersRequestWithDefaults

`func NewEligibilityFetchOffersRequestWithDefaults() *EligibilityFetchOffersRequest`

NewEligibilityFetchOffersRequestWithDefaults instantiates a new EligibilityFetchOffersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *EligibilityFetchOffersRequest) GetQueries() OfferQueries`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *EligibilityFetchOffersRequest) GetQueriesOk() (*OfferQueries, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *EligibilityFetchOffersRequest) SetQueries(v OfferQueries)`

SetQueries sets Queries field to given value.


### GetFilters

`func (o *EligibilityFetchOffersRequest) GetFilters() OfferFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EligibilityFetchOffersRequest) GetFiltersOk() (*OfferFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EligibilityFetchOffersRequest) SetFilters(v OfferFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EligibilityFetchOffersRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


