# DeletedInstrumentEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Customer ID against which the card is saved. | [optional] 
**AfaReference** | Pointer to **string** | cf_payment_id of the successful transaction done while saving instrument. | [optional] 
**InstrumentId** | Pointer to **string** | Identifier for the card saved at Cashfree, which was requested to be deleted. | [optional] 
**InstrumentType** | Pointer to **string** | Type of the saved instrument. | [optional] 
**InstrumentUid** | Pointer to **string** | Unique identifier for the saved card, used to identify a specific card. | [optional] 
**InstrumentDisplay** | Pointer to **string** | Last four digits of actual card number. | [optional] 
**InstrumentStatus** | Pointer to **string** | Status of the saved instrument. This would be &#x60;INACTIVE&#x60; when the instrument is successfully deleted. | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp at which instrument was saved. | [optional] 
**InstrumentMeta** | Pointer to [**SavedInstrumentMeta**](SavedInstrumentMeta.md) |  | [optional] 

## Methods

### NewDeletedInstrumentEntity

`func NewDeletedInstrumentEntity() *DeletedInstrumentEntity`

NewDeletedInstrumentEntity instantiates a new DeletedInstrumentEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedInstrumentEntityWithDefaults

`func NewDeletedInstrumentEntityWithDefaults() *DeletedInstrumentEntity`

NewDeletedInstrumentEntityWithDefaults instantiates a new DeletedInstrumentEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DeletedInstrumentEntity) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DeletedInstrumentEntity) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DeletedInstrumentEntity) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DeletedInstrumentEntity) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAfaReference

`func (o *DeletedInstrumentEntity) GetAfaReference() string`

GetAfaReference returns the AfaReference field if non-nil, zero value otherwise.

### GetAfaReferenceOk

`func (o *DeletedInstrumentEntity) GetAfaReferenceOk() (*string, bool)`

GetAfaReferenceOk returns a tuple with the AfaReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfaReference

`func (o *DeletedInstrumentEntity) SetAfaReference(v string)`

SetAfaReference sets AfaReference field to given value.

### HasAfaReference

`func (o *DeletedInstrumentEntity) HasAfaReference() bool`

HasAfaReference returns a boolean if a field has been set.

### GetInstrumentId

`func (o *DeletedInstrumentEntity) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *DeletedInstrumentEntity) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *DeletedInstrumentEntity) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *DeletedInstrumentEntity) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetInstrumentType

`func (o *DeletedInstrumentEntity) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *DeletedInstrumentEntity) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *DeletedInstrumentEntity) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *DeletedInstrumentEntity) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetInstrumentUid

`func (o *DeletedInstrumentEntity) GetInstrumentUid() string`

GetInstrumentUid returns the InstrumentUid field if non-nil, zero value otherwise.

### GetInstrumentUidOk

`func (o *DeletedInstrumentEntity) GetInstrumentUidOk() (*string, bool)`

GetInstrumentUidOk returns a tuple with the InstrumentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentUid

`func (o *DeletedInstrumentEntity) SetInstrumentUid(v string)`

SetInstrumentUid sets InstrumentUid field to given value.

### HasInstrumentUid

`func (o *DeletedInstrumentEntity) HasInstrumentUid() bool`

HasInstrumentUid returns a boolean if a field has been set.

### GetInstrumentDisplay

`func (o *DeletedInstrumentEntity) GetInstrumentDisplay() string`

GetInstrumentDisplay returns the InstrumentDisplay field if non-nil, zero value otherwise.

### GetInstrumentDisplayOk

`func (o *DeletedInstrumentEntity) GetInstrumentDisplayOk() (*string, bool)`

GetInstrumentDisplayOk returns a tuple with the InstrumentDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDisplay

`func (o *DeletedInstrumentEntity) SetInstrumentDisplay(v string)`

SetInstrumentDisplay sets InstrumentDisplay field to given value.

### HasInstrumentDisplay

`func (o *DeletedInstrumentEntity) HasInstrumentDisplay() bool`

HasInstrumentDisplay returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *DeletedInstrumentEntity) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *DeletedInstrumentEntity) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *DeletedInstrumentEntity) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *DeletedInstrumentEntity) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeletedInstrumentEntity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeletedInstrumentEntity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeletedInstrumentEntity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeletedInstrumentEntity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInstrumentMeta

`func (o *DeletedInstrumentEntity) GetInstrumentMeta() SavedInstrumentMeta`

GetInstrumentMeta returns the InstrumentMeta field if non-nil, zero value otherwise.

### GetInstrumentMetaOk

`func (o *DeletedInstrumentEntity) GetInstrumentMetaOk() (*SavedInstrumentMeta, bool)`

GetInstrumentMetaOk returns a tuple with the InstrumentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentMeta

`func (o *DeletedInstrumentEntity) SetInstrumentMeta(v SavedInstrumentMeta)`

SetInstrumentMeta sets InstrumentMeta field to given value.

### HasInstrumentMeta

`func (o *DeletedInstrumentEntity) HasInstrumentMeta() bool`

HasInstrumentMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


