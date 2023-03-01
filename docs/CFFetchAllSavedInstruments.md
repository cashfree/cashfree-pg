# CFFetchAllSavedInstruments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | customer_id for which the instrument was saved | [optional] 
**AfaReference** | Pointer to **int32** | cf_payment_id of the successful transaction done while saving instrument | [optional] 
**InstrumentId** | Pointer to **string** | saved instrument id | [optional] 
**InstrumentType** | Pointer to **string** | Type of the saved instrument | [optional] 
**InstrumentUid** | Pointer to **string** | Unique id for the saved instrument | [optional] 
**InstrumentDisplay** | Pointer to **string** | masked card number displayed to the customer | [optional] 
**InstrumentStatus** | Pointer to **string** | status of the saved instrument | [optional] 
**CreatedAt** | Pointer to **string** | timestamp at which instrument was saved | [optional] 
**InstrumentMeta** | Pointer to [**CFSavedInstrumentMeta**](CFSavedInstrumentMeta.md) |  | [optional] 

## Methods

### NewCFFetchAllSavedInstruments

`func NewCFFetchAllSavedInstruments() *CFFetchAllSavedInstruments`

NewCFFetchAllSavedInstruments instantiates a new CFFetchAllSavedInstruments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFFetchAllSavedInstrumentsWithDefaults

`func NewCFFetchAllSavedInstrumentsWithDefaults() *CFFetchAllSavedInstruments`

NewCFFetchAllSavedInstrumentsWithDefaults instantiates a new CFFetchAllSavedInstruments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CFFetchAllSavedInstruments) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CFFetchAllSavedInstruments) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CFFetchAllSavedInstruments) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CFFetchAllSavedInstruments) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAfaReference

`func (o *CFFetchAllSavedInstruments) GetAfaReference() int32`

GetAfaReference returns the AfaReference field if non-nil, zero value otherwise.

### GetAfaReferenceOk

`func (o *CFFetchAllSavedInstruments) GetAfaReferenceOk() (*int32, bool)`

GetAfaReferenceOk returns a tuple with the AfaReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfaReference

`func (o *CFFetchAllSavedInstruments) SetAfaReference(v int32)`

SetAfaReference sets AfaReference field to given value.

### HasAfaReference

`func (o *CFFetchAllSavedInstruments) HasAfaReference() bool`

HasAfaReference returns a boolean if a field has been set.

### GetInstrumentId

`func (o *CFFetchAllSavedInstruments) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *CFFetchAllSavedInstruments) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *CFFetchAllSavedInstruments) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetInstrumentType

`func (o *CFFetchAllSavedInstruments) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *CFFetchAllSavedInstruments) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *CFFetchAllSavedInstruments) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetInstrumentUid

`func (o *CFFetchAllSavedInstruments) GetInstrumentUid() string`

GetInstrumentUid returns the InstrumentUid field if non-nil, zero value otherwise.

### GetInstrumentUidOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentUidOk() (*string, bool)`

GetInstrumentUidOk returns a tuple with the InstrumentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentUid

`func (o *CFFetchAllSavedInstruments) SetInstrumentUid(v string)`

SetInstrumentUid sets InstrumentUid field to given value.

### HasInstrumentUid

`func (o *CFFetchAllSavedInstruments) HasInstrumentUid() bool`

HasInstrumentUid returns a boolean if a field has been set.

### GetInstrumentDisplay

`func (o *CFFetchAllSavedInstruments) GetInstrumentDisplay() string`

GetInstrumentDisplay returns the InstrumentDisplay field if non-nil, zero value otherwise.

### GetInstrumentDisplayOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentDisplayOk() (*string, bool)`

GetInstrumentDisplayOk returns a tuple with the InstrumentDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDisplay

`func (o *CFFetchAllSavedInstruments) SetInstrumentDisplay(v string)`

SetInstrumentDisplay sets InstrumentDisplay field to given value.

### HasInstrumentDisplay

`func (o *CFFetchAllSavedInstruments) HasInstrumentDisplay() bool`

HasInstrumentDisplay returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *CFFetchAllSavedInstruments) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *CFFetchAllSavedInstruments) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *CFFetchAllSavedInstruments) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CFFetchAllSavedInstruments) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CFFetchAllSavedInstruments) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CFFetchAllSavedInstruments) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CFFetchAllSavedInstruments) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInstrumentMeta

`func (o *CFFetchAllSavedInstruments) GetInstrumentMeta() CFSavedInstrumentMeta`

GetInstrumentMeta returns the InstrumentMeta field if non-nil, zero value otherwise.

### GetInstrumentMetaOk

`func (o *CFFetchAllSavedInstruments) GetInstrumentMetaOk() (*CFSavedInstrumentMeta, bool)`

GetInstrumentMetaOk returns a tuple with the InstrumentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentMeta

`func (o *CFFetchAllSavedInstruments) SetInstrumentMeta(v CFSavedInstrumentMeta)`

SetInstrumentMeta sets InstrumentMeta field to given value.

### HasInstrumentMeta

`func (o *CFFetchAllSavedInstruments) HasInstrumentMeta() bool`

HasInstrumentMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


