# ManageSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The unique ID which was used to create subscription. | 
**Action** | **string** | Action to be performed on the subscription. Possible values - CANCEL, PAUSE, ACTIVATE, CHANGE_PLAN. | 
**ActionDetails** | Pointer to [**ManageSubscriptionRequestActionDetails**](ManageSubscriptionRequestActionDetails.md) |  | [optional] 

## Methods

### NewManageSubscriptionRequest

`func NewManageSubscriptionRequest(subscriptionId string, action string, ) *ManageSubscriptionRequest`

NewManageSubscriptionRequest instantiates a new ManageSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageSubscriptionRequestWithDefaults

`func NewManageSubscriptionRequestWithDefaults() *ManageSubscriptionRequest`

NewManageSubscriptionRequestWithDefaults instantiates a new ManageSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *ManageSubscriptionRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ManageSubscriptionRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ManageSubscriptionRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetAction

`func (o *ManageSubscriptionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageSubscriptionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageSubscriptionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionDetails

`func (o *ManageSubscriptionRequest) GetActionDetails() ManageSubscriptionRequestActionDetails`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *ManageSubscriptionRequest) GetActionDetailsOk() (*ManageSubscriptionRequestActionDetails, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *ManageSubscriptionRequest) SetActionDetails(v ManageSubscriptionRequestActionDetails)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *ManageSubscriptionRequest) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


