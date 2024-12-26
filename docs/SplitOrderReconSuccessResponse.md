# SplitOrderReconSuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settlement** | Pointer to [**SplitOrderReconSuccessResponseSettlement**](SplitOrderReconSuccessResponseSettlement.md) |  | [optional] 
**Refunds** | Pointer to **[]map[string]interface{}** | List of refunds associated with the order, if any. | [optional] 
**Vendors** | Pointer to [**[]SplitOrderReconSuccessResponseVendorsInner**](SplitOrderReconSuccessResponseVendorsInner.md) | List of vendor settlements associated with the split settlement. | [optional] 

## Methods

### NewSplitOrderReconSuccessResponse

`func NewSplitOrderReconSuccessResponse() *SplitOrderReconSuccessResponse`

NewSplitOrderReconSuccessResponse instantiates a new SplitOrderReconSuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitOrderReconSuccessResponseWithDefaults

`func NewSplitOrderReconSuccessResponseWithDefaults() *SplitOrderReconSuccessResponse`

NewSplitOrderReconSuccessResponseWithDefaults instantiates a new SplitOrderReconSuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlement

`func (o *SplitOrderReconSuccessResponse) GetSettlement() SplitOrderReconSuccessResponseSettlement`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *SplitOrderReconSuccessResponse) GetSettlementOk() (*SplitOrderReconSuccessResponseSettlement, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *SplitOrderReconSuccessResponse) SetSettlement(v SplitOrderReconSuccessResponseSettlement)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *SplitOrderReconSuccessResponse) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### GetRefunds

`func (o *SplitOrderReconSuccessResponse) GetRefunds() []map[string]interface{}`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *SplitOrderReconSuccessResponse) GetRefundsOk() (*[]map[string]interface{}, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *SplitOrderReconSuccessResponse) SetRefunds(v []map[string]interface{})`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *SplitOrderReconSuccessResponse) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetVendors

`func (o *SplitOrderReconSuccessResponse) GetVendors() []SplitOrderReconSuccessResponseVendorsInner`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *SplitOrderReconSuccessResponse) GetVendorsOk() (*[]SplitOrderReconSuccessResponseVendorsInner, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *SplitOrderReconSuccessResponse) SetVendors(v []SplitOrderReconSuccessResponseVendorsInner)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *SplitOrderReconSuccessResponse) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


