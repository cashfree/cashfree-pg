# UpdateOrderExtendedDataEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfOrderId** | Pointer to **string** | unique id generated by cashfree for your order | [optional] 
**OrderId** | Pointer to **string** | order_id sent during the api request | [optional] 
**ShipmentDetails** | Pointer to [**[]ShipmentDetails**](ShipmentDetails.md) |  | [optional] 
**OrderDeliveryStatus** | Pointer to [**OrderDeliveryStatus**](OrderDeliveryStatus.md) |  | [optional] 

## Methods

### NewUpdateOrderExtendedDataEntity

`func NewUpdateOrderExtendedDataEntity() *UpdateOrderExtendedDataEntity`

NewUpdateOrderExtendedDataEntity instantiates a new UpdateOrderExtendedDataEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderExtendedDataEntityWithDefaults

`func NewUpdateOrderExtendedDataEntityWithDefaults() *UpdateOrderExtendedDataEntity`

NewUpdateOrderExtendedDataEntityWithDefaults instantiates a new UpdateOrderExtendedDataEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfOrderId

`func (o *UpdateOrderExtendedDataEntity) GetCfOrderId() string`

GetCfOrderId returns the CfOrderId field if non-nil, zero value otherwise.

### GetCfOrderIdOk

`func (o *UpdateOrderExtendedDataEntity) GetCfOrderIdOk() (*string, bool)`

GetCfOrderIdOk returns a tuple with the CfOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfOrderId

`func (o *UpdateOrderExtendedDataEntity) SetCfOrderId(v string)`

SetCfOrderId sets CfOrderId field to given value.

### HasCfOrderId

`func (o *UpdateOrderExtendedDataEntity) HasCfOrderId() bool`

HasCfOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *UpdateOrderExtendedDataEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UpdateOrderExtendedDataEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UpdateOrderExtendedDataEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UpdateOrderExtendedDataEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *UpdateOrderExtendedDataEntity) GetShipmentDetails() []ShipmentDetails`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *UpdateOrderExtendedDataEntity) GetShipmentDetailsOk() (*[]ShipmentDetails, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *UpdateOrderExtendedDataEntity) SetShipmentDetails(v []ShipmentDetails)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *UpdateOrderExtendedDataEntity) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetOrderDeliveryStatus

`func (o *UpdateOrderExtendedDataEntity) GetOrderDeliveryStatus() OrderDeliveryStatus`

GetOrderDeliveryStatus returns the OrderDeliveryStatus field if non-nil, zero value otherwise.

### GetOrderDeliveryStatusOk

`func (o *UpdateOrderExtendedDataEntity) GetOrderDeliveryStatusOk() (*OrderDeliveryStatus, bool)`

GetOrderDeliveryStatusOk returns a tuple with the OrderDeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDeliveryStatus

`func (o *UpdateOrderExtendedDataEntity) SetOrderDeliveryStatus(v OrderDeliveryStatus)`

SetOrderDeliveryStatus sets OrderDeliveryStatus field to given value.

### HasOrderDeliveryStatus

`func (o *UpdateOrderExtendedDataEntity) HasOrderDeliveryStatus() bool`

HasOrderDeliveryStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

