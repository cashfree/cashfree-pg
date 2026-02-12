# \SoftPOSAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SposCreateTerminal**](SoftPOSAPI.md#SposCreateTerminal) | **Post** /terminal | Create Terminal
[**SposCreateTerminalTransaction**](SoftPOSAPI.md#SposCreateTerminalTransaction) | **Post** /terminal/transactions | Create Terminal Transaction
[**SposDemapSoundboxVpa**](SoftPOSAPI.md#SposDemapSoundboxVpa) | **Post** /terminal/demap/soundbox | Demap Soundbox Vpa
[**SposFetchTerminal**](SoftPOSAPI.md#SposFetchTerminal) | **Get** /terminal/{terminal_phone_no} | Get Terminal Status using Phone Number
[**SposFetchTerminalQRCodes**](SoftPOSAPI.md#SposFetchTerminalQRCodes) | **Get** /terminal/qrcodes | Fetch Terminal QR Codes
[**SposFetchTerminalSoundboxVpa**](SoftPOSAPI.md#SposFetchTerminalSoundboxVpa) | **Get** /terminal/soundbox/qrcodes | Fetch Terminal Soundbox vpa
[**SposFetchTerminalTransaction**](SoftPOSAPI.md#SposFetchTerminalTransaction) | **Get** /terminal/{cf_terminal_id}/payments | Get Terminal Transaction
[**SposOnboardSoundboxVpa**](SoftPOSAPI.md#SposOnboardSoundboxVpa) | **Post** /terminal/soundbox | Onboard Soundbox Vpa
[**SposUpdateSoundboxVpa**](SoftPOSAPI.md#SposUpdateSoundboxVpa) | **Patch** /terminal/{cf_terminal_id}/soundbox | Update Soundbox Vpa
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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	createTerminalRequest := *openapiclient.NewCreateTerminalRequest("TerminalId_example", "TerminalPhoneNo_example", "TerminalName_example", "TerminalEmail_example", "TerminalType_example") // CreateTerminalRequest | Request Body to Create Terminal for SPOS
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposCreateTerminal(context.Background()).XApiVersion(xApiVersion).CreateTerminalRequest(createTerminalRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposCreateTerminal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposCreateTerminal`: TerminalEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposCreateTerminal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposCreateTerminalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	createTerminalTransactionRequest := *openapiclient.NewCreateTerminalTransactionRequest("CfOrderId_example", "PaymentMethod_example") // CreateTerminalTransactionRequest | Request body to create a terminal transaction
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposCreateTerminalTransaction(context.Background()).XApiVersion(xApiVersion).CreateTerminalTransactionRequest(createTerminalTransactionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposCreateTerminalTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposCreateTerminalTransaction`: TerminalTransactionEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposCreateTerminalTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposCreateTerminalTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
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


## SposDemapSoundboxVpa

> []SoundboxVpaEntity SposDemapSoundboxVpa(ctx).XApiVersion(xApiVersion).DemapSoundboxVpaRequest(demapSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Demap Soundbox Vpa



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	demapSoundboxVpaRequest := *openapiclient.NewDemapSoundboxVpaRequest("CfTerminalId_example", "DeviceSerialNo_example") // DemapSoundboxVpaRequest | Request body to demap soundbox vpa
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposDemapSoundboxVpa(context.Background()).XApiVersion(xApiVersion).DemapSoundboxVpaRequest(demapSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposDemapSoundboxVpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposDemapSoundboxVpa`: []SoundboxVpaEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposDemapSoundboxVpa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposDemapSoundboxVpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **demapSoundboxVpaRequest** | [**DemapSoundboxVpaRequest**](DemapSoundboxVpaRequest.md) | Request body to demap soundbox vpa | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]SoundboxVpaEntity**](SoundboxVpaEntity.md)

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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	terminalPhoneNo := "6309291183" // string | The terminal for which you want to view the order details.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposFetchTerminal(context.Background(), terminalPhoneNo).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposFetchTerminal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposFetchTerminal`: TerminalEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposFetchTerminal`: %v\n", resp)
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	terminalPhoneNo := "9876543214" // string | Phone number assigned to the terminal. Required if you are not providing the cf_terminal_id in the request.
	cfTerminalId := "123344" // string | Cashfree terminal id for which you want to get staticQRs. Required if you are not providing the terminal_phone_number in the request.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposFetchTerminalQRCodes(context.Background()).XApiVersion(xApiVersion).TerminalPhoneNo(terminalPhoneNo).CfTerminalId(cfTerminalId).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposFetchTerminalQRCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposFetchTerminalQRCodes`: []FetchTerminalQRCodesEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposFetchTerminalQRCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposFetchTerminalQRCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
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


## SposFetchTerminalSoundboxVpa

> []SoundboxVpaEntity SposFetchTerminalSoundboxVpa(ctx).XApiVersion(xApiVersion).DeviceSerialNo(deviceSerialNo).CfTerminalId(cfTerminalId).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Terminal Soundbox vpa



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	deviceSerialNo := "9876543214" // string | Device Serial No assinged. Required if you are not providing the cf_terminal_id in the request.
	cfTerminalId := "123344" // string | Cashfree terminal id for which you want to get Soundbox Vpa. Required if you are not providing the device_serial_no in the request.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposFetchTerminalSoundboxVpa(context.Background()).XApiVersion(xApiVersion).DeviceSerialNo(deviceSerialNo).CfTerminalId(cfTerminalId).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposFetchTerminalSoundboxVpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposFetchTerminalSoundboxVpa`: []SoundboxVpaEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposFetchTerminalSoundboxVpa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposFetchTerminalSoundboxVpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **deviceSerialNo** | **string** | Device Serial No assinged. Required if you are not providing the cf_terminal_id in the request. | 
 **cfTerminalId** | **string** | Cashfree terminal id for which you want to get Soundbox Vpa. Required if you are not providing the device_serial_no in the request. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]SoundboxVpaEntity**](SoundboxVpaEntity.md)

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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	utr := "testUTR001" // string | Utr of the transaction.
	cfTerminalId := "123344" // string | Provide the Cashfree terminal ID for which the details have to be updated.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposFetchTerminalTransaction(context.Background(), cfTerminalId).XApiVersion(xApiVersion).Utr(utr).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposFetchTerminalTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposFetchTerminalTransaction`: TerminalPaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposFetchTerminalTransaction`: %v\n", resp)
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
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


## SposOnboardSoundboxVpa

> SoundboxVpaEntity SposOnboardSoundboxVpa(ctx).XApiVersion(xApiVersion).OnboardSoundboxVpaRequest(onboardSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Onboard Soundbox Vpa



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	onboardSoundboxVpaRequest := *openapiclient.NewOnboardSoundboxVpaRequest("Vpa_example", "CfTerminalId_example", "DeviceSerialNo_example") // OnboardSoundboxVpaRequest | Request body to onboard soundbox vpa
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposOnboardSoundboxVpa(context.Background()).XApiVersion(xApiVersion).OnboardSoundboxVpaRequest(onboardSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposOnboardSoundboxVpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposOnboardSoundboxVpa`: SoundboxVpaEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposOnboardSoundboxVpa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSposOnboardSoundboxVpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **onboardSoundboxVpaRequest** | [**OnboardSoundboxVpaRequest**](OnboardSoundboxVpaRequest.md) | Request body to onboard soundbox vpa | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SoundboxVpaEntity**](SoundboxVpaEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SposUpdateSoundboxVpa

> SoundboxVpaEntity SposUpdateSoundboxVpa(ctx, cfTerminalId).XApiVersion(xApiVersion).UpdateSoundboxVpaRequest(updateSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Update Soundbox Vpa



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	cfTerminalId := "123344" // string | Provide the Cashfree terminal ID for which the details have to be updated.
	updateSoundboxVpaRequest := *openapiclient.NewUpdateSoundboxVpaRequest("Vpa_example", "CfTerminalId_example") // UpdateSoundboxVpaRequest | Request body to update soundbox vpa
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposUpdateSoundboxVpa(context.Background(), cfTerminalId).XApiVersion(xApiVersion).UpdateSoundboxVpaRequest(updateSoundboxVpaRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposUpdateSoundboxVpa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposUpdateSoundboxVpa`: SoundboxVpaEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposUpdateSoundboxVpa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTerminalId** | **string** | Provide the Cashfree terminal ID for which the details have to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSposUpdateSoundboxVpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **updateSoundboxVpaRequest** | [**UpdateSoundboxVpaRequest**](UpdateSoundboxVpaRequest.md) | Request body to update soundbox vpa | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SoundboxVpaEntity**](SoundboxVpaEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	cfTerminalId := "123344" // string | Provide the Cashfree terminal ID for which the details have to be updated.
	updateTerminalRequest := *openapiclient.NewUpdateTerminalRequest("TerminalType_example") // UpdateTerminalRequest | Request Body to update terminal for SPOS.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposUpdateTerminal(context.Background(), cfTerminalId).XApiVersion(xApiVersion).UpdateTerminalRequest(updateTerminalRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposUpdateTerminal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposUpdateTerminal`: []UpdateTerminalEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposUpdateTerminal`: %v\n", resp)
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	cfTerminalId := "123344" // string | Provide the Cashfree terminal ID for which the details have to be updated.
	updateTerminalStatusRequest := *openapiclient.NewUpdateTerminalStatusRequest("TerminalStatus_example") // UpdateTerminalStatusRequest | Request Body to update terminal status for SPOS.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposUpdateTerminalStatus(context.Background(), cfTerminalId).XApiVersion(xApiVersion).UpdateTerminalStatusRequest(updateTerminalStatusRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposUpdateTerminalStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposUpdateTerminalStatus`: []UpdateTerminalEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposUpdateTerminalStatus`: %v\n", resp)
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

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
	openapiclient "github.com/cashfree/cashfree-pg"
)

func main() {
	xApiVersion := "2025-01-01" // string | API version to be used. Format is in YYYY-MM-DD (default to "2025-01-01")
	cfTerminalId := "123344" // string | Provide the Cashfree terminal ID for which the details have to be updated.
	uploadTerminalDocs := *openapiclient.NewUploadTerminalDocs("DocType_example", "DocValue_example", "File_example") // UploadTerminalDocs | Request Body to update terminal documents for SPOS.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SoftPOSAPI.SposUploadTerminalDocs(context.Background(), cfTerminalId).XApiVersion(xApiVersion).UploadTerminalDocs(uploadTerminalDocs).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SoftPOSAPI.SposUploadTerminalDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SposUploadTerminalDocs`: []UploadTerminalDocsEntity
	fmt.Fprintf(os.Stdout, "Response from `SoftPOSAPI.SposUploadTerminalDocs`: %v\n", resp)
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
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

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

