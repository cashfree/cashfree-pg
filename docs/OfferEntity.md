# OfferEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | Pointer to **string** |  | [optional] 
**OfferStatus** | Pointer to **string** |  | [optional] 
**OrderAmount** | Pointer to **float32** |  | [optional] 
**PayableAmount** | Pointer to **float32** |  | [optional] 
**OfferMeta** | Pointer to [**OfferMetaResponse**](OfferMetaResponse.md) |  | [optional] 
**OfferTnc** | Pointer to [**OfferTncResponse**](OfferTncResponse.md) |  | [optional] 
**OfferDetails** | Pointer to [**OfferDetailsResponse**](OfferDetailsResponse.md) |  | [optional] 
**OfferValidations** | Pointer to [**OfferValidationsResponse**](OfferValidationsResponse.md) |  | [optional] 

## Methods

### NewOfferEntity

`func NewOfferEntity() *OfferEntity`

NewOfferEntity instantiates a new OfferEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferEntityWithDefaults

`func NewOfferEntityWithDefaults() *OfferEntity`

NewOfferEntityWithDefaults instantiates a new OfferEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *OfferEntity) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *OfferEntity) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *OfferEntity) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *OfferEntity) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetOfferStatus

`func (o *OfferEntity) GetOfferStatus() string`

GetOfferStatus returns the OfferStatus field if non-nil, zero value otherwise.

### GetOfferStatusOk

`func (o *OfferEntity) GetOfferStatusOk() (*string, bool)`

GetOfferStatusOk returns a tuple with the OfferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferStatus

`func (o *OfferEntity) SetOfferStatus(v string)`

SetOfferStatus sets OfferStatus field to given value.

### HasOfferStatus

`func (o *OfferEntity) HasOfferStatus() bool`

HasOfferStatus returns a boolean if a field has been set.

### GetOrderAmount

`func (o *OfferEntity) GetOrderAmount() float32`

GetOrderAmount returns the OrderAmount field if non-nil, zero value otherwise.

### GetOrderAmountOk

`func (o *OfferEntity) GetOrderAmountOk() (*float32, bool)`

GetOrderAmountOk returns a tuple with the OrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmount

`func (o *OfferEntity) SetOrderAmount(v float32)`

SetOrderAmount sets OrderAmount field to given value.

### HasOrderAmount

`func (o *OfferEntity) HasOrderAmount() bool`

HasOrderAmount returns a boolean if a field has been set.

### GetPayableAmount

`func (o *OfferEntity) GetPayableAmount() float32`

GetPayableAmount returns the PayableAmount field if non-nil, zero value otherwise.

### GetPayableAmountOk

`func (o *OfferEntity) GetPayableAmountOk() (*float32, bool)`

GetPayableAmountOk returns a tuple with the PayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableAmount

`func (o *OfferEntity) SetPayableAmount(v float32)`

SetPayableAmount sets PayableAmount field to given value.

### HasPayableAmount

`func (o *OfferEntity) HasPayableAmount() bool`

HasPayableAmount returns a boolean if a field has been set.

### GetOfferMeta

`func (o *OfferEntity) GetOfferMeta() OfferMetaResponse`

GetOfferMeta returns the OfferMeta field if non-nil, zero value otherwise.

### GetOfferMetaOk

`func (o *OfferEntity) GetOfferMetaOk() (*OfferMetaResponse, bool)`

GetOfferMetaOk returns a tuple with the OfferMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferMeta

`func (o *OfferEntity) SetOfferMeta(v OfferMetaResponse)`

SetOfferMeta sets OfferMeta field to given value.

### HasOfferMeta

`func (o *OfferEntity) HasOfferMeta() bool`

HasOfferMeta returns a boolean if a field has been set.

### GetOfferTnc

`func (o *OfferEntity) GetOfferTnc() OfferTncResponse`

GetOfferTnc returns the OfferTnc field if non-nil, zero value otherwise.

### GetOfferTncOk

`func (o *OfferEntity) GetOfferTncOk() (*OfferTncResponse, bool)`

GetOfferTncOk returns a tuple with the OfferTnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferTnc

`func (o *OfferEntity) SetOfferTnc(v OfferTncResponse)`

SetOfferTnc sets OfferTnc field to given value.

### HasOfferTnc

`func (o *OfferEntity) HasOfferTnc() bool`

HasOfferTnc returns a boolean if a field has been set.

### GetOfferDetails

`func (o *OfferEntity) GetOfferDetails() OfferDetailsResponse`

GetOfferDetails returns the OfferDetails field if non-nil, zero value otherwise.

### GetOfferDetailsOk

`func (o *OfferEntity) GetOfferDetailsOk() (*OfferDetailsResponse, bool)`

GetOfferDetailsOk returns a tuple with the OfferDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferDetails

`func (o *OfferEntity) SetOfferDetails(v OfferDetailsResponse)`

SetOfferDetails sets OfferDetails field to given value.

### HasOfferDetails

`func (o *OfferEntity) HasOfferDetails() bool`

HasOfferDetails returns a boolean if a field has been set.

### GetOfferValidations

`func (o *OfferEntity) GetOfferValidations() OfferValidationsResponse`

GetOfferValidations returns the OfferValidations field if non-nil, zero value otherwise.

### GetOfferValidationsOk

`func (o *OfferEntity) GetOfferValidationsOk() (*OfferValidationsResponse, bool)`

GetOfferValidationsOk returns a tuple with the OfferValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferValidations

`func (o *OfferEntity) SetOfferValidations(v OfferValidationsResponse)`

SetOfferValidations sets OfferValidations field to given value.

### HasOfferValidations

`func (o *OfferEntity) HasOfferValidations() bool`

HasOfferValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


