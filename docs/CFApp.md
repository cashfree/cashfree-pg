# CFApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel has to be set to link | 
**Provider** | **string** | Specify the channel through which the payment must be processed. | 
**Phone** | **string** | Customer phone number associated with a wallet for payment. | 

## Methods

### NewCFApp

`func NewCFApp(channel string, provider string, phone string, ) *CFApp`

NewCFApp instantiates a new CFApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFAppWithDefaults

`func NewCFAppWithDefaults() *CFApp`

NewCFAppWithDefaults instantiates a new CFApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFApp) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFApp) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFApp) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetProvider

`func (o *CFApp) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CFApp) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CFApp) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPhone

`func (o *CFApp) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CFApp) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CFApp) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


