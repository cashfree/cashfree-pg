package cashfree_pg_test

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	cashfree "github.com/cashfree/cashfree-pg/v3"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func Test_cashfree_pg_payments(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	orderId := "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6"
	paymentId := "14909870498"
	XApiVersion := "2022-09-01"
	ctx := context.Background()

	t.Run("PGOrderFetchPayments should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchPayments(&XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPayments should fail when xapiversion is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchPayments(nil, orderId, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchPayments should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGOrderFetchPayments(&XApiVersion, orderId+"gosdk", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPayments shoudl give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		resp, httpRes, err := cashfree.PGOrderFetchPayments(&XApiVersion, orderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###########################
	// Fetch Payments With Context
	// ###########################

	t.Run("PGOrderFetchPaymentsWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGOrderFetchPaymentsWithContext(ctx, &XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPaymentsWithContext should fail when xapiversion is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchPaymentsWithContext(ctx, nil, orderId, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchPaymentsWithContext should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGOrderFetchPaymentsWithContext(ctx, &XApiVersion, orderId+"gosdk", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPaymentsWithContext shoudl give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		resp, httpRes, err := cashfree.PGOrderFetchPaymentsWithContext(ctx, &XApiVersion, orderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###########################
	// Fetch Payments By Id
	// ###########################

	t.Run("PGOrderFetchPayment should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGOrderFetchPayment(&XApiVersion, orderId, paymentId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPayment should fail when xapiversion is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchPayment(nil, orderId, paymentId, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchPayment should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGOrderFetchPayment(&XApiVersion, orderId+"gosdk", paymentId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPayment shoudl give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		resp, httpRes, err := cashfree.PGOrderFetchPayment(&XApiVersion, orderId, paymentId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###########################
	// Fetch Payments By Id With Context
	// ###########################

	t.Run("PGOrderFetchPaymentWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGOrderFetchPaymentWithContext(ctx, &XApiVersion, orderId, paymentId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPaymentWithContext should fail when xapiversion is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchPaymentWithContext(ctx, nil, orderId, paymentId, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchPaymentWithContext should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGOrderFetchPaymentWithContext(ctx, &XApiVersion, orderId+"gosdk", paymentId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchPaymentWithContext shoudl give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		resp, httpRes, err := cashfree.PGOrderFetchPaymentWithContext(ctx, &XApiVersion, orderId, paymentId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###################
	// Order Pay
	// ###################

	t.Run("PGOrderPay should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, _, _ := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		upiId := "testsuccess@gocash"
		upiPayment := cashfree.PayOrderRequestPaymentMethod{
			UPIPaymentMethod: &cashfree.UPIPaymentMethod{
				Upi: cashfree.Upi{
					Channel: "collect",
					UpiId:   &upiId,
				},
			},
		}

		cardPayOrderRequest := cashfree.PayOrderRequest{
			PaymentSessionId: *resp.PaymentSessionId,
			PaymentMethod:    upiPayment,
		}

		card, httpRes, err := cashfree.PGPayOrder(&XApiVersion, &cardPayOrderRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, card)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderPay should fail when order request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGPayOrder(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGOrderPay should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGPayOrder(nil, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGOrderPay should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		upiId := "testsuccess@gocash"
		upiPayment := cashfree.PayOrderRequestPaymentMethod{
			UPIPaymentMethod: &cashfree.UPIPaymentMethod{
				Upi: cashfree.Upi{
					Channel: "collect",
					UpiId:   &upiId,
				},
			},
		}

		cardPayOrderRequest := cashfree.PayOrderRequest{
			PaymentSessionId: "invalid",
			PaymentMethod:    upiPayment,
		}

		card, httpRes, err := cashfree.PGPayOrder(&XApiVersion, &cardPayOrderRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, card)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	// ###################
	// Order Pay With Context
	// ###################

	t.Run("PGOrderPayWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, _, _ := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		upiId := "testsuccess@gocash"
		upiPayment := cashfree.PayOrderRequestPaymentMethod{
			UPIPaymentMethod: &cashfree.UPIPaymentMethod{
				Upi: cashfree.Upi{
					Channel: "collect",
					UpiId:   &upiId,
				},
			},
		}

		cardPayOrderRequest := cashfree.PayOrderRequest{
			PaymentSessionId: *resp.PaymentSessionId,
			PaymentMethod:    upiPayment,
		}

		card, httpRes, err := cashfree.PGPayOrderWithContext(ctx, &XApiVersion, &cardPayOrderRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, card)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderPayWithContext should fail when order request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGPayOrderWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGOrderPayWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGPayOrderWithContext(ctx, nil, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGOrderPayWithContext should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		upiId := "testsuccess@gocash"
		upiPayment := cashfree.PayOrderRequestPaymentMethod{
			UPIPaymentMethod: &cashfree.UPIPaymentMethod{
				Upi: cashfree.Upi{
					Channel: "collect",
					UpiId:   &upiId,
				},
			},
		}

		cardPayOrderRequest := cashfree.PayOrderRequest{
			PaymentSessionId: "invalid",
			PaymentMethod:    upiPayment,
		}

		card, httpRes, err := cashfree.PGPayOrderWithContext(ctx, &XApiVersion, &cardPayOrderRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, card)
		assert.Equal(t, 404, httpRes.StatusCode)

	})
}
