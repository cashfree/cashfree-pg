# SplitAfterPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Split** | [**[]SplitAfterPaymentRequestSplitInner**](SplitAfterPaymentRequestSplitInner.md) | Specify the vendors order split details. | 
**DisableSplit** | Pointer to **bool** | Specify if you want to end the split or continue creating further splits in future. | [optional] 

## Methods

### NewSplitAfterPaymentRequest

`func NewSplitAfterPaymentRequest(split []SplitAfterPaymentRequestSplitInner, ) *SplitAfterPaymentRequest`

NewSplitAfterPaymentRequest instantiates a new SplitAfterPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitAfterPaymentRequestWithDefaults

`func NewSplitAfterPaymentRequestWithDefaults() *SplitAfterPaymentRequest`

NewSplitAfterPaymentRequestWithDefaults instantiates a new SplitAfterPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplit

`func (o *SplitAfterPaymentRequest) GetSplit() []SplitAfterPaymentRequestSplitInner`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *SplitAfterPaymentRequest) GetSplitOk() (*[]SplitAfterPaymentRequestSplitInner, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *SplitAfterPaymentRequest) SetSplit(v []SplitAfterPaymentRequestSplitInner)`

SetSplit sets Split field to given value.


### GetDisableSplit

`func (o *SplitAfterPaymentRequest) GetDisableSplit() bool`

GetDisableSplit returns the DisableSplit field if non-nil, zero value otherwise.

### GetDisableSplitOk

`func (o *SplitAfterPaymentRequest) GetDisableSplitOk() (*bool, bool)`

GetDisableSplitOk returns a tuple with the DisableSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSplit

`func (o *SplitAfterPaymentRequest) SetDisableSplit(v bool)`

SetDisableSplit sets DisableSplit field to given value.

### HasDisableSplit

`func (o *SplitAfterPaymentRequest) HasDisableSplit() bool`

HasDisableSplit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


