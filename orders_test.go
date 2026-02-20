package cashfree_pg_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	cashfree "github.com/cashfree/cashfree-pg/v3"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func Test_cashfree_pg_orders(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	XApiVersion := "2023-08-01"
	ctx := context.Background()
	seedOrderId := "order_" + uniqueSuffix()
	invalidOrderId := "invalid_order_" + uniqueSuffix()

	seedCreateOrderRequest := cashfree.CreateOrderRequest{
		OrderId:       &seedOrderId,
		OrderAmount:   1.0,
		OrderCurrency: "INR",
		CustomerDetails: cashfree.CustomerDetails{
			CustomerId:    "suhas-test",
			CustomerPhone: "9999999999",
		},
	}
	var seedOrderHTTPRes *http.Response
	var seedOrderErr error
	for attempt := 0; attempt < eventualConsistencyMaxAttempts; attempt++ {
		_, seedOrderHTTPRes, seedOrderErr = cashfree.PGCreateOrder(&XApiVersion, &seedCreateOrderRequest, nil, nil, nil)
		if seedOrderHTTPRes != nil && seedOrderHTTPRes.StatusCode == http.StatusOK {
			break
		}
		sleepBeforeEventualRetry(attempt)
	}
	requireSuccessOrDecodeError(t, seedOrderHTTPRes, seedOrderErr)

	t.Run("PGCreateOrder should give status 200", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		_, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGCreateOrder should give status code 409", func(t *testing.T) {

		orderId := seedOrderId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderId:       &orderId,
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 409, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should give status 200", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		_, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	// t.Run("PGCreateOrder should give status 200", func(t *testing.T) {

	// 	requestId := "test"
	// 	idempotency := "test"

	// 	createOrderRequest := cashfree.CreateOrderRequest{
	// 		OrderAmount:   1.0,
	// 		OrderCurrency: "INR",
	// 		CustomerDetails: cashfree.CustomerDetails{
	// 			CustomerId:    "suhas-test",
	// 			CustomerPhone: "9999999999",
	// 		},
	// 	}
	// 	resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, &requestId, &idempotency, nil)

	// 	require.Nil(t, err)
	// 	require.NotNil(t, resp)
	// 	assert.Equal(t, 200, httpRes.StatusCode)

	// })

	t.Run("PGCreateOrder should give status 422", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "100",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, &requestId, &idempotency, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should fail when order amount is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should fail when order currency is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount: 1.0,
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount: 1.0,
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, _, err := cashfree.PGCreateOrder(nil, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrder should fail when request is missing", func(t *testing.T) {
		resp, _, err := cashfree.PGCreateOrder(&XApiVersion, nil, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrder should fail when request is missing", func(t *testing.T) {
		resp, _, err := cashfree.PGCreateOrder(&XApiVersion, nil, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrder should give 401 status code", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ############################
	// With Context
	// ############################

	t.Run("PGCreateOrderWithContext should give status 200", func(t *testing.T) {
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		_, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGCreateOrderWithContext should give status code 409", func(t *testing.T) {

		orderId := seedOrderId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderId:       &orderId,
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 409, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should give status 200", func(t *testing.T) {
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		_, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGCreateOrderWithContext should give status 422", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "100",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, &requestId, &idempotency, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should fail when order amount is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should fail when order currency is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount: 1.0,
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount: 1.0,
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, _, err := cashfree.PGCreateOrderWithContext(ctx, nil, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrderWithContext should fail when request is missing", func(t *testing.T) {
		resp, _, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, nil, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrderWithContext should fail when request is missing", func(t *testing.T) {
		resp, _, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, nil, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGCreateOrderWithContext should give 401 status code", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###############################
	// Fetch Order
	// ###############################

	t.Run("PGFetchOrder should status code 200", func(t *testing.T) {
		cashfree.XClientId = &clientId

		_, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, seedOrderId, nil, nil, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrder should give status code 200 with client", func(t *testing.T) {
		cashfree.XClientId = &clientId

		_, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, seedOrderId, nil, nil, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrder should give status code 200 with request id", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		_, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, seedOrderId, &requestId, &idempotency, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrder should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, invalidOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGFetchOrder should fail when order id missing", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGFetchOrder should fail when xApiVerion missing", func(t *testing.T) {
		resp, _, err := cashfree.PGFetchOrder(nil, seedOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGFetchOrder should give status code 401", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, seedOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###############################
	// Fetch Order With Context
	// ###############################

	t.Run("PGFetchOrder should status code 200", func(t *testing.T) {
		cashfree.XClientId = &clientId

		_, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, seedOrderId, nil, nil, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrderWithContext should give status code 200 with client", func(t *testing.T) {
		cashfree.XClientId = &clientId

		_, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, seedOrderId, nil, nil, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrderWithContext should give status code 200 with request id", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		_, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, seedOrderId, &requestId, &idempotency, nil)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("PGFetchOrderWithContext should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, invalidOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGFetchOrderWithContext should fail when order id missing", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGFetchOrderWithContext should fail when xApiVerion missing", func(t *testing.T) {
		resp, _, err := cashfree.PGFetchOrderWithContext(ctx, nil, seedOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGFetchOrderWithContext should give status code 401", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, seedOrderId, nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
