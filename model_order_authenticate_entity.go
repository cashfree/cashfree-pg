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

// checks if the OrderAuthenticateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAuthenticateEntity{}

// OrderAuthenticateEntity This is the response shared when merchant inovkes the OTP submit or resend API
type OrderAuthenticateEntity struct {
	// The payment id for which this request was sent
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	// The action that was invoked for this request.
	Action *string `json:"action,omitempty"`
	// Status of the is action. Will be either failed or successful. If the action is successful, you should still call the authorization status to verify the final payment status.
	AuthenticateStatus *string `json:"authenticate_status,omitempty"`
	// Human readable message which describes the status in more detail
	PaymentMessage *string `json:"payment_message,omitempty"`
}


func (o OrderAuthenticateEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAuthenticateEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AuthenticateStatus) {
		toSerialize["authenticate_status"] = o.AuthenticateStatus
	}
	if !IsNil(o.PaymentMessage) {
		toSerialize["payment_message"] = o.PaymentMessage
	}
	return toSerialize, nil
}



