# WHdata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**WHorder**](WHorder.md) |  | [optional] 
**Payment** | Pointer to [**PaymentEntity**](PaymentEntity.md) |  | [optional] 
**CustomerDetails** | Pointer to [**WHcustomerDetails**](WHcustomerDetails.md) |  | [optional] 

## Methods

### NewWHdata

`func NewWHdata() *WHdata`

NewWHdata instantiates a new WHdata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWHdataWithDefaults

`func NewWHdataWithDefaults() *WHdata`

NewWHdataWithDefaults instantiates a new WHdata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *WHdata) GetOrder() WHorder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WHdata) GetOrderOk() (*WHorder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WHdata) SetOrder(v WHorder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WHdata) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPayment

`func (o *WHdata) GetPayment() PaymentEntity`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *WHdata) GetPaymentOk() (*PaymentEntity, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *WHdata) SetPayment(v PaymentEntity)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *WHdata) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *WHdata) GetCustomerDetails() WHcustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *WHdata) GetCustomerDetailsOk() (*WHcustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *WHdata) SetCustomerDetails(v WHcustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *WHdata) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


