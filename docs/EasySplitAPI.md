# \EasySplitAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGESCreateOnDemandTransfer**](EasySplitAPI.md#PGESCreateOnDemandTransfer) | **Post** /easy-split/vendors/{vendor_id}/transfer | Create On Demand Transfer
[**PGESCreateVendors**](EasySplitAPI.md#PGESCreateVendors) | **Post** /easy-split/vendors | Create vendor
[**PGESDownloadVendorsDocs**](EasySplitAPI.md#PGESDownloadVendorsDocs) | **Get** /easy-split/vendor-docs/{vendor_id}/download/{doc_type} | Download Vendor Documents
[**PGESFetchVendors**](EasySplitAPI.md#PGESFetchVendors) | **Get** /easy-split/vendors/{vendor_id} | Get Vendor All Details
[**PGESGetVendorBalance**](EasySplitAPI.md#PGESGetVendorBalance) | **Get** /easy-split/vendors/{vendor_id}/balances | Get On Demand Balance
[**PGESGetVendorBalanceTransferCharges**](EasySplitAPI.md#PGESGetVendorBalanceTransferCharges) | **Get** /easy-split/amount/{amount}/charges | Get Vendor Balance Transfer Charges
[**PGESGetVendorsDocs**](EasySplitAPI.md#PGESGetVendorsDocs) | **Get** /easy-split/vendor-docs/{vendor_id} | Get Vendor All Documents Status
[**PGESOrderRecon**](EasySplitAPI.md#PGESOrderRecon) | **Post** /split/order/vendor/recon | Get Split and Settlement Details by OrderID v2.0
[**PGESUpdateVendors**](EasySplitAPI.md#PGESUpdateVendors) | **Patch** /easy-split/vendors/{vendor_id} | Update vendor Details
[**PGESUploadVendorsDocs**](EasySplitAPI.md#PGESUploadVendorsDocs) | **Post** /easy-split/vendor-docs/{vendor_id} | Upload Vendor Docs
[**PGOrderSplitAfterPayment**](EasySplitAPI.md#PGOrderSplitAfterPayment) | **Post** /easy-split/orders/{order_id}/split | Split After Payment
[**PGOrderStaticSplit**](EasySplitAPI.md#PGOrderStaticSplit) | **Post** /easy-split/static-split | Create Static Split Configuration



## PGESCreateOnDemandTransfer

> AdjustVendorBalanceResponse PGESCreateOnDemandTransfer(ctx, vendorId).XApiVersion(xApiVersion).AdjustVendorBalanceRequest(adjustVendorBalanceRequest).Execute()

Create On Demand Transfer



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
    vendorId := "your-vendor-id" 
    adjustVendorBalanceRequest := *cashfree.NewAdjustVendorBalanceRequest("TransferFrom_example", "TransferType_example", float32(123)) 

    resp, r, err := cashfree.PGESCreateOnDemandTransfer(&xApiVersion, &vendorId, &adjustVendorBalanceRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESCreateOnDemandTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESCreateOnDemandTransfer`: AdjustVendorBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESCreateOnDemandTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESCreateOnDemandTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **adjustVendorBalanceRequest** | [**AdjustVendorBalanceRequest**](AdjustVendorBalanceRequest.md) | Adjust Vendor Balance Request Body. | 

### Return type

[**AdjustVendorBalanceResponse**](AdjustVendorBalanceResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESCreateVendors

> CreateVendorResponse PGESCreateVendors(ctx).XApiVersion(xApiVersion).CreateVendorRequest(createVendorRequest).Execute()

Create vendor



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
    createVendorRequest := *cashfree.NewCreateVendorRequest("VendorId_example", "Status_example", "Name_example", "Email_example", "Phone_example", []cashfree.KycDetails{*cashfree.NewKycDetails()}) 

    resp, r, err := cashfree.PGESCreateVendors(&xApiVersion, &createVendorRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESCreateVendors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESCreateVendors`: CreateVendorResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESCreateVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGESCreateVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **createVendorRequest** | [**CreateVendorRequest**](CreateVendorRequest.md) | Create Vendor Request Body. | 

### Return type

[**CreateVendorResponse**](CreateVendorResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESDownloadVendorsDocs

> VendorDocumentDownloadResponse PGESDownloadVendorsDocs(ctx, docType, vendorId).Execute()

Download Vendor Documents



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

    docType := "docType_example" 
    vendorId := "your-vendor-id" 

    resp, r, err := cashfree.PGESDownloadVendorsDocs(&docType, &vendorId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESDownloadVendorsDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESDownloadVendorsDocs`: VendorDocumentDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESDownloadVendorsDocs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**docType** | **string** | Mention the document type that has to be downloaded. Only an uploaded document can be downloaded. | 
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESDownloadVendorsDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VendorDocumentDownloadResponse**](VendorDocumentDownloadResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESFetchVendors

> VendorEntity PGESFetchVendors(ctx, vendorId).XApiVersion(xApiVersion).Execute()

Get Vendor All Details



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
    vendorId := "your-vendor-id" 

    resp, r, err := cashfree.PGESFetchVendors(&xApiVersion, &vendorId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESFetchVendors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESFetchVendors`: VendorEntity
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESFetchVendors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESFetchVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


### Return type

[**VendorEntity**](VendorEntity.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESGetVendorBalance

> VendorBalance PGESGetVendorBalance(ctx, vendorId).XApiVersion(xApiVersion).Execute()

Get On Demand Balance



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
    vendorId := "your-vendor-id" 

    resp, r, err := cashfree.PGESGetVendorBalance(&xApiVersion, &vendorId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESGetVendorBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESGetVendorBalance`: VendorBalance
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESGetVendorBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESGetVendorBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


### Return type

[**VendorBalance**](VendorBalance.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESGetVendorBalanceTransferCharges

> VendorBalanceTransferCharges PGESGetVendorBalanceTransferCharges(ctx, amount).XApiVersion(xApiVersion).RateType(rateType).Execute()

Get Vendor Balance Transfer Charges



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
    amount := float32(1000) 
    rateType := "VENDOR_ON_DEMAND" 

    resp, r, err := cashfree.PGESGetVendorBalanceTransferCharges(&xApiVersion, &amount, &rateType, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESGetVendorBalanceTransferCharges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESGetVendorBalanceTransferCharges`: VendorBalanceTransferCharges
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESGetVendorBalanceTransferCharges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amount** | **float32** | Specify the amount for which you want to view the service charges and service taxes in the response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESGetVendorBalanceTransferChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **rateType** | **string** | Mention the type of rate for which you want to check the charges. Possible value: VENDOR_ON_DEMAND | 

### Return type

[**VendorBalanceTransferCharges**](VendorBalanceTransferCharges.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESGetVendorsDocs

> VendorDocumentsResponse PGESGetVendorsDocs(ctx, vendorId).XApiVersion(xApiVersion).Execute()

Get Vendor All Documents Status



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
    vendorId := "your-vendor-id" 

    resp, r, err := cashfree.PGESGetVendorsDocs(&xApiVersion, &vendorId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESGetVendorsDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESGetVendorsDocs`: VendorDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESGetVendorsDocs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESGetVendorsDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]


### Return type

[**VendorDocumentsResponse**](VendorDocumentsResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESOrderRecon

> ESOrderReconResponse PGESOrderRecon(ctx).XApiVersion(xApiVersion).ESOrderReconRequest(eSOrderReconRequest).Execute()

Get Split and Settlement Details by OrderID v2.0



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
    eSOrderReconRequest := *cashfree.NewESOrderReconRequest(*cashfree.NewESOrderReconRequestFilters(), *cashfree.NewESOrderReconRequestPagination()) 

    resp, r, err := cashfree.PGESOrderRecon(&xApiVersion, &eSOrderReconRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESOrderRecon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESOrderRecon`: ESOrderReconResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESOrderRecon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGESOrderReconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **eSOrderReconRequest** | [**ESOrderReconRequest**](ESOrderReconRequest.md) | Get Split and Settlement Details by OrderID v2.0 | 

### Return type

[**ESOrderReconResponse**](ESOrderReconResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESUpdateVendors

> UpdateVendorResponse PGESUpdateVendors(ctx, vendorId).XApiVersion(xApiVersion).UpdateVendorRequest(updateVendorRequest).Execute()

Update vendor Details



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
    vendorId := "your-vendor-id" 
    updateVendorRequest := *cashfree.NewUpdateVendorRequest("Status_example", "Name_example", "Email_example", "Phone_example", float32(123), []cashfree.KycDetails{*cashfree.NewKycDetails()}) 

    resp, r, err := cashfree.PGESUpdateVendors(&xApiVersion, &vendorId, &updateVendorRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESUpdateVendors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESUpdateVendors`: UpdateVendorResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESUpdateVendors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESUpdateVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **updateVendorRequest** | [**UpdateVendorRequest**](UpdateVendorRequest.md) | Create Vendor Request Body. | 

### Return type

[**UpdateVendorResponse**](UpdateVendorResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGESUploadVendorsDocs

> UploadVendorDocumentsResponse PGESUploadVendorsDocs(ctx, vendorId).XApiVersion(xApiVersion).UploadVendorDocsRequest(uploadVendorDocsRequest).Execute()

Upload Vendor Docs



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
    vendorId := "your-vendor-id" 
    uploadVendorDocsRequest := TODO 

    resp, r, err := cashfree.PGESUploadVendorsDocs(&xApiVersion, &vendorId, &uploadVendorDocsRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGESUploadVendorsDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGESUploadVendorsDocs`: UploadVendorDocumentsResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGESUploadVendorsDocs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | **string** | The id which uniquely identifies your vendor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGESUploadVendorsDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **uploadVendorDocsRequest** | [**UploadVendorDocsRequest**](UploadVendorDocsRequest.md) | Upload Vendor Docs Request Body. | 

### Return type

[**UploadVendorDocumentsResponse**](UploadVendorDocumentsResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGOrderSplitAfterPayment

> SplitAfterPaymentResponse PGOrderSplitAfterPayment(ctx, orderId).XApiVersion(xApiVersion).SplitAfterPaymentRequest(splitAfterPaymentRequest).Execute()

Split After Payment



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
    orderId := "your-order-id" 
    splitAfterPaymentRequest := *cashfree.NewSplitAfterPaymentRequest([]cashfree.SplitAfterPaymentRequestSplitInner{*cashfree.NewSplitAfterPaymentRequestSplitInner()}) 

    resp, r, err := cashfree.PGOrderSplitAfterPayment(&xApiVersion, &orderId, &splitAfterPaymentRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGOrderSplitAfterPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGOrderSplitAfterPayment`: SplitAfterPaymentResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGOrderSplitAfterPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id which uniquely identifies your order | 

### Other Parameters

Other parameters are passed through a pointer to a apiPGOrderSplitAfterPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]

 **splitAfterPaymentRequest** | [**SplitAfterPaymentRequest**](SplitAfterPaymentRequest.md) | Request Body to Create Split for an order. | 

### Return type

[**SplitAfterPaymentResponse**](SplitAfterPaymentResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PGOrderStaticSplit

> StaticSplitResponse PGOrderStaticSplit(ctx).XApiVersion(xApiVersion).StaticSplitRequest(staticSplitRequest).Execute()

Create Static Split Configuration



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
    staticSplitRequest := *cashfree.NewStaticSplitRequest(false, "ProductType_example", []cashfree.StaticSplitRequestSchemeInner{*cashfree.NewStaticSplitRequestSchemeInner()}) 

    resp, r, err := cashfree.PGOrderStaticSplit(&xApiVersion, &staticSplitRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PGOrderStaticSplit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PGOrderStaticSplit`: StaticSplitResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PGOrderStaticSplit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPGOrderStaticSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2023-08-01&quot;]
 **staticSplitRequest** | [**StaticSplitRequest**](StaticSplitRequest.md) | Static Split | 

### Return type

[**StaticSplitResponse**](StaticSplitResponse.md)

### Authorization

[XPartnerAPIKey](../README.md#XPartnerAPIKey), [XClientSecret](../README.md#XClientSecret), [XPartnerMerchantID](../README.md#XPartnerMerchantID), [XClientID](../README.md#XClientID), [XClientSignatureHeader](../README.md#XClientSignatureHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

