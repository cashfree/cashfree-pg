# SettlementReconEntityDataInner

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
**CfPaymentId** | Pointer to **int32** | Cashfree Payments unique ID to identify a payment. | [optional] 
**CfSettlementId** | Pointer to **int32** | Unique ID to identify the settlement. | [optional] 
**SettlementDate** | Pointer to **string** | Date and time when the settlement was processed. | [optional] 
**SettlementUtr** | Pointer to **string** | Unique transaction reference number of the settlement. | [optional] 
**SplitServiceCharge** | Pointer to **float32** | Service charge that is applicable for splitting the payment. | [optional] 
**SplitServiceTax** | Pointer to **float32** | Service tax applicable for splitting the amount to vendors. | [optional] 
**VendorCommission** | Pointer to **float32** | Vendor commission applicable for this transaction. | [optional] 
**ClosedInFavorOf** | Pointer to **string** | Specifies whether the dispute was closed in favor of the merchant or customer. Possible values - Merchant, Customer. | [optional] 
**DisputeResolvedOn** | Pointer to **string** | Date and time when the dispute was resolved. | [optional] 
**DisputeCategory** | Pointer to **string** | Category of the dispute - Dispute code and the reason for dispute is shown. | [optional] 
**DisputeNote** | Pointer to **string** | Note regarding the dispute. | [optional] 
**RefundProcessedAt** | Pointer to **string** | Date and time when the refund was processed. | [optional] 
**RefundArn** | Pointer to **string** | The bank reference number for refund. | [optional] 
**RefundNote** | Pointer to **string** | A refund note for your reference. | [optional] 
**RefundId** | Pointer to **string** | An unique ID associated with the refund. | [optional] 
**AdjustmentRemarks** | Pointer to **string** | Other adjustment remarks. | [optional] 

## Methods

### NewSettlementReconEntityDataInner

`func NewSettlementReconEntityDataInner() *SettlementReconEntityDataInner`

NewSettlementReconEntityDataInner instantiates a new SettlementReconEntityDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerWithDefaults

`func NewSettlementReconEntityDataInnerWithDefaults() *SettlementReconEntityDataInner`

NewSettlementReconEntityDataInnerWithDefaults instantiates a new SettlementReconEntityDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SettlementReconEntityDataInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SettlementReconEntityDataInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SettlementReconEntityDataInner) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SettlementReconEntityDataInner) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *SettlementReconEntityDataInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SettlementReconEntityDataInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SettlementReconEntityDataInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SettlementReconEntityDataInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventSettlementAmount

`func (o *SettlementReconEntityDataInner) GetEventSettlementAmount() float32`

GetEventSettlementAmount returns the EventSettlementAmount field if non-nil, zero value otherwise.

### GetEventSettlementAmountOk

`func (o *SettlementReconEntityDataInner) GetEventSettlementAmountOk() (*float32, bool)`

GetEventSettlementAmountOk returns a tuple with the EventSettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSettlementAmount

`func (o *SettlementReconEntityDataInner) SetEventSettlementAmount(v float32)`

SetEventSettlementAmount sets EventSettlementAmount field to given value.

### HasEventSettlementAmount

`func (o *SettlementReconEntityDataInner) HasEventSettlementAmount() bool`

HasEventSettlementAmount returns a boolean if a field has been set.

### GetEventAmount

`func (o *SettlementReconEntityDataInner) GetEventAmount() float32`

GetEventAmount returns the EventAmount field if non-nil, zero value otherwise.

### GetEventAmountOk

`func (o *SettlementReconEntityDataInner) GetEventAmountOk() (*float32, bool)`

GetEventAmountOk returns a tuple with the EventAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventAmount

`func (o *SettlementReconEntityDataInner) SetEventAmount(v float32)`

SetEventAmount sets EventAmount field to given value.

### HasEventAmount

`func (o *SettlementReconEntityDataInner) HasEventAmount() bool`

HasEventAmount returns a boolean if a field has been set.

### GetSaleType

`func (o *SettlementReconEntityDataInner) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *SettlementReconEntityDataInner) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *SettlementReconEntityDataInner) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *SettlementReconEntityDataInner) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetEventStatus

`func (o *SettlementReconEntityDataInner) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *SettlementReconEntityDataInner) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *SettlementReconEntityDataInner) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *SettlementReconEntityDataInner) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEntity

`func (o *SettlementReconEntityDataInner) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SettlementReconEntityDataInner) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SettlementReconEntityDataInner) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *SettlementReconEntityDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetEventTime

`func (o *SettlementReconEntityDataInner) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SettlementReconEntityDataInner) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SettlementReconEntityDataInner) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *SettlementReconEntityDataInner) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventCurrency

