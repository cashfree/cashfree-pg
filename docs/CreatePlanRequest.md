# CreatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | Unique ID to identify the plan. Only alpha-numerics, dot, hyphen and underscore allowed. | 
**PlanName** | **string** | Name of the plan. | 
**PlanType** | **string** | Type of the plan. Possible values - PERIODIC, ON_DEMAND. | 
**PlanCurrency** | Pointer to **string** | Currency of the plan. | [optional] 
**PlanRecurringAmount** | Pointer to **float32** | Recurring amount for the plan. Required for PERIODIC plan_type. | [optional] 
**PlanMaxAmount** | **float32** | Maximum amount for the plan. | 
**PlanMaxCycles** | Pointer to **int32** | Maximum number of payment cycles for the plan. | [optional] 
**PlanIntervals** | Pointer to **int32** | Number of billing cycles between charges. For instance, if set to 2 and the interval type is &#39;week&#39;, the service will be billed every 2 weeks. Similarly, if set to 3 and the interval type is &#39;month&#39;, the service will be billed every 3 months. Required for PERIODIC plan_type. | [optional] 
**PlanIntervalType** | Pointer to **string** | Interval type for the plan. Possible values - DAY, WEEK, MONTH, YEAR. | [optional] 
**PlanNote** | Pointer to **string** | Note for the plan. | [optional] 

## Methods

### NewCreatePlanRequest

`func NewCreatePlanRequest(planId string, planName string, planType string, planMaxAmount float32, ) *CreatePlanRequest`

NewCreatePlanRequest instantiates a new CreatePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanRequestWithDefaults

`func NewCreatePlanRequestWithDefaults() *CreatePlanRequest`

NewCreatePlanRequestWithDefaults instantiates a new CreatePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreatePlanRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreatePlanRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreatePlanRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanName

`func (o *CreatePlanRequest) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CreatePlanRequest) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CreatePlanRequest) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetPlanType

`func (o *CreatePlanRequest) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CreatePlanRequest) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CreatePlanRequest) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetPlanCurrency

`func (o *CreatePlanRequest) GetPlanCurrency() string`

GetPlanCurrency returns the PlanCurrency field if non-nil, zero value otherwise.

### GetPlanCurrencyOk

`func (o *CreatePlanRequest) GetPlanCurrencyOk() (*string, bool)`

GetPlanCurrencyOk returns a tuple with the PlanCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCurrency

`func (o *CreatePlanRequest) SetPlanCurrency(v string)`

SetPlanCurrency sets PlanCurrency field to given value.

### HasPlanCurrency

`func (o *CreatePlanRequest) HasPlanCurrency() bool`

HasPlanCurrency returns a boolean if a field has been set.

### GetPlanRecurringAmount

`func (o *CreatePlanRequest) GetPlanRecurringAmount() float32`

GetPlanRecurringAmount returns the PlanRecurringAmount field if non-nil, zero value otherwise.

### GetPlanRecurringAmountOk

`func (o *CreatePlanRequest) GetPlanRecurringAmountOk() (*float32, bool)`

GetPlanRecurringAmountOk returns a tuple with the PlanRecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRecurringAmount

`func (o *CreatePlanRequest) SetPlanRecurringAmount(v float32)`

SetPlanRecurringAmount sets PlanRecurringAmount field to given value.

### HasPlanRecurringAmount

`func (o *CreatePlanRequest) HasPlanRecurringAmount() bool`

HasPlanRecurringAmount returns a boolean if a field has been set.

### GetPlanMaxAmount

`func (o *CreatePlanRequest) GetPlanMaxAmount() float32`

GetPlanMaxAmount returns the PlanMaxAmount field if non-nil, zero value otherwise.

### GetPlanMaxAmountOk

`func (o *CreatePlanRequest) GetPlanMaxAmountOk() (*float32, bool)`

GetPlanMaxAmountOk returns a tuple with the PlanMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxAmount

`func (o *CreatePlanRequest) SetPlanMaxAmount(v float32)`

SetPlanMaxAmount sets PlanMaxAmount field to given value.


### GetPlanMaxCycles

`func (o *CreatePlanRequest) GetPlanMaxCycles() int32`

GetPlanMaxCycles returns the PlanMaxCycles field if non-nil, zero value otherwise.

### GetPlanMaxCyclesOk

`func (o *CreatePlanRequest) GetPlanMaxCyclesOk() (*int32, bool)`

GetPlanMaxCyclesOk returns a tuple with the PlanMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxCycles

`func (o *CreatePlanRequest) SetPlanMaxCycles(v int32)`

SetPlanMaxCycles sets PlanMaxCycles field to given value.

### HasPlanMaxCycles

`func (o *CreatePlanRequest) HasPlanMaxCycles() bool`

HasPlanMaxCycles returns a boolean if a field has been set.

### GetPlanIntervals

`func (o *CreatePlanRequest) GetPlanIntervals() int32`

GetPlanIntervals returns the PlanIntervals field if non-nil, zero value otherwise.

### GetPlanIntervalsOk

`func (o *CreatePlanRequest) GetPlanIntervalsOk() (*int32, bool)`

GetPlanIntervalsOk returns a tuple with the PlanIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervals

`func (o *CreatePlanRequest) SetPlanIntervals(v int32)`

SetPlanIntervals sets PlanIntervals field to given value.

### HasPlanIntervals

`func (o *CreatePlanRequest) HasPlanIntervals() bool`

HasPlanIntervals returns a boolean if a field has been set.

### GetPlanIntervalType

`func (o *CreatePlanRequest) GetPlanIntervalType() string`

GetPlanIntervalType returns the PlanIntervalType field if non-nil, zero value otherwise.

### GetPlanIntervalTypeOk

`func (o *CreatePlanRequest) GetPlanIntervalTypeOk() (*string, bool)`

GetPlanIntervalTypeOk returns a tuple with the PlanIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervalType

`func (o *CreatePlanRequest) SetPlanIntervalType(v string)`

SetPlanIntervalType sets PlanIntervalType field to given value.

### HasPlanIntervalType

`func (o *CreatePlanRequest) HasPlanIntervalType() bool`

HasPlanIntervalType returns a boolean if a field has been set.

### GetPlanNote

`func (o *CreatePlanRequest) GetPlanNote() string`

GetPlanNote returns the PlanNote field if non-nil, zero value otherwise.

### GetPlanNoteOk

`func (o *CreatePlanRequest) GetPlanNoteOk() (*string, bool)`

GetPlanNoteOk returns a tuple with the PlanNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNote

`func (o *CreatePlanRequest) SetPlanNote(v string)`

SetPlanNote sets PlanNote field to given value.

### HasPlanNote

`func (o *CreatePlanRequest) HasPlanNote() bool`

HasPlanNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


