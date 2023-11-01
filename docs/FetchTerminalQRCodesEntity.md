# FetchTerminalQRCodesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to **string** | Name of the bank that is linked to the Static QR. | [optional] 
**QrCode** | Pointer to **string** | Base-64 Encoded QR Code URL | [optional] 
**QrCodeUrl** | Pointer to **string** | URL of the qr Code. | [optional] 
**Status** | Pointer to **string** | Status of the static QR. | [optional] 

## Methods

### NewFetchTerminalQRCodesEntity

`func NewFetchTerminalQRCodesEntity() *FetchTerminalQRCodesEntity`

NewFetchTerminalQRCodesEntity instantiates a new FetchTerminalQRCodesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchTerminalQRCodesEntityWithDefaults

`func NewFetchTerminalQRCodesEntityWithDefaults() *FetchTerminalQRCodesEntity`

NewFetchTerminalQRCodesEntityWithDefaults instantiates a new FetchTerminalQRCodesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *FetchTerminalQRCodesEntity) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *FetchTerminalQRCodesEntity) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *FetchTerminalQRCodesEntity) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *FetchTerminalQRCodesEntity) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetQrCode

`func (o *FetchTerminalQRCodesEntity) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *FetchTerminalQRCodesEntity) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *FetchTerminalQRCodesEntity) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *FetchTerminalQRCodesEntity) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetQrCodeUrl

`func (o *FetchTerminalQRCodesEntity) GetQrCodeUrl() string`

GetQrCodeUrl returns the QrCodeUrl field if non-nil, zero value otherwise.

### GetQrCodeUrlOk

`func (o *FetchTerminalQRCodesEntity) GetQrCodeUrlOk() (*string, bool)`

GetQrCodeUrlOk returns a tuple with the QrCodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeUrl

`func (o *FetchTerminalQRCodesEntity) SetQrCodeUrl(v string)`

SetQrCodeUrl sets QrCodeUrl field to given value.

### HasQrCodeUrl

`func (o *FetchTerminalQRCodesEntity) HasQrCodeUrl() bool`

HasQrCodeUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FetchTerminalQRCodesEntity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchTerminalQRCodesEntity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchTerminalQRCodesEntity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FetchTerminalQRCodesEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


