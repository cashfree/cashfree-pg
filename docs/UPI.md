# UPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpiId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** | Channel. can be link, qrcode, or collect | [optional] 

## Methods

### NewUPI

`func NewUPI() *UPI`

NewUPI instantiates a new UPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUPIWithDefaults

`func NewUPIWithDefaults() *UPI`

NewUPIWithDefaults instantiates a new UPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpiId

`func (o *UPI) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *UPI) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *UPI) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *UPI) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetChannel

`func (o *UPI) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UPI) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UPI) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UPI) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


