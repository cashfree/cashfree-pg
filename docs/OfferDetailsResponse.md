# OfferDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashbackDetails** | Pointer to [**CashbackDetails**](CashbackDetails.md) |  | [optional] 
**DiscountDetails** | Pointer to [**DiscountDetails**](DiscountDetails.md) |  | [optional] 
**OfferType** | Pointer to **string** | Offer Type for the Offer. | [optional] 

## Methods

### NewOfferDetailsResponse

`func NewOfferDetailsResponse() *OfferDetailsResponse`

NewOfferDetailsResponse instantiates a new OfferDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferDetailsResponseWithDefaults

`func NewOfferDetailsResponseWithDefaults() *OfferDetailsResponse`

NewOfferDetailsResponseWithDefaults instantiates a new OfferDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashbackDetails

`func (o *OfferDetailsResponse) GetCashbackDetails() CashbackDetails`

GetCashbackDetails returns the CashbackDetails field if non-nil, zero value otherwise.

### GetCashbackDetailsOk

`func (o *OfferDetailsResponse) GetCashbackDetailsOk() (*CashbackDetails, bool)`

GetCashbackDetailsOk returns a tuple with the CashbackDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackDetails

`func (o *OfferDetailsResponse) SetCashbackDetails(v CashbackDetails)`

SetCashbackDetails sets CashbackDetails field to given value.

### HasCashbackDetails

`func (o *OfferDetailsResponse) HasCashbackDetails() bool`

HasCashbackDetails returns a boolean if a field has been set.

### GetDiscountDetails

`func (o *OfferDetailsResponse) GetDiscountDetails() DiscountDetails`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *OfferDetailsResponse) GetDiscountDetailsOk() (*DiscountDetails, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *OfferDetailsResponse) SetDiscountDetails(v DiscountDetails)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *OfferDetailsResponse) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetOfferType

`func (o *OfferDetailsResponse) GetOfferType() string`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *OfferDetailsResponse) GetOfferTypeOk() (*string, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *OfferDetailsResponse) SetOfferType(v string)`

SetOfferType sets OfferType field to given value.

### HasOfferType

`func (o *OfferDetailsResponse) HasOfferType() bool`

HasOfferType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


