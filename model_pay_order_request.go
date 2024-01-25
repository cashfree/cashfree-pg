/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the PayOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayOrderRequest{}

// PayOrderRequest Complete object for the pay api that uses payment method objects
type PayOrderRequest struct {
	PaymentSessionId string `json:"payment_session_id"`
	PaymentMethod PayOrderRequestPaymentMethod `json:"payment_method"`
	SaveInstrument *bool `json:"save_instrument,omitempty"`
	// This is required if any offers needs to be applied to the order.
	OfferId *string `json:"offer_id,omitempty"`
}


func (o PayOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payment_session_id"] = o.PaymentSessionId
	toSerialize["payment_method"] = o.PaymentMethod
	if !IsNil(o.SaveInstrument) {
		toSerialize["save_instrument"] = o.SaveInstrument
	}
	if !IsNil(o.OfferId) {
		toSerialize["offer_id"] = o.OfferId
	}
	return toSerialize, nil
}



