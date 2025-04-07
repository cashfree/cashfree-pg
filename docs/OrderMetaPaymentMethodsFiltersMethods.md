# OrderMetaPaymentMethodsFiltersMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | It accepts value of \&quot;ALLOW\&quot; and allows only those modes present in it&#39;s neighbouring parameter \&quot;values\&quot; | [optional] 
**Values** | Pointer to **[]string** | The accepted entries for this paramter are \&quot;debit_card, credit_card, prepaid_card, upi, wallet, netbanking, banktransfer, paylater, paypal, debit_card_emi, credit_card_emi, upi_credit_card, upi_ppi, cardless_emi, account_based_payment, corporate_credit_card, sbc_debit_card, sbc_emandate, sbc_upi, sbc_credit_card\&quot; | [optional] 

## Methods

### NewOrderMetaPaymentMethodsFiltersMethods

`func NewOrderMetaPaymentMethodsFiltersMethods() *OrderMetaPaymentMethodsFiltersMethods`

NewOrderMetaPaymentMethodsFiltersMethods instantiates a new OrderMetaPaymentMethodsFiltersMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderMetaPaymentMethodsFiltersMethodsWithDefaults

`func NewOrderMetaPaymentMethodsFiltersMethodsWithDefaults() *OrderMetaPaymentMethodsFiltersMethods`

NewOrderMetaPaymentMethodsFiltersMethodsWithDefaults instantiates a new OrderMetaPaymentMethodsFiltersMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrderMetaPaymentMethodsFiltersMethods) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderMetaPaymentMethodsFiltersMethods) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderMetaPaymentMethodsFiltersMethods) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderMetaPaymentMethodsFiltersMethods) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValues

`func (o *OrderMetaPaymentMethodsFiltersMethods) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OrderMetaPaymentMethodsFiltersMethods) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OrderMetaPaymentMethodsFiltersMethods) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *OrderMetaPaymentMethodsFiltersMethods) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


