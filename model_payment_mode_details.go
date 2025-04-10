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

// checks if the PaymentModeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentModeDetails{}

// PaymentModeDetails payment mode eligiblity object
type PaymentModeDetails struct {
	Nick *string `json:"nick,omitempty"`
	Display *string `json:"display,omitempty"`
	Eligibility *bool `json:"eligibility,omitempty"`
	Code *float32 `json:"code,omitempty"`
}


func (o PaymentModeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentModeDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nick) {
		toSerialize["nick"] = o.Nick
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Eligibility) {
		toSerialize["eligibility"] = o.Eligibility
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}



