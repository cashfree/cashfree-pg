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

// checks if the UPIAuthorizeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UPIAuthorizeDetails{}

// UPIAuthorizeDetails object when you are using preauth in UPI in order pay
type UPIAuthorizeDetails struct {
	// Time by which this authorization should be approved by the customer.
	ApproveBy *string `json:"approve_by,omitempty"`
	// This is the time when the UPI one time mandate will start
	StartTime *string `json:"start_time,omitempty"`
	// This is the time when the UPI mandate will be over. If the mandate has not been executed by this time, the funds will be returned back to the customer after this time.
	EndTime *string `json:"end_time,omitempty"`
}


func (o UPIAuthorizeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UPIAuthorizeDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApproveBy) {
		toSerialize["approve_by"] = o.ApproveBy
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	return toSerialize, nil
}



