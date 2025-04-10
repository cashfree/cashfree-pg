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

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App App payment method
type App struct {
	// Specify the channel through which the payment must be processed.
	Channel string `json:"channel"`
	// Specify the provider through which the payment must be processed.
	Provider string `json:"provider"`
	// Customer phone number associated with a wallet for payment.
	Phone string `json:"phone"`
}


func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["provider"] = o.Provider
	toSerialize["phone"] = o.Phone
	return toSerialize, nil
}



