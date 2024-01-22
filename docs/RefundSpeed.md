# RefundSpeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | Pointer to **string** | Requested speed of refund. | [optional] 
**Accepted** | Pointer to **string** | Accepted speed of refund. | [optional] 
**Processed** | Pointer to **string** | Processed speed of refund. | [optional] 
**Message** | Pointer to **string** | Error message, if any for refund_speed request | [optional] 

## Methods

### NewRefundSpeed

`func NewRefundSpeed() *RefundSpeed`

NewRefundSpeed instantiates a new RefundSpeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundSpeedWithDefaults

`func NewRefundSpeedWithDefaults() *RefundSpeed`

NewRefundSpeedWithDefaults instantiates a new RefundSpeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *RefundSpeed) GetRequested() string`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RefundSpeed) GetRequestedOk() (*string, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RefundSpeed) SetRequested(v string)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *RefundSpeed) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetAccepted

`func (o *RefundSpeed) GetAccepted() string`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *RefundSpeed) GetAcceptedOk() (*string, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *RefundSpeed) SetAccepted(v string)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *RefundSpeed) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetProcessed

`func (o *RefundSpeed) GetProcessed() string`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *RefundSpeed) GetProcessedOk() (*string, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *RefundSpeed) SetProcessed(v string)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *RefundSpeed) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetMessage

`func (o *RefundSpeed) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RefundSpeed) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RefundSpeed) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RefundSpeed) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


