# SettlementFetchReconRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfSettlementIds** | Pointer to **[]int32** | List of settlement IDs for which you want the settlement reconciliation details. | [optional] 
**SettlementUtrs** | Pointer to **[]string** | List of settlement UTRs for which you want the settlement reconciliation details. | [optional] 
**StartDate** | Pointer to **string** | Specify the start date from when you want the settlement reconciliation details. | [optional] 
**EndDate** | Pointer to **string** | Specify the end date till when you want the settlement reconciliation details. | [optional] 

## Methods

### NewSettlementFetchReconRequestFilters

`func NewSettlementFetchReconRequestFilters() *SettlementFetchReconRequestFilters`

NewSettlementFetchReconRequestFilters instantiates a new SettlementFetchReconRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementFetchReconRequestFiltersWithDefaults

`func NewSettlementFetchReconRequestFiltersWithDefaults() *SettlementFetchReconRequestFilters`

NewSettlementFetchReconRequestFiltersWithDefaults instantiates a new SettlementFetchReconRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementIds

`func (o *SettlementFetchReconRequestFilters) GetCfSettlementIds() []int32`

GetCfSettlementIds returns the CfSettlementIds field if non-nil, zero value otherwise.

### GetCfSettlementIdsOk

`func (o *SettlementFetchReconRequestFilters) GetCfSettlementIdsOk() (*[]int32, bool)`

GetCfSettlementIdsOk returns a tuple with the CfSettlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementIds

`func (o *SettlementFetchReconRequestFilters) SetCfSettlementIds(v []int32)`

SetCfSettlementIds sets CfSettlementIds field to given value.

### HasCfSettlementIds

`func (o *SettlementFetchReconRequestFilters) HasCfSettlementIds() bool`

HasCfSettlementIds returns a boolean if a field has been set.

### GetSettlementUtrs

`func (o *SettlementFetchReconRequestFilters) GetSettlementUtrs() []string`

GetSettlementUtrs returns the SettlementUtrs field if non-nil, zero value otherwise.

### GetSettlementUtrsOk

`func (o *SettlementFetchReconRequestFilters) GetSettlementUtrsOk() (*[]string, bool)`

GetSettlementUtrsOk returns a tuple with the SettlementUtrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtrs

`func (o *SettlementFetchReconRequestFilters) SetSettlementUtrs(v []string)`

SetSettlementUtrs sets SettlementUtrs field to given value.

### HasSettlementUtrs

`func (o *SettlementFetchReconRequestFilters) HasSettlementUtrs() bool`

HasSettlementUtrs returns a boolean if a field has been set.

### GetStartDate

`func (o *SettlementFetchReconRequestFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SettlementFetchReconRequestFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SettlementFetchReconRequestFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SettlementFetchReconRequestFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SettlementFetchReconRequestFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SettlementFetchReconRequestFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SettlementFetchReconRequestFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SettlementFetchReconRequestFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


