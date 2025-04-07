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

// checks if the OrderMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderMeta{}

// OrderMeta Optional meta details to control how the customer pays and how payment journey completes
type OrderMeta struct {
	// The URL to which user will be redirected to after the payment on bank OTP page. Maximum length: 250. We suggest to keep context of order_id in your return_url so that you can identify the order when customer lands on your page. Example of return_url format could be https://www.cashfree.com/devstudio/thankyou
	ReturnUrl *string `json:"return_url,omitempty"`
	// Notification URL for server-server communication. Useful when user's connection drops while re-directing. NotifyUrl should be an https URL. Maximum length: 250.
	NotifyUrl *string `json:"notify_url,omitempty"`
	// Allowed payment modes for this order. Pass comma-separated values among following options - \"cc\", \"dc\", \"ccc\", \"ppc\",\"nb\",\"upi\",\"paypal\",\"app\",\"paylater\",\"cardlessemi\",\"dcemi\",\"ccemi\",\"banktransfer\". Leave it blank to show all available payment methods
	PaymentMethods interface{} `json:"payment_methods,omitempty"`
	PaymentMethodsFilters *OrderMetaPaymentMethodsFilters `json:"payment_methods_filters,omitempty"`
}


func (o OrderMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderMeta) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.PaymentMethodsFilters) {
		toSerialize["payment_methods_filters"] = o.PaymentMethodsFilters
	}
	return toSerialize, nil
}



