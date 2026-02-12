# \PaymentsAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGAuthorizeOrder**](PaymentsAPI.md#PGAuthorizeOrder) | **Post** /orders/{order_id}/authorization | Preauthorization
[**PGOrderAuthenticatePayment**](PaymentsAPI.md#PGOrderAuthenticatePayment) | **Post** /orders/pay/authenticate/{cf_payment_id} | Submit or Resend OTP
[**PGOrderFetchPayment**](PaymentsAPI.md#PGOrderFetchPayment) | **Get** /orders/{order_id}/payments/{cf_payment_id} | Get Payment by ID
[**PGOrderFetchPayments**](PaymentsAPI.md#PGOrderFetchPayments) | **Get** /orders/{order_id}/payments | Get Payments for an Order
[**PGPayOrder**](PaymentsAPI.md#PGPayOrder) | **Post** /orders/sessions | Order Pay



## PGAuthorizeOrder

> PaymentEntity PGAuthorizeOrder(ctx, orderId).XApiVersion(xApiVersion).AuthorizeOrderRequest(authorizeOrderRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Preauthorization



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
	orderId := "your-order-id" // string | The id which uniquely identifies your order
	authorizeOrderRequest := *openapiclient.NewAuthorizeOrderRequest() // AuthorizeOrderRequest | Request to Capture or Void Transactions
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PGAuthorizeOrder(context.Background(), orderId).XApiVersion(xApiVersion).AuthorizeOrderRequest(authorizeOrderRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PGAuthorizeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGAuthorizeOrder`: PaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PGAuthorizeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id which uniquely identifies your order | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGAuthorizeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **authorizeOrderRequest** | [**AuthorizeOrderRequest**](AuthorizeOrderRequest.md) | Request to Capture or Void Transactions | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PaymentEntity**](PaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGOrderAuthenticatePayment

> OrderAuthenticateEntity PGOrderAuthenticatePayment(ctx, cfPaymentId).XApiVersion(xApiVersion).OrderAuthenticatePaymentRequest(orderAuthenticatePaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Submit or Resend OTP



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
	cfPaymentId := "121224562" // string | The Cashfree payment or transaction ID.
	orderAuthenticatePaymentRequest := *openapiclient.NewOrderAuthenticatePaymentRequest("Otp_example", "Action_example") // OrderAuthenticatePaymentRequest | Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PGOrderAuthenticatePayment(context.Background(), cfPaymentId).XApiVersion(xApiVersion).OrderAuthenticatePaymentRequest(orderAuthenticatePaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PGOrderAuthenticatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGOrderAuthenticatePayment`: OrderAuthenticateEntity
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PGOrderAuthenticatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfPaymentId** | **string** | The Cashfree payment or transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGOrderAuthenticatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **orderAuthenticatePaymentRequest** | [**OrderAuthenticatePaymentRequest**](OrderAuthenticatePaymentRequest.md) | Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**OrderAuthenticateEntity**](OrderAuthenticateEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGOrderFetchPayment

> PaymentEntity PGOrderFetchPayment(ctx, orderId, cfPaymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Payment by ID



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
	orderId := "your-order-id" // string | The id which uniquely identifies your order
	cfPaymentId := "121224562" // string | The Cashfree payment or transaction ID.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PGOrderFetchPayment(context.Background(), orderId, cfPaymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PGOrderFetchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGOrderFetchPayment`: PaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PGOrderFetchPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id which uniquely identifies your order | 
**cfPaymentId** | **string** | The Cashfree payment or transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGOrderFetchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PaymentEntity**](PaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGOrderFetchPayments

> []PaymentEntity PGOrderFetchPayments(ctx, orderId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Payments for an Order



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
	orderId := "your-order-id" // string | The id which uniquely identifies your order
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PGOrderFetchPayments(context.Background(), orderId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PGOrderFetchPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGOrderFetchPayments`: []PaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PGOrderFetchPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id which uniquely identifies your order | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGOrderFetchPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]PaymentEntity**](PaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGPayOrder

> PayOrderEntity PGPayOrder(ctx).XApiVersion(xApiVersion).PayOrderRequest(payOrderRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Order Pay



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
	payOrderRequest := *openapiclient.NewPayOrderRequest("session__CvcEmNKDkmERQrxnx39ibhJ3Ii034pjc8ZVxf3qcgEXCWlgDDlHRgz2XYZCqpajDQSXMMtCusPgOIxYP2LZx0-05p39gC2Vgmq1RAj--gcn", openapiclient.PayOrderRequest_payment_method{AppPaymentMethod: openapiclient.NewAppPaymentMethod(*openapiclient.NewApp("Channel_example", "Provider_example", "Phone_example"))}) // PayOrderRequest | Request body to create a transaction at cashfree using `payment_session_id`
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PGPayOrder(context.Background()).XApiVersion(xApiVersion).PayOrderRequest(payOrderRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PGPayOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGPayOrder`: PayOrderEntity
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PGPayOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGPayOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **payOrderRequest** | [**PayOrderRequest**](PayOrderRequest.md) | Request body to create a transaction at cashfree using &#x60;payment_session_id&#x60; | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PayOrderEntity**](PayOrderEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

