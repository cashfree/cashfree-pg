# ReconEntityDataInnerEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Unique ID associated with the event. | [optional] 
**EventType** | Pointer to **string** | The event type can be PAYMENT, REFUND, REFUND_REVERSAL, DISPUTE, DISPUTE_REVERSAL, CHARGEBACK, CHARGEBACK_REVERSAL, OTHER_ADJUSTMENT. | [optional] 
**EventSettlementAmount** | Pointer to **float32** | Amount that is part of the settlement corresponding to the event. | [optional] 
**EventAmount** | Pointer to **float32** | Amount corresponding to the event. Example, refund amount, dispute amount, payment amount, etc. | [optional] 
**SaleType** | Pointer to **string** | Indicates if it is CREDIT/DEBIT sale. | [optional] 
**EventStatus** | Pointer to **string** | Status of the event. Example - SUCCESS, FAILED, PENDING, CANCELLED. | [optional] 
**Entity** | Pointer to **string** | Recon. | [optional] 
**EventTime** | Pointer to **string** | Time associated with the event. Example, transaction time, dispute initiation time. | [optional] 
**EventCurrency** | Pointer to **string** | Curreny type - INR. | [optional] 
**EventServiceCharge** | Pointer to **float32** | Service charge for above event_type. | [optional] 
**EventServiceTax** | Pointer to **float32** | Service tax for above event_type. | [optional] 
**EventRemarks** | Pointer to **float32** | Remarks for above event_type. | [optional] 

## Methods

### NewReconEntityDataInnerEventDetails

`func NewReconEntityDataInnerEventDetails() *ReconEntityDataInnerEventDetails`

NewReconEntityDataInnerEventDetails instantiates a new ReconEntityDataInnerEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityDataInnerEventDetailsWithDefaults

`func NewReconEntityDataInnerEventDetailsWithDefaults() *ReconEntityDataInnerEventDetails`

NewReconEntityDataInnerEventDetailsWithDefaults instantiates a new ReconEntityDataInnerEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ReconEntityDataInnerEventDetails) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ReconEntityDataInnerEventDetails) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ReconEntityDataInnerEventDetails) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *ReconEntityDataInnerEventDetails) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *ReconEntityDataInnerEventDetails) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ReconEntityDataInnerEventDetails) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ReconEntityDataInnerEventDetails) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ReconEntityDataInnerEventDetails) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventSettlementAmount

`func (o *ReconEntityDataInnerEventDetails) GetEventSettlementAmount() float32`

GetEventSettlementAmount returns the EventSettlementAmount field if non-nil, zero value otherwise.

### GetEventSettlementAmountOk

`func (o *ReconEntityDataInnerEventDetails) GetEventSettlementAmountOk() (*float32, bool)`

GetEventSettlementAmountOk returns a tuple with the EventSettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSettlementAmount

`func (o *ReconEntityDataInnerEventDetails) SetEventSettlementAmount(v float32)`

SetEventSettlementAmount sets EventSettlementAmount field to given value.

### HasEventSettlementAmount

`func (o *ReconEntityDataInnerEventDetails) HasEventSettlementAmount() bool`

HasEventSettlementAmount returns a boolean if a field has been set.

### GetEventAmount

`func (o *ReconEntityDataInnerEventDetails) GetEventAmount() float32`

GetEventAmount returns the EventAmount field if non-nil, zero value otherwise.

### GetEventAmountOk

`func (o *ReconEntityDataInnerEventDetails) GetEventAmountOk() (*float32, bool)`

GetEventAmountOk returns a tuple with the EventAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAmount

`func (o *ReconEntityDataInnerEventDetails) SetEventAmount(v float32)`

SetEventAmount sets EventAmount field to given value.

### HasEventAmount

`func (o *ReconEntityDataInnerEventDetails) HasEventAmount() bool`

HasEventAmount returns a boolean if a field has been set.

### GetSaleType

`func (o *ReconEntityDataInnerEventDetails) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *ReconEntityDataInnerEventDetails) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *ReconEntityDataInnerEventDetails) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *ReconEntityDataInnerEventDetails) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetEventStatus

`func (o *ReconEntityDataInnerEventDetails) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *ReconEntityDataInnerEventDetails) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *ReconEntityDataInnerEventDetails) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *ReconEntityDataInnerEventDetails) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEntity

`func (o *ReconEntityDataInnerEventDetails) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ReconEntityDataInnerEventDetails) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ReconEntityDataInnerEventDetails) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ReconEntityDataInnerEventDetails) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventTime

`func (o *ReconEntityDataInnerEventDetails) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ReconEntityDataInnerEventDetails) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ReconEntityDataInnerEventDetails) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ReconEntityDataInnerEventDetails) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventCurrency

`func (o *ReconEntityDataInnerEventDetails) GetEventCurrency() string`

GetEventCurrency returns the EventCurrency field if non-nil, zero value otherwise.

### GetEventCurrencyOk

`func (o *ReconEntityDataInnerEventDetails) GetEventCurrencyOk() (*string, bool)`

GetEventCurrencyOk returns a tuple with the EventCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCurrency

`func (o *ReconEntityDataInnerEventDetails) SetEventCurrency(v string)`

SetEventCurrency sets EventCurrency field to given value.

### HasEventCurrency

`func (o *ReconEntityDataInnerEventDetails) HasEventCurrency() bool`

HasEventCurrency returns a boolean if a field has been set.

### GetEventServiceCharge

`func (o *ReconEntityDataInnerEventDetails) GetEventServiceCharge() float32`

GetEventServiceCharge returns the EventServiceCharge field if non-nil, zero value otherwise.

### GetEventServiceChargeOk

`func (o *ReconEntityDataInnerEventDetails) GetEventServiceChargeOk() (*float32, bool)`

GetEventServiceChargeOk returns a tuple with the EventServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceCharge

`func (o *ReconEntityDataInnerEventDetails) SetEventServiceCharge(v float32)`

SetEventServiceCharge sets EventServiceCharge field to given value.

### HasEventServiceCharge

`func (o *ReconEntityDataInnerEventDetails) HasEventServiceCharge() bool`

HasEventServiceCharge returns a boolean if a field has been set.

### GetEventServiceTax

`func (o *ReconEntityDataInnerEventDetails) GetEventServiceTax() float32`

GetEventServiceTax returns the EventServiceTax field if non-nil, zero value otherwise.

### GetEventServiceTaxOk

`func (o *ReconEntityDataInnerEventDetails) GetEventServiceTaxOk() (*float32, bool)`

GetEventServiceTaxOk returns a tuple with the EventServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceTax

`func (o *ReconEntityDataInnerEventDetails) SetEventServiceTax(v float32)`

SetEventServiceTax sets EventServiceTax field to given value.

### HasEventServiceTax

`func (o *ReconEntityDataInnerEventDetails) HasEventServiceTax() bool`

HasEventServiceTax returns a boolean if a field has been set.

### GetEventRemarks

`func (o *ReconEntityDataInnerEventDetails) GetEventRemarks() float32`

GetEventRemarks returns the EventRemarks field if non-nil, zero value otherwise.

### GetEventRemarksOk

`func (o *ReconEntityDataInnerEventDetails) GetEventRemarksOk() (*float32, bool)`

GetEventRemarksOk returns a tuple with the EventRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRemarks

`func (o *ReconEntityDataInnerEventDetails) SetEventRemarks(v float32)`

SetEventRemarks sets EventRemarks field to given value.

### HasEventRemarks

`func (o *ReconEntityDataInnerEventDetails) HasEventRemarks() bool`

HasEventRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


