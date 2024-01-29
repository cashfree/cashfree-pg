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

// checks if the PaymentMethodsFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodsFilters{}

// PaymentMethodsFilters Filter for Payment Methods
type PaymentMethodsFilters struct {
	// Array of payment methods to be filtered. This is optional, by default all payment methods will be returned. Possible values in [ 'debit_card', 'credit_card', 'prepaid_card', 'corporate_credit_card', 'upi', 'wallet', 'netbanking', 'banktransfer', 'paylater', 'paypal', 'debit_card_emi', 'credit_card_emi', 'upi_credit_card', 'upi_ppi', 'cardless_emi', 'account_based_payment' ] 
	PaymentMethods []string `json:"payment_methods,omitempty"`
}


func (o PaymentMethodsFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return toSerialize, nil
}



