# ReconEntityDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Unique ID associated with the event. | [optional] 
**EventType** | Pointer to **string** | The event type can be SETTLEMENT, PAYMENT, REFUND, REFUND_REVERSAL, DISPUTE, DISPUTE_REVERSAL, CHARGEBACK, CHARGEBACK_REVERSAL, OTHER_ADJUSTMENT. | [optional] 
**EventSettlementAmount** | Pointer to **float32** | Amount that is part of the settlement corresponding to the event. | [optional] 
**EventAmount** | Pointer to **float32** | Amount of the event. Example, refund amount, dispute amount, payment amount, etc. | [optional] 
**SaleType** | Pointer to **string** | Indicates if it is CREDIT/DEBIT sale. | [optional] 
**EventStatus** | Pointer to **string** | Status of the event. Example - SUCCESS, FAILED, PENDING, CANCELLED. | [optional] 
**Entity** | Pointer to **string** | Recon | [optional] 
**EventTime** | Pointer to **string** | Time associated with the event. Example, transaction time, dispute initiation time | [optional] 
**EventCurrency** | Pointer to **string** | Curreny type - INR. | [optional] 
**OrderId** | Pointer to **string** | Unique order ID. Alphanumeric and only &#39;-&#39; and &#39;_&#39; allowed. | [optional] 
**OrderAmount** | Pointer to **float32** | The amount which was passed at the order creation time. | [optional] 
**CustomerPhone** | Pointer to **string** | Customer phone number. | [optional] 
**CustomerEmail** | Pointer to **string** | Customer email. | [optional] 
**CustomerName** | Pointer to **string** | Customer name. | [optional] 
**PaymentAmount** | Pointer to **float32** | Payment amount captured. | [optional] 
**PaymentUtr** | Pointer to **string** | Unique transaction reference number of the payment. | [optional] 
**PaymentTime** | Pointer to **string** | Date and time when the payment was initiated. | [optional] 
**PaymentServiceCharge** | Pointer to **float32** | Service charge applicable for the payment. | [optional] 
**PaymentServiceTax** | Pointer to **float32** | Service tax applicable on the payment. | [optional] 
**CfPaymentId** | Pointer to **string** | Cashfree Payments unique ID to identify a payment. | [optional] 
**CfSettlementId** | Pointer to **string** | Unique ID to identify the settlement. | [optional] 
**SettlementDate** | Pointer to **string** | Date and time when the settlement was processed. | [optional] 
**SettlementUtr** | Pointer to **string** | Unique transaction reference number of the settlement. | [optional] 
**SplitServiceCharge** | Pointer to **float32** | Service charge that is applicable for splitting the payment. | [optional] 
**SplitServiceTax** | Pointer to **float32** | Service tax applicable for splitting the amount to vendors. | [optional] 
**VendorCommission** | Pointer to **float32** | Vendor commission applicable for this transaction. | [optional] 
**ClosedInFavorOf** | Pointer to **string** | Specifies whether the dispute was closed in favor of the merchant or customer. /n Possible values - Merchant, Customer | [optional] 
**DisputeResolvedOn** | Pointer to **string** | Date and time when the dispute was resolved. | [optional] 
**DisputeCategory** | Pointer to **string** | Category of the dispute - Dispute code and the reason for dispute is shown. | [optional] 
**DisputeNote** | Pointer to **string** | Note regarding the dispute. | [optional] 
**RefundProcessedAt** | Pointer to **string** | Date and time when the refund was processed. | [optional] 
**RefundArn** | Pointer to **string** | The bank reference number for the refund. | [optional] 
**RefundNote** | Pointer to **string** | A refund note for your reference. | [optional] 
**RefundId** | Pointer to **string** | An unique ID to associate the refund with. | [optional] 
**AdjustmentRemarks** | Pointer to **string** | Other adjustment remarks. | [optional] 
**Adjustment** | Pointer to **float32** | Amount that is adjusted from the settlement amount because of any credit/debit event such as refund, refund_reverse etc. | [optional] 
**ServiceTax** | Pointer to **float32** | Service tax applicable on the settlement amount. | [optional] 
**ServiceCharge** | Pointer to **float32** | Service charge applicable on the settlement amount. | [optional] 
**AmountSettled** | Pointer to **float32** | Net amount that is settled after considering the adjustments, settlement charge and tax. | [optional] 
**PaymentFrom** | Pointer to **string** | The start time of the time range of the payments considered for the settlement. | [optional] 
**PaymentTill** | Pointer to **string** | The end time of time range of the payments considered for the settlement. | [optional] 
**Reason** | Pointer to **string** | Reason for settlement failure. | [optional] 
**SettlementInitiatedOn** | Pointer to **string** | Date and time when the settlement was initiated. | [optional] 
**SettlementType** | Pointer to **string** | Type of settlement. Possible values - Standard, Instant, On demand. | [optional] 
**SettlementCharge** | Pointer to **float32** | Settlement charges applicable on the settlement. | [optional] 
**SettlementTax** | Pointer to **float32** | Settlement tax applicable on the settlement. | [optional] 
**Remarks** | Pointer to **string** | Remarks on the settlement. | [optional] 

