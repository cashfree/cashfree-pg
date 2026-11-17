# OrderSettlementEntitySettlementDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfSettlementId** | Pointer to **string** | Unique ID to identify the settlement. | [optional] 
**Status** | Pointer to **string** | Status of the linked settlement. | [optional] 
**StatusDescription** | Pointer to **string** | Human-readable description of the settlement status. | [optional] 
**SettlementUtr** | Pointer to **string** | Unique transaction reference number of the settlement. | [optional] 
**SettlementInitiatedOn** | Pointer to **string** | Date and time when the settlement was initiated. | [optional] 
**SettlementProcessedOn** | Pointer to **string** | Date and time when the settlement was processed. | [optional] 
**SettlementCurrency** | Pointer to **string** | Currency of the settlement. | [optional] 
**SettlementType** | Pointer to **string** | Type of settlement. Example - NORMAL_SETTLEMENT. | [optional] 

## Methods

### NewOrderSettlementEntitySettlementDetails

`func NewOrderSettlementEntitySettlementDetails() *OrderSettlementEntitySettlementDetails`

NewOrderSettlementEntitySettlementDetails instantiates a new OrderSettlementEntitySettlementDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSettlementEntitySettlementDetailsWithDefaults

`func NewOrderSettlementEntitySettlementDetailsWithDefaults() *OrderSettlementEntitySettlementDetails`

NewOrderSettlementEntitySettlementDetailsWithDefaults instantiates a new OrderSettlementEntitySettlementDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfSettlementId

`func (o *OrderSettlementEntitySettlementDetails) GetCfSettlementId() string`

GetCfSettlementId returns the CfSettlementId field if non-nil, zero value otherwise.

### GetCfSettlementIdOk

`func (o *OrderSettlementEntitySettlementDetails) GetCfSettlementIdOk() (*string, bool)`

GetCfSettlementIdOk returns a tuple with the CfSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfSettlementId

`func (o *OrderSettlementEntitySettlementDetails) SetCfSettlementId(v string)`

SetCfSettlementId sets CfSettlementId field to given value.

### HasCfSettlementId

`func (o *OrderSettlementEntitySettlementDetails) HasCfSettlementId() bool`

HasCfSettlementId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderSettlementEntitySettlementDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderSettlementEntitySettlementDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderSettlementEntitySettlementDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderSettlementEntitySettlementDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *OrderSettlementEntitySettlementDetails) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *OrderSettlementEntitySettlementDetails) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *OrderSettlementEntitySettlementDetails) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *OrderSettlementEntitySettlementDetails) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetSettlementUtr

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementUtr() string`

GetSettlementUtr returns the SettlementUtr field if non-nil, zero value otherwise.

### GetSettlementUtrOk

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementUtrOk() (*string, bool)`

GetSettlementUtrOk returns a tuple with the SettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtr

`func (o *OrderSettlementEntitySettlementDetails) SetSettlementUtr(v string)`

SetSettlementUtr sets SettlementUtr field to given value.

### HasSettlementUtr

`func (o *OrderSettlementEntitySettlementDetails) HasSettlementUtr() bool`

HasSettlementUtr returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *OrderSettlementEntitySettlementDetails) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *OrderSettlementEntitySettlementDetails) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementProcessedOn

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementProcessedOn() string`

GetSettlementProcessedOn returns the SettlementProcessedOn field if non-nil, zero value otherwise.

### GetSettlementProcessedOnOk

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementProcessedOnOk() (*string, bool)`

GetSettlementProcessedOnOk returns a tuple with the SettlementProcessedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementProcessedOn

`func (o *OrderSettlementEntitySettlementDetails) SetSettlementProcessedOn(v string)`

SetSettlementProcessedOn sets SettlementProcessedOn field to given value.

### HasSettlementProcessedOn

`func (o *OrderSettlementEntitySettlementDetails) HasSettlementProcessedOn() bool`

HasSettlementProcessedOn returns a boolean if a field has been set.

### GetSettlementCurrency

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementCurrency() string`

GetSettlementCurrency returns the SettlementCurrency field if non-nil, zero value otherwise.

### GetSettlementCurrencyOk

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementCurrencyOk() (*string, bool)`

GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrency

`func (o *OrderSettlementEntitySettlementDetails) SetSettlementCurrency(v string)`

SetSettlementCurrency sets SettlementCurrency field to given value.

### HasSettlementCurrency

`func (o *OrderSettlementEntitySettlementDetails) HasSettlementCurrency() bool`

HasSettlementCurrency returns a boolean if a field has been set.

### GetSettlementType

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *OrderSettlementEntitySettlementDetails) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *OrderSettlementEntitySettlementDetails) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *OrderSettlementEntitySettlementDetails) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


