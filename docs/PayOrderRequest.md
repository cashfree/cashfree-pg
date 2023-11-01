# PayOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSessionId** | **string** |  | 
**PaymentMethod** | [**PayOrderRequestPaymentMethod**](PayOrderRequestPaymentMethod.md) |  | 
**SaveInstrument** | Pointer to **bool** |  | [optional] 
**OfferId** | Pointer to **string** | This is required if any offers needs to be applied to the order. | [optional] 

## Methods

### NewPayOrderRequest

`func NewPayOrderRequest(paymentSessionId string, paymentMethod PayOrderRequestPaymentMethod, ) *PayOrderRequest`

NewPayOrderRequest instantiates a new PayOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderRequestWithDefaults

`func NewPayOrderRequestWithDefaults() *PayOrderRequest`

NewPayOrderRequestWithDefaults instantiates a new PayOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSessionId

`func (o *PayOrderRequest) GetPaymentSessionId() string`

GetPaymentSessionId returns the PaymentSessionId field if non-nil, zero value otherwise.

### GetPaymentSessionIdOk

`func (o *PayOrderRequest) GetPaymentSessionIdOk() (*string, bool)`

GetPaymentSessionIdOk returns a tuple with the PaymentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSessionId

`func (o *PayOrderRequest) SetPaymentSessionId(v string)`

SetPaymentSessionId sets PaymentSessionId field to given value.


### GetPaymentMethod

`func (o *PayOrderRequest) GetPaymentMethod() PayOrderRequestPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PayOrderRequest) GetPaymentMethodOk() (*PayOrderRequestPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PayOrderRequest) SetPaymentMethod(v PayOrderRequestPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetSaveInstrument

`func (o *PayOrderRequest) GetSaveInstrument() bool`

GetSaveInstrument returns the SaveInstrument field if non-nil, zero value otherwise.

### GetSaveInstrumentOk

`func (o *PayOrderRequest) GetSaveInstrumentOk() (*bool, bool)`

GetSaveInstrumentOk returns a tuple with the SaveInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveInstrument

`func (o *PayOrderRequest) SetSaveInstrument(v bool)`

SetSaveInstrument sets SaveInstrument field to given value.

### HasSaveInstrument

`func (o *PayOrderRequest) HasSaveInstrument() bool`

HasSaveInstrument returns a boolean if a field has been set.

### GetOfferId

`func (o *PayOrderRequest) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *PayOrderRequest) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *PayOrderRequest) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *PayOrderRequest) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


