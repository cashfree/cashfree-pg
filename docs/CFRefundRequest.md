# CFRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundAmount** | **float64** | Amount to be refunded. Should be lesser than or equal to the transaction amount. (Decimals allowed) | 
**RefundId** | **string** | An unique ID to associate the refund with. Provie alphanumeric values | 
**RefundNote** | Pointer to **string** | A refund note for your reference. | [optional] 
**RefundSplits** | Pointer to [**[]CFVendorSplit**](CFVendorSplit.md) |  | [optional] 

## Methods

### NewCFRefundRequest

`func NewCFRefundRequest(refundAmount float64, refundId string, ) *CFRefundRequest`

NewCFRefundRequest instantiates a new CFRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFRefundRequestWithDefaults

`func NewCFRefundRequestWithDefaults() *CFRefundRequest`

NewCFRefundRequestWithDefaults instantiates a new CFRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundAmount

`func (o *CFRefundRequest) GetRefundAmount() float64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *CFRefundRequest) GetRefundAmountOk() (*float64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *CFRefundRequest) SetRefundAmount(v float64)`

SetRefundAmount sets RefundAmount field to given value.


### GetRefundId

`func (o *CFRefundRequest) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *CFRefundRequest) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *CFRefundRequest) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetRefundNote

`func (o *CFRefundRequest) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *CFRefundRequest) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *CFRefundRequest) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *CFRefundRequest) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundSplits

`func (o *CFRefundRequest) GetRefundSplits() []CFVendorSplit`

GetRefundSplits returns the RefundSplits field if non-nil, zero value otherwise.

### GetRefundSplitsOk

`func (o *CFRefundRequest) GetRefundSplitsOk() (*[]CFVendorSplit, bool)`

GetRefundSplitsOk returns a tuple with the RefundSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSplits

`func (o *CFRefundRequest) SetRefundSplits(v []CFVendorSplit)`

SetRefundSplits sets RefundSplits field to given value.

### HasRefundSplits

`func (o *CFRefundRequest) HasRefundSplits() bool`

HasRefundSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


