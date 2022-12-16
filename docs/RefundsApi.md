# \RefundsApi

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createrefund**](RefundsApi.md#Createrefund) | **Post** /orders/{order_id}/refunds | Create Refund
[**GetRefund**](RefundsApi.md#GetRefund) | **Get** /orders/{order_id}/refunds/{refund_id} | Get Refund
[**Getallrefundsfororder**](RefundsApi.md#Getallrefundsfororder) | **Get** /orders/{order_id}/refunds | Get All Refunds for an Order



## Createrefund

> CFRefund Createrefund(ctx, orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFRefundRequest(cFRefundRequest).Execute()

Create Refund



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
    orderId := "orderId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)
    cFRefundRequest := *openapiclient.NewCFRefundRequest(float64(123), "RefundId_example") // CFRefundRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsApi.Createrefund(context.Background(), orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFRefundRequest(cFRefundRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.Createrefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Createrefund`: CFRefund
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.Createrefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreaterefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 
 **cFRefundRequest** | [**CFRefundRequest**](CFRefundRequest.md) |  | 

### Return type

[**CFRefund**](CFRefund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefund

> CFRefund GetRefund(ctx, orderId, refundId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Get Refund



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
    orderId := "orderId_example" // string | 
    refundId := "refundId_example" // string | Refund Id of the refund you want to fetch.
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsApi.GetRefund(context.Background(), orderId, refundId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.GetRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefund`: CFRefund
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.GetRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**refundId** | **string** | Refund Id of the refund you want to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 


 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFRefund**](CFRefund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getallrefundsfororder

> []CFRefund Getallrefundsfororder(ctx, orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).Execute()

Get All Refunds for an Order



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
    orderId := "orderId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RefundsApi.Getallrefundsfororder(context.Background(), orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundsApi.Getallrefundsfororder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getallrefundsfororder`: []CFRefund
    fmt.Fprintf(os.Stdout, "Response from `RefundsApi.Getallrefundsfororder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetallrefundsfororderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]

### Return type

[**[]CFRefund**](CFRefund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

