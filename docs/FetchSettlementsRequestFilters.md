# FetchSettlementsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfSettlementIds** | Pointer to **[]int64** | List of settlement IDs for which you want the settlement reconciliation details. | [optional] 
**SettlementUtrs** | Pointer to **[]string** | List of settlement UTRs for which you want the settlement reconciliation details. | [optional] 
**StartDate** | Pointer to **string** | Specify the start date from when you want the settlement reconciliation details. | [optional] 
**EndDate** | Pointer to **string** | Specify the end date till when you want the settlement reconciliation details. | [optional] 

## Methods

### NewFetchSettlementsRequestFilters

`func NewFetchSettlementsRequestFilters() *FetchSettlementsRequestFilters`

NewFetchSettlementsRequestFilters instantiates a new FetchSettlementsRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSettlementsRequestFiltersWithDefaults

`func NewFetchSettlementsRequestFiltersWithDefaults() *FetchSettlementsRequestFilters`

NewFetchSettlementsRequestFiltersWithDefaults instantiates a new FetchSettlementsRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementIds

`func (o *FetchSettlementsRequestFilters) GetCfSettlementIds() []int64`

GetCfSettlementIds returns the CfSettlementIds field if non-nil, zero value otherwise.

### GetCfSettlementIdsOk

`func (o *FetchSettlementsRequestFilters) GetCfSettlementIdsOk() (*[]int64, bool)`

GetCfSettlementIdsOk returns a tuple with the CfSettlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementIds

`func (o *FetchSettlementsRequestFilters) SetCfSettlementIds(v []int64)`

SetCfSettlementIds sets CfSettlementIds field to given value.

### HasCfSettlementIds

`func (o *FetchSettlementsRequestFilters) HasCfSettlementIds() bool`

HasCfSettlementIds returns a boolean if a field has been set.

### GetSettlementUtrs

`func (o *FetchSettlementsRequestFilters) GetSettlementUtrs() []string`

GetSettlementUtrs returns the SettlementUtrs field if non-nil, zero value otherwise.

### GetSettlementUtrsOk

`func (o *FetchSettlementsRequestFilters) GetSettlementUtrsOk() (*[]string, bool)`

GetSettlementUtrsOk returns a tuple with the SettlementUtrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtrs

`func (o *FetchSettlementsRequestFilters) SetSettlementUtrs(v []string)`

SetSettlementUtrs sets SettlementUtrs field to given value.

### HasSettlementUtrs

`func (o *FetchSettlementsRequestFilters) HasSettlementUtrs() bool`

HasSettlementUtrs returns a boolean if a field has been set.

### GetStartDate

`func (o *FetchSettlementsRequestFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FetchSettlementsRequestFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FetchSettlementsRequestFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FetchSettlementsRequestFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *FetchSettlementsRequestFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FetchSettlementsRequestFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FetchSettlementsRequestFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FetchSettlementsRequestFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


