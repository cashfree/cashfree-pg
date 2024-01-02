# RefundEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **int64** | Cashfree Payments ID of the payment for which refund is initiated | [optional] 
**CfRefundId** | Pointer to **string** | Cashfree Payments ID for a refund | [optional] 
**OrderId** | Pointer to **string** | Merchant’s order Id of the order for which refund is initiated | [optional] 
**RefundId** | Pointer to **string** | Merchant’s refund ID of the refund | [optional] 
**Entity** | Pointer to **string** | Type of object | [optional] 
**RefundAmount** | Pointer to **float32** | Amount that is refunded | [optional] 
**RefundCurrency** | Pointer to **string** | Currency of the refund amount | [optional] 
**RefundNote** | Pointer to **string** | Note added by merchant for the refund | [optional] 
**RefundStatus** | Pointer to **string** | This can be one of [\&quot;SUCCESS\&quot;, \&quot;PENDING\&quot;, \&quot;CANCELLED\&quot;, \&quot;ONHOLD\&quot;, \&quot;FAILED\&quot;] | [optional] 
**RefundArn** | Pointer to **string** | The bank reference number for refund | [optional] 
**RefundCharge** | Pointer to **float32** | Charges in INR for processing refund | [optional] 
**StatusDescription** | Pointer to **string** | Description of refund status | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs | [optional] 
**RefundSplits** | Pointer to [**[]VendorSplit**](VendorSplit.md) |  | [optional] 
**RefundType** | Pointer to **string** | This can be one of [\&quot;PAYMENT_AUTO_REFUND\&quot;, \&quot;MERCHANT_INITIATED\&quot;, \&quot;UNRECONCILED_AUTO_REFUND\&quot;] | [optional] 
**RefundMode** | Pointer to **string** | Method or speed of processing refund | [optional] 
**CreatedAt** | Pointer to **string** | Time of refund creation | [optional] 
**ProcessedAt** | Pointer to **string** | Time when refund was processed successfully | [optional] 
**RefundSpeed** | Pointer to [**RefundSpeed**](RefundSpeed.md) |  | [optional] 

## Methods

### NewRefundEntity

`func NewRefundEntity() *RefundEntity`

NewRefundEntity instantiates a new RefundEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundEntityWithDefaults

`func NewRefundEntityWithDefaults() *RefundEntity`

NewRefundEntityWithDefaults instantiates a new RefundEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *RefundEntity) GetCfPaymentId() int64`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *RefundEntity) GetCfPaymentIdOk() (*int64, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *RefundEntity) SetCfPaymentId(v int64)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *RefundEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfRefundId

`func (o *RefundEntity) GetCfRefundId() string`

GetCfRefundId returns the CfRefundId field if non-nil, zero value otherwise.

### GetCfRefundIdOk

`func (o *RefundEntity) GetCfRefundIdOk() (*string, bool)`

GetCfRefundIdOk returns a tuple with the CfRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfRefundId

`func (o *RefundEntity) SetCfRefundId(v string)`

SetCfRefundId sets CfRefundId field to given value.

### HasCfRefundId

`func (o *RefundEntity) HasCfRefundId() bool`

HasCfRefundId returns a boolean if a field has been set.

### GetOrderId

`func (o *RefundEntity) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *RefundEntity) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *RefundEntity) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *RefundEntity) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetRefundId

`func (o *RefundEntity) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *RefundEntity) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *RefundEntity) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *RefundEntity) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetEntity

`func (o *RefundEntity) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RefundEntity) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RefundEntity) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RefundEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRefundAmount

`func (o *RefundEntity) GetRefundAmount() float32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RefundEntity) GetRefundAmountOk() (*float32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RefundEntity) SetRefundAmount(v float32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *RefundEntity) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundCurrency

`func (o *RefundEntity) GetRefundCurrency() string`

GetRefundCurrency returns the RefundCurrency field if non-nil, zero value otherwise.

### GetRefundCurrencyOk

`func (o *RefundEntity) GetRefundCurrencyOk() (*string, bool)`

GetRefundCurrencyOk returns a tuple with the RefundCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCurrency

`func (o *RefundEntity) SetRefundCurrency(v string)`

SetRefundCurrency sets RefundCurrency field to given value.

### HasRefundCurrency

`func (o *RefundEntity) HasRefundCurrency() bool`

HasRefundCurrency returns a boolean if a field has been set.

### GetRefundNote

`func (o *RefundEntity) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *RefundEntity) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *RefundEntity) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *RefundEntity) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundStatus

