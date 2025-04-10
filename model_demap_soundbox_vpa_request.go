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

// checks if the DemapSoundboxVpaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemapSoundboxVpaRequest{}

// DemapSoundboxVpaRequest Request body to demap soundbox vpa
type DemapSoundboxVpaRequest struct {
	// cashfree terminal id.
	CfTerminalId string `json:"cf_terminal_id"`
	// Device Serial No of soundbox that need to demap.
	DeviceSerialNo string `json:"device_serial_no"`
}


func (o DemapSoundboxVpaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemapSoundboxVpaRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["cf_terminal_id"] = o.CfTerminalId
	toSerialize["device_serial_no"] = o.DeviceSerialNo
	return toSerialize, nil
}



