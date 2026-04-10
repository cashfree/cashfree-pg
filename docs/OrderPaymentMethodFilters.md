# OrderPaymentMethodFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardEmiTenure** | Pointer to **int32** | Allowed card EMI tenure for the order. | [optional] 
**CardEmiBins** | Pointer to **[]int32** | Allowed card EMI bins for the order. | [optional] 
**CardEmiSchemes** | Pointer to **[]string** | Allowed card EMI schemes for the order. | [optional] 
**CardEmiSuffix** | Pointer to **[]int32** | Allowed card EMI suffixes for the order. | [optional] 
**CardEmiIssuingBank** | Pointer to **[]string** | Allowed card EMI issuing bank for the order. | [optional] 
**CardBins** | Pointer to **[]int32** | Allowed card bins for the order. | [optional] 
**CardSchemes** | Pointer to **[]string** | Allowed card schemes for the order. | [optional] 
**CardSuffix** | Pointer to **[]int32** | Allowed card suffixes for the order. | [optional] 
**CardIssuingBank** | Pointer to **[]string** | Allowed card issuing bank for the order. | [optional] 

## Methods

### NewOrderPaymentMethodFilters

`func NewOrderPaymentMethodFilters() *OrderPaymentMethodFilters`

NewOrderPaymentMethodFilters instantiates a new OrderPaymentMethodFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPaymentMethodFiltersWithDefaults

`func NewOrderPaymentMethodFiltersWithDefaults() *OrderPaymentMethodFilters`

NewOrderPaymentMethodFiltersWithDefaults instantiates a new OrderPaymentMethodFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardEmiTenure

`func (o *OrderPaymentMethodFilters) GetCardEmiTenure() int32`

GetCardEmiTenure returns the CardEmiTenure field if non-nil, zero value otherwise.

### GetCardEmiTenureOk

`func (o *OrderPaymentMethodFilters) GetCardEmiTenureOk() (*int32, bool)`

GetCardEmiTenureOk returns a tuple with the CardEmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardEmiTenure

`func (o *OrderPaymentMethodFilters) SetCardEmiTenure(v int32)`

SetCardEmiTenure sets CardEmiTenure field to given value.

### HasCardEmiTenure

`func (o *OrderPaymentMethodFilters) HasCardEmiTenure() bool`

HasCardEmiTenure returns a boolean if a field has been set.

### GetCardEmiBins

`func (o *OrderPaymentMethodFilters) GetCardEmiBins() []int32`

GetCardEmiBins returns the CardEmiBins field if non-nil, zero value otherwise.

### GetCardEmiBinsOk

`func (o *OrderPaymentMethodFilters) GetCardEmiBinsOk() (*[]int32, bool)`

GetCardEmiBinsOk returns a tuple with the CardEmiBins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardEmiBins

`func (o *OrderPaymentMethodFilters) SetCardEmiBins(v []int32)`

SetCardEmiBins sets CardEmiBins field to given value.

### HasCardEmiBins

`func (o *OrderPaymentMethodFilters) HasCardEmiBins() bool`

HasCardEmiBins returns a boolean if a field has been set.

### GetCardEmiSchemes

`func (o *OrderPaymentMethodFilters) GetCardEmiSchemes() []string`

GetCardEmiSchemes returns the CardEmiSchemes field if non-nil, zero value otherwise.

### GetCardEmiSchemesOk

`func (o *OrderPaymentMethodFilters) GetCardEmiSchemesOk() (*[]string, bool)`

GetCardEmiSchemesOk returns a tuple with the CardEmiSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardEmiSchemes

`func (o *OrderPaymentMethodFilters) SetCardEmiSchemes(v []string)`

SetCardEmiSchemes sets CardEmiSchemes field to given value.

### HasCardEmiSchemes

`func (o *OrderPaymentMethodFilters) HasCardEmiSchemes() bool`

HasCardEmiSchemes returns a boolean if a field has been set.

### GetCardEmiSuffix

`func (o *OrderPaymentMethodFilters) GetCardEmiSuffix() []int32`

GetCardEmiSuffix returns the CardEmiSuffix field if non-nil, zero value otherwise.

### GetCardEmiSuffixOk

