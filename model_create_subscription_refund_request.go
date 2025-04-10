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

// checks if the CreateSubscriptionRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionRefundRequest{}

// CreateSubscriptionRefundRequest Request body to create a subscription refund.
type CreateSubscriptionRefundRequest struct {
	// Cashfree subscription payment reference number.
	CfPaymentId *float32 `json:"cf_payment_id,omitempty"`
	// A unique ID passed by merchant for identifying the subscription.
	SubscriptionId string `json:"subscription_id"`
	// A unique ID passed by merchant for identifying the transaction.
	PaymentId *string `json:"payment_id,omitempty"`
	// A unique ID passed by merchant for identifying the refund.
	RefundId string `json:"refund_id"`
	// The amount to be refunded. Can be partial or full amount of the payment.
	RefundAmount float32 `json:"refund_amount"`
	// Refund note.
	RefundNote *string `json:"refund_note,omitempty"`
	// Refund speed. Can be INSTANT or STANDARD. UPI supports only STANDARD refunds, Enach and Pnach supports only INSTANT refunds.
	RefundSpeed *string `json:"refund_speed,omitempty"`
}


func (o CreateSubscriptionRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionRefundRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	toSerialize["subscription_id"] = o.SubscriptionId
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	toSerialize["refund_id"] = o.RefundId
	toSerialize["refund_amount"] = o.RefundAmount
	if !IsNil(o.RefundNote) {
		toSerialize["refund_note"] = o.RefundNote
	}
	if !IsNil(o.RefundSpeed) {
		toSerialize["refund_speed"] = o.RefundSpeed
	}
	return toSerialize, nil
}



