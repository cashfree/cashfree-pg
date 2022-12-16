# CFOrderPayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderToken** | **string** |  | 
**PaymentMethod** | [**CFPaymentMethod**](CFPaymentMethod.md) |  | 
**SaveInstrument** | Pointer to **bool** |  | [optional] 

## Methods

### NewCFOrderPayRequest

`func NewCFOrderPayRequest(orderToken string, paymentMethod CFPaymentMethod, ) *CFOrderPayRequest`

NewCFOrderPayRequest instantiates a new CFOrderPayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFOrderPayRequestWithDefaults

`func NewCFOrderPayRequestWithDefaults() *CFOrderPayRequest`

NewCFOrderPayRequestWithDefaults instantiates a new CFOrderPayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderToken

`func (o *CFOrderPayRequest) GetOrderToken() string`

GetOrderToken returns the OrderToken field if non-nil, zero value otherwise.

### GetOrderTokenOk

`func (o *CFOrderPayRequest) GetOrderTokenOk() (*string, bool)`

GetOrderTokenOk returns a tuple with the OrderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderToken

`func (o *CFOrderPayRequest) SetOrderToken(v string)`

SetOrderToken sets OrderToken field to given value.


### GetPaymentMethod

`func (o *CFOrderPayRequest) GetPaymentMethod() CFPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CFOrderPayRequest) GetPaymentMethodOk() (*CFPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CFOrderPayRequest) SetPaymentMethod(v CFPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetSaveInstrument

`func (o *CFOrderPayRequest) GetSaveInstrument() bool`

GetSaveInstrument returns the SaveInstrument field if non-nil, zero value otherwise.

### GetSaveInstrumentOk

`func (o *CFOrderPayRequest) GetSaveInstrumentOk() (*bool, bool)`

GetSaveInstrumentOk returns a tuple with the SaveInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveInstrument

`func (o *CFOrderPayRequest) SetSaveInstrument(v bool)`

SetSaveInstrument sets SaveInstrument field to given value.

### HasSaveInstrument

`func (o *CFOrderPayRequest) HasSaveInstrument() bool`

HasSaveInstrument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


