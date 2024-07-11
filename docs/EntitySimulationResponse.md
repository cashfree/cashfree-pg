# EntitySimulationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentStatus** | **string** | Payment Status | 
**PaymentErrorCode** | Pointer to **string** | Payment Error Code | [optional] 

## Methods

### NewEntitySimulationResponse

`func NewEntitySimulationResponse(paymentStatus string, ) *EntitySimulationResponse`

NewEntitySimulationResponse instantiates a new EntitySimulationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySimulationResponseWithDefaults

`func NewEntitySimulationResponseWithDefaults() *EntitySimulationResponse`

NewEntitySimulationResponseWithDefaults instantiates a new EntitySimulationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentStatus

`func (o *EntitySimulationResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *EntitySimulationResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *EntitySimulationResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetPaymentErrorCode

`func (o *EntitySimulationResponse) GetPaymentErrorCode() string`

GetPaymentErrorCode returns the PaymentErrorCode field if non-nil, zero value otherwise.

### GetPaymentErrorCodeOk

`func (o *EntitySimulationResponse) GetPaymentErrorCodeOk() (*string, bool)`

GetPaymentErrorCodeOk returns a tuple with the PaymentErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentErrorCode

`func (o *EntitySimulationResponse) SetPaymentErrorCode(v string)`

SetPaymentErrorCode sets PaymentErrorCode field to given value.

### HasPaymentErrorCode

`func (o *EntitySimulationResponse) HasPaymentErrorCode() bool`

HasPaymentErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


