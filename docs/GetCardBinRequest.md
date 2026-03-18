# GetCardBinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | Card number, minimum first 8 digits are required. | 

## Methods

### NewGetCardBinRequest

`func NewGetCardBinRequest(cardNumber string, ) *GetCardBinRequest`

NewGetCardBinRequest instantiates a new GetCardBinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCardBinRequestWithDefaults

`func NewGetCardBinRequestWithDefaults() *GetCardBinRequest`

NewGetCardBinRequestWithDefaults instantiates a new GetCardBinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *GetCardBinRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *GetCardBinRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *GetCardBinRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


