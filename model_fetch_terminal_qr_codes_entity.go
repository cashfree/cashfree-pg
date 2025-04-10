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

// checks if the FetchTerminalQRCodesEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchTerminalQRCodesEntity{}

// FetchTerminalQRCodesEntity Fetch Static QR Codes using terminal ID or phone number
type FetchTerminalQRCodesEntity struct {
	// Name of the bank that is linked to the Static QR.
	Bank *string `json:"bank,omitempty"`
	// Base-64 Encoded QR Code URL
	QrCode *string `json:"qrCode,omitempty"`
	// URL of the qr Code.
	QrCodeUrl *string `json:"qrCodeUrl,omitempty"`
	// Status of the static QR.
	Status *string `json:"status,omitempty"`
}


func (o FetchTerminalQRCodesEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchTerminalQRCodesEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.QrCode) {
		toSerialize["qrCode"] = o.QrCode
	}
	if !IsNil(o.QrCodeUrl) {
		toSerialize["qrCodeUrl"] = o.QrCodeUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}



