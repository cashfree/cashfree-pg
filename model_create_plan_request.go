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

// checks if the CreatePlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePlanRequest{}

// CreatePlanRequest Request body to create a plan.
type CreatePlanRequest struct {
	// Unique ID to identify the plan. Only alpha-numerics, dot, hyphen and underscore allowed.
	PlanId string `json:"plan_id"`
	// Name of the plan.
	PlanName string `json:"plan_name"`
	// Type of the plan. Possible values - PERIODIC, ON_DEMAND.
	PlanType string `json:"plan_type"`
	// Currency of the plan.
	PlanCurrency *string `json:"plan_currency,omitempty"`
	// Recurring amount for the plan. Required for PERIODIC plan_type.
	PlanRecurringAmount *float32 `json:"plan_recurring_amount,omitempty"`
	// Maximum amount for the plan.
	PlanMaxAmount float32 `json:"plan_max_amount"`
	// Maximum number of payment cycles for the plan.
	PlanMaxCycles *int32 `json:"plan_max_cycles,omitempty"`
	// Number of billing cycles between charges. For instance, if set to 2 and the interval type is 'week', the service will be billed every 2 weeks. Similarly, if set to 3 and the interval type is 'month', the service will be billed every 3 months. Required for PERIODIC plan_type.
	PlanIntervals *int32 `json:"plan_intervals,omitempty"`
	// Interval type for the plan. Possible values - DAY, WEEK, MONTH, YEAR.
	PlanIntervalType *string `json:"plan_interval_type,omitempty"`
	// Note for the plan.
	PlanNote *string `json:"plan_note,omitempty"`
}


func (o CreatePlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePlanRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["plan_id"] = o.PlanId
	toSerialize["plan_name"] = o.PlanName
	toSerialize["plan_type"] = o.PlanType
	if !IsNil(o.PlanCurrency) {
		toSerialize["plan_currency"] = o.PlanCurrency
	}
	if !IsNil(o.PlanRecurringAmount) {
		toSerialize["plan_recurring_amount"] = o.PlanRecurringAmount
	}
	toSerialize["plan_max_amount"] = o.PlanMaxAmount
	if !IsNil(o.PlanMaxCycles) {
		toSerialize["plan_max_cycles"] = o.PlanMaxCycles
	}
	if !IsNil(o.PlanIntervals) {
		toSerialize["plan_intervals"] = o.PlanIntervals
	}
	if !IsNil(o.PlanIntervalType) {
		toSerialize["plan_interval_type"] = o.PlanIntervalType
	}
	if !IsNil(o.PlanNote) {
		toSerialize["plan_note"] = o.PlanNote
	}
	return toSerialize, nil
}



