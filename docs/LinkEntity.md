# LinkEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfLinkId** | Pointer to **int64** |  | [optional] 
**LinkId** | Pointer to **string** |  | [optional] 
**LinkStatus** | Pointer to **string** |  | [optional] 
**LinkCurrency** | Pointer to **string** |  | [optional] 
**LinkAmount** | Pointer to **float32** |  | [optional] 
**LinkAmountPaid** | Pointer to **float32** |  | [optional] 
**LinkPartialPayments** | Pointer to **bool** |  | [optional] 
**LinkMinimumPartialAmount** | Pointer to **float32** |  | [optional] 
**LinkPurpose** | Pointer to **string** |  | [optional] 
**LinkCreatedAt** | Pointer to **string** |  | [optional] 
**CustomerDetails** | Pointer to [**LinkCustomerDetailsEntity**](LinkCustomerDetailsEntity.md) |  | [optional] 
**LinkMeta** | Pointer to **map[string]string** | Payment link meta information object. | [optional] 
**LinkUrl** | Pointer to **string** |  | [optional] 
**LinkExpiryTime** | Pointer to **string** |  | [optional] 
**LinkNotes** | Pointer to **map[string]string** | Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs | [optional] 
**LinkAutoReminders** | Pointer to **bool** |  | [optional] 
**LinkNotify** | Pointer to [**LinkNotifyEntity**](LinkNotifyEntity.md) |  | [optional] 

## Methods

### NewLinkEntity

`func NewLinkEntity() *LinkEntity`

NewLinkEntity instantiates a new LinkEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkEntityWithDefaults

`func NewLinkEntityWithDefaults() *LinkEntity`

NewLinkEntityWithDefaults instantiates a new LinkEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfLinkId

`func (o *LinkEntity) GetCfLinkId() int64`

GetCfLinkId returns the CfLinkId field if non-nil, zero value otherwise.

### GetCfLinkIdOk

`func (o *LinkEntity) GetCfLinkIdOk() (*int64, bool)`

GetCfLinkIdOk returns a tuple with the CfLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfLinkId

`func (o *LinkEntity) SetCfLinkId(v int64)`

SetCfLinkId sets CfLinkId field to given value.

### HasCfLinkId

`func (o *LinkEntity) HasCfLinkId() bool`

HasCfLinkId returns a boolean if a field has been set.

### GetLinkId

`func (o *LinkEntity) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *LinkEntity) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *LinkEntity) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *LinkEntity) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetLinkStatus

`func (o *LinkEntity) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *LinkEntity) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *LinkEntity) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *LinkEntity) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLinkCurrency

`func (o *LinkEntity) GetLinkCurrency() string`

GetLinkCurrency returns the LinkCurrency field if non-nil, zero value otherwise.

### GetLinkCurrencyOk

`func (o *LinkEntity) GetLinkCurrencyOk() (*string, bool)`

GetLinkCurrencyOk returns a tuple with the LinkCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCurrency

`func (o *LinkEntity) SetLinkCurrency(v string)`

SetLinkCurrency sets LinkCurrency field to given value.

### HasLinkCurrency

`func (o *LinkEntity) HasLinkCurrency() bool`

HasLinkCurrency returns a boolean if a field has been set.

### GetLinkAmount

`func (o *LinkEntity) GetLinkAmount() float32`

GetLinkAmount returns the LinkAmount field if non-nil, zero value otherwise.

### GetLinkAmountOk

`func (o *LinkEntity) GetLinkAmountOk() (*float32, bool)`

GetLinkAmountOk returns a tuple with the LinkAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmount

`func (o *LinkEntity) SetLinkAmount(v float32)`

SetLinkAmount sets LinkAmount field to given value.

### HasLinkAmount

`func (o *LinkEntity) HasLinkAmount() bool`

HasLinkAmount returns a boolean if a field has been set.

### GetLinkAmountPaid

`func (o *LinkEntity) GetLinkAmountPaid() float32`

GetLinkAmountPaid returns the LinkAmountPaid field if non-nil, zero value otherwise.

### GetLinkAmountPaidOk

`func (o *LinkEntity) GetLinkAmountPaidOk() (*float32, bool)`

GetLinkAmountPaidOk returns a tuple with the LinkAmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmountPaid

`func (o *LinkEntity) SetLinkAmountPaid(v float32)`

SetLinkAmountPaid sets LinkAmountPaid field to given value.

### HasLinkAmountPaid

`func (o *LinkEntity) HasLinkAmountPaid() bool`

HasLinkAmountPaid returns a boolean if a field has been set.

### GetLinkPartialPayments

`func (o *LinkEntity) GetLinkPartialPayments() bool`

GetLinkPartialPayments returns the LinkPartialPayments field if non-nil, zero value otherwise.

### GetLinkPartialPaymentsOk

`func (o *LinkEntity) GetLinkPartialPaymentsOk() (*bool, bool)`

GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPartialPayments

`func (o *LinkEntity) SetLinkPartialPayments(v bool)`

SetLinkPartialPayments sets LinkPartialPayments field to given value.

### HasLinkPartialPayments

`func (o *LinkEntity) HasLinkPartialPayments() bool`

HasLinkPartialPayments returns a boolean if a field has been set.

### GetLinkMinimumPartialAmount

`func (o *LinkEntity) GetLinkMinimumPartialAmount() float32`

GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field if non-nil, zero value otherwise.

### GetLinkMinimumPartialAmountOk

