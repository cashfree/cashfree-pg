# VBASimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VbaIfsc** | **string** | IFSC of the issues VBA. | 
**Utr** | **string** | UTR of the transaction. | 
**Amount** | **float32** | Amount of the transaction. | 
**RemitterAccount** | **string** | Account from which transaction was done. | 
**RemitterIfsc** | **string** | IFSC of the remitter. | 
**RemitterName** | **string** | Name of the remitter. | 
**Phone** | **string** | Phone number of the remitter. | 

## Methods

### NewVBASimulationRequest

`func NewVBASimulationRequest(vbaIfsc string, utr string, amount float32, remitterAccount string, remitterIfsc string, remitterName string, phone string, ) *VBASimulationRequest`

NewVBASimulationRequest instantiates a new VBASimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVBASimulationRequestWithDefaults

`func NewVBASimulationRequestWithDefaults() *VBASimulationRequest`

NewVBASimulationRequestWithDefaults instantiates a new VBASimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVbaIfsc

`func (o *VBASimulationRequest) GetVbaIfsc() string`

GetVbaIfsc returns the VbaIfsc field if non-nil, zero value otherwise.

### GetVbaIfscOk

`func (o *VBASimulationRequest) GetVbaIfscOk() (*string, bool)`

GetVbaIfscOk returns a tuple with the VbaIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbaIfsc

`func (o *VBASimulationRequest) SetVbaIfsc(v string)`

SetVbaIfsc sets VbaIfsc field to given value.


### GetUtr

`func (o *VBASimulationRequest) GetUtr() string`

GetUtr returns the Utr field if non-nil, zero value otherwise.

### GetUtrOk

`func (o *VBASimulationRequest) GetUtrOk() (*string, bool)`

GetUtrOk returns a tuple with the Utr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtr

`func (o *VBASimulationRequest) SetUtr(v string)`

SetUtr sets Utr field to given value.


### GetAmount

`func (o *VBASimulationRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VBASimulationRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VBASimulationRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetRemitterAccount

`func (o *VBASimulationRequest) GetRemitterAccount() string`

GetRemitterAccount returns the RemitterAccount field if non-nil, zero value otherwise.

### GetRemitterAccountOk

`func (o *VBASimulationRequest) GetRemitterAccountOk() (*string, bool)`

GetRemitterAccountOk returns a tuple with the RemitterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitterAccount

`func (o *VBASimulationRequest) SetRemitterAccount(v string)`

SetRemitterAccount sets RemitterAccount field to given value.


### GetRemitterIfsc

`func (o *VBASimulationRequest) GetRemitterIfsc() string`

GetRemitterIfsc returns the RemitterIfsc field if non-nil, zero value otherwise.

### GetRemitterIfscOk

`func (o *VBASimulationRequest) GetRemitterIfscOk() (*string, bool)`

GetRemitterIfscOk returns a tuple with the RemitterIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitterIfsc

`func (o *VBASimulationRequest) SetRemitterIfsc(v string)`

SetRemitterIfsc sets RemitterIfsc field to given value.


### GetRemitterName

`func (o *VBASimulationRequest) GetRemitterName() string`

GetRemitterName returns the RemitterName field if non-nil, zero value otherwise.

### GetRemitterNameOk

`func (o *VBASimulationRequest) GetRemitterNameOk() (*string, bool)`

GetRemitterNameOk returns a tuple with the RemitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitterName

`func (o *VBASimulationRequest) SetRemitterName(v string)`

SetRemitterName sets RemitterName field to given value.


### GetPhone

`func (o *VBASimulationRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *VBASimulationRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *VBASimulationRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


