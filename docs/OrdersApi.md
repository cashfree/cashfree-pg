# \OrdersApi

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /orders | Create Order
[**GetOrder**](OrdersApi.md#GetOrder) | **Get** /orders/{order_id} | Get Order
[**OrderPay**](OrdersApi.md#OrderPay) | **Post** /orders/pay | Order Pay
[**Preauthorization**](OrdersApi.md#Preauthorization) | **Post** /orders/{order_id}/authorization | Preauthorization



## CreateOrder

> CFOrder CreateOrder(ctx).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFOrderRequest(cFOrderRequest).Execute()

Create Order



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
    cFOrderRequest := *openapiclient.NewCFOrderRequest(float64(10.15), "INR", *openapiclient.NewCFCustomerDetails("CustomerId_example", "CustomerEmail_example", "CustomerPhone_example", "CustomerBankAccountNumber_example", "CustomerBankIfsc_example", int32(123))) // CFOrderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateOrder(context.Background()).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFOrderRequest(cFOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrder`: CFOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 
 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 
 **cFOrderRequest** | [**CFOrderRequest**](CFOrderRequest.md) |  | 

### Return type

[**CFOrder**](CFOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> CFOrder GetOrder(ctx, orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Get Order



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
    resp, r, err := apiClient.OrdersApi.GetOrder(context.Background(), orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: CFOrder
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFOrder**](CFOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPay

> CFOrderPayResponse OrderPay(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CFOrderPayRequest(cFOrderPayRequest).Execute()

Order Pay



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
    xApiVersion := "xApiVersion_example" // string | 
    xRequestId := "xRequestId_example" // string |  (optional)
    cFOrderPayRequest := *openapiclient.NewCFOrderPayRequest("hyj120nbvt12831", openapiclient.CFPaymentMethod{CFAppPayment: openapiclient.NewCFAppPayment(*openapiclient.NewCFApp("Channel_example", "Provider_example", "Phone_example"))}) // CFOrderPayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderPay(context.Background()).XApiVersion(xApiVersion).XRequestId(xRequestId).CFOrderPayRequest(cFOrderPayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderPay`: CFOrderPayResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderPay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** |  | 
 **xRequestId** | **string** |  | 
 **cFOrderPayRequest** | [**CFOrderPayRequest**](CFOrderPayRequest.md) |  | 

### Return type

[**CFOrderPayResponse**](CFOrderPayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Preauthorization

> CFPaymentsEntity Preauthorization(ctx, orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFAuthorizationRequest(cFAuthorizationRequest).Execute()

Preauthorization



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
    cFAuthorizationRequest := *openapiclient.NewCFAuthorizationRequest() // CFAuthorizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.Preauthorization(context.Background(), orderId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).CFAuthorizationRequest(cFAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.Preauthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Preauthorization`: CFPaymentsEntity
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.Preauthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreauthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 
 **cFAuthorizationRequest** | [**CFAuthorizationRequest**](CFAuthorizationRequest.md) |  | 

### Return type

[**CFPaymentsEntity**](CFPaymentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

