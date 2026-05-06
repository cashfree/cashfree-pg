# OrderUpdateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundStatus** | **string** | Allowed values: [\&quot;CANCELLED\&quot;]. | 
**Remarks** | Pointer to **string** | Add remarks for your reference. | [optional] 

## Methods

### NewOrderUpdateRefundRequest

`func NewOrderUpdateRefundRequest(refundStatus string, ) *OrderUpdateRefundRequest`

NewOrderUpdateRefundRequest instantiates a new OrderUpdateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateRefundRequestWithDefaults

`func NewOrderUpdateRefundRequestWithDefaults() *OrderUpdateRefundRequest`

NewOrderUpdateRefundRequestWithDefaults instantiates a new OrderUpdateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundStatus

`func (o *OrderUpdateRefundRequest) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *OrderUpdateRefundRequest) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *OrderUpdateRefundRequest) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.


### GetRemarks

`func (o *OrderUpdateRefundRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *OrderUpdateRefundRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *OrderUpdateRefundRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *OrderUpdateRefundRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


