# CFPaylater

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for cardless EMI is always &#x60;link&#x60; | 
**Provider** | **string** | One of [\&quot;kotak\&quot;, \&quot;flexipay\&quot;, \&quot;zestmoney\&quot;, \&quot;lazypay\&quot;, \&quot;olapostpaid\&quot;]. Please note that Flexipay is offered by HDFC bank | 
**Phone** | **string** | Customers phone number for this payment instrument. If the customer is not eligible you will receive a 400 error with type as &#39;invalid_request_error&#39; and code as &#39;invalid_request_error&#39; | 

## Methods

### NewCFPaylater

`func NewCFPaylater(channel string, provider string, phone string, ) *CFPaylater`

NewCFPaylater instantiates a new CFPaylater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFPaylaterWithDefaults

`func NewCFPaylaterWithDefaults() *CFPaylater`

NewCFPaylaterWithDefaults instantiates a new CFPaylater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFPaylater) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFPaylater) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFPaylater) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetProvider

`func (o *CFPaylater) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CFPaylater) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CFPaylater) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPhone

`func (o *CFPaylater) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CFPaylater) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CFPaylater) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


