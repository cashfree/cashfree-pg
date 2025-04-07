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

// checks if the OfferMetaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferMetaResponse{}

// OfferMetaResponse Offer meta response details object
type OfferMetaResponse struct {
	// Unique identifier for the Offer.
	OfferCode *string `json:"offer_code,omitempty"`
	// Description for the Offer.
	OfferDescription *string `json:"offer_description,omitempty"`
	// Expiry Time for the Offer
	OfferEndTime *string `json:"offer_end_time,omitempty"`
	// Start Time for the Offer
	OfferStartTime *string `json:"offer_start_time,omitempty"`
	// Title for the Offer.
	OfferTitle *string `json:"offer_title,omitempty"`
}


func (o OfferMetaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferMetaResponse) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OfferCode) {
		toSerialize["offer_code"] = o.OfferCode
	}
	if !IsNil(o.OfferDescription) {
		toSerialize["offer_description"] = o.OfferDescription
	}
	if !IsNil(o.OfferEndTime) {
		toSerialize["offer_end_time"] = o.OfferEndTime
	}
	if !IsNil(o.OfferStartTime) {
		toSerialize["offer_start_time"] = o.OfferStartTime
	}
	if !IsNil(o.OfferTitle) {
		toSerialize["offer_title"] = o.OfferTitle
	}
	return toSerialize, nil
}



