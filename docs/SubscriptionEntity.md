# SubscriptionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorisationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**CfSubscriptionId** | Pointer to **string** | Cashfree subscription reference number | [optional] 
**CustomerDetails** | Pointer to [**SubscriptionCustomerDetails**](SubscriptionCustomerDetails.md) |  | [optional] 
**PlanDetails** | Pointer to [**PlanEntity**](PlanEntity.md) |  | [optional] 
**SubscriptionExpiryTime** | Pointer to **string** | Time at which the subscription will expire. | [optional] 
**SubscriptionFirstChargeTime** | Pointer to **string** | Time at which the first charge will be made for the subscription. Applicable only for PERIODIC plans. | [optional] 
**SubscriptionId** | Pointer to **string** | A unique ID passed by merchant for identifying the subscription. | [optional] 
**SubscriptionMeta** | Pointer to [**SubscriptionEntitySubscriptionMeta**](SubscriptionEntitySubscriptionMeta.md) |  | [optional] 
**SubscriptionNote** | Pointer to **string** | Note for the subscription. | [optional] 
**SubscriptionPaymentSplits** | Pointer to [**[]SubscriptionPaymentSplitItem**](SubscriptionPaymentSplitItem.md) | Payment splits for the subscription. | [optional] 
**SubscriptionStatus** | Pointer to **string** | Status of the subscription. | [optional] 
**SubscriptionTags** | Pointer to **map[string]interface{}** | Tags for the subscription. | [optional] 

## Methods

### NewSubscriptionEntity

`func NewSubscriptionEntity() *SubscriptionEntity`

NewSubscriptionEntity instantiates a new SubscriptionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEntityWithDefaults

`func NewSubscriptionEntityWithDefaults() *SubscriptionEntity`

NewSubscriptionEntityWithDefaults instantiates a new SubscriptionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorisationDetails

`func (o *SubscriptionEntity) GetAuthorisationDetails() AuthorizationDetails`

GetAuthorisationDetails returns the AuthorisationDetails field if non-nil, zero value otherwise.

### GetAuthorisationDetailsOk

