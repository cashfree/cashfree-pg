# ProductConditionsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The Action key in the conditions array specifies whether a condition is allowed or denied for the specified rule or feature | [optional] 
**Key** | Pointer to **string** | key of the condition | [optional] 
**Values** | Pointer to **[]string** | Values set for the condition | [optional] 

## Methods

### NewProductConditionsEntity

`func NewProductConditionsEntity() *ProductConditionsEntity`

NewProductConditionsEntity instantiates a new ProductConditionsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConditionsEntityWithDefaults

`func NewProductConditionsEntityWithDefaults() *ProductConditionsEntity`

NewProductConditionsEntityWithDefaults instantiates a new ProductConditionsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProductConditionsEntity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProductConditionsEntity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProductConditionsEntity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProductConditionsEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetKey

`func (o *ProductConditionsEntity) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProductConditionsEntity) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProductConditionsEntity) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProductConditionsEntity) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValues

`func (o *ProductConditionsEntity) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProductConditionsEntity) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProductConditionsEntity) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProductConditionsEntity) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


