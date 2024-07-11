# ManageSubscriptionPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The unique ID which was used to create subscription. | 
**PaymentId** | **string** | The unique ID which was used to create payment. | 
**Action** | **string** | Action to be performed on the payment. Possible values - CANCEL, RETRY. | 
**ActionDetails** | Pointer to [**ManageSubscriptionPaymentRequestActionDetails**](ManageSubscriptionPaymentRequestActionDetails.md) |  | [optional] 

## Methods

### NewManageSubscriptionPaymentRequest

`func NewManageSubscriptionPaymentRequest(subscriptionId string, paymentId string, action string, ) *ManageSubscriptionPaymentRequest`

NewManageSubscriptionPaymentRequest instantiates a new ManageSubscriptionPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageSubscriptionPaymentRequestWithDefaults

`func NewManageSubscriptionPaymentRequestWithDefaults() *ManageSubscriptionPaymentRequest`

NewManageSubscriptionPaymentRequestWithDefaults instantiates a new ManageSubscriptionPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *ManageSubscriptionPaymentRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ManageSubscriptionPaymentRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ManageSubscriptionPaymentRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPaymentId

`func (o *ManageSubscriptionPaymentRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ManageSubscriptionPaymentRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ManageSubscriptionPaymentRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAction

`func (o *ManageSubscriptionPaymentRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageSubscriptionPaymentRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageSubscriptionPaymentRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionDetails

`func (o *ManageSubscriptionPaymentRequest) GetActionDetails() ManageSubscriptionPaymentRequestActionDetails`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *ManageSubscriptionPaymentRequest) GetActionDetailsOk() (*ManageSubscriptionPaymentRequestActionDetails, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *ManageSubscriptionPaymentRequest) SetActionDetails(v ManageSubscriptionPaymentRequestActionDetails)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *ManageSubscriptionPaymentRequest) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


