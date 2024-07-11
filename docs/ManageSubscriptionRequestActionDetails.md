# ManageSubscriptionRequestActionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextScheduledTime** | Pointer to **string** | Next scheduled time for the action. Required for ACTIVATE action. | [optional] 
**PlanId** | Pointer to **string** | Plan ID to update. Required for CHANGE_PLAN action. | [optional] 

## Methods

### NewManageSubscriptionRequestActionDetails

`func NewManageSubscriptionRequestActionDetails() *ManageSubscriptionRequestActionDetails`

NewManageSubscriptionRequestActionDetails instantiates a new ManageSubscriptionRequestActionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageSubscriptionRequestActionDetailsWithDefaults

`func NewManageSubscriptionRequestActionDetailsWithDefaults() *ManageSubscriptionRequestActionDetails`

NewManageSubscriptionRequestActionDetailsWithDefaults instantiates a new ManageSubscriptionRequestActionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextScheduledTime

`func (o *ManageSubscriptionRequestActionDetails) GetNextScheduledTime() string`

GetNextScheduledTime returns the NextScheduledTime field if non-nil, zero value otherwise.

### GetNextScheduledTimeOk

`func (o *ManageSubscriptionRequestActionDetails) GetNextScheduledTimeOk() (*string, bool)`

GetNextScheduledTimeOk returns a tuple with the NextScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledTime

`func (o *ManageSubscriptionRequestActionDetails) SetNextScheduledTime(v string)`

SetNextScheduledTime sets NextScheduledTime field to given value.

### HasNextScheduledTime

`func (o *ManageSubscriptionRequestActionDetails) HasNextScheduledTime() bool`

HasNextScheduledTime returns a boolean if a field has been set.

### GetPlanId

`func (o *ManageSubscriptionRequestActionDetails) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ManageSubscriptionRequestActionDetails) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ManageSubscriptionRequestActionDetails) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ManageSubscriptionRequestActionDetails) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


