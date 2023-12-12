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

func Test_cashfree_pg_refunds(t *testing.T) {

	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	XApiVersion := "2022-09-01"
	ctx := context.Background()
	orderId := "order_342Z7qpz85EsLl3nj0DxSchfzzx19"
	refundId := "gosdktestrefund"

	t.Run("PGCreateRefund should give status code 200", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, &createOrderRefundRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateRefund should fail when request is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, nil, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefund should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
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
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGCreateRefund should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
		}
		resp, httpRes, err := cashfree.PGOrderCreateRefund(&XApiVersion, orderId, &createOrderRefundRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###########################
	// Create Refund With Context
	// ###########################

	t.Run("PGCreateRefundWithContext should give status code 200", func(t *testing.T) {

		cashfree.XClientId = &clientId

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, &createOrderRefundRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateRefundWithContext should fail when request is missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderCreateRefundWithContext(ctx, &XApiVersion, orderId, nil, nil, nil, nil)

		require.NotNil(t, err)

	})

	t.Run("PGCreateRefundWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
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
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGCreateRefundWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c

		createOrderRefundRequest := cashfree.OrderCreateRefundRequest{
			RefundAmount: 1.0,
			RefundId:     strconv.Itoa(int(time.Now().Unix())),
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
		xReq := strconv.Itoa(int(time.Now().Unix()))

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
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefunds(&XApiVersion, "", &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefunds should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := strconv.Itoa(int(time.Now().Unix()))

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
		xReq := strconv.Itoa(int(time.Now().Unix()))

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
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundsWithContext(ctx, &XApiVersion, "", &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundsWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundsWithContext(ctx, &XApiVersion, orderId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch Refunds by Refund Id
	// ########################

	t.Run("PGOrderFetchRefund should give status code 200", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, orderId, refundId, &xReq, &xReq, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefund should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefund(nil, orderId, refundId, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefund should give status code 404", func(t *testing.T) {
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, "", refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefund should give status code 400", func(t *testing.T) {
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, "=", refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefund should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefund(&XApiVersion, orderId, refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Fetch All Refunds With Context
	// ########################

	t.Run("PGOrderFetchRefundWithContext should give status code 200", func(t *testing.T) {

		cashfree.XClientId = &clientId
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, orderId, refundId, &xReq, &xReq, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundWithContext should fail when xApiVersion missing", func(t *testing.T) {

		_, _, err := cashfree.PGOrderFetchRefundWithContext(ctx, nil, orderId, refundId, nil, nil, nil)
		require.NotNil(t, err)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 404", func(t *testing.T) {
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, "", refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 400", func(t *testing.T) {
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, "=", refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchRefundWithContext should give status code 401", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		xReq := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchRefundWithContext(ctx, &XApiVersion, orderId, refundId, &xReq, &xReq, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
