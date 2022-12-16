# CFLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **string** | Unique Identifier (provided by merchant) for the Link. Alphanumeric and only - and _ allowed (50 character limit). Use this for other link-related APIs. | 
**LinkAmount** | **float64** | Amount to be collected using this link. Provide upto two decimals for paise. | 
**LinkCurrency** | **string** | Currency for the payment link. Default is INR. Contact care@cashfree.com to enable new currencies. | 
**LinkPurpose** | **string** | A brief description for which payment must be collected. This is shown to the customer. | 
**CustomerDetails** | [**CFLinkCustomerDetailsEntity**](CFLinkCustomerDetailsEntity.md) |  | 
**LinkPartialPayments** | Pointer to **bool** | If \&quot;true\&quot;, customer can make partial payments for the link. | [optional] 
**LinkMinimumPartialAmount** | Pointer to **float64** | Minimum amount in first installment that needs to be paid by the customer if partial payments are enabled. This should be less than the link_amount. | [optional] 
**LinkExpiryTime** | Pointer to **string** | Time after which the link expires. Customers will not be able to make the payment beyond the time specified here. You can provide them in a valid ISO 8601 time format. Default is 30 days. | [optional] 
**LinkNotify** | Pointer to [**CFLinkNotifyEntity**](CFLinkNotifyEntity.md) |  | [optional] 
**LinkAutoReminders** | Pointer to **bool** | If \&quot;true\&quot;, reminders will be sent to customers for collecting payments. | [optional] 
**LinkNotes** | Pointer to [**CFLinkNotesEntity**](CFLinkNotesEntity.md) |  | [optional] 
**LinkMeta** | Pointer to [**CFLinkMetaEntity**](CFLinkMetaEntity.md) |  | [optional] 

## Methods

### NewCFLinkRequest

`func NewCFLinkRequest(linkId string, linkAmount float64, linkCurrency string, linkPurpose string, customerDetails CFLinkCustomerDetailsEntity, ) *CFLinkRequest`

NewCFLinkRequest instantiates a new CFLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFLinkRequestWithDefaults

`func NewCFLinkRequestWithDefaults() *CFLinkRequest`

NewCFLinkRequestWithDefaults instantiates a new CFLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *CFLinkRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *CFLinkRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *CFLinkRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetLinkAmount

`func (o *CFLinkRequest) GetLinkAmount() float64`

GetLinkAmount returns the LinkAmount field if non-nil, zero value otherwise.

### GetLinkAmountOk

`func (o *CFLinkRequest) GetLinkAmountOk() (*float64, bool)`

GetLinkAmountOk returns a tuple with the LinkAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmount

`func (o *CFLinkRequest) SetLinkAmount(v float64)`

SetLinkAmount sets LinkAmount field to given value.


### GetLinkCurrency

`func (o *CFLinkRequest) GetLinkCurrency() string`

GetLinkCurrency returns the LinkCurrency field if non-nil, zero value otherwise.

### GetLinkCurrencyOk

`func (o *CFLinkRequest) GetLinkCurrencyOk() (*string, bool)`

GetLinkCurrencyOk returns a tuple with the LinkCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCurrency

`func (o *CFLinkRequest) SetLinkCurrency(v string)`

SetLinkCurrency sets LinkCurrency field to given value.


### GetLinkPurpose

`func (o *CFLinkRequest) GetLinkPurpose() string`

GetLinkPurpose returns the LinkPurpose field if non-nil, zero value otherwise.

### GetLinkPurposeOk

`func (o *CFLinkRequest) GetLinkPurposeOk() (*string, bool)`

GetLinkPurposeOk returns a tuple with the LinkPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPurpose

`func (o *CFLinkRequest) SetLinkPurpose(v string)`

SetLinkPurpose sets LinkPurpose field to given value.


### GetCustomerDetails

`func (o *CFLinkRequest) GetCustomerDetails() CFLinkCustomerDetailsEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CFLinkRequest) GetCustomerDetailsOk() (*CFLinkCustomerDetailsEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CFLinkRequest) SetCustomerDetails(v CFLinkCustomerDetailsEntity)`

SetCustomerDetails sets CustomerDetails field to given value.


### GetLinkPartialPayments

`func (o *CFLinkRequest) GetLinkPartialPayments() bool`

GetLinkPartialPayments returns the LinkPartialPayments field if non-nil, zero value otherwise.

### GetLinkPartialPaymentsOk

`func (o *CFLinkRequest) GetLinkPartialPaymentsOk() (*bool, bool)`

GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPartialPayments

