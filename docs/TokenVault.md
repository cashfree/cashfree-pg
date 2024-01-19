# TokenVault

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGCustomerDeleteInstrument**](TokenVault.md#PGCustomerDeleteInstrument) | **Delete** /customers/{customer_id}/instruments/{instrument_id} | Delete Saved Card Instrument
[**PGCustomerFetchInstrument**](TokenVault.md#PGCustomerFetchInstrument) | **Get** /customers/{customer_id}/instruments/{instrument_id} | Fetch Specific Saved Card Instrument
[**PGCustomerFetchInstruments**](TokenVault.md#PGCustomerFetchInstruments) | **Get** /customers/{customer_id}/instruments | Fetch All Saved Card Instrument
[**PGCustomerInstrumentsFetchCryptogram**](TokenVault.md#PGCustomerInstrumentsFetchCryptogram) | **Get** /customers/{customer_id}/instruments/{instrument_id}/cryptogram | Fetch cryptogram for a saved card instrument



## PGCustomerDeleteInstrument

> PGCustomerDeleteInstrument(xApiVersion *string, customerId string, instrumentId string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*InstrumentEntity, *http.Response, error)

Delete Saved Card Instrument ([Docs](https://docs.cashfree.com/reference/pgcustomerdeleteinstrument))

### Example

```go
version := "2022-09-01"
instrumentEntity, httpResponse, err := cashfree.PGCustomerDeleteInstrument(&version, "<customer_id>", "<instrument_id>", nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(instrumentEntity)
}
```

### Parameters
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**customerId** | **string*** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string*** | The instrument_id which needs to be deleted | 
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response

```json
{
  "customer_id": "siddhesh_desai",
  "afa_reference": "740324562",
  "instrument_id": "54deabb4-ba45-4a60-9e6a-9c016fe7ab10",
  "instrument_type": "card",
  "instrument_uid": "0d8f70838cc5af8b1cd2bc0fe71278551fd3f1101e40020d89ad22ceba4f933c",
  "instrument_display": "xxxxxxxxxxxx4375",
  "instrument_status": "ACTIVE",
  "created_at": "2021-11-11 16:57:57",
  "instrument_meta": {
    "card_network": "VISA",
    "card_bank_name": "HDFC Bank Limited",
    "card_country": "IN",
    "card_type": "DEBIT_CARD",
    "card_token_details": {
      "par": "somepar",
      "expiry_month": "12",
      "expiry_year": "23"
    }
  }
}
```

## PGCustomerFetchInstrument

> PGCustomerFetchInstrument(xApiVersion *string, customerId string, instrumentId string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*InstrumentEntity, *http.Response, error)

Fetch Specific Saved Card Instrument ([Docs](https://docs.cashfree.com/reference/pgcustomerfetchinstrument))



### Example

```go
version := "2022-09-01"
instrumentEntity, httpResponse, err := cashfree.PGCustomerFetchInstrument(&version, "<customer_id>", "<instrument_id>", nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(instrumentEntity)
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**customerId** | **string** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string** | The instrument_id of the saved instrument which needs to be queried | 
**xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response
```json
{
  "customer_id": "siddhesh_desai",
  "afa_reference": "740324562",
  "instrument_id": "54deabb4-ba45-4a60-9e6a-9c016fe7ab10",
  "instrument_type": "card",
  "instrument_uid": "0d8f70838cc5af8b1cd2bc0fe71278551fd3f1101e40020d89ad22ceba4f933c",
  "instrument_display": "xxxxxxxxxxxx4375",
  "instrument_status": "ACTIVE",
  "created_at": "2021-11-11 16:57:57",
  "instrument_meta": {
    "card_network": "VISA",
    "card_bank_name": "HDFC Bank Limited",
    "card_country": "IN",
    "card_type": "DEBIT_CARD",
    "card_token_details": {
      "par": "somepar",
      "expiry_month": "12",
      "expiry_year": "23"
    }
  }
}
```

## PGCustomerFetchInstruments

> PGCustomerFetchInstruments(xApiVersion *string, customerId string, instrumentType *string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]InstrumentEntity, *http.Response, error)

Fetch All Saved Card Instrument ([Docs](https://docs.cashfree.com/reference/pgcustomerfetchinstruments))

### Example

```go
version := "2022-09-01"
instrumentType := "card"
instrumentEntity, httpResponse, err := cashfree.PGCustomerFetchInstruments(&version, "<customer_id>", &instrumentType, nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(instrumentEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**customerId** | **string*** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**instrumentType** | **string*** | Payment mode or type of saved instrument  | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response
```json
[
  {
  "customer_id": "siddhesh_desai",
  "afa_reference": "740324562",
  "instrument_id": "54deabb4-ba45-4a60-9e6a-9c016fe7ab10",
  "instrument_type": "card",
  "instrument_uid": "0d8f70838cc5af8b1cd2bc0fe71278551fd3f1101e40020d89ad22ceba4f933c",
  "instrument_display": "xxxxxxxxxxxx4375",
  "instrument_status": "ACTIVE",
  "created_at": "2021-11-11 16:57:57",
  "instrument_meta": {
    "card_network": "VISA",
    "card_bank_name": "HDFC Bank Limited",
    "card_country": "IN",
    "card_type": "DEBIT_CARD",
    "card_token_details": {
      "par": "somepar",
      "expiry_month": "12",
      "expiry_year": "23"
    }
  }
}
]
```


## PGCustomerInstrumentsFetchCryptogram

> PGCustomerInstrumentsFetchCryptogram(xApiVersion *string, customerId string, instrumentId string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*CryptogramEntity, *http.Response, error)

Fetch cryptogram for a saved card instrument ([Docs](https://docs.cashfree.com/reference/pgcustomerinstrumentsfetchcryptogram))



### Example

```go
version := "2022-09-01"
cryptogramEntity, httpResponse, err := cashfree.PGCustomerInstrumentsFetchCryptogram(&version, "<customer_id>", "<instrument_id>", nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(cryptogramEntity)
}
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**customerId** | **string*** | Your Customer ID that you had sent during create order API &#x60;POST/orders&#x60; | 
**instrumentId** | **string*** | The instrument_id of the saved card instrument which needs to be queried | 
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response

```json
{
  "instrument_id": "54deabb4-ba45-4a60-9e6a-9c016fe7ab10",
  "token_requestor_id": "22457512314",
  "card_number": "4491365621601472",
  "card_expiry_mm": "06",
  "card_expiry_yy": "2025",
  "cryptogram": "AQBBBBBBZatIlaIAmWKSghwBBBB=",
  "card_display": "1234"
}
```

