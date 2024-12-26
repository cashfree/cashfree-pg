# DisputesEntityMerchantAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisputeId** | Pointer to **int32** |  | [optional] 
**DisputeType** | Pointer to **string** |  | [optional] 
**ReasonCode** | Pointer to **string** |  | [optional] 
**ReasonDescription** | Pointer to **string** |  | [optional] 
**DisputeAmount** | Pointer to **float32** | Dispute amount may differ from transaction amount for partial cases. | [optional] 
**CreatedAt** | Pointer to **string** | This is the time when the dispute was created. | [optional] 
**RespondBy** | Pointer to **string** | This is the time by which evidence should be submitted to contest the dispute. | [optional] 
**UpdatedAt** | Pointer to **string** | This is the time when the dispute case was updated. | [optional] 
**ResolvedAt** | Pointer to **string** | This is the time when the dispute case was closed. | [optional] 
**DisputeStatus** | Pointer to **string** |  | [optional] 
**CfDisputeRemarks** | Pointer to **string** |  | [optional] 
**PreferredEvidence** | Pointer to  |  | [optional] 
**DisputeEvidence** | Pointer to  |  | [optional] 
**OrderDetails** | Pointer to [**OrderDetailsInDisputesEntity**](OrderDetailsInDisputesEntity.md) |  | [optional] 
**CustomerDetails** | Pointer to [**CustomerDetailsInDisputesEntity**](CustomerDetailsInDisputesEntity.md) |  | [optional] 

## Methods

### NewDisputesEntityMerchantAccepted

`func NewDisputesEntityMerchantAccepted() *DisputesEntityMerchantAccepted`

NewDisputesEntityMerchantAccepted instantiates a new DisputesEntityMerchantAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputesEntityMerchantAcceptedWithDefaults

`func NewDisputesEntityMerchantAcceptedWithDefaults() *DisputesEntityMerchantAccepted`

NewDisputesEntityMerchantAcceptedWithDefaults instantiates a new DisputesEntityMerchantAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputeId

`func (o *DisputesEntityMerchantAccepted) GetDisputeId() int32`

GetDisputeId returns the DisputeId field if non-nil, zero value otherwise.

### GetDisputeIdOk

`func (o *DisputesEntityMerchantAccepted) GetDisputeIdOk() (*int32, bool)`

GetDisputeIdOk returns a tuple with the DisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeId

`func (o *DisputesEntityMerchantAccepted) SetDisputeId(v int32)`

SetDisputeId sets DisputeId field to given value.

### HasDisputeId

`func (o *DisputesEntityMerchantAccepted) HasDisputeId() bool`

HasDisputeId returns a boolean if a field has been set.

### GetDisputeType

`func (o *DisputesEntityMerchantAccepted) GetDisputeType() string`

GetDisputeType returns the DisputeType field if non-nil, zero value otherwise.

### GetDisputeTypeOk

`func (o *DisputesEntityMerchantAccepted) GetDisputeTypeOk() (*string, bool)`

GetDisputeTypeOk returns a tuple with the DisputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeType

`func (o *DisputesEntityMerchantAccepted) SetDisputeType(v string)`

SetDisputeType sets DisputeType field to given value.

### HasDisputeType

`func (o *DisputesEntityMerchantAccepted) HasDisputeType() bool`

HasDisputeType returns a boolean if a field has been set.

### GetReasonCode

`func (o *DisputesEntityMerchantAccepted) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DisputesEntityMerchantAccepted) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DisputesEntityMerchantAccepted) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DisputesEntityMerchantAccepted) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReasonDescription

`func (o *DisputesEntityMerchantAccepted) GetReasonDescription() string`

GetReasonDescription returns the ReasonDescription field if non-nil, zero value otherwise.

### GetReasonDescriptionOk

`func (o *DisputesEntityMerchantAccepted) GetReasonDescriptionOk() (*string, bool)`

GetReasonDescriptionOk returns a tuple with the ReasonDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDescription

`func (o *DisputesEntityMerchantAccepted) SetReasonDescription(v string)`

SetReasonDescription sets ReasonDescription field to given value.

### HasReasonDescription

`func (o *DisputesEntityMerchantAccepted) HasReasonDescription() bool`

HasReasonDescription returns a boolean if a field has been set.

### GetDisputeAmount

`func (o *DisputesEntityMerchantAccepted) GetDisputeAmount() float32`

GetDisputeAmount returns the DisputeAmount field if non-nil, zero value otherwise.

### GetDisputeAmountOk

`func (o *DisputesEntityMerchantAccepted) GetDisputeAmountOk() (*float32, bool)`

GetDisputeAmountOk returns a tuple with the DisputeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeAmount

`func (o *DisputesEntityMerchantAccepted) SetDisputeAmount(v float32)`

SetDisputeAmount sets DisputeAmount field to given value.

### HasDisputeAmount

`func (o *DisputesEntityMerchantAccepted) HasDisputeAmount() bool`

HasDisputeAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DisputesEntityMerchantAccepted) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DisputesEntityMerchantAccepted) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DisputesEntityMerchantAccepted) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DisputesEntityMerchantAccepted) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRespondBy

`func (o *DisputesEntityMerchantAccepted) GetRespondBy() string`

GetRespondBy returns the RespondBy field if non-nil, zero value otherwise.

### GetRespondByOk

`func (o *DisputesEntityMerchantAccepted) GetRespondByOk() (*string, bool)`

GetRespondByOk returns a tuple with the RespondBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondBy

`func (o *DisputesEntityMerchantAccepted) SetRespondBy(v string)`

SetRespondBy sets RespondBy field to given value.

### HasRespondBy

`func (o *DisputesEntityMerchantAccepted) HasRespondBy() bool`

HasRespondBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DisputesEntityMerchantAccepted) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DisputesEntityMerchantAccepted) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DisputesEntityMerchantAccepted) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DisputesEntityMerchantAccepted) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetResolvedAt

`func (o *DisputesEntityMerchantAccepted) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DisputesEntityMerchantAccepted) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DisputesEntityMerchantAccepted) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DisputesEntityMerchantAccepted) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetDisputeStatus

