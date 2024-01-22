# \OffersAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGCreateOffer**](OffersAPI.md#PGCreateOffer) | **Post** /offers | Create Offer
[**PGFetchOffer**](OffersAPI.md#PGFetchOffer) | **Get** /offers/{offer_id} | Get Offer by ID



## PGCreateOffer

> OfferEntity PGCreateOffer(ctx).XApiVersion(xApiVersion).CreateOfferRequest(createOfferRequest).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Create Offer



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
    createOfferRequest := *cashfree.NewCreateOfferRequest(*cashfree.NewOfferMeta("Test Offer", "Lorem ipsum dolor sit amet, consectetur adipiscing elit", "CFTESTOFFER", "2023-03-21T08:09:51Z", "2023-03-29T08:09:51Z"), *cashfree.NewOfferTnc("text", "Lorem ipsum dolor sit amet, consectetur adipiscing elit"), *cashfree.NewOfferDetails("DISCOUNT_AND_CASHBACK"), *cashfree.NewOfferValidations(cashfree.OfferValidations_payment_method{OfferAll: cashfree.NewOfferAll(map[string]interface{}(123))})) 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGCreateOffer(&xApiVersion, &createOfferRequest, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGCreateOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGCreateOffer`: OfferEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGCreateOffer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGCreateOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **createOfferRequest** | [**CreateOfferRequest**](CreateOfferRequest.md) | Request body to create an offer at Cashfree | 
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**OfferEntity**](OfferEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGFetchOffer

> OfferEntity PGFetchOffer(ctx, offerId).XApiVersion(xApiVersion).XRequestId(xRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Offer by ID



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
    offerId := "d2b430fb-1afe-455a-af31-66d00377b29a" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    xIdempotencyKey := "47bf8872-46fe-11ee-be56-0242ac120002" 

    resp, r, err := cashfree.PGFetchOffer(&xApiVersion, &offerId, &xRequestId, &xIdempotencyKey, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGFetchOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGFetchOffer`: OfferEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGFetchOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**offerId** | **string** | The offer ID for which you want to view the offer details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGFetchOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 **xIdempotencyKey** | **string** | An idempotency key is a unique identifier you include with your API call. If the request fails or times out, you can safely retry it using the same key to avoid duplicate actions.   | 

### Return type

[**OfferEntity**](OfferEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

