# SettlementReconEntityDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDetails** | Pointer to [**ReconEntityDataInnerEventDetails**](ReconEntityDataInnerEventDetails.md) |  | [optional] 
**OrderDetails** | Pointer to [**ReconEntityDataInnerOrderDetails**](ReconEntityDataInnerOrderDetails.md) |  | [optional] 
**CustomerDetails** | Pointer to [**ReconEntityDataInnerCustomerDetails**](ReconEntityDataInnerCustomerDetails.md) |  | [optional] 
**PaymentDetails** | Pointer to [**SettlementReconEntityDataInnerPaymentDetails**](SettlementReconEntityDataInnerPaymentDetails.md) |  | [optional] 
**SettlementDetails** | Pointer to [**ReconEntityDataInnerSettlementDetails**](ReconEntityDataInnerSettlementDetails.md) |  | [optional] 
**DisputeDetails** | Pointer to [**ReconEntityDataInnerDisputeDetails**](ReconEntityDataInnerDisputeDetails.md) |  | [optional] 
**RefundDetails** | Pointer to [**ReconEntityDataInnerRefundDetails**](ReconEntityDataInnerRefundDetails.md) |  | [optional] 

## Methods

### NewSettlementReconEntityDataInner

`func NewSettlementReconEntityDataInner() *SettlementReconEntityDataInner`

NewSettlementReconEntityDataInner instantiates a new SettlementReconEntityDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementReconEntityDataInnerWithDefaults

`func NewSettlementReconEntityDataInnerWithDefaults() *SettlementReconEntityDataInner`

NewSettlementReconEntityDataInnerWithDefaults instantiates a new SettlementReconEntityDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDetails

`func (o *SettlementReconEntityDataInner) GetEventDetails() ReconEntityDataInnerEventDetails`

GetEventDetails returns the EventDetails field if non-nil, zero value otherwise.

### GetEventDetailsOk

`func (o *SettlementReconEntityDataInner) GetEventDetailsOk() (*ReconEntityDataInnerEventDetails, bool)`

GetEventDetailsOk returns a tuple with the EventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetails

`func (o *SettlementReconEntityDataInner) SetEventDetails(v ReconEntityDataInnerEventDetails)`

SetEventDetails sets EventDetails field to given value.

### HasEventDetails

`func (o *SettlementReconEntityDataInner) HasEventDetails() bool`

HasEventDetails returns a boolean if a field has been set.

### GetOrderDetails

`func (o *SettlementReconEntityDataInner) GetOrderDetails() ReconEntityDataInnerOrderDetails`

GetOrderDetails returns the OrderDetails field if non-nil, zero value otherwise.

### GetOrderDetailsOk

`func (o *SettlementReconEntityDataInner) GetOrderDetailsOk() (*ReconEntityDataInnerOrderDetails, bool)`

GetOrderDetailsOk returns a tuple with the OrderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetails

`func (o *SettlementReconEntityDataInner) SetOrderDetails(v ReconEntityDataInnerOrderDetails)`

SetOrderDetails sets OrderDetails field to given value.

### HasOrderDetails

`func (o *SettlementReconEntityDataInner) HasOrderDetails() bool`

HasOrderDetails returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *SettlementReconEntityDataInner) GetCustomerDetails() ReconEntityDataInnerCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *SettlementReconEntityDataInner) GetCustomerDetailsOk() (*ReconEntityDataInnerCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *SettlementReconEntityDataInner) SetCustomerDetails(v ReconEntityDataInnerCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *SettlementReconEntityDataInner) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetPaymentDetails

`func (o *SettlementReconEntityDataInner) GetPaymentDetails() SettlementReconEntityDataInnerPaymentDetails`

GetPaymentDetails returns the PaymentDetails field if non-nil, zero value otherwise.

### GetPaymentDetailsOk

`func (o *SettlementReconEntityDataInner) GetPaymentDetailsOk() (*SettlementReconEntityDataInnerPaymentDetails, bool)`

GetPaymentDetailsOk returns a tuple with the PaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetails

`func (o *SettlementReconEntityDataInner) SetPaymentDetails(v SettlementReconEntityDataInnerPaymentDetails)`

SetPaymentDetails sets PaymentDetails field to given value.

### HasPaymentDetails

`func (o *SettlementReconEntityDataInner) HasPaymentDetails() bool`

HasPaymentDetails returns a boolean if a field has been set.

### GetSettlementDetails

`func (o *SettlementReconEntityDataInner) GetSettlementDetails() ReconEntityDataInnerSettlementDetails`

GetSettlementDetails returns the SettlementDetails field if non-nil, zero value otherwise.

### GetSettlementDetailsOk

`func (o *SettlementReconEntityDataInner) GetSettlementDetailsOk() (*ReconEntityDataInnerSettlementDetails, bool)`

GetSettlementDetailsOk returns a tuple with the SettlementDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDetails

`func (o *SettlementReconEntityDataInner) SetSettlementDetails(v ReconEntityDataInnerSettlementDetails)`

SetSettlementDetails sets SettlementDetails field to given value.

### HasSettlementDetails

`func (o *SettlementReconEntityDataInner) HasSettlementDetails() bool`

HasSettlementDetails returns a boolean if a field has been set.

### GetDisputeDetails

`func (o *SettlementReconEntityDataInner) GetDisputeDetails() ReconEntityDataInnerDisputeDetails`

GetDisputeDetails returns the DisputeDetails field if non-nil, zero value otherwise.

### GetDisputeDetailsOk

`func (o *SettlementReconEntityDataInner) GetDisputeDetailsOk() (*ReconEntityDataInnerDisputeDetails, bool)`

GetDisputeDetailsOk returns a tuple with the DisputeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDetails

`func (o *SettlementReconEntityDataInner) SetDisputeDetails(v ReconEntityDataInnerDisputeDetails)`

SetDisputeDetails sets DisputeDetails field to given value.

### HasDisputeDetails

`func (o *SettlementReconEntityDataInner) HasDisputeDetails() bool`

HasDisputeDetails returns a boolean if a field has been set.

### GetRefundDetails

`func (o *SettlementReconEntityDataInner) GetRefundDetails() ReconEntityDataInnerRefundDetails`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *SettlementReconEntityDataInner) GetRefundDetailsOk() (*ReconEntityDataInnerRefundDetails, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *SettlementReconEntityDataInner) SetRefundDetails(v ReconEntityDataInnerRefundDetails)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *SettlementReconEntityDataInner) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


