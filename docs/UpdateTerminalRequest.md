# UpdateTerminalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalEmail** | Pointer to **string** | Mention the updated email ID of the terminal. | [optional] 
**TerminalPhoneNo** | Pointer to **string** | Terminal phone number to be updated. | [optional] 
**TerminalMeta** | Pointer to [**UpdateTerminalRequestTerminalMeta**](UpdateTerminalRequestTerminalMeta.md) |  | [optional] 
**TerminalType** | **string** | Mention the terminal type to be updated. Possible values - AGENT, STOREFRONT. | 

## Methods

### NewUpdateTerminalRequest

`func NewUpdateTerminalRequest(terminalType string, ) *UpdateTerminalRequest`

NewUpdateTerminalRequest instantiates a new UpdateTerminalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTerminalRequestWithDefaults

`func NewUpdateTerminalRequestWithDefaults() *UpdateTerminalRequest`

NewUpdateTerminalRequestWithDefaults instantiates a new UpdateTerminalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerminalEmail

`func (o *UpdateTerminalRequest) GetTerminalEmail() string`

GetTerminalEmail returns the TerminalEmail field if non-nil, zero value otherwise.

### GetTerminalEmailOk

`func (o *UpdateTerminalRequest) GetTerminalEmailOk() (*string, bool)`

GetTerminalEmailOk returns a tuple with the TerminalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalEmail

`func (o *UpdateTerminalRequest) SetTerminalEmail(v string)`

SetTerminalEmail sets TerminalEmail field to given value.

### HasTerminalEmail

`func (o *UpdateTerminalRequest) HasTerminalEmail() bool`

HasTerminalEmail returns a boolean if a field has been set.

### GetTerminalPhoneNo

`func (o *UpdateTerminalRequest) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *UpdateTerminalRequest) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *UpdateTerminalRequest) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.

### HasTerminalPhoneNo

`func (o *UpdateTerminalRequest) HasTerminalPhoneNo() bool`

HasTerminalPhoneNo returns a boolean if a field has been set.

### GetTerminalMeta

`func (o *UpdateTerminalRequest) GetTerminalMeta() UpdateTerminalRequestTerminalMeta`

GetTerminalMeta returns the TerminalMeta field if non-nil, zero value otherwise.

### GetTerminalMetaOk

`func (o *UpdateTerminalRequest) GetTerminalMetaOk() (*UpdateTerminalRequestTerminalMeta, bool)`

GetTerminalMetaOk returns a tuple with the TerminalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalMeta

`func (o *UpdateTerminalRequest) SetTerminalMeta(v UpdateTerminalRequestTerminalMeta)`

SetTerminalMeta sets TerminalMeta field to given value.

### HasTerminalMeta

`func (o *UpdateTerminalRequest) HasTerminalMeta() bool`

HasTerminalMeta returns a boolean if a field has been set.

### GetTerminalType

`func (o *UpdateTerminalRequest) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *UpdateTerminalRequest) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *UpdateTerminalRequest) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


