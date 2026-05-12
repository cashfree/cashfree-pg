# \IntellisenseAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDowntimeById**](IntellisenseAPI.md#FetchDowntimeById) | **Get** /incident/{id} | Fetch Downtime by ID



## FetchDowntimeById

> IncidentByIdResponse FetchDowntimeById(ctx, id).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Downtime by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v6"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    id := "id_example" 
    xApiVersion := "2026-01-01" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.FetchDowntimeById(&id, &xApiVersion, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.FetchDowntimeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDowntimeById`: IncidentByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.FetchDowntimeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the downtime. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchDowntimeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD. | [default to &quot;2026-01-01&quot;]
 **xRequestId** | **string** | Request ID for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to Cashfree. | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.  | 

### Return type

[**IncidentByIdResponse**](IncidentByIdResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

