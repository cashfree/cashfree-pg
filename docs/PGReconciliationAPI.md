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
    cashfree "github.com/cashfree/cashfree_pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2022-09-01" 
    fetchReconRequest := *cashfree.NewFetchReconRequest(*cashfree.NewFetchReconRequestPagination(int32(123)), *cashfree.NewFetchReconRequestFilters("StartDate_example", "EndDate_example")) 
    contentType := "application/json" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 
    accept := "application/json" 

    resp, r, err := cashfree.PGFetchRecon(&xApiVersion, &fetchReconRequest, &contentType, &xRequestId, &xIdempotencyKey, &accept, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGFetchRecon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGFetchRecon`: ReconEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGFetchRecon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchReconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
 **fetchReconRequest** | [**FetchReconRequest**](FetchReconRequest.md) | Request Body for the reconciliation | 
 **contentType** | **string** | application/json | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | Idempotency works by saving the resulting status code and body of the first request made for any given idempotency key, regardless of whether it succeeded or failed. Subsequent requests with the same key return the same result, including 500 errors.  Currently supported on all POST calls that uses x-client-id &amp; x-client-secret. To use enable, pass x-idempotency-key in the request header. The value of this header must be unique to each operation you are trying to do. One example can be to use the same order_id that you pass while creating orders   | 
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

