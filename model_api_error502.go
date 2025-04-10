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

// checks if the ApiError502 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiError502{}

// ApiError502 Error when there is error at partner bank
type ApiError502 struct {
	Message *string `json:"message,omitempty"`
	Help *string `json:"help,omitempty"`
	// `bank_processing_failure` will be returned here to denote failure at bank. 
	Code *string `json:"code,omitempty"`
	// api_error
	Type *string `json:"type,omitempty"`
}


func (o ApiError502) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiError502) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}



