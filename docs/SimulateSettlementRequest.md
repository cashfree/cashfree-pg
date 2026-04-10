# SimulateSettlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderIds** | Pointer to **[]string** | A list of orders for which you want to simulate settlement. | [optional] 
**TxnTime** | Pointer to **string** | The start time (YYYY-MM-DD HH:mm:ss) from which transactions are picked for simulating settlement. You can pass a &#x60;txnTime&#x60; value for up to the last seven days. | [optional] 
**Status** | Pointer to **string** | The simulation status. Possible values are SUCCESS, FAILED or PENDING. | [optional] 
**ErrorReason** | Pointer to **string** | Specifies the reason for settlement failure. The default value is used if this is not provided. This is required only if the status is FAILED. | [optional] 

## Methods

### NewSimulateSettlementRequest

`func NewSimulateSettlementRequest() *SimulateSettlementRequest`

NewSimulateSettlementRequest instantiates a new SimulateSettlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateSettlementRequestWithDefaults

`func NewSimulateSettlementRequestWithDefaults() *SimulateSettlementRequest`

NewSimulateSettlementRequestWithDefaults instantiates a new SimulateSettlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderIds

`func (o *SimulateSettlementRequest) GetMerchantOrderIds() []string`

GetMerchantOrderIds returns the MerchantOrderIds field if non-nil, zero value otherwise.

### GetMerchantOrderIdsOk

`func (o *SimulateSettlementRequest) GetMerchantOrderIdsOk() (*[]string, bool)`

GetMerchantOrderIdsOk returns a tuple with the MerchantOrderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderIds

`func (o *SimulateSettlementRequest) SetMerchantOrderIds(v []string)`

SetMerchantOrderIds sets MerchantOrderIds field to given value.

### HasMerchantOrderIds

`func (o *SimulateSettlementRequest) HasMerchantOrderIds() bool`

HasMerchantOrderIds returns a boolean if a field has been set.

### GetTxnTime

`func (o *SimulateSettlementRequest) GetTxnTime() string`

GetTxnTime returns the TxnTime field if non-nil, zero value otherwise.

### GetTxnTimeOk

`func (o *SimulateSettlementRequest) GetTxnTimeOk() (*string, bool)`

GetTxnTimeOk returns a tuple with the TxnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnTime

`func (o *SimulateSettlementRequest) SetTxnTime(v string)`

SetTxnTime sets TxnTime field to given value.

### HasTxnTime

`func (o *SimulateSettlementRequest) HasTxnTime() bool`

HasTxnTime returns a boolean if a field has been set.

### GetStatus

`func (o *SimulateSettlementRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimulateSettlementRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimulateSettlementRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimulateSettlementRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorReason

`func (o *SimulateSettlementRequest) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *SimulateSettlementRequest) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *SimulateSettlementRequest) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *SimulateSettlementRequest) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


