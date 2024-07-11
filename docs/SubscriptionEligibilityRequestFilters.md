# SubscriptionEligibilityRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to **[]string** | Possbile values in array - enach, pnach, upi, card. | [optional] 

## Methods

### NewSubscriptionEligibilityRequestFilters

`func NewSubscriptionEligibilityRequestFilters() *SubscriptionEligibilityRequestFilters`

NewSubscriptionEligibilityRequestFilters instantiates a new SubscriptionEligibilityRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEligibilityRequestFiltersWithDefaults

`func NewSubscriptionEligibilityRequestFiltersWithDefaults() *SubscriptionEligibilityRequestFilters`

NewSubscriptionEligibilityRequestFiltersWithDefaults instantiates a new SubscriptionEligibilityRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *SubscriptionEligibilityRequestFilters) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *SubscriptionEligibilityRequestFilters) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *SubscriptionEligibilityRequestFilters) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *SubscriptionEligibilityRequestFilters) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


