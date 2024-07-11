# CreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | A unique ID for the subscription. It can include alphanumeric characters, underscore, dot, hyphen, and space. Maximum characters allowed is 250. | 
**CustomerDetails** | [**SubscriptionCustomerDetails**](SubscriptionCustomerDetails.md) |  | 
**PlanDetails** | [**CreateSubscriptionRequestPlanDetails**](CreateSubscriptionRequestPlanDetails.md) |  | 
**AuthorizationDetails** | Pointer to [**CreateSubscriptionRequestAuthorizationDetails**](CreateSubscriptionRequestAuthorizationDetails.md) |  | [optional] 
**SubscriptionMeta** | Pointer to [**CreateSubscriptionRequestSubscriptionMeta**](CreateSubscriptionRequestSubscriptionMeta.md) |  | [optional] 
**SubscriptionExpiryTime** | Pointer to **string** | Expiry date for the subscription. | [optional] 
**SubscriptionFirstChargeTime** | Pointer to **string** | Time at which the first charge will be made for the subscription after authorization. Applicable only for PERIODIC plans. | [optional] 
**SubscriptionNote** | Pointer to **string** | Note for the subscription. | [optional] 
**SubscriptionTags** | Pointer to **map[string]interface{}** | Tags for the subscription. | [optional] 
**SubscriptionPaymentSplits** | Pointer to [**[]SubscriptionPaymentSplitItem**](SubscriptionPaymentSplitItem.md) | Payment splits for the subscription. | [optional] 

## Methods

### NewCreateSubscriptionRequest

`func NewCreateSubscriptionRequest(subscriptionId string, customerDetails SubscriptionCustomerDetails, planDetails CreateSubscriptionRequestPlanDetails, ) *CreateSubscriptionRequest`

NewCreateSubscriptionRequest instantiates a new CreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestWithDefaults

`func NewCreateSubscriptionRequestWithDefaults() *CreateSubscriptionRequest`

NewCreateSubscriptionRequestWithDefaults instantiates a new CreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *CreateSubscriptionRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerDetails

`func (o *CreateSubscriptionRequest) GetCustomerDetails() SubscriptionCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CreateSubscriptionRequest) GetCustomerDetailsOk() (*SubscriptionCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CreateSubscriptionRequest) SetCustomerDetails(v SubscriptionCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.


### GetPlanDetails

`func (o *CreateSubscriptionRequest) GetPlanDetails() CreateSubscriptionRequestPlanDetails`

GetPlanDetails returns the PlanDetails field if non-nil, zero value otherwise.

### GetPlanDetailsOk

`func (o *CreateSubscriptionRequest) GetPlanDetailsOk() (*CreateSubscriptionRequestPlanDetails, bool)`

GetPlanDetailsOk returns a tuple with the PlanDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDetails

`func (o *CreateSubscriptionRequest) SetPlanDetails(v CreateSubscriptionRequestPlanDetails)`

SetPlanDetails sets PlanDetails field to given value.


### GetAuthorizationDetails

`func (o *CreateSubscriptionRequest) GetAuthorizationDetails() CreateSubscriptionRequestAuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *CreateSubscriptionRequest) GetAuthorizationDetailsOk() (*CreateSubscriptionRequestAuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *CreateSubscriptionRequest) SetAuthorizationDetails(v CreateSubscriptionRequestAuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *CreateSubscriptionRequest) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetSubscriptionMeta

`func (o *CreateSubscriptionRequest) GetSubscriptionMeta() CreateSubscriptionRequestSubscriptionMeta`

GetSubscriptionMeta returns the SubscriptionMeta field if non-nil, zero value otherwise.

### GetSubscriptionMetaOk

`func (o *CreateSubscriptionRequest) GetSubscriptionMetaOk() (*CreateSubscriptionRequestSubscriptionMeta, bool)`

GetSubscriptionMetaOk returns a tuple with the SubscriptionMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionMeta

`func (o *CreateSubscriptionRequest) SetSubscriptionMeta(v CreateSubscriptionRequestSubscriptionMeta)`

SetSubscriptionMeta sets SubscriptionMeta field to given value.

### HasSubscriptionMeta

`func (o *CreateSubscriptionRequest) HasSubscriptionMeta() bool`

HasSubscriptionMeta returns a boolean if a field has been set.

### GetSubscriptionExpiryTime

