# VendorRecon200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**SettlementEligibilityTime** | Pointer to **string** |  | [optional] 
**MerchantOrderId** | Pointer to **string** |  | [optional] 
**TxTime** | Pointer to **string** |  | [optional] 
**SettlementId** | Pointer to **int32** |  | [optional] 
**Settled** | Pointer to **bool** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**Tax** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**MerchantVendorId** | Pointer to **string** |  | [optional] 
**AddedOnTime** | Pointer to **string** |  | [optional] 
**SettlementTime** | Pointer to **string** |  | [optional] 
**SettlementUtr** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Debit** | Pointer to **string** |  | [optional] 
**Credit** | Pointer to **string** |  | [optional] 
**RefundArn** | Pointer to **string** |  | [optional] 

## Methods

### NewVendorRecon200ResponseDataInner

`func NewVendorRecon200ResponseDataInner() *VendorRecon200ResponseDataInner`

NewVendorRecon200ResponseDataInner instantiates a new VendorRecon200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorRecon200ResponseDataInnerWithDefaults

`func NewVendorRecon200ResponseDataInnerWithDefaults() *VendorRecon200ResponseDataInner`

NewVendorRecon200ResponseDataInnerWithDefaults instantiates a new VendorRecon200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VendorRecon200ResponseDataInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VendorRecon200ResponseDataInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VendorRecon200ResponseDataInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VendorRecon200ResponseDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSettlementEligibilityTime

`func (o *VendorRecon200ResponseDataInner) GetSettlementEligibilityTime() string`

GetSettlementEligibilityTime returns the SettlementEligibilityTime field if non-nil, zero value otherwise.

### GetSettlementEligibilityTimeOk

`func (o *VendorRecon200ResponseDataInner) GetSettlementEligibilityTimeOk() (*string, bool)`

GetSettlementEligibilityTimeOk returns a tuple with the SettlementEligibilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementEligibilityTime

`func (o *VendorRecon200ResponseDataInner) SetSettlementEligibilityTime(v string)`

SetSettlementEligibilityTime sets SettlementEligibilityTime field to given value.

### HasSettlementEligibilityTime

`func (o *VendorRecon200ResponseDataInner) HasSettlementEligibilityTime() bool`

HasSettlementEligibilityTime returns a boolean if a field has been set.

### GetMerchantOrderId

`func (o *VendorRecon200ResponseDataInner) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *VendorRecon200ResponseDataInner) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *VendorRecon200ResponseDataInner) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.

### HasMerchantOrderId

`func (o *VendorRecon200ResponseDataInner) HasMerchantOrderId() bool`

HasMerchantOrderId returns a boolean if a field has been set.

### GetTxTime

`func (o *VendorRecon200ResponseDataInner) GetTxTime() string`

GetTxTime returns the TxTime field if non-nil, zero value otherwise.

### GetTxTimeOk

`func (o *VendorRecon200ResponseDataInner) GetTxTimeOk() (*string, bool)`

GetTxTimeOk returns a tuple with the TxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTime

`func (o *VendorRecon200ResponseDataInner) SetTxTime(v string)`

SetTxTime sets TxTime field to given value.

### HasTxTime

`func (o *VendorRecon200ResponseDataInner) HasTxTime() bool`

HasTxTime returns a boolean if a field has been set.

### GetSettlementId

`func (o *VendorRecon200ResponseDataInner) GetSettlementId() int32`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *VendorRecon200ResponseDataInner) GetSettlementIdOk() (*int32, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *VendorRecon200ResponseDataInner) SetSettlementId(v int32)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *VendorRecon200ResponseDataInner) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### GetSettled

`func (o *VendorRecon200ResponseDataInner) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *VendorRecon200ResponseDataInner) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *VendorRecon200ResponseDataInner) SetSettled(v bool)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *VendorRecon200ResponseDataInner) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetFee

`func (o *VendorRecon200ResponseDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *VendorRecon200ResponseDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *VendorRecon200ResponseDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *VendorRecon200ResponseDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetTax

`func (o *VendorRecon200ResponseDataInner) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *VendorRecon200ResponseDataInner) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *VendorRecon200ResponseDataInner) SetTax(v string)`

SetTax sets Tax field to given value.

### HasTax

`func (o *VendorRecon200ResponseDataInner) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetEntityId

