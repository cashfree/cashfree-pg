# Go API client for cashfree_pg_sdk_go

Use our Golang SDK to integrate the Cashfree Payment Gateway into your application.

- API version: 2022-09-01
- Package version: 2.0.2

## Installation and Importing

Run the following command in the root level of the project

```shell
go get github.com/cashfree/cashfree-pg-sdk-go/v2
```

```golang
import cashfreeSDK "github.com/cashfree/cashfree-pg-sdk-go/v2/implementation"
```
---

# Version <= 0.0.5

## Creating a CFSession instance

The `CFSession` type consists of properties that are necessary for every method call that is exposed by the SDK. The following code snippet can be used to create a CFSession instance :-

```
environment := cashfreeSDK.PRODUCTION // cashfreeSDK.SANDBOX -> for sandbox environment
xApiVersion := "2022-01-01"
session := cashfreeSDK.CFSession{
	Environment:  &environment,
	ApiVersion:   &xApiVersion,
	ClientId:     &xClientId, // Send respective client id depending on the environment set
	ClientSecret: &xClientSecret, // Send respective secret key depending on the environment set
}
```

`NOTE` The environment in CFSession is of type `CFEnvironment` and the values are `PRODUCION` and `SANDBOX`

---

# Version >= 0.0.6

## Creating a CFConfig instance

The `CFConfig` type consists of properties that are necessary for every method call that is exposed by the SDK. The following code snippet can be used to create a CFConfig instance :-

```
environment := cashfreeSDK.PRODUCTION // cashfreeSDK.SANDBOX -> for sandbox environment
xApiVersion := "2022-09-01"
session := cashfreeSDK.CFConfig{
	Environment:  &environment,
	ApiVersion:   &xApiVersion,
	ClientId:     &xClientId, // Send respective client id depending on the environment set
	ClientSecret: &xClientSecret, // Send respective secret key depending on the environment set
}
```

`NOTE` The environment in CFSession is of type `CFEnvironment` and the values are `PRODUCION` and `SANDBOX`
`NOTE` CFConfig also supports setting timeout and proxy.

---

## Create a CFHeader instance

The `CFHeader` type consists of properties that are sent in the headers as part of every request and they are optional values. The following code snippet can be used to create a CFHeader instance :-

```
idempotencyKey := "some_random_string"
requestId := "request_id"
header := cashfreeSDK.CFHeader{
	RequestID:      &requestId,
	IdempotencyKey: &idempotencyKey,
}
```
---

## Create an Order with Cashfree

To process any payment on Cashfree PG, the merchant needs to create an order in the cashfree system. An order can be created using the following code snippet :

