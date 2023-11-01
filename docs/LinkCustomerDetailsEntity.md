# LinkCustomerDetailsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPhone** | **string** | Customer phone number | 
**CustomerEmail** | Pointer to **string** | Customer email address | [optional] 
**CustomerName** | Pointer to **string** | Customer name | [optional] 

## Methods

### NewLinkCustomerDetailsEntity

`func NewLinkCustomerDetailsEntity(customerPhone string, ) *LinkCustomerDetailsEntity`

NewLinkCustomerDetailsEntity instantiates a new LinkCustomerDetailsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkCustomerDetailsEntityWithDefaults

`func NewLinkCustomerDetailsEntityWithDefaults() *LinkCustomerDetailsEntity`

NewLinkCustomerDetailsEntityWithDefaults instantiates a new LinkCustomerDetailsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPhone

`func (o *LinkCustomerDetailsEntity) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *LinkCustomerDetailsEntity) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *LinkCustomerDetailsEntity) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.


### GetCustomerEmail

`func (o *LinkCustomerDetailsEntity) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *LinkCustomerDetailsEntity) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *LinkCustomerDetailsEntity) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *LinkCustomerDetailsEntity) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *LinkCustomerDetailsEntity) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *LinkCustomerDetailsEntity) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *LinkCustomerDetailsEntity) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *LinkCustomerDetailsEntity) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


