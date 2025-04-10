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

// checks if the CreateSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionRequest{}

// CreateSubscriptionRequest Request body to create a new subscription.
type CreateSubscriptionRequest struct {
	// A unique ID for the subscription. It can include alphanumeric characters, underscore, dot, hyphen, and space. Maximum characters allowed is 250.
	SubscriptionId string `json:"subscription_id"`
	CustomerDetails SubscriptionCustomerDetails `json:"customer_details"`
	PlanDetails CreateSubscriptionRequestPlanDetails `json:"plan_details"`
	AuthorizationDetails *CreateSubscriptionRequestAuthorizationDetails `json:"authorization_details,omitempty"`
	SubscriptionMeta *CreateSubscriptionRequestSubscriptionMeta `json:"subscription_meta,omitempty"`
	// Expiry date for the subscription.
	SubscriptionExpiryTime *string `json:"subscription_expiry_time,omitempty"`
	// Time at which the first charge will be made for the subscription after authorization. Applicable only for PERIODIC plans.
	SubscriptionFirstChargeTime *string `json:"subscription_first_charge_time,omitempty"`
	// Note for the subscription.
	SubscriptionNote *string `json:"subscription_note,omitempty"`
	// Tags for the subscription.
	SubscriptionTags map[string]interface{} `json:"subscription_tags,omitempty"`
	// Payment splits for the subscription.
	SubscriptionPaymentSplits []SubscriptionPaymentSplitItem `json:"subscription_payment_splits,omitempty"`
}


func (o CreateSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["customer_details"] = o.CustomerDetails
	toSerialize["plan_details"] = o.PlanDetails
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorization_details"] = o.AuthorizationDetails
	}
	if !IsNil(o.SubscriptionMeta) {
		toSerialize["subscription_meta"] = o.SubscriptionMeta
	}
	if !IsNil(o.SubscriptionExpiryTime) {
		toSerialize["subscription_expiry_time"] = o.SubscriptionExpiryTime
	}
	if !IsNil(o.SubscriptionFirstChargeTime) {
		toSerialize["subscription_first_charge_time"] = o.SubscriptionFirstChargeTime
	}
	if !IsNil(o.SubscriptionNote) {
		toSerialize["subscription_note"] = o.SubscriptionNote
	}
	if !IsNil(o.SubscriptionTags) {
		toSerialize["subscription_tags"] = o.SubscriptionTags
	}
	if !IsNil(o.SubscriptionPaymentSplits) {
		toSerialize["subscription_payment_splits"] = o.SubscriptionPaymentSplits
	}
	return toSerialize, nil
}



