# SubscriptionBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | Pointer to **string** | ID of the bank. | [optional] 
**BankName** | Pointer to **string** | Name of the bank. | [optional] 
**AccountAuthModes** | Pointer to **[]string** | List of account authentication modes supported by the bank. (e.g. DEBIT_CARD, NET_BANKING, AADHAAR) | [optional] 

## Methods

### NewSubscriptionBankDetails

`func NewSubscriptionBankDetails() *SubscriptionBankDetails`

NewSubscriptionBankDetails instantiates a new SubscriptionBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionBankDetailsWithDefaults

`func NewSubscriptionBankDetailsWithDefaults() *SubscriptionBankDetails`

NewSubscriptionBankDetailsWithDefaults instantiates a new SubscriptionBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *SubscriptionBankDetails) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *SubscriptionBankDetails) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *SubscriptionBankDetails) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *SubscriptionBankDetails) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetBankName

`func (o *SubscriptionBankDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *SubscriptionBankDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *SubscriptionBankDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *SubscriptionBankDetails) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetAccountAuthModes

`func (o *SubscriptionBankDetails) GetAccountAuthModes() []string`

GetAccountAuthModes returns the AccountAuthModes field if non-nil, zero value otherwise.

### GetAccountAuthModesOk

`func (o *SubscriptionBankDetails) GetAccountAuthModesOk() (*[]string, bool)`

GetAccountAuthModesOk returns a tuple with the AccountAuthModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAuthModes

`func (o *SubscriptionBankDetails) SetAccountAuthModes(v []string)`

SetAccountAuthModes sets AccountAuthModes field to given value.

### HasAccountAuthModes

`func (o *SubscriptionBankDetails) HasAccountAuthModes() bool`

HasAccountAuthModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


