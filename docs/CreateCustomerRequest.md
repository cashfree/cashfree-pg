# CreateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPhone** | **string** | Customer Phone Number | 
**CustomerEmail** | Pointer to **string** | Customer Email | [optional] 
**CustomerName** | Pointer to **string** | Customer Name | [optional] 

## Methods

### NewCreateCustomerRequest

`func NewCreateCustomerRequest(customerPhone string, ) *CreateCustomerRequest`

NewCreateCustomerRequest instantiates a new CreateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerRequestWithDefaults

`func NewCreateCustomerRequestWithDefaults() *CreateCustomerRequest`

NewCreateCustomerRequestWithDefaults instantiates a new CreateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPhone

`func (o *CreateCustomerRequest) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *CreateCustomerRequest) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *CreateCustomerRequest) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerEmail

`func (o *CreateCustomerRequest) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CreateCustomerRequest) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CreateCustomerRequest) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CreateCustomerRequest) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *CreateCustomerRequest) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *CreateCustomerRequest) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *CreateCustomerRequest) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *CreateCustomerRequest) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


