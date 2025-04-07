# TerminalEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedOn** | Pointer to **string** |  | [optional] 
**CfTerminalId** | Pointer to **int32** |  | [optional] 
**LastUpdatedOn** | Pointer to **string** |  | [optional] 
**TerminalAddress** | Pointer to **string** |  | [optional] 
**TerminalEmail** | Pointer to **string** |  | [optional] 
**TerminalType** | Pointer to **string** |  | [optional] 
**TeminalId** | Pointer to **string** |  | [optional] 
**TerminalName** | Pointer to **string** |  | [optional] 
**TerminalNote** | Pointer to **string** |  | [optional] 
**TerminalPhoneNo** | Pointer to **string** |  | [optional] 
**TerminalStatus** | Pointer to **string** |  | [optional] 
**TerminalMeta** | Pointer to [**TerminalEntityTerminalMeta**](TerminalEntityTerminalMeta.md) |  | [optional] 

## Methods

### NewTerminalEntity

`func NewTerminalEntity() *TerminalEntity`

NewTerminalEntity instantiates a new TerminalEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalEntityWithDefaults

`func NewTerminalEntityWithDefaults() *TerminalEntity`

NewTerminalEntityWithDefaults instantiates a new TerminalEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedOn

`func (o *TerminalEntity) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *TerminalEntity) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *TerminalEntity) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *TerminalEntity) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetCfTerminalId

`func (o *TerminalEntity) GetCfTerminalId() int32`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *TerminalEntity) GetCfTerminalIdOk() (*int32, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *TerminalEntity) SetCfTerminalId(v int32)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *TerminalEntity) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetLastUpdatedOn

`func (o *TerminalEntity) GetLastUpdatedOn() string`

GetLastUpdatedOn returns the LastUpdatedOn field if non-nil, zero value otherwise.

### GetLastUpdatedOnOk

`func (o *TerminalEntity) GetLastUpdatedOnOk() (*string, bool)`

GetLastUpdatedOnOk returns a tuple with the LastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedOn

`func (o *TerminalEntity) SetLastUpdatedOn(v string)`

SetLastUpdatedOn sets LastUpdatedOn field to given value.

### HasLastUpdatedOn

`func (o *TerminalEntity) HasLastUpdatedOn() bool`

HasLastUpdatedOn returns a boolean if a field has been set.

### GetTerminalAddress

`func (o *TerminalEntity) GetTerminalAddress() string`

GetTerminalAddress returns the TerminalAddress field if non-nil, zero value otherwise.

### GetTerminalAddressOk

`func (o *TerminalEntity) GetTerminalAddressOk() (*string, bool)`

GetTerminalAddressOk returns a tuple with the TerminalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAddress

`func (o *TerminalEntity) SetTerminalAddress(v string)`

SetTerminalAddress sets TerminalAddress field to given value.

### HasTerminalAddress

`func (o *TerminalEntity) HasTerminalAddress() bool`

HasTerminalAddress returns a boolean if a field has been set.

### GetTerminalEmail

`func (o *TerminalEntity) GetTerminalEmail() string`

GetTerminalEmail returns the TerminalEmail field if non-nil, zero value otherwise.

### GetTerminalEmailOk

`func (o *TerminalEntity) GetTerminalEmailOk() (*string, bool)`

GetTerminalEmailOk returns a tuple with the TerminalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalEmail

`func (o *TerminalEntity) SetTerminalEmail(v string)`

SetTerminalEmail sets TerminalEmail field to given value.

### HasTerminalEmail

`func (o *TerminalEntity) HasTerminalEmail() bool`

HasTerminalEmail returns a boolean if a field has been set.

### GetTerminalType

`func (o *TerminalEntity) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *TerminalEntity) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *TerminalEntity) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.

### HasTerminalType

`func (o *TerminalEntity) HasTerminalType() bool`

HasTerminalType returns a boolean if a field has been set.

### GetTeminalId

