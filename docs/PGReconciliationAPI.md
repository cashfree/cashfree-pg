# \PGReconciliationAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGFetchRecon**](PGReconciliationAPI.md#PGFetchRecon) | **Post** /recon | PG Reconciliation



## PGFetchRecon

> ReconEntity PGFetchRecon(ctx).XApiVersion(xApiVersion).FetchReconRequest(fetchReconRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()

PG Reconciliation



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
	fetchReconRequest := *openapiclient.NewFetchReconRequest(*openapiclient.NewFetchReconRequestPagination(int32(123)), *openapiclient.NewFetchReconRequestFilters("StartDate_example", "EndDate_example")) // FetchReconRequest | Request Body for the reconciliation
	contentType := "application/json" // string | application/json (optional)
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)
	accept := "application/json" // string | application/json (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PGReconciliationAPI.PGFetchRecon(context.Background()).XApiVersion(xApiVersion).FetchReconRequest(fetchReconRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PGReconciliationAPI.PGFetchRecon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGFetchRecon`: ReconEntity
	fmt.Fprintf(os.Stdout, "Response from `PGReconciliationAPI.PGFetchRecon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchReconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **fetchReconRequest** | [**FetchReconRequest**](FetchReconRequest.md) | Request Body for the reconciliation | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 
 **accept** | **string** | application/json | 

### Return type

[**ReconEntity**](ReconEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