## Methods

### NewReconEntityDataInner

`func NewReconEntityDataInner() *ReconEntityDataInner`

NewReconEntityDataInner instantiates a new ReconEntityDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityDataInnerWithDefaults

`func NewReconEntityDataInnerWithDefaults() *ReconEntityDataInner`

NewReconEntityDataInnerWithDefaults instantiates a new ReconEntityDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ReconEntityDataInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ReconEntityDataInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ReconEntityDataInner) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *ReconEntityDataInner) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *ReconEntityDataInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ReconEntityDataInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ReconEntityDataInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ReconEntityDataInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventSettlementAmount

`func (o *ReconEntityDataInner) GetEventSettlementAmount() float32`

GetEventSettlementAmount returns the EventSettlementAmount field if non-nil, zero value otherwise.

### GetEventSettlementAmountOk

`func (o *ReconEntityDataInner) GetEventSettlementAmountOk() (*float32, bool)`

GetEventSettlementAmountOk returns a tuple with the EventSettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSettlementAmount

`func (o *ReconEntityDataInner) SetEventSettlementAmount(v float32)`

SetEventSettlementAmount sets EventSettlementAmount field to given value.

### HasEventSettlementAmount

`func (o *ReconEntityDataInner) HasEventSettlementAmount() bool`

HasEventSettlementAmount returns a boolean if a field has been set.

### GetEventAmount

`func (o *ReconEntityDataInner) GetEventAmount() float32`

GetEventAmount returns the EventAmount field if non-nil, zero value otherwise.

### GetEventAmountOk

`func (o *ReconEntityDataInner) GetEventAmountOk() (*float32, bool)`

GetEventAmountOk returns a tuple with the EventAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAmount

`func (o *ReconEntityDataInner) SetEventAmount(v float32)`

SetEventAmount sets EventAmount field to given value.

### HasEventAmount

`func (o *ReconEntityDataInner) HasEventAmount() bool`

HasEventAmount returns a boolean if a field has been set.

### GetSaleType

`func (o *ReconEntityDataInner) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *ReconEntityDataInner) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *ReconEntityDataInner) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *ReconEntityDataInner) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetEventStatus

`func (o *ReconEntityDataInner) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *ReconEntityDataInner) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *ReconEntityDataInner) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *ReconEntityDataInner) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEntity

`func (o *ReconEntityDataInner) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ReconEntityDataInner) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ReconEntityDataInner) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ReconEntityDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventTime

`func (o *ReconEntityDataInner) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ReconEntityDataInner) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ReconEntityDataInner) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ReconEntityDataInner) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventCurrency

`func (o *ReconEntityDataInner) GetEventCurrency() string`

GetEventCurrency returns the EventCurrency field if non-nil, zero value otherwise.

### GetEventCurrencyOk

