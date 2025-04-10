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

// checks if the ManageSubscriptionPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageSubscriptionPaymentRequest{}

// ManageSubscriptionPaymentRequest Request body to manage a subscription payment.
type ManageSubscriptionPaymentRequest struct {
	// The unique ID which was used to create subscription.
	SubscriptionId string `json:"subscription_id"`
	// The unique ID which was used to create payment.
	PaymentId string `json:"payment_id"`
	// Action to be performed on the payment. Possible values - CANCEL, RETRY.
	Action string `json:"action"`
	ActionDetails *ManageSubscriptionPaymentRequestActionDetails `json:"action_details,omitempty"`
}


func (o ManageSubscriptionPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageSubscriptionPaymentRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["action"] = o.Action
	if !IsNil(o.ActionDetails) {
		toSerialize["action_details"] = o.ActionDetails
	}
	return toSerialize, nil
}



