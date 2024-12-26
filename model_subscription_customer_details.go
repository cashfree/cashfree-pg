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

// checks if the SubscriptionCustomerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCustomerDetails{}

// SubscriptionCustomerDetails Subscription customer details.
type SubscriptionCustomerDetails struct {
	// Name of the customer.
	CustomerName *string `json:"customer_name,omitempty"`
	// Email of the customer.
	CustomerEmail string `json:"customer_email"`
	// Phone number of the customer.
	CustomerPhone string `json:"customer_phone"`
	// Bank holder name of the customer.
	CustomerBankAccountHolderName *string `json:"customer_bank_account_holder_name,omitempty"`
	// Bank account number of the customer.
	CustomerBankAccountNumber *string `json:"customer_bank_account_number,omitempty"`
	// IFSC code of the customer.
	CustomerBankIfsc *string `json:"customer_bank_ifsc,omitempty"`
	// Bank code of the customer. Refer to https://www.npci.org.in/PDF/nach/live-members-e-mandates/Live-Banks-in-API-E-Mandate.pdf
	CustomerBankCode *string `json:"customer_bank_code,omitempty"`
	// Bank account type of the customer.
	CustomerBankAccountType *string `json:"customer_bank_account_type,omitempty"`
}


func (o SubscriptionCustomerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCustomerDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	toSerialize["customer_email"] = o.CustomerEmail
	toSerialize["customer_phone"] = o.CustomerPhone
	if !IsNil(o.CustomerBankAccountHolderName) {
		toSerialize["customer_bank_account_holder_name"] = o.CustomerBankAccountHolderName
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
	if !IsNil(o.CustomerBankAccountType) {
		toSerialize["customer_bank_account_type"] = o.CustomerBankAccountType
	}
	return toSerialize, nil
}



