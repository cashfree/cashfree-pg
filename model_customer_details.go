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

// checks if the CustomerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDetails{}

// CustomerDetails The customer details that are necessary. Note that you can pass dummy details if your use case does not require the customer details.
type CustomerDetails struct {
	// A unique identifier for the customer. Use alphanumeric values only.
	CustomerId string `json:"customer_id"`
	// Customer email address.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Customer phone number.
	CustomerPhone string `json:"customer_phone"`
	// Name of the customer.
	CustomerName *string `json:"customer_name,omitempty"`
	// Customer bank account. Required if you want to do a bank account check (TPV)
	CustomerBankAccountNumber *string `json:"customer_bank_account_number,omitempty"`
	// Customer bank IFSC. Required if you want to do a bank account check (TPV)
	CustomerBankIfsc *string `json:"customer_bank_ifsc,omitempty"`
	// Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV)
	CustomerBankCode *float32 `json:"customer_bank_code,omitempty"`
	// Customer identifier at Cashfree. You will get this when you create/get customer
	CustomerUid *string `json:"customer_uid,omitempty"`
}


func (o CustomerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["customer_id"] = o.CustomerId
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	toSerialize["customer_phone"] = o.CustomerPhone
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !IsNil(o.CustomerBankAccountNumber) {
		toSerialize["customer_bank_account_number"] = o.CustomerBankAccountNumber
	}
	if !IsNil(o.CustomerBankIfsc) {
		toSerialize["customer_bank_ifsc"] = o.CustomerBankIfsc
	}
	if !IsNil(o.CustomerBankCode) {
		toSerialize["customer_bank_code"] = o.CustomerBankCode
	}
	if !IsNil(o.CustomerUid) {
		toSerialize["customer_uid"] = o.CustomerUid
	}
	return toSerialize, nil
}



