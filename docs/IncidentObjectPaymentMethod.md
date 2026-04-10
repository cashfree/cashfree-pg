# IncidentObjectPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upi** | Pointer to [**UPIDowntimeUpi**](UPIDowntimeUpi.md) |  | [optional] 
**Card** | Pointer to [**CardDowntimeCard**](CardDowntimeCard.md) |  | [optional] 
**NetBanking** | Pointer to [**NetBankingDowntimeNetBanking**](NetBankingDowntimeNetBanking.md) |  | [optional] 
**Wallet** | Pointer to [**WalletDowntimeWallet**](WalletDowntimeWallet.md) |  | [optional] 

## Methods

### NewIncidentObjectPaymentMethod

`func NewIncidentObjectPaymentMethod() *IncidentObjectPaymentMethod`

NewIncidentObjectPaymentMethod instantiates a new IncidentObjectPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentObjectPaymentMethodWithDefaults

`func NewIncidentObjectPaymentMethodWithDefaults() *IncidentObjectPaymentMethod`

NewIncidentObjectPaymentMethodWithDefaults instantiates a new IncidentObjectPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpi

`func (o *IncidentObjectPaymentMethod) GetUpi() UPIDowntimeUpi`

GetUpi returns the Upi field if non-nil, zero value otherwise.

### GetUpiOk

`func (o *IncidentObjectPaymentMethod) GetUpiOk() (*UPIDowntimeUpi, bool)`

GetUpiOk returns a tuple with the Upi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpi

`func (o *IncidentObjectPaymentMethod) SetUpi(v UPIDowntimeUpi)`

SetUpi sets Upi field to given value.

### HasUpi

`func (o *IncidentObjectPaymentMethod) HasUpi() bool`

HasUpi returns a boolean if a field has been set.

### GetCard

`func (o *IncidentObjectPaymentMethod) GetCard() CardDowntimeCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *IncidentObjectPaymentMethod) GetCardOk() (*CardDowntimeCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *IncidentObjectPaymentMethod) SetCard(v CardDowntimeCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *IncidentObjectPaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetNetBanking

`func (o *IncidentObjectPaymentMethod) GetNetBanking() NetBankingDowntimeNetBanking`

GetNetBanking returns the NetBanking field if non-nil, zero value otherwise.

### GetNetBankingOk

`func (o *IncidentObjectPaymentMethod) GetNetBankingOk() (*NetBankingDowntimeNetBanking, bool)`

GetNetBankingOk returns a tuple with the NetBanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBanking

`func (o *IncidentObjectPaymentMethod) SetNetBanking(v NetBankingDowntimeNetBanking)`

SetNetBanking sets NetBanking field to given value.

### HasNetBanking

`func (o *IncidentObjectPaymentMethod) HasNetBanking() bool`

HasNetBanking returns a boolean if a field has been set.

### GetWallet

`func (o *IncidentObjectPaymentMethod) GetWallet() WalletDowntimeWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *IncidentObjectPaymentMethod) GetWalletOk() (*WalletDowntimeWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *IncidentObjectPaymentMethod) SetWallet(v WalletDowntimeWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *IncidentObjectPaymentMethod) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


