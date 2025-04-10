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

// checks if the FetchSettlementsRequestFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSettlementsRequestFilters{}

// FetchSettlementsRequestFilters Specify either the Settlement ID, Settlement UTR, or start date and end date to fetch the settlement details.
type FetchSettlementsRequestFilters struct {
	// List of settlement IDs for which you want the settlement reconciliation details.
	CfSettlementIds []string `json:"cf_settlement_ids,omitempty"`
	// List of settlement UTRs for which you want the settlement reconciliation details.
	SettlementUtrs []string `json:"settlement_utrs,omitempty"`
	// Specify the start date from when you want the settlement reconciliation details.
	StartDate *string `json:"start_date,omitempty"`
	// Specify the end date till when you want the settlement reconciliation details.
	EndDate *string `json:"end_date,omitempty"`
}


func (o FetchSettlementsRequestFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSettlementsRequestFilters) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CfSettlementIds) {
		toSerialize["cf_settlement_ids"] = o.CfSettlementIds
	}
	if !IsNil(o.SettlementUtrs) {
		toSerialize["settlement_utrs"] = o.SettlementUtrs
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}



