# CreateOrderRequestTerminal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedOn** | Pointer to **string** | date time at which terminal is added | [optional] 
**CfTerminalId** | Pointer to **int32** | cashfree terminal id | [optional] 
**LastUpdatedOn** | Pointer to **string** | last instant when this terminal was updated | [optional] 
**TerminalAddress** | Pointer to **string** | location of terminal | [optional] 
**TerminalId** | **string** | terminal id for merchant reference | 
**TerminalName** | Pointer to **string** | name of terminal/agent/storefront | [optional] 
**TerminalNote** | Pointer to **string** | note given by merchant while creating the terminal | [optional] 
**TerminalPhoneNo** | **string** | mobile num of the terminal/agent/storefront | 
**TerminalStatus** | Pointer to **string** | status of terminal active/inactive | [optional] 
**TerminalType** | **string** | To identify the type of terminal product in use, in this case it is SPOS. | 

## Methods

### NewCreateOrderRequestTerminal

`func NewCreateOrderRequestTerminal(terminalId string, terminalPhoneNo string, terminalType string, ) *CreateOrderRequestTerminal`

NewCreateOrderRequestTerminal instantiates a new CreateOrderRequestTerminal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestTerminalWithDefaults

`func NewCreateOrderRequestTerminalWithDefaults() *CreateOrderRequestTerminal`

NewCreateOrderRequestTerminalWithDefaults instantiates a new CreateOrderRequestTerminal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedOn

`func (o *CreateOrderRequestTerminal) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *CreateOrderRequestTerminal) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *CreateOrderRequestTerminal) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *CreateOrderRequestTerminal) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetCfTerminalId

`func (o *CreateOrderRequestTerminal) GetCfTerminalId() int32`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *CreateOrderRequestTerminal) GetCfTerminalIdOk() (*int32, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *CreateOrderRequestTerminal) SetCfTerminalId(v int32)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *CreateOrderRequestTerminal) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *CreateOrderRequestTerminal) GetLastUpdatedOn() string`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *CreateOrderRequestTerminal) GetLastUpdatedOnOk() (*string, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *CreateOrderRequestTerminal) SetLastUpdatedOn(v string)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *CreateOrderRequestTerminal) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetTerminalAddress

`func (o *CreateOrderRequestTerminal) GetTerminalAddress() string`

GetTerminalAddress returns the TerminalAddress field if non-nil, zero value otherwise.

### GetTerminalAddressOk

`func (o *CreateOrderRequestTerminal) GetTerminalAddressOk() (*string, bool)`

GetTerminalAddressOk returns a tuple with the TerminalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAddress

`func (o *CreateOrderRequestTerminal) SetTerminalAddress(v string)`

SetTerminalAddress sets TerminalAddress field to given value.

### HasTerminalAddress

`func (o *CreateOrderRequestTerminal) HasTerminalAddress() bool`

HasTerminalAddress returns a boolean if a field has been set.

### GetTerminalId

`func (o *CreateOrderRequestTerminal) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CreateOrderRequestTerminal) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CreateOrderRequestTerminal) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.


### GetTerminalName

`func (o *CreateOrderRequestTerminal) GetTerminalName() string`

GetTerminalName returns the TerminalName field if non-nil, zero value otherwise.

### GetTerminalNameOk

`func (o *CreateOrderRequestTerminal) GetTerminalNameOk() (*string, bool)`

GetTerminalNameOk returns a tuple with the TerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalName

`func (o *CreateOrderRequestTerminal) SetTerminalName(v string)`

SetTerminalName sets TerminalName field to given value.

### HasTerminalName

`func (o *CreateOrderRequestTerminal) HasTerminalName() bool`

HasTerminalName returns a boolean if a field has been set.

### GetTerminalNote

`func (o *CreateOrderRequestTerminal) GetTerminalNote() string`

GetTerminalNote returns the TerminalNote field if non-nil, zero value otherwise.

### GetTerminalNoteOk

`func (o *CreateOrderRequestTerminal) GetTerminalNoteOk() (*string, bool)`

GetTerminalNoteOk returns a tuple with the TerminalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalNote

`func (o *CreateOrderRequestTerminal) SetTerminalNote(v string)`

SetTerminalNote sets TerminalNote field to given value.

### HasTerminalNote

`func (o *CreateOrderRequestTerminal) HasTerminalNote() bool`

HasTerminalNote returns a boolean if a field has been set.

### GetTerminalPhoneNo

`func (o *CreateOrderRequestTerminal) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *CreateOrderRequestTerminal) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *CreateOrderRequestTerminal) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.


### GetTerminalStatus

`func (o *CreateOrderRequestTerminal) GetTerminalStatus() string`

GetTerminalStatus returns the TerminalStatus field if non-nil, zero value otherwise.

### GetTerminalStatusOk

`func (o *CreateOrderRequestTerminal) GetTerminalStatusOk() (*string, bool)`

GetTerminalStatusOk returns a tuple with the TerminalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStatus

`func (o *CreateOrderRequestTerminal) SetTerminalStatus(v string)`

SetTerminalStatus sets TerminalStatus field to given value.

### HasTerminalStatus

`func (o *CreateOrderRequestTerminal) HasTerminalStatus() bool`

HasTerminalStatus returns a boolean if a field has been set.

### GetTerminalType

`func (o *CreateOrderRequestTerminal) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *CreateOrderRequestTerminal) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *CreateOrderRequestTerminal) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


