# UpdateSoundboxVpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpa** | **string** | Terminal Vpa,for which we need to update details. | 
**CfTerminalId** | **string** | cashfree terminal id. | 
**MerchantName** | Pointer to **string** | Merchant Name that need to updated on soundbox | [optional] 
**Language** | Pointer to **string** | language of soundbox,currently English, Hindi, Tamil | [optional] 

## Methods

### NewUpdateSoundboxVpaRequest

`func NewUpdateSoundboxVpaRequest(vpa string, cfTerminalId string, ) *UpdateSoundboxVpaRequest`

NewUpdateSoundboxVpaRequest instantiates a new UpdateSoundboxVpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSoundboxVpaRequestWithDefaults

`func NewUpdateSoundboxVpaRequestWithDefaults() *UpdateSoundboxVpaRequest`

NewUpdateSoundboxVpaRequestWithDefaults instantiates a new UpdateSoundboxVpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpa

`func (o *UpdateSoundboxVpaRequest) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *UpdateSoundboxVpaRequest) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *UpdateSoundboxVpaRequest) SetVpa(v string)`

SetVpa sets Vpa field to given value.


### GetCfTerminalId

`func (o *UpdateSoundboxVpaRequest) GetCfTerminalId() string`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *UpdateSoundboxVpaRequest) GetCfTerminalIdOk() (*string, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *UpdateSoundboxVpaRequest) SetCfTerminalId(v string)`

SetCfTerminalId sets CfTerminalId field to given value.


### GetMerchantName

`func (o *UpdateSoundboxVpaRequest) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *UpdateSoundboxVpaRequest) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *UpdateSoundboxVpaRequest) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *UpdateSoundboxVpaRequest) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetLanguage

`func (o *UpdateSoundboxVpaRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateSoundboxVpaRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateSoundboxVpaRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateSoundboxVpaRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


