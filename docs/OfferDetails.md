# OfferDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferType** | **string** | Offer Type for the Offer. | 
**DiscountDetails** | Pointer to [**DiscountDetails**](DiscountDetails.md) |  | [optional] 
**CashbackDetails** | Pointer to [**CashbackDetails**](CashbackDetails.md) |  | [optional] 

## Methods

### NewOfferDetails

`func NewOfferDetails(offerType string, ) *OfferDetails`

NewOfferDetails instantiates a new OfferDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferDetailsWithDefaults

`func NewOfferDetailsWithDefaults() *OfferDetails`

NewOfferDetailsWithDefaults instantiates a new OfferDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferType

`func (o *OfferDetails) GetOfferType() string`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *OfferDetails) GetOfferTypeOk() (*string, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *OfferDetails) SetOfferType(v string)`

SetOfferType sets OfferType field to given value.


### GetDiscountDetails

`func (o *OfferDetails) GetDiscountDetails() DiscountDetails`

GetDiscountDetails returns the DiscountDetails field if non-nil, zero value otherwise.

### GetDiscountDetailsOk

`func (o *OfferDetails) GetDiscountDetailsOk() (*DiscountDetails, bool)`

GetDiscountDetailsOk returns a tuple with the DiscountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDetails

`func (o *OfferDetails) SetDiscountDetails(v DiscountDetails)`

SetDiscountDetails sets DiscountDetails field to given value.

### HasDiscountDetails

`func (o *OfferDetails) HasDiscountDetails() bool`

HasDiscountDetails returns a boolean if a field has been set.

### GetCashbackDetails

`func (o *OfferDetails) GetCashbackDetails() CashbackDetails`

GetCashbackDetails returns the CashbackDetails field if non-nil, zero value otherwise.

### GetCashbackDetailsOk

`func (o *OfferDetails) GetCashbackDetailsOk() (*CashbackDetails, bool)`

GetCashbackDetailsOk returns a tuple with the CashbackDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackDetails

`func (o *OfferDetails) SetCashbackDetails(v CashbackDetails)`

SetCashbackDetails sets CashbackDetails field to given value.

### HasCashbackDetails

`func (o *OfferDetails) HasCashbackDetails() bool`

HasCashbackDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


