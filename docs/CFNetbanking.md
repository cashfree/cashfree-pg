# CFNetbanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel for netbanking will always be &#x60;link&#x60; | 
**NetbankingBankCode** | **int32** | Bank code | 

## Methods

### NewCFNetbanking

`func NewCFNetbanking(channel string, netbankingBankCode int32, ) *CFNetbanking`

NewCFNetbanking instantiates a new CFNetbanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFNetbankingWithDefaults

`func NewCFNetbankingWithDefaults() *CFNetbanking`

NewCFNetbankingWithDefaults instantiates a new CFNetbanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFNetbanking) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFNetbanking) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFNetbanking) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetNetbankingBankCode

`func (o *CFNetbanking) GetNetbankingBankCode() int32`

GetNetbankingBankCode returns the NetbankingBankCode field if non-nil, zero value otherwise.

### GetNetbankingBankCodeOk

`func (o *CFNetbanking) GetNetbankingBankCodeOk() (*int32, bool)`

GetNetbankingBankCodeOk returns a tuple with the NetbankingBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankCode

`func (o *CFNetbanking) SetNetbankingBankCode(v int32)`

SetNetbankingBankCode sets NetbankingBankCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


