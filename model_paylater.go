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

// checks if the Paylater type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Paylater{}

// Paylater Paylater payment method
type Paylater struct {
	// The channel for cardless EMI is always `link`
	Channel *string `json:"channel,omitempty"`
	// One of [\"kotak\", \"flexipay\", \"zestmoney\", \"lazypay\", \"olapostpaid\",\"simpl\", \"freechargepaylater\"]. Please note that Flexipay is offered by HDFC bank
	Provider *string `json:"provider,omitempty"`
	// Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as 'invalid_request_error' and code as 'invalid_request_error'
	Phone *string `json:"phone,omitempty"`
}


func (o Paylater) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Paylater) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}



