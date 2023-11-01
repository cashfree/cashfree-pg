# FetchReconRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | Specify the start date from when you want the settlement reconciliation details. | 
**EndDate** | **string** | Specify the end date till when you want the settlement reconciliation details. | 

## Methods

### NewFetchReconRequestFilters

`func NewFetchReconRequestFilters(startDate string, endDate string, ) *FetchReconRequestFilters`

NewFetchReconRequestFilters instantiates a new FetchReconRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchReconRequestFiltersWithDefaults

`func NewFetchReconRequestFiltersWithDefaults() *FetchReconRequestFilters`

NewFetchReconRequestFiltersWithDefaults instantiates a new FetchReconRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *FetchReconRequestFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FetchReconRequestFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FetchReconRequestFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *FetchReconRequestFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FetchReconRequestFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FetchReconRequestFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


