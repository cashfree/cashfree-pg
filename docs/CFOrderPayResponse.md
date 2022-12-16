# CFOrderPayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfPaymentId** | Pointer to **int32** | Payment identifier created by Cashfree | [optional] 
**PaymentMethod** | Pointer to **string** | One of [\&quot;upi\&quot;, \&quot;netbanking\&quot;, \&quot;card\&quot;, \&quot;app\&quot;, \&quot;cardless_emi\&quot;, \&quot;paylater\&quot;]  | [optional] 
**Channel** | Pointer to **string** | One of [\&quot;link\&quot;, \&quot;collect\&quot;, \&quot;qrcode\&quot;]. In an older version we used to support different channels like &#39;gpay&#39;, &#39;phonepe&#39; etc. However, we now support only the following channels - link, collect and qrcode. To process payments using gpay, you will have to provide channel as &#39;link&#39; and provider as &#39;gpay&#39; | [optional] 
**Action** | Pointer to **NullableString** | One of [\&quot;link\&quot;, \&quot;custom\&quot;, \&quot;form\&quot;] | [optional] 
**Data** | Pointer to [**CFOrderPayData**](CFOrderPayData.md) |  | [optional] 

## Methods

### NewCFOrderPayResponse

`func NewCFOrderPayResponse() *CFOrderPayResponse`

NewCFOrderPayResponse instantiates a new CFOrderPayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFOrderPayResponseWithDefaults

`func NewCFOrderPayResponseWithDefaults() *CFOrderPayResponse`

NewCFOrderPayResponseWithDefaults instantiates a new CFOrderPayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfPaymentId

`func (o *CFOrderPayResponse) GetCfPaymentId() int32`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *CFOrderPayResponse) GetCfPaymentIdOk() (*int32, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *CFOrderPayResponse) SetCfPaymentId(v int32)`

SetCfPaymentId sets CfPaymentId field to given value.

### HasCfPaymentId

`func (o *CFOrderPayResponse) HasCfPaymentId() bool`

HasCfPaymentId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CFOrderPayResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CFOrderPayResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CFOrderPayResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CFOrderPayResponse) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetChannel

`func (o *CFOrderPayResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFOrderPayResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFOrderPayResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CFOrderPayResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetAction

`func (o *CFOrderPayResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CFOrderPayResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CFOrderPayResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CFOrderPayResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CFOrderPayResponse) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CFOrderPayResponse) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetData

`func (o *CFOrderPayResponse) GetData() CFOrderPayData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CFOrderPayResponse) GetDataOk() (*CFOrderPayData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CFOrderPayResponse) SetData(v CFOrderPayData)`

SetData sets Data field to given value.

### HasData

`func (o *CFOrderPayResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


