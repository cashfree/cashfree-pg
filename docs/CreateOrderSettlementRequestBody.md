# CreateOrderSettlementRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | OrderId of the order. | 
**MetaData** | [**CreateOrderSettlementRequestBodyMetaData**](CreateOrderSettlementRequestBodyMetaData.md) |  | 

## Methods

### NewCreateOrderSettlementRequestBody

`func NewCreateOrderSettlementRequestBody(orderId string, metaData CreateOrderSettlementRequestBodyMetaData, ) *CreateOrderSettlementRequestBody`

NewCreateOrderSettlementRequestBody instantiates a new CreateOrderSettlementRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderSettlementRequestBodyWithDefaults

`func NewCreateOrderSettlementRequestBodyWithDefaults() *CreateOrderSettlementRequestBody`

NewCreateOrderSettlementRequestBodyWithDefaults instantiates a new CreateOrderSettlementRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderSettlementRequestBody) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderSettlementRequestBody) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderSettlementRequestBody) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetMetaData

`func (o *CreateOrderSettlementRequestBody) GetMetaData() CreateOrderSettlementRequestBodyMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateOrderSettlementRequestBody) GetMetaDataOk() (*CreateOrderSettlementRequestBodyMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateOrderSettlementRequestBody) SetMetaData(v CreateOrderSettlementRequestBodyMetaData)`

SetMetaData sets MetaData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


