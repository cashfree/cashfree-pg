# CustomerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | A unique identifier for the customer. Use alphanumeric values only. | 
**CustomerEmail** | Pointer to **NullableString** | Customer email address. | [optional] 
**CustomerPhone** | **string** | Customer phone number. | 
**CustomerName** | Pointer to **NullableString** | Name of the customer. | [optional] 
**CustomerBankAccountNumber** | Pointer to **NullableString** | Customer bank account. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankIfsc** | Pointer to **NullableString** | Customer bank IFSC. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankCode** | Pointer to **NullableFloat32** | Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV) | [optional] 

## Methods

### NewCustomerDetails

`func NewCustomerDetails(customerId string, customerPhone string, ) *CustomerDetails`

NewCustomerDetails instantiates a new CustomerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDetailsWithDefaults

`func NewCustomerDetailsWithDefaults() *CustomerDetails`

NewCustomerDetailsWithDefaults instantiates a new CustomerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomerDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerEmail

`func (o *CustomerDetails) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CustomerDetails) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CustomerDetails) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CustomerDetails) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *CustomerDetails) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *CustomerDetails) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerPhone

`func (o *CustomerDetails) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *CustomerDetails) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *CustomerDetails) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerName

`func (o *CustomerDetails) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *CustomerDetails) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *CustomerDetails) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *CustomerDetails) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### SetCustomerNameNil

`func (o *CustomerDetails) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *CustomerDetails) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetCustomerBankAccountNumber

`func (o *CustomerDetails) GetCustomerBankAccountNumber() string`

GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field if non-nil, zero value otherwise.

### GetCustomerBankAccountNumberOk

`func (o *CustomerDetails) GetCustomerBankAccountNumberOk() (*string, bool)`

GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountNumber

`func (o *CustomerDetails) SetCustomerBankAccountNumber(v string)`

SetCustomerBankAccountNumber sets CustomerBankAccountNumber field to given value.

### HasCustomerBankAccountNumber

`func (o *CustomerDetails) HasCustomerBankAccountNumber() bool`

HasCustomerBankAccountNumber returns a boolean if a field has been set.

### SetCustomerBankAccountNumberNil

`func (o *CustomerDetails) SetCustomerBankAccountNumberNil(b bool)`

 SetCustomerBankAccountNumberNil sets the value for CustomerBankAccountNumber to be an explicit nil

### UnsetCustomerBankAccountNumber
`func (o *CustomerDetails) UnsetCustomerBankAccountNumber()`

UnsetCustomerBankAccountNumber ensures that no value is present for CustomerBankAccountNumber, not even an explicit nil
### GetCustomerBankIfsc

`func (o *CustomerDetails) GetCustomerBankIfsc() string`

GetCustomerBankIfsc returns the CustomerBankIfsc field if non-nil, zero value otherwise.

### GetCustomerBankIfscOk

`func (o *CustomerDetails) GetCustomerBankIfscOk() (*string, bool)`

GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankIfsc

`func (o *CustomerDetails) SetCustomerBankIfsc(v string)`

SetCustomerBankIfsc sets CustomerBankIfsc field to given value.

### HasCustomerBankIfsc

`func (o *CustomerDetails) HasCustomerBankIfsc() bool`

HasCustomerBankIfsc returns a boolean if a field has been set.

### SetCustomerBankIfscNil

`func (o *CustomerDetails) SetCustomerBankIfscNil(b bool)`

 SetCustomerBankIfscNil sets the value for CustomerBankIfsc to be an explicit nil

### UnsetCustomerBankIfsc
`func (o *CustomerDetails) UnsetCustomerBankIfsc()`

UnsetCustomerBankIfsc ensures that no value is present for CustomerBankIfsc, not even an explicit nil
### GetCustomerBankCode

`func (o *CustomerDetails) GetCustomerBankCode() float32`

GetCustomerBankCode returns the CustomerBankCode field if non-nil, zero value otherwise.

### GetCustomerBankCodeOk

`func (o *CustomerDetails) GetCustomerBankCodeOk() (*float32, bool)`

GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankCode

`func (o *CustomerDetails) SetCustomerBankCode(v float32)`

SetCustomerBankCode sets CustomerBankCode field to given value.

### HasCustomerBankCode

`func (o *CustomerDetails) HasCustomerBankCode() bool`

HasCustomerBankCode returns a boolean if a field has been set.

### SetCustomerBankCodeNil

`func (o *CustomerDetails) SetCustomerBankCodeNil(b bool)`

 SetCustomerBankCodeNil sets the value for CustomerBankCode to be an explicit nil

### UnsetCustomerBankCode
`func (o *CustomerDetails) UnsetCustomerBankCode()`

UnsetCustomerBankCode ensures that no value is present for CustomerBankCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


