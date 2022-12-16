/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-01-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"encoding/json"
)

// CFLinkOrders struct for CFLinkOrders
type CFLinkOrders struct {
	CfOrderId *int32 `json:"cf_order_id,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	Entity *string `json:"entity,omitempty"`
	OrderCurrency *string `json:"order_currency,omitempty"`
	OrderAmount *float32 `json:"order_amount,omitempty"`
	OrderStatus *string `json:"order_status,omitempty"`
	OrderToken *string `json:"order_token,omitempty"`
	OrderExpiryTime *string `json:"order_expiry_time,omitempty"`
	OrderNote *string `json:"order_note,omitempty"`
	CustomerDetails *CFLinkCustomerDetailsEntity `json:"customer_details,omitempty"`
	Payments *CFPaymentURLObject `json:"payments,omitempty"`
	Settlements *CFSettlementURLObject `json:"settlements,omitempty"`
	Refunds *CFRefundURLObject `json:"refunds,omitempty"`
}

// NewCFLinkOrders instantiates a new CFLinkOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFLinkOrders() *CFLinkOrders {
	this := CFLinkOrders{}
	return &this
}

// NewCFLinkOrdersWithDefaults instantiates a new CFLinkOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFLinkOrdersWithDefaults() *CFLinkOrders {
	this := CFLinkOrders{}
	return &this
}

// GetCfOrderId returns the CfOrderId field value if set, zero value otherwise.
func (o *CFLinkOrders) GetCfOrderId() int32 {
	if o == nil || o.CfOrderId == nil {
		var ret int32
		return ret
	}
	return *o.CfOrderId
}

// GetCfOrderIdOk returns a tuple with the CfOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetCfOrderIdOk() (*int32, bool) {
	if o == nil || o.CfOrderId == nil {
		return nil, false
	}
	return o.CfOrderId, true
}

// HasCfOrderId returns a boolean if a field has been set.
func (o *CFLinkOrders) HasCfOrderId() bool {
	if o != nil && o.CfOrderId != nil {
		return true
	}

	return false
}

// SetCfOrderId gets a reference to the given int32 and assigns it to the CfOrderId field.
func (o *CFLinkOrders) SetCfOrderId(v int32) {
	o.CfOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CFLinkOrders) SetOrderId(v string) {
	o.OrderId = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *CFLinkOrders) GetEntity() string {
	if o == nil || o.Entity == nil {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetEntityOk() (*string, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *CFLinkOrders) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *CFLinkOrders) SetEntity(v string) {
	o.Entity = &v
}

// GetOrderCurrency returns the OrderCurrency field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderCurrency() string {
	if o == nil || o.OrderCurrency == nil {
		var ret string
		return ret
	}
	return *o.OrderCurrency
}

// GetOrderCurrencyOk returns a tuple with the OrderCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderCurrencyOk() (*string, bool) {
	if o == nil || o.OrderCurrency == nil {
		return nil, false
	}
	return o.OrderCurrency, true
}

// HasOrderCurrency returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderCurrency() bool {
	if o != nil && o.OrderCurrency != nil {
		return true
	}

	return false
}

// SetOrderCurrency gets a reference to the given string and assigns it to the OrderCurrency field.
func (o *CFLinkOrders) SetOrderCurrency(v string) {
	o.OrderCurrency = &v
}

// GetOrderAmount returns the OrderAmount field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderAmount() float32 {
	if o == nil || o.OrderAmount == nil {
		var ret float32
		return ret
	}
	return *o.OrderAmount
}

// GetOrderAmountOk returns a tuple with the OrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderAmountOk() (*float32, bool) {
	if o == nil || o.OrderAmount == nil {
		return nil, false
	}
	return o.OrderAmount, true
}

// HasOrderAmount returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderAmount() bool {
	if o != nil && o.OrderAmount != nil {
		return true
	}

	return false
}

// SetOrderAmount gets a reference to the given float32 and assigns it to the OrderAmount field.
func (o *CFLinkOrders) SetOrderAmount(v float32) {
	o.OrderAmount = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderStatus() string {
	if o == nil || o.OrderStatus == nil {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderStatusOk() (*string, bool) {
	if o == nil || o.OrderStatus == nil {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderStatus() bool {
	if o != nil && o.OrderStatus != nil {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *CFLinkOrders) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOrderToken returns the OrderToken field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderToken() string {
	if o == nil || o.OrderToken == nil {
		var ret string
		return ret
	}
	return *o.OrderToken
}

// GetOrderTokenOk returns a tuple with the OrderToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderTokenOk() (*string, bool) {
	if o == nil || o.OrderToken == nil {
		return nil, false
	}
	return o.OrderToken, true
}

// HasOrderToken returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderToken() bool {
	if o != nil && o.OrderToken != nil {
		return true
	}

	return false
}

// SetOrderToken gets a reference to the given string and assigns it to the OrderToken field.
func (o *CFLinkOrders) SetOrderToken(v string) {
	o.OrderToken = &v
}

// GetOrderExpiryTime returns the OrderExpiryTime field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderExpiryTime() string {
	if o == nil || o.OrderExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.OrderExpiryTime
}

// GetOrderExpiryTimeOk returns a tuple with the OrderExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderExpiryTimeOk() (*string, bool) {
	if o == nil || o.OrderExpiryTime == nil {
		return nil, false
	}
	return o.OrderExpiryTime, true
}

// HasOrderExpiryTime returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderExpiryTime() bool {
	if o != nil && o.OrderExpiryTime != nil {
		return true
	}

	return false
}

// SetOrderExpiryTime gets a reference to the given string and assigns it to the OrderExpiryTime field.
func (o *CFLinkOrders) SetOrderExpiryTime(v string) {
	o.OrderExpiryTime = &v
}

// GetOrderNote returns the OrderNote field value if set, zero value otherwise.
func (o *CFLinkOrders) GetOrderNote() string {
	if o == nil || o.OrderNote == nil {
		var ret string
		return ret
	}
	return *o.OrderNote
}

// GetOrderNoteOk returns a tuple with the OrderNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetOrderNoteOk() (*string, bool) {
	if o == nil || o.OrderNote == nil {
		return nil, false
	}
	return o.OrderNote, true
}

// HasOrderNote returns a boolean if a field has been set.
func (o *CFLinkOrders) HasOrderNote() bool {
	if o != nil && o.OrderNote != nil {
		return true
	}

	return false
}

// SetOrderNote gets a reference to the given string and assigns it to the OrderNote field.
func (o *CFLinkOrders) SetOrderNote(v string) {
	o.OrderNote = &v
}

// GetCustomerDetails returns the CustomerDetails field value if set, zero value otherwise.
func (o *CFLinkOrders) GetCustomerDetails() CFLinkCustomerDetailsEntity {
	if o == nil || o.CustomerDetails == nil {
		var ret CFLinkCustomerDetailsEntity
		return ret
	}
	return *o.CustomerDetails
}

// GetCustomerDetailsOk returns a tuple with the CustomerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetCustomerDetailsOk() (*CFLinkCustomerDetailsEntity, bool) {
	if o == nil || o.CustomerDetails == nil {
		return nil, false
	}
	return o.CustomerDetails, true
}

// HasCustomerDetails returns a boolean if a field has been set.
func (o *CFLinkOrders) HasCustomerDetails() bool {
	if o != nil && o.CustomerDetails != nil {
		return true
	}

	return false
}

// SetCustomerDetails gets a reference to the given CFLinkCustomerDetailsEntity and assigns it to the CustomerDetails field.
func (o *CFLinkOrders) SetCustomerDetails(v CFLinkCustomerDetailsEntity) {
	o.CustomerDetails = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *CFLinkOrders) GetPayments() CFPaymentURLObject {
	if o == nil || o.Payments == nil {
		var ret CFPaymentURLObject
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetPaymentsOk() (*CFPaymentURLObject, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *CFLinkOrders) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given CFPaymentURLObject and assigns it to the Payments field.
func (o *CFLinkOrders) SetPayments(v CFPaymentURLObject) {
	o.Payments = &v
}

// GetSettlements returns the Settlements field value if set, zero value otherwise.
func (o *CFLinkOrders) GetSettlements() CFSettlementURLObject {
	if o == nil || o.Settlements == nil {
		var ret CFSettlementURLObject
		return ret
	}
	return *o.Settlements
}

// GetSettlementsOk returns a tuple with the Settlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetSettlementsOk() (*CFSettlementURLObject, bool) {
	if o == nil || o.Settlements == nil {
		return nil, false
	}
	return o.Settlements, true
}

// HasSettlements returns a boolean if a field has been set.
func (o *CFLinkOrders) HasSettlements() bool {
	if o != nil && o.Settlements != nil {
		return true
	}

	return false
}

// SetSettlements gets a reference to the given CFSettlementURLObject and assigns it to the Settlements field.
func (o *CFLinkOrders) SetSettlements(v CFSettlementURLObject) {
	o.Settlements = &v
}

// GetRefunds returns the Refunds field value if set, zero value otherwise.
func (o *CFLinkOrders) GetRefunds() CFRefundURLObject {
	if o == nil || o.Refunds == nil {
		var ret CFRefundURLObject
		return ret
	}
	return *o.Refunds
}

// GetRefundsOk returns a tuple with the Refunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFLinkOrders) GetRefundsOk() (*CFRefundURLObject, bool) {
	if o == nil || o.Refunds == nil {
		return nil, false
	}
	return o.Refunds, true
}

// HasRefunds returns a boolean if a field has been set.
func (o *CFLinkOrders) HasRefunds() bool {
	if o != nil && o.Refunds != nil {
		return true
	}

	return false
}

// SetRefunds gets a reference to the given CFRefundURLObject and assigns it to the Refunds field.
func (o *CFLinkOrders) SetRefunds(v CFRefundURLObject) {
	o.Refunds = &v
}

func (o CFLinkOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CfOrderId != nil {
		toSerialize["cf_order_id"] = o.CfOrderId
	}
	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.OrderCurrency != nil {
		toSerialize["order_currency"] = o.OrderCurrency
	}
	if o.OrderAmount != nil {
		toSerialize["order_amount"] = o.OrderAmount
	}
	if o.OrderStatus != nil {
		toSerialize["order_status"] = o.OrderStatus
	}
	if o.OrderToken != nil {
		toSerialize["order_token"] = o.OrderToken
	}
	if o.OrderExpiryTime != nil {
		toSerialize["order_expiry_time"] = o.OrderExpiryTime
	}
	if o.OrderNote != nil {
		toSerialize["order_note"] = o.OrderNote
	}
	if o.CustomerDetails != nil {
		toSerialize["customer_details"] = o.CustomerDetails
	}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	if o.Settlements != nil {
		toSerialize["settlements"] = o.Settlements
	}
	if o.Refunds != nil {
		toSerialize["refunds"] = o.Refunds
	}
	return json.Marshal(toSerialize)
}

type NullableCFLinkOrders struct {
	value *CFLinkOrders
	isSet bool
}

func (v NullableCFLinkOrders) Get() *CFLinkOrders {
	return v.value
}

func (v *NullableCFLinkOrders) Set(val *CFLinkOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableCFLinkOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableCFLinkOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFLinkOrders(val *CFLinkOrders) *NullableCFLinkOrders {
	return &NullableCFLinkOrders{value: val, isSet: true}
}

func (v NullableCFLinkOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFLinkOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