`func (o *CreateSubscriptionRequest) GetSubscriptionExpiryTime() string`

GetSubscriptionExpiryTime returns the SubscriptionExpiryTime field if non-nil, zero value otherwise.

### GetSubscriptionExpiryTimeOk

`func (o *CreateSubscriptionRequest) GetSubscriptionExpiryTimeOk() (*string, bool)`

GetSubscriptionExpiryTimeOk returns a tuple with the SubscriptionExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiryTime

`func (o *CreateSubscriptionRequest) SetSubscriptionExpiryTime(v string)`

SetSubscriptionExpiryTime sets SubscriptionExpiryTime field to given value.

### HasSubscriptionExpiryTime

`func (o *CreateSubscriptionRequest) HasSubscriptionExpiryTime() bool`

HasSubscriptionExpiryTime returns a boolean if a field has been set.

### GetSubscriptionFirstChargeTime

`func (o *CreateSubscriptionRequest) GetSubscriptionFirstChargeTime() string`

GetSubscriptionFirstChargeTime returns the SubscriptionFirstChargeTime field if non-nil, zero value otherwise.

### GetSubscriptionFirstChargeTimeOk

`func (o *CreateSubscriptionRequest) GetSubscriptionFirstChargeTimeOk() (*string, bool)`

GetSubscriptionFirstChargeTimeOk returns a tuple with the SubscriptionFirstChargeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFirstChargeTime

`func (o *CreateSubscriptionRequest) SetSubscriptionFirstChargeTime(v string)`

SetSubscriptionFirstChargeTime sets SubscriptionFirstChargeTime field to given value.

### HasSubscriptionFirstChargeTime

`func (o *CreateSubscriptionRequest) HasSubscriptionFirstChargeTime() bool`

HasSubscriptionFirstChargeTime returns a boolean if a field has been set.

### GetSubscriptionNote

`func (o *CreateSubscriptionRequest) GetSubscriptionNote() string`

GetSubscriptionNote returns the SubscriptionNote field if non-nil, zero value otherwise.

### GetSubscriptionNoteOk

`func (o *CreateSubscriptionRequest) GetSubscriptionNoteOk() (*string, bool)`

GetSubscriptionNoteOk returns a tuple with the SubscriptionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionNote

`func (o *CreateSubscriptionRequest) SetSubscriptionNote(v string)`

SetSubscriptionNote sets SubscriptionNote field to given value.

### HasSubscriptionNote

`func (o *CreateSubscriptionRequest) HasSubscriptionNote() bool`

HasSubscriptionNote returns a boolean if a field has been set.

### GetSubscriptionTags

`func (o *CreateSubscriptionRequest) GetSubscriptionTags() map[string]interface{}`

GetSubscriptionTags returns the SubscriptionTags field if non-nil, zero value otherwise.

### GetSubscriptionTagsOk

`func (o *CreateSubscriptionRequest) GetSubscriptionTagsOk() (*map[string]interface{}, bool)`

GetSubscriptionTagsOk returns a tuple with the SubscriptionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTags

`func (o *CreateSubscriptionRequest) SetSubscriptionTags(v map[string]interface{})`

SetSubscriptionTags sets SubscriptionTags field to given value.

### HasSubscriptionTags

`func (o *CreateSubscriptionRequest) HasSubscriptionTags() bool`

HasSubscriptionTags returns a boolean if a field has been set.

### GetSubscriptionPaymentSplits

`func (o *CreateSubscriptionRequest) GetSubscriptionPaymentSplits() []SubscriptionPaymentSplitItem`

GetSubscriptionPaymentSplits returns the SubscriptionPaymentSplits field if non-nil, zero value otherwise.

### GetSubscriptionPaymentSplitsOk

`func (o *CreateSubscriptionRequest) GetSubscriptionPaymentSplitsOk() (*[]SubscriptionPaymentSplitItem, bool)`

GetSubscriptionPaymentSplitsOk returns a tuple with the SubscriptionPaymentSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPaymentSplits

`func (o *CreateSubscriptionRequest) SetSubscriptionPaymentSplits(v []SubscriptionPaymentSplitItem)`

SetSubscriptionPaymentSplits sets SubscriptionPaymentSplits field to given value.

### HasSubscriptionPaymentSplits

`func (o *CreateSubscriptionRequest) HasSubscriptionPaymentSplits() bool`

HasSubscriptionPaymentSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


