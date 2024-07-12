# EligibilityMethodItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligibility** | Pointer to **bool** | Indicates whether the payment method is eligible. | [optional] 
**EntityType** | Pointer to **string** | Type of entity (e.g., \&quot;payment_methods\&quot;). | [optional] 
**EntityValue** | Pointer to **string** | Payment method (e.g., enach, pnach, upi, card). | [optional] 
**EntityDetails** | Pointer to [**EligibilityMethodItemEntityDetails**](EligibilityMethodItemEntityDetails.md) |  | [optional] 

## Methods

### NewEligibilityMethodItem

`func NewEligibilityMethodItem() *EligibilityMethodItem`

NewEligibilityMethodItem instantiates a new EligibilityMethodItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityMethodItemWithDefaults

`func NewEligibilityMethodItemWithDefaults() *EligibilityMethodItem`

NewEligibilityMethodItemWithDefaults instantiates a new EligibilityMethodItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibility

`func (o *EligibilityMethodItem) GetEligibility() bool`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *EligibilityMethodItem) GetEligibilityOk() (*bool, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *EligibilityMethodItem) SetEligibility(v bool)`

SetEligibility sets Eligibility field to given value.

### HasEligibility

`func (o *EligibilityMethodItem) HasEligibility() bool`

HasEligibility returns a boolean if a field has been set.

### GetEntityType

`func (o *EligibilityMethodItem) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EligibilityMethodItem) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EligibilityMethodItem) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EligibilityMethodItem) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityValue

`func (o *EligibilityMethodItem) GetEntityValue() string`

GetEntityValue returns the EntityValue field if non-nil, zero value otherwise.

### GetEntityValueOk

`func (o *EligibilityMethodItem) GetEntityValueOk() (*string, bool)`

GetEntityValueOk returns a tuple with the EntityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityValue

`func (o *EligibilityMethodItem) SetEntityValue(v string)`

SetEntityValue sets EntityValue field to given value.

### HasEntityValue

`func (o *EligibilityMethodItem) HasEntityValue() bool`

HasEntityValue returns a boolean if a field has been set.

### GetEntityDetails

`func (o *EligibilityMethodItem) GetEntityDetails() EligibilityMethodItemEntityDetails`

GetEntityDetails returns the EntityDetails field if non-nil, zero value otherwise.

### GetEntityDetailsOk

`func (o *EligibilityMethodItem) GetEntityDetailsOk() (*EligibilityMethodItemEntityDetails, bool)`

GetEntityDetailsOk returns a tuple with the EntityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDetails

`func (o *EligibilityMethodItem) SetEntityDetails(v EligibilityMethodItemEntityDetails)`

SetEntityDetails sets EntityDetails field to given value.

### HasEntityDetails

`func (o *EligibilityMethodItem) HasEntityDetails() bool`

HasEntityDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


