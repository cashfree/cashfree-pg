# AuthenticationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | authentication_error | [optional] 

## Methods

### NewAuthenticationError

`func NewAuthenticationError() *AuthenticationError`

NewAuthenticationError instantiates a new AuthenticationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationErrorWithDefaults

`func NewAuthenticationErrorWithDefaults() *AuthenticationError`

NewAuthenticationErrorWithDefaults instantiates a new AuthenticationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AuthenticationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthenticationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthenticationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthenticationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *AuthenticationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthenticationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthenticationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthenticationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *AuthenticationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticationError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


