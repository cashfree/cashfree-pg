# AddressDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Full Name of the customer associated with the address. | [optional] 
**AddressLineOne** | Pointer to **string** | First line of the address. | [optional] 
**AddressLineTwo** | Pointer to **string** | Second line of the address. | [optional] 
**Country** | Pointer to **string** | Country Name. | [optional] 
**CountryCode** | Pointer to **string** | Country Code. | [optional] 
**State** | Pointer to **string** | State Name. | [optional] 
**StateCode** | Pointer to **string** | State Code. | [optional] 
**City** | Pointer to **string** | City Name. | [optional] 
**PinCode** | Pointer to **string** | Pin Code/Zip Code. | [optional] 
**Phone** | Pointer to **string** | Customer Phone Number. | [optional] 
**Email** | Pointer to **string** | Cutomer Email Address. | [optional] 

## Methods

### NewAddressDetails

`func NewAddressDetails() *AddressDetails`

NewAddressDetails instantiates a new AddressDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDetailsWithDefaults

`func NewAddressDetailsWithDefaults() *AddressDetails`

NewAddressDetailsWithDefaults instantiates a new AddressDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddressDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressLineOne

`func (o *AddressDetails) GetAddressLineOne() string`

GetAddressLineOne returns the AddressLineOne field if non-nil, zero value otherwise.

### GetAddressLineOneOk

`func (o *AddressDetails) GetAddressLineOneOk() (*string, bool)`

GetAddressLineOneOk returns a tuple with the AddressLineOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLineOne

`func (o *AddressDetails) SetAddressLineOne(v string)`

SetAddressLineOne sets AddressLineOne field to given value.

### HasAddressLineOne

`func (o *AddressDetails) HasAddressLineOne() bool`

HasAddressLineOne returns a boolean if a field has been set.

### GetAddressLineTwo

`func (o *AddressDetails) GetAddressLineTwo() string`

GetAddressLineTwo returns the AddressLineTwo field if non-nil, zero value otherwise.

### GetAddressLineTwoOk

`func (o *AddressDetails) GetAddressLineTwoOk() (*string, bool)`

GetAddressLineTwoOk returns a tuple with the AddressLineTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLineTwo

`func (o *AddressDetails) SetAddressLineTwo(v string)`

SetAddressLineTwo sets AddressLineTwo field to given value.

### HasAddressLineTwo

`func (o *AddressDetails) HasAddressLineTwo() bool`

HasAddressLineTwo returns a boolean if a field has been set.

### GetCountry

`func (o *AddressDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddressDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressDetails) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressDetails) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressDetails) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AddressDetails) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetState

`func (o *AddressDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateCode

`func (o *AddressDetails) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *AddressDetails) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *AddressDetails) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *AddressDetails) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCity

`func (o *AddressDetails) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressDetails) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressDetails) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressDetails) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPinCode

`func (o *AddressDetails) GetPinCode() string`

GetPinCode returns the PinCode field if non-nil, zero value otherwise.

### GetPinCodeOk

`func (o *AddressDetails) GetPinCodeOk() (*string, bool)`

GetPinCodeOk returns a tuple with the PinCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCode

`func (o *AddressDetails) SetPinCode(v string)`

SetPinCode sets PinCode field to given value.

### HasPinCode

`func (o *AddressDetails) HasPinCode() bool`

HasPinCode returns a boolean if a field has been set.

### GetPhone

`func (o *AddressDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressDetails) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *AddressDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


