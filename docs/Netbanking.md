# Netbanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for netbanking will always be &#x60;link&#x60; | 
**NetbankingBankCode** | Pointer to **int32** | Bank code | [optional] 
**NetbankingBankName** | Pointer to **string** | String code for bank | [optional] 

## Methods

### NewNetbanking

`func NewNetbanking(channel string, ) *Netbanking`

NewNetbanking instantiates a new Netbanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetbankingWithDefaults

`func NewNetbankingWithDefaults() *Netbanking`

NewNetbankingWithDefaults instantiates a new Netbanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *Netbanking) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Netbanking) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Netbanking) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetNetbankingBankCode

`func (o *Netbanking) GetNetbankingBankCode() int32`

GetNetbankingBankCode returns the NetbankingBankCode field if non-nil, zero value otherwise.

### GetNetbankingBankCodeOk

`func (o *Netbanking) GetNetbankingBankCodeOk() (*int32, bool)`

GetNetbankingBankCodeOk returns a tuple with the NetbankingBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankCode

`func (o *Netbanking) SetNetbankingBankCode(v int32)`

SetNetbankingBankCode sets NetbankingBankCode field to given value.

### HasNetbankingBankCode

`func (o *Netbanking) HasNetbankingBankCode() bool`

HasNetbankingBankCode returns a boolean if a field has been set.

### GetNetbankingBankName

`func (o *Netbanking) GetNetbankingBankName() string`

GetNetbankingBankName returns the NetbankingBankName field if non-nil, zero value otherwise.

### GetNetbankingBankNameOk

`func (o *Netbanking) GetNetbankingBankNameOk() (*string, bool)`

GetNetbankingBankNameOk returns a tuple with the NetbankingBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankName

`func (o *Netbanking) SetNetbankingBankName(v string)`

SetNetbankingBankName sets NetbankingBankName field to given value.

### HasNetbankingBankName

`func (o *Netbanking) HasNetbankingBankName() bool`

HasNetbankingBankName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


