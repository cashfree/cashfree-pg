# CFRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **int32** | Cashfree Payments ID of the payment for which refund is initiated | [optional] 
**CfRefundId** | Pointer to **string** | Cashfree Payments ID for a refund | [optional] 
**OrderId** | Pointer to **string** | Merchant’s order Id of the order for which refund is initiated | [optional] 
**RefundId** | Pointer to **string** | Merchant’s refund ID of the refund | [optional] 
**Entity** | Pointer to **string** | Type of object | [optional] 
**RefundAmount** | Pointer to **float32** | Amount that is refunded | [optional] 
**RefundCurrency** | Pointer to **string** | Currency of the refund amount | [optional] 
**RefundNote** | Pointer to **string** | Note added by merchant for the refund | [optional] 
**RefundStatus** | Pointer to **string** | This can be one of [\&quot;SUCCESS\&quot;, \&quot;PENDING\&quot;, \&quot;CANCELLED\&quot;, \&quot;ONHOLD\&quot;] | [optional] 
**RefundArn** | Pointer to **string** | The bank reference number for refund | [optional] 
**RefundCharge** | Pointer to **float32** | Charges in INR for processing refund | [optional] 
**StatusDescription** | Pointer to **string** | Description of refund status | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs | [optional] 
**RefundSplits** | Pointer to [**[]CFVendorSplit**](CFVendorSplit.md) |  | [optional] 
**RefundType** | Pointer to **string** | This can be one of [\&quot;PAYMENT_AUTO_REFUND\&quot;, \&quot;MERCHANT_INITIATED\&quot;, \&quot;UNRECONCILED_AUTO_REFUND\&quot;] | [optional] 
**RefundMode** | Pointer to **string** | Method or speed of processing refund | [optional] 
**CreatedAt** | Pointer to **string** | Time of refund creation | [optional] 
**ProcessedAt** | Pointer to **string** | Time when refund was processed successfully | [optional] 

## Methods

### NewCFRefund

`func NewCFRefund() *CFRefund`

NewCFRefund instantiates a new CFRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFRefundWithDefaults

`func NewCFRefundWithDefaults() *CFRefund`

NewCFRefundWithDefaults instantiates a new CFRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CFRefund) GetCfPaymentId() int32`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CFRefund) GetCfPaymentIdOk() (*int32, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CFRefund) SetCfPaymentId(v int32)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CFRefund) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfRefundId

`func (o *CFRefund) GetCfRefundId() string`

GetCfRefundId returns the CfRefundId field if non-nil, zero value otherwise.

### GetCfRefundIdOk

`func (o *CFRefund) GetCfRefundIdOk() (*string, bool)`

GetCfRefundIdOk returns a tuple with the CfRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfRefundId

`func (o *CFRefund) SetCfRefundId(v string)`

SetCfRefundId sets CfRefundId field to given value.

### HasCfRefundId

`func (o *CFRefund) HasCfRefundId() bool`

HasCfRefundId returns a boolean if a field has been set.

### GetOrderId

`func (o *CFRefund) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CFRefund) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CFRefund) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CFRefund) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetRefundId

`func (o *CFRefund) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *CFRefund) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *CFRefund) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *CFRefund) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetEntity

`func (o *CFRefund) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CFRefund) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CFRefund) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CFRefund) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRefundAmount

`func (o *CFRefund) GetRefundAmount() float32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *CFRefund) GetRefundAmountOk() (*float32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *CFRefund) SetRefundAmount(v float32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *CFRefund) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundCurrency

`func (o *CFRefund) GetRefundCurrency() string`

GetRefundCurrency returns the RefundCurrency field if non-nil, zero value otherwise.

### GetRefundCurrencyOk

`func (o *CFRefund) GetRefundCurrencyOk() (*string, bool)`

GetRefundCurrencyOk returns a tuple with the RefundCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCurrency

`func (o *CFRefund) SetRefundCurrency(v string)`

SetRefundCurrency sets RefundCurrency field to given value.

### HasRefundCurrency

`func (o *CFRefund) HasRefundCurrency() bool`

HasRefundCurrency returns a boolean if a field has been set.

### GetRefundNote

`func (o *CFRefund) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *CFRefund) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *CFRefund) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *CFRefund) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundStatus

`func (o *CFRefund) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *CFRefund) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *CFRefund) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.

### HasRefundStatus

`func (o *CFRefund) HasRefundStatus() bool`

HasRefundStatus returns a boolean if a field has been set.

### GetRefundArn

`func (o *CFRefund) GetRefundArn() string`

GetRefundArn returns the RefundArn field if non-nil, zero value otherwise.

### GetRefundArnOk

`func (o *CFRefund) GetRefundArnOk() (*string, bool)`

GetRefundArnOk returns a tuple with the RefundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundArn

`func (o *CFRefund) SetRefundArn(v string)`

SetRefundArn sets RefundArn field to given value.

### HasRefundArn

`func (o *CFRefund) HasRefundArn() bool`

HasRefundArn returns a boolean if a field has been set.

### GetRefundCharge

`func (o *CFRefund) GetRefundCharge() float32`

GetRefundCharge returns the RefundCharge field if non-nil, zero value otherwise.

### GetRefundChargeOk

`func (o *CFRefund) GetRefundChargeOk() (*float32, bool)`

GetRefundChargeOk returns a tuple with the RefundCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCharge

`func (o *CFRefund) SetRefundCharge(v float32)`

SetRefundCharge sets RefundCharge field to given value.

### HasRefundCharge

`func (o *CFRefund) HasRefundCharge() bool`

HasRefundCharge returns a boolean if a field has been set.

### GetStatusDescription

`func (o *CFRefund) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *CFRefund) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *CFRefund) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *CFRefund) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CFRefund) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CFRefund) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CFRefund) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CFRefund) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefundSplits

`func (o *CFRefund) GetRefundSplits() []CFVendorSplit`

GetRefundSplits returns the RefundSplits field if non-nil, zero value otherwise.

### GetRefundSplitsOk

`func (o *CFRefund) GetRefundSplitsOk() (*[]CFVendorSplit, bool)`

GetRefundSplitsOk returns a tuple with the RefundSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSplits

`func (o *CFRefund) SetRefundSplits(v []CFVendorSplit)`

SetRefundSplits sets RefundSplits field to given value.

### HasRefundSplits

`func (o *CFRefund) HasRefundSplits() bool`

HasRefundSplits returns a boolean if a field has been set.

### GetRefundType

`func (o *CFRefund) GetRefundType() string`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *CFRefund) GetRefundTypeOk() (*string, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *CFRefund) SetRefundType(v string)`

SetRefundType sets RefundType field to given value.

### HasRefundType

`func (o *CFRefund) HasRefundType() bool`

HasRefundType returns a boolean if a field has been set.

### GetRefundMode

`func (o *CFRefund) GetRefundMode() string`

GetRefundMode returns the RefundMode field if non-nil, zero value otherwise.

### GetRefundModeOk

`func (o *CFRefund) GetRefundModeOk() (*string, bool)`

GetRefundModeOk returns a tuple with the RefundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundMode

`func (o *CFRefund) SetRefundMode(v string)`

SetRefundMode sets RefundMode field to given value.

### HasRefundMode

`func (o *CFRefund) HasRefundMode() bool`

HasRefundMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CFRefund) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CFRefund) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CFRefund) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CFRefund) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProcessedAt

`func (o *CFRefund) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *CFRefund) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *CFRefund) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *CFRefund) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


