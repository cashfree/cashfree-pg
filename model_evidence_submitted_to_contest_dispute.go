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

// checks if the EvidenceSubmittedToContestDispute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvidenceSubmittedToContestDispute{}

// EvidenceSubmittedToContestDispute struct for EvidenceSubmittedToContestDispute
type EvidenceSubmittedToContestDispute struct {
	DocumentId *int32 `json:"documentId,omitempty"`
	DocumentName *string `json:"documentName,omitempty"`
	DocumentType *string `json:"documentType,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}


func (o EvidenceSubmittedToContestDispute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvidenceSubmittedToContestDispute) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentId) {
		toSerialize["documentId"] = o.DocumentId
	}
	if !IsNil(o.DocumentName) {
		toSerialize["documentName"] = o.DocumentName
	}
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	return toSerialize, nil
}