`func (o *TerminalEntity) GetTeminalId() string`

GetTeminalId returns the TeminalId field if non-nil, zero value otherwise.

### GetTeminalIdOk

`func (o *TerminalEntity) GetTeminalIdOk() (*string, bool)`

GetTeminalIdOk returns a tuple with the TeminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeminalId

`func (o *TerminalEntity) SetTeminalId(v string)`

SetTeminalId sets TeminalId field to given value.

### HasTeminalId

`func (o *TerminalEntity) HasTeminalId() bool`

HasTeminalId returns a boolean if a field has been set.

### GetTerminalName

`func (o *TerminalEntity) GetTerminalName() string`

GetTerminalName returns the TerminalName field if non-nil, zero value otherwise.

### GetTerminalNameOk

`func (o *TerminalEntity) GetTerminalNameOk() (*string, bool)`

GetTerminalNameOk returns a tuple with the TerminalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalName

`func (o *TerminalEntity) SetTerminalName(v string)`

SetTerminalName sets TerminalName field to given value.

### HasTerminalName

`func (o *TerminalEntity) HasTerminalName() bool`

HasTerminalName returns a boolean if a field has been set.

### GetTerminalNote

`func (o *TerminalEntity) GetTerminalNote() string`

GetTerminalNote returns the TerminalNote field if non-nil, zero value otherwise.

### GetTerminalNoteOk

`func (o *TerminalEntity) GetTerminalNoteOk() (*string, bool)`

GetTerminalNoteOk returns a tuple with the TerminalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalNote

`func (o *TerminalEntity) SetTerminalNote(v string)`

SetTerminalNote sets TerminalNote field to given value.

### HasTerminalNote

`func (o *TerminalEntity) HasTerminalNote() bool`

HasTerminalNote returns a boolean if a field has been set.

### GetTerminalPhoneNo

`func (o *TerminalEntity) GetTerminalPhoneNo() string`

GetTerminalPhoneNo returns the TerminalPhoneNo field if non-nil, zero value otherwise.

### GetTerminalPhoneNoOk

`func (o *TerminalEntity) GetTerminalPhoneNoOk() (*string, bool)`

GetTerminalPhoneNoOk returns a tuple with the TerminalPhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPhoneNo

`func (o *TerminalEntity) SetTerminalPhoneNo(v string)`

SetTerminalPhoneNo sets TerminalPhoneNo field to given value.

### HasTerminalPhoneNo

`func (o *TerminalEntity) HasTerminalPhoneNo() bool`

HasTerminalPhoneNo returns a boolean if a field has been set.

### GetTerminalStatus

`func (o *TerminalEntity) GetTerminalStatus() string`

GetTerminalStatus returns the TerminalStatus field if non-nil, zero value otherwise.

### GetTerminalStatusOk

`func (o *TerminalEntity) GetTerminalStatusOk() (*string, bool)`

GetTerminalStatusOk returns a tuple with the TerminalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalStatus

`func (o *TerminalEntity) SetTerminalStatus(v string)`

SetTerminalStatus sets TerminalStatus field to given value.

### HasTerminalStatus

`func (o *TerminalEntity) HasTerminalStatus() bool`

HasTerminalStatus returns a boolean if a field has been set.

### GetTerminalMeta

`func (o *TerminalEntity) GetTerminalMeta() TerminalEntityTerminalMeta`

GetTerminalMeta returns the TerminalMeta field if non-nil, zero value otherwise.

### GetTerminalMetaOk

`func (o *TerminalEntity) GetTerminalMetaOk() (*TerminalEntityTerminalMeta, bool)`

GetTerminalMetaOk returns a tuple with the TerminalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalMeta

`func (o *TerminalEntity) SetTerminalMeta(v TerminalEntityTerminalMeta)`

SetTerminalMeta sets TerminalMeta field to given value.

### HasTerminalMeta

`func (o *TerminalEntity) HasTerminalMeta() bool`

HasTerminalMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


