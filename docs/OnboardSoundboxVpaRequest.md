# OnboardSoundboxVpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpa** | **string** | Terminal Vpa ,that need to onboard on soundbox | 
**CfTerminalId** | **string** | cashfree terminal id. | 
**DeviceSerialNo** | **string** | Device Serial No of soundbox | 
**MerchantName** | Pointer to **string** | Merchant Name that need to onboard on soundbox | [optional] 
**Language** | Pointer to **string** | language of soundbox,currently English, Hindi, Tamil | [optional] 

## Methods

### NewOnboardSoundboxVpaRequest

`func NewOnboardSoundboxVpaRequest(vpa string, cfTerminalId string, deviceSerialNo string, ) *OnboardSoundboxVpaRequest`

NewOnboardSoundboxVpaRequest instantiates a new OnboardSoundboxVpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardSoundboxVpaRequestWithDefaults

`func NewOnboardSoundboxVpaRequestWithDefaults() *OnboardSoundboxVpaRequest`

NewOnboardSoundboxVpaRequestWithDefaults instantiates a new OnboardSoundboxVpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpa

`func (o *OnboardSoundboxVpaRequest) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *OnboardSoundboxVpaRequest) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *OnboardSoundboxVpaRequest) SetVpa(v string)`

SetVpa sets Vpa field to given value.


### GetCfTerminalId

`func (o *OnboardSoundboxVpaRequest) GetCfTerminalId() string`

GetCfTerminalId returns the CfTerminalId field if non-nil, zero value otherwise.

### GetCfTerminalIdOk

`func (o *OnboardSoundboxVpaRequest) GetCfTerminalIdOk() (*string, bool)`

GetCfTerminalIdOk returns a tuple with the CfTerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTerminalId

`func (o *OnboardSoundboxVpaRequest) SetCfTerminalId(v string)`

SetCfTerminalId sets CfTerminalId field to given value.


### GetDeviceSerialNo

`func (o *OnboardSoundboxVpaRequest) GetDeviceSerialNo() string`

GetDeviceSerialNo returns the DeviceSerialNo field if non-nil, zero value otherwise.

### GetDeviceSerialNoOk

`func (o *OnboardSoundboxVpaRequest) GetDeviceSerialNoOk() (*string, bool)`

GetDeviceSerialNoOk returns a tuple with the DeviceSerialNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerialNo

`func (o *OnboardSoundboxVpaRequest) SetDeviceSerialNo(v string)`

SetDeviceSerialNo sets DeviceSerialNo field to given value.


### GetMerchantName

`func (o *OnboardSoundboxVpaRequest) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *OnboardSoundboxVpaRequest) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *OnboardSoundboxVpaRequest) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *OnboardSoundboxVpaRequest) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetLanguage

`func (o *OnboardSoundboxVpaRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OnboardSoundboxVpaRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OnboardSoundboxVpaRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OnboardSoundboxVpaRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


