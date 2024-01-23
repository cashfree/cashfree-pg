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
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2022-09-01" 
    fetchSettlementsRequest := *cashfree.NewFetchSettlementsRequest(*cashfree.NewFetchSettlementsRequestPagination(int32(123)), *cashfree.NewFetchSettlementsRequestFilters()) 
    contentType := "application/json" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 
    accept := "application/json" 

    resp, r, err := cashfree.PGFetchSettlements(&xApiVersion, &fetchSettlementsRequest, &contentType, &xRequestId, &xIdempotencyKey, &accept, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGFetchSettlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGFetchSettlements`: SettlementEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGFetchSettlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchSettlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
 **fetchSettlementsRequest** | [**FetchSettlementsRequest**](FetchSettlementsRequest.md) | Request Body to get the settlements | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions. | 
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
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2022-09-01" 
    settlementFetchReconRequest := *cashfree.NewSettlementFetchReconRequest(*cashfree.NewFetchSettlementsRequestPagination(int32(123)), *cashfree.NewSettlementFetchReconRequestFilters()) 
    contentType := "application/json" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 
    accept := "application/json" 

    resp, r, err := cashfree.PGSettlementFetchRecon(&xApiVersion, &settlementFetchReconRequest, &contentType, &xRequestId, &xIdempotencyKey, &accept, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGSettlementFetchRecon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGSettlementFetchRecon`: SettlementReconEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGSettlementFetchRecon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGSettlementFetchReconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
 **settlementFetchReconRequest** | [**SettlementFetchReconRequest**](SettlementFetchReconRequest.md) | Request Body for the settlement reconciliation | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions. | 
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

