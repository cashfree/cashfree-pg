# CreateSubscriptionRequestPlanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | Pointer to **string** | The unique identifier used to create plan. You only need to pass this field if you had already created plan. Otherwise use the other fields here to define the plan. | [optional] 
**PlanName** | Pointer to **string** | Specify plan name for easy reference. | [optional] 
**PlanType** | Pointer to **string** | Possible values ON_DEMAND or PERIODIC. PERIODIC - Payments are triggered automatically at fixed intervals defined by the merchant. ON_DEMAND - Merchant needs to trigger/charge the customer explicitly with the required amount. | [optional] 
**PlanCurrency** | Pointer to **string** | INR by default. | [optional] 
**PlanAmount** | Pointer to **float32** | The amount to be charged for PERIODIC plan. This is a conditional parameter, only required for PERIODIC plans. | [optional] 
**PlanMaxAmount** | Pointer to **float32** | This is the maximum amount that can be charged on a subscription. | [optional] 
**PlanMaxCycles** | Pointer to **int32** | Maximum number of debits set for the plan. The subscription will automatically change to COMPLETED status once this limit is reached. | [optional] 
**PlanIntervals** | Pointer to **int32** | Number of intervals of intervalType between every subscription payment. For example, to charge a customer bi-weekly use intervalType as “week” and intervals as 2. Required for PERIODIC plan. The default value is 1. | [optional] 
**PlanIntervalType** | Pointer to **string** | The type of interval for a PERIODIC plan like DAY, WEEK, MONTH, or YEAR. This is a conditional parameter only applicable for PERIODIC plans. | [optional] 
**PlanNote** | Pointer to **string** | Note for the plan. | [optional] 

## Methods

### NewCreateSubscriptionRequestPlanDetails

`func NewCreateSubscriptionRequestPlanDetails() *CreateSubscriptionRequestPlanDetails`

NewCreateSubscriptionRequestPlanDetails instantiates a new CreateSubscriptionRequestPlanDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestPlanDetailsWithDefaults

`func NewCreateSubscriptionRequestPlanDetailsWithDefaults() *CreateSubscriptionRequestPlanDetails`

NewCreateSubscriptionRequestPlanDetailsWithDefaults instantiates a new CreateSubscriptionRequestPlanDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanType

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPlanCurrency

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanCurrency() string`

GetPlanCurrency returns the PlanCurrency field if non-nil, zero value otherwise.

### GetPlanCurrencyOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanCurrencyOk() (*string, bool)`

GetPlanCurrencyOk returns a tuple with the PlanCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCurrency

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanCurrency(v string)`

SetPlanCurrency sets PlanCurrency field to given value.

### HasPlanCurrency

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanCurrency() bool`

HasPlanCurrency returns a boolean if a field has been set.

### GetPlanAmount

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanAmount() float32`

GetPlanAmount returns the PlanAmount field if non-nil, zero value otherwise.

### GetPlanAmountOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanAmountOk() (*float32, bool)`

GetPlanAmountOk returns a tuple with the PlanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAmount

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanAmount(v float32)`

SetPlanAmount sets PlanAmount field to given value.

### HasPlanAmount

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanAmount() bool`

HasPlanAmount returns a boolean if a field has been set.

### GetPlanMaxAmount

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanMaxAmount() float32`

GetPlanMaxAmount returns the PlanMaxAmount field if non-nil, zero value otherwise.

### GetPlanMaxAmountOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanMaxAmountOk() (*float32, bool)`

GetPlanMaxAmountOk returns a tuple with the PlanMaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxAmount

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanMaxAmount(v float32)`

SetPlanMaxAmount sets PlanMaxAmount field to given value.

### HasPlanMaxAmount

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanMaxAmount() bool`

HasPlanMaxAmount returns a boolean if a field has been set.

### GetPlanMaxCycles

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanMaxCycles() int32`

GetPlanMaxCycles returns the PlanMaxCycles field if non-nil, zero value otherwise.

### GetPlanMaxCyclesOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanMaxCyclesOk() (*int32, bool)`

GetPlanMaxCyclesOk returns a tuple with the PlanMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMaxCycles

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanMaxCycles(v int32)`

SetPlanMaxCycles sets PlanMaxCycles field to given value.

### HasPlanMaxCycles

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanMaxCycles() bool`

HasPlanMaxCycles returns a boolean if a field has been set.

### GetPlanIntervals

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanIntervals() int32`

GetPlanIntervals returns the PlanIntervals field if non-nil, zero value otherwise.

### GetPlanIntervalsOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanIntervalsOk() (*int32, bool)`

GetPlanIntervalsOk returns a tuple with the PlanIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervals

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanIntervals(v int32)`

SetPlanIntervals sets PlanIntervals field to given value.

### HasPlanIntervals

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanIntervals() bool`

HasPlanIntervals returns a boolean if a field has been set.

### GetPlanIntervalType

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanIntervalType() string`

GetPlanIntervalType returns the PlanIntervalType field if non-nil, zero value otherwise.

### GetPlanIntervalTypeOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanIntervalTypeOk() (*string, bool)`

GetPlanIntervalTypeOk returns a tuple with the PlanIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIntervalType

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanIntervalType(v string)`

SetPlanIntervalType sets PlanIntervalType field to given value.

### HasPlanIntervalType

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanIntervalType() bool`

HasPlanIntervalType returns a boolean if a field has been set.

### GetPlanNote

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanNote() string`

GetPlanNote returns the PlanNote field if non-nil, zero value otherwise.

### GetPlanNoteOk

`func (o *CreateSubscriptionRequestPlanDetails) GetPlanNoteOk() (*string, bool)`

GetPlanNoteOk returns a tuple with the PlanNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanNote

`func (o *CreateSubscriptionRequestPlanDetails) SetPlanNote(v string)`

SetPlanNote sets PlanNote field to given value.

### HasPlanNote

`func (o *CreateSubscriptionRequestPlanDetails) HasPlanNote() bool`

HasPlanNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


