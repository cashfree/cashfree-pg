# InstrumentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | customer_id for which the instrument was saved | [optional] 
**AfaReference** | Pointer to **string** | cf_payment_id of the successful transaction done while saving instrument | [optional] 
**InstrumentId** | Pointer to **string** | saved instrument id | [optional] 
**InstrumentType** | Pointer to **string** | Type of the saved instrument | [optional] 
**InstrumentUid** | Pointer to **string** | Unique id for the saved instrument | [optional] 
**InstrumentDisplay** | Pointer to **string** | masked card number displayed to the customer | [optional] 
**InstrumentStatus** | Pointer to **string** | Status of the saved instrument. | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp at which instrument was saved. | [optional] 
**InstrumentMeta** | Pointer to [**SavedInstrumentMeta**](SavedInstrumentMeta.md) |  | [optional] 

## Methods

### NewInstrumentEntity

`func NewInstrumentEntity() *InstrumentEntity`

NewInstrumentEntity instantiates a new InstrumentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentEntityWithDefaults

`func NewInstrumentEntityWithDefaults() *InstrumentEntity`

NewInstrumentEntityWithDefaults instantiates a new InstrumentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *InstrumentEntity) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InstrumentEntity) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InstrumentEntity) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InstrumentEntity) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAfaReference

`func (o *InstrumentEntity) GetAfaReference() string`

GetAfaReference returns the AfaReference field if non-nil, zero value otherwise.

### GetAfaReferenceOk

`func (o *InstrumentEntity) GetAfaReferenceOk() (*string, bool)`

GetAfaReferenceOk returns a tuple with the AfaReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfaReference

`func (o *InstrumentEntity) SetAfaReference(v string)`

SetAfaReference sets AfaReference field to given value.

### HasAfaReference

`func (o *InstrumentEntity) HasAfaReference() bool`

HasAfaReference returns a boolean if a field has been set.

### GetInstrumentId

`func (o *InstrumentEntity) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *InstrumentEntity) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *InstrumentEntity) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *InstrumentEntity) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetInstrumentType

`func (o *InstrumentEntity) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *InstrumentEntity) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *InstrumentEntity) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *InstrumentEntity) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetInstrumentUid

`func (o *InstrumentEntity) GetInstrumentUid() string`

GetInstrumentUid returns the InstrumentUid field if non-nil, zero value otherwise.

### GetInstrumentUidOk

`func (o *InstrumentEntity) GetInstrumentUidOk() (*string, bool)`

GetInstrumentUidOk returns a tuple with the InstrumentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentUid

`func (o *InstrumentEntity) SetInstrumentUid(v string)`

SetInstrumentUid sets InstrumentUid field to given value.

### HasInstrumentUid

`func (o *InstrumentEntity) HasInstrumentUid() bool`

HasInstrumentUid returns a boolean if a field has been set.

### GetInstrumentDisplay

`func (o *InstrumentEntity) GetInstrumentDisplay() string`

GetInstrumentDisplay returns the InstrumentDisplay field if non-nil, zero value otherwise.

### GetInstrumentDisplayOk

`func (o *InstrumentEntity) GetInstrumentDisplayOk() (*string, bool)`

GetInstrumentDisplayOk returns a tuple with the InstrumentDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDisplay

`func (o *InstrumentEntity) SetInstrumentDisplay(v string)`

SetInstrumentDisplay sets InstrumentDisplay field to given value.

### HasInstrumentDisplay

`func (o *InstrumentEntity) HasInstrumentDisplay() bool`

HasInstrumentDisplay returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *InstrumentEntity) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *InstrumentEntity) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *InstrumentEntity) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *InstrumentEntity) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstrumentEntity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstrumentEntity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstrumentEntity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstrumentEntity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInstrumentMeta

`func (o *InstrumentEntity) GetInstrumentMeta() SavedInstrumentMeta`

GetInstrumentMeta returns the InstrumentMeta field if non-nil, zero value otherwise.

### GetInstrumentMetaOk

`func (o *InstrumentEntity) GetInstrumentMetaOk() (*SavedInstrumentMeta, bool)`

GetInstrumentMetaOk returns a tuple with the InstrumentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentMeta

`func (o *InstrumentEntity) SetInstrumentMeta(v SavedInstrumentMeta)`

SetInstrumentMeta sets InstrumentMeta field to given value.

### HasInstrumentMeta

`func (o *InstrumentEntity) HasInstrumentMeta() bool`

HasInstrumentMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


