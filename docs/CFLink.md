# CFLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfLinkId** | Pointer to **int32** |  | [optional] 
**LinkId** | Pointer to **string** |  | [optional] 
**LinkStatus** | Pointer to **string** |  | [optional] 
**LinkCurrency** | Pointer to **string** |  | [optional] 
**LinkAmount** | Pointer to **float32** |  | [optional] 
**LinkAmountPaid** | Pointer to **float32** |  | [optional] 
**LinkPartialPayments** | Pointer to **bool** |  | [optional] 
**LinkMinimumPartialAmount** | Pointer to **float32** |  | [optional] 
**LinkPurpose** | Pointer to **string** |  | [optional] 
**LinkCreatedAt** | Pointer to **string** |  | [optional] 
**CustomerDetails** | Pointer to [**CFLinkCustomerDetailsEntity**](CFLinkCustomerDetailsEntity.md) |  | [optional] 
**LinkMeta** | Pointer to [**CFLinkMetaEntity**](CFLinkMetaEntity.md) |  | [optional] 
**LinkUrl** | Pointer to **string** |  | [optional] 
**LinkExpiryTime** | Pointer to **string** |  | [optional] 
**LinkNotes** | Pointer to **map[string]string** | Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs | [optional] 
**LinkAutoReminders** | Pointer to **bool** |  | [optional] 
**LinkNotify** | Pointer to [**CFLinkNotifyEntity**](CFLinkNotifyEntity.md) |  | [optional] 

## Methods

### NewCFLink

`func NewCFLink() *CFLink`

NewCFLink instantiates a new CFLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkWithDefaults

`func NewCFLinkWithDefaults() *CFLink`

NewCFLinkWithDefaults instantiates a new CFLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfLinkId

`func (o *CFLink) GetCfLinkId() int32`

GetCfLinkId returns the CfLinkId field if non-nil, zero value otherwise.

### GetCfLinkIdOk

`func (o *CFLink) GetCfLinkIdOk() (*int32, bool)`

GetCfLinkIdOk returns a tuple with the CfLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfLinkId

`func (o *CFLink) SetCfLinkId(v int32)`

SetCfLinkId sets CfLinkId field to given value.

### HasCfLinkId

`func (o *CFLink) HasCfLinkId() bool`

HasCfLinkId returns a boolean if a field has been set.

### GetLinkId

`func (o *CFLink) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *CFLink) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *CFLink) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *CFLink) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetLinkStatus

`func (o *CFLink) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *CFLink) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *CFLink) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *CFLink) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLinkCurrency

`func (o *CFLink) GetLinkCurrency() string`

GetLinkCurrency returns the LinkCurrency field if non-nil, zero value otherwise.

### GetLinkCurrencyOk

`func (o *CFLink) GetLinkCurrencyOk() (*string, bool)`

GetLinkCurrencyOk returns a tuple with the LinkCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCurrency

`func (o *CFLink) SetLinkCurrency(v string)`

SetLinkCurrency sets LinkCurrency field to given value.

### HasLinkCurrency

`func (o *CFLink) HasLinkCurrency() bool`

HasLinkCurrency returns a boolean if a field has been set.

### GetLinkAmount

`func (o *CFLink) GetLinkAmount() float32`

GetLinkAmount returns the LinkAmount field if non-nil, zero value otherwise.

### GetLinkAmountOk

`func (o *CFLink) GetLinkAmountOk() (*float32, bool)`

GetLinkAmountOk returns a tuple with the LinkAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmount

`func (o *CFLink) SetLinkAmount(v float32)`

SetLinkAmount sets LinkAmount field to given value.

### HasLinkAmount

`func (o *CFLink) HasLinkAmount() bool`

HasLinkAmount returns a boolean if a field has been set.

### GetLinkAmountPaid

`func (o *CFLink) GetLinkAmountPaid() float32`

GetLinkAmountPaid returns the LinkAmountPaid field if non-nil, zero value otherwise.

### GetLinkAmountPaidOk

`func (o *CFLink) GetLinkAmountPaidOk() (*float32, bool)`

GetLinkAmountPaidOk returns a tuple with the LinkAmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmountPaid

`func (o *CFLink) SetLinkAmountPaid(v float32)`

SetLinkAmountPaid sets LinkAmountPaid field to given value.

### HasLinkAmountPaid

`func (o *CFLink) HasLinkAmountPaid() bool`

HasLinkAmountPaid returns a boolean if a field has been set.

### GetLinkPartialPayments

`func (o *CFLink) GetLinkPartialPayments() bool`

GetLinkPartialPayments returns the LinkPartialPayments field if non-nil, zero value otherwise.

### GetLinkPartialPaymentsOk

`func (o *CFLink) GetLinkPartialPaymentsOk() (*bool, bool)`

GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPartialPayments

`func (o *CFLink) SetLinkPartialPayments(v bool)`

SetLinkPartialPayments sets LinkPartialPayments field to given value.

### HasLinkPartialPayments

`func (o *CFLink) HasLinkPartialPayments() bool`

HasLinkPartialPayments returns a boolean if a field has been set.

### GetLinkMinimumPartialAmount

`func (o *CFLink) GetLinkMinimumPartialAmount() float32`

GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field if non-nil, zero value otherwise.

### GetLinkMinimumPartialAmountOk

