# UpdateOrderExtendedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentDetails** | [**[]ShipmentDetails**](ShipmentDetails.md) | Shipment details, such as the tracking company, tracking number, and tracking URLs, associated with the shipping of an order. Either &#x60;shipment_details&#x60; or &#x60;order_delivery_status&#x60; is required. | 
**OrderDeliveryStatus** | Pointer to [**OrderDeliveryStatus**](OrderDeliveryStatus.md) |  | [optional] 

## Methods

### NewUpdateOrderExtendedRequest

`func NewUpdateOrderExtendedRequest(shipmentDetails []ShipmentDetails, ) *UpdateOrderExtendedRequest`

NewUpdateOrderExtendedRequest instantiates a new UpdateOrderExtendedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderExtendedRequestWithDefaults

`func NewUpdateOrderExtendedRequestWithDefaults() *UpdateOrderExtendedRequest`

NewUpdateOrderExtendedRequestWithDefaults instantiates a new UpdateOrderExtendedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentDetails

`func (o *UpdateOrderExtendedRequest) GetShipmentDetails() []ShipmentDetails`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *UpdateOrderExtendedRequest) GetShipmentDetailsOk() (*[]ShipmentDetails, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *UpdateOrderExtendedRequest) SetShipmentDetails(v []ShipmentDetails)`

SetShipmentDetails sets ShipmentDetails field to given value.


### GetOrderDeliveryStatus

`func (o *UpdateOrderExtendedRequest) GetOrderDeliveryStatus() OrderDeliveryStatus`

GetOrderDeliveryStatus returns the OrderDeliveryStatus field if non-nil, zero value otherwise.

### GetOrderDeliveryStatusOk

`func (o *UpdateOrderExtendedRequest) GetOrderDeliveryStatusOk() (*OrderDeliveryStatus, bool)`

GetOrderDeliveryStatusOk returns a tuple with the OrderDeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDeliveryStatus

`func (o *UpdateOrderExtendedRequest) SetOrderDeliveryStatus(v OrderDeliveryStatus)`

SetOrderDeliveryStatus sets OrderDeliveryStatus field to given value.

### HasOrderDeliveryStatus

`func (o *UpdateOrderExtendedRequest) HasOrderDeliveryStatus() bool`

HasOrderDeliveryStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


