# PaymentMethodInPaymentsEntityPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**CardNumber** | Pointer to **string** |  | [optional] 
**CardNetwork** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**CardCountry** | Pointer to **string** |  | [optional] 
**CardBankName** | Pointer to **string** |  | [optional] 
**CardNetworkReferenceId** | Pointer to **string** |  | [optional] 
**NetbankingBankCode** | **int32** |  | 
**NetbankingBankName** | **string** |  | 
**UpiId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodInPaymentsEntityPaymentMethod

`func NewPaymentMethodInPaymentsEntityPaymentMethod(channel string, netbankingBankCode int32, netbankingBankName string, ) *PaymentMethodInPaymentsEntityPaymentMethod`

NewPaymentMethodInPaymentsEntityPaymentMethod instantiates a new PaymentMethodInPaymentsEntityPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults

`func NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults() *PaymentMethodInPaymentsEntityPaymentMethod`

NewPaymentMethodInPaymentsEntityPaymentMethodWithDefaults instantiates a new PaymentMethodInPaymentsEntityPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCardNumber

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNetwork

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardCountry

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardBankName

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetCardNetworkReferenceId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNetworkReferenceId() string`

GetCardNetworkReferenceId returns the CardNetworkReferenceId field if non-nil, zero value otherwise.

### GetCardNetworkReferenceIdOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetCardNetworkReferenceIdOk() (*string, bool)`

GetCardNetworkReferenceIdOk returns a tuple with the CardNetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetworkReferenceId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetCardNetworkReferenceId(v string)`

SetCardNetworkReferenceId sets CardNetworkReferenceId field to given value.

### HasCardNetworkReferenceId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasCardNetworkReferenceId() bool`

HasCardNetworkReferenceId returns a boolean if a field has been set.

### GetNetbankingBankCode

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbankingBankCode() int32`

GetNetbankingBankCode returns the NetbankingBankCode field if non-nil, zero value otherwise.

### GetNetbankingBankCodeOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbankingBankCodeOk() (*int32, bool)`

GetNetbankingBankCodeOk returns a tuple with the NetbankingBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankCode

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetNetbankingBankCode(v int32)`

SetNetbankingBankCode sets NetbankingBankCode field to given value.


### GetNetbankingBankName

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbankingBankName() string`

GetNetbankingBankName returns the NetbankingBankName field if non-nil, zero value otherwise.

### GetNetbankingBankNameOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetNetbankingBankNameOk() (*string, bool)`

GetNetbankingBankNameOk returns a tuple with the NetbankingBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankName

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetNetbankingBankName(v string)`

SetNetbankingBankName sets NetbankingBankName field to given value.


### GetUpiId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetProvider

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPhone

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PaymentMethodInPaymentsEntityPaymentMethod) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


