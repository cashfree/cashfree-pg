# UpdateTerminalStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalStatus** | **string** | Status of the terminal to be updated. possible values - ACTIVE, INACTIVE. | 

## Methods

### NewUpdateTerminalStatusRequest

`func NewUpdateTerminalStatusRequest(terminalStatus string, ) *UpdateTerminalStatusRequest`

NewUpdateTerminalStatusRequest instantiates a new UpdateTerminalStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTerminalStatusRequestWithDefaults

`func NewUpdateTerminalStatusRequestWithDefaults() *UpdateTerminalStatusRequest`

NewUpdateTerminalStatusRequestWithDefaults instantiates a new UpdateTerminalStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalStatus

`func (o *UpdateTerminalStatusRequest) GetTerminalStatus() string`

GetTerminalStatus returns the TerminalStatus field if non-nil, zero value otherwise.

### GetTerminalStatusOk

`func (o *UpdateTerminalStatusRequest) GetTerminalStatusOk() (*string, bool)`

GetTerminalStatusOk returns a tuple with the TerminalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStatus

`func (o *UpdateTerminalStatusRequest) SetTerminalStatus(v string)`

SetTerminalStatus sets TerminalStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


