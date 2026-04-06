# OrderMetaOfferFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defines whether to allow or deny the specified offers. | [optional] 
**Values** | Pointer to **[]string** | List of offer identifiers. | [optional] 

## Methods

### NewOrderMetaOfferFilters

`func NewOrderMetaOfferFilters() *OrderMetaOfferFilters`

NewOrderMetaOfferFilters instantiates a new OrderMetaOfferFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderMetaOfferFiltersWithDefaults

`func NewOrderMetaOfferFiltersWithDefaults() *OrderMetaOfferFilters`

NewOrderMetaOfferFiltersWithDefaults instantiates a new OrderMetaOfferFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OrderMetaOfferFilters) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OrderMetaOfferFilters) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OrderMetaOfferFilters) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *OrderMetaOfferFilters) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValues

`func (o *OrderMetaOfferFilters) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OrderMetaOfferFilters) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OrderMetaOfferFilters) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *OrderMetaOfferFilters) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


