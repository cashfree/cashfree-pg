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

// checks if the CardPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardPaymentMethod{}

// CardPaymentMethod The card payment object is used to make payment using either plain card number, saved card instrument id or using cryptogram 
type CardPaymentMethod struct {
	Card Card `json:"card"`
}


func (o CardPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardPaymentMethod) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["card"] = o.Card
	return toSerialize, nil
}



