# CardlessEMIEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | Pointer to **string** |  | [optional] 
**EmiPlans** | Pointer to [**[]EMIPlansArray**](EMIPlansArray.md) |  | [optional] 

## Methods

### NewCardlessEMIEntity

`func NewCardlessEMIEntity() *CardlessEMIEntity`

NewCardlessEMIEntity instantiates a new CardlessEMIEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardlessEMIEntityWithDefaults

`func NewCardlessEMIEntityWithDefaults() *CardlessEMIEntity`

NewCardlessEMIEntityWithDefaults instantiates a new CardlessEMIEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *CardlessEMIEntity) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CardlessEMIEntity) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CardlessEMIEntity) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CardlessEMIEntity) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetEmiPlans

`func (o *CardlessEMIEntity) GetEmiPlans() []EMIPlansArray`

GetEmiPlans returns the EmiPlans field if non-nil, zero value otherwise.

### GetEmiPlansOk

`func (o *CardlessEMIEntity) GetEmiPlansOk() (*[]EMIPlansArray, bool)`

GetEmiPlansOk returns a tuple with the EmiPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiPlans

`func (o *CardlessEMIEntity) SetEmiPlans(v []EMIPlansArray)`

SetEmiPlans sets EmiPlans field to given value.

### HasEmiPlans

`func (o *CardlessEMIEntity) HasEmiPlans() bool`

HasEmiPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


