# \TokenVaultAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGCustomerDeleteInstrument**](TokenVaultAPI.md#PGCustomerDeleteInstrument) | **Delete** /customers/{customer_id}/instruments/{instrument_id} | Delete Saved Card Instrument
[**PGCustomerFetchInstrument**](TokenVaultAPI.md#PGCustomerFetchInstrument) | **Get** /customers/{customer_id}/instruments/{instrument_id} | Fetch Specific Saved Card Instrument
[**PGCustomerFetchInstruments**](TokenVaultAPI.md#PGCustomerFetchInstruments) | **Get** /customers/{customer_id}/instruments | Fetch All Saved Card Instrument
[**PGCustomerInstrumentsFetchCryptogram**](TokenVaultAPI.md#PGCustomerInstrumentsFetchCryptogram) | **Get** /customers/{customer_id}/instruments/{instrument_id}/cryptogram | Fetch cryptogram for a saved card instrument



## PGCustomerDeleteInstrument

> InstrumentEntity PGCustomerDeleteInstrument(ctx, customerId, instrumentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Delete Saved Card Instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2023-08-01" 
    customerId := "your-customer-id" 
    instrumentId := "some-instrument-id" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCustomerDeleteInstrument(&xApiVersion, &customerId, &instrumentId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCustomerDeleteInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCustomerDeleteInstrument`: InstrumentEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCustomerDeleteInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string** | The instrument_id which needs to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGCustomerDeleteInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**InstrumentEntity**](InstrumentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGCustomerFetchInstrument

> InstrumentEntity PGCustomerFetchInstrument(ctx, customerId, instrumentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Specific Saved Card Instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2023-08-01" 
    customerId := "your-customer-id" 
    instrumentId := "some-instrument-id" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCustomerFetchInstrument(&xApiVersion, &customerId, &instrumentId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCustomerFetchInstrument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCustomerFetchInstrument`: InstrumentEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCustomerFetchInstrument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string** | The instrument_id of the saved instrument which needs to be queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGCustomerFetchInstrumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**InstrumentEntity**](InstrumentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGCustomerFetchInstruments

> []InstrumentEntity PGCustomerFetchInstruments(ctx, customerId).XApiVersion(xApiVersion).InstrumentType(instrumentType).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch All Saved Card Instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2023-08-01" 
    customerId := "your-customer-id" 
    instrumentType := "card" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCustomerFetchInstruments(&xApiVersion, &customerId, &instrumentType, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCustomerFetchInstruments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCustomerFetchInstruments`: []InstrumentEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCustomerFetchInstruments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGCustomerFetchInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **instrumentType** | **string** | Payment mode or type of saved instrument  | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]InstrumentEntity**](InstrumentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGCustomerInstrumentsFetchCryptogram

> CryptogramEntity PGCustomerInstrumentsFetchCryptogram(ctx, customerId, instrumentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch cryptogram for a saved card instrument



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2023-08-01" 
    customerId := "your-customer-id" 
    instrumentId := "some-instrument-id" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCustomerInstrumentsFetchCryptogram(&xApiVersion, &customerId, &instrumentId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCustomerInstrumentsFetchCryptogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCustomerInstrumentsFetchCryptogram`: CryptogramEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCustomerInstrumentsFetchCryptogram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string** | The instrument_id of the saved card instrument which needs to be queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGCustomerInstrumentsFetchCryptogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**CryptogramEntity**](CryptogramEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

