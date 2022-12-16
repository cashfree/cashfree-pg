# \TokenVaultApi

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpecificSavedInstrument**](TokenVaultApi.md#DeleteSpecificSavedInstrument) | **Delete** /customers/{customer_id}/instruments/{instrument_id} | Delete Saved Instrument
[**FetchAllSavedInstruments**](TokenVaultApi.md#FetchAllSavedInstruments) | **Get** /customers/{customer_id}/instruments | Fetch All Saved Instruments
[**FetchCryptogram**](TokenVaultApi.md#FetchCryptogram) | **Get** /customers/{customer_id}/instruments/{instrument_id}/cryptogram | Fetch cryptogram for saved instrument
[**FetchSpecificSavedInstrument**](TokenVaultApi.md#FetchSpecificSavedInstrument) | **Get** /customers/{customer_id}/instruments/{instrument_id} | Fetch Single Saved Instrument



## DeleteSpecificSavedInstrument

> CFFetchAllSavedInstruments DeleteSpecificSavedInstrument(ctx, customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).Execute()

Delete Saved Instrument



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
    customerId := "customerId_example" // string | 
    instrumentId := "instrumentId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenVaultApi.DeleteSpecificSavedInstrument(context.Background(), customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenVaultApi.DeleteSpecificSavedInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSpecificSavedInstrument`: CFFetchAllSavedInstruments
    fmt.Fprintf(os.Stdout, "Response from `TokenVaultApi.DeleteSpecificSavedInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**instrumentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpecificSavedInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 


 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]

### Return type

[**CFFetchAllSavedInstruments**](CFFetchAllSavedInstruments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllSavedInstruments

> []CFFetchAllSavedInstruments FetchAllSavedInstruments(ctx, customerId).XClientId(xClientId).XClientSecret(xClientSecret).InstrumentType(instrumentType).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Fetch All Saved Instruments



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
    customerId := "customerId_example" // string | 
    instrumentType := "instrumentType_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenVaultApi.FetchAllSavedInstruments(context.Background(), customerId).XClientId(xClientId).XClientSecret(xClientSecret).InstrumentType(instrumentType).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenVaultApi.FetchAllSavedInstruments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllSavedInstruments`: []CFFetchAllSavedInstruments
    fmt.Fprintf(os.Stdout, "Response from `TokenVaultApi.FetchAllSavedInstruments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllSavedInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 

 **instrumentType** | **string** |  | 
 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**[]CFFetchAllSavedInstruments**](CFFetchAllSavedInstruments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCryptogram

> CFCryptogram FetchCryptogram(ctx, customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Fetch cryptogram for saved instrument



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
    customerId := "customerId_example" // string | 
    instrumentId := "instrumentId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenVaultApi.FetchCryptogram(context.Background(), customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenVaultApi.FetchCryptogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCryptogram`: CFCryptogram
    fmt.Fprintf(os.Stdout, "Response from `TokenVaultApi.FetchCryptogram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**instrumentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchCryptogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 


 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFCryptogram**](CFCryptogram.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSpecificSavedInstrument

> CFFetchAllSavedInstruments FetchSpecificSavedInstrument(ctx, customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()

Fetch Single Saved Instrument



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
    customerId := "customerId_example" // string | 
    instrumentId := "instrumentId_example" // string | 
    xApiVersion := "xApiVersion_example" // string |  (optional) (default to "2022-01-01")
    xIdempotencyReplayed := true // bool |  (optional) (default to false)
    xIdempotencyKey := "xIdempotencyKey_example" // string |  (optional)
    xRequestId := "xRequestId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenVaultApi.FetchSpecificSavedInstrument(context.Background(), customerId, instrumentId).XClientId(xClientId).XClientSecret(xClientSecret).XApiVersion(xApiVersion).XIdempotencyReplayed(xIdempotencyReplayed).XIdempotencyKey(xIdempotencyKey).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenVaultApi.FetchSpecificSavedInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSpecificSavedInstrument`: CFFetchAllSavedInstruments
    fmt.Fprintf(os.Stdout, "Response from `TokenVaultApi.FetchSpecificSavedInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** |  | 
**instrumentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSpecificSavedInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xClientId** | **string** |  | 
 **xClientSecret** | **string** |  | 


 **xApiVersion** | **string** |  | [default to &quot;2022-01-01&quot;]
 **xIdempotencyReplayed** | **bool** |  | [default to false]
 **xIdempotencyKey** | **string** |  | 
 **xRequestId** | **string** |  | 

### Return type

[**CFFetchAllSavedInstruments**](CFFetchAllSavedInstruments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

