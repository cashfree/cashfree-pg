# UploadPnachImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** | The payment_id against which the pnach image is uploaded. | [optional] 
**AuthorizationStatus** | Pointer to **string** | Authorization status of the subscription. | [optional] 
**Action** | Pointer to **string** | Action performed on the file. | [optional] 
**PaymentMessage** | Pointer to **string** | Message of the API. | [optional] 

## Methods

### NewUploadPnachImageResponse

`func NewUploadPnachImageResponse() *UploadPnachImageResponse`

NewUploadPnachImageResponse instantiates a new UploadPnachImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPnachImageResponseWithDefaults

`func NewUploadPnachImageResponseWithDefaults() *UploadPnachImageResponse`

NewUploadPnachImageResponseWithDefaults instantiates a new UploadPnachImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *UploadPnachImageResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UploadPnachImageResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UploadPnachImageResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UploadPnachImageResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetAuthorizationStatus

`func (o *UploadPnachImageResponse) GetAuthorizationStatus() string`

GetAuthorizationStatus returns the AuthorizationStatus field if non-nil, zero value otherwise.

### GetAuthorizationStatusOk

`func (o *UploadPnachImageResponse) GetAuthorizationStatusOk() (*string, bool)`

GetAuthorizationStatusOk returns a tuple with the AuthorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationStatus

`func (o *UploadPnachImageResponse) SetAuthorizationStatus(v string)`

SetAuthorizationStatus sets AuthorizationStatus field to given value.

### HasAuthorizationStatus

`func (o *UploadPnachImageResponse) HasAuthorizationStatus() bool`

HasAuthorizationStatus returns a boolean if a field has been set.

### GetAction

`func (o *UploadPnachImageResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UploadPnachImageResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UploadPnachImageResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UploadPnachImageResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPaymentMessage

`func (o *UploadPnachImageResponse) GetPaymentMessage() string`

GetPaymentMessage returns the PaymentMessage field if non-nil, zero value otherwise.

### GetPaymentMessageOk

`func (o *UploadPnachImageResponse) GetPaymentMessageOk() (*string, bool)`

GetPaymentMessageOk returns a tuple with the PaymentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMessage

`func (o *UploadPnachImageResponse) SetPaymentMessage(v string)`

SetPaymentMessage sets PaymentMessage field to given value.

### HasPaymentMessage

`func (o *UploadPnachImageResponse) HasPaymentMessage() bool`

HasPaymentMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


