# OrderMetaPaymentMethodsFiltersFiltersCardSchemes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | It accepts value of \&quot;ALLOW\&quot; and allows only those schemes present in it&#39;s neighbouring parameter \&quot;values\&quot; | [optional] 
**Values** | Pointer to **[]string** | List of card schemes to be allowed for the order | [optional] 

## Methods

### NewOrderMetaPaymentMethodsFiltersFiltersCardSchemes

`func NewOrderMetaPaymentMethodsFiltersFiltersCardSchemes() *OrderMetaPaymentMethodsFiltersFiltersCardSchemes`

NewOrderMetaPaymentMethodsFiltersFiltersCardSchemes instantiates a new OrderMetaPaymentMethodsFiltersFiltersCardSchemes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderMetaPaymentMethodsFiltersFiltersCardSchemesWithDefaults

`func NewOrderMetaPaymentMethodsFiltersFiltersCardSchemesWithDefaults() *OrderMetaPaymentMethodsFiltersFiltersCardSchemes`

NewOrderMetaPaymentMethodsFiltersFiltersCardSchemesWithDefaults instantiates a new OrderMetaPaymentMethodsFiltersFiltersCardSchemes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *OrderMetaPaymentMethodsFiltersFiltersCardSchemes) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


