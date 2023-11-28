# PaymentMethodsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to **[]string** | Array of payment methods to be filtered. This is optional, by default all payment methods will be returned. Possible values in [ &#39;debit_card&#39;, &#39;credit_card&#39;, &#39;prepaid_card&#39;, &#39;corporate_credit_card&#39;, &#39;upi&#39;, &#39;wallet&#39;, &#39;netbanking&#39;, &#39;banktransfer&#39;, &#39;paylater&#39;, &#39;paypal&#39;, &#39;debit_card_emi&#39;, &#39;credit_card_emi&#39;, &#39;upi_credit_card&#39;, &#39;upi_ppi&#39;, &#39;cardless_emi&#39;, &#39;account_based_payment&#39; ]  | [optional] 

## Methods

### NewPaymentMethodsFilters

`func NewPaymentMethodsFilters() *PaymentMethodsFilters`

NewPaymentMethodsFilters instantiates a new PaymentMethodsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodsFiltersWithDefaults

`func NewPaymentMethodsFiltersWithDefaults() *PaymentMethodsFilters`

NewPaymentMethodsFiltersWithDefaults instantiates a new PaymentMethodsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *PaymentMethodsFilters) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentMethodsFilters) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentMethodsFilters) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentMethodsFilters) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