`func (o *ReconEntityDataInner) GetEventCurrencyOk() (*string, bool)`

GetEventCurrencyOk returns a tuple with the EventCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCurrency

`func (o *ReconEntityDataInner) SetEventCurrency(v string)`

SetEventCurrency sets EventCurrency field to given value.

### HasEventCurrency

`func (o *ReconEntityDataInner) HasEventCurrency() bool`

HasEventCurrency returns a boolean if a field has been set.

### GetOrderId

`func (o *ReconEntityDataInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ReconEntityDataInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ReconEntityDataInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ReconEntityDataInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *ReconEntityDataInner) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *ReconEntityDataInner) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *ReconEntityDataInner) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *ReconEntityDataInner) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *ReconEntityDataInner) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *ReconEntityDataInner) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *ReconEntityDataInner) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *ReconEntityDataInner) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *ReconEntityDataInner) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *ReconEntityDataInner) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *ReconEntityDataInner) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *ReconEntityDataInner) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *ReconEntityDataInner) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *ReconEntityDataInner) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *ReconEntityDataInner) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *ReconEntityDataInner) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *ReconEntityDataInner) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *ReconEntityDataInner) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *ReconEntityDataInner) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *ReconEntityDataInner) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentUtr

`func (o *ReconEntityDataInner) GetPaymentUtr() string`

GetPaymentUtr returns the PaymentUtr field if non-nil, zero value otherwise.

### GetPaymentUtrOk

`func (o *ReconEntityDataInner) GetPaymentUtrOk() (*string, bool)`

GetPaymentUtrOk returns a tuple with the PaymentUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUtr

`func (o *ReconEntityDataInner) SetPaymentUtr(v string)`

SetPaymentUtr sets PaymentUtr field to given value.

### HasPaymentUtr

`func (o *ReconEntityDataInner) HasPaymentUtr() bool`

HasPaymentUtr returns a boolean if a field has been set.

### GetPaymentTime

`func (o *ReconEntityDataInner) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *ReconEntityDataInner) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *ReconEntityDataInner) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *ReconEntityDataInner) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentServiceCharge

`func (o *ReconEntityDataInner) GetPaymentServiceCharge() float32`

GetPaymentServiceCharge returns the PaymentServiceCharge field if non-nil, zero value otherwise.

### GetPaymentServiceChargeOk

`func (o *ReconEntityDataInner) GetPaymentServiceChargeOk() (*float32, bool)`

GetPaymentServiceChargeOk returns a tuple with the PaymentServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceCharge

`func (o *ReconEntityDataInner) SetPaymentServiceCharge(v float32)`

SetPaymentServiceCharge sets PaymentServiceCharge field to given value.

### HasPaymentServiceCharge

`func (o *ReconEntityDataInner) HasPaymentServiceCharge() bool`

HasPaymentServiceCharge returns a boolean if a field has been set.

### GetPaymentServiceTax

`func (o *ReconEntityDataInner) GetPaymentServiceTax() float32`

GetPaymentServiceTax returns the PaymentServiceTax field if non-nil, zero value otherwise.

### GetPaymentServiceTaxOk

`func (o *ReconEntityDataInner) GetPaymentServiceTaxOk() (*float32, bool)`

GetPaymentServiceTaxOk returns a tuple with the PaymentServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTax

`func (o *ReconEntityDataInner) SetPaymentServiceTax(v float32)`

SetPaymentServiceTax sets PaymentServiceTax field to given value.

### HasPaymentServiceTax

`func (o *ReconEntityDataInner) HasPaymentServiceTax() bool`

HasPaymentServiceTax returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *ReconEntityDataInner) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *ReconEntityDataInner) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *ReconEntityDataInner) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *ReconEntityDataInner) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSettlementId

`func (o *ReconEntityDataInner) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *ReconEntityDataInner) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *ReconEntityDataInner) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *ReconEntityDataInner) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementDate

`func (o *ReconEntityDataInner) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *ReconEntityDataInner) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *ReconEntityDataInner) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *ReconEntityDataInner) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetSettlementUtr

