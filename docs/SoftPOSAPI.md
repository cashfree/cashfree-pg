# \SoftPOSAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SposCreateTerminal**](SoftPOSAPI.md#SposCreateTerminal) | **Post** /terminal | Create Terminal
[**SposCreateTerminalTransaction**](SoftPOSAPI.md#SposCreateTerminalTransaction) | **Post** /terminal/transactions | Create Terminal Transaction
[**SposFetchTerminal**](SoftPOSAPI.md#SposFetchTerminal) | **Get** /terminal/{terminal_phone_no} | Get Terminal Status using Phone Number
[**SposFetchTerminalQRCodes**](SoftPOSAPI.md#SposFetchTerminalQRCodes) | **Get** /terminal/qrcodes | Fetch Terminal QR Codes
[**SposFetchTerminalTransaction**](SoftPOSAPI.md#SposFetchTerminalTransaction) | **Get** /terminal/{cf_terminal_id}/payments | Get Terminal Transaction
[**SposUpdateTerminal**](SoftPOSAPI.md#SposUpdateTerminal) | **Patch** /terminal/{cf_terminal_id} | Update Terminal
[**SposUpdateTerminalStatus**](SoftPOSAPI.md#SposUpdateTerminalStatus) | **Patch** /terminal/{cf_terminal_id}/status | Update Terminal Status
[**SposUploadTerminalDocs**](SoftPOSAPI.md#SposUploadTerminalDocs) | **Post** /terminal/{cf_terminal_id}/docs | Upload Terminal Docs



## SposCreateTerminal

> TerminalEntity SposCreateTerminal(ctx).XApiVersion(xApiVersion).CreateTerminalRequest(createTerminalRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create Terminal



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

    xApiVersion := "2023-08-01" 
    createTerminalRequest := *cashfree.NewCreateTerminalRequest("TerminalId_example", "TerminalPhoneNo_example", "TerminalName_example", "TerminalEmail_example", "TerminalType_example") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposCreateTerminal(&xApiVersion, &createTerminalRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposCreateTerminal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposCreateTerminal`: TerminalEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposCreateTerminal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposCreateTerminalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **createTerminalRequest** | [**CreateTerminalRequest**](CreateTerminalRequest.md) | Request Body to Create Terminal for SPOS | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**TerminalEntity**](TerminalEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposCreateTerminalTransaction

> TerminalTransactionEntity SposCreateTerminalTransaction(ctx).XApiVersion(xApiVersion).CreateTerminalTransactionRequest(createTerminalTransactionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create Terminal Transaction



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

    xApiVersion := "2023-08-01" 
    createTerminalTransactionRequest := *cashfree.NewCreateTerminalTransactionRequest("CfOrderId_example", "PaymentMethod_example") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposCreateTerminalTransaction(&xApiVersion, &createTerminalTransactionRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposCreateTerminalTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposCreateTerminalTransaction`: TerminalTransactionEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposCreateTerminalTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposCreateTerminalTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **createTerminalTransactionRequest** | [**CreateTerminalTransactionRequest**](CreateTerminalTransactionRequest.md) | Request body to create a terminal transaction | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**TerminalTransactionEntity**](TerminalTransactionEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposFetchTerminal

> TerminalEntity SposFetchTerminal(ctx, terminalPhoneNo).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Terminal Status using Phone Number



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

    xApiVersion := "2023-08-01" 
    terminalPhoneNo := "6309291183" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposFetchTerminal(&xApiVersion, &terminalPhoneNo, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposFetchTerminal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposFetchTerminal`: TerminalEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposFetchTerminal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**terminalPhoneNo** | **string** | The terminal for which you want to view the order details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposFetchTerminalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**TerminalEntity**](TerminalEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposFetchTerminalQRCodes

> []FetchTerminalQRCodesEntity SposFetchTerminalQRCodes(ctx).XApiVersion(xApiVersion).TerminalPhoneNo(terminalPhoneNo).CfTerminalId(cfTerminalId).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Terminal QR Codes



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

    xApiVersion := "2023-08-01" 
    terminalPhoneNo := "9876543214" 
    cfTerminalId := "123344" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposFetchTerminalQRCodes(&xApiVersion, &terminalPhoneNo, &cfTerminalId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposFetchTerminalQRCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposFetchTerminalQRCodes`: []FetchTerminalQRCodesEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposFetchTerminalQRCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposFetchTerminalQRCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **terminalPhoneNo** | **string** | Phone number assigned to the terminal. Required if you are not providing the cf_terminal_id in the request. | 
 **cfTerminalId** | **string** | Cashfree terminal id for which you want to get staticQRs. Required if you are not providing the terminal_phone_number in the request. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]FetchTerminalQRCodesEntity**](FetchTerminalQRCodesEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposFetchTerminalTransaction

> TerminalPaymentEntity SposFetchTerminalTransaction(ctx, cfTerminalId).XApiVersion(xApiVersion).Utr(utr).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Terminal Transaction



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

    xApiVersion := "2023-08-01" 
    utr := "testUTR001" 
    cfTerminalId := "123344" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposFetchTerminalTransaction(&xApiVersion, &utr, &cfTerminalId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposFetchTerminalTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposFetchTerminalTransaction`: TerminalPaymentEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposFetchTerminalTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTerminalId** | **string** | Provide the Cashfree terminal ID for which the details have to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposFetchTerminalTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **utr** | **string** | Utr of the transaction. | 

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**TerminalPaymentEntity**](TerminalPaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposUpdateTerminal

> []UpdateTerminalEntity SposUpdateTerminal(ctx, cfTerminalId).XApiVersion(xApiVersion).UpdateTerminalRequest(updateTerminalRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Update Terminal



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

    xApiVersion := "2023-08-01" 
    cfTerminalId := "123344" 
    updateTerminalRequest := *cashfree.NewUpdateTerminalRequest("TerminalType_example") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposUpdateTerminal(&xApiVersion, &cfTerminalId, &updateTerminalRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposUpdateTerminal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposUpdateTerminal`: []UpdateTerminalEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposUpdateTerminal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTerminalId** | **string** | Provide the Cashfree terminal ID for which the details have to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposUpdateTerminalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **updateTerminalRequest** | [**UpdateTerminalRequest**](UpdateTerminalRequest.md) | Request Body to update terminal for SPOS. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]UpdateTerminalEntity**](UpdateTerminalEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposUpdateTerminalStatus

> []UpdateTerminalEntity SposUpdateTerminalStatus(ctx, cfTerminalId).XApiVersion(xApiVersion).UpdateTerminalStatusRequest(updateTerminalStatusRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Update Terminal Status



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

    xApiVersion := "2023-08-01" 
    cfTerminalId := "123344" 
    updateTerminalStatusRequest := *cashfree.NewUpdateTerminalStatusRequest("TerminalStatus_example") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposUpdateTerminalStatus(&xApiVersion, &cfTerminalId, &updateTerminalStatusRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposUpdateTerminalStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposUpdateTerminalStatus`: []UpdateTerminalEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposUpdateTerminalStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTerminalId** | **string** | Provide the Cashfree terminal ID for which the details have to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposUpdateTerminalStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **updateTerminalStatusRequest** | [**UpdateTerminalStatusRequest**](UpdateTerminalStatusRequest.md) | Request Body to update terminal status for SPOS. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]UpdateTerminalEntity**](UpdateTerminalEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposUploadTerminalDocs

> []UploadTerminalDocsEntity SposUploadTerminalDocs(ctx, cfTerminalId).XApiVersion(xApiVersion).UploadTerminalDocs(uploadTerminalDocs).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Upload Terminal Docs



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

    xApiVersion := "2023-08-01" 
    cfTerminalId := "123344" 
    uploadTerminalDocs := *cashfree.NewUploadTerminalDocs("DocType_example", "DocValue_example", "File_example") 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.SposUploadTerminalDocs(&xApiVersion, &cfTerminalId, &uploadTerminalDocs, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.SposUploadTerminalDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SposUploadTerminalDocs`: []UploadTerminalDocsEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.SposUploadTerminalDocs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTerminalId** | **string** | Provide the Cashfree terminal ID for which the details have to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposUploadTerminalDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **uploadTerminalDocs** | [**UploadTerminalDocs**](UploadTerminalDocs.md) | Request Body to update terminal documents for SPOS. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]UploadTerminalDocsEntity**](UploadTerminalDocsEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

