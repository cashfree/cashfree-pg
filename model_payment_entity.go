/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2025-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the PaymentEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentEntity{}

// PaymentEntity payment entity full object
type PaymentEntity struct {
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	Entity *string `json:"entity,omitempty"`
	ErrorDetails *ErrorDetailsInPaymentsEntity `json:"error_details,omitempty"`
	IsCaptured *bool `json:"is_captured,omitempty"`
	// Order amount can be different from payment amount if you collect service fee from the customer
	OrderAmount *float32 `json:"order_amount,omitempty"`
	// Type of payment group. One of ['prepaid_card', 'upi_ppi_offline', 'cash', 'upi_credit_card', 'paypal', 'net_banking', 'cardless_emi', 'credit_card', 'bank_transfer', 'pay_later', 'debit_card_emi', 'debit_card', 'wallet', 'upi_ppi', 'upi', 'credit_card_emi']
	PaymentGroup *string `json:"payment_group,omitempty"`
	PaymentCurrency *string `json:"payment_currency,omitempty"`
	PaymentAmount *float32 `json:"payment_amount,omitempty"`
	// This is the time when the payment was initiated
	PaymentTime *string `json:"payment_time,omitempty"`
	// This is the time when the payment reaches its terminal state
	PaymentCompletionTime *string `json:"payment_completion_time,omitempty"`
	// The transaction status can be one of  [\"SUCCESS\", \"NOT_ATTEMPTED\", \"FAILED\", \"USER_DROPPED\", \"VOID\", \"CANCELLED\", \"PENDING\"]
	PaymentStatus *string `json:"payment_status,omitempty"`
	PaymentMessage *string `json:"payment_message,omitempty"`
	BankReference *string `json:"bank_reference,omitempty"`
	AuthId *string `json:"auth_id,omitempty"`
	OrderCurrency *string `json:"order_currency,omitempty"`
	Authorization *AuthorizationInPaymentsEntity `json:"authorization,omitempty"`
	PaymentMethod *PaymentEntityPaymentMethod `json:"payment_method,omitempty"`
	InternationalPayment *InternationalPaymentEntity `json:"international_payment,omitempty"`
	PaymentGatewayDetails *PaymentGatewayDetails `json:"payment_gateway_details,omitempty"`
	PaymentSurcharge *PaymentEntityPaymentSurcharge `json:"payment_surcharge,omitempty"`
}


func (o PaymentEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["error_details"] = o.ErrorDetails
	}
	if !IsNil(o.IsCaptured) {
		toSerialize["is_captured"] = o.IsCaptured
	}
	if !IsNil(o.OrderAmount) {
		toSerialize["order_amount"] = o.OrderAmount
	}
	if !IsNil(o.PaymentGroup) {
		toSerialize["payment_group"] = o.PaymentGroup
	}
	if !IsNil(o.PaymentCurrency) {
		toSerialize["payment_currency"] = o.PaymentCurrency
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentTime) {
		toSerialize["payment_time"] = o.PaymentTime
	}
	if !IsNil(o.PaymentCompletionTime) {
		toSerialize["payment_completion_time"] = o.PaymentCompletionTime
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if !IsNil(o.PaymentMessage) {
		toSerialize["payment_message"] = o.PaymentMessage
	}
	if !IsNil(o.BankReference) {
		toSerialize["bank_reference"] = o.BankReference
	}
	if !IsNil(o.AuthId) {
		toSerialize["auth_id"] = o.AuthId
	}
	if !IsNil(o.OrderCurrency) {
		toSerialize["order_currency"] = o.OrderCurrency
	}
	if !IsNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.InternationalPayment) {
		toSerialize["international_payment"] = o.InternationalPayment
	}
	if !IsNil(o.PaymentGatewayDetails) {
		toSerialize["payment_gateway_details"] = o.PaymentGatewayDetails
	}
	if !IsNil(o.PaymentSurcharge) {
		toSerialize["payment_surcharge"] = o.PaymentSurcharge
	}
	return toSerialize, nil
}



