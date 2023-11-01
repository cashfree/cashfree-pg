# ApiError409

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | invalid_request_error | [optional] 

## Methods

### NewApiError409

`func NewApiError409() *ApiError409`

NewApiError409 instantiates a new ApiError409 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiError409WithDefaults

`func NewApiError409WithDefaults() *ApiError409`

NewApiError409WithDefaults instantiates a new ApiError409 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiError409) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError409) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError409) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiError409) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *ApiError409) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError409) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError409) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError409) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ApiError409) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiError409) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiError409) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiError409) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


