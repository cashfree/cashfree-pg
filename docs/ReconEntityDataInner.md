# ReconEntityDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDetails** | Pointer to [**ReconEntityDataInnerEventDetails**](ReconEntityDataInnerEventDetails.md) |  | [optional] 
**OrderDetails** | Pointer to [**ReconEntityDataInnerOrderDetails**](ReconEntityDataInnerOrderDetails.md) |  | [optional] 
**CustomerDetails** | Pointer to [**ReconEntityDataInnerCustomerDetails**](ReconEntityDataInnerCustomerDetails.md) |  | [optional] 
**PaymentDetails** | Pointer to [**ReconEntityDataInnerPaymentDetails**](ReconEntityDataInnerPaymentDetails.md) |  | [optional] 
**SettlementDetails** | Pointer to [**ReconEntityDataInnerSettlementDetails**](ReconEntityDataInnerSettlementDetails.md) |  | [optional] 
**DisputeDetails** | Pointer to [**ReconEntityDataInnerDisputeDetails**](ReconEntityDataInnerDisputeDetails.md) |  | [optional] 
**RefundDetails** | Pointer to [**ReconEntityDataInnerRefundDetails**](ReconEntityDataInnerRefundDetails.md) |  | [optional] 

## Methods

### NewReconEntityDataInner

`func NewReconEntityDataInner() *ReconEntityDataInner`

NewReconEntityDataInner instantiates a new ReconEntityDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconEntityDataInnerWithDefaults

`func NewReconEntityDataInnerWithDefaults() *ReconEntityDataInner`

NewReconEntityDataInnerWithDefaults instantiates a new ReconEntityDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDetails

`func (o *ReconEntityDataInner) GetEventDetails() ReconEntityDataInnerEventDetails`

GetEventDetails returns the EventDetails field if non-nil, zero value otherwise.

### GetEventDetailsOk

`func (o *ReconEntityDataInner) GetEventDetailsOk() (*ReconEntityDataInnerEventDetails, bool)`

GetEventDetailsOk returns a tuple with the EventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetails

`func (o *ReconEntityDataInner) SetEventDetails(v ReconEntityDataInnerEventDetails)`

SetEventDetails sets EventDetails field to given value.

### HasEventDetails

`func (o *ReconEntityDataInner) HasEventDetails() bool`

HasEventDetails returns a boolean if a field has been set.

### GetOrderDetails

`func (o *ReconEntityDataInner) GetOrderDetails() ReconEntityDataInnerOrderDetails`

GetOrderDetails returns the OrderDetails field if non-nil, zero value otherwise.

### GetOrderDetailsOk

`func (o *ReconEntityDataInner) GetOrderDetailsOk() (*ReconEntityDataInnerOrderDetails, bool)`

GetOrderDetailsOk returns a tuple with the OrderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetails

`func (o *ReconEntityDataInner) SetOrderDetails(v ReconEntityDataInnerOrderDetails)`

SetOrderDetails sets OrderDetails field to given value.

### HasOrderDetails

`func (o *ReconEntityDataInner) HasOrderDetails() bool`

HasOrderDetails returns a boolean if a field has been set.

### GetCustomerDetails

`func (o *ReconEntityDataInner) GetCustomerDetails() ReconEntityDataInnerCustomerDetails`

GetCustomerDetails returns the CustomerDetails field if non-nil, zero value otherwise.

### GetCustomerDetailsOk

`func (o *ReconEntityDataInner) GetCustomerDetailsOk() (*ReconEntityDataInnerCustomerDetails, bool)`

GetCustomerDetailsOk returns a tuple with the CustomerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDetails

`func (o *ReconEntityDataInner) SetCustomerDetails(v ReconEntityDataInnerCustomerDetails)`

SetCustomerDetails sets CustomerDetails field to given value.

### HasCustomerDetails

`func (o *ReconEntityDataInner) HasCustomerDetails() bool`

HasCustomerDetails returns a boolean if a field has been set.

### GetPaymentDetails

`func (o *ReconEntityDataInner) GetPaymentDetails() ReconEntityDataInnerPaymentDetails`

GetPaymentDetails returns the PaymentDetails field if non-nil, zero value otherwise.

### GetPaymentDetailsOk

`func (o *ReconEntityDataInner) GetPaymentDetailsOk() (*ReconEntityDataInnerPaymentDetails, bool)`

GetPaymentDetailsOk returns a tuple with the PaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetails

`func (o *ReconEntityDataInner) SetPaymentDetails(v ReconEntityDataInnerPaymentDetails)`

SetPaymentDetails sets PaymentDetails field to given value.

### HasPaymentDetails

`func (o *ReconEntityDataInner) HasPaymentDetails() bool`

HasPaymentDetails returns a boolean if a field has been set.

### GetSettlementDetails

`func (o *ReconEntityDataInner) GetSettlementDetails() ReconEntityDataInnerSettlementDetails`

GetSettlementDetails returns the SettlementDetails field if non-nil, zero value otherwise.

### GetSettlementDetailsOk

`func (o *ReconEntityDataInner) GetSettlementDetailsOk() (*ReconEntityDataInnerSettlementDetails, bool)`

GetSettlementDetailsOk returns a tuple with the SettlementDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDetails

`func (o *ReconEntityDataInner) SetSettlementDetails(v ReconEntityDataInnerSettlementDetails)`

SetSettlementDetails sets SettlementDetails field to given value.

### HasSettlementDetails

`func (o *ReconEntityDataInner) HasSettlementDetails() bool`

HasSettlementDetails returns a boolean if a field has been set.

### GetDisputeDetails

`func (o *ReconEntityDataInner) GetDisputeDetails() ReconEntityDataInnerDisputeDetails`

GetDisputeDetails returns the DisputeDetails field if non-nil, zero value otherwise.

### GetDisputeDetailsOk

`func (o *ReconEntityDataInner) GetDisputeDetailsOk() (*ReconEntityDataInnerDisputeDetails, bool)`

GetDisputeDetailsOk returns a tuple with the DisputeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDetails

`func (o *ReconEntityDataInner) SetDisputeDetails(v ReconEntityDataInnerDisputeDetails)`

SetDisputeDetails sets DisputeDetails field to given value.

### HasDisputeDetails

`func (o *ReconEntityDataInner) HasDisputeDetails() bool`

HasDisputeDetails returns a boolean if a field has been set.

### GetRefundDetails

`func (o *ReconEntityDataInner) GetRefundDetails() ReconEntityDataInnerRefundDetails`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *ReconEntityDataInner) GetRefundDetailsOk() (*ReconEntityDataInnerRefundDetails, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *ReconEntityDataInner) SetRefundDetails(v ReconEntityDataInnerRefundDetails)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *ReconEntityDataInner) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


