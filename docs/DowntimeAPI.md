# \DowntimeAPI

All URIs are relative to *https://sandbox.cashfree.com/pg*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchActiveEcosystemDowntimes**](DowntimeAPI.md#FetchActiveEcosystemDowntimes) | **Get** /incident | Fetch All Downtimes



## FetchActiveEcosystemDowntimes

> FetchActiveEcosystemDowntimes200Response FetchActiveEcosystemDowntimes(ctx).XApiVersion(xApiVersion).IncidentId(incidentId).IncidentStatus(incidentStatus).IncidentImpact(incidentImpact).IncidentType(incidentType).IncidentStartTime(incidentStartTime).IncidentEndTime(incidentEndTime).PaymentMethod(paymentMethod).Execute()

Fetch All Downtimes



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
    incidentId := "your-incident-id" 
    incidentStatus := []string{"IncidentStatus_example"} 
    incidentImpact := []string{"IncidentImpact_example"} 
    incidentType := []string{"IncidentType_example"} 
    incidentStartTime := time.Now() 
    incidentEndTime := time.Now() 
    paymentMethod := []string{"PaymentMethod_example"} 

    resp, r, err := cashfree.FetchActiveEcosystemDowntimes(&xApiVersion, &incidentId, &incidentStatus, &incidentImpact, &incidentType, &incidentStartTime, &incidentEndTime, &paymentMethod, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.FetchActiveEcosystemDowntimes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchActiveEcosystemDowntimes`: FetchActiveEcosystemDowntimes200Response
    fmt.Fprintf(os.Stdout, "Response from `cashfree.FetchActiveEcosystemDowntimes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchActiveEcosystemDowntimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD. | [default to &quot;2026-01-01&quot;]
 **incidentId** | **string** | Valid incident id for fetching incident details. | 
 **incidentStatus** | **[]string** | Filter incidents by status. Possible values: ACTIVE, UPCOMING, RESOLVED. | 
 **incidentImpact** | **[]string** | Filter incidents by impact level. Possible values: HIGH, MEDIUM, LOW. | 
 **incidentType** | **[]string** | Filter incidents by type. Possible values: SCHEDULED, UNSCHEDULED. | 
 **incidentStartTime** | **time.Time** | Filter incidents by start time. Format: YYYY-MM-DD HH:MM:SS. | 
 **incidentEndTime** | **time.Time** | Filter incidents by end time. Format: YYYY-MM-DD HH:MM:SS. | 
 **paymentMethod** | **[]string** | Filter incidents by payment method. Possible values: UPI, CARD, NET_BANKING, WALLET. | 

### Return type

[**FetchActiveEcosystemDowntimes200Response**](FetchActiveEcosystemDowntimes200Response.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