`func (o *RefundEntity) GetRefundStatus() string`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *RefundEntity) GetRefundStatusOk() (*string, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *RefundEntity) SetRefundStatus(v string)`

SetRefundStatus sets RefundStatus field to given value.

### HasRefundStatus

`func (o *RefundEntity) HasRefundStatus() bool`

HasRefundStatus returns a boolean if a field has been set.

### GetRefundArn

`func (o *RefundEntity) GetRefundArn() string`

GetRefundArn returns the RefundArn field if non-nil, zero value otherwise.

### GetRefundArnOk

`func (o *RefundEntity) GetRefundArnOk() (*string, bool)`

GetRefundArnOk returns a tuple with the RefundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundArn

`func (o *RefundEntity) SetRefundArn(v string)`

SetRefundArn sets RefundArn field to given value.

### HasRefundArn

`func (o *RefundEntity) HasRefundArn() bool`

HasRefundArn returns a boolean if a field has been set.

### GetRefundCharge

`func (o *RefundEntity) GetRefundCharge() float32`

GetRefundCharge returns the RefundCharge field if non-nil, zero value otherwise.

### GetRefundChargeOk

`func (o *RefundEntity) GetRefundChargeOk() (*float32, bool)`

GetRefundChargeOk returns a tuple with the RefundCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCharge

`func (o *RefundEntity) SetRefundCharge(v float32)`

SetRefundCharge sets RefundCharge field to given value.

### HasRefundCharge

`func (o *RefundEntity) HasRefundCharge() bool`

HasRefundCharge returns a boolean if a field has been set.

### GetStatusDescription

`func (o *RefundEntity) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *RefundEntity) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *RefundEntity) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *RefundEntity) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *RefundEntity) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RefundEntity) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RefundEntity) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RefundEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefundSplits

`func (o *RefundEntity) GetRefundSplits() []VendorSplit`

GetRefundSplits returns the RefundSplits field if non-nil, zero value otherwise.

### GetRefundSplitsOk

`func (o *RefundEntity) GetRefundSplitsOk() (*[]VendorSplit, bool)`

GetRefundSplitsOk returns a tuple with the RefundSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSplits

`func (o *RefundEntity) SetRefundSplits(v []VendorSplit)`

SetRefundSplits sets RefundSplits field to given value.

### HasRefundSplits

`func (o *RefundEntity) HasRefundSplits() bool`

HasRefundSplits returns a boolean if a field has been set.

### GetRefundType

`func (o *RefundEntity) GetRefundType() string`

GetRefundType returns the RefundType field if non-nil, zero value otherwise.

### GetRefundTypeOk

`func (o *RefundEntity) GetRefundTypeOk() (*string, bool)`

GetRefundTypeOk returns a tuple with the RefundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundType

`func (o *RefundEntity) SetRefundType(v string)`

SetRefundType sets RefundType field to given value.

### HasRefundType

`func (o *RefundEntity) HasRefundType() bool`

HasRefundType returns a boolean if a field has been set.

### GetRefundMode

`func (o *RefundEntity) GetRefundMode() string`

GetRefundMode returns the RefundMode field if non-nil, zero value otherwise.

### GetRefundModeOk

`func (o *RefundEntity) GetRefundModeOk() (*string, bool)`

GetRefundModeOk returns a tuple with the RefundMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundMode

`func (o *RefundEntity) SetRefundMode(v string)`

SetRefundMode sets RefundMode field to given value.

### HasRefundMode

`func (o *RefundEntity) HasRefundMode() bool`

HasRefundMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RefundEntity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefundEntity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefundEntity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RefundEntity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProcessedAt

`func (o *RefundEntity) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *RefundEntity) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *RefundEntity) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *RefundEntity) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetRefundSpeed

`func (o *RefundEntity) GetRefundSpeed() RefundSpeed`

GetRefundSpeed returns the RefundSpeed field if non-nil, zero value otherwise.

### GetRefundSpeedOk

`func (o *RefundEntity) GetRefundSpeedOk() (*RefundSpeed, bool)`

GetRefundSpeedOk returns a tuple with the RefundSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSpeed

`func (o *RefundEntity) SetRefundSpeed(v RefundSpeed)`

SetRefundSpeed sets RefundSpeed field to given value.

### HasRefundSpeed

`func (o *RefundEntity) HasRefundSpeed() bool`

HasRefundSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


