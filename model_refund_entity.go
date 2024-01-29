/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the RefundEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundEntity{}

// RefundEntity The refund entity
type RefundEntity struct {
	// Cashfree Payments ID of the payment for which refund is initiated
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	// Cashfree Payments ID for a refund
	CfRefundId *string `json:"cf_refund_id,omitempty"`
	// Merchant’s order Id of the order for which refund is initiated
	OrderId *string `json:"order_id,omitempty"`
	// Merchant’s refund ID of the refund
	RefundId *string `json:"refund_id,omitempty"`
	// Type of object
	Entity *string `json:"entity,omitempty"`
	// Amount that is refunded
	RefundAmount *float32 `json:"refund_amount,omitempty"`
	// Currency of the refund amount
	RefundCurrency *string `json:"refund_currency,omitempty"`
	// Note added by merchant for the refund
	RefundNote *string `json:"refund_note,omitempty"`
	// This can be one of [\"SUCCESS\", \"PENDING\", \"CANCELLED\", \"ONHOLD\", \"FAILED\"]
	RefundStatus *string `json:"refund_status,omitempty"`
	// The bank reference number for refund
	RefundArn *string `json:"refund_arn,omitempty"`
	// Charges in INR for processing refund
	RefundCharge *float32 `json:"refund_charge,omitempty"`
	// Description of refund status
	StatusDescription *string `json:"status_description,omitempty"`
	// Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	RefundSplits []VendorSplit `json:"refund_splits,omitempty"`
	// This can be one of [\"PAYMENT_AUTO_REFUND\", \"MERCHANT_INITIATED\", \"UNRECONCILED_AUTO_REFUND\"]
	RefundType *string `json:"refund_type,omitempty"`
	// Method or speed of processing refund
	RefundMode *string `json:"refund_mode,omitempty"`
	// Time of refund creation
	CreatedAt *string `json:"created_at,omitempty"`
	// Time when refund was processed successfully
	ProcessedAt *string `json:"processed_at,omitempty"`
	RefundSpeed *RefundSpeed `json:"refund_speed,omitempty"`
}


func (o RefundEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.CfRefundId) {
		toSerialize["cf_refund_id"] = o.CfRefundId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.RefundId) {
		toSerialize["refund_id"] = o.RefundId
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.RefundCurrency) {
		toSerialize["refund_currency"] = o.RefundCurrency
	}
	if !IsNil(o.RefundNote) {
		toSerialize["refund_note"] = o.RefundNote
	}
	if !IsNil(o.RefundStatus) {
		toSerialize["refund_status"] = o.RefundStatus
	}
	if !IsNil(o.RefundArn) {
		toSerialize["refund_arn"] = o.RefundArn
	}
	if !IsNil(o.RefundCharge) {
		toSerialize["refund_charge"] = o.RefundCharge
	}
	if !IsNil(o.StatusDescription) {
		toSerialize["status_description"] = o.StatusDescription
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RefundSplits) {
		toSerialize["refund_splits"] = o.RefundSplits
	}
	if !IsNil(o.RefundType) {
		toSerialize["refund_type"] = o.RefundType
	}
	if !IsNil(o.RefundMode) {
		toSerialize["refund_mode"] = o.RefundMode
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ProcessedAt) {
		toSerialize["processed_at"] = o.ProcessedAt
	}
	if !IsNil(o.RefundSpeed) {
		toSerialize["refund_speed"] = o.RefundSpeed
	}
	return toSerialize, nil
}



