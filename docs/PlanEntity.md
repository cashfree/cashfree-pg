# PlanEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanCurrency** | Pointer to **string** | Currency for the plan. | [optional] 
**PlanId** | Pointer to **string** | Plan ID provided by merchant. | [optional] 
**PlanIntervalType** | Pointer to **string** | Interval type for the plan. | [optional] 
**PlanIntervals** | Pointer to **int32** | Number of intervals for the plan. | [optional] 
**PlanMaxAmount** | Pointer to **float32** | Maximum amount for the plan. | [optional] 
**PlanMaxCycles** | Pointer to **int32** | Maximum number of payment cycles for the plan. | [optional] 
**PlanName** | Pointer to **string** | Name of the plan. | [optional] 
**PlanNote** | Pointer to **string** | Note for the plan. | [optional] 
**PlanRecurringAmount** | Pointer to **float32** | Recurring amount for the plan. | [optional] 
**PlanStatus** | Pointer to **string** | Status of the plan. | [optional] 
**PlanType** | Pointer to **string** | Type of the plan. | [optional] 

## Methods

### NewPlanEntity

`func NewPlanEntity() *PlanEntity`

NewPlanEntity instantiates a new PlanEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanEntityWithDefaults

`func NewPlanEntityWithDefaults() *PlanEntity`

NewPlanEntityWithDefaults instantiates a new PlanEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanCurrency

`func (o *PlanEntity) GetPlanCurrency() string`

GetPlanCurrency returns the PlanCurrency field if non-nil, zero value otherwise.

### GetPlanCurrencyOk

`func (o *PlanEntity) GetPlanCurrencyOk() (*string, bool)`

GetPlanCurrencyOk returns a tuple with the PlanCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCurrency

`func (o *PlanEntity) SetPlanCurrency(v string)`

SetPlanCurrency sets PlanCurrency field to given value.

### HasPlanCurrency

`func (o *PlanEntity) HasPlanCurrency() bool`

HasPlanCurrency returns a boolean if a field has been set.

### GetPlanId

`func (o *PlanEntity) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanEntity) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanEntity) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanEntity) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanIntervalType

`func (o *PlanEntity) GetPlanIntervalType() string`

GetPlanIntervalType returns the PlanIntervalType field if non-nil, zero value otherwise.

### GetPlanIntervalTypeOk

`func (o *PlanEntity) GetPlanIntervalTypeOk() (*string, bool)`

GetPlanIntervalTypeOk returns a tuple with the PlanIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervalType

`func (o *PlanEntity) SetPlanIntervalType(v string)`

SetPlanIntervalType sets PlanIntervalType field to given value.

### HasPlanIntervalType

`func (o *PlanEntity) HasPlanIntervalType() bool`

HasPlanIntervalType returns a boolean if a field has been set.

### GetPlanIntervals

`func (o *PlanEntity) GetPlanIntervals() int32`

GetPlanIntervals returns the PlanIntervals field if non-nil, zero value otherwise.

### GetPlanIntervalsOk

`func (o *PlanEntity) GetPlanIntervalsOk() (*int32, bool)`

GetPlanIntervalsOk returns a tuple with the PlanIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervals

`func (o *PlanEntity) SetPlanIntervals(v int32)`

SetPlanIntervals sets PlanIntervals field to given value.

### HasPlanIntervals

`func (o *PlanEntity) HasPlanIntervals() bool`

HasPlanIntervals returns a boolean if a field has been set.

### GetPlanMaxAmount

`func (o *PlanEntity) GetPlanMaxAmount() float32`

GetPlanMaxAmount returns the PlanMaxAmount field if non-nil, zero value otherwise.

### GetPlanMaxAmountOk

`func (o *PlanEntity) GetPlanMaxAmountOk() (*float32, bool)`

GetPlanMaxAmountOk returns a tuple with the PlanMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxAmount

`func (o *PlanEntity) SetPlanMaxAmount(v float32)`

SetPlanMaxAmount sets PlanMaxAmount field to given value.

### HasPlanMaxAmount

`func (o *PlanEntity) HasPlanMaxAmount() bool`

HasPlanMaxAmount returns a boolean if a field has been set.

### GetPlanMaxCycles

`func (o *PlanEntity) GetPlanMaxCycles() int32`

GetPlanMaxCycles returns the PlanMaxCycles field if non-nil, zero value otherwise.

### GetPlanMaxCyclesOk

`func (o *PlanEntity) GetPlanMaxCyclesOk() (*int32, bool)`

GetPlanMaxCyclesOk returns a tuple with the PlanMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxCycles

`func (o *PlanEntity) SetPlanMaxCycles(v int32)`

SetPlanMaxCycles sets PlanMaxCycles field to given value.

### HasPlanMaxCycles

`func (o *PlanEntity) HasPlanMaxCycles() bool`

HasPlanMaxCycles returns a boolean if a field has been set.

### GetPlanName

`func (o *PlanEntity) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *PlanEntity) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *PlanEntity) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *PlanEntity) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanNote

`func (o *PlanEntity) GetPlanNote() string`

GetPlanNote returns the PlanNote field if non-nil, zero value otherwise.

### GetPlanNoteOk

`func (o *PlanEntity) GetPlanNoteOk() (*string, bool)`

GetPlanNoteOk returns a tuple with the PlanNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNote

`func (o *PlanEntity) SetPlanNote(v string)`

SetPlanNote sets PlanNote field to given value.

### HasPlanNote

`func (o *PlanEntity) HasPlanNote() bool`

HasPlanNote returns a boolean if a field has been set.

### GetPlanRecurringAmount

`func (o *PlanEntity) GetPlanRecurringAmount() float32`

GetPlanRecurringAmount returns the PlanRecurringAmount field if non-nil, zero value otherwise.

### GetPlanRecurringAmountOk

`func (o *PlanEntity) GetPlanRecurringAmountOk() (*float32, bool)`

GetPlanRecurringAmountOk returns a tuple with the PlanRecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRecurringAmount

`func (o *PlanEntity) SetPlanRecurringAmount(v float32)`

SetPlanRecurringAmount sets PlanRecurringAmount field to given value.

### HasPlanRecurringAmount

`func (o *PlanEntity) HasPlanRecurringAmount() bool`

HasPlanRecurringAmount returns a boolean if a field has been set.

### GetPlanStatus

`func (o *PlanEntity) GetPlanStatus() string`

GetPlanStatus returns the PlanStatus field if non-nil, zero value otherwise.

### GetPlanStatusOk

`func (o *PlanEntity) GetPlanStatusOk() (*string, bool)`

GetPlanStatusOk returns a tuple with the PlanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanStatus

`func (o *PlanEntity) SetPlanStatus(v string)`

SetPlanStatus sets PlanStatus field to given value.

### HasPlanStatus

`func (o *PlanEntity) HasPlanStatus() bool`

HasPlanStatus returns a boolean if a field has been set.

### GetPlanType

`func (o *PlanEntity) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *PlanEntity) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *PlanEntity) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *PlanEntity) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


