# CFAuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of authorization to run. Can be one of &#39;CAPTURE&#39; , &#39;VOID&#39; | [optional] 
**Amount** | Pointer to **float32** | The amount if you are running a &#39;CAPTURE&#39; | [optional] 

## Methods

### NewCFAuthorizationRequest

`func NewCFAuthorizationRequest() *CFAuthorizationRequest`

NewCFAuthorizationRequest instantiates a new CFAuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFAuthorizationRequestWithDefaults

`func NewCFAuthorizationRequestWithDefaults() *CFAuthorizationRequest`

NewCFAuthorizationRequestWithDefaults instantiates a new CFAuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CFAuthorizationRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CFAuthorizationRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CFAuthorizationRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CFAuthorizationRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAmount

`func (o *CFAuthorizationRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CFAuthorizationRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CFAuthorizationRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CFAuthorizationRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


