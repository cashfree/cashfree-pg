# Paylater

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel for cardless EMI is always &#x60;link&#x60; | [optional] 
**Provider** | Pointer to **string** | One of [\&quot;kotak\&quot;, \&quot;flexipay\&quot;, \&quot;zestmoney\&quot;, \&quot;lazypay\&quot;, \&quot;olapostpaid\&quot;,\&quot;simpl\&quot;, \&quot;freechargepaylater\&quot;]. Please note that Flexipay is offered by HDFC bank | [optional] 
**Phone** | Pointer to **string** | Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as &#39;invalid_request_error&#39; and code as &#39;invalid_request_error&#39; | [optional] 

## Methods

### NewPaylater

`func NewPaylater() *Paylater`

NewPaylater instantiates a new Paylater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylaterWithDefaults

`func NewPaylaterWithDefaults() *Paylater`

NewPaylaterWithDefaults instantiates a new Paylater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Paylater) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Paylater) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Paylater) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Paylater) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetProvider

`func (o *Paylater) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Paylater) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Paylater) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Paylater) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPhone

`func (o *Paylater) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Paylater) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Paylater) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Paylater) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


