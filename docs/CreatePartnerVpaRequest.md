# CreatePartnerVpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpaCount** | **int32** | count of vpa , to create in bulk, max limit:50 | 

## Methods

### NewCreatePartnerVpaRequest

`func NewCreatePartnerVpaRequest(vpaCount int32, ) *CreatePartnerVpaRequest`

NewCreatePartnerVpaRequest instantiates a new CreatePartnerVpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePartnerVpaRequestWithDefaults

`func NewCreatePartnerVpaRequestWithDefaults() *CreatePartnerVpaRequest`

NewCreatePartnerVpaRequestWithDefaults instantiates a new CreatePartnerVpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpaCount

`func (o *CreatePartnerVpaRequest) GetVpaCount() int32`

GetVpaCount returns the VpaCount field if non-nil, zero value otherwise.

### GetVpaCountOk

`func (o *CreatePartnerVpaRequest) GetVpaCountOk() (*int32, bool)`

GetVpaCountOk returns a tuple with the VpaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpaCount

`func (o *CreatePartnerVpaRequest) SetVpaCount(v int32)`

SetVpaCount sets VpaCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


