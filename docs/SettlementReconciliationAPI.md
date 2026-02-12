# \SettlementReconciliationAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGFetchSettlements**](SettlementReconciliationAPI.md#PGFetchSettlements) | **Post** /settlements | Get All Settlements
[**PGSettlementFetchRecon**](SettlementReconciliationAPI.md#PGSettlementFetchRecon) | **Post** /settlement/recon | Settlement Reconciliation



## PGFetchSettlements

> SettlementEntity PGFetchSettlements(ctx).XApiVersion(xApiVersion).FetchSettlementsRequest(fetchSettlementsRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()

Get All Settlements



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
	fetchSettlementsRequest := *openapiclient.NewFetchSettlementsRequest(*openapiclient.NewFetchSettlementsRequestPagination(int32(123)), *openapiclient.NewFetchSettlementsRequestFilters()) // FetchSettlementsRequest | Request Body to get the settlements
	contentType := "application/json" // string | application/json (optional)
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)
	accept := "application/json" // string | application/json (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementReconciliationAPI.PGFetchSettlements(context.Background()).XApiVersion(xApiVersion).FetchSettlementsRequest(fetchSettlementsRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementReconciliationAPI.PGFetchSettlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGFetchSettlements`: SettlementEntity
	fmt.Fprintf(os.Stdout, "Response from `SettlementReconciliationAPI.PGFetchSettlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchSettlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **fetchSettlementsRequest** | [**FetchSettlementsRequest**](FetchSettlementsRequest.md) | Request Body to get the settlements | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 
 **accept** | **string** | application/json | 

### Return type

[**SettlementEntity**](SettlementEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGSettlementFetchRecon

> SettlementReconEntity PGSettlementFetchRecon(ctx).XApiVersion(xApiVersion).SettlementFetchReconRequest(settlementFetchReconRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()

Settlement Reconciliation



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
	settlementFetchReconRequest := *openapiclient.NewSettlementFetchReconRequest(*openapiclient.NewFetchSettlementsRequestPagination(int32(123)), *openapiclient.NewFetchSettlementsRequestFilters()) // SettlementFetchReconRequest | Request Body for the settlement reconciliation
	contentType := "application/json" // string | application/json (optional)
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)
	accept := "application/json" // string | application/json (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementReconciliationAPI.PGSettlementFetchRecon(context.Background()).XApiVersion(xApiVersion).SettlementFetchReconRequest(settlementFetchReconRequest).ContentType(contentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementReconciliationAPI.PGSettlementFetchRecon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGSettlementFetchRecon`: SettlementReconEntity
	fmt.Fprintf(os.Stdout, "Response from `SettlementReconciliationAPI.PGSettlementFetchRecon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGSettlementFetchReconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **settlementFetchReconRequest** | [**SettlementFetchReconRequest**](SettlementFetchReconRequest.md) | Request Body for the settlement reconciliation | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 
 **accept** | **string** | application/json | 

### Return type

[**SettlementReconEntity**](SettlementReconEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

