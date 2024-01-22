# RefundWebhookDataEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refund** | Pointer to [**RefundEntity**](RefundEntity.md) |  | [optional] 

## Methods

### NewRefundWebhookDataEntity

`func NewRefundWebhookDataEntity() *RefundWebhookDataEntity`

NewRefundWebhookDataEntity instantiates a new RefundWebhookDataEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundWebhookDataEntityWithDefaults

`func NewRefundWebhookDataEntityWithDefaults() *RefundWebhookDataEntity`

NewRefundWebhookDataEntityWithDefaults instantiates a new RefundWebhookDataEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefund

`func (o *RefundWebhookDataEntity) GetRefund() RefundEntity`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *RefundWebhookDataEntity) GetRefundOk() (*RefundEntity, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *RefundWebhookDataEntity) SetRefund(v RefundEntity)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *RefundWebhookDataEntity) HasRefund() bool`

HasRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


