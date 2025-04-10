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

// checks if the EMIPlansArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EMIPlansArray{}

// EMIPlansArray Single EMI object
type EMIPlansArray struct {
	Tenure *int32 `json:"tenure,omitempty"`
	InterestRate *float32 `json:"interest_rate,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Emi *int32 `json:"emi,omitempty"`
	TotalInterest *int32 `json:"total_interest,omitempty"`
	TotalAmount *int32 `json:"total_amount,omitempty"`
}


func (o EMIPlansArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EMIPlansArray) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenure) {
		toSerialize["tenure"] = o.Tenure
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interest_rate"] = o.InterestRate
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Emi) {
		toSerialize["emi"] = o.Emi
	}
	if !IsNil(o.TotalInterest) {
		toSerialize["total_interest"] = o.TotalInterest
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["total_amount"] = o.TotalAmount
	}
	return toSerialize, nil
}



