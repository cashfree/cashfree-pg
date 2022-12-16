# CFUPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Specify the channel through which the payment must be processed. Can be one of [\&quot;link\&quot;, \&quot;collect\&quot;, \&quot;qrcode\&quot;] | 
**UpiId** | **string** | Customer UPI VPA to process payment. | 
**AuthorizeOnly** | Pointer to **bool** | For one time mandate on UPI. Set this as authorize_only &#x3D; true. Please note that you can only use the \&quot;collect\&quot; channel if you are sending a one time mandate request | [optional] 
**Authorization** | Pointer to [**CFUPIAuthorizeDetails**](CFUPIAuthorizeDetails.md) |  | [optional] 

## Methods

### NewCFUPI

`func NewCFUPI(channel string, upiId string, ) *CFUPI`

NewCFUPI instantiates a new CFUPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFUPIWithDefaults

`func NewCFUPIWithDefaults() *CFUPI`

NewCFUPIWithDefaults instantiates a new CFUPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFUPI) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFUPI) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFUPI) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetUpiId

`func (o *CFUPI) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *CFUPI) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *CFUPI) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.


### GetAuthorizeOnly

`func (o *CFUPI) GetAuthorizeOnly() bool`

GetAuthorizeOnly returns the AuthorizeOnly field if non-nil, zero value otherwise.

### GetAuthorizeOnlyOk

`func (o *CFUPI) GetAuthorizeOnlyOk() (*bool, bool)`

GetAuthorizeOnlyOk returns a tuple with the AuthorizeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeOnly

`func (o *CFUPI) SetAuthorizeOnly(v bool)`

SetAuthorizeOnly sets AuthorizeOnly field to given value.

### HasAuthorizeOnly

`func (o *CFUPI) HasAuthorizeOnly() bool`

HasAuthorizeOnly returns a boolean if a field has been set.

### GetAuthorization

`func (o *CFUPI) GetAuthorization() CFUPIAuthorizeDetails`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *CFUPI) GetAuthorizationOk() (*CFUPIAuthorizeDetails, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *CFUPI) SetAuthorization(v CFUPIAuthorizeDetails)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *CFUPI) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


