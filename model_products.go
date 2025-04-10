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

// checks if the Products type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Products{}

// Products Use this to set configurations for the products like One Click Checkout, Verify and Pay, if they are enabled for your account
type Products struct {
	OneClickCheckout *ProductDetails `json:"one_click_checkout,omitempty"`
	VerifyPay *ProductDetails `json:"verify_pay,omitempty"`
}


func (o Products) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Products) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OneClickCheckout) {
		toSerialize["one_click_checkout"] = o.OneClickCheckout
	}
	if !IsNil(o.VerifyPay) {
		toSerialize["verify_pay"] = o.VerifyPay
	}
	return toSerialize, nil
}



