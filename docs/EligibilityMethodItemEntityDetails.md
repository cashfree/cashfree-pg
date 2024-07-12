# EligibilityMethodItemEntityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountTypes** | Pointer to **[]string** | List of account types associated with the payment method. (e.g. SAVINGS or CURRENT) | [optional] 
**FrequentBankDetails** | Pointer to [**[]SubscriptionBankDetails**](SubscriptionBankDetails.md) | List of the most frequently used banks. | [optional] 
**AllBankDetails** | Pointer to [**[]SubscriptionBankDetails**](SubscriptionBankDetails.md) | Details about all banks associated with the payment method. | [optional] 
**AvailableHandles** | Pointer to [**[]EligibilityMethodItemEntityDetailsAvailableHandlesInner**](EligibilityMethodItemEntityDetailsAvailableHandlesInner.md) | List of supported VPA handles. | [optional] 
**AllowedCardTypes** | Pointer to **[]string** | List of allowed card types. (e.g. DEBIT_CARD, CREDIT_CARD) | [optional] 

## Methods

### NewEligibilityMethodItemEntityDetails

`func NewEligibilityMethodItemEntityDetails() *EligibilityMethodItemEntityDetails`

NewEligibilityMethodItemEntityDetails instantiates a new EligibilityMethodItemEntityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityMethodItemEntityDetailsWithDefaults

`func NewEligibilityMethodItemEntityDetailsWithDefaults() *EligibilityMethodItemEntityDetails`

NewEligibilityMethodItemEntityDetailsWithDefaults instantiates a new EligibilityMethodItemEntityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountTypes

`func (o *EligibilityMethodItemEntityDetails) GetAccountTypes() []string`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *EligibilityMethodItemEntityDetails) GetAccountTypesOk() (*[]string, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *EligibilityMethodItemEntityDetails) SetAccountTypes(v []string)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *EligibilityMethodItemEntityDetails) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetFrequentBankDetails

`func (o *EligibilityMethodItemEntityDetails) GetFrequentBankDetails() []SubscriptionBankDetails`

GetFrequentBankDetails returns the FrequentBankDetails field if non-nil, zero value otherwise.

### GetFrequentBankDetailsOk

`func (o *EligibilityMethodItemEntityDetails) GetFrequentBankDetailsOk() (*[]SubscriptionBankDetails, bool)`

GetFrequentBankDetailsOk returns a tuple with the FrequentBankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentBankDetails

`func (o *EligibilityMethodItemEntityDetails) SetFrequentBankDetails(v []SubscriptionBankDetails)`

SetFrequentBankDetails sets FrequentBankDetails field to given value.

### HasFrequentBankDetails

`func (o *EligibilityMethodItemEntityDetails) HasFrequentBankDetails() bool`

HasFrequentBankDetails returns a boolean if a field has been set.

### GetAllBankDetails

`func (o *EligibilityMethodItemEntityDetails) GetAllBankDetails() []SubscriptionBankDetails`

GetAllBankDetails returns the AllBankDetails field if non-nil, zero value otherwise.

### GetAllBankDetailsOk

`func (o *EligibilityMethodItemEntityDetails) GetAllBankDetailsOk() (*[]SubscriptionBankDetails, bool)`

GetAllBankDetailsOk returns a tuple with the AllBankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBankDetails

`func (o *EligibilityMethodItemEntityDetails) SetAllBankDetails(v []SubscriptionBankDetails)`

SetAllBankDetails sets AllBankDetails field to given value.

### HasAllBankDetails

`func (o *EligibilityMethodItemEntityDetails) HasAllBankDetails() bool`

HasAllBankDetails returns a boolean if a field has been set.

### GetAvailableHandles

`func (o *EligibilityMethodItemEntityDetails) GetAvailableHandles() []EligibilityMethodItemEntityDetailsAvailableHandlesInner`

GetAvailableHandles returns the AvailableHandles field if non-nil, zero value otherwise.

### GetAvailableHandlesOk

`func (o *EligibilityMethodItemEntityDetails) GetAvailableHandlesOk() (*[]EligibilityMethodItemEntityDetailsAvailableHandlesInner, bool)`

GetAvailableHandlesOk returns a tuple with the AvailableHandles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableHandles

`func (o *EligibilityMethodItemEntityDetails) SetAvailableHandles(v []EligibilityMethodItemEntityDetailsAvailableHandlesInner)`

SetAvailableHandles sets AvailableHandles field to given value.

### HasAvailableHandles

`func (o *EligibilityMethodItemEntityDetails) HasAvailableHandles() bool`

HasAvailableHandles returns a boolean if a field has been set.

### GetAllowedCardTypes

`func (o *EligibilityMethodItemEntityDetails) GetAllowedCardTypes() []string`

GetAllowedCardTypes returns the AllowedCardTypes field if non-nil, zero value otherwise.

### GetAllowedCardTypesOk

`func (o *EligibilityMethodItemEntityDetails) GetAllowedCardTypesOk() (*[]string, bool)`

GetAllowedCardTypesOk returns a tuple with the AllowedCardTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCardTypes

`func (o *EligibilityMethodItemEntityDetails) SetAllowedCardTypes(v []string)`

SetAllowedCardTypes sets AllowedCardTypes field to given value.

### HasAllowedCardTypes

`func (o *EligibilityMethodItemEntityDetails) HasAllowedCardTypes() bool`

HasAllowedCardTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


