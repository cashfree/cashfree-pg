# LinkPlanEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | Unique ID to identify the plan. Only alpha-numerics, dot, hyphen and underscore allowed. | [optional] 
**PlanName** | Pointer to **string** | Name of the plan. | [optional] 
**PlanType** | Pointer to **string** | Type of the plan. Possible values - PERIODIC, ON_DEMAND. | [optional] 
**PlanCurrency** | Pointer to **string** | Currency of the plan. | [optional] 
**PlanAmount** | Pointer to **float32** | The amount to be charged for PERIODIC plan. This is a conditional parameter, only required for PERIODIC plans. | [optional] 
**PlanMaxAmount** | Pointer to **float32** | Maximum amount for the plan. | [optional] 
**PlanMaxCycles** | Pointer to **int32** | Maximum number of payment cycles for the plan. | [optional] 
**PlanIntervals** | Pointer to **int32** | Number of billing cycles between charges. For instance, if set to 2 and the interval type is &#39;week&#39;, the service will be billed every 2 weeks. Similarly, if set to 3 and the interval type is &#39;month&#39;, the service will be billed every 3 months. Required for PERIODIC plan_type. | [optional] 
**PlanIntervalType** | Pointer to **string** | Interval type for the plan. Possible values - DAY, WEEK, MONTH, YEAR. | [optional] 
**PlanNote** | Pointer to **string** | Note for the plan. | [optional] 

## Methods

### NewLinkPlanEntity

`func NewLinkPlanEntity() *LinkPlanEntity`

NewLinkPlanEntity instantiates a new LinkPlanEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkPlanEntityWithDefaults

`func NewLinkPlanEntityWithDefaults() *LinkPlanEntity`

NewLinkPlanEntityWithDefaults instantiates a new LinkPlanEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *LinkPlanEntity) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *LinkPlanEntity) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *LinkPlanEntity) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *LinkPlanEntity) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *LinkPlanEntity) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *LinkPlanEntity) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *LinkPlanEntity) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *LinkPlanEntity) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanType

`func (o *LinkPlanEntity) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *LinkPlanEntity) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *LinkPlanEntity) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *LinkPlanEntity) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPlanCurrency

`func (o *LinkPlanEntity) GetPlanCurrency() string`

GetPlanCurrency returns the PlanCurrency field if non-nil, zero value otherwise.

### GetPlanCurrencyOk

`func (o *LinkPlanEntity) GetPlanCurrencyOk() (*string, bool)`

GetPlanCurrencyOk returns a tuple with the PlanCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCurrency

`func (o *LinkPlanEntity) SetPlanCurrency(v string)`

SetPlanCurrency sets PlanCurrency field to given value.

### HasPlanCurrency

`func (o *LinkPlanEntity) HasPlanCurrency() bool`

HasPlanCurrency returns a boolean if a field has been set.

### GetPlanAmount

`func (o *LinkPlanEntity) GetPlanAmount() float32`

GetPlanAmount returns the PlanAmount field if non-nil, zero value otherwise.

### GetPlanAmountOk

`func (o *LinkPlanEntity) GetPlanAmountOk() (*float32, bool)`

GetPlanAmountOk returns a tuple with the PlanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAmount

`func (o *LinkPlanEntity) SetPlanAmount(v float32)`

SetPlanAmount sets PlanAmount field to given value.

### HasPlanAmount

`func (o *LinkPlanEntity) HasPlanAmount() bool`

HasPlanAmount returns a boolean if a field has been set.

### GetPlanMaxAmount

`func (o *LinkPlanEntity) GetPlanMaxAmount() float32`

GetPlanMaxAmount returns the PlanMaxAmount field if non-nil, zero value otherwise.

### GetPlanMaxAmountOk

`func (o *LinkPlanEntity) GetPlanMaxAmountOk() (*float32, bool)`

GetPlanMaxAmountOk returns a tuple with the PlanMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxAmount

`func (o *LinkPlanEntity) SetPlanMaxAmount(v float32)`

SetPlanMaxAmount sets PlanMaxAmount field to given value.

### HasPlanMaxAmount

`func (o *LinkPlanEntity) HasPlanMaxAmount() bool`

HasPlanMaxAmount returns a boolean if a field has been set.

### GetPlanMaxCycles

`func (o *LinkPlanEntity) GetPlanMaxCycles() int32`

GetPlanMaxCycles returns the PlanMaxCycles field if non-nil, zero value otherwise.

### GetPlanMaxCyclesOk

`func (o *LinkPlanEntity) GetPlanMaxCyclesOk() (*int32, bool)`

GetPlanMaxCyclesOk returns a tuple with the PlanMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxCycles

`func (o *LinkPlanEntity) SetPlanMaxCycles(v int32)`

SetPlanMaxCycles sets PlanMaxCycles field to given value.

### HasPlanMaxCycles

`func (o *LinkPlanEntity) HasPlanMaxCycles() bool`

HasPlanMaxCycles returns a boolean if a field has been set.

### GetPlanIntervals

`func (o *LinkPlanEntity) GetPlanIntervals() int32`

GetPlanIntervals returns the PlanIntervals field if non-nil, zero value otherwise.

### GetPlanIntervalsOk

`func (o *LinkPlanEntity) GetPlanIntervalsOk() (*int32, bool)`

GetPlanIntervalsOk returns a tuple with the PlanIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervals

`func (o *LinkPlanEntity) SetPlanIntervals(v int32)`

SetPlanIntervals sets PlanIntervals field to given value.

### HasPlanIntervals

`func (o *LinkPlanEntity) HasPlanIntervals() bool`

HasPlanIntervals returns a boolean if a field has been set.

### GetPlanIntervalType

`func (o *LinkPlanEntity) GetPlanIntervalType() string`

GetPlanIntervalType returns the PlanIntervalType field if non-nil, zero value otherwise.

### GetPlanIntervalTypeOk

`func (o *LinkPlanEntity) GetPlanIntervalTypeOk() (*string, bool)`

GetPlanIntervalTypeOk returns a tuple with the PlanIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervalType

`func (o *LinkPlanEntity) SetPlanIntervalType(v string)`

SetPlanIntervalType sets PlanIntervalType field to given value.

### HasPlanIntervalType

`func (o *LinkPlanEntity) HasPlanIntervalType() bool`

HasPlanIntervalType returns a boolean if a field has been set.

### GetPlanNote

`func (o *LinkPlanEntity) GetPlanNote() string`

GetPlanNote returns the PlanNote field if non-nil, zero value otherwise.

### GetPlanNoteOk

`func (o *LinkPlanEntity) GetPlanNoteOk() (*string, bool)`

GetPlanNoteOk returns a tuple with the PlanNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNote

`func (o *LinkPlanEntity) SetPlanNote(v string)`

SetPlanNote sets PlanNote field to given value.

### HasPlanNote

`func (o *LinkPlanEntity) HasPlanNote() bool`

HasPlanNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


