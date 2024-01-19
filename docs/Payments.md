# Payments

Method | HTTP request | Description
------------- | ------------- | -------------
[**PGOrderFetchPayment**](Payments.md#PGOrderFetchPayment) | **Get** /orders/{order_id}/payments/{cf_payment_id} | Get Payment by ID
[**PGOrderFetchPayments**](Payments.md#PGOrderFetchPayments) | **Get** /orders/{order_id}/payments | Get Payments for an Order
[**PGPayOrder**](Payments.md#PGPayOrder) | **Post** /orders/sessions | Order Pay
[**PGAuthorizeOrder**](Payments.md#PGAuthorizeOrder) | **Post** /orders/{order_id}/authorization | Preauthorization
[**PGOrderAuthenticatePayment**](Payments.md#PGOrderAuthenticatePayment) | **Post** /orders/pay/authenticate/{cf_payment_id} | Submit or Resend OTP


## PGOrderFetchPayment

> PGOrderFetchPayment(xApiVersion *string, orderId string, cfPaymentId string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*PaymentEntity, *http.Response, error)

Get Payment by ID ([Docs](https://docs.cashfree.com/reference/pgorderfetchpayment))



### Example

```go
version := "2022-09-01"
paymentsEntity, httpResponse, err := cashfree.PGOrderFetchPayment(&version, "order_342bAHtHiGpa2XePHbvRdu22S7p8U", "14910259396", nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(paymentsEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**orderId** | **string*** | The id which uniquely identifies your order | 
**cfPaymentId** | **string*** | The Cashfree payment or transaction ID. | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response
```json
{
  "cf_payment_id": 12376123,
  "order_id": "order_8123",
  "entity": "payment",
  "payment_currency": "INR",
  "error_details": null,
  "order_amount": 10.01,
  "is_captured": true,
  "payment_group": "upi",
  "authorization": {
    "action": "CAPTURE",
    "status": "PENDING",
    "captured_amount": 100,
    "start_time": "2022-02-09T18:04:34+05:30",
    "end_time": "2022-02-19T18:04:34+05:30",
    "approve_by": "2022-02-09T18:04:34+05:30",
    "action_reference": "6595231908096894505959",
    "action_time": "2022-08-03T16:09:51"
  },
  "payment_method": {
    "upi": {
      "channel": "collect",
      "upi_id": "rohit@xcxcx"
    }
  },
  "payment_amount": 10.01,
  "payment_time": "2021-07-23T12:15:06+05:30",
  "payment_completion_time": "2021-07-23T12:18:59+05:30",
  "payment_status": "SUCCESS",
  "payment_message": "Transaction successful",
  "bank_reference": "P78112898712",
  "auth_id": "A898101"
}
```



## PGOrderFetchPayments

> PGOrderFetchPayments(xApiVersion *string, orderId string,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) ([]PaymentEntity, *http.Response, error)

Get Payments for an Order ([Docs](https://docs.cashfree.com/reference/pgorderfetchpayments))



### Example

```javascript
version := "2022-09-01"
paymentsEntity, httpResponse, err := cashfree.PGOrderFetchPayments(&version, "order_342bAHtHiGpa2XePHbvRdu22S7p8U", nil, nil, nil)
if err != nil {
	fmt.Print(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(paymentsEntity)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**xApiVersion** | **string*** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**orderId** | **string*** | The id which uniquely identifies your order | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response
```json
[
  {
    "cf_payment_id": 12376123,
    "order_id": "order_8123",
    "entity": "payment",
    "payment_currency": "INR",
    "error_details": null,
    "order_amount": 10.01,
    "is_captured": true,
    "payment_group": "upi",
    "authorization": {
      "action": "CAPTURE",
      "status": "PENDING",
      "captured_amount": 100,
      "start_time": "2022-02-09T18:04:34+05:30",
      "end_time": "2022-02-19T18:04:34+05:30",
      "approve_by": "2022-02-09T18:04:34+05:30",
      "action_reference": "6595231908096894505959",
      "action_time": "2022-08-03T16:09:51"
    },
    "payment_method": {
      "upi": {
        "channel": "collect",
        "upi_id": "rohit@xcxcx"
      }
    },
    "payment_amount": 10.01,
    "payment_time": "2021-07-23T12:15:06+05:30",
    "payment_completion_time": "2021-07-23T12:18:59+05:30",
    "payment_status": "SUCCESS",
    "payment_message": "Transaction successful",
    "bank_reference": "P78112898712",
    "auth_id": "A898101"
  },
  {
    "cf_payment_id": 12376124,
    "order_id": "order_8123",
    "entity": "payment",
    "payment_currency": "INR",
    "error_details": {
      "error_code": "TRANSACTION_DECLINED",
      "error_description": "issuer bank or payment service provider declined the transaction",
      "error_reason": "auth_declined",
      "error_source": "customer"
    },
    "order_amount": 10.01,
    "is_captured": true,
    "payment_group": "credit_card",
    "authorization": null,
    "payment_method": {
      "card": {
        "channel": "link",
        "card_number": "xxxxxx1111"
      }
    },
    "payment_amount": 10.01,
    "payment_time": "2021-07-23T12:15:06+05:30",
    "payment_completion_time": "2021-07-23T12:18:59+05:30",
    "payment_status": "FAILED",
    "payment_message": "Transaction failed",
    "bank_reference": "P78112898712",
    "auth_id": "A898101"
  }
]
```

## PGPayOrder

> PGPayOrder(xApiVersion *string, payOrderRequest *PayOrderRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*PayOrderEntity, *http.Response, error)

Order Pay ([Docs](https://docs.cashfree.com/reference/pgpayorder))

## Example
#### Card
```go
channel := "link"
cardNumber := "4111111111111111"
cardHolderName := "Test"
cardMonth := "12"
cardYear := "25"
cardCvv := "123"
version := "2022-09-01"

cardPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		CardPaymentMethod: &cashfree.CardPaymentMethod{
			Card: cashfree.Card{
				Channel:        channel,
				CardNumber:     &cardNumber,
				CardExpiryMm:   &cardMonth,
				CardHolderName: &cardHolderName,
				CardCvv:        &cardCvv,
				CardExpiryYy:   &cardYear,
			},
		},
	},
}

orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &cardPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payOrderRequest** | **PayOrderRequest*** | Request body to create a transaction at cashfree using &#x60;payment_session_id&#x60; | 
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 

#### Netbanking
```go
channel := "link"
var netbankingCode int32 = 3010
version := "2022-09-01"
netbankingPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		NetBankingPaymentMethod: &cashfree.NetBankingPaymentMethod{
			Netbanking: cashfree.Netbanking{
				Channel:            channel,
				NetbankingBankCode: &netbankingCode,
			},
		},
	},
}
orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &netbankingPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### UPI
```go
channel := "collect"
version := "2022-09-01"
upi_id := "testsuccess@gocash"
upiPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		UPIPaymentMethod: &cashfree.UPIPaymentMethod{
			Upi: cashfree.Upi{
				Channel: channel,
				UpiId:   &upi_id,
			},
		},
	},
}
orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &upiPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### UPI
```go
channel := "link"
version := "2022-09-01"
upiPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		UPIPaymentMethod: &cashfree.UPIPaymentMethod{
			Upi: cashfree.Upi{
				Channel: channel,
			},
		},
	},
}
orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &upiPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### QR
```go
channel := "qrcode"
version := "2022-09-01"
upiPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		UPIPaymentMethod: &cashfree.UPIPaymentMethod{
			Upi: cashfree.Upi{
				Channel: channel,
			},
		},
	},
}
orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &upiPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### Wallet
```go
channel := "link"
version := "2022-09-01"
walletPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		AppPaymentMethod: &cashfree.AppPaymentMethod{
			App: cashfree.App{
				Channel:  channel,
				Provider: "phonepe",
				Phone:    "9999999999",
			},
		},
	},
}
orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &walletPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### Card (with saving it)
```go
channel := "link"
cardNumber := "4111111111111111"
cardHolderName := "Test"
cardMonth := "12"
cardYear := "25"
cardCvv := "123"
version := "2022-09-01"
flag := true

cardPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		CardPaymentMethod: &cashfree.CardPaymentMethod{
			Card: cashfree.Card{
				Channel:        channel,
				CardNumber:     &cardNumber,
				CardExpiryMm:   &cardMonth,
				CardHolderName: &cardHolderName,
				CardCvv:        &cardCvv,
				CardExpiryYy:   &cardYear,
			},
		},
	},
	SaveInstrument: &flag,
}

orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &cardPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### Card with saved instrument
```go
channel := "link"
cardCvv := "123"
instrument_id := "54deabb4-ba45-4a60-9e6a-9c016fe7ab10"
version := "2022-09-01"

cardPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		CardPaymentMethod: &cashfree.CardPaymentMethod{
			Card: cashfree.Card{
				Channel:        channel,
				CardCvv:        &cardCvv,
				InstrumentId: &instrument_id,
			},
		},
	},
}

orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &cardPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### Card (with native otp)
```go
channel := "post"
cardNumber := "4111111111111111"
cardHolderName := "Test"
cardMonth := "12"
cardYear := "25"
cardCvv := "123"
version := "2022-09-01"
flag := true

cardPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		CardPaymentMethod: &cashfree.CardPaymentMethod{
			Card: cashfree.Card{
				Channel:        channel,
				CardNumber:     &cardNumber,
				CardExpiryMm:   &cardMonth,
				CardHolderName: &cardHolderName,
				CardCvv:        &cardCvv,
				CardExpiryYy:   &cardYear,
			},
		},
	},
	SaveInstrument: &flag,
}

orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &cardPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

#### EMI
```go
channel := "link"
cardNumber := "4111111111111111"
cardHolderName := "Test"
cardMonth := "12"
cardYear := "25"
cardCvv := "123"
version := "2022-09-01"

emiCardPayRequest := cashfree.PayOrderRequest{
	PaymentSessionId: "session_LnO-vcZ9znG_swyugIFmZRtvP3ZC7euzAq4Gq8IfNjt68OFCJ31mbJsN8SWZ169G8y0awciDTv5wSGSgG-EDdG0eQTX1Ra43hdhWE4EtEEIJ",
	PaymentMethod: cashfree.PayOrderRequestPaymentMethod{
		CardEMIPaymentMethod: &cashfree.CardEMIPaymentMethod{
			Emi: cashfree.CardEMI{
				Channel:        channel,
				CardNumber:     cardNumber,
				CardHolderName: &cardHolderName,
				CardExpiryMm:   cardMonth,
				CardExpiryYy:   cardYear,
				CardCvv:        cardCvv,
				EmiTenure:      3,
				CardBankName:   "hdfc",
			},
		},
	},
}

orderPayResponse, httpResponse, err := cashfree.PGPayOrder(&version, &emiCardPayRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(orderPayResponse)
}
```

### Response
```json
{
  "payment_method": "card",
  "channel": "link",
  "action": "link",
  "cf_payment_id": 91235,
  "payment_amount": 22.42,
  "data": {
    "url": "https://sandbox.cashfree.com/pg/view/gateway/FHsuvhayLM5mmhINoqri7ba296e2ebca8b98e6119f6223021a13",
    "payload": {
      "name": "card"
    },
    "content_type": "application/x-www-form-urlencoded",
    "method": "post"
  }
}
```

## PGAuthorizeOrder

> PGAuthorizeOrder(xApiVersion *string, orderId string, authorizeOrderRequest *AuthorizeOrderRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*PaymentEntity, *http.Response, error)

Preauthorization ([Docs](https://docs.cashfree.com/reference/pgauthorizeorder))

### Example

```go
version := "2022-09-01"
action := "CAPTURE"
var amount float32 = 1.0
authorizeOrderRequest := cashfree.AuthorizeOrderRequest{
	Action: &action,
	Amount: &amount,
}

authorizeOrderResponse, httpResponse, err := cashfree.PGAuthorizeOrder(&version, "<order_id>", &authorizeOrderRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(authorizeOrderResponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string*** | The id which uniquely identifies your order | 
 **authorizeOrderRequest** | **AuthorizeOrderRequest*** | Request to Capture or Void Transactions |
 **xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
 **xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 
 

### Response
```json
{
  "cf_payment_id": 12376123,
  "order_id": "order_8123",
  "entity": "payment",
  "payment_currency": "INR",
  "error_details": null,
  "order_amount": 10.01,
  "is_captured": true,
  "payment_group": "upi",
  "authorization": {
    "action": "CAPTURE",
    "status": "PENDING",
    "captured_amount": 100,
    "start_time": "2022-02-09T18:04:34+05:30",
    "end_time": "2022-02-19T18:04:34+05:30",
    "approve_by": "2022-02-09T18:04:34+05:30",
    "action_reference": "6595231908096894505959",
    "action_time": "2022-08-03T16:09:51"
  },
  "payment_method": {
    "upi": {
      "channel": "collect",
      "upi_id": "rohit@xcxcx"
    }
  },
  "payment_amount": 10.01,
  "payment_time": "2021-07-23T12:15:06+05:30",
  "payment_completion_time": "2021-07-23T12:18:59+05:30",
  "payment_status": "SUCCESS",
  "payment_message": "Transaction successful",
  "bank_reference": "P78112898712",
  "auth_id": "A898101"
}
```


## PGOrderAuthenticatePayment

> PGOrderAuthenticatePayment(xApiVersion *string, cfPaymentId string, orderAuthenticatePaymentRequest *OrderAuthenticatePaymentRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*OrderAuthenticateEntity, *http.Response, error)

Submit or Resend OTP ([Docs](https://docs.cashfree.com/reference/pgorderauthenticatepayment))

### Example

```go
version := "2022-09-01"
action := "SUBMIT_OTP"
otp := "111000"
authenticatePaymentRequest := cashfree.OrderAuthenticatePaymentRequest{
	Action: action,
	Otp:    otp,
}

authenticatePaymentReqsponse, httpResponse, err := cashfree.PGOrderAuthenticatePayment(&version, "<cf_payment_id>", &authenticatePaymentRequest, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(authenticatePaymentReqsponse)
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**cfPaymentId** | **string*** | The Cashfree payment or transaction ID. | 
**xApiVersion** | **string** | API version to be used. Format is in YYYY-MM-DD | [default to &quot;2022-09-01&quot;]
**orderAuthenticatePaymentRequest** | **OrderAuthenticatePaymentRequest*** | Request body to submit/resend headless OTP. To use this API make sure you have headless OTP enabled for your account | 
**xRequestId** | **string** | Request id for the API call. Can be used to resolve tech issues. Communicate this in your tech related queries to cashfree | 


### Response
```json
{
  "cf_payment_id": 975654863,
  "authenticate_status": "FAILED",
  "action": "SUBMIT_OTP",
  "payment_message": "otp is invalid"
}
```

