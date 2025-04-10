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

// checks if the UpiDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpiDetails{}

// UpiDetails struct for UpiDetails
type UpiDetails struct {
	Vpa *string `json:"vpa,omitempty"`
	AccountHolder *string `json:"account_holder,omitempty"`
}


func (o UpiDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpiDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vpa) {
		toSerialize["vpa"] = o.Vpa
	}
	if !IsNil(o.AccountHolder) {
		toSerialize["account_holder"] = o.AccountHolder
	}
	return toSerialize, nil
}



