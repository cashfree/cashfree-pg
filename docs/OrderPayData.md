# OrderPayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | For card payments, if the response includes &#x60;&#x60;&#x60;action:link&#x60;&#x60;&#x60;, redirect the customer to the &#x60;&#x60;&#x60;data.url&#x60;&#x60;&#x60; page. If the response includes &#x60;&#x60;&#x60;action:post&#x60;&#x60;&#x60;, display a native OTP UI to collect the OTP and submit it to &#x60;&#x60;&#x60;data.url&#x60;&#x60;&#x60;. | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Key value pairs sent as the request body if the payment link requires a form submission instead of a redirect. | [optional] 
**ContentType** | Pointer to **string** | Specifies the Content-Type header that should be used when submitting the payload, for example, &#x60;&#x60;&#x60;application/x-www-form-urlencoded&#x60;&#x60;&#x60;. | [optional] 
**Method** | Pointer to **string** | The HTTP method to use when submitting the payload to the url. For example, POST. | [optional] 
**RedirectToBank** | Pointer to **string** | This field is available only for card payments when &#x60;&#x60;&#x60;action:post&#x60;&#x60;&#x60; is returned. Display a native OTP UI and also provide an option for the customer to redirect to bank page. When the customer selects this option, redirect them to the &#x60;&#x60;&#x60;data.redirect_to_bank&#x60;&#x60;&#x60; URL. | [optional] 

## Methods

### NewOrderPayData

`func NewOrderPayData() *OrderPayData`

NewOrderPayData instantiates a new OrderPayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPayDataWithDefaults

`func NewOrderPayDataWithDefaults() *OrderPayData`

NewOrderPayDataWithDefaults instantiates a new OrderPayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrderPayData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrderPayData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrderPayData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrderPayData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPayload

`func (o *OrderPayData) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OrderPayData) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OrderPayData) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OrderPayData) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetContentType

`func (o *OrderPayData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *OrderPayData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *OrderPayData) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *OrderPayData) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetMethod

`func (o *OrderPayData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OrderPayData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OrderPayData) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OrderPayData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRedirectToBank

`func (o *OrderPayData) GetRedirectToBank() string`

GetRedirectToBank returns the RedirectToBank field if non-nil, zero value otherwise.

### GetRedirectToBankOk

`func (o *OrderPayData) GetRedirectToBankOk() (*string, bool)`

GetRedirectToBankOk returns a tuple with the RedirectToBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToBank

`func (o *OrderPayData) SetRedirectToBank(v string)`

SetRedirectToBank sets RedirectToBank field to given value.

### HasRedirectToBank

`func (o *OrderPayData) HasRedirectToBank() bool`

HasRedirectToBank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


