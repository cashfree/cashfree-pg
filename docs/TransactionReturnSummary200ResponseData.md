# TransactionReturnSummary200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **int32** | Payment ID for which the summary is generated. | [optional] 
**Status** | Pointer to **string** | Status of the summary generation (GENERATED or INITIALIZED). | [optional] 
**Link** | Pointer to **string** | Download link for the summary (if status is GENERATED). | [optional] 

## Methods

### NewTransactionReturnSummary200ResponseData

`func NewTransactionReturnSummary200ResponseData() *TransactionReturnSummary200ResponseData`

NewTransactionReturnSummary200ResponseData instantiates a new TransactionReturnSummary200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReturnSummary200ResponseDataWithDefaults

`func NewTransactionReturnSummary200ResponseDataWithDefaults() *TransactionReturnSummary200ResponseData`

NewTransactionReturnSummary200ResponseDataWithDefaults instantiates a new TransactionReturnSummary200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *TransactionReturnSummary200ResponseData) GetPaymentId() int32`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *TransactionReturnSummary200ResponseData) GetPaymentIdOk() (*int32, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *TransactionReturnSummary200ResponseData) SetPaymentId(v int32)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *TransactionReturnSummary200ResponseData) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionReturnSummary200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionReturnSummary200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionReturnSummary200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionReturnSummary200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLink

`func (o *TransactionReturnSummary200ResponseData) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *TransactionReturnSummary200ResponseData) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *TransactionReturnSummary200ResponseData) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *TransactionReturnSummary200ResponseData) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


