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

// checks if the DisputesEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisputesEntity{}

// DisputesEntity struct for DisputesEntity
type DisputesEntity struct {
	DisputeId *int32 `json:"dispute_id,omitempty"`
	DisputeType *string `json:"dispute_type,omitempty"`
	ReasonCode *string `json:"reason_code,omitempty"`
	ReasonDescription *string `json:"reason_description,omitempty"`
	// Dispute amount may differ from transaction amount for partial cases.
	DisputeAmount *float32 `json:"dispute_amount,omitempty"`
	// Dispute amount currency for a dispute
	DisputeAmountCurrency *string `json:"dispute_amount_currency,omitempty"`
	// This is the time when the dispute was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// This is the time by which evidence should be submitted to contest the dispute.
	RespondBy *string `json:"respond_by,omitempty"`
	// This is the time when the dispute case was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// This is the time when the dispute case was closed.
	ResolvedAt *string `json:"resolved_at,omitempty"`
	DisputeStatus *string `json:"dispute_status,omitempty"`
	CfDisputeRemarks *string `json:"cf_dispute_remarks,omitempty"`
	PreferredEvidence []PreferredEvidenceInner `json:"preferred_evidence,omitempty"`
	DisputeEvidence []DisputeEvidenceInner `json:"dispute_evidence,omitempty"`
	OrderDetails *OrderDetailsInDisputesEntity `json:"order_details,omitempty"`
	CustomerDetails *CustomerDetailsInDisputesEntity `json:"customer_details,omitempty"`
}


func (o DisputesEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputesEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisputeId) {
		toSerialize["dispute_id"] = o.DisputeId
	}
	if !IsNil(o.DisputeType) {
		toSerialize["dispute_type"] = o.DisputeType
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	if !IsNil(o.ReasonDescription) {
		toSerialize["reason_description"] = o.ReasonDescription
	}
	if !IsNil(o.DisputeAmount) {
		toSerialize["dispute_amount"] = o.DisputeAmount
	}
	if !IsNil(o.DisputeAmountCurrency) {
		toSerialize["dispute_amount_currency"] = o.DisputeAmountCurrency
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.RespondBy) {
		toSerialize["respond_by"] = o.RespondBy
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ResolvedAt) {
		toSerialize["resolved_at"] = o.ResolvedAt
	}
	if !IsNil(o.DisputeStatus) {
		toSerialize["dispute_status"] = o.DisputeStatus
	}
	if !IsNil(o.CfDisputeRemarks) {
		toSerialize["cf_dispute_remarks"] = o.CfDisputeRemarks
	}
	if !IsNil(o.PreferredEvidence) {
		toSerialize["preferred_evidence"] = o.PreferredEvidence
	}
	if !IsNil(o.DisputeEvidence) {
		toSerialize["dispute_evidence"] = o.DisputeEvidence
	}
	if !IsNil(o.OrderDetails) {
		toSerialize["order_details"] = o.OrderDetails
	}
	if !IsNil(o.CustomerDetails) {
		toSerialize["customer_details"] = o.CustomerDetails
	}
	return toSerialize, nil
}



