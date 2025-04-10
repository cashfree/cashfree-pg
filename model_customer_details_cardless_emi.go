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

// checks if the CustomerDetailsCardlessEMI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDetailsCardlessEMI{}

// CustomerDetailsCardlessEMI Details of the customer for whom eligibility is being checked.
type CustomerDetailsCardlessEMI struct {
	// Phone Number of the customer
	CustomerPhone string `json:"customer_phone"`
}


func (o CustomerDetailsCardlessEMI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDetailsCardlessEMI) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["customer_phone"] = o.CustomerPhone
	return toSerialize, nil
}