`func (o *OrderPaymentMethodFilters) GetCardEmiSuffixOk() (*[]int32, bool)`

GetCardEmiSuffixOk returns a tuple with the CardEmiSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardEmiSuffix

`func (o *OrderPaymentMethodFilters) SetCardEmiSuffix(v []int32)`

SetCardEmiSuffix sets CardEmiSuffix field to given value.

### HasCardEmiSuffix

`func (o *OrderPaymentMethodFilters) HasCardEmiSuffix() bool`

HasCardEmiSuffix returns a boolean if a field has been set.

### GetCardEmiIssuingBank

`func (o *OrderPaymentMethodFilters) GetCardEmiIssuingBank() []string`

GetCardEmiIssuingBank returns the CardEmiIssuingBank field if non-nil, zero value otherwise.

### GetCardEmiIssuingBankOk

`func (o *OrderPaymentMethodFilters) GetCardEmiIssuingBankOk() (*[]string, bool)`

GetCardEmiIssuingBankOk returns a tuple with the CardEmiIssuingBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardEmiIssuingBank

`func (o *OrderPaymentMethodFilters) SetCardEmiIssuingBank(v []string)`

SetCardEmiIssuingBank sets CardEmiIssuingBank field to given value.

### HasCardEmiIssuingBank

`func (o *OrderPaymentMethodFilters) HasCardEmiIssuingBank() bool`

HasCardEmiIssuingBank returns a boolean if a field has been set.

### GetCardBins

`func (o *OrderPaymentMethodFilters) GetCardBins() []int32`

GetCardBins returns the CardBins field if non-nil, zero value otherwise.

### GetCardBinsOk

`func (o *OrderPaymentMethodFilters) GetCardBinsOk() (*[]int32, bool)`

GetCardBinsOk returns a tuple with the CardBins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBins

`func (o *OrderPaymentMethodFilters) SetCardBins(v []int32)`

SetCardBins sets CardBins field to given value.

### HasCardBins

`func (o *OrderPaymentMethodFilters) HasCardBins() bool`

HasCardBins returns a boolean if a field has been set.

### GetCardSchemes

`func (o *OrderPaymentMethodFilters) GetCardSchemes() []string`

GetCardSchemes returns the CardSchemes field if non-nil, zero value otherwise.

### GetCardSchemesOk

`func (o *OrderPaymentMethodFilters) GetCardSchemesOk() (*[]string, bool)`

GetCardSchemesOk returns a tuple with the CardSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSchemes

`func (o *OrderPaymentMethodFilters) SetCardSchemes(v []string)`

SetCardSchemes sets CardSchemes field to given value.

### HasCardSchemes

`func (o *OrderPaymentMethodFilters) HasCardSchemes() bool`

HasCardSchemes returns a boolean if a field has been set.

### GetCardSuffix

`func (o *OrderPaymentMethodFilters) GetCardSuffix() []int32`

GetCardSuffix returns the CardSuffix field if non-nil, zero value otherwise.

### GetCardSuffixOk

`func (o *OrderPaymentMethodFilters) GetCardSuffixOk() (*[]int32, bool)`

GetCardSuffixOk returns a tuple with the CardSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSuffix

`func (o *OrderPaymentMethodFilters) SetCardSuffix(v []int32)`

SetCardSuffix sets CardSuffix field to given value.

### HasCardSuffix

`func (o *OrderPaymentMethodFilters) HasCardSuffix() bool`

HasCardSuffix returns a boolean if a field has been set.

### GetCardIssuingBank

`func (o *OrderPaymentMethodFilters) GetCardIssuingBank() []string`

GetCardIssuingBank returns the CardIssuingBank field if non-nil, zero value otherwise.

### GetCardIssuingBankOk

`func (o *OrderPaymentMethodFilters) GetCardIssuingBankOk() (*[]string, bool)`

GetCardIssuingBankOk returns a tuple with the CardIssuingBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuingBank

`func (o *OrderPaymentMethodFilters) SetCardIssuingBank(v []string)`

SetCardIssuingBank sets CardIssuingBank field to given value.

### HasCardIssuingBank

`func (o *OrderPaymentMethodFilters) HasCardIssuingBank() bool`

HasCardIssuingBank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