`func (o *SettlementReconEntityDataInner) GetEventCurrency() string`

GetEventCurrency returns the EventCurrency field if non-nil, zero value otherwise.

### GetEventCurrencyOk

`func (o *SettlementReconEntityDataInner) GetEventCurrencyOk() (*string, bool)`

GetEventCurrencyOk returns a tuple with the EventCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCurrency

`func (o *SettlementReconEntityDataInner) SetEventCurrency(v string)`

SetEventCurrency sets EventCurrency field to given value.

### HasEventCurrency

`func (o *SettlementReconEntityDataInner) HasEventCurrency() bool`

HasEventCurrency returns a boolean if a field has been set.

### GetOrderId

`func (o *SettlementReconEntityDataInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SettlementReconEntityDataInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SettlementReconEntityDataInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SettlementReconEntityDataInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderAmount

`func (o *SettlementReconEntityDataInner) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *SettlementReconEntityDataInner) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *SettlementReconEntityDataInner) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *SettlementReconEntityDataInner) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *SettlementReconEntityDataInner) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *SettlementReconEntityDataInner) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *SettlementReconEntityDataInner) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *SettlementReconEntityDataInner) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *SettlementReconEntityDataInner) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *SettlementReconEntityDataInner) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *SettlementReconEntityDataInner) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *SettlementReconEntityDataInner) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *SettlementReconEntityDataInner) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *SettlementReconEntityDataInner) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *SettlementReconEntityDataInner) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *SettlementReconEntityDataInner) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *SettlementReconEntityDataInner) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *SettlementReconEntityDataInner) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *SettlementReconEntityDataInner) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *SettlementReconEntityDataInner) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentUtr

`func (o *SettlementReconEntityDataInner) GetPaymentUtr() string`

GetPaymentUtr returns the PaymentUtr field if non-nil, zero value otherwise.

### GetPaymentUtrOk

`func (o *SettlementReconEntityDataInner) GetPaymentUtrOk() (*string, bool)`

GetPaymentUtrOk returns a tuple with the PaymentUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUtr

`func (o *SettlementReconEntityDataInner) SetPaymentUtr(v string)`

SetPaymentUtr sets PaymentUtr field to given value.

### HasPaymentUtr

`func (o *SettlementReconEntityDataInner) HasPaymentUtr() bool`

HasPaymentUtr returns a boolean if a field has been set.

### GetPaymentTime

`func (o *SettlementReconEntityDataInner) GetPaymentTime() string`

GetPaymentTime returns the PaymentTime field if non-nil, zero value otherwise.

### GetPaymentTimeOk

`func (o *SettlementReconEntityDataInner) GetPaymentTimeOk() (*string, bool)`

GetPaymentTimeOk returns a tuple with the PaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTime

`func (o *SettlementReconEntityDataInner) SetPaymentTime(v string)`

SetPaymentTime sets PaymentTime field to given value.

### HasPaymentTime

`func (o *SettlementReconEntityDataInner) HasPaymentTime() bool`

HasPaymentTime returns a boolean if a field has been set.

### GetPaymentServiceCharge

`func (o *SettlementReconEntityDataInner) GetPaymentServiceCharge() float32`

GetPaymentServiceCharge returns the PaymentServiceCharge field if non-nil, zero value otherwise.

### GetPaymentServiceChargeOk

`func (o *SettlementReconEntityDataInner) GetPaymentServiceChargeOk() (*float32, bool)`

GetPaymentServiceChargeOk returns a tuple with the PaymentServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceCharge

`func (o *SettlementReconEntityDataInner) SetPaymentServiceCharge(v float32)`

SetPaymentServiceCharge sets PaymentServiceCharge field to given value.

### HasPaymentServiceCharge

`func (o *SettlementReconEntityDataInner) HasPaymentServiceCharge() bool`

HasPaymentServiceCharge returns a boolean if a field has been set.

### GetPaymentServiceTax

`func (o *SettlementReconEntityDataInner) GetPaymentServiceTax() float32`

GetPaymentServiceTax returns the PaymentServiceTax field if non-nil, zero value otherwise.

### GetPaymentServiceTaxOk

`func (o *SettlementReconEntityDataInner) GetPaymentServiceTaxOk() (*float32, bool)`

GetPaymentServiceTaxOk returns a tuple with the PaymentServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTax

`func (o *SettlementReconEntityDataInner) SetPaymentServiceTax(v float32)`

