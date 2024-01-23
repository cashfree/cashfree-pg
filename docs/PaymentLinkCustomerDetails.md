# PaymentLinkCustomerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | A unique identifier for the customer. Use alphanumeric values only. | [optional] 
**CustomerEmail** | Pointer to **string** | Customer email address. | [optional] 
**CustomerPhone** | **string** | Customer phone number. | 
**CustomerName** | Pointer to **string** | Name of the customer. | [optional] 
**CustomerBankAccountNumber** | Pointer to **string** | Customer bank account. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankIfsc** | Pointer to **string** | Customer bank IFSC. Required if you want to do a bank account check (TPV) | [optional] 
**CustomerBankCode** | Pointer to **float32** | Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV) | [optional] 

## Methods

### NewPaymentLinkCustomerDetails

`func NewPaymentLinkCustomerDetails(customerPhone string, ) *PaymentLinkCustomerDetails`

NewPaymentLinkCustomerDetails instantiates a new PaymentLinkCustomerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkCustomerDetailsWithDefaults

`func NewPaymentLinkCustomerDetailsWithDefaults() *PaymentLinkCustomerDetails`

NewPaymentLinkCustomerDetailsWithDefaults instantiates a new PaymentLinkCustomerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *PaymentLinkCustomerDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentLinkCustomerDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentLinkCustomerDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentLinkCustomerDetails) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *PaymentLinkCustomerDetails) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *PaymentLinkCustomerDetails) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *PaymentLinkCustomerDetails) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *PaymentLinkCustomerDetails) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *PaymentLinkCustomerDetails) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *PaymentLinkCustomerDetails) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *PaymentLinkCustomerDetails) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerName

`func (o *PaymentLinkCustomerDetails) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *PaymentLinkCustomerDetails) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *PaymentLinkCustomerDetails) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *PaymentLinkCustomerDetails) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerBankAccountNumber

`func (o *PaymentLinkCustomerDetails) GetCustomerBankAccountNumber() string`

GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field if non-nil, zero value otherwise.

### GetCustomerBankAccountNumberOk

`func (o *PaymentLinkCustomerDetails) GetCustomerBankAccountNumberOk() (*string, bool)`

GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankAccountNumber

`func (o *PaymentLinkCustomerDetails) SetCustomerBankAccountNumber(v string)`

SetCustomerBankAccountNumber sets CustomerBankAccountNumber field to given value.

### HasCustomerBankAccountNumber

`func (o *PaymentLinkCustomerDetails) HasCustomerBankAccountNumber() bool`

HasCustomerBankAccountNumber returns a boolean if a field has been set.

### GetCustomerBankIfsc

`func (o *PaymentLinkCustomerDetails) GetCustomerBankIfsc() string`

GetCustomerBankIfsc returns the CustomerBankIfsc field if non-nil, zero value otherwise.

### GetCustomerBankIfscOk

`func (o *PaymentLinkCustomerDetails) GetCustomerBankIfscOk() (*string, bool)`

GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankIfsc

`func (o *PaymentLinkCustomerDetails) SetCustomerBankIfsc(v string)`

SetCustomerBankIfsc sets CustomerBankIfsc field to given value.

### HasCustomerBankIfsc

`func (o *PaymentLinkCustomerDetails) HasCustomerBankIfsc() bool`

HasCustomerBankIfsc returns a boolean if a field has been set.

### GetCustomerBankCode

`func (o *PaymentLinkCustomerDetails) GetCustomerBankCode() float32`

GetCustomerBankCode returns the CustomerBankCode field if non-nil, zero value otherwise.

### GetCustomerBankCodeOk

`func (o *PaymentLinkCustomerDetails) GetCustomerBankCodeOk() (*float32, bool)`

GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBankCode

`func (o *PaymentLinkCustomerDetails) SetCustomerBankCode(v float32)`

SetCustomerBankCode sets CustomerBankCode field to given value.

### HasCustomerBankCode

`func (o *PaymentLinkCustomerDetails) HasCustomerBankCode() bool`

HasCustomerBankCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


