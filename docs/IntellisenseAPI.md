# \IntellisenseAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchActiveEcosystemDowntimes**](IntellisenseAPI.md#FetchActiveEcosystemDowntimes) | **Get** /intellisense/downtime | Fetch Active Ecosystem Level Downtimes
[**FetchDowntimeById**](IntellisenseAPI.md#FetchDowntimeById) | **Get** /intellisense/downtime/{id} | Fetch Downtime by ID



## FetchActiveEcosystemDowntimes

> FetchActiveEcosystemDowntimes200Response FetchActiveEcosystemDowntimes(ctx).XApiVersion(xApiVersion).Execute()

Fetch Active Ecosystem Level Downtimes



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

    resp, r, err := cashfree.FetchActiveEcosystemDowntimes(&xApiVersion, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.FetchActiveEcosystemDowntimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchActiveEcosystemDowntimes`: FetchActiveEcosystemDowntimes200Response
    fmt.Fprintf(os.Stdout, "Response from `cashfree.FetchActiveEcosystemDowntimes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchActiveEcosystemDowntimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2026-01-01&quot;]

### Return type

[**FetchActiveEcosystemDowntimes200Response**](FetchActiveEcosystemDowntimes200Response.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDowntimeById

> DowntimeByIdResponse FetchDowntimeById(ctx, id).XApiVersion(xApiVersion).Execute()

Fetch Downtime by ID



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

    id := "id_example" 
    xApiVersion := "2026-01-01" 

    resp, r, err := cashfree.FetchDowntimeById(&id, &xApiVersion, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.FetchDowntimeById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDowntimeById`: DowntimeByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.FetchDowntimeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the downtime incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchDowntimeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2026-01-01&quot;]

### Return type

[**DowntimeByIdResponse**](DowntimeByIdResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

