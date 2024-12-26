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

// checks if the ShipmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDetails{}

// ShipmentDetails Shipment details associated with shipping of order like tracking company, tracking number,tracking urls etc.
type ShipmentDetails struct {
	// Tracking company name associated with order.
	TrackingCompany string `json:"tracking_company"`
	// Tracking Urls associated with order.
	TrackingUrls []string `json:"tracking_urls"`
	// Tracking Numbers associated with order.
	TrackingNumbers []string `json:"tracking_numbers"`
}


func (o ShipmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["tracking_company"] = o.TrackingCompany
	toSerialize["tracking_urls"] = o.TrackingUrls
	toSerialize["tracking_numbers"] = o.TrackingNumbers
	return toSerialize, nil
}



