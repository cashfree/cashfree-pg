# Webhook


## PGVerifyWebhookSignature

> PGVerifyWebhookSignature(signature string, rawBody string, timestamp string) (*PGWebhookEvent, error)

Verify Webhook Signatures ([Docs](https://docs.cashfree.com/reference/pg-webhook)

### Example

```go
// Route
e.POST("/webhook", _this.Webhook)

// Controller
func (_this *WebhookRoute) Webhook(c echo.Context) error {
    signature := c.Request().Header.Get("x-webhook-signature")
    timestamp := c.Request().Header.Get("x-webhook-timestamp")
    
    body, _ := ioutil.ReadAll(c.Request().Body)
    rawBody := string(body)
    webhookEvent, err := cashfree.PGVerifyWebhookSignature(signature, rawBody, timestamp)
    if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(webhookEvent.Object)
	}
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**signature** | **string** | The "x-webhook-signature" present in the headers received | 
**rawBody** | **string** | The request body received in the request in string format |
**timestamp** | **string** | The "x-webhook-timestamp" present in the headers received | 


### Response

```json
{
    "type": "PAYMENT_SUCCESS_WEBHOOK",
    "rawBody":"{\"type\":\"PAYMENT_SUCCESS_WEBHOOK\",\"data\":{}}",
    "object": {
        "type": "PAYMENT_SUCCESS_WEBHOOK",
        "data": {
            
        }
    }
}
```

`Note:` The `object` in the response is returned as a JSON object. But it can be of any of the following types:
- [Payment Webhook](https://docs.cashfree.com/docs/payment-webhooks)
- [Refund Webhook](https://docs.cashfree.com/docs/refunds-webhook)
- [Settlement Webhook](https://docs.cashfree.com/docs/settlements-webhook)
- [Instrument Webhook](https://docs.cashfree.com/docs/instrument-webhook)

