# OrderDeliveryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Delivery status of order | 
**Reason** | Pointer to **string** | Reason of provided order delivery status. This is optional field. | [optional] 

## Methods

### NewOrderDeliveryStatus

`func NewOrderDeliveryStatus(status string, ) *OrderDeliveryStatus`

NewOrderDeliveryStatus instantiates a new OrderDeliveryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDeliveryStatusWithDefaults

`func NewOrderDeliveryStatusWithDefaults() *OrderDeliveryStatus`

NewOrderDeliveryStatusWithDefaults instantiates a new OrderDeliveryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OrderDeliveryStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDeliveryStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDeliveryStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *OrderDeliveryStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OrderDeliveryStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OrderDeliveryStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OrderDeliveryStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