`func (o *ReconEntityDataInner) GetSettlementUtr() string`

GetSettlementUtr returns the SettlementUtr field if non-nil, zero value otherwise.

### GetSettlementUtrOk

`func (o *ReconEntityDataInner) GetSettlementUtrOk() (*string, bool)`

GetSettlementUtrOk returns a tuple with the SettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtr

`func (o *ReconEntityDataInner) SetSettlementUtr(v string)`

SetSettlementUtr sets SettlementUtr field to given value.

### HasSettlementUtr

`func (o *ReconEntityDataInner) HasSettlementUtr() bool`

HasSettlementUtr returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *ReconEntityDataInner) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *ReconEntityDataInner) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *ReconEntityDataInner) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *ReconEntityDataInner) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *ReconEntityDataInner) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *ReconEntityDataInner) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *ReconEntityDataInner) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *ReconEntityDataInner) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetVendorCommission

`func (o *ReconEntityDataInner) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *ReconEntityDataInner) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *ReconEntityDataInner) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *ReconEntityDataInner) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetClosedInFavorOf

`func (o *ReconEntityDataInner) GetClosedInFavorOf() string`

GetClosedInFavorOf returns the ClosedInFavorOf field if non-nil, zero value otherwise.

### GetClosedInFavorOfOk

`func (o *ReconEntityDataInner) GetClosedInFavorOfOk() (*string, bool)`

GetClosedInFavorOfOk returns a tuple with the ClosedInFavorOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedInFavorOf

`func (o *ReconEntityDataInner) SetClosedInFavorOf(v string)`

SetClosedInFavorOf sets ClosedInFavorOf field to given value.

### HasClosedInFavorOf

`func (o *ReconEntityDataInner) HasClosedInFavorOf() bool`

HasClosedInFavorOf returns a boolean if a field has been set.

### GetDisputeResolvedOn

`func (o *ReconEntityDataInner) GetDisputeResolvedOn() string`

GetDisputeResolvedOn returns the DisputeResolvedOn field if non-nil, zero value otherwise.

### GetDisputeResolvedOnOk

`func (o *ReconEntityDataInner) GetDisputeResolvedOnOk() (*string, bool)`

GetDisputeResolvedOnOk returns a tuple with the DisputeResolvedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeResolvedOn

`func (o *ReconEntityDataInner) SetDisputeResolvedOn(v string)`

SetDisputeResolvedOn sets DisputeResolvedOn field to given value.

### HasDisputeResolvedOn

`func (o *ReconEntityDataInner) HasDisputeResolvedOn() bool`

HasDisputeResolvedOn returns a boolean if a field has been set.

### GetDisputeCategory

`func (o *ReconEntityDataInner) GetDisputeCategory() string`

GetDisputeCategory returns the DisputeCategory field if non-nil, zero value otherwise.

### GetDisputeCategoryOk

`func (o *ReconEntityDataInner) GetDisputeCategoryOk() (*string, bool)`

GetDisputeCategoryOk returns a tuple with the DisputeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeCategory

`func (o *ReconEntityDataInner) SetDisputeCategory(v string)`

SetDisputeCategory sets DisputeCategory field to given value.

### HasDisputeCategory

`func (o *ReconEntityDataInner) HasDisputeCategory() bool`

HasDisputeCategory returns a boolean if a field has been set.

### GetDisputeNote

`func (o *ReconEntityDataInner) GetDisputeNote() string`

GetDisputeNote returns the DisputeNote field if non-nil, zero value otherwise.

### GetDisputeNoteOk

`func (o *ReconEntityDataInner) GetDisputeNoteOk() (*string, bool)`

GetDisputeNoteOk returns a tuple with the DisputeNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeNote

`func (o *ReconEntityDataInner) SetDisputeNote(v string)`

SetDisputeNote sets DisputeNote field to given value.

### HasDisputeNote

`func (o *ReconEntityDataInner) HasDisputeNote() bool`

HasDisputeNote returns a boolean if a field has been set.

### GetRefundProcessedAt

`func (o *ReconEntityDataInner) GetRefundProcessedAt() string`

GetRefundProcessedAt returns the RefundProcessedAt field if non-nil, zero value otherwise.

### GetRefundProcessedAtOk

`func (o *ReconEntityDataInner) GetRefundProcessedAtOk() (*string, bool)`

GetRefundProcessedAtOk returns a tuple with the RefundProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundProcessedAt

`func (o *ReconEntityDataInner) SetRefundProcessedAt(v string)`

SetRefundProcessedAt sets RefundProcessedAt field to given value.

### HasRefundProcessedAt

`func (o *ReconEntityDataInner) HasRefundProcessedAt() bool`

HasRefundProcessedAt returns a boolean if a field has been set.

### GetRefundArn

`func (o *ReconEntityDataInner) GetRefundArn() string`

GetRefundArn returns the RefundArn field if non-nil, zero value otherwise.

### GetRefundArnOk

`func (o *ReconEntityDataInner) GetRefundArnOk() (*string, bool)`

GetRefundArnOk returns a tuple with the RefundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundArn

`func (o *ReconEntityDataInner) SetRefundArn(v string)`

SetRefundArn sets RefundArn field to given value.

### HasRefundArn

`func (o *ReconEntityDataInner) HasRefundArn() bool`

HasRefundArn returns a boolean if a field has been set.

### GetRefundNote

`func (o *ReconEntityDataInner) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *ReconEntityDataInner) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *ReconEntityDataInner) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *ReconEntityDataInner) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundId

