# ProductDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Option to enable or disable the feature | [optional] 
**Conditions** | Pointer to [**[]ProductConditions**](ProductConditions.md) | The conditions array allows to configure rules by adding condition objects with specific parameters for feature configurations. | [optional] 

## Methods

### NewProductDetails

`func NewProductDetails() *ProductDetails`

NewProductDetails instantiates a new ProductDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDetailsWithDefaults

`func NewProductDetailsWithDefaults() *ProductDetails`

NewProductDetailsWithDefaults instantiates a new ProductDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ProductDetails) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProductDetails) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProductDetails) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProductDetails) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConditions

`func (o *ProductDetails) GetConditions() []ProductConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ProductDetails) GetConditionsOk() (*[]ProductConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ProductDetails) SetConditions(v []ProductConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ProductDetails) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


