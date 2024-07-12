# EntitySimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentStatus** | **string** | Payment Status | 
**PaymentErrorCode** | Pointer to **string** | Payment Error Code | [optional] 

## Methods

### NewEntitySimulationRequest

`func NewEntitySimulationRequest(paymentStatus string, ) *EntitySimulationRequest`

NewEntitySimulationRequest instantiates a new EntitySimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySimulationRequestWithDefaults

`func NewEntitySimulationRequestWithDefaults() *EntitySimulationRequest`

NewEntitySimulationRequestWithDefaults instantiates a new EntitySimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentStatus

`func (o *EntitySimulationRequest) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *EntitySimulationRequest) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *EntitySimulationRequest) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetPaymentErrorCode

`func (o *EntitySimulationRequest) GetPaymentErrorCode() string`

GetPaymentErrorCode returns the PaymentErrorCode field if non-nil, zero value otherwise.

### GetPaymentErrorCodeOk

`func (o *EntitySimulationRequest) GetPaymentErrorCodeOk() (*string, bool)`

GetPaymentErrorCodeOk returns a tuple with the PaymentErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentErrorCode

`func (o *EntitySimulationRequest) SetPaymentErrorCode(v string)`

SetPaymentErrorCode sets PaymentErrorCode field to given value.

### HasPaymentErrorCode

`func (o *EntitySimulationRequest) HasPaymentErrorCode() bool`

HasPaymentErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


