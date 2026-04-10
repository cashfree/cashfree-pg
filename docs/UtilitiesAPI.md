# \UtilitiesAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGCreatePAR**](UtilitiesAPI.md#PGCreatePAR) | **Post** /utilities/pars | Get PAR
[**PGGetCardBin**](UtilitiesAPI.md#PGGetCardBin) | **Post** /utilities/cardbin | Get Card BIN Details



## PGCreatePAR

> PGCreatePAR200Response PGCreatePAR(ctx).XApiVersion(xApiVersion).PARRequest(pARRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get PAR



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v5"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2026-01-01" 
    pARRequest := *cashfree.NewPARRequest("4111111111111111", "123", "12", "27") 
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD. | [default to &quot;2026-01-01&quot;]
 **pARRequest** | [**PARRequest**](PARRequest.md) | Request parameters for creating PAR. | 
 **xRequestId** | **string** | Request ID for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to Cashfree. | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.  | 

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


## PGGetCardBin

> GetCardBinResponseSchema PGGetCardBin(ctx).XApiVersion(xApiVersion).GetCardBinRequest(getCardBinRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Card BIN Details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v5"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2026-01-01" 
    getCardBinRequest := *cashfree.NewGetCardBinRequest("xxxx...xxx") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGGetCardBin(&xApiVersion, &getCardBinRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGGetCardBin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGGetCardBin`: GetCardBinResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGGetCardBin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGGetCardBinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD. | [default to &quot;2026-01-01&quot;]
 **getCardBinRequest** | [**GetCardBinRequest**](GetCardBinRequest.md) | Request parameters to get card BIN details. | 
 **xRequestId** | **string** | Request ID for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to Cashfree. | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.  | 

### Return type

[**GetCardBinResponseSchema**](GetCardBinResponseSchema.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

