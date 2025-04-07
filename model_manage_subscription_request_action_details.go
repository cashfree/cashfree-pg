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

// checks if the ManageSubscriptionRequestActionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManageSubscriptionRequestActionDetails{}

// ManageSubscriptionRequestActionDetails Details of the action to be performed.
type ManageSubscriptionRequestActionDetails struct {
	// Next scheduled time for the action. Required for ACTIVATE action.
	NextScheduledTime *string `json:"next_scheduled_time,omitempty"`
	// Plan ID to update. Required for CHANGE_PLAN action.
	PlanId *string `json:"plan_id,omitempty"`
}


func (o ManageSubscriptionRequestActionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManageSubscriptionRequestActionDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextScheduledTime) {
		toSerialize["next_scheduled_time"] = o.NextScheduledTime
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	return toSerialize, nil
}



