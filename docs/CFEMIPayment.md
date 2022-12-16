# CFEMIPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emi** | Pointer to [**CFCardEMI**](CFCardEMI.md) |  | [optional] 

## Methods

### NewCFEMIPayment

`func NewCFEMIPayment() *CFEMIPayment`

NewCFEMIPayment instantiates a new CFEMIPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFEMIPaymentWithDefaults

`func NewCFEMIPaymentWithDefaults() *CFEMIPayment`

NewCFEMIPaymentWithDefaults instantiates a new CFEMIPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmi

`func (o *CFEMIPayment) GetEmi() CFCardEMI`

GetEmi returns the Emi field if non-nil, zero value otherwise.

### GetEmiOk

`func (o *CFEMIPayment) GetEmiOk() (*CFCardEMI, bool)`

GetEmiOk returns a tuple with the Emi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmi

`func (o *CFEMIPayment) SetEmi(v CFCardEMI)`

SetEmi sets Emi field to given value.

### HasEmi

`func (o *CFEMIPayment) HasEmi() bool`

HasEmi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


