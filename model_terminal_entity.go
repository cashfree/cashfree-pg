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

// checks if the TerminalEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalEntity{}

// TerminalEntity Create terminal response object
type TerminalEntity struct {
	AddedOn *string `json:"added_on,omitempty"`
	CfTerminalId *int32 `json:"cf_terminal_id,omitempty"`
	LastUpdatedOn *string `json:"last_updated_on,omitempty"`
	TerminalAddress *string `json:"terminal_address,omitempty"`
	TerminalEmail *string `json:"terminal_email,omitempty"`
	TerminalType *string `json:"terminal_type,omitempty"`
	TeminalId *string `json:"teminal_id,omitempty"`
	TerminalName *string `json:"terminal_name,omitempty"`
	TerminalNote *string `json:"terminal_note,omitempty"`
	TerminalPhoneNo *string `json:"terminal_phone_no,omitempty"`
	TerminalStatus *string `json:"terminal_status,omitempty"`
	TerminalMeta *string `json:"terminal_meta,omitempty"`
}


func (o TerminalEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	if !IsNil(o.CfTerminalId) {
		toSerialize["cf_terminal_id"] = o.CfTerminalId
	}
	if !IsNil(o.LastUpdatedOn) {
		toSerialize["last_updated_on"] = o.LastUpdatedOn
	}
	if !IsNil(o.TerminalAddress) {
		toSerialize["terminal_address"] = o.TerminalAddress
	}
	if !IsNil(o.TerminalEmail) {
		toSerialize["terminal_email"] = o.TerminalEmail
	}
	if !IsNil(o.TerminalType) {
		toSerialize["terminal_type"] = o.TerminalType
	}
	if !IsNil(o.TeminalId) {
		toSerialize["teminal_id"] = o.TeminalId
	}
	if !IsNil(o.TerminalName) {
		toSerialize["terminal_name"] = o.TerminalName
	}
	if !IsNil(o.TerminalNote) {
		toSerialize["terminal_note"] = o.TerminalNote
	}
	if !IsNil(o.TerminalPhoneNo) {
		toSerialize["terminal_phone_no"] = o.TerminalPhoneNo
	}
	if !IsNil(o.TerminalStatus) {
		toSerialize["terminal_status"] = o.TerminalStatus
	}
	if !IsNil(o.TerminalMeta) {
		toSerialize["terminal_meta"] = o.TerminalMeta
	}
	return toSerialize, nil
}



