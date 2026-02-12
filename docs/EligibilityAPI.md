# \EligibilityAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGEligibilityFetchCardlessEMI**](EligibilityAPI.md#PGEligibilityFetchCardlessEMI) | **Post** /eligibility/cardlessemi | Get Eligible Cardless EMI Payment Methods for a customer on an order
[**PGEligibilityFetchOffers**](EligibilityAPI.md#PGEligibilityFetchOffers) | **Post** /eligibility/offers | Get Eligible Offers for an Order
[**PGEligibilityFetchPaylater**](EligibilityAPI.md#PGEligibilityFetchPaylater) | **Post** /eligibility/paylater | Get Eligible Paylater for a customer on an order
[**PGEligibilityFetchPaymentMethods**](EligibilityAPI.md#PGEligibilityFetchPaymentMethods) | **Post** /eligibility/payment_methods | Get eligible Payment Methods



## PGEligibilityFetchCardlessEMI

> []EligibilityCardlessEMIEntity PGEligibilityFetchCardlessEMI(ctx).XApiVersion(xApiVersion).EligibilityFetchCardlessEMIRequest(eligibilityFetchCardlessEMIRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Eligible Cardless EMI Payment Methods for a customer on an order



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
	eligibilityFetchCardlessEMIRequest := *openapiclient.NewEligibilityFetchCardlessEMIRequest(*openapiclient.NewCardlessEMIQueries()) // EligibilityFetchCardlessEMIRequest | Request Body to get eligible cardless emi options for a customer and order
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EligibilityAPI.PGEligibilityFetchCardlessEMI(context.Background()).XApiVersion(xApiVersion).EligibilityFetchCardlessEMIRequest(eligibilityFetchCardlessEMIRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EligibilityAPI.PGEligibilityFetchCardlessEMI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGEligibilityFetchCardlessEMI`: []EligibilityCardlessEMIEntity
	fmt.Fprintf(os.Stdout, "Response from `EligibilityAPI.PGEligibilityFetchCardlessEMI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGEligibilityFetchCardlessEMIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **eligibilityFetchCardlessEMIRequest** | [**EligibilityFetchCardlessEMIRequest**](EligibilityFetchCardlessEMIRequest.md) | Request Body to get eligible cardless emi options for a customer and order | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]EligibilityCardlessEMIEntity**](EligibilityCardlessEMIEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGEligibilityFetchOffers

> []EligibilityOfferEntity PGEligibilityFetchOffers(ctx).XApiVersion(xApiVersion).EligibilityFetchOffersRequest(eligibilityFetchOffersRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Eligible Offers for an Order



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
	eligibilityFetchOffersRequest := *openapiclient.NewEligibilityFetchOffersRequest(*openapiclient.NewOfferQueries()) // EligibilityFetchOffersRequest | Request Body to get eligible offers for a customer and order
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EligibilityAPI.PGEligibilityFetchOffers(context.Background()).XApiVersion(xApiVersion).EligibilityFetchOffersRequest(eligibilityFetchOffersRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EligibilityAPI.PGEligibilityFetchOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGEligibilityFetchOffers`: []EligibilityOfferEntity
	fmt.Fprintf(os.Stdout, "Response from `EligibilityAPI.PGEligibilityFetchOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGEligibilityFetchOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **eligibilityFetchOffersRequest** | [**EligibilityFetchOffersRequest**](EligibilityFetchOffersRequest.md) | Request Body to get eligible offers for a customer and order | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]EligibilityOfferEntity**](EligibilityOfferEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGEligibilityFetchPaylater

> []EligibilityPaylaterEntity PGEligibilityFetchPaylater(ctx).XApiVersion(xApiVersion).EligibilityFetchPaylaterRequest(eligibilityFetchPaylaterRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Eligible Paylater for a customer on an order



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
	eligibilityFetchPaylaterRequest := *openapiclient.NewEligibilityFetchPaylaterRequest(*openapiclient.NewCardlessEMIQueries()) // EligibilityFetchPaylaterRequest | Request Body to get eligible paylater options for a customer and order
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EligibilityAPI.PGEligibilityFetchPaylater(context.Background()).XApiVersion(xApiVersion).EligibilityFetchPaylaterRequest(eligibilityFetchPaylaterRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EligibilityAPI.PGEligibilityFetchPaylater``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGEligibilityFetchPaylater`: []EligibilityPaylaterEntity
	fmt.Fprintf(os.Stdout, "Response from `EligibilityAPI.PGEligibilityFetchPaylater`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGEligibilityFetchPaylaterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **eligibilityFetchPaylaterRequest** | [**EligibilityFetchPaylaterRequest**](EligibilityFetchPaylaterRequest.md) | Request Body to get eligible paylater options for a customer and order | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]EligibilityPaylaterEntity**](EligibilityPaylaterEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGEligibilityFetchPaymentMethods

> []EligibilityPaymentMethodsEntity PGEligibilityFetchPaymentMethods(ctx).XApiVersion(xApiVersion).EligibilityFetchPaymentMethodsRequest(eligibilityFetchPaymentMethodsRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get eligible Payment Methods



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
	eligibilityFetchPaymentMethodsRequest := *openapiclient.NewEligibilityFetchPaymentMethodsRequest(*openapiclient.NewPaymentMethodsQueries()) // EligibilityFetchPaymentMethodsRequest | Request Body to get eligible payment methods for an account and order
	xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" // string | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree (optional)
	xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" // string | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EligibilityAPI.PGEligibilityFetchPaymentMethods(context.Background()).XApiVersion(xApiVersion).EligibilityFetchPaymentMethodsRequest(eligibilityFetchPaymentMethodsRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EligibilityAPI.PGEligibilityFetchPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PGEligibilityFetchPaymentMethods`: []EligibilityPaymentMethodsEntity
	fmt.Fprintf(os.Stdout, "Response from `EligibilityAPI.PGEligibilityFetchPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGEligibilityFetchPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2025-01-01&quot;]
 **eligibilityFetchPaymentMethodsRequest** | [**EligibilityFetchPaymentMethodsRequest**](EligibilityFetchPaymentMethodsRequest.md) | Request Body to get eligible payment methods for an account and order | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**[]EligibilityPaymentMethodsEntity**](EligibilityPaymentMethodsEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

