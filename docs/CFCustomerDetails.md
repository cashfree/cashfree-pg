# CFCustomerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | A unique identifier for the customer. Use alphanumeric values only. | 
**CustomerEmail** | **string** | Customer email address. | 
**CustomerPhone** | **string** | Customer phone number. | 
**CustomerBankAccountNumber** | **string** | Customer bank account. Required if you want to do a bank account check (TPV) | 
**CustomerBankIfsc** | **string** | Customer bank IFSC. Required if you want to do a bank account check (TPV) | 
**CustomerBankCode** | **int32** | Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV) | 

## Methods

### NewCFCustomerDetails

`func NewCFCustomerDetails(customerId string, customerEmail string, customerPhone string, customerBankAccountNumber string, customerBankIfsc string, customerBankCode int32, ) *CFCustomerDetails`

NewCFCustomerDetails instantiates a new CFCustomerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCustomerDetailsWithDefaults

`func NewCFCustomerDetailsWithDefaults() *CFCustomerDetails`

NewCFCustomerDetailsWithDefaults instantiates a new CFCustomerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CFCustomerDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CFCustomerDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CFCustomerDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerEmail

`func (o *CFCustomerDetails) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CFCustomerDetails) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CFCustomerDetails) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### GetCustomerPhone

`func (o *CFCustomerDetails) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *CFCustomerDetails) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *CFCustomerDetails) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerBankAccountNumber

`func (o *CFCustomerDetails) GetCustomerBankAccountNumber() string`

GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field if non-nil, zero value otherwise.

### GetCustomerBankAccountNumberOk

`func (o *CFCustomerDetails) GetCustomerBankAccountNumberOk() (*string, bool)`

GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountNumber

`func (o *CFCustomerDetails) SetCustomerBankAccountNumber(v string)`

SetCustomerBankAccountNumber sets CustomerBankAccountNumber field to given value.


### GetCustomerBankIfsc

`func (o *CFCustomerDetails) GetCustomerBankIfsc() string`

GetCustomerBankIfsc returns the CustomerBankIfsc field if non-nil, zero value otherwise.

### GetCustomerBankIfscOk

`func (o *CFCustomerDetails) GetCustomerBankIfscOk() (*string, bool)`

GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankIfsc

`func (o *CFCustomerDetails) SetCustomerBankIfsc(v string)`

SetCustomerBankIfsc sets CustomerBankIfsc field to given value.


### GetCustomerBankCode

`func (o *CFCustomerDetails) GetCustomerBankCode() int32`

GetCustomerBankCode returns the CustomerBankCode field if non-nil, zero value otherwise.

### GetCustomerBankCodeOk

`func (o *CFCustomerDetails) GetCustomerBankCodeOk() (*int32, bool)`

GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankCode

`func (o *CFCustomerDetails) SetCustomerBankCode(v int32)`

SetCustomerBankCode sets CustomerBankCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