SetPaymentServiceTax sets PaymentServiceTax field to given value.

### HasPaymentServiceTax

`func (o *SettlementReconEntityDataInner) HasPaymentServiceTax() bool`

HasPaymentServiceTax returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *SettlementReconEntityDataInner) GetCfPaymentId() int32`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *SettlementReconEntityDataInner) GetCfPaymentIdOk() (*int32, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *SettlementReconEntityDataInner) SetCfPaymentId(v int32)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *SettlementReconEntityDataInner) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSettlementId

`func (o *SettlementReconEntityDataInner) GetCfSettlementId() int32`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *SettlementReconEntityDataInner) GetCfSettlementIdOk() (*int32, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *SettlementReconEntityDataInner) SetCfSettlementId(v int32)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *SettlementReconEntityDataInner) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetSettlementDate

`func (o *SettlementReconEntityDataInner) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *SettlementReconEntityDataInner) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *SettlementReconEntityDataInner) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *SettlementReconEntityDataInner) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetSettlementUtr

`func (o *SettlementReconEntityDataInner) GetSettlementUtr() string`

GetSettlementUtr returns the SettlementUtr field if non-nil, zero value otherwise.

### GetSettlementUtrOk

`func (o *SettlementReconEntityDataInner) GetSettlementUtrOk() (*string, bool)`

GetSettlementUtrOk returns a tuple with the SettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtr

`func (o *SettlementReconEntityDataInner) SetSettlementUtr(v string)`

SetSettlementUtr sets SettlementUtr field to given value.

### HasSettlementUtr

`func (o *SettlementReconEntityDataInner) HasSettlementUtr() bool`

HasSettlementUtr returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *SettlementReconEntityDataInner) GetSplitServiceCharge() float32`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *SettlementReconEntityDataInner) GetSplitServiceChargeOk() (*float32, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *SettlementReconEntityDataInner) SetSplitServiceCharge(v float32)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *SettlementReconEntityDataInner) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *SettlementReconEntityDataInner) GetSplitServiceTax() float32`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *SettlementReconEntityDataInner) GetSplitServiceTaxOk() (*float32, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *SettlementReconEntityDataInner) SetSplitServiceTax(v float32)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *SettlementReconEntityDataInner) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetVendorCommission

`func (o *SettlementReconEntityDataInner) GetVendorCommission() float32`

GetVendorCommission returns the VendorCommission field if non-nil, zero value otherwise.

### GetVendorCommissionOk

`func (o *SettlementReconEntityDataInner) GetVendorCommissionOk() (*float32, bool)`

GetVendorCommissionOk returns a tuple with the VendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCommission

`func (o *SettlementReconEntityDataInner) SetVendorCommission(v float32)`

SetVendorCommission sets VendorCommission field to given value.

### HasVendorCommission

`func (o *SettlementReconEntityDataInner) HasVendorCommission() bool`

HasVendorCommission returns a boolean if a field has been set.

### GetClosedInFavorOf

`func (o *SettlementReconEntityDataInner) GetClosedInFavorOf() string`

GetClosedInFavorOf returns the ClosedInFavorOf field if non-nil, zero value otherwise.

### GetClosedInFavorOfOk

`func (o *SettlementReconEntityDataInner) GetClosedInFavorOfOk() (*string, bool)`

GetClosedInFavorOfOk returns a tuple with the ClosedInFavorOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedInFavorOf

`func (o *SettlementReconEntityDataInner) SetClosedInFavorOf(v string)`

SetClosedInFavorOf sets ClosedInFavorOf field to given value.

### HasClosedInFavorOf

`func (o *SettlementReconEntityDataInner) HasClosedInFavorOf() bool`

HasClosedInFavorOf returns a boolean if a field has been set.

### GetDisputeResolvedOn

`func (o *SettlementReconEntityDataInner) GetDisputeResolvedOn() string`

GetDisputeResolvedOn returns the DisputeResolvedOn field if non-nil, zero value otherwise.

### GetDisputeResolvedOnOk

`func (o *SettlementReconEntityDataInner) GetDisputeResolvedOnOk() (*string, bool)`

GetDisputeResolvedOnOk returns a tuple with the DisputeResolvedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeResolvedOn

`func (o *SettlementReconEntityDataInner) SetDisputeResolvedOn(v string)`

SetDisputeResolvedOn sets DisputeResolvedOn field to given value.

### HasDisputeResolvedOn

`func (o *SettlementReconEntityDataInner) HasDisputeResolvedOn() bool`

