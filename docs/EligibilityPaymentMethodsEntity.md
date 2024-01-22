# EligibilityPaymentMethodsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligibility** | Pointer to **bool** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**EntityValue** | Pointer to **string** |  | [optional] 
**EntityDetails** | Pointer to [**EligibilityPaymentMethodsEntityEntityDetails**](EligibilityPaymentMethodsEntityEntityDetails.md) |  | [optional] 

## Methods

### NewEligibilityPaymentMethodsEntity

`func NewEligibilityPaymentMethodsEntity() *EligibilityPaymentMethodsEntity`

NewEligibilityPaymentMethodsEntity instantiates a new EligibilityPaymentMethodsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityPaymentMethodsEntityWithDefaults

`func NewEligibilityPaymentMethodsEntityWithDefaults() *EligibilityPaymentMethodsEntity`

NewEligibilityPaymentMethodsEntityWithDefaults instantiates a new EligibilityPaymentMethodsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibility

`func (o *EligibilityPaymentMethodsEntity) GetEligibility() bool`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *EligibilityPaymentMethodsEntity) GetEligibilityOk() (*bool, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *EligibilityPaymentMethodsEntity) SetEligibility(v bool)`

SetEligibility sets Eligibility field to given value.

### HasEligibility

`func (o *EligibilityPaymentMethodsEntity) HasEligibility() bool`

HasEligibility returns a boolean if a field has been set.

### GetEntityType

`func (o *EligibilityPaymentMethodsEntity) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EligibilityPaymentMethodsEntity) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EligibilityPaymentMethodsEntity) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EligibilityPaymentMethodsEntity) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityValue

`func (o *EligibilityPaymentMethodsEntity) GetEntityValue() string`

GetEntityValue returns the EntityValue field if non-nil, zero value otherwise.

### GetEntityValueOk

`func (o *EligibilityPaymentMethodsEntity) GetEntityValueOk() (*string, bool)`

GetEntityValueOk returns a tuple with the EntityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityValue

`func (o *EligibilityPaymentMethodsEntity) SetEntityValue(v string)`

SetEntityValue sets EntityValue field to given value.

### HasEntityValue

`func (o *EligibilityPaymentMethodsEntity) HasEntityValue() bool`

HasEntityValue returns a boolean if a field has been set.

### GetEntityDetails

`func (o *EligibilityPaymentMethodsEntity) GetEntityDetails() EligibilityPaymentMethodsEntityEntityDetails`

GetEntityDetails returns the EntityDetails field if non-nil, zero value otherwise.

### GetEntityDetailsOk

`func (o *EligibilityPaymentMethodsEntity) GetEntityDetailsOk() (*EligibilityPaymentMethodsEntityEntityDetails, bool)`

GetEntityDetailsOk returns a tuple with the EntityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDetails

`func (o *EligibilityPaymentMethodsEntity) SetEntityDetails(v EligibilityPaymentMethodsEntityEntityDetails)`

SetEntityDetails sets EntityDetails field to given value.

### HasEntityDetails

`func (o *EligibilityPaymentMethodsEntity) HasEntityDetails() bool`

HasEntityDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


