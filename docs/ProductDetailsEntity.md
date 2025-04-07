# ProductDetailsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the feature has been enabled for this order | [optional] 
**Conditions** | Pointer to [**[]ProductConditionsEntity**](ProductConditionsEntity.md) | Configured condtions for the feature | [optional] 

## Methods

### NewProductDetailsEntity

`func NewProductDetailsEntity() *ProductDetailsEntity`

NewProductDetailsEntity instantiates a new ProductDetailsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailsEntityWithDefaults

`func NewProductDetailsEntityWithDefaults() *ProductDetailsEntity`

NewProductDetailsEntityWithDefaults instantiates a new ProductDetailsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ProductDetailsEntity) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProductDetailsEntity) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProductDetailsEntity) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProductDetailsEntity) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConditions

`func (o *ProductDetailsEntity) GetConditions() []ProductConditionsEntity`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ProductDetailsEntity) GetConditionsOk() (*[]ProductConditionsEntity, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ProductDetailsEntity) SetConditions(v []ProductConditionsEntity)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ProductDetailsEntity) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


