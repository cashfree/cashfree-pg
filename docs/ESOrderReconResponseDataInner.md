# ESOrderReconResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**SettlementEligibilityTime** | Pointer to **string** |  | [optional] 
**MerchantOrderId** | Pointer to **string** |  | [optional] 
**TxTime** | Pointer to **string** |  | [optional] 
**Settled** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**MerchantSettlementUtr** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**SaleType** | Pointer to **string** |  | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 
**CustomerEmail** | Pointer to **string** |  | [optional] 
**CustomerPhone** | Pointer to **string** |  | [optional] 
**MerchantVendorCommission** | Pointer to **string** |  | [optional] 
**SplitServiceCharge** | Pointer to **string** |  | [optional] 
**SplitServiceTax** | Pointer to **string** |  | [optional] 
**PgServiceTax** | Pointer to **string** |  | [optional] 
**PgServiceCharge** | Pointer to **string** |  | [optional] 
**PgChargePostpaid** | Pointer to **string** |  | [optional] 
**MerchantSettlementId** | Pointer to **string** |  | [optional] 
**AddedOn** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**SettlementInitiatedOn** | Pointer to **string** |  | [optional] 
**SettlementTime** | Pointer to **string** |  | [optional] 
**OrderSplits** | Pointer to [**[]ESOrderReconResponseDataInnerOrderSplitsInner**](ESOrderReconResponseDataInnerOrderSplitsInner.md) |  | [optional] 
**EligibleSplitBalance** | Pointer to **string** |  | [optional] 

## Methods

### NewESOrderReconResponseDataInner

`func NewESOrderReconResponseDataInner() *ESOrderReconResponseDataInner`

NewESOrderReconResponseDataInner instantiates a new ESOrderReconResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESOrderReconResponseDataInnerWithDefaults

`func NewESOrderReconResponseDataInnerWithDefaults() *ESOrderReconResponseDataInner`

NewESOrderReconResponseDataInnerWithDefaults instantiates a new ESOrderReconResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ESOrderReconResponseDataInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ESOrderReconResponseDataInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ESOrderReconResponseDataInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ESOrderReconResponseDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSettlementEligibilityTime

`func (o *ESOrderReconResponseDataInner) GetSettlementEligibilityTime() string`

GetSettlementEligibilityTime returns the SettlementEligibilityTime field if non-nil, zero value otherwise.

### GetSettlementEligibilityTimeOk

`func (o *ESOrderReconResponseDataInner) GetSettlementEligibilityTimeOk() (*string, bool)`

GetSettlementEligibilityTimeOk returns a tuple with the SettlementEligibilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementEligibilityTime

`func (o *ESOrderReconResponseDataInner) SetSettlementEligibilityTime(v string)`

SetSettlementEligibilityTime sets SettlementEligibilityTime field to given value.

### HasSettlementEligibilityTime

`func (o *ESOrderReconResponseDataInner) HasSettlementEligibilityTime() bool`

HasSettlementEligibilityTime returns a boolean if a field has been set.

### GetMerchantOrderId

