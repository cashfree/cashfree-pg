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
	"fmt"
)

// OfferType Offer Type Object
type OfferType string

// List of OfferType
const (
	DISCOUNT OfferType = "DISCOUNT"
	CASHBACK OfferType = "CASHBACK"
	DISCOUNT_AND_CASHBACK OfferType = "DISCOUNT_AND_CASHBACK"
	NO_COST_EMI OfferType = "NO_COST_EMI"
)

// All allowed values of OfferType enum
var AllowedOfferTypeEnumValues = []OfferType{
	"DISCOUNT",
	"CASHBACK",
	"DISCOUNT_AND_CASHBACK",
	"NO_COST_EMI",
}

func (v *OfferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OfferType(value)
	for _, existing := range AllowedOfferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OfferType", value)
}

// NewOfferTypeFromValue returns a pointer to a valid OfferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOfferTypeFromValue(v string) (*OfferType, error) {
	ev := OfferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OfferType: valid values are %v", v, AllowedOfferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OfferType) IsValid() bool {
	for _, existing := range AllowedOfferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

