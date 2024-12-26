# StaticQrResponseEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfTerminalId** | Pointer to **int32** | cashfree terminal id | [optional] 
**Vpa** | Pointer to **string** | Virtual Address | [optional] 
**Status** | Pointer to **string** | Status of vpa | [optional] 
**QrCode** | Pointer to **string** | qrcode | [optional] 

## Methods

### NewStaticQrResponseEntity

`func NewStaticQrResponseEntity() *StaticQrResponseEntity`

NewStaticQrResponseEntity instantiates a new StaticQrResponseEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticQrResponseEntityWithDefaults

`func NewStaticQrResponseEntityWithDefaults() *StaticQrResponseEntity`

NewStaticQrResponseEntityWithDefaults instantiates a new StaticQrResponseEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfTerminalId

`func (o *StaticQrResponseEntity) GetCfTerminalId() int32`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *StaticQrResponseEntity) GetCfTerminalIdOk() (*int32, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *StaticQrResponseEntity) SetCfTerminalId(v int32)`

SetCfTerminalId sets CfTerminalId field to given value.

### HasCfTerminalId

`func (o *StaticQrResponseEntity) HasCfTerminalId() bool`

HasCfTerminalId returns a boolean if a field has been set.

### GetVpa

`func (o *StaticQrResponseEntity) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *StaticQrResponseEntity) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *StaticQrResponseEntity) SetVpa(v string)`

SetVpa sets Vpa field to given value.

### HasVpa

`func (o *StaticQrResponseEntity) HasVpa() bool`

HasVpa returns a boolean if a field has been set.

### GetStatus

`func (o *StaticQrResponseEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StaticQrResponseEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StaticQrResponseEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StaticQrResponseEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQrCode

`func (o *StaticQrResponseEntity) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *StaticQrResponseEntity) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *StaticQrResponseEntity) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *StaticQrResponseEntity) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