`func (o *CFLinkRequest) SetLinkPartialPayments(v bool)`

SetLinkPartialPayments sets LinkPartialPayments field to given value.

### HasLinkPartialPayments

`func (o *CFLinkRequest) HasLinkPartialPayments() bool`

HasLinkPartialPayments returns a boolean if a field has been set.

### GetLinkMinimumPartialAmount

`func (o *CFLinkRequest) GetLinkMinimumPartialAmount() float64`

GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field if non-nil, zero value otherwise.

### GetLinkMinimumPartialAmountOk

`func (o *CFLinkRequest) GetLinkMinimumPartialAmountOk() (*float64, bool)`

GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMinimumPartialAmount

`func (o *CFLinkRequest) SetLinkMinimumPartialAmount(v float64)`

SetLinkMinimumPartialAmount sets LinkMinimumPartialAmount field to given value.

### HasLinkMinimumPartialAmount

`func (o *CFLinkRequest) HasLinkMinimumPartialAmount() bool`

HasLinkMinimumPartialAmount returns a boolean if a field has been set.

### GetLinkExpiryTime

`func (o *CFLinkRequest) GetLinkExpiryTime() string`

GetLinkExpiryTime returns the LinkExpiryTime field if non-nil, zero value otherwise.

### GetLinkExpiryTimeOk

`func (o *CFLinkRequest) GetLinkExpiryTimeOk() (*string, bool)`

GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryTime

`func (o *CFLinkRequest) SetLinkExpiryTime(v string)`

SetLinkExpiryTime sets LinkExpiryTime field to given value.

### HasLinkExpiryTime

`func (o *CFLinkRequest) HasLinkExpiryTime() bool`

HasLinkExpiryTime returns a boolean if a field has been set.

### GetLinkNotify

`func (o *CFLinkRequest) GetLinkNotify() CFLinkNotifyEntity`

GetLinkNotify returns the LinkNotify field if non-nil, zero value otherwise.

### GetLinkNotifyOk

`func (o *CFLinkRequest) GetLinkNotifyOk() (*CFLinkNotifyEntity, bool)`

GetLinkNotifyOk returns a tuple with the LinkNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotify

`func (o *CFLinkRequest) SetLinkNotify(v CFLinkNotifyEntity)`

SetLinkNotify sets LinkNotify field to given value.

### HasLinkNotify

`func (o *CFLinkRequest) HasLinkNotify() bool`

HasLinkNotify returns a boolean if a field has been set.

### GetLinkAutoReminders

`func (o *CFLinkRequest) GetLinkAutoReminders() bool`

GetLinkAutoReminders returns the LinkAutoReminders field if non-nil, zero value otherwise.

### GetLinkAutoRemindersOk

`func (o *CFLinkRequest) GetLinkAutoRemindersOk() (*bool, bool)`

GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoReminders

`func (o *CFLinkRequest) SetLinkAutoReminders(v bool)`

SetLinkAutoReminders sets LinkAutoReminders field to given value.

### HasLinkAutoReminders

`func (o *CFLinkRequest) HasLinkAutoReminders() bool`

HasLinkAutoReminders returns a boolean if a field has been set.

### GetLinkNotes

`func (o *CFLinkRequest) GetLinkNotes() CFLinkNotesEntity`

GetLinkNotes returns the LinkNotes field if non-nil, zero value otherwise.

### GetLinkNotesOk

`func (o *CFLinkRequest) GetLinkNotesOk() (*CFLinkNotesEntity, bool)`

GetLinkNotesOk returns a tuple with the LinkNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotes

`func (o *CFLinkRequest) SetLinkNotes(v CFLinkNotesEntity)`

SetLinkNotes sets LinkNotes field to given value.

### HasLinkNotes

`func (o *CFLinkRequest) HasLinkNotes() bool`

HasLinkNotes returns a boolean if a field has been set.

### GetLinkMeta

`func (o *CFLinkRequest) GetLinkMeta() CFLinkMetaEntity`

GetLinkMeta returns the LinkMeta field if non-nil, zero value otherwise.

### GetLinkMetaOk

`func (o *CFLinkRequest) GetLinkMetaOk() (*CFLinkMetaEntity, bool)`

GetLinkMetaOk returns a tuple with the LinkMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMeta

`func (o *CFLinkRequest) SetLinkMeta(v CFLinkMetaEntity)`

SetLinkMeta sets LinkMeta field to given value.

### HasLinkMeta

`func (o *CFLinkRequest) HasLinkMeta() bool`

HasLinkMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


