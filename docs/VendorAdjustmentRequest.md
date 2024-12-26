# VendorAdjustmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | **string** | The unique identifier of the vendor to whom the adjustment is applied | 
**AdjustmentId** | **int64** | The unique identifier for the adjustment transaction. | 
**Amount** | **float32** | The adjustment amount to be applied. | 
**Type** | **string** | The type of adjustment. Possible values: CREDIT, DEBIT. | 
**Remarks** | Pointer to **string** | Remarks for the adjustment transaction, if any. | [optional] 

## Methods

### NewVendorAdjustmentRequest

`func NewVendorAdjustmentRequest(vendorId string, adjustmentId int64, amount float32, type_ string, ) *VendorAdjustmentRequest`

NewVendorAdjustmentRequest instantiates a new VendorAdjustmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorAdjustmentRequestWithDefaults

`func NewVendorAdjustmentRequestWithDefaults() *VendorAdjustmentRequest`

NewVendorAdjustmentRequestWithDefaults instantiates a new VendorAdjustmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *VendorAdjustmentRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *VendorAdjustmentRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *VendorAdjustmentRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.


### GetAdjustmentId

`func (o *VendorAdjustmentRequest) GetAdjustmentId() int64`

GetAdjustmentId returns the AdjustmentId field if non-nil, zero value otherwise.

### GetAdjustmentIdOk

`func (o *VendorAdjustmentRequest) GetAdjustmentIdOk() (*int64, bool)`

GetAdjustmentIdOk returns a tuple with the AdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentId

`func (o *VendorAdjustmentRequest) SetAdjustmentId(v int64)`

SetAdjustmentId sets AdjustmentId field to given value.


### GetAmount

`func (o *VendorAdjustmentRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VendorAdjustmentRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VendorAdjustmentRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetType

`func (o *VendorAdjustmentRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorAdjustmentRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorAdjustmentRequest) SetType(v string)`

SetType sets Type field to given value.


### GetRemarks

`func (o *VendorAdjustmentRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *VendorAdjustmentRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *VendorAdjustmentRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *VendorAdjustmentRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


