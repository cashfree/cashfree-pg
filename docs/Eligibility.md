# Eligibility

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGEligibilityFetchCardlessEMI**](Eligibility.md#PGEligibilityFetchCardlessEMI) | **Post** /eligibility/cardlessemi | Get Eligible Cardless EMI
[**PGEligibilityFetchOffers**](Eligibility.md#PGEligibilityFetchOffers) | **Post** /eligibility/offers | Get Eligible Offers
[**PGEligibilityFetchPaylater**](Eligibility.md#PGEligibilityFetchPaylater) | **Post** /eligibility/paylater | Get Eligible Paylater
[**PGEligibilityFetchPaymentMethods**](Eligibility.md#PGEligibilityFetchPaymentMethods) | **Post** /eligibility/payment_methods | Get Eligible Payment Methods



## PGEligibilityFetchCardlessEMI

> PGEligibilityFetchCardlessEMI(xApiVersion *string, eligibilityFetchCardlessEMIRequest *EligibilityFetchCardlessEMIRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]EligibilityCardlessEMIEntity, *http.Response, error)

Get Eligible Cardless EMI ([Docs](https://docs.cashfree.com/reference/pgeligibilityfetchcardlessemi))

### Example

```go
version := "2022-09-01"
var amount float32 = 1000.0
eligibibleCardlessEmiRequest := cashfree.EligibilityFetchCardlessEMIRequest{
	Queries: cashfree.CardlessEMIQueries{
		Amount: &amount,
		CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
			CustomerPhone: "8908908901",
		},
	},
}
eligibilityEntity, httpResponse, err := cashfree.PGEligibilityFetchCardlessEMI(&version, &eligibibleCardlessEmiRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(eligibilityEntity)
}
```

###  Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**eligibilityFetchCardlessEMIRequest** | [**EligibilityFetchCardlessEMIRequest***](Eligibility.md#EligibilityFetchCardlessEMIRequest) | Request Body to get eligible cardless emi options for a customer and order | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree |

### Response

```json
[
  {
    "eligibility": true,
    "entity_type": "cardlessemi",
    "entity_value": "idfc",
    "entity_details": {
      "payment_method": "idfc",
      "emi_plans": [
        {
          "tenure": 1,
          "interest_rate": 10,
          "currency": "INR",
          "emi": 400,
          "total_interest": 10,
          "total_amount": 40
        }
      ]
    }
  }
]
```

## PGEligibilityFetchOffers

> PGEligibilityFetchOffers(xApiVersion *string, eligibilityFetchOffersRequest *EligibilityFetchOffersRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]EligibilityOfferEntity, *http.Response, error)

Get Eligible Offers ([Docs](https://docs.cashfree.com/reference/pgeligibilityfetchoffers))

### Example

```go
version := "2022-09-01"
var amount float32 = 1000.0
eligibibleOffersRequest := cashfree.EligibilityFetchOffersRequest{
	Queries: cashfree.OfferQueries{
		Amount: &amount,
	},
}
eligibilityEntity, httpResponse, err := cashfree.PGEligibilityFetchOffers(&version, &eligibibleOffersRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(eligibilityEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**eligibilityFetchOffersRequest** | [**EligibilityFetchOffersRequest***](EligibilityFetchOffersRequest.md) | Request Body to get eligible offers for a customer and order | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 

### Response

```json
[
  {
    "offer_id": "d2b430fb-1afe-455a-af31-66d00377b29a",
    "offer_status": "active",
    "offer_meta": {
      "offer_title": "some title",
      "offer_description": "some offer description",
      "offer_code": "CFTESTOFFER",
      "offer_start_time": "2023-03-21T08:09:51Z",
      "offer_end_time": "2023-03-29T08:09:51Z"
    },
    "offer_tnc": {
      "offer_tnc_type": "text",
      "offer_tnc_value": "TnC for the Offer."
    },
    "offer_details": {
      "offer_type": "DISCOUNT_AND_CASHBACK",
      "discount_details": {
        "discount_type": "flat",
        "discount_value": "10",
        "max_discount_amount": "10"
      },
      "cashback_details": {
        "cashback_type": "percentage",
        "cashback_value": "20",
        "max_cashback_amount": "150"
      }
    },
    "offer_validations": {
      "min_amount": 10,
      "payment_method": {
        "wallet": {
          "issuer": "paytm"
        }
      },
      "max_allowed": 2
    }
  }
]
```


## PGEligibilityFetchPaylater

> PGEligibilityFetchPaylater(xApiVersion *string, eligibilityFetchPaylaterRequest *EligibilityFetchPaylaterRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]EligibilityPaylaterEntity, *http.Response, error)

Get Eligible Paylater ([Docs](https://docs.cashfree.com/reference/pgeligibilityfetchpaylater))



### Example

```go
version := "2022-09-01"
var amount float32 = 1000.0
eligibiblePayLaterRequest := cashfree.EligibilityFetchPaylaterRequest{
	Queries: cashfree.CardlessEMIQueries{
		Amount: &amount,
		CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
			CustomerPhone: "8908908901",
		},
	},
}
eligibilityEntity, httpResponse, err := cashfree.PGEligibilityFetchPaylater(&version, &eligibiblePayLaterRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(eligibilityEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**eligibilityFetchPaylaterRequest** | [**EligibilityFetchPaylaterRequest***](EligibilityFetchPaylaterRequest.md) | Request Body to get eligible paylater options for a customer and order | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 

### Response

```json
[
  {
    "eligibility": true,
    "entity_type": "paylater",
    "entity_value": "olapostpaid",
    "entity_details": {
      "payment_method": "olapostpaid"
    }
  }
]
```


## PGEligibilityFetchPaymentMethods

> PGEligibilityFetchPaymentMethods(xApiVersion *string, eligibilityFetchPaymentMethodsRequest *EligibilityFetchPaymentMethodsRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]EligibilityPaymentMethodsEntity, *http.Response, error)

Get Eligible Payment Methods ([Docs](https://docs.cashfree.com/reference/pgeligibilityfetchpaymentmethods))



### Example

```go
version := "2022-09-01"
var amount float32 = 1000.0
eligibiblePaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
	Queries: cashfree.PaymentMethodsQueries{
		Amount: &amount,
	},
}
eligibilityEntity, httpResponse, err := cashfree.PGEligibilityFetchPaymentMethods(&version, &eligibiblePaymentMethodsRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(eligibilityEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**eligibilityFetchPaymentMethodsRequest** | [**EligibilityFetchPaymentMethodsRequest***](EligibilityFetchPaymentMethodsRequest.md) | Request Body to get eligible payment methods for an account and order | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 

### Response
```json
[
  {
    "eligibility": true,
    "entity_type": "payment_methods",
    "entity_value": "netbanking",
    "entity_details": {
      "payment_method_details": [
        {
          "nick": "motak_kahindra_bank",
          "display": "Motak Kahindra Bank",
          "eligibility": true,
          "code": 3032
        },
        {
          "nick": "bank_of_india",
          "display": "Bank Of India",
          "eligibility": true,
          "code": 3031
        }
      ]
    }
  }
]
```

