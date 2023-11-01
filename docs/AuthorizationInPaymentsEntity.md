# AuthorizationInPaymentsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | One of CAPTURE or VOID | [optional] 
**Status** | Pointer to **string** | One of SUCCESS or PENDING | [optional] 
**CapturedAmount** | Pointer to **float32** | The captured amount for this authorization request | [optional] 
**StartTime** | Pointer to **string** | Start time of this authorization hold (only for UPI) | [optional] 
**EndTime** | Pointer to **string** | End time of this authorization hold (only for UPI) | [optional] 
**ApproveBy** | Pointer to **string** | Approve by time as passed in the authorization request (only for UPI) | [optional] 
**ActionReference** | Pointer to **string** | CAPTURE or VOID reference number based on action  | [optional] 
**ActionTime** | Pointer to **string** | Time of action (CAPTURE or VOID) | [optional] 

## Methods

### NewAuthorizationInPaymentsEntity

`func NewAuthorizationInPaymentsEntity() *AuthorizationInPaymentsEntity`

NewAuthorizationInPaymentsEntity instantiates a new AuthorizationInPaymentsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationInPaymentsEntityWithDefaults

`func NewAuthorizationInPaymentsEntityWithDefaults() *AuthorizationInPaymentsEntity`

NewAuthorizationInPaymentsEntityWithDefaults instantiates a new AuthorizationInPaymentsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuthorizationInPaymentsEntity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizationInPaymentsEntity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizationInPaymentsEntity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizationInPaymentsEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationInPaymentsEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationInPaymentsEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationInPaymentsEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationInPaymentsEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCapturedAmount

`func (o *AuthorizationInPaymentsEntity) GetCapturedAmount() float32`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *AuthorizationInPaymentsEntity) GetCapturedAmountOk() (*float32, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *AuthorizationInPaymentsEntity) SetCapturedAmount(v float32)`

SetCapturedAmount sets CapturedAmount field to given value.

### HasCapturedAmount

`func (o *AuthorizationInPaymentsEntity) HasCapturedAmount() bool`

HasCapturedAmount returns a boolean if a field has been set.

### GetStartTime

`func (o *AuthorizationInPaymentsEntity) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthorizationInPaymentsEntity) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthorizationInPaymentsEntity) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthorizationInPaymentsEntity) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthorizationInPaymentsEntity) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthorizationInPaymentsEntity) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthorizationInPaymentsEntity) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthorizationInPaymentsEntity) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetApproveBy

`func (o *AuthorizationInPaymentsEntity) GetApproveBy() string`

GetApproveBy returns the ApproveBy field if non-nil, zero value otherwise.

### GetApproveByOk

`func (o *AuthorizationInPaymentsEntity) GetApproveByOk() (*string, bool)`

GetApproveByOk returns a tuple with the ApproveBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveBy

`func (o *AuthorizationInPaymentsEntity) SetApproveBy(v string)`

SetApproveBy sets ApproveBy field to given value.

### HasApproveBy

`func (o *AuthorizationInPaymentsEntity) HasApproveBy() bool`

HasApproveBy returns a boolean if a field has been set.

### GetActionReference

`func (o *AuthorizationInPaymentsEntity) GetActionReference() string`

GetActionReference returns the ActionReference field if non-nil, zero value otherwise.

### GetActionReferenceOk

`func (o *AuthorizationInPaymentsEntity) GetActionReferenceOk() (*string, bool)`

GetActionReferenceOk returns a tuple with the ActionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionReference

`func (o *AuthorizationInPaymentsEntity) SetActionReference(v string)`

SetActionReference sets ActionReference field to given value.

### HasActionReference

`func (o *AuthorizationInPaymentsEntity) HasActionReference() bool`

HasActionReference returns a boolean if a field has been set.

### GetActionTime

`func (o *AuthorizationInPaymentsEntity) GetActionTime() string`

GetActionTime returns the ActionTime field if non-nil, zero value otherwise.

### GetActionTimeOk

`func (o *AuthorizationInPaymentsEntity) GetActionTimeOk() (*string, bool)`

GetActionTimeOk returns a tuple with the ActionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTime

`func (o *AuthorizationInPaymentsEntity) SetActionTime(v string)`

SetActionTime sets ActionTime field to given value.

### HasActionTime

`func (o *AuthorizationInPaymentsEntity) HasActionTime() bool`

HasActionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


