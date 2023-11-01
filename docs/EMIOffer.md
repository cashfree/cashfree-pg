# EMIOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of emi offer. Possible values are &#x60;credit_card_emi&#x60;, &#x60;debit_card_emi&#x60;, &#x60;cardless_emi&#x60; | 
**Issuer** | **string** | Bank Name | 
**Tenures** | **[]int32** |  | 

## Methods

### NewEMIOffer

`func NewEMIOffer(type_ string, issuer string, tenures []int32, ) *EMIOffer`

NewEMIOffer instantiates a new EMIOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEMIOfferWithDefaults

`func NewEMIOfferWithDefaults() *EMIOffer`

NewEMIOfferWithDefaults instantiates a new EMIOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EMIOffer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EMIOffer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EMIOffer) SetType(v string)`

SetType sets Type field to given value.


### GetIssuer

`func (o *EMIOffer) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *EMIOffer) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *EMIOffer) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetTenures

`func (o *EMIOffer) GetTenures() []int32`

GetTenures returns the Tenures field if non-nil, zero value otherwise.

### GetTenuresOk

`func (o *EMIOffer) GetTenuresOk() (*[]int32, bool)`

GetTenuresOk returns a tuple with the Tenures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenures

`func (o *EMIOffer) SetTenures(v []int32)`

SetTenures sets Tenures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