`func (o *DisputesEntityMerchantAccepted) GetDisputeStatus() string`

GetDisputeStatus returns the DisputeStatus field if non-nil, zero value otherwise.

### GetDisputeStatusOk

`func (o *DisputesEntityMerchantAccepted) GetDisputeStatusOk() (*string, bool)`

GetDisputeStatusOk returns a tuple with the DisputeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeStatus

`func (o *DisputesEntityMerchantAccepted) SetDisputeStatus(v string)`

SetDisputeStatus sets DisputeStatus field to given value.

### HasDisputeStatus

`func (o *DisputesEntityMerchantAccepted) HasDisputeStatus() bool`

HasDisputeStatus returns a boolean if a field has been set.

### GetCfDisputeRemarks

`func (o *DisputesEntityMerchantAccepted) GetCfDisputeRemarks() string`

GetCfDisputeRemarks returns the CfDisputeRemarks field if non-nil, zero value otherwise.

### GetCfDisputeRemarksOk

`func (o *DisputesEntityMerchantAccepted) GetCfDisputeRemarksOk() (*string, bool)`

GetCfDisputeRemarksOk returns a tuple with the CfDisputeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfDisputeRemarks

`func (o *DisputesEntityMerchantAccepted) SetCfDisputeRemarks(v string)`

SetCfDisputeRemarks sets CfDisputeRemarks field to given value.

### HasCfDisputeRemarks

`func (o *DisputesEntityMerchantAccepted) HasCfDisputeRemarks() bool`

HasCfDisputeRemarks returns a boolean if a field has been set.

### GetPreferredEvidence

`func (o *DisputesEntityMerchantAccepted) GetPreferredEvidence() []EvidencesToContestDispute`

GetPreferredEvidence returns the PreferredEvidence field if non-nil, zero value otherwise.

### GetPreferredEvidenceOk

`func (o *DisputesEntityMerchantAccepted) GetPreferredEvidenceOk() (*[]EvidencesToContestDispute, bool)`

GetPreferredEvidenceOk returns a tuple with the PreferredEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEvidence

`func (o *DisputesEntityMerchantAccepted) SetPreferredEvidence(v []EvidencesToContestDispute)`

SetPreferredEvidence sets PreferredEvidence field to given value.

### HasPreferredEvidence

`func (o *DisputesEntityMerchantAccepted) HasPreferredEvidence() bool`

HasPreferredEvidence returns a boolean if a field has been set.

### GetDisputeEvidence

`func (o *DisputesEntityMerchantAccepted) GetDisputeEvidence() []Evidence`

GetDisputeEvidence returns the DisputeEvidence field if non-nil, zero value otherwise.

### GetDisputeEvidenceOk

`func (o *DisputesEntityMerchantAccepted) GetDisputeEvidenceOk() (*[]Evidence, bool)`

GetDisputeEvidenceOk returns a tuple with the DisputeEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeEvidence

`func (o *DisputesEntityMerchantAccepted) SetDisputeEvidence(v []Evidence)`

SetDisputeEvidence sets DisputeEvidence field to given value.

### HasDisputeEvidence

`func (o *DisputesEntityMerchantAccepted) HasDisputeEvidence() bool`

HasDisputeEvidence returns a boolean if a field has been set.

### GetOrderDetails

`func (o *DisputesEntityMerchantAccepted) GetOrderDetails() OrderDetailsInDisputesEntity`

GetOrderDetails returns the OrderDetails field if non-nil, zero value otherwise.

### GetOrderDetailsOk

`func (o *DisputesEntityMerchantAccepted) GetOrderDetailsOk() (*OrderDetailsInDisputesEntity, bool)`

GetOrderDetailsOk returns a tuple with the OrderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetails

`func (o *DisputesEntityMerchantAccepted) SetOrderDetails(v OrderDetailsInDisputesEntity)`

SetOrderDetails sets OrderDetails field to given value.

### HasOrderDetails

`func (o *DisputesEntityMerchantAccepted) HasOrderDetails() bool`

HasOrderDetails returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *DisputesEntityMerchantAccepted) GetCustomerDetails() CustomerDetailsInDisputesEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *DisputesEntityMerchantAccepted) GetCustomerDetailsOk() (*CustomerDetailsInDisputesEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *DisputesEntityMerchantAccepted) SetCustomerDetails(v CustomerDetailsInDisputesEntity)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *DisputesEntityMerchantAccepted) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


