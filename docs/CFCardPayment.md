# CFCardPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | [**CFCard**](CFCard.md) |  | 

## Methods

### NewCFCardPayment

`func NewCFCardPayment(card CFCard, ) *CFCardPayment`

NewCFCardPayment instantiates a new CFCardPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCardPaymentWithDefaults

`func NewCFCardPaymentWithDefaults() *CFCardPayment`

NewCFCardPaymentWithDefaults instantiates a new CFCardPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *CFCardPayment) GetCard() CFCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CFCardPayment) GetCardOk() (*CFCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CFCardPayment) SetCard(v CFCard)`

SetCard sets Card field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