`func (o *ReconEntityDataInner) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *ReconEntityDataInner) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *ReconEntityDataInner) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *ReconEntityDataInner) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetAdjustmentRemarks

`func (o *ReconEntityDataInner) GetAdjustmentRemarks() string`

GetAdjustmentRemarks returns the AdjustmentRemarks field if non-nil, zero value otherwise.

### GetAdjustmentRemarksOk

`func (o *ReconEntityDataInner) GetAdjustmentRemarksOk() (*string, bool)`

GetAdjustmentRemarksOk returns a tuple with the AdjustmentRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentRemarks

`func (o *ReconEntityDataInner) SetAdjustmentRemarks(v string)`

SetAdjustmentRemarks sets AdjustmentRemarks field to given value.

### HasAdjustmentRemarks

`func (o *ReconEntityDataInner) HasAdjustmentRemarks() bool`

HasAdjustmentRemarks returns a boolean if a field has been set.

### GetAdjustment

`func (o *ReconEntityDataInner) GetAdjustment() float32`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ReconEntityDataInner) GetAdjustmentOk() (*float32, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ReconEntityDataInner) SetAdjustment(v float32)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *ReconEntityDataInner) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetServiceTax

`func (o *ReconEntityDataInner) GetServiceTax() float32`

GetServiceTax returns the ServiceTax field if non-nil, zero value otherwise.

### GetServiceTaxOk

`func (o *ReconEntityDataInner) GetServiceTaxOk() (*float32, bool)`

GetServiceTaxOk returns a tuple with the ServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTax

`func (o *ReconEntityDataInner) SetServiceTax(v float32)`

SetServiceTax sets ServiceTax field to given value.

### HasServiceTax

`func (o *ReconEntityDataInner) HasServiceTax() bool`

HasServiceTax returns a boolean if a field has been set.

### GetServiceCharge

`func (o *ReconEntityDataInner) GetServiceCharge() float32`

GetServiceCharge returns the ServiceCharge field if non-nil, zero value otherwise.

### GetServiceChargeOk

`func (o *ReconEntityDataInner) GetServiceChargeOk() (*float32, bool)`

GetServiceChargeOk returns a tuple with the ServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCharge

`func (o *ReconEntityDataInner) SetServiceCharge(v float32)`

SetServiceCharge sets ServiceCharge field to given value.

### HasServiceCharge

`func (o *ReconEntityDataInner) HasServiceCharge() bool`

HasServiceCharge returns a boolean if a field has been set.

### GetAmountSettled

`func (o *ReconEntityDataInner) GetAmountSettled() float32`

GetAmountSettled returns the AmountSettled field if non-nil, zero value otherwise.

### GetAmountSettledOk

`func (o *ReconEntityDataInner) GetAmountSettledOk() (*float32, bool)`

GetAmountSettledOk returns a tuple with the AmountSettled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSettled

`func (o *ReconEntityDataInner) SetAmountSettled(v float32)`

SetAmountSettled sets AmountSettled field to given value.

### HasAmountSettled

`func (o *ReconEntityDataInner) HasAmountSettled() bool`

HasAmountSettled returns a boolean if a field has been set.

### GetPaymentFrom

`func (o *ReconEntityDataInner) GetPaymentFrom() string`

GetPaymentFrom returns the PaymentFrom field if non-nil, zero value otherwise.

### GetPaymentFromOk

`func (o *ReconEntityDataInner) GetPaymentFromOk() (*string, bool)`

GetPaymentFromOk returns a tuple with the PaymentFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentFrom

`func (o *ReconEntityDataInner) SetPaymentFrom(v string)`

SetPaymentFrom sets PaymentFrom field to given value.

### HasPaymentFrom

`func (o *ReconEntityDataInner) HasPaymentFrom() bool`

HasPaymentFrom returns a boolean if a field has been set.

### GetPaymentTill

`func (o *ReconEntityDataInner) GetPaymentTill() string`

GetPaymentTill returns the PaymentTill field if non-nil, zero value otherwise.

### GetPaymentTillOk

`func (o *ReconEntityDataInner) GetPaymentTillOk() (*string, bool)`

GetPaymentTillOk returns a tuple with the PaymentTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTill

`func (o *ReconEntityDataInner) SetPaymentTill(v string)`

SetPaymentTill sets PaymentTill field to given value.

### HasPaymentTill

`func (o *ReconEntityDataInner) HasPaymentTill() bool`

HasPaymentTill returns a boolean if a field has been set.

### GetReason

`func (o *ReconEntityDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReconEntityDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReconEntityDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReconEntityDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *ReconEntityDataInner) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *ReconEntityDataInner) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *ReconEntityDataInner) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *ReconEntityDataInner) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementType

`func (o *ReconEntityDataInner) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *ReconEntityDataInner) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *ReconEntityDataInner) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *ReconEntityDataInner) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetSettlementCharge

`func (o *ReconEntityDataInner) GetSettlementCharge() float32`

GetSettlementCharge returns the SettlementCharge field if non-nil, zero value otherwise.

### GetSettlementChargeOk

`func (o *ReconEntityDataInner) GetSettlementChargeOk() (*float32, bool)`

GetSettlementChargeOk returns a tuple with the SettlementCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCharge

`func (o *ReconEntityDataInner) SetSettlementCharge(v float32)`

SetSettlementCharge sets SettlementCharge field to given value.

### HasSettlementCharge

`func (o *ReconEntityDataInner) HasSettlementCharge() bool`

HasSettlementCharge returns a boolean if a field has been set.

### GetSettlementTax

`func (o *ReconEntityDataInner) GetSettlementTax() float32`

GetSettlementTax returns the SettlementTax field if non-nil, zero value otherwise.

### GetSettlementTaxOk

`func (o *ReconEntityDataInner) GetSettlementTaxOk() (*float32, bool)`

GetSettlementTaxOk returns a tuple with the SettlementTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTax

`func (o *ReconEntityDataInner) SetSettlementTax(v float32)`

SetSettlementTax sets SettlementTax field to given value.

### HasSettlementTax

`func (o *ReconEntityDataInner) HasSettlementTax() bool`

HasSettlementTax returns a boolean if a field has been set.

### GetRemarks

`func (o *ReconEntityDataInner) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ReconEntityDataInner) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ReconEntityDataInner) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ReconEntityDataInner) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