`func (o *ESOrderReconResponseDataInner) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *ESOrderReconResponseDataInner) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *ESOrderReconResponseDataInner) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.

### HasMerchantOrderId

`func (o *ESOrderReconResponseDataInner) HasMerchantOrderId() bool`

HasMerchantOrderId returns a boolean if a field has been set.

### GetTxTime

`func (o *ESOrderReconResponseDataInner) GetTxTime() string`

GetTxTime returns the TxTime field if non-nil, zero value otherwise.

### GetTxTimeOk

`func (o *ESOrderReconResponseDataInner) GetTxTimeOk() (*string, bool)`

GetTxTimeOk returns a tuple with the TxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTime

`func (o *ESOrderReconResponseDataInner) SetTxTime(v string)`

SetTxTime sets TxTime field to given value.

### HasTxTime

`func (o *ESOrderReconResponseDataInner) HasTxTime() bool`

HasTxTime returns a boolean if a field has been set.

### GetSettled

`func (o *ESOrderReconResponseDataInner) GetSettled() string`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *ESOrderReconResponseDataInner) GetSettledOk() (*string, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *ESOrderReconResponseDataInner) SetSettled(v string)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *ESOrderReconResponseDataInner) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetEntityId

`func (o *ESOrderReconResponseDataInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ESOrderReconResponseDataInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ESOrderReconResponseDataInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ESOrderReconResponseDataInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMerchantSettlementUtr

`func (o *ESOrderReconResponseDataInner) GetMerchantSettlementUtr() string`

GetMerchantSettlementUtr returns the MerchantSettlementUtr field if non-nil, zero value otherwise.

### GetMerchantSettlementUtrOk

`func (o *ESOrderReconResponseDataInner) GetMerchantSettlementUtrOk() (*string, bool)`

GetMerchantSettlementUtrOk returns a tuple with the MerchantSettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantSettlementUtr

`func (o *ESOrderReconResponseDataInner) SetMerchantSettlementUtr(v string)`

SetMerchantSettlementUtr sets MerchantSettlementUtr field to given value.

### HasMerchantSettlementUtr

`func (o *ESOrderReconResponseDataInner) HasMerchantSettlementUtr() bool`

HasMerchantSettlementUtr returns a boolean if a field has been set.

### GetCurrency

`func (o *ESOrderReconResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ESOrderReconResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ESOrderReconResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ESOrderReconResponseDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSaleType

`func (o *ESOrderReconResponseDataInner) GetSaleType() string`

GetSaleType returns the SaleType field if non-nil, zero value otherwise.

### GetSaleTypeOk

`func (o *ESOrderReconResponseDataInner) GetSaleTypeOk() (*string, bool)`

GetSaleTypeOk returns a tuple with the SaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleType

`func (o *ESOrderReconResponseDataInner) SetSaleType(v string)`

SetSaleType sets SaleType field to given value.

### HasSaleType

`func (o *ESOrderReconResponseDataInner) HasSaleType() bool`

HasSaleType returns a boolean if a field has been set.

### GetCustomerName

`func (o *ESOrderReconResponseDataInner) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *ESOrderReconResponseDataInner) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *ESOrderReconResponseDataInner) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *ESOrderReconResponseDataInner) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *ESOrderReconResponseDataInner) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *ESOrderReconResponseDataInner) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *ESOrderReconResponseDataInner) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *ESOrderReconResponseDataInner) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *ESOrderReconResponseDataInner) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *ESOrderReconResponseDataInner) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *ESOrderReconResponseDataInner) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *ESOrderReconResponseDataInner) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetMerchantVendorCommission

`func (o *ESOrderReconResponseDataInner) GetMerchantVendorCommission() string`

GetMerchantVendorCommission returns the MerchantVendorCommission field if non-nil, zero value otherwise.

### GetMerchantVendorCommissionOk

`func (o *ESOrderReconResponseDataInner) GetMerchantVendorCommissionOk() (*string, bool)`

GetMerchantVendorCommissionOk returns a tuple with the MerchantVendorCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantVendorCommission

`func (o *ESOrderReconResponseDataInner) SetMerchantVendorCommission(v string)`

SetMerchantVendorCommission sets MerchantVendorCommission field to given value.

### HasMerchantVendorCommission

`func (o *ESOrderReconResponseDataInner) HasMerchantVendorCommission() bool`

HasMerchantVendorCommission returns a boolean if a field has been set.

### GetSplitServiceCharge

`func (o *ESOrderReconResponseDataInner) GetSplitServiceCharge() string`

GetSplitServiceCharge returns the SplitServiceCharge field if non-nil, zero value otherwise.

### GetSplitServiceChargeOk

`func (o *ESOrderReconResponseDataInner) GetSplitServiceChargeOk() (*string, bool)`

GetSplitServiceChargeOk returns a tuple with the SplitServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceCharge

`func (o *ESOrderReconResponseDataInner) SetSplitServiceCharge(v string)`

