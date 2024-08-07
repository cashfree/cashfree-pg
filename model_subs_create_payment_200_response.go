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

// SubsCreatePayment200Response - struct for SubsCreatePayment200Response
type SubsCreatePayment200Response struct {
	CreateSubscriptionPaymentAuthResponse *CreateSubscriptionPaymentAuthResponse
	CreateSubscriptionPaymentChargeResponse *CreateSubscriptionPaymentChargeResponse
}

// CreateSubscriptionPaymentAuthResponseAsSubsCreatePayment200Response is a convenience function that returns CreateSubscriptionPaymentAuthResponse wrapped in SubsCreatePayment200Response
func CreateSubscriptionPaymentAuthResponseAsSubsCreatePayment200Response(v *CreateSubscriptionPaymentAuthResponse) SubsCreatePayment200Response {
	return SubsCreatePayment200Response{
		CreateSubscriptionPaymentAuthResponse: v,
	}
}

// CreateSubscriptionPaymentChargeResponseAsSubsCreatePayment200Response is a convenience function that returns CreateSubscriptionPaymentChargeResponse wrapped in SubsCreatePayment200Response
func CreateSubscriptionPaymentChargeResponseAsSubsCreatePayment200Response(v *CreateSubscriptionPaymentChargeResponse) SubsCreatePayment200Response {
	return SubsCreatePayment200Response{
		CreateSubscriptionPaymentChargeResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubsCreatePayment200Response) UnmarshalJSON(data []byte) error {
		var err error





	match := 0


	// try to unmarshal data into CreateSubscriptionPaymentAuthResponse

	err = json.Unmarshal(data, &dst.CreateSubscriptionPaymentAuthResponse)

	if err == nil {

		jsonCreateSubscriptionPaymentAuthResponse, _ := json.Marshal(dst.CreateSubscriptionPaymentAuthResponse)

		if strings.Contains(string(jsonCreateSubscriptionPaymentAuthResponse), "{}") || strings.Contains(string(jsonCreateSubscriptionPaymentAuthResponse), "null") { // empty struct

			dst.CreateSubscriptionPaymentAuthResponse = nil

		} else {

			match++

		}

	} else {

		dst.CreateSubscriptionPaymentAuthResponse = nil

	}


	// try to unmarshal data into CreateSubscriptionPaymentChargeResponse

	err = json.Unmarshal(data, &dst.CreateSubscriptionPaymentChargeResponse)

	if err == nil {

		jsonCreateSubscriptionPaymentChargeResponse, _ := json.Marshal(dst.CreateSubscriptionPaymentChargeResponse)

		if strings.Contains(string(jsonCreateSubscriptionPaymentChargeResponse), "{}") || strings.Contains(string(jsonCreateSubscriptionPaymentChargeResponse), "null") { // empty struct

			dst.CreateSubscriptionPaymentChargeResponse = nil

		} else {

			match++

		}

	} else {

		dst.CreateSubscriptionPaymentChargeResponse = nil

	}


	if match > 1 { // more than 1 match

		// reset to nil


		dst.CreateSubscriptionPaymentAuthResponse = nil


		dst.CreateSubscriptionPaymentChargeResponse = nil


		return fmt.Errorf("data matches more than one schema in oneOf(SubsCreatePayment200Response)")

	} else if match == 1 {

		return nil // exactly one match

	} else { // no match

		return fmt.Errorf("data failed to match schemas in oneOf(SubsCreatePayment200Response)")

	}



}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubsCreatePayment200Response) MarshalJSON() ([]byte, error) {
	if src.CreateSubscriptionPaymentAuthResponse != nil {
		return json.Marshal(&src.CreateSubscriptionPaymentAuthResponse)
	}

	if src.CreateSubscriptionPaymentChargeResponse != nil {
		return json.Marshal(&src.CreateSubscriptionPaymentChargeResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubsCreatePayment200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateSubscriptionPaymentAuthResponse != nil {
		return obj.CreateSubscriptionPaymentAuthResponse
	}

	if obj.CreateSubscriptionPaymentChargeResponse != nil {
		return obj.CreateSubscriptionPaymentChargeResponse
	}

	// all schemas are nil
	return nil
}



