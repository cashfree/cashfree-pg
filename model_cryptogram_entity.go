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

// checks if the CryptogramEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptogramEntity{}

// CryptogramEntity Crytogram Card object
type CryptogramEntity struct {
	// instrument_id of saved instrument
	InstrumentId *string `json:"instrument_id,omitempty"`
	// TRID issued by card networks
	TokenRequestorId *string `json:"token_requestor_id,omitempty"`
	// token pan number
	CardNumber *string `json:"card_number,omitempty"`
	// token pan expiry month
	CardExpiryMm *string `json:"card_expiry_mm,omitempty"`
	// token pan expiry year
	CardExpiryYy *string `json:"card_expiry_yy,omitempty"`
	// cryptogram
	Cryptogram *string `json:"cryptogram,omitempty"`
	// last 4 digits of original card number
	CardDisplay *string `json:"card_display,omitempty"`
}


func (o CryptogramEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptogramEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstrumentId) {
		toSerialize["instrument_id"] = o.InstrumentId
	}
	if !IsNil(o.TokenRequestorId) {
		toSerialize["token_requestor_id"] = o.TokenRequestorId
	}
	if !IsNil(o.CardNumber) {
		toSerialize["card_number"] = o.CardNumber
	}
	if !IsNil(o.CardExpiryMm) {
		toSerialize["card_expiry_mm"] = o.CardExpiryMm
	}
	if !IsNil(o.CardExpiryYy) {
		toSerialize["card_expiry_yy"] = o.CardExpiryYy
	}
	if !IsNil(o.Cryptogram) {
		toSerialize["cryptogram"] = o.Cryptogram
	}
	if !IsNil(o.CardDisplay) {
		toSerialize["card_display"] = o.CardDisplay
	}
	return toSerialize, nil
}



