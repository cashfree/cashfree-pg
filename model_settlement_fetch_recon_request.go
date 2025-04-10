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

// checks if the SettlementFetchReconRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementFetchReconRequest{}

// SettlementFetchReconRequest Recon Request Object
type SettlementFetchReconRequest struct {
	Pagination FetchSettlementsRequestPagination `json:"pagination"`
	Filters FetchSettlementsRequestFilters `json:"filters"`
}


func (o SettlementFetchReconRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementFetchReconRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["filters"] = o.Filters
	return toSerialize, nil
}



