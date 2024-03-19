# AdjustVendorBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementId** | Pointer to **float32** |  | [optional] 
**TransferDetails** | Pointer to [**TransferDetails**](TransferDetails.md) |  | [optional] 
**Balances** | Pointer to [**BalanceDetails**](BalanceDetails.md) |  | [optional] 
**Charges** | Pointer to [**ChargesDetails**](ChargesDetails.md) |  | [optional] 

## Methods

### NewAdjustVendorBalanceResponse

`func NewAdjustVendorBalanceResponse() *AdjustVendorBalanceResponse`

NewAdjustVendorBalanceResponse instantiates a new AdjustVendorBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustVendorBalanceResponseWithDefaults

`func NewAdjustVendorBalanceResponseWithDefaults() *AdjustVendorBalanceResponse`

NewAdjustVendorBalanceResponseWithDefaults instantiates a new AdjustVendorBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementId

`func (o *AdjustVendorBalanceResponse) GetSettlementId() float32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *AdjustVendorBalanceResponse) GetSettlementIdOk() (*float32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *AdjustVendorBalanceResponse) SetSettlementId(v float32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *AdjustVendorBalanceResponse) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetTransferDetails

`func (o *AdjustVendorBalanceResponse) GetTransferDetails() TransferDetails`

GetTransferDetails returns the TransferDetails field if non-nil, zero value otherwise.

### GetTransferDetailsOk

`func (o *AdjustVendorBalanceResponse) GetTransferDetailsOk() (*TransferDetails, bool)`

GetTransferDetailsOk returns a tuple with the TransferDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferDetails

`func (o *AdjustVendorBalanceResponse) SetTransferDetails(v TransferDetails)`

SetTransferDetails sets TransferDetails field to given value.

### HasTransferDetails

`func (o *AdjustVendorBalanceResponse) HasTransferDetails() bool`

HasTransferDetails returns a boolean if a field has been set.

### GetBalances

`func (o *AdjustVendorBalanceResponse) GetBalances() BalanceDetails`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AdjustVendorBalanceResponse) GetBalancesOk() (*BalanceDetails, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AdjustVendorBalanceResponse) SetBalances(v BalanceDetails)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AdjustVendorBalanceResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetCharges

`func (o *AdjustVendorBalanceResponse) GetCharges() ChargesDetails`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *AdjustVendorBalanceResponse) GetChargesOk() (*ChargesDetails, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *AdjustVendorBalanceResponse) SetCharges(v ChargesDetails)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *AdjustVendorBalanceResponse) HasCharges() bool`

HasCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


