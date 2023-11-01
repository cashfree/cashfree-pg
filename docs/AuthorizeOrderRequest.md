# AuthorizeOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of authorization to run. Can be one of &#39;CAPTURE&#39; , &#39;VOID&#39; | [optional] 
**Amount** | Pointer to **float32** | The amount if you are running a &#39;CAPTURE&#39; | [optional] 

## Methods

### NewAuthorizeOrderRequest

`func NewAuthorizeOrderRequest() *AuthorizeOrderRequest`

NewAuthorizeOrderRequest instantiates a new AuthorizeOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeOrderRequestWithDefaults

`func NewAuthorizeOrderRequestWithDefaults() *AuthorizeOrderRequest`

NewAuthorizeOrderRequestWithDefaults instantiates a new AuthorizeOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuthorizeOrderRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizeOrderRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizeOrderRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizeOrderRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAmount

`func (o *AuthorizeOrderRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthorizeOrderRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthorizeOrderRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AuthorizeOrderRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


