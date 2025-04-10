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

// checks if the AuthorizeOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeOrderRequest{}

// AuthorizeOrderRequest Request to capture or void transaction
type AuthorizeOrderRequest struct {
	// Type of authorization to run. Can be one of 'CAPTURE' , 'VOID'
	Action *string `json:"action,omitempty"`
	// The amount if you are running a 'CAPTURE'
	Amount *float32 `json:"amount,omitempty"`
}


func (o AuthorizeOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeOrderRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}



