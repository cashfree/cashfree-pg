# SettlementReconEntityDataInnerEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Unique ID associated with the event. | [optional] 
**EventType** | Pointer to **string** | The event type can be PAYMENT, REFUND, REFUND_REVERSAL, DISPUTE, DISPUTE_REVERSAL, CHARGEBACK, CHARGEBACK_REVERSAL, OTHER_ADJUSTMENT. | [optional] 
**EventSettlementAmount** | Pointer to **float32** | Amount that is part of the settlement corresponding to the event. | [optional] 
**EventAmount** | Pointer to **float32** | Amount corresponding to the event. Example, refund amount, dispute amount, payment amount, etc. | [optional] 
**SaleType** | Pointer to **string** | Indicates if it is CREDIT/DEBIT sale. | [optional] 
**EventStatus** | Pointer to **string** | Status of the event. Example - SUCCESS, FAILED, PENDING, CANCELLED. | [optional] 
**Entity** | Pointer to **string** | Recon | [optional] 
**EventTime** | Pointer to **string** | Time associated with the event. Example, transaction time, dispute initiation time | [optional] 
**EventCurrency** | Pointer to **string** | Curreny type - INR. | [optional] 
**EventServiceCharge** | Pointer to **float32** | Service charge for above event_type. | [optional] 
**EventServiceTax** | Pointer to **float32** | Service tax for above event_type. | [optional] 
**EventRemarks** | Pointer to **float32** | Remarks for above event_type. | [optional] 

## Methods

### NewSettlementReconEntityDataInnerEventDetails

`func NewSettlementReconEntityDataInnerEventDetails() *SettlementReconEntityDataInnerEventDetails`

NewSettlementReconEntityDataInnerEventDetails instantiates a new SettlementReconEntityDataInnerEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerEventDetailsWithDefaults

`func NewSettlementReconEntityDataInnerEventDetailsWithDefaults() *SettlementReconEntityDataInnerEventDetails`

NewSettlementReconEntityDataInnerEventDetailsWithDefaults instantiates a new SettlementReconEntityDataInnerEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventSettlementAmount

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventSettlementAmount() float32`

GetEventSettlementAmount returns the EventSettlementAmount field if non-nil, zero value otherwise.

### GetEventSettlementAmountOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventSettlementAmountOk() (*float32, bool)`

GetEventSettlementAmountOk returns a tuple with the EventSettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSettlementAmount

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventSettlementAmount(v float32)`

SetEventSettlementAmount sets EventSettlementAmount field to given value.

### HasEventSettlementAmount

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventSettlementAmount() bool`

HasEventSettlementAmount returns a boolean if a field has been set.

### GetEventAmount

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventAmount() float32`

GetEventAmount returns the EventAmount field if non-nil, zero value otherwise.

### GetEventAmountOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventAmountOk() (*float32, bool)`

GetEventAmountOk returns a tuple with the EventAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAmount

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventAmount(v float32)`

SetEventAmount sets EventAmount field to given value.

### HasEventAmount

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventAmount() bool`

HasEventAmount returns a boolean if a field has been set.

### GetSaleType

`func (o *SettlementReconEntityDataInnerEventDetails) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *SettlementReconEntityDataInnerEventDetails) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *SettlementReconEntityDataInnerEventDetails) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetEventStatus

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEntity

`func (o *SettlementReconEntityDataInnerEventDetails) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SettlementReconEntityDataInnerEventDetails) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SettlementReconEntityDataInnerEventDetails) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventTime

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventCurrency

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventCurrency() string`

GetEventCurrency returns the EventCurrency field if non-nil, zero value otherwise.

### GetEventCurrencyOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventCurrencyOk() (*string, bool)`

GetEventCurrencyOk returns a tuple with the EventCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCurrency

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventCurrency(v string)`

SetEventCurrency sets EventCurrency field to given value.

### HasEventCurrency

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventCurrency() bool`

HasEventCurrency returns a boolean if a field has been set.

### GetEventServiceCharge

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventServiceCharge() float32`

GetEventServiceCharge returns the EventServiceCharge field if non-nil, zero value otherwise.

### GetEventServiceChargeOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventServiceChargeOk() (*float32, bool)`

GetEventServiceChargeOk returns a tuple with the EventServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceCharge

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventServiceCharge(v float32)`

SetEventServiceCharge sets EventServiceCharge field to given value.

### HasEventServiceCharge

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventServiceCharge() bool`

HasEventServiceCharge returns a boolean if a field has been set.

### GetEventServiceTax

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventServiceTax() float32`

GetEventServiceTax returns the EventServiceTax field if non-nil, zero value otherwise.

### GetEventServiceTaxOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventServiceTaxOk() (*float32, bool)`

GetEventServiceTaxOk returns a tuple with the EventServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceTax

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventServiceTax(v float32)`

SetEventServiceTax sets EventServiceTax field to given value.

### HasEventServiceTax

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventServiceTax() bool`

HasEventServiceTax returns a boolean if a field has been set.

### GetEventRemarks

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventRemarks() float32`

GetEventRemarks returns the EventRemarks field if non-nil, zero value otherwise.

### GetEventRemarksOk

`func (o *SettlementReconEntityDataInnerEventDetails) GetEventRemarksOk() (*float32, bool)`

GetEventRemarksOk returns a tuple with the EventRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRemarks

`func (o *SettlementReconEntityDataInnerEventDetails) SetEventRemarks(v float32)`

SetEventRemarks sets EventRemarks field to given value.

### HasEventRemarks

`func (o *SettlementReconEntityDataInnerEventDetails) HasEventRemarks() bool`

HasEventRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


