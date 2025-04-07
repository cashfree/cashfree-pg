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

// checks if the PayOrderEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayOrderEntity{}

// PayOrderEntity Order Pay response once you create a transaction for that order
type PayOrderEntity struct {
	// total amount payable
	PaymentAmount *float32 `json:"payment_amount,omitempty"`
	// Payment identifier created by Cashfree
	CfPaymentId *string `json:"cf_payment_id,omitempty"`
	// One of [\"upi\", \"netbanking\", \"card\", \"app\", \"cardless_emi\", \"paylater\", \"banktransfer\"] 
	PaymentMethod *string `json:"payment_method,omitempty"`
	// One of [\"link\", \"collect\", \"qrcode\"]. In an older version we used to support different channels like 'gpay', 'phonepe' etc. However, we now support only the following channels - link, collect and qrcode. To process payments using gpay, you will have to provide channel as 'link' and provider as 'gpay'
	Channel *string `json:"channel,omitempty"`
	// One of [\"link\", \"custom\", \"form\"]
	Action *string `json:"action,omitempty"`
	Data *OrderPayData `json:"data,omitempty"`
}


func (o PayOrderEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayOrderEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentAmount) {
		toSerialize["payment_amount"] = o.PaymentAmount
	}
	if !IsNil(o.CfPaymentId) {
		toSerialize["cf_payment_id"] = o.CfPaymentId
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}



