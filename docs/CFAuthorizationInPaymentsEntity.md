# CFAuthorizationInPaymentsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | One of CAPTURE or VOID | [optional] 
**Status** | Pointer to **string** | One of SUCCESS or PENDING | [optional] 
**CapturedAmount** | Pointer to **float32** | The captured amount for this authorization request | [optional] 
**StartTime** | Pointer to **string** | Start time of this authorization hold (only for UPI) | [optional] 
**EndTime** | Pointer to **string** | End time of this authorization hold (only for UPI) | [optional] 
**ApproveBy** | Pointer to **string** | Approve by time as passed in the authorization request (only for UPI) | [optional] 

## Methods

### NewCFAuthorizationInPaymentsEntity

`func NewCFAuthorizationInPaymentsEntity() *CFAuthorizationInPaymentsEntity`

NewCFAuthorizationInPaymentsEntity instantiates a new CFAuthorizationInPaymentsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFAuthorizationInPaymentsEntityWithDefaults

`func NewCFAuthorizationInPaymentsEntityWithDefaults() *CFAuthorizationInPaymentsEntity`

NewCFAuthorizationInPaymentsEntityWithDefaults instantiates a new CFAuthorizationInPaymentsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CFAuthorizationInPaymentsEntity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CFAuthorizationInPaymentsEntity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CFAuthorizationInPaymentsEntity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CFAuthorizationInPaymentsEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *CFAuthorizationInPaymentsEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CFAuthorizationInPaymentsEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CFAuthorizationInPaymentsEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CFAuthorizationInPaymentsEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCapturedAmount

`func (o *CFAuthorizationInPaymentsEntity) GetCapturedAmount() float32`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *CFAuthorizationInPaymentsEntity) GetCapturedAmountOk() (*float32, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *CFAuthorizationInPaymentsEntity) SetCapturedAmount(v float32)`

SetCapturedAmount sets CapturedAmount field to given value.

### HasCapturedAmount

`func (o *CFAuthorizationInPaymentsEntity) HasCapturedAmount() bool`

HasCapturedAmount returns a boolean if a field has been set.

### GetStartTime

`func (o *CFAuthorizationInPaymentsEntity) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CFAuthorizationInPaymentsEntity) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CFAuthorizationInPaymentsEntity) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CFAuthorizationInPaymentsEntity) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CFAuthorizationInPaymentsEntity) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CFAuthorizationInPaymentsEntity) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CFAuthorizationInPaymentsEntity) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CFAuthorizationInPaymentsEntity) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetApproveBy

`func (o *CFAuthorizationInPaymentsEntity) GetApproveBy() string`

GetApproveBy returns the ApproveBy field if non-nil, zero value otherwise.

### GetApproveByOk

`func (o *CFAuthorizationInPaymentsEntity) GetApproveByOk() (*string, bool)`

GetApproveByOk returns a tuple with the ApproveBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveBy

`func (o *CFAuthorizationInPaymentsEntity) SetApproveBy(v string)`

SetApproveBy sets ApproveBy field to given value.

### HasApproveBy

`func (o *CFAuthorizationInPaymentsEntity) HasApproveBy() bool`

HasApproveBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


