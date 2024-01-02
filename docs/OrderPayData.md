# OrderPayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


