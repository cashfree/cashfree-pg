# CFLinkCustomerDetailsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPhone** | **string** | Customer phone number | 
**CustomerEmail** | Pointer to **string** | Customer email address | [optional] 
**CustomerName** | Pointer to **string** | Customer name | [optional] 

## Methods

### NewCFLinkCustomerDetailsEntity

`func NewCFLinkCustomerDetailsEntity(customerPhone string, ) *CFLinkCustomerDetailsEntity`

NewCFLinkCustomerDetailsEntity instantiates a new CFLinkCustomerDetailsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkCustomerDetailsEntityWithDefaults

`func NewCFLinkCustomerDetailsEntityWithDefaults() *CFLinkCustomerDetailsEntity`

NewCFLinkCustomerDetailsEntityWithDefaults instantiates a new CFLinkCustomerDetailsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPhone

`func (o *CFLinkCustomerDetailsEntity) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *CFLinkCustomerDetailsEntity) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *CFLinkCustomerDetailsEntity) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerEmail

`func (o *CFLinkCustomerDetailsEntity) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *CFLinkCustomerDetailsEntity) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *CFLinkCustomerDetailsEntity) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *CFLinkCustomerDetailsEntity) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *CFLinkCustomerDetailsEntity) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *CFLinkCustomerDetailsEntity) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *CFLinkCustomerDetailsEntity) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *CFLinkCustomerDetailsEntity) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


