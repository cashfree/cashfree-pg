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

// checks if the ESOrderReconRequestPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESOrderReconRequestPagination{}

// ESOrderReconRequestPagination Set limit based on your requirement. Pagination limit will fetch a set of orders, next set of orders can be generated using the cursor shared in previous response of the same API.
type ESOrderReconRequestPagination struct {
	Cursor *string `json:"cursor,omitempty"`
	// Set the minimum/maximum limit for number of filtered data. Min value - 10, Max value - 100.
	Limit *int32 `json:"limit,omitempty"`
}


func (o ESOrderReconRequestPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESOrderReconRequestPagination) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}



