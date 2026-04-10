# VendorRecon200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]VendorRecon200ResponseDataInner**](VendorRecon200ResponseDataInner.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewVendorRecon200Response

`func NewVendorRecon200Response() *VendorRecon200Response`

NewVendorRecon200Response instantiates a new VendorRecon200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorRecon200ResponseWithDefaults

`func NewVendorRecon200ResponseWithDefaults() *VendorRecon200Response`

NewVendorRecon200ResponseWithDefaults instantiates a new VendorRecon200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *VendorRecon200Response) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *VendorRecon200Response) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *VendorRecon200Response) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *VendorRecon200Response) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetData

`func (o *VendorRecon200Response) GetData() []VendorRecon200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VendorRecon200Response) GetDataOk() (*[]VendorRecon200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VendorRecon200Response) SetData(v []VendorRecon200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *VendorRecon200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *VendorRecon200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VendorRecon200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VendorRecon200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VendorRecon200Response) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


