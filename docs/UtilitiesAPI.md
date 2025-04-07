# \UtilitiesAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGCreatePAR**](UtilitiesAPI.md#PGCreatePAR) | **Post** /pg/utilities/pars | Create PAR at Cashfree



## PGCreatePAR

> PGCreatePAR200Response PGCreatePAR(ctx).XApiVersion(xApiVersion).PARRequest(pARRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create PAR at Cashfree



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v4"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2025-01-01" 
    pARRequest := *cashfree.NewPARRequest("xxxx...xxx", "xxx", "xx", "xx", "PLAIN_CARD") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCreatePAR(&xApiVersion, &pARRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCreatePAR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCreatePAR`: PGCreatePAR200Response
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCreatePAR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGCreatePARRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **pARRequest** | [**PARRequest**](PARRequest.md) | Request payload for creating PAR. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PGCreatePAR200Response**](PGCreatePAR200Response.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

