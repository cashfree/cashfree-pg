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

// checks if the ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner{}

// ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner struct for ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner
type ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner struct {
	MerchantVendorId *string `json:"merchant_vendor_id,omitempty"`
	Percentage *float32 `json:"percentage,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}


func (o ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESOrderReconResponseDataInnerOrderSplitsInnerSplitInner) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantVendorId) {
		toSerialize["merchant_vendor_id"] = o.MerchantVendorId
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}



