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

// checks if the OfferTnc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferTnc{}

// OfferTnc Offer terms and condition object
type OfferTnc struct {
	// TnC Type for the Offer. It can be either `text` or `link`
	OfferTncType string `json:"offer_tnc_type"`
	// TnC for the Offer.
	OfferTncValue string `json:"offer_tnc_value"`
}


func (o OfferTnc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferTnc) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["offer_tnc_type"] = o.OfferTncType
	toSerialize["offer_tnc_value"] = o.OfferTncValue
	return toSerialize, nil
}



