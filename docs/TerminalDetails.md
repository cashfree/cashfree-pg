# TerminalDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedOn** | Pointer to **string** | date time at which terminal is added | [optional] 
**CfTerminalId** | Pointer to **string** | Cashfree terminal id, this is a required parameter when you do not provide the terminal phone number. | [optional] 
**LastUpdatedOn** | Pointer to **string** | last instant when this terminal was updated | [optional] 
**TerminalAddress** | Pointer to **string** | location of terminal | [optional] 
**TerminalId** | Pointer to **string** | terminal id for merchant reference | [optional] 
**TerminalName** | Pointer to **string** | name of terminal/agent/storefront | [optional] 
**TerminalNote** | Pointer to **string** | note given by merchant while creating the terminal | [optional] 
**TerminalPhoneNo** | Pointer to **string** | mobile num of the terminal/agent/storefront,This is a required parameter when you do not provide the cf_terminal_id. | [optional] 
**TerminalStatus** | Pointer to **string** | status of terminal active/inactive | [optional] 
**TerminalType** | **string** | To identify the type of terminal product in use, in this case it is SPOS. | 

## Methods

### NewTerminalDetails

`func NewTerminalDetails(terminalType string, ) *TerminalDetails`

NewTerminalDetails instantiates a new TerminalDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalDetailsWithDefaults

`func NewTerminalDetailsWithDefaults() *TerminalDetails`

NewTerminalDetailsWithDefaults instantiates a new TerminalDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedOn

`func (o *TerminalDetails) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *TerminalDetails) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *TerminalDetails) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *TerminalDetails) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetCfTerminalId

`func (o *TerminalDetails) GetCfTerminalId() string`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *TerminalDetails) GetCfTerminalIdOk() (*string, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *TerminalDetails) SetCfTerminalId(v string)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *TerminalDetails) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *TerminalDetails) GetLastUpdatedOn() string`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *TerminalDetails) GetLastUpdatedOnOk() (*string, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *TerminalDetails) SetLastUpdatedOn(v string)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *TerminalDetails) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetTerminalAddress

`func (o *TerminalDetails) GetTerminalAddress() string`

GetTerminalAddress returns the TerminalAddress field if non-nil, zero value otherwise.

### GetTerminalAddressOk

`func (o *TerminalDetails) GetTerminalAddressOk() (*string, bool)`

GetTerminalAddressOk returns a tuple with the TerminalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAddress

`func (o *TerminalDetails) SetTerminalAddress(v string)`

SetTerminalAddress sets TerminalAddress field to given value.

### HasTerminalAddress

`func (o *TerminalDetails) HasTerminalAddress() bool`

HasTerminalAddress returns a boolean if a field has been set.

### GetTerminalId

`func (o *TerminalDetails) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *TerminalDetails) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *TerminalDetails) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *TerminalDetails) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalName

`func (o *TerminalDetails) GetTerminalName() string`

GetTerminalName returns the TerminalName field if non-nil, zero value otherwise.

### GetTerminalNameOk

`func (o *TerminalDetails) GetTerminalNameOk() (*string, bool)`

GetTerminalNameOk returns a tuple with the TerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalName

`func (o *TerminalDetails) SetTerminalName(v string)`

SetTerminalName sets TerminalName field to given value.

### HasTerminalName

`func (o *TerminalDetails) HasTerminalName() bool`

HasTerminalName returns a boolean if a field has been set.

### GetTerminalNote

`func (o *TerminalDetails) GetTerminalNote() string`

GetTerminalNote returns the TerminalNote field if non-nil, zero value otherwise.

### GetTerminalNoteOk

`func (o *TerminalDetails) GetTerminalNoteOk() (*string, bool)`

GetTerminalNoteOk returns a tuple with the TerminalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalNote

`func (o *TerminalDetails) SetTerminalNote(v string)`

SetTerminalNote sets TerminalNote field to given value.

### HasTerminalNote

`func (o *TerminalDetails) HasTerminalNote() bool`

HasTerminalNote returns a boolean if a field has been set.

### GetTerminalPhoneNo

`func (o *TerminalDetails) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *TerminalDetails) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *TerminalDetails) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.

### HasTerminalPhoneNo

`func (o *TerminalDetails) HasTerminalPhoneNo() bool`

HasTerminalPhoneNo returns a boolean if a field has been set.

### GetTerminalStatus

`func (o *TerminalDetails) GetTerminalStatus() string`

GetTerminalStatus returns the TerminalStatus field if non-nil, zero value otherwise.

### GetTerminalStatusOk

`func (o *TerminalDetails) GetTerminalStatusOk() (*string, bool)`

GetTerminalStatusOk returns a tuple with the TerminalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStatus

`func (o *TerminalDetails) SetTerminalStatus(v string)`

SetTerminalStatus sets TerminalStatus field to given value.

### HasTerminalStatus

`func (o *TerminalDetails) HasTerminalStatus() bool`

HasTerminalStatus returns a boolean if a field has been set.

### GetTerminalType

`func (o *TerminalDetails) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *TerminalDetails) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *TerminalDetails) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


