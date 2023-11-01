/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
)

// checks if the PaymentMethodUPIInPaymentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodUPIInPaymentsEntity{}

// PaymentMethodUPIInPaymentsEntity UPI payment method for pay api
type PaymentMethodUPIInPaymentsEntity struct {
	Channel string `json:"channel"`
	UpiId *string `json:"upi_id,omitempty"`
}

// NewPaymentMethodUPIInPaymentsEntity instantiates a new PaymentMethodUPIInPaymentsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUPIInPaymentsEntity(channel string) *PaymentMethodUPIInPaymentsEntity {
	this := PaymentMethodUPIInPaymentsEntity{}
	this.Channel = channel
	return &this
}

// NewPaymentMethodUPIInPaymentsEntityWithDefaults instantiates a new PaymentMethodUPIInPaymentsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUPIInPaymentsEntityWithDefaults() *PaymentMethodUPIInPaymentsEntity {
	this := PaymentMethodUPIInPaymentsEntity{}
	return &this
}

// GetChannel returns the Channel field value
func (o *PaymentMethodUPIInPaymentsEntity) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodUPIInPaymentsEntity) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *PaymentMethodUPIInPaymentsEntity) SetChannel(v string) {
	o.Channel = v
}

// GetUpiId returns the UpiId field value if set, zero value otherwise.
func (o *PaymentMethodUPIInPaymentsEntity) GetUpiId() string {
	if o == nil || IsNil(o.UpiId) {
		var ret string
		return ret
	}
	return *o.UpiId
}

// GetUpiIdOk returns a tuple with the UpiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUPIInPaymentsEntity) GetUpiIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpiId) {
		return nil, false
	}
	return o.UpiId, true
}

// HasUpiId returns a boolean if a field has been set.
func (o *PaymentMethodUPIInPaymentsEntity) HasUpiId() bool {
	if o != nil && !IsNil(o.UpiId) {
		return true
	}

	return false
}

// SetUpiId gets a reference to the given string and assigns it to the UpiId field.
func (o *PaymentMethodUPIInPaymentsEntity) SetUpiId(v string) {
	o.UpiId = &v
}

func (o PaymentMethodUPIInPaymentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodUPIInPaymentsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	if !IsNil(o.UpiId) {
		toSerialize["upi_id"] = o.UpiId
	}
	return toSerialize, nil
}

type NullablePaymentMethodUPIInPaymentsEntity struct {
	value *PaymentMethodUPIInPaymentsEntity
	isSet bool
}

func (v NullablePaymentMethodUPIInPaymentsEntity) Get() *PaymentMethodUPIInPaymentsEntity {
	return v.value
}

func (v *NullablePaymentMethodUPIInPaymentsEntity) Set(val *PaymentMethodUPIInPaymentsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUPIInPaymentsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUPIInPaymentsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUPIInPaymentsEntity(val *PaymentMethodUPIInPaymentsEntity) *NullablePaymentMethodUPIInPaymentsEntity {
	return &NullablePaymentMethodUPIInPaymentsEntity{value: val, isSet: true}
}

func (v NullablePaymentMethodUPIInPaymentsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUPIInPaymentsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