`func (o *SubscriptionEntity) GetAuthorisationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorisationDetailsOk returns a tuple with the AuthorisationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationDetails

`func (o *SubscriptionEntity) SetAuthorisationDetails(v AuthorizationDetails)`

SetAuthorisationDetails sets AuthorisationDetails field to given value.

### HasAuthorisationDetails

`func (o *SubscriptionEntity) HasAuthorisationDetails() bool`

HasAuthorisationDetails returns a boolean if a field has been set.

### GetCfSubscriptionId

`func (o *SubscriptionEntity) GetCfSubscriptionId() string`

GetCfSubscriptionId returns the CfSubscriptionId field if non-nil, zero value otherwise.

### GetCfSubscriptionIdOk

`func (o *SubscriptionEntity) GetCfSubscriptionIdOk() (*string, bool)`

GetCfSubscriptionIdOk returns a tuple with the CfSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSubscriptionId

`func (o *SubscriptionEntity) SetCfSubscriptionId(v string)`

SetCfSubscriptionId sets CfSubscriptionId field to given value.

### HasCfSubscriptionId

`func (o *SubscriptionEntity) HasCfSubscriptionId() bool`

HasCfSubscriptionId returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *SubscriptionEntity) GetCustomerDetails() SubscriptionCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *SubscriptionEntity) GetCustomerDetailsOk() (*SubscriptionCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *SubscriptionEntity) SetCustomerDetails(v SubscriptionCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *SubscriptionEntity) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetPlanDetails

`func (o *SubscriptionEntity) GetPlanDetails() PlanEntity`

GetPlanDetails returns the PlanDetails field if non-nil, zero value otherwise.

### GetPlanDetailsOk

`func (o *SubscriptionEntity) GetPlanDetailsOk() (*PlanEntity, bool)`

GetPlanDetailsOk returns a tuple with the PlanDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDetails

`func (o *SubscriptionEntity) SetPlanDetails(v PlanEntity)`

SetPlanDetails sets PlanDetails field to given value.

### HasPlanDetails

`func (o *SubscriptionEntity) HasPlanDetails() bool`

HasPlanDetails returns a boolean if a field has been set.

### GetSubscriptionExpiryTime

`func (o *SubscriptionEntity) GetSubscriptionExpiryTime() string`

GetSubscriptionExpiryTime returns the SubscriptionExpiryTime field if non-nil, zero value otherwise.

### GetSubscriptionExpiryTimeOk

`func (o *SubscriptionEntity) GetSubscriptionExpiryTimeOk() (*string, bool)`

GetSubscriptionExpiryTimeOk returns a tuple with the SubscriptionExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExpiryTime

`func (o *SubscriptionEntity) SetSubscriptionExpiryTime(v string)`

SetSubscriptionExpiryTime sets SubscriptionExpiryTime field to given value.

### HasSubscriptionExpiryTime

`func (o *SubscriptionEntity) HasSubscriptionExpiryTime() bool`

HasSubscriptionExpiryTime returns a boolean if a field has been set.

### GetSubscriptionFirstChargeTime

`func (o *SubscriptionEntity) GetSubscriptionFirstChargeTime() string`

GetSubscriptionFirstChargeTime returns the SubscriptionFirstChargeTime field if non-nil, zero value otherwise.

### GetSubscriptionFirstChargeTimeOk

`func (o *SubscriptionEntity) GetSubscriptionFirstChargeTimeOk() (*string, bool)`

GetSubscriptionFirstChargeTimeOk returns a tuple with the SubscriptionFirstChargeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFirstChargeTime

`func (o *SubscriptionEntity) SetSubscriptionFirstChargeTime(v string)`

SetSubscriptionFirstChargeTime sets SubscriptionFirstChargeTime field to given value.

### HasSubscriptionFirstChargeTime

`func (o *SubscriptionEntity) HasSubscriptionFirstChargeTime() bool`

HasSubscriptionFirstChargeTime returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionEntity) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionEntity) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionEntity) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SubscriptionEntity) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionMeta

`func (o *SubscriptionEntity) GetSubscriptionMeta() SubscriptionEntitySubscriptionMeta`

GetSubscriptionMeta returns the SubscriptionMeta field if non-nil, zero value otherwise.

### GetSubscriptionMetaOk

`func (o *SubscriptionEntity) GetSubscriptionMetaOk() (*SubscriptionEntitySubscriptionMeta, bool)`

GetSubscriptionMetaOk returns a tuple with the SubscriptionMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionMeta

`func (o *SubscriptionEntity) SetSubscriptionMeta(v SubscriptionEntitySubscriptionMeta)`

SetSubscriptionMeta sets SubscriptionMeta field to given value.

### HasSubscriptionMeta

`func (o *SubscriptionEntity) HasSubscriptionMeta() bool`

HasSubscriptionMeta returns a boolean if a field has been set.

### GetSubscriptionNote

`func (o *SubscriptionEntity) GetSubscriptionNote() string`

GetSubscriptionNote returns the SubscriptionNote field if non-nil, zero value otherwise.

### GetSubscriptionNoteOk

`func (o *SubscriptionEntity) GetSubscriptionNoteOk() (*string, bool)`

GetSubscriptionNoteOk returns a tuple with the SubscriptionNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionNote

`func (o *SubscriptionEntity) SetSubscriptionNote(v string)`

SetSubscriptionNote sets SubscriptionNote field to given value.

### HasSubscriptionNote

`func (o *SubscriptionEntity) HasSubscriptionNote() bool`

HasSubscriptionNote returns a boolean if a field has been set.

### GetSubscriptionPaymentSplits

`func (o *SubscriptionEntity) GetSubscriptionPaymentSplits() []SubscriptionPaymentSplitItem`

GetSubscriptionPaymentSplits returns the SubscriptionPaymentSplits field if non-nil, zero value otherwise.

### GetSubscriptionPaymentSplitsOk

`func (o *SubscriptionEntity) GetSubscriptionPaymentSplitsOk() (*[]SubscriptionPaymentSplitItem, bool)`

GetSubscriptionPaymentSplitsOk returns a tuple with the SubscriptionPaymentSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPaymentSplits

`func (o *SubscriptionEntity) SetSubscriptionPaymentSplits(v []SubscriptionPaymentSplitItem)`

SetSubscriptionPaymentSplits sets SubscriptionPaymentSplits field to given value.

### HasSubscriptionPaymentSplits

`func (o *SubscriptionEntity) HasSubscriptionPaymentSplits() bool`

HasSubscriptionPaymentSplits returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *SubscriptionEntity) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *SubscriptionEntity) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *SubscriptionEntity) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *SubscriptionEntity) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetSubscriptionTags

`func (o *SubscriptionEntity) GetSubscriptionTags() map[string]interface{}`

GetSubscriptionTags returns the SubscriptionTags field if non-nil, zero value otherwise.

### GetSubscriptionTagsOk

`func (o *SubscriptionEntity) GetSubscriptionTagsOk() (*map[string]interface{}, bool)`

GetSubscriptionTagsOk returns a tuple with the SubscriptionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTags

`func (o *SubscriptionEntity) SetSubscriptionTags(v map[string]interface{})`

SetSubscriptionTags sets SubscriptionTags field to given value.

### HasSubscriptionTags

`func (o *SubscriptionEntity) HasSubscriptionTags() bool`

HasSubscriptionTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


