# PayOrderEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentAmount** | Pointer to **float32** | Total amount payable. | [optional] 
**CfPaymentId** | Pointer to **string** | Payment identifier created by Cashfree. | [optional] 
**PaymentMethod** | Pointer to **string** | The payment method used for this transaction. - netbanking: Net banking payment. - card: Credit or debit card payment. - upi: UPI payment via collect, intent, or QR code. - app: Wallet-based payment. - cardless_emi: Cardless EMI payment. - paylater: Pay later payment. - banktransfer: Direct bank transfer payment. - applepay: Apple Pay payment.  | [optional] 
**Channel** | Pointer to **string** | The channel used for the payment method. - link: Redirect-based flow where the customer is taken to an external page. - post: Native OTP flow where the merchant renders a custom UI to collect OTP. - collect: UPI collect request sent to the customer&#39;s VPA. - qrcode: UPI QR code for the customer to scan. - podQrCode: Pay on delivery QR code.  | [optional] 
**Action** | Pointer to **string** | The action to complete the payment. - link: Redirect the customer to &#x60;data.url&#x60; using a browser or in-app webview. - post: Render a native UI, collect required input, and POST it to &#x60;data.url&#x60;. - form: Render the form from &#x60;data.payload&#x60; and auto-submit it to &#x60;data.url&#x60;. - custom: Follow integration-specific instructions or SDK handling.  | [optional] 
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

`func (o *PayOrderEntity) GetCfPaymentId() string`

GetCfPaymentId returns the CfPaymentId field if non-nil, zero value otherwise.

### GetCfPaymentIdOk

`func (o *PayOrderEntity) GetCfPaymentIdOk() (*string, bool)`

GetCfPaymentIdOk returns a tuple with the CfPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfPaymentId

`func (o *PayOrderEntity) SetCfPaymentId(v string)`

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


