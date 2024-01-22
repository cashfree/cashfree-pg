# Upi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Specify the channel through which the payment must be processed. Can be one of [\&quot;link\&quot;, \&quot;collect\&quot;, \&quot;qrcode\&quot;] | 
**UpiId** | Pointer to **string** | Customer UPI VPA to process payment.  ### Important This is a required parameter for channel &#x3D; &#x60;collect&#x60;  | [optional] 
**UpiRedirectUrl** | Pointer to **bool** | use this if you want cashfree to show a loader. Sample response below. It is only supported for collect &#x60;action:collect&#x60; will be returned with &#x60;data.url&#x60; having the link for redirection  | [optional] 
**UpiExpiryMinutes** | Pointer to **float32** | The UPI request will be valid for this expiry minutes. This parameter is only applicable for a UPI collect payment. The default value is 5 minutes. You should keep the minimum as 5 minutes, and maximum as 15 minutes | [optional] 
**AuthorizeOnly** | Pointer to **bool** | For one time mandate on UPI. Set this as authorize_only &#x3D; true. Please note that you can only use the \&quot;collect\&quot; channel if you are sending a one time mandate request | [optional] 
**Authorization** | Pointer to [**UPIAuthorizeDetails**](UPIAuthorizeDetails.md) |  | [optional] 

## Methods

### NewUpi

`func NewUpi(channel string, ) *Upi`

NewUpi instantiates a new Upi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpiWithDefaults

`func NewUpiWithDefaults() *Upi`

NewUpiWithDefaults instantiates a new Upi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Upi) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Upi) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Upi) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetUpiId

`func (o *Upi) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *Upi) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *Upi) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *Upi) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetUpiRedirectUrl

`func (o *Upi) GetUpiRedirectUrl() bool`

GetUpiRedirectUrl returns the UpiRedirectUrl field if non-nil, zero value otherwise.

### GetUpiRedirectUrlOk

`func (o *Upi) GetUpiRedirectUrlOk() (*bool, bool)`

GetUpiRedirectUrlOk returns a tuple with the UpiRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiRedirectUrl

`func (o *Upi) SetUpiRedirectUrl(v bool)`

SetUpiRedirectUrl sets UpiRedirectUrl field to given value.

### HasUpiRedirectUrl

`func (o *Upi) HasUpiRedirectUrl() bool`

HasUpiRedirectUrl returns a boolean if a field has been set.

### GetUpiExpiryMinutes

`func (o *Upi) GetUpiExpiryMinutes() float32`

GetUpiExpiryMinutes returns the UpiExpiryMinutes field if non-nil, zero value otherwise.

### GetUpiExpiryMinutesOk

`func (o *Upi) GetUpiExpiryMinutesOk() (*float32, bool)`

GetUpiExpiryMinutesOk returns a tuple with the UpiExpiryMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiExpiryMinutes

`func (o *Upi) SetUpiExpiryMinutes(v float32)`

SetUpiExpiryMinutes sets UpiExpiryMinutes field to given value.

### HasUpiExpiryMinutes

`func (o *Upi) HasUpiExpiryMinutes() bool`

HasUpiExpiryMinutes returns a boolean if a field has been set.

### GetAuthorizeOnly

`func (o *Upi) GetAuthorizeOnly() bool`

GetAuthorizeOnly returns the AuthorizeOnly field if non-nil, zero value otherwise.

### GetAuthorizeOnlyOk

`func (o *Upi) GetAuthorizeOnlyOk() (*bool, bool)`

GetAuthorizeOnlyOk returns a tuple with the AuthorizeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeOnly

`func (o *Upi) SetAuthorizeOnly(v bool)`

SetAuthorizeOnly sets AuthorizeOnly field to given value.

### HasAuthorizeOnly

`func (o *Upi) HasAuthorizeOnly() bool`

HasAuthorizeOnly returns a boolean if a field has been set.

### GetAuthorization

`func (o *Upi) GetAuthorization() UPIAuthorizeDetails`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *Upi) GetAuthorizationOk() (*UPIAuthorizeDetails, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *Upi) SetAuthorization(v UPIAuthorizeDetails)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *Upi) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


