# PayOrderAuthorizeOnlyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSessionId** | **string** |  | 
**AuthorizationData** | [**PayOrderAuthorizeOnlyRequestAuthorizationData**](PayOrderAuthorizeOnlyRequestAuthorizationData.md) |  | 

## Methods

### NewPayOrderAuthorizeOnlyRequest

`func NewPayOrderAuthorizeOnlyRequest(paymentSessionId string, authorizationData PayOrderAuthorizeOnlyRequestAuthorizationData, ) *PayOrderAuthorizeOnlyRequest`

NewPayOrderAuthorizeOnlyRequest instantiates a new PayOrderAuthorizeOnlyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderAuthorizeOnlyRequestWithDefaults

`func NewPayOrderAuthorizeOnlyRequestWithDefaults() *PayOrderAuthorizeOnlyRequest`

NewPayOrderAuthorizeOnlyRequestWithDefaults instantiates a new PayOrderAuthorizeOnlyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSessionId

`func (o *PayOrderAuthorizeOnlyRequest) GetPaymentSessionId() string`

GetPaymentSessionId returns the PaymentSessionId field if non-nil, zero value otherwise.

### GetPaymentSessionIdOk

`func (o *PayOrderAuthorizeOnlyRequest) GetPaymentSessionIdOk() (*string, bool)`

GetPaymentSessionIdOk returns a tuple with the PaymentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSessionId

`func (o *PayOrderAuthorizeOnlyRequest) SetPaymentSessionId(v string)`

SetPaymentSessionId sets PaymentSessionId field to given value.


### GetAuthorizationData

`func (o *PayOrderAuthorizeOnlyRequest) GetAuthorizationData() PayOrderAuthorizeOnlyRequestAuthorizationData`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *PayOrderAuthorizeOnlyRequest) GetAuthorizationDataOk() (*PayOrderAuthorizeOnlyRequestAuthorizationData, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *PayOrderAuthorizeOnlyRequest) SetAuthorizationData(v PayOrderAuthorizeOnlyRequestAuthorizationData)`

SetAuthorizationData sets AuthorizationData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