`func (o *CFLink) GetLinkMinimumPartialAmountOk() (*float32, bool)`

GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMinimumPartialAmount

`func (o *CFLink) SetLinkMinimumPartialAmount(v float32)`

SetLinkMinimumPartialAmount sets LinkMinimumPartialAmount field to given value.

### HasLinkMinimumPartialAmount

`func (o *CFLink) HasLinkMinimumPartialAmount() bool`

HasLinkMinimumPartialAmount returns a boolean if a field has been set.

### GetLinkPurpose

`func (o *CFLink) GetLinkPurpose() string`

GetLinkPurpose returns the LinkPurpose field if non-nil, zero value otherwise.

### GetLinkPurposeOk

`func (o *CFLink) GetLinkPurposeOk() (*string, bool)`

GetLinkPurposeOk returns a tuple with the LinkPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPurpose

`func (o *CFLink) SetLinkPurpose(v string)`

SetLinkPurpose sets LinkPurpose field to given value.

### HasLinkPurpose

`func (o *CFLink) HasLinkPurpose() bool`

HasLinkPurpose returns a boolean if a field has been set.

### GetLinkCreatedAt

`func (o *CFLink) GetLinkCreatedAt() string`

GetLinkCreatedAt returns the LinkCreatedAt field if non-nil, zero value otherwise.

### GetLinkCreatedAtOk

`func (o *CFLink) GetLinkCreatedAtOk() (*string, bool)`

GetLinkCreatedAtOk returns a tuple with the LinkCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCreatedAt

`func (o *CFLink) SetLinkCreatedAt(v string)`

SetLinkCreatedAt sets LinkCreatedAt field to given value.

### HasLinkCreatedAt

`func (o *CFLink) HasLinkCreatedAt() bool`

HasLinkCreatedAt returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *CFLink) GetCustomerDetails() CFLinkCustomerDetailsEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CFLink) GetCustomerDetailsOk() (*CFLinkCustomerDetailsEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CFLink) SetCustomerDetails(v CFLinkCustomerDetailsEntity)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *CFLink) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetLinkMeta

`func (o *CFLink) GetLinkMeta() CFLinkMetaEntity`

GetLinkMeta returns the LinkMeta field if non-nil, zero value otherwise.

### GetLinkMetaOk

`func (o *CFLink) GetLinkMetaOk() (*CFLinkMetaEntity, bool)`

GetLinkMetaOk returns a tuple with the LinkMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMeta

`func (o *CFLink) SetLinkMeta(v CFLinkMetaEntity)`

SetLinkMeta sets LinkMeta field to given value.

### HasLinkMeta

`func (o *CFLink) HasLinkMeta() bool`

HasLinkMeta returns a boolean if a field has been set.

### GetLinkUrl

`func (o *CFLink) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *CFLink) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *CFLink) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *CFLink) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetLinkExpiryTime

`func (o *CFLink) GetLinkExpiryTime() string`

GetLinkExpiryTime returns the LinkExpiryTime field if non-nil, zero value otherwise.

### GetLinkExpiryTimeOk

`func (o *CFLink) GetLinkExpiryTimeOk() (*string, bool)`

GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryTime

`func (o *CFLink) SetLinkExpiryTime(v string)`

SetLinkExpiryTime sets LinkExpiryTime field to given value.

### HasLinkExpiryTime

`func (o *CFLink) HasLinkExpiryTime() bool`

HasLinkExpiryTime returns a boolean if a field has been set.

### GetLinkNotes

`func (o *CFLink) GetLinkNotes() map[string]string`

GetLinkNotes returns the LinkNotes field if non-nil, zero value otherwise.

### GetLinkNotesOk

`func (o *CFLink) GetLinkNotesOk() (*map[string]string, bool)`

GetLinkNotesOk returns a tuple with the LinkNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotes

`func (o *CFLink) SetLinkNotes(v map[string]string)`

SetLinkNotes sets LinkNotes field to given value.

### HasLinkNotes

`func (o *CFLink) HasLinkNotes() bool`

HasLinkNotes returns a boolean if a field has been set.

### GetLinkAutoReminders

`func (o *CFLink) GetLinkAutoReminders() bool`

GetLinkAutoReminders returns the LinkAutoReminders field if non-nil, zero value otherwise.

### GetLinkAutoRemindersOk

`func (o *CFLink) GetLinkAutoRemindersOk() (*bool, bool)`

GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoReminders

`func (o *CFLink) SetLinkAutoReminders(v bool)`

SetLinkAutoReminders sets LinkAutoReminders field to given value.

### HasLinkAutoReminders

`func (o *CFLink) HasLinkAutoReminders() bool`

HasLinkAutoReminders returns a boolean if a field has been set.

### GetLinkNotify

`func (o *CFLink) GetLinkNotify() CFLinkNotifyEntity`

GetLinkNotify returns the LinkNotify field if non-nil, zero value otherwise.

### GetLinkNotifyOk

`func (o *CFLink) GetLinkNotifyOk() (*CFLinkNotifyEntity, bool)`

GetLinkNotifyOk returns a tuple with the LinkNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotify

`func (o *CFLink) SetLinkNotify(v CFLinkNotifyEntity)`

SetLinkNotify sets LinkNotify field to given value.

### HasLinkNotify

`func (o *CFLink) HasLinkNotify() bool`

HasLinkNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


