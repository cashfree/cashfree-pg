# OrderCreateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundAmount** | **float64** | Amount to be refunded. Should be lesser than or equal to the transaction amount. (Decimals allowed) | 
**RefundId** | **string** | An unique ID to associate the refund with. Provie alphanumeric values | 
**RefundNote** | Pointer to **string** | A refund note for your reference. | [optional] 
**RefundSpeed** | Pointer to **string** | Speed at which the refund is processed. It&#39;s an optional field with default being STANDARD | [optional] 
**RefundSplits** | Pointer to [**[]VendorSplit**](VendorSplit.md) |  | [optional] 

## Methods

### NewOrderCreateRefundRequest

`func NewOrderCreateRefundRequest(refundAmount float64, refundId string, ) *OrderCreateRefundRequest`

NewOrderCreateRefundRequest instantiates a new OrderCreateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRefundRequestWithDefaults

`func NewOrderCreateRefundRequestWithDefaults() *OrderCreateRefundRequest`

NewOrderCreateRefundRequestWithDefaults instantiates a new OrderCreateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundAmount

`func (o *OrderCreateRefundRequest) GetRefundAmount() float64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *OrderCreateRefundRequest) GetRefundAmountOk() (*float64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *OrderCreateRefundRequest) SetRefundAmount(v float64)`

SetRefundAmount sets RefundAmount field to given value.


### GetRefundId

`func (o *OrderCreateRefundRequest) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *OrderCreateRefundRequest) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *OrderCreateRefundRequest) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetRefundNote

`func (o *OrderCreateRefundRequest) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *OrderCreateRefundRequest) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *OrderCreateRefundRequest) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *OrderCreateRefundRequest) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundSpeed

`func (o *OrderCreateRefundRequest) GetRefundSpeed() string`

GetRefundSpeed returns the RefundSpeed field if non-nil, zero value otherwise.

### GetRefundSpeedOk

`func (o *OrderCreateRefundRequest) GetRefundSpeedOk() (*string, bool)`

GetRefundSpeedOk returns a tuple with the RefundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSpeed

`func (o *OrderCreateRefundRequest) SetRefundSpeed(v string)`

SetRefundSpeed sets RefundSpeed field to given value.

### HasRefundSpeed

`func (o *OrderCreateRefundRequest) HasRefundSpeed() bool`

HasRefundSpeed returns a boolean if a field has been set.

### GetRefundSplits

`func (o *OrderCreateRefundRequest) GetRefundSplits() []VendorSplit`

GetRefundSplits returns the RefundSplits field if non-nil, zero value otherwise.

### GetRefundSplitsOk

`func (o *OrderCreateRefundRequest) GetRefundSplitsOk() (*[]VendorSplit, bool)`

GetRefundSplitsOk returns a tuple with the RefundSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSplits

`func (o *OrderCreateRefundRequest) SetRefundSplits(v []VendorSplit)`

SetRefundSplits sets RefundSplits field to given value.

### HasRefundSplits

`func (o *OrderCreateRefundRequest) HasRefundSplits() bool`

HasRefundSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


