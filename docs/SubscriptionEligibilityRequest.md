# SubscriptionEligibilityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**SubscriptionEligibilityRequestQueries**](SubscriptionEligibilityRequestQueries.md) |  | 
**Filters** | Pointer to [**SubscriptionEligibilityRequestFilters**](SubscriptionEligibilityRequestFilters.md) |  | [optional] 

## Methods

### NewSubscriptionEligibilityRequest

`func NewSubscriptionEligibilityRequest(queries SubscriptionEligibilityRequestQueries, ) *SubscriptionEligibilityRequest`

NewSubscriptionEligibilityRequest instantiates a new SubscriptionEligibilityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEligibilityRequestWithDefaults

`func NewSubscriptionEligibilityRequestWithDefaults() *SubscriptionEligibilityRequest`

NewSubscriptionEligibilityRequestWithDefaults instantiates a new SubscriptionEligibilityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *SubscriptionEligibilityRequest) GetQueries() SubscriptionEligibilityRequestQueries`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SubscriptionEligibilityRequest) GetQueriesOk() (*SubscriptionEligibilityRequestQueries, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SubscriptionEligibilityRequest) SetQueries(v SubscriptionEligibilityRequestQueries)`

SetQueries sets Queries field to given value.


### GetFilters

`func (o *SubscriptionEligibilityRequest) GetFilters() SubscriptionEligibilityRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SubscriptionEligibilityRequest) GetFiltersOk() (*SubscriptionEligibilityRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SubscriptionEligibilityRequest) SetFilters(v SubscriptionEligibilityRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SubscriptionEligibilityRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


