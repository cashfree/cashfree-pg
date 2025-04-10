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

// checks if the EvidencesToContestDispute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvidencesToContestDispute{}

// EvidencesToContestDispute struct for EvidencesToContestDispute
type EvidencesToContestDispute struct {
	DocumentType *string `json:"document_type,omitempty"`
	DocumentDescription *string `json:"document_description,omitempty"`
}


func (o EvidencesToContestDispute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvidencesToContestDispute) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentType) {
		toSerialize["document_type"] = o.DocumentType
	}
	if !IsNil(o.DocumentDescription) {
		toSerialize["document_description"] = o.DocumentDescription
	}
	return toSerialize, nil
}



