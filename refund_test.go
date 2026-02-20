package cashfree_pg_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	cashfree "github.com/cashfree/cashfree-pg/v3"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func Test_cashfree_pg_refunds(t *testing.T) {

	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	XApiVersion := "2023-08-01"
	ctx := context.Background()
	unpaidOrderId := "order_" + uniqueSuffix()
	orderId := "order_" + uniqueSuffix()

	createOrder := func(orderID string) (*cashfree.OrderEntity, *http.Response, error) {
		createOrderRequest := cashfree.CreateOrderRequest{
			OrderId:       &orderID,
			OrderAmount:   20.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		var resp *cashfree.OrderEntity
		var httpRes *http.Response
		var err error
		for attempt := 0; attempt < eventualConsistencyMaxAttempts; attempt++ {
			resp, httpRes, err = cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)
			if httpRes != nil && httpRes.StatusCode == http.StatusOK {
				return resp, httpRes, err
			}
			sleepBeforeEventualRetry(attempt)
		}
		return resp, httpRes, err
	}

	_, unpaidOrderHTTPRes, unpaidOrderErr := createOrder(unpaidOrderId)
	requireSuccessOrDecodeError(t, unpaidOrderHTTPRes, unpaidOrderErr)

	orderResp, orderHTTPRes, orderErr := createOrder(orderId)
	requireSuccessOrDecodeError(t, orderHTTPRes, orderErr)

	paymentSessionID := ""
	if orderResp != nil && orderResp.PaymentSessionId != nil {
		paymentSessionID = *orderResp.PaymentSessionId
	}
	if paymentSessionID == "" {
		paymentSessionID = extractStringFromErrorBody(orderErr, "payment_session_id")
	}
	require.NotEmpty(t, paymentSessionID)

	upiID := "testsuccess@gocash"
	upiPayment := cashfree.PayOrderRequestPaymentMethod{
		UPIPaymentMethod: &cashfree.UPIPaymentMethod{
			Upi: cashfree.Upi{
				Channel: "collect",
				UpiId:   &upiID,
			},
		},
	}
	payOrderRequest := cashfree.PayOrderRequest{
		PaymentSessionId: paymentSessionID,
		PaymentMethod:    upiPayment,
	}
	seedReq := "seed-" + uniqueSuffix()
	seedIdempotency := uniqueSuffix()
	var payHTTPRes *http.Response
	var payErr error
	for attempt := 0; attempt < eventualConsistencyMaxAttempts; attempt++ {
		_, payHTTPRes, payErr = cashfree.PGPayOrder(&XApiVersion, &payOrderRequest, &seedReq, &seedIdempotency, http.DefaultClient)
		if payHTTPRes != nil && payHTTPRes.StatusCode == http.StatusOK {
			break
		}
		sleepBeforeEventualRetry(attempt)
	}
	requireSuccessOrDecodeError(t, payHTTPRes, payErr)

	// Let the sandbox state settle before issuing refund calls.
	time.Sleep(2 * time.Second)

	t.Run("PGCreateRefund should give status code 400 when order is unpaid", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		req := "test"
		idemp := uniqueSuffix()
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, unpaidOrderId, &createOrderRefundRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefund should fail when request is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, nil, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefund should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		_, _, err := cashfree.PGOrderCreateRefund(nil, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefund should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefund should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 999999999.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefund should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId+"111", &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGCreateRefund should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###########################
	// Create Refund With Context
	// ###########################

	t.Run("PGCreateRefundWithContext should give status code 400 when order is unpaid", func(t *testing.T) {

		cashfree.XClientId = &clientId

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		req := "test"
		idemp := uniqueSuffix()
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, unpaidOrderId, &createOrderRefundRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefundWithContext should fail when request is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, nil, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefundWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		_, _, err := cashfree.PGOrderCreateRefundWithContext(ctx, nil, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefundWithContext should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefundWithContext should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 999999999.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateRefundWithContext should give status code 400", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     "sadaaskaahaaas1daaaasab123",
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId+"111", &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGCreateRefundWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     uniqueSuffix(),
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch All Refunds
	// ########################

	t.Run("PGOrderFetchRefunds should give status code 200", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefunds(&XApiVersion, orderId, &xReq, &xReq, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefunds should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefunds(nil, orderId, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefunds should give status code 404", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefunds(&XApiVersion, "", &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefunds should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefunds(&XApiVersion, orderId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch All Refunds With Context
	// ########################

	t.Run("PGOrderFetchRefundsWithContext should give status code 200", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundsWithContext(ctx, &XApiVersion, orderId, &xReq, &xReq, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundsWithContext should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefundsWithContext(ctx, nil, orderId, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefundsWithContext should give status code 404", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundsWithContext(ctx, &XApiVersion, "", &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefundsWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundsWithContext(ctx, &XApiVersion, orderId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch Refunds by Refund Id
	// ########################

	refundID := "refund_" + uniqueSuffix()

	t.Run("PGOrderFetchRefund should give status code 404", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, orderId, refundID, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefund should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefund(nil, orderId, refundID, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefund should give status code 404", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, "", refundID, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefund should give status code 400", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, "=", refundID, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefund should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, orderId, refundID, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch All Refunds With Context
	// ########################

	refundIDWithContext := "refund_" + uniqueSuffix()

	t.Run("PGOrderFetchRefundWithContext should give status code 404", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, orderId, refundIDWithContext, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefundWithContext should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefundWithContext(ctx, nil, orderId, refundIDWithContext, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 404", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, "", refundIDWithContext, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 400", func(t *testing.T) {
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, "=", refundIDWithContext, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, orderId, refundIDWithContext, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
