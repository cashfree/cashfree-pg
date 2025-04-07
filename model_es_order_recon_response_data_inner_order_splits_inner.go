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

// checks if the ESOrderReconResponseDataInnerOrderSplitsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESOrderReconResponseDataInnerOrderSplitsInner{}

// ESOrderReconResponseDataInnerOrderSplitsInner struct for ESOrderReconResponseDataInnerOrderSplitsInner
type ESOrderReconResponseDataInnerOrderSplitsInner struct {
	Split []ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner `json:"split,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}


func (o ESOrderReconResponseDataInnerOrderSplitsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESOrderReconResponseDataInnerOrderSplitsInner) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Split) {
		toSerialize["split"] = o.Split
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}



