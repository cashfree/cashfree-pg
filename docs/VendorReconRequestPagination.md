# VendorReconRequestPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Set the minimum or maximum limit for the number of filtered data. Minimum value: 10, Maximum value: 100. | 
**Cursor** | Pointer to **string** | Specifies from where the next set of records should be fetched. | [optional] 

## Methods

### NewVendorReconRequestPagination

`func NewVendorReconRequestPagination(limit int32, ) *VendorReconRequestPagination`

NewVendorReconRequestPagination instantiates a new VendorReconRequestPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorReconRequestPaginationWithDefaults

`func NewVendorReconRequestPaginationWithDefaults() *VendorReconRequestPagination`

NewVendorReconRequestPaginationWithDefaults instantiates a new VendorReconRequestPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *VendorReconRequestPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VendorReconRequestPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VendorReconRequestPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetCursor

`func (o *VendorReconRequestPagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *VendorReconRequestPagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *VendorReconRequestPagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *VendorReconRequestPagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