`func (o *VendorRecon200ResponseDataInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *VendorRecon200ResponseDataInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *VendorRecon200ResponseDataInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *VendorRecon200ResponseDataInner) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetType

`func (o *VendorRecon200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorRecon200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorRecon200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorRecon200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMerchantVendorId

`func (o *VendorRecon200ResponseDataInner) GetMerchantVendorId() string`

GetMerchantVendorId returns the MerchantVendorId field if non-nil, zero value otherwise.

### GetMerchantVendorIdOk

`func (o *VendorRecon200ResponseDataInner) GetMerchantVendorIdOk() (*string, bool)`

GetMerchantVendorIdOk returns a tuple with the MerchantVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantVendorId

`func (o *VendorRecon200ResponseDataInner) SetMerchantVendorId(v string)`

SetMerchantVendorId sets MerchantVendorId field to given value.

### HasMerchantVendorId

`func (o *VendorRecon200ResponseDataInner) HasMerchantVendorId() bool`

HasMerchantVendorId returns a boolean if a field has been set.

### GetAddedOnTime

`func (o *VendorRecon200ResponseDataInner) GetAddedOnTime() string`

GetAddedOnTime returns the AddedOnTime field if non-nil, zero value otherwise.

### GetAddedOnTimeOk

`func (o *VendorRecon200ResponseDataInner) GetAddedOnTimeOk() (*string, bool)`

GetAddedOnTimeOk returns a tuple with the AddedOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOnTime

`func (o *VendorRecon200ResponseDataInner) SetAddedOnTime(v string)`

SetAddedOnTime sets AddedOnTime field to given value.

### HasAddedOnTime

`func (o *VendorRecon200ResponseDataInner) HasAddedOnTime() bool`

HasAddedOnTime returns a boolean if a field has been set.

### GetSettlementTime

`func (o *VendorRecon200ResponseDataInner) GetSettlementTime() string`

GetSettlementTime returns the SettlementTime field if non-nil, zero value otherwise.

### GetSettlementTimeOk

`func (o *VendorRecon200ResponseDataInner) GetSettlementTimeOk() (*string, bool)`

GetSettlementTimeOk returns a tuple with the SettlementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTime

`func (o *VendorRecon200ResponseDataInner) SetSettlementTime(v string)`

SetSettlementTime sets SettlementTime field to given value.

### HasSettlementTime

`func (o *VendorRecon200ResponseDataInner) HasSettlementTime() bool`

HasSettlementTime returns a boolean if a field has been set.

### GetSettlementUtr

`func (o *VendorRecon200ResponseDataInner) GetSettlementUtr() string`

GetSettlementUtr returns the SettlementUtr field if non-nil, zero value otherwise.

### GetSettlementUtrOk

`func (o *VendorRecon200ResponseDataInner) GetSettlementUtrOk() (*string, bool)`

GetSettlementUtrOk returns a tuple with the SettlementUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementUtr

`func (o *VendorRecon200ResponseDataInner) SetSettlementUtr(v string)`

SetSettlementUtr sets SettlementUtr field to given value.

### HasSettlementUtr

`func (o *VendorRecon200ResponseDataInner) HasSettlementUtr() bool`

HasSettlementUtr returns a boolean if a field has been set.

### GetCurrency

`func (o *VendorRecon200ResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VendorRecon200ResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VendorRecon200ResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VendorRecon200ResponseDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDebit

`func (o *VendorRecon200ResponseDataInner) GetDebit() string`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *VendorRecon200ResponseDataInner) GetDebitOk() (*string, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *VendorRecon200ResponseDataInner) SetDebit(v string)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *VendorRecon200ResponseDataInner) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetCredit

`func (o *VendorRecon200ResponseDataInner) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *VendorRecon200ResponseDataInner) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *VendorRecon200ResponseDataInner) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *VendorRecon200ResponseDataInner) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetRefundArn

`func (o *VendorRecon200ResponseDataInner) GetRefundArn() string`

GetRefundArn returns the RefundArn field if non-nil, zero value otherwise.

### GetRefundArnOk

`func (o *VendorRecon200ResponseDataInner) GetRefundArnOk() (*string, bool)`

GetRefundArnOk returns a tuple with the RefundArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundArn

`func (o *VendorRecon200ResponseDataInner) SetRefundArn(v string)`

SetRefundArn sets RefundArn field to given value.

### HasRefundArn

`func (o *VendorRecon200ResponseDataInner) HasRefundArn() bool`

HasRefundArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


