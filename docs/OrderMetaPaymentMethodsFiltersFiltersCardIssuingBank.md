# OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | It accepts value of \&quot;ALLOW\&quot; and allows only those issuing bank present in it&#39;s neighbouring parameter \&quot;values\&quot; | [optional] 
**Values** | Pointer to **[]string** | List of card issuing bank to be allowed for the order | [optional] 

## Methods

### NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBank

`func NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBank() *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank`

NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBank instantiates a new OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBankWithDefaults

`func NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBankWithDefaults() *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank`

NewOrderMetaPaymentMethodsFiltersFiltersCardIssuingBankWithDefaults instantiates a new OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardIssuingBank) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


