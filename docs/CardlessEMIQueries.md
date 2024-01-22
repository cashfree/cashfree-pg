# CardlessEMIQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | OrderId of the order. Either of &#x60;order_id&#x60; or &#x60;amount&#x60; is mandatory. | [optional] 
**Amount** | Pointer to **float32** | Amount of the order. OrderId of the order. Either of &#x60;order_id&#x60; or &#x60;amount&#x60; is mandatory. | [optional] 
**CustomerDetails** | Pointer to [**CustomerDetailsCardlessEMI**](CustomerDetailsCardlessEMI.md) |  | [optional] 

## Methods

### NewCardlessEMIQueries

`func NewCardlessEMIQueries() *CardlessEMIQueries`

NewCardlessEMIQueries instantiates a new CardlessEMIQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardlessEMIQueriesWithDefaults

`func NewCardlessEMIQueriesWithDefaults() *CardlessEMIQueries`

NewCardlessEMIQueriesWithDefaults instantiates a new CardlessEMIQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CardlessEMIQueries) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CardlessEMIQueries) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CardlessEMIQueries) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CardlessEMIQueries) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAmount

`func (o *CardlessEMIQueries) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardlessEMIQueries) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardlessEMIQueries) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CardlessEMIQueries) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *CardlessEMIQueries) GetCustomerDetails() CustomerDetailsCardlessEMI`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CardlessEMIQueries) GetCustomerDetailsOk() (*CustomerDetailsCardlessEMI, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CardlessEMIQueries) SetCustomerDetails(v CustomerDetailsCardlessEMI)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *CardlessEMIQueries) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


