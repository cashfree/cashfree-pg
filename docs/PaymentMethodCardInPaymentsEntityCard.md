# PaymentMethodCardInPaymentsEntityCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The requested channel, can be &#x60;link&#x60; or &#x60;post&#x60;. | [optional] 
**CardNumber** | Pointer to **string** | The last four digits of the customer&#39;s card number. For external token transactions or external Alt ID transactions, this value is passed only when the merchant includes &#x60;card_display&#x60; in the [Order Pay API](https://www.cashfree.com/docs/api-reference/payments/latest/payments/pay) request. | [optional] 
**CardNetwork** | Pointer to **string** | The card scheme or network of the card. For example, &#x60;visa&#x60;, &#x60;mastercard&#x60;, &#x60;rupay&#x60;, &#x60;amex&#x60;, or &#x60;diners&#x60;. | [optional] 
**CardType** | Pointer to **string** | The type of card. For example, &#x60;credit_card&#x60;, &#x60;debit_card&#x60;, or &#x60;prepaid_card&#x60;. | [optional] 
**CardSubType** | Pointer to **string** | The sub-type of card. &#x60;R&#x60; is Retail card, &#x60;P&#x60; is Premium card, &#x60;C&#x60; is Corporate card. | [optional] 
**CardCountry** | Pointer to **string** | The issuing country of the card. For example, &#x60;IN&#x60;. | [optional] 
**CardBankName** | Pointer to **string** | The issuing bank of the card. For example, &#x60;HDFC BANK&#x60;, &#x60;AXIS BANK&#x60;, or &#x60;ICICI BANK&#x60;. | [optional] 
**CardNetworkReferenceId** | Pointer to **string** | The authentication reference ID provided by the respective card network. | [optional] 
**InstrumentId** | Pointer to **string** | The identifier for the card saved at Cashfree. This value is sent only for CF token transactions. | [optional] 

## Methods

### NewPaymentMethodCardInPaymentsEntityCard

`func NewPaymentMethodCardInPaymentsEntityCard() *PaymentMethodCardInPaymentsEntityCard`

NewPaymentMethodCardInPaymentsEntityCard instantiates a new PaymentMethodCardInPaymentsEntityCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardInPaymentsEntityCardWithDefaults

`func NewPaymentMethodCardInPaymentsEntityCardWithDefaults() *PaymentMethodCardInPaymentsEntityCard`

NewPaymentMethodCardInPaymentsEntityCardWithDefaults instantiates a new PaymentMethodCardInPaymentsEntityCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PaymentMethodCardInPaymentsEntityCard) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PaymentMethodCardInPaymentsEntityCard) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PaymentMethodCardInPaymentsEntityCard) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCardNumber

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCardNetwork

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.

### HasCardNetwork

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardNetwork() bool`

HasCardNetwork returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardSubType

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardSubType() string`

GetCardSubType returns the CardSubType field if non-nil, zero value otherwise.

### GetCardSubTypeOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardSubTypeOk() (*string, bool)`

GetCardSubTypeOk returns a tuple with the CardSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSubType

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardSubType(v string)`

SetCardSubType sets CardSubType field to given value.

### HasCardSubType

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardSubType() bool`

HasCardSubType returns a boolean if a field has been set.

### GetCardCountry

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardCountry() string`

GetCardCountry returns the CardCountry field if non-nil, zero value otherwise.

### GetCardCountryOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardCountryOk() (*string, bool)`

GetCardCountryOk returns a tuple with the CardCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCountry

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardCountry(v string)`

SetCardCountry sets CardCountry field to given value.

### HasCardCountry

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardCountry() bool`

HasCardCountry returns a boolean if a field has been set.

### GetCardBankName

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardBankName() string`

GetCardBankName returns the CardBankName field if non-nil, zero value otherwise.

### GetCardBankNameOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardBankNameOk() (*string, bool)`

GetCardBankNameOk returns a tuple with the CardBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBankName

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardBankName(v string)`

SetCardBankName sets CardBankName field to given value.

### HasCardBankName

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardBankName() bool`

HasCardBankName returns a boolean if a field has been set.

### GetCardNetworkReferenceId

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNetworkReferenceId() string`

GetCardNetworkReferenceId returns the CardNetworkReferenceId field if non-nil, zero value otherwise.

### GetCardNetworkReferenceIdOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetCardNetworkReferenceIdOk() (*string, bool)`

GetCardNetworkReferenceIdOk returns a tuple with the CardNetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetworkReferenceId

`func (o *PaymentMethodCardInPaymentsEntityCard) SetCardNetworkReferenceId(v string)`

SetCardNetworkReferenceId sets CardNetworkReferenceId field to given value.

### HasCardNetworkReferenceId

`func (o *PaymentMethodCardInPaymentsEntityCard) HasCardNetworkReferenceId() bool`

HasCardNetworkReferenceId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *PaymentMethodCardInPaymentsEntityCard) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *PaymentMethodCardInPaymentsEntityCard) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *PaymentMethodCardInPaymentsEntityCard) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *PaymentMethodCardInPaymentsEntityCard) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


