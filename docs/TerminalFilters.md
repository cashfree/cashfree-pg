# TerminalFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Start date and time for transaction filtering. Use ISO8601 format  (example: 2021-07-02T10:20:12+05:30 for IST, 2021-07-02T10:20:12Z for UTC). | [optional] 
**EndDate** | Pointer to **string** | End date and time for transaction filtering. Use ISO8601 format  (example: 2021-07-02T10:20:12+05:30 for IST, 2021-07-02T10:20:12Z for UTC). | [optional] 
**CfTerminalId** | Pointer to **int64** | Unique Cashfree terminal identifier. | [optional] 
**TerminalVpa** | Pointer to **string** | Virtual payment address (VPA) associated with the terminal. | [optional] 
**TerminalPhoneNo** | Pointer to **string** | Mobile phone number registered with the terminal. | [optional] 
**PaymentStatus** | Pointer to **string** | Payment transaction status. Possible values are SUCCESS, FAILED, or PENDING. | [optional] 
**PaymentGroup** | Pointer to **string** | Payment method group (for example, UPI, CARD, NET_BANKING). | [optional] 
**SortBy** | Pointer to **string** | Field to sort results by. Currently supports sorting by start_date. | [optional] 
**SortOrder** | Pointer to **string** | Sort order for the results. Use ASC for ascending or DESC for descending order. | [optional] 

## Methods

### NewTerminalFilters

`func NewTerminalFilters() *TerminalFilters`

NewTerminalFilters instantiates a new TerminalFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalFiltersWithDefaults

`func NewTerminalFiltersWithDefaults() *TerminalFilters`

NewTerminalFiltersWithDefaults instantiates a new TerminalFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *TerminalFilters) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TerminalFilters) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TerminalFilters) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TerminalFilters) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TerminalFilters) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TerminalFilters) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TerminalFilters) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TerminalFilters) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCfTerminalId

`func (o *TerminalFilters) GetCfTerminalId() int64`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *TerminalFilters) GetCfTerminalIdOk() (*int64, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *TerminalFilters) SetCfTerminalId(v int64)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *TerminalFilters) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetTerminalVpa

`func (o *TerminalFilters) GetTerminalVpa() string`

GetTerminalVpa returns the TerminalVpa field if non-nil, zero value otherwise.

### GetTerminalVpaOk

`func (o *TerminalFilters) GetTerminalVpaOk() (*string, bool)`

GetTerminalVpaOk returns a tuple with the TerminalVpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalVpa

`func (o *TerminalFilters) SetTerminalVpa(v string)`

SetTerminalVpa sets TerminalVpa field to given value.

### HasTerminalVpa

`func (o *TerminalFilters) HasTerminalVpa() bool`

HasTerminalVpa returns a boolean if a field has been set.

### GetTerminalPhoneNo

`func (o *TerminalFilters) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *TerminalFilters) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *TerminalFilters) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.

### HasTerminalPhoneNo

`func (o *TerminalFilters) HasTerminalPhoneNo() bool`

HasTerminalPhoneNo returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *TerminalFilters) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *TerminalFilters) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *TerminalFilters) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *TerminalFilters) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentGroup

`func (o *TerminalFilters) GetPaymentGroup() string`

GetPaymentGroup returns the PaymentGroup field if non-nil, zero value otherwise.

### GetPaymentGroupOk

`func (o *TerminalFilters) GetPaymentGroupOk() (*string, bool)`

GetPaymentGroupOk returns a tuple with the PaymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroup

`func (o *TerminalFilters) SetPaymentGroup(v string)`

SetPaymentGroup sets PaymentGroup field to given value.

### HasPaymentGroup

`func (o *TerminalFilters) HasPaymentGroup() bool`

HasPaymentGroup returns a boolean if a field has been set.

### GetSortBy

`func (o *TerminalFilters) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *TerminalFilters) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *TerminalFilters) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *TerminalFilters) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *TerminalFilters) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TerminalFilters) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TerminalFilters) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *TerminalFilters) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


