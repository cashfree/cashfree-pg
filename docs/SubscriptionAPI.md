# \SubscriptionAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubsCreatePayment**](SubscriptionAPI.md#SubsCreatePayment) | **Post** /subscriptions/pay | Raise a charge or create an auth.
[**SubsCreatePlan**](SubscriptionAPI.md#SubsCreatePlan) | **Post** /plans | Create a plan.
[**SubsCreateRefund**](SubscriptionAPI.md#SubsCreateRefund) | **Post** /subscriptions/{subscription_id}/refunds | Create a refund.
[**SubsCreateSubscription**](SubscriptionAPI.md#SubsCreateSubscription) | **Post** /subscriptions | Create Subscription
[**SubsFetchPlan**](SubscriptionAPI.md#SubsFetchPlan) | **Get** /plans/{plan_id} | Fetch Plan
[**SubsFetchSubscription**](SubscriptionAPI.md#SubsFetchSubscription) | **Get** /subscriptions/{subscription_id} | Fetch Subscription
[**SubsFetchSubscriptionPayment**](SubscriptionAPI.md#SubsFetchSubscriptionPayment) | **Get** /subscriptions/{subscription_id}/payments/{payment_id} | Fetch details of a single payment.
[**SubsFetchSubscriptionPayments**](SubscriptionAPI.md#SubsFetchSubscriptionPayments) | **Get** /subscriptions/{subscription_id}/payments | Fetch details of all payments of a subscription.
[**SubsFetchSubscriptionRefund**](SubscriptionAPI.md#SubsFetchSubscriptionRefund) | **Get** /subscriptions/{subscription_id}/refunds/{refund_id} | Fetch details of a refund.
[**SubsManageSubscription**](SubscriptionAPI.md#SubsManageSubscription) | **Post** /subscriptions/{subscription_id}/manage | Manage a subscription.
[**SubsManageSubscriptionPayment**](SubscriptionAPI.md#SubsManageSubscriptionPayment) | **Post** /subscriptions/{subscription_id}/payments/{payment_id}/manage | Manage a single payment.
[**SubscriptionDocumentUpload**](SubscriptionAPI.md#SubscriptionDocumentUpload) | **Post** /subscriptions/pay/documents/{payment_id} | API to upload file for Physical Nach Authorization.
[**SubscriptionEligibility**](SubscriptionAPI.md#SubscriptionEligibility) | **Post** /subscriptions/eligibility/payment_methods | API to get all the payment method details available for subscription payments.



## SubsCreatePayment

> CreateSubscriptionPaymentResponse SubsCreatePayment(ctx).XApiVersion(xApiVersion).CreateSubscriptionPaymentRequest(createSubscriptionPaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Raise a charge or create an auth.



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
	createSubscriptionPaymentRequest := *openapiclient.NewCreateSubscriptionPaymentRequest("SubscriptionId_example", "PaymentId_example", "PaymentType_example") // CreateSubscriptionPaymentRequest | Request body to create a subscription payment.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsCreatePayment(context.Background()).XApiVersion(xApiVersion).CreateSubscriptionPaymentRequest(createSubscriptionPaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsCreatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsCreatePayment`: CreateSubscriptionPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsCreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubsCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **createSubscriptionPaymentRequest** | [**CreateSubscriptionPaymentRequest**](CreateSubscriptionPaymentRequest.md) | Request body to create a subscription payment. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**CreateSubscriptionPaymentResponse**](CreateSubscriptionPaymentResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsCreatePlan

> PlanEntity SubsCreatePlan(ctx).XApiVersion(xApiVersion).CreatePlanRequest(createPlanRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create a plan.



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
	createPlanRequest := *openapiclient.NewCreatePlanRequest("PlanId_example", "PlanName_example", "PlanType_example", float32(123)) // CreatePlanRequest | Request body to create a plan.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsCreatePlan(context.Background()).XApiVersion(xApiVersion).CreatePlanRequest(createPlanRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsCreatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsCreatePlan`: PlanEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsCreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubsCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **createPlanRequest** | [**CreatePlanRequest**](CreatePlanRequest.md) | Request body to create a plan. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsCreateRefund

> SubscriptionPaymentRefundEntity SubsCreateRefund(ctx, subscriptionId).XApiVersion(xApiVersion).CreateSubscriptionRefundRequest(createSubscriptionRefundRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create a refund.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	createSubscriptionRefundRequest := *openapiclient.NewCreateSubscriptionRefundRequest("SubscriptionId_example", "RefundId_example", float32(123)) // CreateSubscriptionRefundRequest | Request body to create a subscription refund.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsCreateRefund(context.Background(), subscriptionId).XApiVersion(xApiVersion).CreateSubscriptionRefundRequest(createSubscriptionRefundRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsCreateRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsCreateRefund`: SubscriptionPaymentRefundEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsCreateRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsCreateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **createSubscriptionRefundRequest** | [**CreateSubscriptionRefundRequest**](CreateSubscriptionRefundRequest.md) | Request body to create a subscription refund. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionPaymentRefundEntity**](SubscriptionPaymentRefundEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsCreateSubscription

> SubscriptionEntity SubsCreateSubscription(ctx).XApiVersion(xApiVersion).CreateSubscriptionRequest(createSubscriptionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create Subscription



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
	createSubscriptionRequest := *openapiclient.NewCreateSubscriptionRequest("SubscriptionId_example", *openapiclient.NewSubscriptionCustomerDetails("CustomerEmail_example", "CustomerPhone_example"), *openapiclient.NewCreateSubscriptionRequestPlanDetails()) // CreateSubscriptionRequest | Request body to create a subscription.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsCreateSubscription(context.Background()).XApiVersion(xApiVersion).CreateSubscriptionRequest(createSubscriptionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsCreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsCreateSubscription`: SubscriptionEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubsCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **createSubscriptionRequest** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md) | Request body to create a subscription. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionEntity**](SubscriptionEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsFetchPlan

> PlanEntity SubsFetchPlan(ctx, planId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Plan



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
	planId := "plan-id-123" // string | Provide the PlanId for which the details have to be fetched.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsFetchPlan(context.Background(), planId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsFetchPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsFetchPlan`: PlanEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsFetchPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Provide the PlanId for which the details have to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsFetchPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsFetchSubscription

> SubscriptionEntity SubsFetchSubscription(ctx, subscriptionId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch Subscription



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsFetchSubscription(context.Background(), subscriptionId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsFetchSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsFetchSubscription`: SubscriptionEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsFetchSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsFetchSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionEntity**](SubscriptionEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsFetchSubscriptionPayment

> SubscriptionPaymentEntity SubsFetchSubscriptionPayment(ctx, subscriptionId, paymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch details of a single payment.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	paymentId := "payment-id-123" // string | Provide the PaymentId using which the payment was created.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsFetchSubscriptionPayment(context.Background(), subscriptionId, paymentId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsFetchSubscriptionPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsFetchSubscriptionPayment`: SubscriptionPaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsFetchSubscriptionPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 
**paymentId** | **string** | Provide the PaymentId using which the payment was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsFetchSubscriptionPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionPaymentEntity**](SubscriptionPaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsFetchSubscriptionPayments

> []SubscriptionPaymentEntity SubsFetchSubscriptionPayments(ctx, subscriptionId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch details of all payments of a subscription.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsFetchSubscriptionPayments(context.Background(), subscriptionId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsFetchSubscriptionPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsFetchSubscriptionPayments`: []SubscriptionPaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsFetchSubscriptionPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsFetchSubscriptionPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]SubscriptionPaymentEntity**](SubscriptionPaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsFetchSubscriptionRefund

> SubscriptionPaymentRefundEntity SubsFetchSubscriptionRefund(ctx, subscriptionId, refundId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Fetch details of a refund.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	refundId := "refund-id-123" // string | Provide the PaymentId for which the details have to be fetched.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsFetchSubscriptionRefund(context.Background(), subscriptionId, refundId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsFetchSubscriptionRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsFetchSubscriptionRefund`: SubscriptionPaymentRefundEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsFetchSubscriptionRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 
**refundId** | **string** | Provide the PaymentId for which the details have to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsFetchSubscriptionRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]


 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionPaymentRefundEntity**](SubscriptionPaymentRefundEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsManageSubscription

> SubscriptionEntity SubsManageSubscription(ctx, subscriptionId).XApiVersion(xApiVersion).ManageSubscriptionRequest(manageSubscriptionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Manage a subscription.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	manageSubscriptionRequest := *openapiclient.NewManageSubscriptionRequest("SubscriptionId_example", "Action_example") // ManageSubscriptionRequest | Request body to manage a subscription.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsManageSubscription(context.Background(), subscriptionId).XApiVersion(xApiVersion).ManageSubscriptionRequest(manageSubscriptionRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsManageSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsManageSubscription`: SubscriptionEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsManageSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsManageSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **manageSubscriptionRequest** | [**ManageSubscriptionRequest**](ManageSubscriptionRequest.md) | Request body to manage a subscription. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionEntity**](SubscriptionEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubsManageSubscriptionPayment

> SubscriptionPaymentEntity SubsManageSubscriptionPayment(ctx, subscriptionId, paymentId).XApiVersion(xApiVersion).ManageSubscriptionPaymentRequest(manageSubscriptionPaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Manage a single payment.



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
	subscriptionId := "subscription-id-123" // string | Provide the SubscriptionId using which the subscription was created.
	paymentId := "payment-id-123" // string | Provide the PaymentId using which the payment was created.
	manageSubscriptionPaymentRequest := *openapiclient.NewManageSubscriptionPaymentRequest("SubscriptionId_example", "PaymentId_example", "Action_example") // ManageSubscriptionPaymentRequest | Request body to manage a subscription payment.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubsManageSubscriptionPayment(context.Background(), subscriptionId, paymentId).XApiVersion(xApiVersion).ManageSubscriptionPaymentRequest(manageSubscriptionPaymentRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubsManageSubscriptionPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubsManageSubscriptionPayment`: SubscriptionPaymentEntity
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubsManageSubscriptionPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Provide the SubscriptionId using which the subscription was created. | 
**paymentId** | **string** | Provide the PaymentId using which the payment was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubsManageSubscriptionPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]


 **manageSubscriptionPaymentRequest** | [**ManageSubscriptionPaymentRequest**](ManageSubscriptionPaymentRequest.md) | Request body to manage a subscription payment. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionPaymentEntity**](SubscriptionPaymentEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDocumentUpload

> UploadPnachImageResponse SubscriptionDocumentUpload(ctx, paymentId).XApiVersion(xApiVersion).File(file).PaymentId2(paymentId2).Action(action).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

API to upload file for Physical Nach Authorization.



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
	paymentId := "payment-id-123" // string | Provide the PaymentId using which the payment was created.
	file := os.NewFile(1234, "some_file") // *os.File | Select the .jpg file that should be uploaded or provide the path of that file. You cannot upload a file that is more than 1MB in size.
	paymentId2 := "paymentId_example" // string | Authorization Payment Id for physical nach authorization
	action := "action_example" // string | Action to be performed on the file. Can be SUBMIT_DOCUMENT
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionDocumentUpload(context.Background(), paymentId).XApiVersion(xApiVersion).File(file).PaymentId2(paymentId2).Action(action).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionDocumentUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionDocumentUpload`: UploadPnachImageResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionDocumentUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** | Provide the PaymentId using which the payment was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDocumentUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]

 **file** | ***os.File** | Select the .jpg file that should be uploaded or provide the path of that file. You cannot upload a file that is more than 1MB in size. | 
 **paymentId2** | **string** | Authorization Payment Id for physical nach authorization | 
 **action** | **string** | Action to be performed on the file. Can be SUBMIT_DOCUMENT | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**UploadPnachImageResponse**](UploadPnachImageResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionEligibility

> SubscriptionEligibilityResponse SubscriptionEligibility(ctx).XApiVersion(xApiVersion).SubscriptionEligibilityRequest(subscriptionEligibilityRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

API to get all the payment method details available for subscription payments.



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
	subscriptionEligibilityRequest := *openapiclient.NewSubscriptionEligibilityRequest(*openapiclient.NewSubscriptionEligibilityRequestQueries("SubscriptionId_example")) // SubscriptionEligibilityRequest | Request body to fetch subscription eligibile payment method details.
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionEligibility(context.Background()).XApiVersion(xApiVersion).SubscriptionEligibilityRequest(subscriptionEligibilityRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionEligibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionEligibility`: SubscriptionEligibilityResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionEligibility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **subscriptionEligibilityRequest** | [**SubscriptionEligibilityRequest**](SubscriptionEligibilityRequest.md) | Request body to fetch subscription eligibile payment method details. | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**SubscriptionEligibilityResponse**](SubscriptionEligibilityResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

