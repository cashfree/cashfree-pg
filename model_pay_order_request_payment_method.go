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
	"fmt"
)

// PayOrderRequestPaymentMethod - struct for PayOrderRequestPaymentMethod
type PayOrderRequestPaymentMethod struct {
	AppPaymentMethod *AppPaymentMethod
	BanktransferPaymentMethod *BanktransferPaymentMethod
	CardEMIPaymentMethod *CardEMIPaymentMethod
	CardPaymentMethod *CardPaymentMethod
	CardlessEMIPaymentMethod *CardlessEMIPaymentMethod
	NetBankingPaymentMethod *NetBankingPaymentMethod
	PaylaterPaymentMethod *PaylaterPaymentMethod
	UPIPaymentMethod *UPIPaymentMethod
}

// AppPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns AppPaymentMethod wrapped in PayOrderRequestPaymentMethod
func AppPaymentMethodAsPayOrderRequestPaymentMethod(v *AppPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		AppPaymentMethod: v,
	}
}

// BanktransferPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns BanktransferPaymentMethod wrapped in PayOrderRequestPaymentMethod
func BanktransferPaymentMethodAsPayOrderRequestPaymentMethod(v *BanktransferPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		BanktransferPaymentMethod: v,
	}
}

// CardEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardEMIPaymentMethod: v,
	}
}

// CardPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardPaymentMethodAsPayOrderRequestPaymentMethod(v *CardPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardPaymentMethod: v,
	}
}

// CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardlessEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardlessEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardlessEMIPaymentMethod: v,
	}
}

// NetBankingPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns NetBankingPaymentMethod wrapped in PayOrderRequestPaymentMethod
func NetBankingPaymentMethodAsPayOrderRequestPaymentMethod(v *NetBankingPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		NetBankingPaymentMethod: v,
	}
}

// PaylaterPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns PaylaterPaymentMethod wrapped in PayOrderRequestPaymentMethod
func PaylaterPaymentMethodAsPayOrderRequestPaymentMethod(v *PaylaterPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		PaylaterPaymentMethod: v,
	}
}

// UPIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns UPIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func UPIPaymentMethodAsPayOrderRequestPaymentMethod(v *UPIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		UPIPaymentMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PayOrderRequestPaymentMethod) UnmarshalJSON(data []byte) error {
		var err error





	match := 0


	// try to unmarshal data into AppPaymentMethod

	err = json.Unmarshal(data, &dst.AppPaymentMethod)

	if err == nil {

		jsonAppPaymentMethod, _ := json.Marshal(dst.AppPaymentMethod)

		if strings.Contains(string(jsonAppPaymentMethod), "{}") || strings.Contains(string(jsonAppPaymentMethod), "null") { // empty struct

			dst.AppPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.AppPaymentMethod = nil

	}


	// try to unmarshal data into BanktransferPaymentMethod

	err = json.Unmarshal(data, &dst.BanktransferPaymentMethod)

	if err == nil {

		jsonBanktransferPaymentMethod, _ := json.Marshal(dst.BanktransferPaymentMethod)

		if strings.Contains(string(jsonBanktransferPaymentMethod), "{}") || strings.Contains(string(jsonBanktransferPaymentMethod), "null") { // empty struct

			dst.BanktransferPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.BanktransferPaymentMethod = nil

	}


	// try to unmarshal data into CardEMIPaymentMethod

	err = json.Unmarshal(data, &dst.CardEMIPaymentMethod)

	if err == nil {

		jsonCardEMIPaymentMethod, _ := json.Marshal(dst.CardEMIPaymentMethod)

		if strings.Contains(string(jsonCardEMIPaymentMethod), "{}") || strings.Contains(string(jsonCardEMIPaymentMethod), "null") { // empty struct

			dst.CardEMIPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.CardEMIPaymentMethod = nil

	}


	// try to unmarshal data into CardPaymentMethod

	err = json.Unmarshal(data, &dst.CardPaymentMethod)

	if err == nil {

		jsonCardPaymentMethod, _ := json.Marshal(dst.CardPaymentMethod)

		if strings.Contains(string(jsonCardPaymentMethod), "{}") || strings.Contains(string(jsonCardPaymentMethod), "null") { // empty struct

			dst.CardPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.CardPaymentMethod = nil

	}


	// try to unmarshal data into CardlessEMIPaymentMethod

	err = json.Unmarshal(data, &dst.CardlessEMIPaymentMethod)

	if err == nil {

		jsonCardlessEMIPaymentMethod, _ := json.Marshal(dst.CardlessEMIPaymentMethod)

		if strings.Contains(string(jsonCardlessEMIPaymentMethod), "{}") || strings.Contains(string(jsonCardlessEMIPaymentMethod), "null") { // empty struct

			dst.CardlessEMIPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.CardlessEMIPaymentMethod = nil

	}


	// try to unmarshal data into NetBankingPaymentMethod

	err = json.Unmarshal(data, &dst.NetBankingPaymentMethod)

	if err == nil {

		jsonNetBankingPaymentMethod, _ := json.Marshal(dst.NetBankingPaymentMethod)

		if strings.Contains(string(jsonNetBankingPaymentMethod), "{}") || strings.Contains(string(jsonNetBankingPaymentMethod), "null") { // empty struct

			dst.NetBankingPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.NetBankingPaymentMethod = nil

	}


	// try to unmarshal data into PaylaterPaymentMethod

	err = json.Unmarshal(data, &dst.PaylaterPaymentMethod)

	if err == nil {

		jsonPaylaterPaymentMethod, _ := json.Marshal(dst.PaylaterPaymentMethod)

		if strings.Contains(string(jsonPaylaterPaymentMethod), "{}") || strings.Contains(string(jsonPaylaterPaymentMethod), "null") { // empty struct

			dst.PaylaterPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.PaylaterPaymentMethod = nil

	}


	// try to unmarshal data into UPIPaymentMethod

	err = json.Unmarshal(data, &dst.UPIPaymentMethod)

	if err == nil {

		jsonUPIPaymentMethod, _ := json.Marshal(dst.UPIPaymentMethod)

		if strings.Contains(string(jsonUPIPaymentMethod), "{}") || strings.Contains(string(jsonUPIPaymentMethod), "null") { // empty struct

			dst.UPIPaymentMethod = nil

		} else {

			match++

		}

	} else {

		dst.UPIPaymentMethod = nil

	}


	if match > 1 { // more than 1 match

		// reset to nil


		dst.AppPaymentMethod = nil


		dst.BanktransferPaymentMethod = nil


		dst.CardEMIPaymentMethod = nil


		dst.CardPaymentMethod = nil


		dst.CardlessEMIPaymentMethod = nil


		dst.NetBankingPaymentMethod = nil


		dst.PaylaterPaymentMethod = nil


		dst.UPIPaymentMethod = nil


		return fmt.Errorf("data matches more than one schema in oneOf(PayOrderRequestPaymentMethod)")

	} else if match == 1 {

		return nil // exactly one match

	} else { // no match

		return fmt.Errorf("data failed to match schemas in oneOf(PayOrderRequestPaymentMethod)")

	}



}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PayOrderRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.AppPaymentMethod != nil {
		return json.Marshal(&src.AppPaymentMethod)
	}

	if src.BanktransferPaymentMethod != nil {
		return json.Marshal(&src.BanktransferPaymentMethod)
	}

	if src.CardEMIPaymentMethod != nil {
		return json.Marshal(&src.CardEMIPaymentMethod)
	}

	if src.CardPaymentMethod != nil {
		return json.Marshal(&src.CardPaymentMethod)
	}

	if src.CardlessEMIPaymentMethod != nil {
		return json.Marshal(&src.CardlessEMIPaymentMethod)
	}

	if src.NetBankingPaymentMethod != nil {
		return json.Marshal(&src.NetBankingPaymentMethod)
	}

	if src.PaylaterPaymentMethod != nil {
		return json.Marshal(&src.PaylaterPaymentMethod)
	}

	if src.UPIPaymentMethod != nil {
		return json.Marshal(&src.UPIPaymentMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PayOrderRequestPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppPaymentMethod != nil {
		return obj.AppPaymentMethod
	}

	if obj.BanktransferPaymentMethod != nil {
		return obj.BanktransferPaymentMethod
	}

	if obj.CardEMIPaymentMethod != nil {
		return obj.CardEMIPaymentMethod
	}

	if obj.CardPaymentMethod != nil {
		return obj.CardPaymentMethod
	}

	if obj.CardlessEMIPaymentMethod != nil {
		return obj.CardlessEMIPaymentMethod
	}

	if obj.NetBankingPaymentMethod != nil {
		return obj.NetBankingPaymentMethod
	}

	if obj.PaylaterPaymentMethod != nil {
		return obj.PaylaterPaymentMethod
	}

	if obj.UPIPaymentMethod != nil {
		return obj.UPIPaymentMethod
	}

	// all schemas are nil
	return nil
}



