# CreateLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | **string** | Unique Identifier (provided by merchant) for the Link. Alphanumeric and only - and _ allowed (50 character limit). Use this for other link-related APIs. | 
**LinkAmount** | **float64** | Amount to be collected using this link. Provide upto two decimals for paise. | 
**LinkCurrency** | **string** | Currency for the payment link. Default is INR. Contact care@cashfree.com to enable new currencies. | 
**LinkPurpose** | **string** | A brief description for which payment must be collected. This is shown to the customer. | 
**CustomerDetails** | [**LinkCustomerDetailsEntity**](LinkCustomerDetailsEntity.md) |  | 
**LinkPartialPayments** | Pointer to **bool** | If \&quot;true\&quot;, customer can make partial payments for the link. | [optional] 
**LinkMinimumPartialAmount** | Pointer to **float64** | Minimum amount in first installment that needs to be paid by the customer if partial payments are enabled. This should be less than the link_amount. | [optional] 
**LinkExpiryTime** | Pointer to **string** | Time after which the link expires. Customers will not be able to make the payment beyond the time specified here. You can provide them in a valid ISO 8601 time format. Default is 30 days. | [optional] 
**LinkNotify** | Pointer to [**LinkNotifyEntity**](LinkNotifyEntity.md) |  | [optional] 
**LinkAutoReminders** | Pointer to **bool** | If \&quot;true\&quot;, reminders will be sent to customers for collecting payments. | [optional] 
**LinkNotes** | Pointer to **map[string]string** | Key-value pair that can be used to store additional information about the entity. Maximum 5 key-value pairs | [optional] 
**LinkMeta** | Pointer to [**LinkMetaResponseEntity**](LinkMetaResponseEntity.md) |  | [optional] 
**OrderSplits** | Pointer to [**[]VendorSplit**](VendorSplit.md) | If you have Easy split enabled in your Cashfree account then you can use this option to split the order amount. | [optional] 

## Methods

### NewCreateLinkRequest

`func NewCreateLinkRequest(linkId string, linkAmount float64, linkCurrency string, linkPurpose string, customerDetails LinkCustomerDetailsEntity, ) *CreateLinkRequest`

NewCreateLinkRequest instantiates a new CreateLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkRequestWithDefaults

`func NewCreateLinkRequestWithDefaults() *CreateLinkRequest`

NewCreateLinkRequestWithDefaults instantiates a new CreateLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *CreateLinkRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *CreateLinkRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *CreateLinkRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetLinkAmount

`func (o *CreateLinkRequest) GetLinkAmount() float64`

GetLinkAmount returns the LinkAmount field if non-nil, zero value otherwise.

### GetLinkAmountOk

`func (o *CreateLinkRequest) GetLinkAmountOk() (*float64, bool)`

GetLinkAmountOk returns a tuple with the LinkAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAmount

`func (o *CreateLinkRequest) SetLinkAmount(v float64)`

SetLinkAmount sets LinkAmount field to given value.


### GetLinkCurrency

`func (o *CreateLinkRequest) GetLinkCurrency() string`

GetLinkCurrency returns the LinkCurrency field if non-nil, zero value otherwise.

### GetLinkCurrencyOk

`func (o *CreateLinkRequest) GetLinkCurrencyOk() (*string, bool)`

GetLinkCurrencyOk returns a tuple with the LinkCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCurrency

`func (o *CreateLinkRequest) SetLinkCurrency(v string)`

SetLinkCurrency sets LinkCurrency field to given value.


### GetLinkPurpose

`func (o *CreateLinkRequest) GetLinkPurpose() string`

GetLinkPurpose returns the LinkPurpose field if non-nil, zero value otherwise.

### GetLinkPurposeOk

`func (o *CreateLinkRequest) GetLinkPurposeOk() (*string, bool)`

GetLinkPurposeOk returns a tuple with the LinkPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPurpose

`func (o *CreateLinkRequest) SetLinkPurpose(v string)`

SetLinkPurpose sets LinkPurpose field to given value.


### GetCustomerDetails

