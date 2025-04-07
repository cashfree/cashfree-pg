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

// checks if the OrderPayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPayData{}

// OrderPayData the data object pay api
type OrderPayData struct {
	Url *string `json:"url,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Method *string `json:"method,omitempty"`
	RedirectToBank *string `json:"redirect_to_bank,omitempty"`
}


func (o OrderPayData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPayData) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.RedirectToBank) {
		toSerialize["redirect_to_bank"] = o.RedirectToBank
	}
	return toSerialize, nil
}



