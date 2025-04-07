# ProductConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The Action key in the conditions array specifies whether a condition should \&quot;ALLOW\&quot; or \&quot;DENY\&quot; the specified rule or feature | [optional] 
**Key** | Pointer to **string** | Specify what you&#39;re trying to configure, such as \&quot;features\&quot; | [optional] 
**Values** | Pointer to **[]string** | Define the values you need to set within the conditions in this array, such as \&quot;checkoutCollectAddress\&quot;, \&quot;checkoutAuthenticate\&quot; | [optional] 

## Methods

### NewProductConditions

`func NewProductConditions() *ProductConditions`

NewProductConditions instantiates a new ProductConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConditionsWithDefaults

`func NewProductConditionsWithDefaults() *ProductConditions`

NewProductConditionsWithDefaults instantiates a new ProductConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProductConditions) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProductConditions) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProductConditions) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ProductConditions) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetKey

`func (o *ProductConditions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProductConditions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProductConditions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProductConditions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValues

`func (o *ProductConditions) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProductConditions) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProductConditions) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProductConditions) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


