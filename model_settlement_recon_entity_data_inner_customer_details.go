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

// checks if the SettlementReconEntityDataInnerCustomerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementReconEntityDataInnerCustomerDetails{}

// SettlementReconEntityDataInnerCustomerDetails struct for SettlementReconEntityDataInnerCustomerDetails
type SettlementReconEntityDataInnerCustomerDetails struct {
	// Customer phone number.
	CustomerPhone *string `json:"customer_phone,omitempty"`
	// Customer email.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Customer name.
	CustomerName *string `json:"customer_name,omitempty"`
	// Customer's id.
	CustomerId *string `json:"customer_id,omitempty"`
	// Customer bank account number.
	CustomerBankAccountNumber *string `json:"customer_bank_account_number,omitempty"`
	// Customer bank code.
	CustomerBankCode *string `json:"customer_bank_code,omitempty"`
	// Customer bank ifsc\"
	CustomerBankIfsc *string `json:"customer_bank_ifsc,omitempty"`
}


func (o SettlementReconEntityDataInnerCustomerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementReconEntityDataInnerCustomerDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerPhone) {
		toSerialize["customer_phone"] = o.CustomerPhone
	}
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.CustomerBankAccountNumber) {
		toSerialize["customer_bank_account_number"] = o.CustomerBankAccountNumber
	}
	if !IsNil(o.CustomerBankCode) {
		toSerialize["customer_bank_code"] = o.CustomerBankCode
	}
	if !IsNil(o.CustomerBankIfsc) {
		toSerialize["customer_bank_ifsc"] = o.CustomerBankIfsc
	}
	return toSerialize, nil
}



