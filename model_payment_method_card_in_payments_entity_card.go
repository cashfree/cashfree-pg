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

// checks if the PaymentMethodCardInPaymentsEntityCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCardInPaymentsEntityCard{}

// PaymentMethodCardInPaymentsEntityCard struct for PaymentMethodCardInPaymentsEntityCard
type PaymentMethodCardInPaymentsEntityCard struct {
	Channel *string `json:"channel,omitempty"`
	CardNumber *string `json:"card_number,omitempty"`
	CardNetwork *string `json:"card_network,omitempty"`
	CardType *string `json:"card_type,omitempty"`
	CardCountry *string `json:"card_country,omitempty"`
	CardBankName *string `json:"card_bank_name,omitempty"`
	CardNetworkReferenceId *string `json:"card_network_reference_id,omitempty"`
	InstrumentId *string `json:"instrument_id,omitempty"`
}


func (o PaymentMethodCardInPaymentsEntityCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCardInPaymentsEntityCard) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.CardNumber) {
		toSerialize["card_number"] = o.CardNumber
	}
	if !IsNil(o.CardNetwork) {
		toSerialize["card_network"] = o.CardNetwork
	}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.CardCountry) {
		toSerialize["card_country"] = o.CardCountry
	}
	if !IsNil(o.CardBankName) {
		toSerialize["card_bank_name"] = o.CardBankName
	}
	if !IsNil(o.CardNetworkReferenceId) {
		toSerialize["card_network_reference_id"] = o.CardNetworkReferenceId
	}
	if !IsNil(o.InstrumentId) {
		toSerialize["instrument_id"] = o.InstrumentId
	}
	return toSerialize, nil
}



