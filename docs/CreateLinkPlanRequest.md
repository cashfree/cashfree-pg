# CreateLinkPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** | Unique ID to identify the plan. Only alpha-numerics, dot, hyphen and underscore allowed. | 
**PlanName** | **string** | Name of the plan. | 
**PlanType** | **string** | Type of the plan. Possible values - PERIODIC, ON_DEMAND. | 
**PlanCurrency** | Pointer to **string** | Currency of the plan. | [optional] 
**PlanAmount** | Pointer to **float32** | The amount to be charged for PERIODIC plan. This is a conditional parameter, only required for PERIODIC plans. | [optional] 
**PlanMaxAmount** | **float32** | Maximum amount for the plan. | 
**PlanMaxCycles** | Pointer to **int32** | Maximum number of payment cycles for the plan. | [optional] 
**PlanIntervals** | Pointer to **int32** | Number of billing cycles between charges. | [optional] 
**PlanIntervalType** | Pointer to **string** | Interval type for the plan. Possible values - DAY, WEEK, MONTH, YEAR. | [optional] 
**PlanNote** | Pointer to **string** | Note for the plan. | [optional] 

## Methods

### NewCreateLinkPlanRequest

`func NewCreateLinkPlanRequest(planId string, planName string, planType string, planMaxAmount float32, ) *CreateLinkPlanRequest`

NewCreateLinkPlanRequest instantiates a new CreateLinkPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkPlanRequestWithDefaults

`func NewCreateLinkPlanRequestWithDefaults() *CreateLinkPlanRequest`

NewCreateLinkPlanRequestWithDefaults instantiates a new CreateLinkPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreateLinkPlanRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateLinkPlanRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateLinkPlanRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanName

`func (o *CreateLinkPlanRequest) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CreateLinkPlanRequest) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CreateLinkPlanRequest) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetPlanType

`func (o *CreateLinkPlanRequest) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CreateLinkPlanRequest) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CreateLinkPlanRequest) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetPlanCurrency

`func (o *CreateLinkPlanRequest) GetPlanCurrency() string`

GetPlanCurrency returns the PlanCurrency field if non-nil, zero value otherwise.

### GetPlanCurrencyOk

`func (o *CreateLinkPlanRequest) GetPlanCurrencyOk() (*string, bool)`

GetPlanCurrencyOk returns a tuple with the PlanCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCurrency

`func (o *CreateLinkPlanRequest) SetPlanCurrency(v string)`

SetPlanCurrency sets PlanCurrency field to given value.

### HasPlanCurrency

`func (o *CreateLinkPlanRequest) HasPlanCurrency() bool`

HasPlanCurrency returns a boolean if a field has been set.

### GetPlanAmount

`func (o *CreateLinkPlanRequest) GetPlanAmount() float32`

GetPlanAmount returns the PlanAmount field if non-nil, zero value otherwise.

### GetPlanAmountOk

`func (o *CreateLinkPlanRequest) GetPlanAmountOk() (*float32, bool)`

GetPlanAmountOk returns a tuple with the PlanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAmount

`func (o *CreateLinkPlanRequest) SetPlanAmount(v float32)`

SetPlanAmount sets PlanAmount field to given value.

### HasPlanAmount

`func (o *CreateLinkPlanRequest) HasPlanAmount() bool`

HasPlanAmount returns a boolean if a field has been set.

### GetPlanMaxAmount

`func (o *CreateLinkPlanRequest) GetPlanMaxAmount() float32`

GetPlanMaxAmount returns the PlanMaxAmount field if non-nil, zero value otherwise.

### GetPlanMaxAmountOk

`func (o *CreateLinkPlanRequest) GetPlanMaxAmountOk() (*float32, bool)`

GetPlanMaxAmountOk returns a tuple with the PlanMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxAmount

`func (o *CreateLinkPlanRequest) SetPlanMaxAmount(v float32)`

SetPlanMaxAmount sets PlanMaxAmount field to given value.


### GetPlanMaxCycles

`func (o *CreateLinkPlanRequest) GetPlanMaxCycles() int32`

GetPlanMaxCycles returns the PlanMaxCycles field if non-nil, zero value otherwise.

### GetPlanMaxCyclesOk

`func (o *CreateLinkPlanRequest) GetPlanMaxCyclesOk() (*int32, bool)`

GetPlanMaxCyclesOk returns a tuple with the PlanMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxCycles

`func (o *CreateLinkPlanRequest) SetPlanMaxCycles(v int32)`

SetPlanMaxCycles sets PlanMaxCycles field to given value.

### HasPlanMaxCycles

`func (o *CreateLinkPlanRequest) HasPlanMaxCycles() bool`

HasPlanMaxCycles returns a boolean if a field has been set.

### GetPlanIntervals

`func (o *CreateLinkPlanRequest) GetPlanIntervals() int32`

GetPlanIntervals returns the PlanIntervals field if non-nil, zero value otherwise.

### GetPlanIntervalsOk

`func (o *CreateLinkPlanRequest) GetPlanIntervalsOk() (*int32, bool)`

GetPlanIntervalsOk returns a tuple with the PlanIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervals

`func (o *CreateLinkPlanRequest) SetPlanIntervals(v int32)`

SetPlanIntervals sets PlanIntervals field to given value.

### HasPlanIntervals

`func (o *CreateLinkPlanRequest) HasPlanIntervals() bool`

HasPlanIntervals returns a boolean if a field has been set.

### GetPlanIntervalType

`func (o *CreateLinkPlanRequest) GetPlanIntervalType() string`

GetPlanIntervalType returns the PlanIntervalType field if non-nil, zero value otherwise.

### GetPlanIntervalTypeOk

`func (o *CreateLinkPlanRequest) GetPlanIntervalTypeOk() (*string, bool)`

GetPlanIntervalTypeOk returns a tuple with the PlanIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervalType

`func (o *CreateLinkPlanRequest) SetPlanIntervalType(v string)`

SetPlanIntervalType sets PlanIntervalType field to given value.

### HasPlanIntervalType

`func (o *CreateLinkPlanRequest) HasPlanIntervalType() bool`

HasPlanIntervalType returns a boolean if a field has been set.

### GetPlanNote

`func (o *CreateLinkPlanRequest) GetPlanNote() string`

GetPlanNote returns the PlanNote field if non-nil, zero value otherwise.

### GetPlanNoteOk

`func (o *CreateLinkPlanRequest) GetPlanNoteOk() (*string, bool)`

GetPlanNoteOk returns a tuple with the PlanNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNote

`func (o *CreateLinkPlanRequest) SetPlanNote(v string)`

SetPlanNote sets PlanNote field to given value.

### HasPlanNote

`func (o *CreateLinkPlanRequest) HasPlanNote() bool`

HasPlanNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


