# CFPaymentsEntityMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**CardNumber** | Pointer to **string** |  | [optional] 
**CardNetwork** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**CardCountry** | Pointer to **string** |  | [optional] 
**CardBankName** | Pointer to **string** |  | [optional] 
**NetbankingBankCode** | **string** |  | 
**NetbankingBankName** | **string** |  | 
**UpiId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewCFPaymentsEntityMethod

`func NewCFPaymentsEntityMethod(channel string, netbankingBankCode string, netbankingBankName string, ) *CFPaymentsEntityMethod`

NewCFPaymentsEntityMethod instantiates a new CFPaymentsEntityMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFPaymentsEntityMethodWithDefaults

`func NewCFPaymentsEntityMethodWithDefaults() *CFPaymentsEntityMethod`

NewCFPaymentsEntityMethodWithDefaults instantiates a new CFPaymentsEntityMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CFPaymentsEntityMethod) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CFPaymentsEntityMethod) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CFPaymentsEntityMethod) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCardNumber

`func (o *CFPaymentsEntityMethod) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *CFPaymentsEntityMethod) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *CFPaymentsEntityMethod) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *CFPaymentsEntityMethod) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNetwork

`func (o *CFPaymentsEntityMethod) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *CFPaymentsEntityMethod) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *CFPaymentsEntityMethod) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *CFPaymentsEntityMethod) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardType

`func (o *CFPaymentsEntityMethod) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CFPaymentsEntityMethod) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CFPaymentsEntityMethod) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CFPaymentsEntityMethod) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardCountry

`func (o *CFPaymentsEntityMethod) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *CFPaymentsEntityMethod) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *CFPaymentsEntityMethod) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *CFPaymentsEntityMethod) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardBankName

`func (o *CFPaymentsEntityMethod) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *CFPaymentsEntityMethod) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *CFPaymentsEntityMethod) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *CFPaymentsEntityMethod) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetNetbankingBankCode

`func (o *CFPaymentsEntityMethod) GetNetbankingBankCode() string`

GetNetbankingBankCode returns the NetbankingBankCode field if non-nil, zero value otherwise.

### GetNetbankingBankCodeOk

`func (o *CFPaymentsEntityMethod) GetNetbankingBankCodeOk() (*string, bool)`

GetNetbankingBankCodeOk returns a tuple with the NetbankingBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankCode

`func (o *CFPaymentsEntityMethod) SetNetbankingBankCode(v string)`

SetNetbankingBankCode sets NetbankingBankCode field to given value.


### GetNetbankingBankName

`func (o *CFPaymentsEntityMethod) GetNetbankingBankName() string`

GetNetbankingBankName returns the NetbankingBankName field if non-nil, zero value otherwise.

### GetNetbankingBankNameOk

`func (o *CFPaymentsEntityMethod) GetNetbankingBankNameOk() (*string, bool)`

GetNetbankingBankNameOk returns a tuple with the NetbankingBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbankingBankName

`func (o *CFPaymentsEntityMethod) SetNetbankingBankName(v string)`

SetNetbankingBankName sets NetbankingBankName field to given value.


### GetUpiId

`func (o *CFPaymentsEntityMethod) GetUpiId() string`

GetUpiId returns the UpiId field if non-nil, zero value otherwise.

### GetUpiIdOk

`func (o *CFPaymentsEntityMethod) GetUpiIdOk() (*string, bool)`

GetUpiIdOk returns a tuple with the UpiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiId

`func (o *CFPaymentsEntityMethod) SetUpiId(v string)`

SetUpiId sets UpiId field to given value.

### HasUpiId

`func (o *CFPaymentsEntityMethod) HasUpiId() bool`

HasUpiId returns a boolean if a field has been set.

### GetProvider

`func (o *CFPaymentsEntityMethod) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CFPaymentsEntityMethod) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CFPaymentsEntityMethod) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CFPaymentsEntityMethod) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPhone

`func (o *CFPaymentsEntityMethod) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CFPaymentsEntityMethod) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CFPaymentsEntityMethod) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CFPaymentsEntityMethod) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


