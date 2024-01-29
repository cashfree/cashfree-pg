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

// checks if the TerminalTransactionEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalTransactionEntity{}

// TerminalTransactionEntity Create terminal response object
type TerminalTransactionEntity struct {
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	PaymentAmount *int32 `json:"payment_amount,omitempty"`
	PaymentMethod *string `json:"payment_method,omitempty"`
	PaymentUrl *string `json:"payment_url,omitempty"`
	Qrcode *string `json:"qrcode,omitempty"`
	Timeout *string `json:"timeout,omitempty"`
}


func (o TerminalTransactionEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalTransactionEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentUrl) {
		toSerialize["payment_url"] = o.PaymentUrl
	}
	if !IsNil(o.Qrcode) {
		toSerialize["qrcode"] = o.Qrcode
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}



