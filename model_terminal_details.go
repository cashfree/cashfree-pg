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

// checks if the TerminalDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalDetails{}

// TerminalDetails Use this if you are creating an order for cashfree's softPOS
type TerminalDetails struct {
	// date time at which terminal is added
	AddedOn *string `json:"added_on,omitempty"`
	// Cashfree terminal id, this is a required parameter when you do not provide the terminal phone number.
	CfTerminalId *string `json:"cf_terminal_id,omitempty"`
	// last instant when this terminal was updated
	LastUpdatedOn *string `json:"last_updated_on,omitempty"`
	// location of terminal
	TerminalAddress *string `json:"terminal_address,omitempty"`
	// terminal id for merchant reference
	TerminalId *string `json:"terminal_id,omitempty"`
	// name of terminal/agent/storefront
	TerminalName *string `json:"terminal_name,omitempty"`
	// note given by merchant while creating the terminal
	TerminalNote *string `json:"terminal_note,omitempty"`
	// mobile num of the terminal/agent/storefront,This is a required parameter when you do not provide the cf_terminal_id.
	TerminalPhoneNo *string `json:"terminal_phone_no,omitempty"`
	// status of terminal active/inactive
	TerminalStatus *string `json:"terminal_status,omitempty"`
	// To identify the type of terminal product in use, in this case it is SPOS.
	TerminalType string `json:"terminal_type"`
}


func (o TerminalDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
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
	if !IsNil(o.TerminalId) {
		toSerialize["terminal_id"] = o.TerminalId
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
	toSerialize["terminal_type"] = o.TerminalType
	return toSerialize, nil
}



