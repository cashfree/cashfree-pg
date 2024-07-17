# CreateSubscriptionPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** | Cashfree subscription payment reference number | [optional] 
**FailureDetails** | Pointer to [**SubscriptionPaymentEntityFailureDetails**](SubscriptionPaymentEntityFailureDetails.md) |  | [optional] 
**PaymentAmount** | Pointer to **float32** | The charge amount of the payment. | [optional] 
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**PaymentInitiatedDate** | Pointer to **string** | The date on which the payment was initiated. | [optional] 
**PaymentStatus** | Pointer to **string** | Status of the payment. | [optional] 
**PaymentType** | Pointer to **string** | Payment type. Can be AUTH or CHARGE. | [optional] 
**SubscriptionId** | Pointer to **string** | A unique ID passed by merchant for identifying the subscription. | [optional] 
**Data** | Pointer to **map[string]interface{}** | Contains a payload for auth app links in case of AUTH. For charge, the payload is empty. | [optional] 
**PaymentMethod** | Pointer to **string** | Payment method used for the authorization. | [optional] 

## Methods

### NewCreateSubscriptionPaymentResponse

`func NewCreateSubscriptionPaymentResponse() *CreateSubscriptionPaymentResponse`

NewCreateSubscriptionPaymentResponse instantiates a new CreateSubscriptionPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentResponseWithDefaults

`func NewCreateSubscriptionPaymentResponseWithDefaults() *CreateSubscriptionPaymentResponse`

NewCreateSubscriptionPaymentResponseWithDefaults instantiates a new CreateSubscriptionPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CreateSubscriptionPaymentResponse) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CreateSubscriptionPaymentResponse) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CreateSubscriptionPaymentResponse) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CreateSubscriptionPaymentResponse) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetFailureDetails

`func (o *CreateSubscriptionPaymentResponse) GetFailureDetails() SubscriptionPaymentEntityFailureDetails`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *CreateSubscriptionPaymentResponse) GetFailureDetailsOk() (*SubscriptionPaymentEntityFailureDetails, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *CreateSubscriptionPaymentResponse) SetFailureDetails(v SubscriptionPaymentEntityFailureDetails)`

SetFailureDetails sets FailureDetails field to given value.

### HasFailureDetails

`func (o *CreateSubscriptionPaymentResponse) HasFailureDetails() bool`

HasFailureDetails returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *CreateSubscriptionPaymentResponse) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *CreateSubscriptionPaymentResponse) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *CreateSubscriptionPaymentResponse) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *CreateSubscriptionPaymentResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateSubscriptionPaymentResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *CreateSubscriptionPaymentResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentResponse) GetPaymentInitiatedDate() string`

GetPaymentInitiatedDate returns the PaymentInitiatedDate field if non-nil, zero value otherwise.

### GetPaymentInitiatedDateOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentInitiatedDateOk() (*string, bool)`

GetPaymentInitiatedDateOk returns a tuple with the PaymentInitiatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentResponse) SetPaymentInitiatedDate(v string)`

SetPaymentInitiatedDate sets PaymentInitiatedDate field to given value.

### HasPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentResponse) HasPaymentInitiatedDate() bool`

HasPaymentInitiatedDate returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *CreateSubscriptionPaymentResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CreateSubscriptionPaymentResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *CreateSubscriptionPaymentResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentType

`func (o *CreateSubscriptionPaymentResponse) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CreateSubscriptionPaymentResponse) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *CreateSubscriptionPaymentResponse) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreateSubscriptionPaymentResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionPaymentResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionPaymentResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateSubscriptionPaymentResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetData

`func (o *CreateSubscriptionPaymentResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateSubscriptionPaymentResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateSubscriptionPaymentResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateSubscriptionPaymentResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreateSubscriptionPaymentResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateSubscriptionPaymentResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateSubscriptionPaymentResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreateSubscriptionPaymentResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


