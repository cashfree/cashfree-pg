# CashbackDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashbackType** | **string** | Type of discount | 
**CashbackValue** | **float32** | Value of Discount. | 
**MaxCashbackAmount** | **float32** | Maximum Value of Cashback allowed. | 

## Methods

### NewCashbackDetails

`func NewCashbackDetails(cashbackType string, cashbackValue float32, maxCashbackAmount float32, ) *CashbackDetails`

NewCashbackDetails instantiates a new CashbackDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashbackDetailsWithDefaults

`func NewCashbackDetailsWithDefaults() *CashbackDetails`

NewCashbackDetailsWithDefaults instantiates a new CashbackDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashbackType

`func (o *CashbackDetails) GetCashbackType() string`

GetCashbackType returns the CashbackType field if non-nil, zero value otherwise.

### GetCashbackTypeOk

`func (o *CashbackDetails) GetCashbackTypeOk() (*string, bool)`

GetCashbackTypeOk returns a tuple with the CashbackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackType

`func (o *CashbackDetails) SetCashbackType(v string)`

SetCashbackType sets CashbackType field to given value.


### GetCashbackValue

`func (o *CashbackDetails) GetCashbackValue() float32`

GetCashbackValue returns the CashbackValue field if non-nil, zero value otherwise.

### GetCashbackValueOk

`func (o *CashbackDetails) GetCashbackValueOk() (*float32, bool)`

GetCashbackValueOk returns a tuple with the CashbackValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackValue

`func (o *CashbackDetails) SetCashbackValue(v float32)`

SetCashbackValue sets CashbackValue field to given value.


### GetMaxCashbackAmount

`func (o *CashbackDetails) GetMaxCashbackAmount() float32`

GetMaxCashbackAmount returns the MaxCashbackAmount field if non-nil, zero value otherwise.

### GetMaxCashbackAmountOk

`func (o *CashbackDetails) GetMaxCashbackAmountOk() (*float32, bool)`

GetMaxCashbackAmountOk returns a tuple with the MaxCashbackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCashbackAmount

`func (o *CashbackDetails) SetMaxCashbackAmount(v float32)`

SetMaxCashbackAmount sets MaxCashbackAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


