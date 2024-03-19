# StaticSplitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Specify if the split is to be active or not. Possible values: true/false | 
**TerminalId** | Pointer to **string** | For Subscription payments, the subscription reference ID is to be shared as the terminal ID. Incase for Payment Gateway terminal ID is non-mandatory. Mention as 0 if not applicable. | [optional] 
**TerminalReferenceId** | Pointer to **float32** | You can share additional information using the reference ID. | [optional] 
**ProductType** | **string** | Specify the product for which the split should be created. If you want split to be created for Payment Gateway pass value as \&quot;PG\&quot;. If you want split to be created for Subscription, pass value as \&quot;SBC\&quot;. Accepted values - \&quot;STATIC_QR\&quot;, \&quot;SBC\&quot;, \&quot;PG\&quot;, \&quot;EPOS\&quot;. | 
**Scheme** | [**[]StaticSplitRequestSchemeInner**](StaticSplitRequestSchemeInner.md) | Provide the split scheme details. | 

## Methods

### NewStaticSplitRequest

`func NewStaticSplitRequest(active bool, productType string, scheme []StaticSplitRequestSchemeInner, ) *StaticSplitRequest`

NewStaticSplitRequest instantiates a new StaticSplitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSplitRequestWithDefaults

`func NewStaticSplitRequestWithDefaults() *StaticSplitRequest`

NewStaticSplitRequestWithDefaults instantiates a new StaticSplitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *StaticSplitRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StaticSplitRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StaticSplitRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetTerminalId

`func (o *StaticSplitRequest) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *StaticSplitRequest) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *StaticSplitRequest) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *StaticSplitRequest) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalReferenceId

`func (o *StaticSplitRequest) GetTerminalReferenceId() float32`

GetTerminalReferenceId returns the TerminalReferenceId field if non-nil, zero value otherwise.

### GetTerminalReferenceIdOk

`func (o *StaticSplitRequest) GetTerminalReferenceIdOk() (*float32, bool)`

GetTerminalReferenceIdOk returns a tuple with the TerminalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalReferenceId

`func (o *StaticSplitRequest) SetTerminalReferenceId(v float32)`

SetTerminalReferenceId sets TerminalReferenceId field to given value.

### HasTerminalReferenceId

`func (o *StaticSplitRequest) HasTerminalReferenceId() bool`

HasTerminalReferenceId returns a boolean if a field has been set.

### GetProductType

`func (o *StaticSplitRequest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *StaticSplitRequest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *StaticSplitRequest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetScheme

`func (o *StaticSplitRequest) GetScheme() []StaticSplitRequestSchemeInner`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *StaticSplitRequest) GetSchemeOk() (*[]StaticSplitRequestSchemeInner, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *StaticSplitRequest) SetScheme(v []StaticSplitRequestSchemeInner)`

SetScheme sets Scheme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


