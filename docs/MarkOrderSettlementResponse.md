# MarkOrderSettlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | The order ID that was marked for settlement. | [optional] 
**MetaData** | Pointer to [**MarkOrderSettlementResponseMetaData**](MarkOrderSettlementResponseMetaData.md) |  | [optional] 

## Methods

### NewMarkOrderSettlementResponse

`func NewMarkOrderSettlementResponse() *MarkOrderSettlementResponse`

NewMarkOrderSettlementResponse instantiates a new MarkOrderSettlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkOrderSettlementResponseWithDefaults

`func NewMarkOrderSettlementResponseWithDefaults() *MarkOrderSettlementResponse`

NewMarkOrderSettlementResponseWithDefaults instantiates a new MarkOrderSettlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *MarkOrderSettlementResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarkOrderSettlementResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarkOrderSettlementResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarkOrderSettlementResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMetaData

`func (o *MarkOrderSettlementResponse) GetMetaData() MarkOrderSettlementResponseMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *MarkOrderSettlementResponse) GetMetaDataOk() (*MarkOrderSettlementResponseMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *MarkOrderSettlementResponse) SetMetaData(v MarkOrderSettlementResponseMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *MarkOrderSettlementResponse) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


