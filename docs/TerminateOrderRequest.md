# TerminateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderStatus** | **string** | To terminate an order, pass order_status as \&quot;TERMINATE\&quot;. Please note, order might not be terminated - confirm with the order_status in response. \&quot;TERMINATION_REQUESTED\&quot; states that the request is recieved and we are working on it. If the order terminates successfully, status will change to \&quot;TERMINATED\&quot;. Incase there&#39;s any active transaction which moved to success - order might not get terminated. | 

## Methods

### NewTerminateOrderRequest

`func NewTerminateOrderRequest(orderStatus string, ) *TerminateOrderRequest`

NewTerminateOrderRequest instantiates a new TerminateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateOrderRequestWithDefaults

`func NewTerminateOrderRequestWithDefaults() *TerminateOrderRequest`

NewTerminateOrderRequestWithDefaults instantiates a new TerminateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderStatus

`func (o *TerminateOrderRequest) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *TerminateOrderRequest) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *TerminateOrderRequest) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


