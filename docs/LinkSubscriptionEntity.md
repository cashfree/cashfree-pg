# LinkSubscriptionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanDetails** | Pointer to [**LinkPlanEntity**](LinkPlanEntity.md) |  | [optional] 
**AuthorizationAmountRefund** | Pointer to **bool** | Indicates whether the authorization amount should be refunded to the customer automatically. Merchants can use this field to specify if the authorized funds should be returned to the customer after authorization of the subscription. | [optional] 
**SubscriptionExpiryTime** | Pointer to **string** | Time at which the subscription will expire. | [optional] 
**SubscriptionFirstChargeTime** | Pointer to **string** | Time at which the first charge for the subscription will be processed. | [optional] 

## Methods

### NewLinkSubscriptionEntity

`func NewLinkSubscriptionEntity() *LinkSubscriptionEntity`

NewLinkSubscriptionEntity instantiates a new LinkSubscriptionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkSubscriptionEntityWithDefaults

`func NewLinkSubscriptionEntityWithDefaults() *LinkSubscriptionEntity`

NewLinkSubscriptionEntityWithDefaults instantiates a new LinkSubscriptionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanDetails

`func (o *LinkSubscriptionEntity) GetPlanDetails() LinkPlanEntity`

GetPlanDetails returns the PlanDetails field if non-nil, zero value otherwise.

### GetPlanDetailsOk

`func (o *LinkSubscriptionEntity) GetPlanDetailsOk() (*LinkPlanEntity, bool)`

GetPlanDetailsOk returns a tuple with the PlanDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDetails

`func (o *LinkSubscriptionEntity) SetPlanDetails(v LinkPlanEntity)`

SetPlanDetails sets PlanDetails field to given value.

### HasPlanDetails

`func (o *LinkSubscriptionEntity) HasPlanDetails() bool`

HasPlanDetails returns a boolean if a field has been set.

### GetAuthorizationAmountRefund

`func (o *LinkSubscriptionEntity) GetAuthorizationAmountRefund() bool`

GetAuthorizationAmountRefund returns the AuthorizationAmountRefund field if non-nil, zero value otherwise.

### GetAuthorizationAmountRefundOk

`func (o *LinkSubscriptionEntity) GetAuthorizationAmountRefundOk() (*bool, bool)`

GetAuthorizationAmountRefundOk returns a tuple with the AuthorizationAmountRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAmountRefund

`func (o *LinkSubscriptionEntity) SetAuthorizationAmountRefund(v bool)`

SetAuthorizationAmountRefund sets AuthorizationAmountRefund field to given value.

### HasAuthorizationAmountRefund

`func (o *LinkSubscriptionEntity) HasAuthorizationAmountRefund() bool`

HasAuthorizationAmountRefund returns a boolean if a field has been set.

### GetSubscriptionExpiryTime

`func (o *LinkSubscriptionEntity) GetSubscriptionExpiryTime() string`

GetSubscriptionExpiryTime returns the SubscriptionExpiryTime field if non-nil, zero value otherwise.

### GetSubscriptionExpiryTimeOk

`func (o *LinkSubscriptionEntity) GetSubscriptionExpiryTimeOk() (*string, bool)`

GetSubscriptionExpiryTimeOk returns a tuple with the SubscriptionExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiryTime

`func (o *LinkSubscriptionEntity) SetSubscriptionExpiryTime(v string)`

SetSubscriptionExpiryTime sets SubscriptionExpiryTime field to given value.

### HasSubscriptionExpiryTime

`func (o *LinkSubscriptionEntity) HasSubscriptionExpiryTime() bool`

HasSubscriptionExpiryTime returns a boolean if a field has been set.

### GetSubscriptionFirstChargeTime

`func (o *LinkSubscriptionEntity) GetSubscriptionFirstChargeTime() string`

GetSubscriptionFirstChargeTime returns the SubscriptionFirstChargeTime field if non-nil, zero value otherwise.

### GetSubscriptionFirstChargeTimeOk

`func (o *LinkSubscriptionEntity) GetSubscriptionFirstChargeTimeOk() (*string, bool)`

GetSubscriptionFirstChargeTimeOk returns a tuple with the SubscriptionFirstChargeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFirstChargeTime

`func (o *LinkSubscriptionEntity) SetSubscriptionFirstChargeTime(v string)`

SetSubscriptionFirstChargeTime sets SubscriptionFirstChargeTime field to given value.

### HasSubscriptionFirstChargeTime

`func (o *LinkSubscriptionEntity) HasSubscriptionFirstChargeTime() bool`

HasSubscriptionFirstChargeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


