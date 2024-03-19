# ESOrderReconRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Specify the start data from which you want to get the recon data. | [optional] 
**EndDate** | Pointer to **string** | Specify the end data till which you want to get the recon data. | [optional] 
**OrderIds** | Pointer to **[]string** | Please provide list of order ids for which you want to get the recon data. | [optional] 

## Methods

### NewESOrderReconRequestFilters

`func NewESOrderReconRequestFilters() *ESOrderReconRequestFilters`

NewESOrderReconRequestFilters instantiates a new ESOrderReconRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESOrderReconRequestFiltersWithDefaults

`func NewESOrderReconRequestFiltersWithDefaults() *ESOrderReconRequestFilters`

NewESOrderReconRequestFiltersWithDefaults instantiates a new ESOrderReconRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ESOrderReconRequestFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ESOrderReconRequestFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ESOrderReconRequestFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ESOrderReconRequestFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ESOrderReconRequestFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ESOrderReconRequestFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ESOrderReconRequestFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ESOrderReconRequestFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOrderIds

`func (o *ESOrderReconRequestFilters) GetOrderIds() []string`

GetOrderIds returns the OrderIds field if non-nil, zero value otherwise.

### GetOrderIdsOk

`func (o *ESOrderReconRequestFilters) GetOrderIdsOk() (*[]string, bool)`

GetOrderIdsOk returns a tuple with the OrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIds

`func (o *ESOrderReconRequestFilters) SetOrderIds(v []string)`

SetOrderIds sets OrderIds field to given value.

### HasOrderIds

`func (o *ESOrderReconRequestFilters) HasOrderIds() bool`

HasOrderIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


