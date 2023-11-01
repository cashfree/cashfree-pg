# PayOrderEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAmount** | Pointer to **float32** | total amount payable | [optional] 
**CfPaymentId** | Pointer to **int64** | Payment identifier created by Cashfree | [optional] 
**PaymentMethod** | Pointer to **string** | One of [\&quot;upi\&quot;, \&quot;netbanking\&quot;, \&quot;card\&quot;, \&quot;app\&quot;, \&quot;cardless_emi\&quot;, \&quot;paylater\&quot;, \&quot;banktransfer\&quot;]  | [optional] 
**Channel** | Pointer to **string** | One of [\&quot;link\&quot;, \&quot;collect\&quot;, \&quot;qrcode\&quot;]. In an older version we used to support different channels like &#39;gpay&#39;, &#39;phonepe&#39; etc. However, we now support only the following channels - link, collect and qrcode. To process payments using gpay, you will have to provide channel as &#39;link&#39; and provider as &#39;gpay&#39; | [optional] 
**Action** | Pointer to **NullableString** | One of [\&quot;link\&quot;, \&quot;custom\&quot;, \&quot;form\&quot;] | [optional] 
**Data** | Pointer to [**OrderPayData**](OrderPayData.md) |  | [optional] 

## Methods

### NewPayOrderEntity

`func NewPayOrderEntity() *PayOrderEntity`

NewPayOrderEntity instantiates a new PayOrderEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderEntityWithDefaults

`func NewPayOrderEntityWithDefaults() *PayOrderEntity`

NewPayOrderEntityWithDefaults instantiates a new PayOrderEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentAmount

`func (o *PayOrderEntity) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PayOrderEntity) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PayOrderEntity) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PayOrderEntity) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetCfPaymentId

`func (o *PayOrderEntity) GetCfPaymentId() int64`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *PayOrderEntity) GetCfPaymentIdOk() (*int64, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *PayOrderEntity) SetCfPaymentId(v int64)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *PayOrderEntity) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PayOrderEntity) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PayOrderEntity) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PayOrderEntity) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PayOrderEntity) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetChannel

`func (o *PayOrderEntity) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PayOrderEntity) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PayOrderEntity) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PayOrderEntity) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetAction

`func (o *PayOrderEntity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PayOrderEntity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PayOrderEntity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PayOrderEntity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *PayOrderEntity) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *PayOrderEntity) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetData

`func (o *PayOrderEntity) GetData() OrderPayData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PayOrderEntity) GetDataOk() (*OrderPayData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PayOrderEntity) SetData(v OrderPayData)`

SetData sets Data field to given value.

### HasData

`func (o *PayOrderEntity) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


