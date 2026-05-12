# GetAllTerminalTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**TerminalFilters**](TerminalFilters.md) |  | [optional] 
**Pagination** | Pointer to [**TerminalPagination**](TerminalPagination.md) |  | [optional] 

## Methods

### NewGetAllTerminalTransactionRequest

`func NewGetAllTerminalTransactionRequest() *GetAllTerminalTransactionRequest`

NewGetAllTerminalTransactionRequest instantiates a new GetAllTerminalTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllTerminalTransactionRequestWithDefaults

`func NewGetAllTerminalTransactionRequestWithDefaults() *GetAllTerminalTransactionRequest`

NewGetAllTerminalTransactionRequestWithDefaults instantiates a new GetAllTerminalTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *GetAllTerminalTransactionRequest) GetFilter() TerminalFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetAllTerminalTransactionRequest) GetFilterOk() (*TerminalFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetAllTerminalTransactionRequest) SetFilter(v TerminalFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetAllTerminalTransactionRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *GetAllTerminalTransactionRequest) GetPagination() TerminalPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAllTerminalTransactionRequest) GetPaginationOk() (*TerminalPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAllTerminalTransactionRequest) SetPagination(v TerminalPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAllTerminalTransactionRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


