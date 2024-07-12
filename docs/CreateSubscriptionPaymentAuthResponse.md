# CreateSubscriptionPaymentAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **string** | Cashfree subscription payment reference number | [optional] 
**CfSubscriptionId** | Pointer to **string** | Cashfree subscription reference number | [optional] 
**FailureDetails** | Pointer to [**CreateSubscriptionPaymentAuthResponseFailureDetails**](CreateSubscriptionPaymentAuthResponseFailureDetails.md) |  | [optional] 
**PaymentAmount** | Pointer to **float32** | The charge amount of the payment. | [optional] 
**PaymentId** | Pointer to **string** | A unique ID passed by merchant for identifying the transaction. | [optional] 
**PaymentInitiatedDate** | Pointer to **string** | The date on which the payment was initiated. | [optional] 
**PaymentStatus** | Pointer to **string** | Status of the payment. | [optional] 
**PaymentType** | Pointer to **string** | Payment type. Can be AUTH or CHARGE. | [optional] 
**SubscriptionId** | Pointer to **string** | A unique ID passed by merchant for identifying the subscription. | [optional] 
**PaymentMethod** | Pointer to **string** | Payment method used for the authorization. | [optional] 

## Methods

### NewCreateSubscriptionPaymentAuthResponse

`func NewCreateSubscriptionPaymentAuthResponse() *CreateSubscriptionPaymentAuthResponse`

NewCreateSubscriptionPaymentAuthResponse instantiates a new CreateSubscriptionPaymentAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPaymentAuthResponseWithDefaults

`func NewCreateSubscriptionPaymentAuthResponseWithDefaults() *CreateSubscriptionPaymentAuthResponse`

NewCreateSubscriptionPaymentAuthResponseWithDefaults instantiates a new CreateSubscriptionPaymentAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) SetCfPaymentId(v string)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetCfSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) GetCfSubscriptionId() string`

GetCfSubscriptionId returns the CfSubscriptionId field if non-nil, zero value otherwise.

### GetCfSubscriptionIdOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetCfSubscriptionIdOk() (*string, bool)`

GetCfSubscriptionIdOk returns a tuple with the CfSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) SetCfSubscriptionId(v string)`

SetCfSubscriptionId sets CfSubscriptionId field to given value.

### HasCfSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) HasCfSubscriptionId() bool`

HasCfSubscriptionId returns a boolean if a field has been set.

### GetFailureDetails

`func (o *CreateSubscriptionPaymentAuthResponse) GetFailureDetails() CreateSubscriptionPaymentAuthResponseFailureDetails`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetFailureDetailsOk() (*CreateSubscriptionPaymentAuthResponseFailureDetails, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *CreateSubscriptionPaymentAuthResponse) SetFailureDetails(v CreateSubscriptionPaymentAuthResponseFailureDetails)`

SetFailureDetails sets FailureDetails field to given value.

### HasFailureDetails

`func (o *CreateSubscriptionPaymentAuthResponse) HasFailureDetails() bool`

HasFailureDetails returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentInitiatedDate() string`

GetPaymentInitiatedDate returns the PaymentInitiatedDate field if non-nil, zero value otherwise.

### GetPaymentInitiatedDateOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentInitiatedDateOk() (*string, bool)`

GetPaymentInitiatedDateOk returns a tuple with the PaymentInitiatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentInitiatedDate(v string)`

SetPaymentInitiatedDate sets PaymentInitiatedDate field to given value.

### HasPaymentInitiatedDate

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentInitiatedDate() bool`

HasPaymentInitiatedDate returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentType

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateSubscriptionPaymentAuthResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateSubscriptionPaymentAuthResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateSubscriptionPaymentAuthResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreateSubscriptionPaymentAuthResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


