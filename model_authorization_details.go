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

// checks if the AuthorizationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDetails{}

// AuthorizationDetails Details of the authorization done for the subscription. Returned in Get subscription and auth payments.
type AuthorizationDetails struct {
	// Authorization amount for the auth payment.
	AuthorizationAmount *float32 `json:"authorization_amount,omitempty"`
	// Indicates whether the authorization amount should be refunded to the customer automatically. Merchants can use this field to specify if the authorized funds should be returned to the customer after authorization of the subscription.
	AuthorizationAmountRefund *bool `json:"authorization_amount_refund,omitempty"`
	// Authorization reference. UMN for UPI, UMRN for EMandate/Physical Mandate and Enrollment ID for cards.
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	// Authorization time.
	AuthorizationTime *string `json:"authorization_time,omitempty"`
	// Status of the authorization.
	AuthorizationStatus *string `json:"authorization_status,omitempty"`
	// A unique ID passed by merchant for identifying the transaction.
	PaymentId *string `json:"payment_id,omitempty"`
	// Payment method used for the authorization.
	PaymentMethod *string `json:"payment_method,omitempty"`
}


func (o AuthorizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationAmount) {
		toSerialize["authorization_amount"] = o.AuthorizationAmount
	}
	if !IsNil(o.AuthorizationAmountRefund) {
		toSerialize["authorization_amount_refund"] = o.AuthorizationAmountRefund
	}
	if !IsNil(o.AuthorizationReference) {
		toSerialize["authorization_reference"] = o.AuthorizationReference
	}
	if !IsNil(o.AuthorizationTime) {
		toSerialize["authorization_time"] = o.AuthorizationTime
	}
	if !IsNil(o.AuthorizationStatus) {
		toSerialize["authorization_status"] = o.AuthorizationStatus
	}
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return toSerialize, nil
}



