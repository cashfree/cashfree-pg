# Banktransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel for cardless EMI is always &#x60;link&#x60; | [optional] 

## Methods

### NewBanktransfer

`func NewBanktransfer() *Banktransfer`

NewBanktransfer instantiates a new Banktransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanktransferWithDefaults

`func NewBanktransferWithDefaults() *Banktransfer`

NewBanktransferWithDefaults instantiates a new Banktransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Banktransfer) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Banktransfer) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Banktransfer) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Banktransfer) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


