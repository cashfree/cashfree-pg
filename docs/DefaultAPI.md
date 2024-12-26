# \DefaultAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerminalCreateQRCodes**](DefaultAPI.md#TerminalCreateQRCodes) | **Post** /partners/merchant/qrcodes | Create Pre-Activated Vpas for partner
[**TerminalGetQRCodes**](DefaultAPI.md#TerminalGetQRCodes) | **Get** /partners/merchant/qrcodes | Get Pre-Activated Vpas for partner



## TerminalCreateQRCodes

> []StaticQrResponseEntity TerminalCreateQRCodes(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).CreatePartnerVpaRequest(createPartnerVpaRequest).Execute()

Create Pre-Activated Vpas for partner



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
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 
    createPartnerVpaRequest := *cashfree.NewCreatePartnerVpaRequest(int32(123)) 

    resp, r, err := cashfree.TerminalCreateQRCodes(&xApiVersion, &xRequestId, &xIdempotencyKey, &createPartnerVpaRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.TerminalCreateQRCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminalCreateQRCodes`: []StaticQrResponseEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.TerminalCreateQRCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminalCreateQRCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 
 **createPartnerVpaRequest** | [**CreatePartnerVpaRequest**](CreatePartnerVpaRequest.md) |  | 

### Return type

[**[]StaticQrResponseEntity**](StaticQrResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminalGetQRCodes

> []StaticQrResponseEntity TerminalGetQRCodes(ctx).XApiVersion(xApiVersion).Status(status).CfTerminalId(cfTerminalId).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Pre-Activated Vpas for partner



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
    status := "MAPPED" 
    cfTerminalId := "123344" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.TerminalGetQRCodes(&xApiVersion, &status, &cfTerminalId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.TerminalGetQRCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminalGetQRCodes`: []StaticQrResponseEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.TerminalGetQRCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminalGetQRCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **status** | **string** | Status of pre-created Qr. | 
 **cfTerminalId** | **string** | Cashfree terminal id for which you want to get pre-generated staticQRs. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]StaticQrResponseEntity**](StaticQrResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

