# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Help** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | api_error | [optional] 

## Methods

### NewApiError

`func NewApiError() *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *ApiError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHelp

`func (o *ApiError) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ApiError) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ApiError) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ApiError) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetType

`func (o *ApiError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


