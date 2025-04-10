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

// checks if the CreateSubscriptionPaymentRequestEnach type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionPaymentRequestEnach{}

// CreateSubscriptionPaymentRequestEnach payment method enach.
type CreateSubscriptionPaymentRequestEnach struct {
	// Account bank code (required without AccountIFSC)
	AccountBankCode *string `json:"account_bank_code,omitempty"`
	// Account holder name
	AccountHolderName *string `json:"account_holder_name,omitempty"`
	// Account IFSC
	AccountIfsc *string `json:"account_ifsc,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Account type
	AccountType *string `json:"account_type,omitempty"`
	// Authentication mode. can be debit_card, aadhaar, or net_banking
	AuthMode *string `json:"auth_mode,omitempty"`
	// Channel. can be link
	Channel *string `json:"channel,omitempty"`
}


func (o CreateSubscriptionPaymentRequestEnach) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionPaymentRequestEnach) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountBankCode) {
		toSerialize["account_bank_code"] = o.AccountBankCode
	}
	if !IsNil(o.AccountHolderName) {
		toSerialize["account_holder_name"] = o.AccountHolderName
	}
	if !IsNil(o.AccountIfsc) {
		toSerialize["account_ifsc"] = o.AccountIfsc
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.AuthMode) {
		toSerialize["auth_mode"] = o.AuthMode
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}



