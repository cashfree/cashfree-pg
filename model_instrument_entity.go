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

// checks if the InstrumentEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstrumentEntity{}

// InstrumentEntity Saved card instrument object
type InstrumentEntity struct {
	// customer_id for which the instrument was saved
	CustomerId *string `json:"customer_id,omitempty"`
	// cf_payment_id of the successful transaction done while saving instrument
	AfaReference *string `json:"afa_reference,omitempty"`
	// saved instrument id
	InstrumentId *string `json:"instrument_id,omitempty"`
	// Type of the saved instrument
	InstrumentType *string `json:"instrument_type,omitempty"`
	// Unique id for the saved instrument
	InstrumentUid *string `json:"instrument_uid,omitempty"`
	// masked card number displayed to the customer
	InstrumentDisplay *string `json:"instrument_display,omitempty"`
	// Status of the saved instrument.
	InstrumentStatus *string `json:"instrument_status,omitempty"`
	// Timestamp at which instrument was saved.
	CreatedAt *string `json:"created_at,omitempty"`
	InstrumentMeta *SavedInstrumentMeta `json:"instrument_meta,omitempty"`
}


func (o InstrumentEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstrumentEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.AfaReference) {
		toSerialize["afa_reference"] = o.AfaReference
	}
	if !IsNil(o.InstrumentId) {
		toSerialize["instrument_id"] = o.InstrumentId
	}
	if !IsNil(o.InstrumentType) {
		toSerialize["instrument_type"] = o.InstrumentType
	}
	if !IsNil(o.InstrumentUid) {
		toSerialize["instrument_uid"] = o.InstrumentUid
	}
	if !IsNil(o.InstrumentDisplay) {
		toSerialize["instrument_display"] = o.InstrumentDisplay
	}
	if !IsNil(o.InstrumentStatus) {
		toSerialize["instrument_status"] = o.InstrumentStatus
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.InstrumentMeta) {
		toSerialize["instrument_meta"] = o.InstrumentMeta
	}
	return toSerialize, nil
}



