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

// checks if the CreateSubscriptonPaymentRequestUpi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptonPaymentRequestUpi{}

// CreateSubscriptonPaymentRequestUpi payment method upi.
type CreateSubscriptonPaymentRequestUpi struct {
	// Channel. can be link, qrcode, or collect
	Channel *string `json:"channel,omitempty"`
	UpiId *string `json:"upi_id,omitempty"`
}


func (o CreateSubscriptonPaymentRequestUpi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptonPaymentRequestUpi) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.UpiId) {
		toSerialize["upi_id"] = o.UpiId
	}
	return toSerialize, nil
}



