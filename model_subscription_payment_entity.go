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

// checks if the SubscriptionPaymentEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPaymentEntity{}

// SubscriptionPaymentEntity The response returned in Get, Create or Manage Subscription Payment APIs.
type SubscriptionPaymentEntity struct {
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"`
	// Cashfree subscription payment reference number
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	// Cashfree subscription reference number
	CfSubscriptionId *string `json:"cf_subscription_id,omitempty"`
	// Cashfree subscription payment transaction ID
	CfTxnId *string `json:"cf_txn_id,omitempty"`
	// Cashfree subscription payment order ID
	CfOrderId *string `json:"cf_order_id,omitempty"`
	FailureDetails *SubscriptionPaymentEntityFailureDetails `json:"failure_details,omitempty"`
	// The charge amount of the payment.
	PaymentAmount *float32 `json:"payment_amount,omitempty"`
	// A unique ID passed by merchant for identifying the transaction.
	PaymentId *string `json:"payment_id,omitempty"`
	// The date on which the payment was initiated.
	PaymentInitiatedDate *string `json:"payment_initiated_date,omitempty"`
	// Payment remarks.
	PaymentRemarks *string `json:"payment_remarks,omitempty"`
	// The date on which the payment is scheduled to be processed.
	PaymentScheduleDate *string `json:"payment_schedule_date,omitempty"`
	// Status of the payment.
	PaymentStatus *string `json:"payment_status,omitempty"`
	// Payment type. Can be AUTH or CHARGE.
	PaymentType *string `json:"payment_type,omitempty"`
	// Retry attempts.
	RetryAttempts *int32 `json:"retry_attempts,omitempty"`
	// A unique ID passed by merchant for identifying the subscription.
	SubscriptionId *string `json:"subscription_id,omitempty"`
}


func (o SubscriptionPaymentEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPaymentEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorization_details"] = o.AuthorizationDetails
	}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.CfSubscriptionId) {
		toSerialize["cf_subscription_id"] = o.CfSubscriptionId
	}
	if !IsNil(o.CfTxnId) {
		toSerialize["cf_txn_id"] = o.CfTxnId
	}
	if !IsNil(o.CfOrderId) {
		toSerialize["cf_order_id"] = o.CfOrderId
	}
	if !IsNil(o.FailureDetails) {
		toSerialize["failure_details"] = o.FailureDetails
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !IsNil(o.PaymentInitiatedDate) {
		toSerialize["payment_initiated_date"] = o.PaymentInitiatedDate
	}
	if !IsNil(o.PaymentRemarks) {
		toSerialize["payment_remarks"] = o.PaymentRemarks
	}
	if !IsNil(o.PaymentScheduleDate) {
		toSerialize["payment_schedule_date"] = o.PaymentScheduleDate
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if !IsNil(o.PaymentType) {
		toSerialize["payment_type"] = o.PaymentType
	}
	if !IsNil(o.RetryAttempts) {
		toSerialize["retry_attempts"] = o.RetryAttempts
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	return toSerialize, nil
}



