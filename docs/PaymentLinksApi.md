# \PaymentLinksApi

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentLink**](PaymentLinksApi.md#CreatePaymentLink) | **Post** /links | Create Payment Link
[**GetPaymentLinkDetails**](PaymentLinksApi.md#GetPaymentLinkDetails) | **Get** /links/{link_id} | Fetch Payment Link Details
[**GetPaymentLinkOrders**](PaymentLinksApi.md#GetPaymentLinkOrders) | **Get** /links/{link_id}/orders | Get Orders for a Payment Link



## CreatePaymentLink

> CFLink CreatePaymentLink(ctx).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFLinkRequest(cFLinkRequest).Execute()

Create Payment Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xClientId := "xClientId_example" // string | 
    xClientSecret := "xClientSecret_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)
    cFLinkRequest := *openapiclient.NewCFLinkRequest("LinkId_example", float64(123), "LinkCurrency_example", "LinkPurpose_example", *openapiclient.NewCFLinkCustomerDetailsEntity("CustomerPhone_example")) // CFLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.CreatePaymentLink(context.Background()).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFLinkRequest(cFLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.CreatePaymentLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentLink`: CFLink
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.CreatePaymentLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 
 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 
 **cFLinkRequest** | [**CFLinkRequest**](CFLinkRequest.md) |  | 

### Return type

[**CFLink**](CFLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentLinkDetails

> CFLink GetPaymentLinkDetails(ctx, linkId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Fetch Payment Link Details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xClientId := "xClientId_example" // string | 
    xClientSecret := "xClientSecret_example" // string | 
    linkId := "linkId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.GetPaymentLinkDetails(context.Background(), linkId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.GetPaymentLinkDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentLinkDetails`: CFLink
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.GetPaymentLinkDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinkDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFLink**](CFLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentLinkOrders

> []CFLinkOrders GetPaymentLinkOrders(ctx, linkId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Get Orders for a Payment Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xClientId := "xClientId_example" // string | 
    xClientSecret := "xClientSecret_example" // string | 
    linkId := "linkId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentLinksApi.GetPaymentLinkOrders(context.Background(), linkId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.GetPaymentLinkOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentLinkOrders`: []CFLinkOrders
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.GetPaymentLinkOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinkOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**[]CFLinkOrders**](CFLinkOrders.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

