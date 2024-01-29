/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the Netbanking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Netbanking{}

// Netbanking Netbanking payment method request body
type Netbanking struct {
	// The channel for netbanking will always be `link`
	Channel string `json:"channel"`
	// Bank code
	NetbankingBankCode *int32 `json:"netbanking_bank_code,omitempty"`
	// String code for bank
	NetbankingBankName *string `json:"netbanking_bank_name,omitempty"`
}


func (o Netbanking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Netbanking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.NetbankingBankCode) {
		toSerialize["netbanking_bank_code"] = o.NetbankingBankCode
	}
	if !IsNil(o.NetbankingBankName) {
		toSerialize["netbanking_bank_name"] = o.NetbankingBankName
	}
	return toSerialize, nil
}



