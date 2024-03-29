/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
)

// checks if the WHcustomerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WHcustomerDetails{}

// WHcustomerDetails customer details object in webhook
type WHcustomerDetails struct {
	CustomerName *string `json:"customer_name,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerEmail *string `json:"customer_email,omitempty"`
	CustomerPhone *string `json:"customer_phone,omitempty"`
}


func (o WHcustomerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WHcustomerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.CustomerPhone) {
		toSerialize["customer_phone"] = o.CustomerPhone
	}
	return toSerialize, nil
}