`func (o *LinkEntity) GetLinkMinimumPartialAmountOk() (*float32, bool)`

GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMinimumPartialAmount

`func (o *LinkEntity) SetLinkMinimumPartialAmount(v float32)`

SetLinkMinimumPartialAmount sets LinkMinimumPartialAmount field to given value.

### HasLinkMinimumPartialAmount

`func (o *LinkEntity) HasLinkMinimumPartialAmount() bool`

HasLinkMinimumPartialAmount returns a boolean if a field has been set.

### GetLinkPurpose

`func (o *LinkEntity) GetLinkPurpose() string`

GetLinkPurpose returns the LinkPurpose field if non-nil, zero value otherwise.

### GetLinkPurposeOk

`func (o *LinkEntity) GetLinkPurposeOk() (*string, bool)`

GetLinkPurposeOk returns a tuple with the LinkPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPurpose

`func (o *LinkEntity) SetLinkPurpose(v string)`

SetLinkPurpose sets LinkPurpose field to given value.

### HasLinkPurpose

`func (o *LinkEntity) HasLinkPurpose() bool`

HasLinkPurpose returns a boolean if a field has been set.

### GetLinkCreatedAt

`func (o *LinkEntity) GetLinkCreatedAt() string`

GetLinkCreatedAt returns the LinkCreatedAt field if non-nil, zero value otherwise.

### GetLinkCreatedAtOk

`func (o *LinkEntity) GetLinkCreatedAtOk() (*string, bool)`

GetLinkCreatedAtOk returns a tuple with the LinkCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCreatedAt

`func (o *LinkEntity) SetLinkCreatedAt(v string)`

SetLinkCreatedAt sets LinkCreatedAt field to given value.

### HasLinkCreatedAt

`func (o *LinkEntity) HasLinkCreatedAt() bool`

HasLinkCreatedAt returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *LinkEntity) GetCustomerDetails() LinkCustomerDetailsEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *LinkEntity) GetCustomerDetailsOk() (*LinkCustomerDetailsEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *LinkEntity) SetCustomerDetails(v LinkCustomerDetailsEntity)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *LinkEntity) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetLinkMeta

`func (o *LinkEntity) GetLinkMeta() map[string]string`

GetLinkMeta returns the LinkMeta field if non-nil, zero value otherwise.

### GetLinkMetaOk

`func (o *LinkEntity) GetLinkMetaOk() (*map[string]string, bool)`

GetLinkMetaOk returns a tuple with the LinkMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMeta

`func (o *LinkEntity) SetLinkMeta(v map[string]string)`

SetLinkMeta sets LinkMeta field to given value.

### HasLinkMeta

`func (o *LinkEntity) HasLinkMeta() bool`

HasLinkMeta returns a boolean if a field has been set.

### GetLinkUrl

`func (o *LinkEntity) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *LinkEntity) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *LinkEntity) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *LinkEntity) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetLinkExpiryTime

`func (o *LinkEntity) GetLinkExpiryTime() string`

GetLinkExpiryTime returns the LinkExpiryTime field if non-nil, zero value otherwise.

### GetLinkExpiryTimeOk

`func (o *LinkEntity) GetLinkExpiryTimeOk() (*string, bool)`

GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryTime

`func (o *LinkEntity) SetLinkExpiryTime(v string)`

SetLinkExpiryTime sets LinkExpiryTime field to given value.

### HasLinkExpiryTime

`func (o *LinkEntity) HasLinkExpiryTime() bool`

HasLinkExpiryTime returns a boolean if a field has been set.

### GetLinkNotes

`func (o *LinkEntity) GetLinkNotes() map[string]string`

GetLinkNotes returns the LinkNotes field if non-nil, zero value otherwise.

### GetLinkNotesOk

`func (o *LinkEntity) GetLinkNotesOk() (*map[string]string, bool)`

GetLinkNotesOk returns a tuple with the LinkNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotes

`func (o *LinkEntity) SetLinkNotes(v map[string]string)`

SetLinkNotes sets LinkNotes field to given value.

### HasLinkNotes

`func (o *LinkEntity) HasLinkNotes() bool`

HasLinkNotes returns a boolean if a field has been set.

### GetLinkAutoReminders

`func (o *LinkEntity) GetLinkAutoReminders() bool`

GetLinkAutoReminders returns the LinkAutoReminders field if non-nil, zero value otherwise.

### GetLinkAutoRemindersOk

`func (o *LinkEntity) GetLinkAutoRemindersOk() (*bool, bool)`

GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoReminders

`func (o *LinkEntity) SetLinkAutoReminders(v bool)`

SetLinkAutoReminders sets LinkAutoReminders field to given value.

### HasLinkAutoReminders

`func (o *LinkEntity) HasLinkAutoReminders() bool`

HasLinkAutoReminders returns a boolean if a field has been set.

### GetLinkNotify

`func (o *LinkEntity) GetLinkNotify() LinkNotifyEntity`

GetLinkNotify returns the LinkNotify field if non-nil, zero value otherwise.

### GetLinkNotifyOk

`func (o *LinkEntity) GetLinkNotifyOk() (*LinkNotifyEntity, bool)`

GetLinkNotifyOk returns a tuple with the LinkNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotify

`func (o *LinkEntity) SetLinkNotify(v LinkNotifyEntity)`

SetLinkNotify sets LinkNotify field to given value.

### HasLinkNotify

`func (o *LinkEntity) HasLinkNotify() bool`

HasLinkNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