```
returnUrl := "https://merchant.in/pg/process_return?cf_id={order_id}"
orderMeta := cashfreeSDK.CFOrderMeta{
	ReturnUrl:      returnUrl,
	NotifyUrl:      returnUrl,
	PaymentMethods: "upi",
}
request := cashfreeSDK.CFOrderRequest{
	OrderId:       &or,
	OrderAmount:   1.0,
	OrderCurrency: "INR",
	CustomerDetails: cashfreeSDK.CFCustomerDetails{
		CustomerId:    customerId,
		CustomerEmail: "sample@gmail.com",
		CustomerPhone: "9999999999",
	},
	OrderNote: &orderId,
	OrderMeta: &orderMeta,
}

session := getSession() // above created session
header := getHeaders() // above created header, this is optional
cfOrder, responseHeader, cfError := cashfreeSDK.CreateOrder(&session, &header, request)
```
`Note:` For more information about order properties, visit [here](https://docs.cashfree.com/docs/create-order#curl-request)

---

## Pay Order

Once you have created the order, you can use the order to initiate payment. Order creation API returns "payment_session_id" which contains information about the order and that has to be used in payment initiation stage. Cashfree provides multiple payment methods to choose to make payments for an order, namely, UPI, Netbanking, Wallet, Card, Card EMI, Cardless EMI and Pay later.

### Pay Order with Card

Below is the code to initiate payment with Card

```
returnUrl := "https://merchant.in/pg/process_return?cf_id={order_id}"
orderMeta := cashfreeSDK.CFOrderMeta{
	ReturnUrl:      returnUrl,
	NotifyUrl:      returnUrl,
	PaymentMethods: "upi",
}
request := cashfreeSDK.CFOrderRequest{
	OrderId:       &or,
	OrderAmount:   1.0,
	OrderCurrency: "INR",
	CustomerDetails: cashfreeSDK.CFCustomerDetails{
		CustomerId:    customerId,
		CustomerEmail: "sample@gmail.com",
		CustomerPhone: "9999999999",
	},
	OrderNote: &orderId,
	OrderMeta: &orderMeta,
}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfOrder, responseHeader, cfError := cashfreeSDK.CreateOrder(&config, &header, request)

// Once the order is created and "payment_session_id" is generated, that should be used to initiate the payment

card := cashfreeSDK.CFOrderPaySessionIdRequest{
		SaveInstrument: &save,
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFCardPayment: &cashfreeSDK.CFCardPayment{
				Card: cashfreeSDK.CFCard{
					Channel:        "link",
					CardNumber:     cardNumber,
					CardHolderName: cardHolderName,
					CardExpiryMm:   cardExpiryMonth,
					CardExpiryYy:   cardExpiryYear,
					CardCvv:        cardCVV,
				},
			},
		},
	}

orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&session, &header, opr)
```

`Note:` Order has to be created and then the "payment_session_id" has to be used to make the payments for all the other payment methods as well. This step is covered in the above example and the same step has to be followed for all other payment methods.

---
### Pay Order with Saved Card

Below is the code to initiate payment with Saved Card

```
opr := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFCardPayment: &cashfreeSDK.CFCardPayment{
				Card: cashfreeSDK.CFCard{
					Channel:      "link",
					InstrumentId: instrument_id,
					CardCvv:      cardCVV,
				},
			},
		},
	}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, opr)
```
---

### Pay Order with UPI

Below is the code to initiate payment with UPI - `Collect`

```
upi := "testsuccess@gocash"
opr := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFUPIPayment: &cashfreeSDK.CFUPIPayment{
				Upi: cashfreeSDK.CFUPI{
					Channel: "collect",
					UpiId:   upi,
				},
			},
		},
	}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, opr)

```

Below is the code to initiate payment with UPI - `Intent`

```
upi := "testsuccess@gocash"
opr := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFUPIPayment: &cashfreeSDK.CFUPIPayment{
				Upi: cashfreeSDK.CFUPI{
					Channel: "link",
				},
			},
		},
	}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, opr)

```

Below is the code to initiate payment with UPI - `QRCode`

```
upi := "testsuccess@gocash"
opr := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFUPIPayment: &cashfreeSDK.CFUPIPayment{
				Upi: cashfreeSDK.CFUPI{
					Channel: "qrcode",
				},
			},
		},
	}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, opr)

```
---

### Pay Order with Netbanking

Below is the code to initiate payment with Netbanking

```
netbanking := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFNetbankingPayment: &cashfreeSDK.CFNetbankingPayment{
				Netbanking: cashfreeSDK.CFNetbanking{
					Channel:            "link",
					NetbankingBankCode: 3058,
				},
			},
		},
	}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, netbanking)
```

`Note:` For all bank codes, visit [here](https://docs.cashfree.com/docs/net-banking)

---

### Pay Order with App (Wallet)

Below is the code to initiate payment with App (Wallet)

```
app := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFAppPayment: &cashfreeSDK.CFAppPayment{
				App: cashfreeSDK.CFApp{
					Channel:  "link",
					Provider: "amazon",
					Phone:    "9999999999",
				},
			},
		},
	}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, app)
```

`Note:` Below is the list of all values for all the wallet providers supported by Cashfree:-
`phonepe` `paytm` `amazon` `airtel` `freecharge` `mobikwik` `jio` `ola`

---
### Pay Order with Paylater

Below is the code to initiate payment with Paylater

```
payLater := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFPaylaterPayment: &cashfreeSDK.CFPaylaterPayment{
				Paylater: cashfreeSDK.CFPaylater{
					Channel:  "link",
					Provider: "lazypay",
					Phone:    "9999999999",
				},
			},
		},
	}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, payLater)
```

`Note:` Currently we support "kotak", "flexipay", "zestmoney", "lazypay", "olapostpaid"

---
### Pay Order with CardlessEMI

Below is the code to initiate payment with CardlessEMI

```
cardlessEMI := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFCardlessEMIPayment: &cashfreeSDK.CFCardlessEMIPayment{
				CardlessEmi: cashfreeSDK.CFCardlessEMI{
					Channel:  "link",
					Provider: "zestmoney",
					Phone:    "9999999999",
				},
			},
		},
	}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, cardlessEMI)
```

`Note:` Currently we support "flexmoney", "zestmoney"

---
### Pay Order with EMI - Card

Below is the code to initiate payment with EMI - Card

```
opr := cashfreeSDK.CFOrderPaySessionIdRequest{
		PaymentSessionId: *cfOrder.PaymentSessionId,
		PaymentMethod: cashfreeSDK.CFPaymentMethod{
			CFEMIPayment: &cashfreeSDK.CFEMIPayment{
				Emi: &cashfreeSDK.CFCardEMI{
					Channel:        "link",
					CardNumber:     cardNumber,
					CardHolderName: &cardHolderName,
					CardExpiryMm:   cardEzpiryMonth,
					CardExpiryYy:   cardExpiryYear,
					CardCvv:        cardCVV,
					EmiTenure:      3,
					CardBankName:   "ICICI",
				},
			},
		},
	}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
orderPayResponse, responseHeader, cfError := cashfreeSDK.OrderPay(&config, &header, opr)
```

`Note:` Currently we support "Kotak", "ICICI", "RBL", "BOB", "Standard Chartered", "HDFC"

---

### Get Order

The details and status of the order can be fetched using this API. Below is the code snippet to retrieve order using `order_id`

```
orderId := "orderId"
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfOrder, responseHeader, cfError := cashfreeSDK.GetOrder(&session, &config, orderId)
```

---

### Get Payments For An Order

Once the payment process is initiated, all the payment information for a particular order can be retrieved through this API. Below is the code snippet to retrieve payments for a particular order.

```
orderId := "orderId"
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfPaymentsEntities, responseHeader, cfError := cashfreeSDK.GetPaymentsForOrder(&config, &header, orderId)
```
---

### Get Payment By CFPaymentID

Payment information can be retrieved using a unique ID generated by Cashfree (CFPaymentID). We can find get this ID from `GetOrder` API and can be used here to fetch Payment Information. Below is the code snippet to fetch Payment Information.

```
cfPaymentId := 965726202
orderId := "orderId"
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfPaymentsEntity, responseHeader, cfError := cashfreeSDK.GetPaymentByPaymentID(&config, &header, orderId, int32(cfPaymentId))
```

---

### Create Payment Links

A Payment link can be created and shared with users through `email` or `sms`. Below is the code snippet to create a payment link.

```
createLinkRequest := cashfreeSDK.CFLinkRequest{
	LinkId:       linkId,
	LinkAmount:   1.0,
	LinkCurrency: "INR",
	LinkPurpose:  linkPurpose,
	LinkNotify: &cashfreeSDK.CFLinkNotifyEntity{
		SendSms:   &flag,
		SendEmail: &flag,
	},
	CustomerDetails: cashfreeSDK.CFLinkCustomerDetailsEntity{
		CustomerPhone: "9999999999",
	},
}
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional

cfLink, responseHeader, cfError := cashfreeSDK.CreateLink(&config, &header, createLinkRequest)
```

---

### Get Payment Links By LinkID

A Payment link which was created can be retrieved using this API by specifying the `LinkID`. Below is the code snippet to get the payment link information.

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfLink, responseHeader, cfError := cashfreeSDK.GetLinkByLinkID(&config, &header, linkId)
```

---

### Get Orders By LinkID

All the orders created using links can be fetched using linkID by using this API. Below is the code snippet to get all the orders associated with the linkID.

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfLinkOrders, responseHeader, cfError := cashfreeSDK.GetOrdersByLinkID(&config, &header, linkID)
```

---

### Initiate Refund

Amount associated with a particular orderID where the transaction has gone through can be refunded using this API. Below is the code snippet to initiate the refund.

```
refundRequest := cashfreeSDK.CFRefundRequest{
	RefundAmount: refundAmount,
	RefundId:     refundId,
	RefundNote:   &refundNote,
}

config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfRefund, responseHeader, cfError := cashfreeSDK.InitiateRefund(&config, &header, refundRequest, orderId)
```

---

### Fetch Refund Information

Once the refund is initiated, we can check the status of that refund using this API. Below is the code snippet to check refund information.

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfRefund, responseHeader, cfError := cashfreeSDK.FetchRefundData(&config, &header, refundId, orderId)
```

---

### Get Settlements

Below is the code snippet to retrieve order settlement information

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfSettlementEntity, responseHeader, cfError := cashfreeSDK.Getsettlements(&config, &header, orderId)
```

---

### Get Saved Instruments By Customer ID

Saved instrument information can be retrieved using this API. Below is the code snippet to retrieve instrument information.

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
instrumentType := "card"
cfInstruments, responseHeader, cfError := cashfreeSDK.GetInstrumentsByCustomerID(&config, &header, customerId, instrumentType)
```

---

### Get Instrument by Instrument ID

Saved instrument information can be retrieved using this API by using the instrument ID. Below is the code snippet to retrieve instrument information.

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfInstruments, responseHeader, cfError := cashfreeSDK.GetInstrumentsByInstrumentID(&config, &header, customerId, instrument_id)
```

---

### Get Instrument Cryptogram by Instrument ID

```
config := getConfig() // above created config
header := getHeaders() // above created header, this is optional
cfCryptogram, responseHeader, cfError := cashfreeSDK.GetInstrumentCryptogramByInstrumentID(&config, &header, customerId, instrument_id)
```

---
## Documentation For Models

 - [CFApp](docs/CFApp.md)
 - [CFAppPayment](docs/CFAppPayment.md)
 - [CFAuthorizationInPaymentsEntity](docs/CFAuthorizationInPaymentsEntity.md)
 - [CFAuthorizationRequest](docs/CFAuthorizationRequest.md)
 - [CFCard](docs/CFCard.md)
 - [CFCardEMI](docs/CFCardEMI.md)
 - [CFCardPayment](docs/CFCardPayment.md)
 - [CFCardlessEMI](docs/CFCardlessEMI.md)
 - [CFCardlessEMIPayment](docs/CFCardlessEMIPayment.md)
 - [CFCryptogram](docs/CFCryptogram.md)
 - [CFCustomerDetails](docs/CFCustomerDetails.md)
 - [CFEMIPayment](docs/CFEMIPayment.md)
 - [CFError](docs/CFError.md)
 - [CFFetchAllSavedInstruments](docs/CFFetchAllSavedInstruments.md)
 - [CFLink](docs/CFLink.md)
 - [CFLinkCustomerDetailsEntity](docs/CFLinkCustomerDetailsEntity.md)
 - [CFLinkMetaEntity](docs/CFLinkMetaEntity.md)
 - [CFLinkNotifyEntity](docs/CFLinkNotifyEntity.md)
 - [CFLinkOrders](docs/CFLinkOrders.md)
 - [CFLinkRequest](docs/CFLinkRequest.md)
 - [CFNetbanking](docs/CFNetbanking.md)
 - [CFNetbankingPayment](docs/CFNetbankingPayment.md)
 - [CFOrder](docs/CFOrder.md)
 - [CFOrderMeta](docs/CFOrderMeta.md)
 - [CFOrderPayData](docs/CFOrderPayData.md)
 - [CFOrderPayRequest](docs/CFOrderPayRequest.md)
 - [CFOrderPayResponse](docs/CFOrderPayResponse.md)
 - [CFOrderRequest](docs/CFOrderRequest.md)
 - [CFPaylater](docs/CFPaylater.md)
 - [CFPaylaterPayment](docs/CFPaylaterPayment.md)
 - [CFPaymentMethod](docs/CFPaymentMethod.md)
 - [CFPaymentURLObject](docs/CFPaymentURLObject.md)
 - [CFPaymentsEntity](docs/CFPaymentsEntity.md)
 - [CFPaymentsEntityAppPayment](docs/CFPaymentsEntityAppPayment.md)
 - [CFPaymentsEntityCardPayment](docs/CFPaymentsEntityCardPayment.md)
 - [CFPaymentsEntityCardlessEMIPayment](docs/CFPaymentsEntityCardlessEMIPayment.md)
 - [CFPaymentsEntityMethod](docs/CFPaymentsEntityMethod.md)
 - [CFPaymentsEntityNetbankingPayment](docs/CFPaymentsEntityNetbankingPayment.md)
 - [CFPaymentsEntityPaylaterPayment](docs/CFPaymentsEntityPaylaterPayment.md)
 - [CFPaymentsEntityPayment](docs/CFPaymentsEntityPayment.md)
 - [CFPaymentsEntityUPIPayment](docs/CFPaymentsEntityUPIPayment.md)
 - [CFRefund](docs/CFRefund.md)
 - [CFRefundRequest](docs/CFRefundRequest.md)
 - [CFRefundURLObject](docs/CFRefundURLObject.md)
 - [CFSavedInstrumentMeta](docs/CFSavedInstrumentMeta.md)
 - [CFSettlementURLObject](docs/CFSettlementURLObject.md)
 - [CFSettlementsEntity](docs/CFSettlementsEntity.md)
 - [CFUPI](docs/CFUPI.md)
 - [CFUPIAuthorizeDetails](docs/CFUPIAuthorizeDetails.md)
 - [CFUPIPayment](docs/CFUPIPayment.md)
 - [CFVendorSplit](docs/CFVendorSplit.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

nextgenapi@cashfree.com

# Support

## Ask a question about Integration

You can ask questions, and participate in discussions about Cashfree-related topics ian the Cashfree Discord channel.

<a href="https://discord.com/invite/VxW7geAEd2"><img src="https://cashfreelogo.cashfree.com/discord_banner_purple.svg" /></a>