HasDisputeResolvedOn returns a boolean if a field has been set.

### GetDisputeCategory

`func (o *SettlementReconEntityDataInner) GetDisputeCategory() string`

GetDisputeCategory returns the DisputeCategory field if non-nil, zero value otherwise.

### GetDisputeCategoryOk

`func (o *SettlementReconEntityDataInner) GetDisputeCategoryOk() (*string, bool)`

GetDisputeCategoryOk returns a tuple with the DisputeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeCategory

`func (o *SettlementReconEntityDataInner) SetDisputeCategory(v string)`

SetDisputeCategory sets DisputeCategory field to given value.

### HasDisputeCategory

`func (o *SettlementReconEntityDataInner) HasDisputeCategory() bool`

HasDisputeCategory returns a boolean if a field has been set.

### GetDisputeNote

`func (o *SettlementReconEntityDataInner) GetDisputeNote() string`

GetDisputeNote returns the DisputeNote field if non-nil, zero value otherwise.

### GetDisputeNoteOk

`func (o *SettlementReconEntityDataInner) GetDisputeNoteOk() (*string, bool)`

GetDisputeNoteOk returns a tuple with the DisputeNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeNote

`func (o *SettlementReconEntityDataInner) SetDisputeNote(v string)`

SetDisputeNote sets DisputeNote field to given value.

### HasDisputeNote

`func (o *SettlementReconEntityDataInner) HasDisputeNote() bool`

HasDisputeNote returns a boolean if a field has been set.

### GetRefundProcessedAt

`func (o *SettlementReconEntityDataInner) GetRefundProcessedAt() string`

GetRefundProcessedAt returns the RefundProcessedAt field if non-nil, zero value otherwise.

### GetRefundProcessedAtOk

`func (o *SettlementReconEntityDataInner) GetRefundProcessedAtOk() (*string, bool)`

GetRefundProcessedAtOk returns a tuple with the RefundProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundProcessedAt

`func (o *SettlementReconEntityDataInner) SetRefundProcessedAt(v string)`

SetRefundProcessedAt sets RefundProcessedAt field to given value.

### HasRefundProcessedAt

`func (o *SettlementReconEntityDataInner) HasRefundProcessedAt() bool`

HasRefundProcessedAt returns a boolean if a field has been set.

### GetRefundArn

`func (o *SettlementReconEntityDataInner) GetRefundArn() string`

GetRefundArn returns the RefundArn field if non-nil, zero value otherwise.

### GetRefundArnOk

`func (o *SettlementReconEntityDataInner) GetRefundArnOk() (*string, bool)`

GetRefundArnOk returns a tuple with the RefundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundArn

`func (o *SettlementReconEntityDataInner) SetRefundArn(v string)`

SetRefundArn sets RefundArn field to given value.

### HasRefundArn

`func (o *SettlementReconEntityDataInner) HasRefundArn() bool`

HasRefundArn returns a boolean if a field has been set.

### GetRefundNote

`func (o *SettlementReconEntityDataInner) GetRefundNote() string`

GetRefundNote returns the RefundNote field if non-nil, zero value otherwise.

### GetRefundNoteOk

`func (o *SettlementReconEntityDataInner) GetRefundNoteOk() (*string, bool)`

GetRefundNoteOk returns a tuple with the RefundNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNote

`func (o *SettlementReconEntityDataInner) SetRefundNote(v string)`

SetRefundNote sets RefundNote field to given value.

### HasRefundNote

`func (o *SettlementReconEntityDataInner) HasRefundNote() bool`

HasRefundNote returns a boolean if a field has been set.

### GetRefundId

`func (o *SettlementReconEntityDataInner) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *SettlementReconEntityDataInner) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *SettlementReconEntityDataInner) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *SettlementReconEntityDataInner) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetAdjustmentRemarks

`func (o *SettlementReconEntityDataInner) GetAdjustmentRemarks() string`

GetAdjustmentRemarks returns the AdjustmentRemarks field if non-nil, zero value otherwise.

### GetAdjustmentRemarksOk

`func (o *SettlementReconEntityDataInner) GetAdjustmentRemarksOk() (*string, bool)`

GetAdjustmentRemarksOk returns a tuple with the AdjustmentRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentRemarks

`func (o *SettlementReconEntityDataInner) SetAdjustmentRemarks(v string)`

SetAdjustmentRemarks sets AdjustmentRemarks field to given value.

### HasAdjustmentRemarks

`func (o *SettlementReconEntityDataInner) HasAdjustmentRemarks() bool`

HasAdjustmentRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


