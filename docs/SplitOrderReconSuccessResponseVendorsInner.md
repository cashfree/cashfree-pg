# SplitOrderReconSuccessResponseVendorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** | Unique identifier for the vendor. | [optional] 
**SettlementId** | Pointer to **int64** | Settlement ID associated with the vendor. | [optional] 
**SettlementAmount** | Pointer to **float32** | Settlement amount allocated to the vendor. | [optional] 
**SettlementEligibilityDate** | Pointer to **time.Time** | Date and time when the vendor is eligible for the settlement. | [optional] 

## Methods

### NewSplitOrderReconSuccessResponseVendorsInner

`func NewSplitOrderReconSuccessResponseVendorsInner() *SplitOrderReconSuccessResponseVendorsInner`

NewSplitOrderReconSuccessResponseVendorsInner instantiates a new SplitOrderReconSuccessResponseVendorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitOrderReconSuccessResponseVendorsInnerWithDefaults

`func NewSplitOrderReconSuccessResponseVendorsInnerWithDefaults() *SplitOrderReconSuccessResponseVendorsInner`

NewSplitOrderReconSuccessResponseVendorsInnerWithDefaults instantiates a new SplitOrderReconSuccessResponseVendorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *SplitOrderReconSuccessResponseVendorsInner) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *SplitOrderReconSuccessResponseVendorsInner) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetSettlementId

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementId() int64`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementIdOk() (*int64, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *SplitOrderReconSuccessResponseVendorsInner) SetSettlementId(v int64)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *SplitOrderReconSuccessResponseVendorsInner) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *SplitOrderReconSuccessResponseVendorsInner) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *SplitOrderReconSuccessResponseVendorsInner) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetSettlementEligibilityDate

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementEligibilityDate() time.Time`

GetSettlementEligibilityDate returns the SettlementEligibilityDate field if non-nil, zero value otherwise.

### GetSettlementEligibilityDateOk

`func (o *SplitOrderReconSuccessResponseVendorsInner) GetSettlementEligibilityDateOk() (*time.Time, bool)`

GetSettlementEligibilityDateOk returns a tuple with the SettlementEligibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementEligibilityDate

`func (o *SplitOrderReconSuccessResponseVendorsInner) SetSettlementEligibilityDate(v time.Time)`

SetSettlementEligibilityDate sets SettlementEligibilityDate field to given value.

### HasSettlementEligibilityDate

`func (o *SplitOrderReconSuccessResponseVendorsInner) HasSettlementEligibilityDate() bool`

HasSettlementEligibilityDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


