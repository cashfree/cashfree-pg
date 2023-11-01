# CreateTerminalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalId** | **string** | merchant’s internal terminal id | 
**TerminalPhoneNo** | **string** | phone number assigned to the terminal | 
**TerminalName** | **string** | terminal name to be assigned by merchants | 
**TerminalAddress** | Pointer to **string** | address of the terminal. required for STOREFRONT | [optional] 
**TerminalEmail** | **string** | terminal email ID of the AGENT/STOREFRONT assigned by merchants. | 
**TerminalNote** | Pointer to **string** | additional note for terminal | [optional] 
**TerminalType** | **string** | mention the terminal type. possible values - AGENT, STOREFRONT. | 
**TerminalMeta** | Pointer to [**CreateTerminalRequestTerminalMeta**](CreateTerminalRequestTerminalMeta.md) |  | [optional] 

## Methods

### NewCreateTerminalRequest

`func NewCreateTerminalRequest(terminalId string, terminalPhoneNo string, terminalName string, terminalEmail string, terminalType string, ) *CreateTerminalRequest`

NewCreateTerminalRequest instantiates a new CreateTerminalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTerminalRequestWithDefaults

`func NewCreateTerminalRequestWithDefaults() *CreateTerminalRequest`

NewCreateTerminalRequestWithDefaults instantiates a new CreateTerminalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalId

`func (o *CreateTerminalRequest) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CreateTerminalRequest) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CreateTerminalRequest) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.


### GetTerminalPhoneNo

`func (o *CreateTerminalRequest) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *CreateTerminalRequest) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *CreateTerminalRequest) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.


### GetTerminalName

`func (o *CreateTerminalRequest) GetTerminalName() string`

GetTerminalName returns the TerminalName field if non-nil, zero value otherwise.

### GetTerminalNameOk

`func (o *CreateTerminalRequest) GetTerminalNameOk() (*string, bool)`

GetTerminalNameOk returns a tuple with the TerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalName

`func (o *CreateTerminalRequest) SetTerminalName(v string)`

SetTerminalName sets TerminalName field to given value.


### GetTerminalAddress

`func (o *CreateTerminalRequest) GetTerminalAddress() string`

GetTerminalAddress returns the TerminalAddress field if non-nil, zero value otherwise.

### GetTerminalAddressOk

`func (o *CreateTerminalRequest) GetTerminalAddressOk() (*string, bool)`

GetTerminalAddressOk returns a tuple with the TerminalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAddress

`func (o *CreateTerminalRequest) SetTerminalAddress(v string)`

SetTerminalAddress sets TerminalAddress field to given value.

### HasTerminalAddress

`func (o *CreateTerminalRequest) HasTerminalAddress() bool`

HasTerminalAddress returns a boolean if a field has been set.

### GetTerminalEmail

`func (o *CreateTerminalRequest) GetTerminalEmail() string`

GetTerminalEmail returns the TerminalEmail field if non-nil, zero value otherwise.

### GetTerminalEmailOk

`func (o *CreateTerminalRequest) GetTerminalEmailOk() (*string, bool)`

GetTerminalEmailOk returns a tuple with the TerminalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalEmail

`func (o *CreateTerminalRequest) SetTerminalEmail(v string)`

SetTerminalEmail sets TerminalEmail field to given value.


### GetTerminalNote

`func (o *CreateTerminalRequest) GetTerminalNote() string`

GetTerminalNote returns the TerminalNote field if non-nil, zero value otherwise.

### GetTerminalNoteOk

`func (o *CreateTerminalRequest) GetTerminalNoteOk() (*string, bool)`

GetTerminalNoteOk returns a tuple with the TerminalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalNote

`func (o *CreateTerminalRequest) SetTerminalNote(v string)`

SetTerminalNote sets TerminalNote field to given value.

### HasTerminalNote

`func (o *CreateTerminalRequest) HasTerminalNote() bool`

HasTerminalNote returns a boolean if a field has been set.

### GetTerminalType

`func (o *CreateTerminalRequest) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *CreateTerminalRequest) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *CreateTerminalRequest) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.


### GetTerminalMeta

`func (o *CreateTerminalRequest) GetTerminalMeta() CreateTerminalRequestTerminalMeta`

GetTerminalMeta returns the TerminalMeta field if non-nil, zero value otherwise.

### GetTerminalMetaOk

`func (o *CreateTerminalRequest) GetTerminalMetaOk() (*CreateTerminalRequestTerminalMeta, bool)`

GetTerminalMetaOk returns a tuple with the TerminalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalMeta

`func (o *CreateTerminalRequest) SetTerminalMeta(v CreateTerminalRequestTerminalMeta)`

SetTerminalMeta sets TerminalMeta field to given value.

### HasTerminalMeta

`func (o *CreateTerminalRequest) HasTerminalMeta() bool`

HasTerminalMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


