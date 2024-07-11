/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the EntitySimulationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySimulationResponse{}

// EntitySimulationResponse Entity Simulation it contains payment_status and payment_error_code
type EntitySimulationResponse struct {
	// Payment Status
	PaymentStatus string `json:"payment_status"`
	// Payment Error Code
	PaymentErrorCode *string `json:"payment_error_code,omitempty"`
}


func (o EntitySimulationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySimulationResponse) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["payment_status"] = o.PaymentStatus
	if !IsNil(o.PaymentErrorCode) {
		toSerialize["payment_error_code"] = o.PaymentErrorCode
	}
	return toSerialize, nil
}



