# CustomerDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | A unique identifier for the customer. Use alphanumeric values only. | [optional] 
**CustomerEmail** | Pointer to **string** | Customer email address. | [optional] 
**CustomerPhone** | Pointer to **string** | Customer phone number. | [optional] 
**CustomerName** | Pointer to **string** | Name of the customer. | [optional] 
**CustomerBankAccountNumber** | Pointer to **string** | Customer bank account. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankIfsc** | Pointer to **string** | Customer bank IFSC. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankCode** | Pointer to **float32** | Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV) | [optional] 
**CustomerUid** | Pointer to **string** | Customer identifier at Cashfree. You will get this when you create/get customer         | [optional] 

## Methods

### NewCustomerDetailsResponse

`func NewCustomerDetailsResponse() *CustomerDetailsResponse`

NewCustomerDetailsResponse instantiates a new CustomerDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDetailsResponseWithDefaults

`func NewCustomerDetailsResponseWithDefaults() *CustomerDetailsResponse`

NewCustomerDetailsResponseWithDefaults instantiates a new CustomerDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomerDetailsResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerDetailsResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerDetailsResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerDetailsResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *CustomerDetailsResponse) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CustomerDetailsResponse) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CustomerDetailsResponse) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CustomerDetailsResponse) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *CustomerDetailsResponse) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *CustomerDetailsResponse) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *CustomerDetailsResponse) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *CustomerDetailsResponse) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerName

`func (o *CustomerDetailsResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *CustomerDetailsResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *CustomerDetailsResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *CustomerDetailsResponse) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerBankAccountNumber

`func (o *CustomerDetailsResponse) GetCustomerBankAccountNumber() string`

GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field if non-nil, zero value otherwise.

### GetCustomerBankAccountNumberOk

`func (o *CustomerDetailsResponse) GetCustomerBankAccountNumberOk() (*string, bool)`

GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountNumber

`func (o *CustomerDetailsResponse) SetCustomerBankAccountNumber(v string)`

SetCustomerBankAccountNumber sets CustomerBankAccountNumber field to given value.

### HasCustomerBankAccountNumber

`func (o *CustomerDetailsResponse) HasCustomerBankAccountNumber() bool`

HasCustomerBankAccountNumber returns a boolean if a field has been set.

### GetCustomerBankIfsc

`func (o *CustomerDetailsResponse) GetCustomerBankIfsc() string`

GetCustomerBankIfsc returns the CustomerBankIfsc field if non-nil, zero value otherwise.

### GetCustomerBankIfscOk

`func (o *CustomerDetailsResponse) GetCustomerBankIfscOk() (*string, bool)`

GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankIfsc

`func (o *CustomerDetailsResponse) SetCustomerBankIfsc(v string)`

SetCustomerBankIfsc sets CustomerBankIfsc field to given value.

### HasCustomerBankIfsc

`func (o *CustomerDetailsResponse) HasCustomerBankIfsc() bool`

HasCustomerBankIfsc returns a boolean if a field has been set.

### GetCustomerBankCode

`func (o *CustomerDetailsResponse) GetCustomerBankCode() float32`

GetCustomerBankCode returns the CustomerBankCode field if non-nil, zero value otherwise.

### GetCustomerBankCodeOk

`func (o *CustomerDetailsResponse) GetCustomerBankCodeOk() (*float32, bool)`

GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankCode

`func (o *CustomerDetailsResponse) SetCustomerBankCode(v float32)`

SetCustomerBankCode sets CustomerBankCode field to given value.

### HasCustomerBankCode

`func (o *CustomerDetailsResponse) HasCustomerBankCode() bool`

HasCustomerBankCode returns a boolean if a field has been set.

### GetCustomerUid

`func (o *CustomerDetailsResponse) GetCustomerUid() string`

GetCustomerUid returns the CustomerUid field if non-nil, zero value otherwise.

### GetCustomerUidOk

`func (o *CustomerDetailsResponse) GetCustomerUidOk() (*string, bool)`

GetCustomerUidOk returns a tuple with the CustomerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUid

`func (o *CustomerDetailsResponse) SetCustomerUid(v string)`

SetCustomerUid sets CustomerUid field to given value.

### HasCustomerUid

`func (o *CustomerDetailsResponse) HasCustomerUid() bool`

HasCustomerUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


