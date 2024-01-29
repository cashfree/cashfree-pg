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

// checks if the PaymentMethodNetBankingInPaymentsEntityNetbanking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodNetBankingInPaymentsEntityNetbanking{}

// PaymentMethodNetBankingInPaymentsEntityNetbanking struct for PaymentMethodNetBankingInPaymentsEntityNetbanking
type PaymentMethodNetBankingInPaymentsEntityNetbanking struct {
	Channel *string `json:"channel,omitempty"`
	NetbankingBankCode *int32 `json:"netbanking_bank_code,omitempty"`
	NetbankingBankName *string `json:"netbanking_bank_name,omitempty"`
}


func (o PaymentMethodNetBankingInPaymentsEntityNetbanking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodNetBankingInPaymentsEntityNetbanking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.NetbankingBankCode) {
		toSerialize["netbanking_bank_code"] = o.NetbankingBankCode
	}
	if !IsNil(o.NetbankingBankName) {
		toSerialize["netbanking_bank_name"] = o.NetbankingBankName
	}
	return toSerialize, nil
}



