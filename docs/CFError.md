# CFError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | One of [\&quot;invalid_request_error\&quot;, \&quot;authentication_error\&quot;, \&quot;rate_limit_error\&quot;, \&quot;validation_error\&quot;, \&quot;api_error\&quot;] | [optional] 

## Methods

### NewCFError

`func NewCFError() *CFError`

NewCFError instantiates a new CFError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFErrorWithDefaults

`func NewCFErrorWithDefaults() *CFError`

NewCFErrorWithDefaults instantiates a new CFError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CFError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CFError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CFError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CFError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *CFError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CFError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CFError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CFError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *CFError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CFError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CFError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CFError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