`func (o *CreateLinkRequest) GetCustomerDetails() LinkCustomerDetailsEntity`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *CreateLinkRequest) GetCustomerDetailsOk() (*LinkCustomerDetailsEntity, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *CreateLinkRequest) SetCustomerDetails(v LinkCustomerDetailsEntity)`

SetCustomerDetails sets CustomerDetails field to given value.


### GetLinkPartialPayments

`func (o *CreateLinkRequest) GetLinkPartialPayments() bool`

GetLinkPartialPayments returns the LinkPartialPayments field if non-nil, zero value otherwise.

### GetLinkPartialPaymentsOk

`func (o *CreateLinkRequest) GetLinkPartialPaymentsOk() (*bool, bool)`

GetLinkPartialPaymentsOk returns a tuple with the LinkPartialPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPartialPayments

`func (o *CreateLinkRequest) SetLinkPartialPayments(v bool)`

SetLinkPartialPayments sets LinkPartialPayments field to given value.

### HasLinkPartialPayments

`func (o *CreateLinkRequest) HasLinkPartialPayments() bool`

HasLinkPartialPayments returns a boolean if a field has been set.

### GetLinkMinimumPartialAmount

`func (o *CreateLinkRequest) GetLinkMinimumPartialAmount() float64`

GetLinkMinimumPartialAmount returns the LinkMinimumPartialAmount field if non-nil, zero value otherwise.

### GetLinkMinimumPartialAmountOk

`func (o *CreateLinkRequest) GetLinkMinimumPartialAmountOk() (*float64, bool)`

GetLinkMinimumPartialAmountOk returns a tuple with the LinkMinimumPartialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMinimumPartialAmount

`func (o *CreateLinkRequest) SetLinkMinimumPartialAmount(v float64)`

SetLinkMinimumPartialAmount sets LinkMinimumPartialAmount field to given value.

### HasLinkMinimumPartialAmount

`func (o *CreateLinkRequest) HasLinkMinimumPartialAmount() bool`

HasLinkMinimumPartialAmount returns a boolean if a field has been set.

### GetLinkExpiryTime

`func (o *CreateLinkRequest) GetLinkExpiryTime() string`

GetLinkExpiryTime returns the LinkExpiryTime field if non-nil, zero value otherwise.

### GetLinkExpiryTimeOk

`func (o *CreateLinkRequest) GetLinkExpiryTimeOk() (*string, bool)`

GetLinkExpiryTimeOk returns a tuple with the LinkExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkExpiryTime

`func (o *CreateLinkRequest) SetLinkExpiryTime(v string)`

SetLinkExpiryTime sets LinkExpiryTime field to given value.

### HasLinkExpiryTime

`func (o *CreateLinkRequest) HasLinkExpiryTime() bool`

HasLinkExpiryTime returns a boolean if a field has been set.

### GetLinkNotify

`func (o *CreateLinkRequest) GetLinkNotify() LinkNotifyEntity`

GetLinkNotify returns the LinkNotify field if non-nil, zero value otherwise.

### GetLinkNotifyOk

`func (o *CreateLinkRequest) GetLinkNotifyOk() (*LinkNotifyEntity, bool)`

GetLinkNotifyOk returns a tuple with the LinkNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotify

`func (o *CreateLinkRequest) SetLinkNotify(v LinkNotifyEntity)`

SetLinkNotify sets LinkNotify field to given value.

### HasLinkNotify

`func (o *CreateLinkRequest) HasLinkNotify() bool`

HasLinkNotify returns a boolean if a field has been set.

### GetLinkAutoReminders

`func (o *CreateLinkRequest) GetLinkAutoReminders() bool`

GetLinkAutoReminders returns the LinkAutoReminders field if non-nil, zero value otherwise.

### GetLinkAutoRemindersOk

`func (o *CreateLinkRequest) GetLinkAutoRemindersOk() (*bool, bool)`

GetLinkAutoRemindersOk returns a tuple with the LinkAutoReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoReminders

`func (o *CreateLinkRequest) SetLinkAutoReminders(v bool)`

SetLinkAutoReminders sets LinkAutoReminders field to given value.

### HasLinkAutoReminders

`func (o *CreateLinkRequest) HasLinkAutoReminders() bool`

HasLinkAutoReminders returns a boolean if a field has been set.

### GetLinkNotes

`func (o *CreateLinkRequest) GetLinkNotes() map[string]string`

GetLinkNotes returns the LinkNotes field if non-nil, zero value otherwise.

### GetLinkNotesOk

`func (o *CreateLinkRequest) GetLinkNotesOk() (*map[string]string, bool)`

GetLinkNotesOk returns a tuple with the LinkNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkNotes

`func (o *CreateLinkRequest) SetLinkNotes(v map[string]string)`

SetLinkNotes sets LinkNotes field to given value.

### HasLinkNotes

`func (o *CreateLinkRequest) HasLinkNotes() bool`

HasLinkNotes returns a boolean if a field has been set.

### GetLinkMeta

`func (o *CreateLinkRequest) GetLinkMeta() LinkMetaResponseEntity`

GetLinkMeta returns the LinkMeta field if non-nil, zero value otherwise.

### GetLinkMetaOk

`func (o *CreateLinkRequest) GetLinkMetaOk() (*LinkMetaResponseEntity, bool)`

GetLinkMetaOk returns a tuple with the LinkMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkMeta

`func (o *CreateLinkRequest) SetLinkMeta(v LinkMetaResponseEntity)`

SetLinkMeta sets LinkMeta field to given value.

### HasLinkMeta

`func (o *CreateLinkRequest) HasLinkMeta() bool`

HasLinkMeta returns a boolean if a field has been set.

### GetOrderSplits

`func (o *CreateLinkRequest) GetOrderSplits() []VendorSplit`

GetOrderSplits returns the OrderSplits field if non-nil, zero value otherwise.

### GetOrderSplitsOk

`func (o *CreateLinkRequest) GetOrderSplitsOk() (*[]VendorSplit, bool)`

GetOrderSplitsOk returns a tuple with the OrderSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplits

`func (o *CreateLinkRequest) SetOrderSplits(v []VendorSplit)`

SetOrderSplits sets OrderSplits field to given value.

### HasOrderSplits

`func (o *CreateLinkRequest) HasOrderSplits() bool`

HasOrderSplits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


