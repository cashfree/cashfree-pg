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

// checks if the LinkNotifyEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkNotifyEntity{}

// LinkNotifyEntity Payment link Notify Object for SMS and Email
type LinkNotifyEntity struct {
	// If \"true\", Cashfree will send sms on customer_phone
	SendSms *bool `json:"send_sms,omitempty"`
	// If \"true\", Cashfree will send email on customer_email
	SendEmail *bool `json:"send_email,omitempty"`
}


func (o LinkNotifyEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkNotifyEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SendSms) {
		toSerialize["send_sms"] = o.SendSms
	}
	if !IsNil(o.SendEmail) {
		toSerialize["send_email"] = o.SendEmail
	}
	return toSerialize, nil
}



