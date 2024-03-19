# BankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**AccountHolder** | Pointer to **string** |  | [optional] 
**Ifsc** | Pointer to **string** |  | [optional] 

## Methods

### NewBankDetails

`func NewBankDetails() *BankDetails`

NewBankDetails instantiates a new BankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankDetailsWithDefaults

`func NewBankDetailsWithDefaults() *BankDetails`

NewBankDetailsWithDefaults instantiates a new BankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *BankDetails) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankDetails) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankDetails) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *BankDetails) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountHolder

`func (o *BankDetails) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *BankDetails) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *BankDetails) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *BankDetails) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetIfsc

`func (o *BankDetails) GetIfsc() string`

GetIfsc returns the Ifsc field if non-nil, zero value otherwise.

### GetIfscOk

`func (o *BankDetails) GetIfscOk() (*string, bool)`

GetIfscOk returns a tuple with the Ifsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfsc

`func (o *BankDetails) SetIfsc(v string)`

SetIfsc sets Ifsc field to given value.

### HasIfsc

`func (o *BankDetails) HasIfsc() bool`

HasIfsc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


