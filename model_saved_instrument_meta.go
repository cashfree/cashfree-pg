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

// checks if the SavedInstrumentMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedInstrumentMeta{}

// SavedInstrumentMeta Card instrument meta information
type SavedInstrumentMeta struct {
	// card scheme/network of the saved card. Example visa, mastercard
	CardNetwork *string `json:"card_network,omitempty"`
	// Issuing bank name of saved card
	CardBankName *string `json:"card_bank_name,omitempty"`
	// Issuing country of saved card
	CardCountry *string `json:"card_country,omitempty"`
	// Type of saved card
	CardType *string `json:"card_type,omitempty"`
	CardTokenDetails map[string]interface{} `json:"card_token_details,omitempty"`
}


func (o SavedInstrumentMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedInstrumentMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardNetwork) {
		toSerialize["card_network"] = o.CardNetwork
	}
	if !IsNil(o.CardBankName) {
		toSerialize["card_bank_name"] = o.CardBankName
	}
	if !IsNil(o.CardCountry) {
		toSerialize["card_country"] = o.CardCountry
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.CardTokenDetails) {
		toSerialize["card_token_details"] = o.CardTokenDetails
	}
	return toSerialize, nil
}



