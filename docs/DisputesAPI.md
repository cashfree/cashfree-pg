# \DisputesAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGAcceptDisputeByID**](DisputesAPI.md#PGAcceptDisputeByID) | **Put** /disputes/{dispute_id}/accept | Accept Dispute by Dispute ID
[**PGFetchDisputeByID**](DisputesAPI.md#PGFetchDisputeByID) | **Get** /disputes/{dispute_id} | Get Disputes by Dispute ID
[**PGFetchOrderDisputes**](DisputesAPI.md#PGFetchOrderDisputes) | **Get** /orders/{order_id}/disputes | Get Disputes by Order Id
[**PGFetchPaymentDisputes**](DisputesAPI.md#PGFetchPaymentDisputes) | **Get** /payments/{cf_payment_id}/disputes | Get Disputes by Payment ID
[**PGUploadDisputesDocuments**](DisputesAPI.md#PGUploadDisputesDocuments) | **Post** /disputes/{dispute_id}/documents | Submit Evidence to contest the Dispute by Dispute ID



## PGAcceptDisputeByID

> DisputesEntityMerchantAccepted PGAcceptDisputeByID(ctx, disputeId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Accept Dispute by Dispute ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	disputeId := int32(56) // int32 | 
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.PGAcceptDisputeByID(context.Background(), disputeId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.PGAcceptDisputeByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGAcceptDisputeByID`: DisputesEntityMerchantAccepted
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.PGAcceptDisputeByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGAcceptDisputeByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**DisputesEntityMerchantAccepted**](DisputesEntityMerchantAccepted.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGFetchDisputeByID

> DisputesEntity PGFetchDisputeByID(ctx, disputeId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Disputes by Dispute ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	disputeId := int32(56) // int32 | 
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.PGFetchDisputeByID(context.Background(), disputeId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.PGFetchDisputeByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGFetchDisputeByID`: DisputesEntity
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.PGFetchDisputeByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchDisputeByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**DisputesEntity**](DisputesEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGFetchOrderDisputes

> []DisputesEntity PGFetchOrderDisputes(ctx, orderId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Disputes by Order Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	orderId := "orderId_example" // string | 
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.PGFetchOrderDisputes(context.Background(), orderId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.PGFetchOrderDisputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGFetchOrderDisputes`: []DisputesEntity
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.PGFetchOrderDisputes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchOrderDisputesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]DisputesEntity**](DisputesEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGFetchPaymentDisputes

> []DisputesEntity PGFetchPaymentDisputes(ctx, cfPaymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Disputes by Payment ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	cfPaymentId := int32(56) // int32 | 
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.PGFetchPaymentDisputes(context.Background(), cfPaymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.PGFetchPaymentDisputes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGFetchPaymentDisputes`: []DisputesEntity
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.PGFetchPaymentDisputes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfPaymentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchPaymentDisputesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]DisputesEntity**](DisputesEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGUploadDisputesDocuments

> []DisputesEntity PGUploadDisputesDocuments(ctx, disputeId).XApiVersion(xApiVersion).File(file).DocType(docType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Note(note).Execute()

Submit Evidence to contest the Dispute by Dispute ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	disputeId := int32(56) // int32 | 
	file := "file_example" // string | File types supported are jpeg, jpg, png, pdf and maximum file size allowed is 20 MB.
	docType := "docType_example" // string | Mention the type of the document you are uploading. Possible values :- Delivery/Service Proof, Shipping Proof, Statement of Service, Proof of Service Used, Cancellation of Service Proof, Refund Proof, Business model explanation, Extra Charges Declaration, Terms & Conditions, Customer Withdrawal Letter, Certificate of Authenticity, Reseller Agreement. You can use get evidences to contest dispute API to fetch set of documents required to contest particular dispute.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)
	note := "note_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputesAPI.PGUploadDisputesDocuments(context.Background(), disputeId).XApiVersion(xApiVersion).File(file).DocType(docType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputesAPI.PGUploadDisputesDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGUploadDisputesDocuments`: []DisputesEntity
	fmt.Fprintf(os.Stdout, "Response from `DisputesAPI.PGUploadDisputesDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disputeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGUploadDisputesDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **file** | **string** | File types supported are jpeg, jpg, png, pdf and maximum file size allowed is 20 MB. | 
 **docType** | **string** | Mention the type of the document you are uploading. Possible values :- Delivery/Service Proof, Shipping Proof, Statement of Service, Proof of Service Used, Cancellation of Service Proof, Refund Proof, Business model explanation, Extra Charges Declaration, Terms &amp; Conditions, Customer Withdrawal Letter, Certificate of Authenticity, Reseller Agreement. You can use get evidences to contest dispute API to fetch set of documents required to contest particular dispute. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 
 **note** | **string** |  | 

### Return type

[**[]DisputesEntity**](DisputesEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

