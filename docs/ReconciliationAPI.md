# \ReconciliationAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGESReconVendor**](ReconciliationAPI.md#PGESReconVendor) | **Post** /recon/vendor | Vendor Reconciliation API



## PGESReconVendor

> VendorRecon200Response PGESReconVendor(ctx).XApiVersion(xApiVersion).VendorReconRequest(vendorReconRequest).ContentType(contentType).Execute()

Vendor Reconciliation API



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
    vendorReconRequest := *cashfree.NewVendorReconRequest(*cashfree.NewVendorReconRequestPagination(int32(123)), *cashfree.NewVendorReconRequestFilters()) 
    contentType := "application/json" 

    resp, r, err := cashfree.PGESReconVendor(&xApiVersion, &vendorReconRequest, &contentType, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESReconVendor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESReconVendor`: VendorRecon200Response
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESReconVendor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGESReconVendorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD. | [default to &quot;2026-01-01&quot;]
 **vendorReconRequest** | [**VendorReconRequest**](VendorReconRequest.md) | Vendor Recon Request body. | 
 **contentType** | **string** | application/json. | 

### Return type

[**VendorRecon200Response**](VendorRecon200Response.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