SetSplitServiceCharge sets SplitServiceCharge field to given value.

### HasSplitServiceCharge

`func (o *ESOrderReconResponseDataInner) HasSplitServiceCharge() bool`

HasSplitServiceCharge returns a boolean if a field has been set.

### GetSplitServiceTax

`func (o *ESOrderReconResponseDataInner) GetSplitServiceTax() string`

GetSplitServiceTax returns the SplitServiceTax field if non-nil, zero value otherwise.

### GetSplitServiceTaxOk

`func (o *ESOrderReconResponseDataInner) GetSplitServiceTaxOk() (*string, bool)`

GetSplitServiceTaxOk returns a tuple with the SplitServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitServiceTax

`func (o *ESOrderReconResponseDataInner) SetSplitServiceTax(v string)`

SetSplitServiceTax sets SplitServiceTax field to given value.

### HasSplitServiceTax

`func (o *ESOrderReconResponseDataInner) HasSplitServiceTax() bool`

HasSplitServiceTax returns a boolean if a field has been set.

### GetPgServiceTax

`func (o *ESOrderReconResponseDataInner) GetPgServiceTax() string`

GetPgServiceTax returns the PgServiceTax field if non-nil, zero value otherwise.

### GetPgServiceTaxOk

`func (o *ESOrderReconResponseDataInner) GetPgServiceTaxOk() (*string, bool)`

GetPgServiceTaxOk returns a tuple with the PgServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceTax

`func (o *ESOrderReconResponseDataInner) SetPgServiceTax(v string)`

SetPgServiceTax sets PgServiceTax field to given value.

### HasPgServiceTax

`func (o *ESOrderReconResponseDataInner) HasPgServiceTax() bool`

HasPgServiceTax returns a boolean if a field has been set.

### GetPgServiceCharge

`func (o *ESOrderReconResponseDataInner) GetPgServiceCharge() string`

GetPgServiceCharge returns the PgServiceCharge field if non-nil, zero value otherwise.

### GetPgServiceChargeOk

`func (o *ESOrderReconResponseDataInner) GetPgServiceChargeOk() (*string, bool)`

GetPgServiceChargeOk returns a tuple with the PgServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgServiceCharge

`func (o *ESOrderReconResponseDataInner) SetPgServiceCharge(v string)`

SetPgServiceCharge sets PgServiceCharge field to given value.

### HasPgServiceCharge

`func (o *ESOrderReconResponseDataInner) HasPgServiceCharge() bool`

HasPgServiceCharge returns a boolean if a field has been set.

### GetPgChargePostpaid

`func (o *ESOrderReconResponseDataInner) GetPgChargePostpaid() string`

GetPgChargePostpaid returns the PgChargePostpaid field if non-nil, zero value otherwise.

### GetPgChargePostpaidOk

`func (o *ESOrderReconResponseDataInner) GetPgChargePostpaidOk() (*string, bool)`

GetPgChargePostpaidOk returns a tuple with the PgChargePostpaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgChargePostpaid

`func (o *ESOrderReconResponseDataInner) SetPgChargePostpaid(v string)`

SetPgChargePostpaid sets PgChargePostpaid field to given value.

### HasPgChargePostpaid

`func (o *ESOrderReconResponseDataInner) HasPgChargePostpaid() bool`

HasPgChargePostpaid returns a boolean if a field has been set.

### GetMerchantSettlementId

`func (o *ESOrderReconResponseDataInner) GetMerchantSettlementId() string`

GetMerchantSettlementId returns the MerchantSettlementId field if non-nil, zero value otherwise.

### GetMerchantSettlementIdOk

`func (o *ESOrderReconResponseDataInner) GetMerchantSettlementIdOk() (*string, bool)`

GetMerchantSettlementIdOk returns a tuple with the MerchantSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantSettlementId

`func (o *ESOrderReconResponseDataInner) SetMerchantSettlementId(v string)`

SetMerchantSettlementId sets MerchantSettlementId field to given value.

