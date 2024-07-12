# CardlessEMI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel for cardless EMI is always &#x60;link&#x60; | [optional] 
**Provider** | Pointer to **string** | One of [&#x60;flexmoney&#x60;, &#x60;zestmoney&#x60;, &#x60;hdfc&#x60;, &#x60;icici&#x60;, &#x60;cashe&#x60;, &#x60;idfc&#x60;, &#x60;kotak&#x60;, &#x60;snapmint&#x60;, &#x60;bharatx&#x60;] | [optional] 
**Phone** | Pointer to **string** | Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as &#39;invalid_request_error&#39; and code as &#39;invalid_request_error&#39; | [optional] 
**EmiTenure** | Pointer to **int32** | EMI tenure for the selected provider. This is mandatory when provider is one of [&#x60;hdfc&#x60;, &#x60;icici&#x60;, &#x60;cashe&#x60;, &#x60;idfc&#x60;, &#x60;kotak&#x60;] | [optional] 

## Methods

### NewCardlessEMI

`func NewCardlessEMI() *CardlessEMI`

NewCardlessEMI instantiates a new CardlessEMI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardlessEMIWithDefaults

`func NewCardlessEMIWithDefaults() *CardlessEMI`

NewCardlessEMIWithDefaults instantiates a new CardlessEMI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CardlessEMI) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CardlessEMI) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CardlessEMI) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CardlessEMI) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetProvider

`func (o *CardlessEMI) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CardlessEMI) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CardlessEMI) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CardlessEMI) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPhone

`func (o *CardlessEMI) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CardlessEMI) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CardlessEMI) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CardlessEMI) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmiTenure

`func (o *CardlessEMI) GetEmiTenure() int32`

GetEmiTenure returns the EmiTenure field if non-nil, zero value otherwise.

### GetEmiTenureOk

`func (o *CardlessEMI) GetEmiTenureOk() (*int32, bool)`

GetEmiTenureOk returns a tuple with the EmiTenure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiTenure

`func (o *CardlessEMI) SetEmiTenure(v int32)`

SetEmiTenure sets EmiTenure field to given value.

### HasEmiTenure

`func (o *CardlessEMI) HasEmiTenure() bool`

HasEmiTenure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


