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

// checks if the CreateOrderRequestOrderMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrderRequestOrderMeta{}

// CreateOrderRequestOrderMeta struct for CreateOrderRequestOrderMeta
type CreateOrderRequestOrderMeta struct {
	// The URL to which user will be redirected to after the payment on bank OTP page. Maximum length: 250. The return_url must contain placeholder {order_id}. When redirecting the customer back to the return url from the bank’s OTP page, Cashfree will replace this placeholder with the actual value for that order.
	ReturnUrl NullableString `json:"return_url,omitempty"`
	// Notification URL for server-server communication. Useful when user's connection drops while re-directing. NotifyUrl should be an https URL. Maximum length: 250.
	NotifyUrl NullableString `json:"notify_url,omitempty"`
	// Allowed payment modes for this order. Pass comma-separated values among following options - \"cc\", \"dc\", \"ccc\", \"ppc\",\"nb\",\"upi\",\"paypal\",\"app\",\"paylater\",\"cardlessemi\",\"dcemi\",\"ccemi\",\"banktransfer\". Leave it blank to show all available payment methods
	PaymentMethods interface{} `json:"payment_methods,omitempty"`
}

// NewCreateOrderRequestOrderMeta instantiates a new CreateOrderRequestOrderMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequestOrderMeta() *CreateOrderRequestOrderMeta {
	this := CreateOrderRequestOrderMeta{}
	return &this
}

// NewCreateOrderRequestOrderMetaWithDefaults instantiates a new CreateOrderRequestOrderMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestOrderMetaWithDefaults() *CreateOrderRequestOrderMeta {
	this := CreateOrderRequestOrderMeta{}
	return &this
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrderRequestOrderMeta) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ReturnUrl.Get()
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrderRequestOrderMeta) GetReturnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnUrl.Get(), o.ReturnUrl.IsSet()
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *CreateOrderRequestOrderMeta) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given NullableString and assigns it to the ReturnUrl field.
func (o *CreateOrderRequestOrderMeta) SetReturnUrl(v string) {
	o.ReturnUrl.Set(&v)
}
// SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil
func (o *CreateOrderRequestOrderMeta) SetReturnUrlNil() {
	o.ReturnUrl.Set(nil)
}

// UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
func (o *CreateOrderRequestOrderMeta) UnsetReturnUrl() {
	o.ReturnUrl.Unset()
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrderRequestOrderMeta) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NotifyUrl.Get()
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrderRequestOrderMeta) GetNotifyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyUrl.Get(), o.NotifyUrl.IsSet()
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *CreateOrderRequestOrderMeta) HasNotifyUrl() bool {
	if o != nil && o.NotifyUrl.IsSet() {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given NullableString and assigns it to the NotifyUrl field.
func (o *CreateOrderRequestOrderMeta) SetNotifyUrl(v string) {
	o.NotifyUrl.Set(&v)
}
// SetNotifyUrlNil sets the value for NotifyUrl to be an explicit nil
func (o *CreateOrderRequestOrderMeta) SetNotifyUrlNil() {
	o.NotifyUrl.Set(nil)
}

// UnsetNotifyUrl ensures that no value is present for NotifyUrl, not even an explicit nil
func (o *CreateOrderRequestOrderMeta) UnsetNotifyUrl() {
	o.NotifyUrl.Unset()
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrderRequestOrderMeta) GetPaymentMethods() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrderRequestOrderMeta) GetPaymentMethodsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return &o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *CreateOrderRequestOrderMeta) HasPaymentMethods() bool {
	if o != nil && IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given interface{} and assigns it to the PaymentMethods field.
func (o *CreateOrderRequestOrderMeta) SetPaymentMethods(v interface{}) {
	o.PaymentMethods = v
}

func (o CreateOrderRequestOrderMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderRequestOrderMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReturnUrl.IsSet() {
		toSerialize["return_url"] = o.ReturnUrl.Get()
	}
	if o.NotifyUrl.IsSet() {
		toSerialize["notify_url"] = o.NotifyUrl.Get()
	}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return toSerialize, nil
}

type NullableCreateOrderRequestOrderMeta struct {
	value *CreateOrderRequestOrderMeta
	isSet bool
}

func (v NullableCreateOrderRequestOrderMeta) Get() *CreateOrderRequestOrderMeta {
	return v.value
}

func (v *NullableCreateOrderRequestOrderMeta) Set(val *CreateOrderRequestOrderMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequestOrderMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequestOrderMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequestOrderMeta(val *CreateOrderRequestOrderMeta) *NullableCreateOrderRequestOrderMeta {
	return &NullableCreateOrderRequestOrderMeta{value: val, isSet: true}
}

func (v NullableCreateOrderRequestOrderMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequestOrderMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


