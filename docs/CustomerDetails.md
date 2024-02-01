# CustomerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | A unique identifier for the customer. Use alphanumeric values only. | 
**CustomerEmail** | Pointer to **string** | Customer email address. | [optional] 
**CustomerPhone** | **string** | Customer phone number. | 
**CustomerName** | Pointer to **string** | Name of the customer. | [optional] 
**CustomerBankAccountNumber** | Pointer to **string** | Customer bank account. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankIfsc** | Pointer to **string** | Customer bank IFSC. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankCode** | Pointer to **float32** | Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV) | [optional] 
**CustomerUid** | Pointer to **string** | Customer identifier at Cashfree. You will get this when you create/get customer | [optional] 

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

### GetCustomerUid

`func (o *CustomerDetails) GetCustomerUid() string`

GetCustomerUid returns the CustomerUid field if non-nil, zero value otherwise.

### GetCustomerUidOk

`func (o *CustomerDetails) GetCustomerUidOk() (*string, bool)`

GetCustomerUidOk returns a tuple with the CustomerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUid

`func (o *CustomerDetails) SetCustomerUid(v string)`

SetCustomerUid sets CustomerUid field to given value.

### HasCustomerUid

`func (o *CustomerDetails) HasCustomerUid() bool`

HasCustomerUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


