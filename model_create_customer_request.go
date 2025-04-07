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

// checks if the CreateCustomerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerRequest{}

// CreateCustomerRequest Request body to create a customer at cashfree
type CreateCustomerRequest struct {
	// Customer Phone Number
	CustomerPhone string `json:"customer_phone"`
	// Customer Email
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Customer Name
	CustomerName *string `json:"customer_name,omitempty"`
}


func (o CreateCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["customer_phone"] = o.CustomerPhone
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	return toSerialize, nil
}



