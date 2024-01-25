# DiscountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountType** | **string** | Type of discount | 
**DiscountValue** | **float32** | Value of Discount. | 
**MaxDiscountAmount** | **float32** | Maximum Value of Discount allowed. | 

## Methods

### NewDiscountDetails

`func NewDiscountDetails(discountType string, discountValue float32, maxDiscountAmount float32, ) *DiscountDetails`

NewDiscountDetails instantiates a new DiscountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountDetailsWithDefaults

`func NewDiscountDetailsWithDefaults() *DiscountDetails`

NewDiscountDetailsWithDefaults instantiates a new DiscountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountType

`func (o *DiscountDetails) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *DiscountDetails) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *DiscountDetails) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetDiscountValue

`func (o *DiscountDetails) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *DiscountDetails) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *DiscountDetails) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.


### GetMaxDiscountAmount

`func (o *DiscountDetails) GetMaxDiscountAmount() float32`

GetMaxDiscountAmount returns the MaxDiscountAmount field if non-nil, zero value otherwise.

### GetMaxDiscountAmountOk

`func (o *DiscountDetails) GetMaxDiscountAmountOk() (*float32, bool)`

GetMaxDiscountAmountOk returns a tuple with the MaxDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDiscountAmount

`func (o *DiscountDetails) SetMaxDiscountAmount(v float32)`

SetMaxDiscountAmount sets MaxDiscountAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


