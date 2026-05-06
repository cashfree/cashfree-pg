# InstrumentEntityForAllSavedCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Customer ID that merchant sends during [Create Order API](https://www.cashfree.com/docs/api-reference/payments/latest/orders/create), against which the cards are saved for the customer. | [optional] 
**AfaReference** | Pointer to **string** | cf_payment_id of the successful transaction done while saving instrument. | [optional] 
**InstrumentId** | Pointer to **string** | Identifier for the card saved at Cashfree. It is used for [cryptogram generation](https://www.cashfree.com/docs/api-reference/payments/latest/token-vault/generate-cryptogram) and in [order pay](https://www.cashfree.com/docs/api-reference/payments/latest/payments/pay) request for saved cards.. | [optional] 
**InstrumentType** | Pointer to **string** | Type of the saved instrument. Available option is &#x60;card&#x60;. | [optional] 
**InstrumentUid** | Pointer to **string** | Unique identifier for the saved card, used to identify a specific card. | [optional] 
**InstrumentDisplay** | Pointer to **string** | Last 4 digits of actual card number, to be displayed to the customer for card identification. | [optional] 
**InstrumentStatus** | Pointer to **string** | Status of the saved instrument. Available options are &#x60;ACTIVE&#x60;, &#x60;INACTIVE&#x60;. | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp at which instrument was saved. | [optional] 
**InstrumentMeta** | Pointer to [**SavedInstrumentMeta**](SavedInstrumentMeta.md) |  | [optional] 

## Methods

### NewInstrumentEntityForAllSavedCard

`func NewInstrumentEntityForAllSavedCard() *InstrumentEntityForAllSavedCard`

NewInstrumentEntityForAllSavedCard instantiates a new InstrumentEntityForAllSavedCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentEntityForAllSavedCardWithDefaults

`func NewInstrumentEntityForAllSavedCardWithDefaults() *InstrumentEntityForAllSavedCard`

NewInstrumentEntityForAllSavedCardWithDefaults instantiates a new InstrumentEntityForAllSavedCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *InstrumentEntityForAllSavedCard) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InstrumentEntityForAllSavedCard) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InstrumentEntityForAllSavedCard) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InstrumentEntityForAllSavedCard) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAfaReference

`func (o *InstrumentEntityForAllSavedCard) GetAfaReference() string`

GetAfaReference returns the AfaReference field if non-nil, zero value otherwise.

### GetAfaReferenceOk

`func (o *InstrumentEntityForAllSavedCard) GetAfaReferenceOk() (*string, bool)`

GetAfaReferenceOk returns a tuple with the AfaReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfaReference

`func (o *InstrumentEntityForAllSavedCard) SetAfaReference(v string)`

SetAfaReference sets AfaReference field to given value.

### HasAfaReference

`func (o *InstrumentEntityForAllSavedCard) HasAfaReference() bool`

HasAfaReference returns a boolean if a field has been set.

### GetInstrumentId

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetInstrumentType

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetInstrumentUid

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentUid() string`

GetInstrumentUid returns the InstrumentUid field if non-nil, zero value otherwise.

### GetInstrumentUidOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentUidOk() (*string, bool)`

GetInstrumentUidOk returns a tuple with the InstrumentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentUid

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentUid(v string)`

SetInstrumentUid sets InstrumentUid field to given value.

### HasInstrumentUid

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentUid() bool`

HasInstrumentUid returns a boolean if a field has been set.

### GetInstrumentDisplay

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentDisplay() string`

GetInstrumentDisplay returns the InstrumentDisplay field if non-nil, zero value otherwise.

### GetInstrumentDisplayOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentDisplayOk() (*string, bool)`

GetInstrumentDisplayOk returns a tuple with the InstrumentDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDisplay

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentDisplay(v string)`

SetInstrumentDisplay sets InstrumentDisplay field to given value.

### HasInstrumentDisplay

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentDisplay() bool`

HasInstrumentDisplay returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstrumentEntityForAllSavedCard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstrumentEntityForAllSavedCard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstrumentEntityForAllSavedCard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstrumentEntityForAllSavedCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetInstrumentMeta

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentMeta() SavedInstrumentMeta`

GetInstrumentMeta returns the InstrumentMeta field if non-nil, zero value otherwise.

### GetInstrumentMetaOk

`func (o *InstrumentEntityForAllSavedCard) GetInstrumentMetaOk() (*SavedInstrumentMeta, bool)`

GetInstrumentMetaOk returns a tuple with the InstrumentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentMeta

`func (o *InstrumentEntityForAllSavedCard) SetInstrumentMeta(v SavedInstrumentMeta)`

SetInstrumentMeta sets InstrumentMeta field to given value.

### HasInstrumentMeta

`func (o *InstrumentEntityForAllSavedCard) HasInstrumentMeta() bool`

HasInstrumentMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


