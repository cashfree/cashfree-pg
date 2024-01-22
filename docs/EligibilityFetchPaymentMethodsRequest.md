# EligibilityFetchPaymentMethodsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**PaymentMethodsQueries**](PaymentMethodsQueries.md) |  | 
**Filters** | Pointer to [**PaymentMethodsFilters**](PaymentMethodsFilters.md) |  | [optional] 

## Methods

### NewEligibilityFetchPaymentMethodsRequest

`func NewEligibilityFetchPaymentMethodsRequest(queries PaymentMethodsQueries, ) *EligibilityFetchPaymentMethodsRequest`

NewEligibilityFetchPaymentMethodsRequest instantiates a new EligibilityFetchPaymentMethodsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityFetchPaymentMethodsRequestWithDefaults

`func NewEligibilityFetchPaymentMethodsRequestWithDefaults() *EligibilityFetchPaymentMethodsRequest`

NewEligibilityFetchPaymentMethodsRequestWithDefaults instantiates a new EligibilityFetchPaymentMethodsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *EligibilityFetchPaymentMethodsRequest) GetQueries() PaymentMethodsQueries`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *EligibilityFetchPaymentMethodsRequest) GetQueriesOk() (*PaymentMethodsQueries, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *EligibilityFetchPaymentMethodsRequest) SetQueries(v PaymentMethodsQueries)`

SetQueries sets Queries field to given value.


### GetFilters

`func (o *EligibilityFetchPaymentMethodsRequest) GetFilters() PaymentMethodsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *EligibilityFetchPaymentMethodsRequest) GetFiltersOk() (*PaymentMethodsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *EligibilityFetchPaymentMethodsRequest) SetFilters(v PaymentMethodsFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *EligibilityFetchPaymentMethodsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


