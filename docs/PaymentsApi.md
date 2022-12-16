# \PaymentsApi

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentbyId**](PaymentsApi.md#GetPaymentbyId) | **Get** /orders/{order_id}/payments/{cf_payment_id} | Get Payment by ID
[**GetPaymentsfororder**](PaymentsApi.md#GetPaymentsfororder) | **Get** /orders/{order_id}/payments | Get Payments for an Order



## GetPaymentbyId

> CFPaymentsEntity GetPaymentbyId(ctx, orderId, cfPaymentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Get Payment by ID



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
    cfPaymentId := int32(56) // int32 | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.GetPaymentbyId(context.Background(), orderId, cfPaymentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.GetPaymentbyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentbyId`: CFPaymentsEntity
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.GetPaymentbyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**cfPaymentId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentbyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 


 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFPaymentsEntity**](CFPaymentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsfororder

> []CFPaymentsEntity GetPaymentsfororder(ctx, orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Get Payments for an Order



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.GetPaymentsfororder(context.Background(), orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.GetPaymentsfororder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsfororder`: []CFPaymentsEntity
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.GetPaymentsfororder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsfororderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**[]CFPaymentsEntity**](CFPaymentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

