# ApiError502

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Help** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | &#x60;bank_processing_failure&#x60; will be returned here to denote failure at bank.  | [optional] 
**Type** | Pointer to **string** | api_error | [optional] 

## Methods

### NewApiError502

`func NewApiError502() *ApiError502`

NewApiError502 instantiates a new ApiError502 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiError502WithDefaults

`func NewApiError502WithDefaults() *ApiError502`

NewApiError502WithDefaults instantiates a new ApiError502 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiError502) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError502) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError502) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiError502) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHelp

`func (o *ApiError502) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ApiError502) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ApiError502) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ApiError502) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetCode

`func (o *ApiError502) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError502) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError502) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError502) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ApiError502) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiError502) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiError502) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiError502) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


