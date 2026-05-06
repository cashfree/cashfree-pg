# TransactionReturnSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TransactionReturnSummary200ResponseData**](TransactionReturnSummary200ResponseData.md) |  | [optional] 
**Message** | Pointer to **string** | Status message. | [optional] 
**Status** | Pointer to **int32** | Status code of the response. | [optional] 

## Methods

### NewTransactionReturnSummary200Response

`func NewTransactionReturnSummary200Response() *TransactionReturnSummary200Response`

NewTransactionReturnSummary200Response instantiates a new TransactionReturnSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReturnSummary200ResponseWithDefaults

`func NewTransactionReturnSummary200ResponseWithDefaults() *TransactionReturnSummary200Response`

NewTransactionReturnSummary200ResponseWithDefaults instantiates a new TransactionReturnSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionReturnSummary200Response) GetData() TransactionReturnSummary200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionReturnSummary200Response) GetDataOk() (*TransactionReturnSummary200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionReturnSummary200Response) SetData(v TransactionReturnSummary200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionReturnSummary200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *TransactionReturnSummary200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransactionReturnSummary200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransactionReturnSummary200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransactionReturnSummary200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionReturnSummary200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionReturnSummary200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionReturnSummary200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionReturnSummary200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