### HasMerchantSettlementId

`func (o *ESOrderReconResponseDataInner) HasMerchantSettlementId() bool`

HasMerchantSettlementId returns a boolean if a field has been set.

### GetAddedOn

`func (o *ESOrderReconResponseDataInner) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *ESOrderReconResponseDataInner) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *ESOrderReconResponseDataInner) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *ESOrderReconResponseDataInner) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetTags

`func (o *ESOrderReconResponseDataInner) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ESOrderReconResponseDataInner) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ESOrderReconResponseDataInner) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ESOrderReconResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEntityType

`func (o *ESOrderReconResponseDataInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ESOrderReconResponseDataInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ESOrderReconResponseDataInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ESOrderReconResponseDataInner) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetSettlementInitiatedOn

`func (o *ESOrderReconResponseDataInner) GetSettlementInitiatedOn() string`

GetSettlementInitiatedOn returns the SettlementInitiatedOn field if non-nil, zero value otherwise.

### GetSettlementInitiatedOnOk

`func (o *ESOrderReconResponseDataInner) GetSettlementInitiatedOnOk() (*string, bool)`

GetSettlementInitiatedOnOk returns a tuple with the SettlementInitiatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInitiatedOn

`func (o *ESOrderReconResponseDataInner) SetSettlementInitiatedOn(v string)`

SetSettlementInitiatedOn sets SettlementInitiatedOn field to given value.

### HasSettlementInitiatedOn

`func (o *ESOrderReconResponseDataInner) HasSettlementInitiatedOn() bool`

HasSettlementInitiatedOn returns a boolean if a field has been set.

### GetSettlementTime

`func (o *ESOrderReconResponseDataInner) GetSettlementTime() string`

GetSettlementTime returns the SettlementTime field if non-nil, zero value otherwise.

### GetSettlementTimeOk

`func (o *ESOrderReconResponseDataInner) GetSettlementTimeOk() (*string, bool)`

GetSettlementTimeOk returns a tuple with the SettlementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTime

`func (o *ESOrderReconResponseDataInner) SetSettlementTime(v string)`

SetSettlementTime sets SettlementTime field to given value.

### HasSettlementTime

`func (o *ESOrderReconResponseDataInner) HasSettlementTime() bool`

HasSettlementTime returns a boolean if a field has been set.

### GetOrderSplits

`func (o *ESOrderReconResponseDataInner) GetOrderSplits() []ESOrderReconResponseDataInnerOrderSplitsInner`

GetOrderSplits returns the OrderSplits field if non-nil, zero value otherwise.

### GetOrderSplitsOk

`func (o *ESOrderReconResponseDataInner) GetOrderSplitsOk() (*[]ESOrderReconResponseDataInnerOrderSplitsInner, bool)`

GetOrderSplitsOk returns a tuple with the OrderSplits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSplits

`func (o *ESOrderReconResponseDataInner) SetOrderSplits(v []ESOrderReconResponseDataInnerOrderSplitsInner)`

SetOrderSplits sets OrderSplits field to given value.

### HasOrderSplits

`func (o *ESOrderReconResponseDataInner) HasOrderSplits() bool`

HasOrderSplits returns a boolean if a field has been set.

### GetEligibleSplitBalance

`func (o *ESOrderReconResponseDataInner) GetEligibleSplitBalance() string`

GetEligibleSplitBalance returns the EligibleSplitBalance field if non-nil, zero value otherwise.

### GetEligibleSplitBalanceOk

`func (o *ESOrderReconResponseDataInner) GetEligibleSplitBalanceOk() (*string, bool)`

GetEligibleSplitBalanceOk returns a tuple with the EligibleSplitBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleSplitBalance

`func (o *ESOrderReconResponseDataInner) SetEligibleSplitBalance(v string)`

SetEligibleSplitBalance sets EligibleSplitBalance field to given value.

### HasEligibleSplitBalance

`func (o *ESOrderReconResponseDataInner) HasEligibleSplitBalance() bool`

HasEligibleSplitBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


