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

// checks if the EligibilityPaymentMethodsEntityEntityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityPaymentMethodsEntityEntityDetails{}

// EligibilityPaymentMethodsEntityEntityDetails struct for EligibilityPaymentMethodsEntityEntityDetails
type EligibilityPaymentMethodsEntityEntityDetails struct {
	PaymentMethodDetails []PaymentModeDetails `json:"payment_method_details,omitempty"`
}


func (o EligibilityPaymentMethodsEntityEntityDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityPaymentMethodsEntityEntityDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethodDetails) {
		toSerialize["payment_method_details"] = o.PaymentMethodDetails
	}
	return toSerialize, nil
}



