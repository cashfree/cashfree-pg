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
	XApiVersion := "2022-09-01"
	ctx := context.Background()

	t.Run("PGCreateOrder should give status 200", func(t *testing.T) {

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should give status code 409", func(t *testing.T) {

		orderId := "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderId:       *cashfree.NewNullableString(&orderId),
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
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, nil, nil, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOrder should give status 200", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrder(&XApiVersion, &createOrderRequest, &requestId, &idempotency, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

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
		assert.Equal(t, 422, httpRes.StatusCode)

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
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should give status code 409", func(t *testing.T) {

		orderId := "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderId:       *cashfree.NewNullableString(&orderId),
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
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, nil, nil, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOrderWithContext should give status 200", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		createOrderRequest := cashfree.CreateOrderRequest{
			OrderAmount:   1.0,
			OrderCurrency: "INR",
			CustomerDetails: cashfree.CustomerDetails{
				CustomerId:    "suhas-test",
				CustomerPhone: "9999999999",
			},
		}
		resp, httpRes, err := cashfree.PGCreateOrderWithContext(ctx, &XApiVersion, &createOrderRequest, &requestId, &idempotency, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

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
		assert.Equal(t, 422, httpRes.StatusCode)

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

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrder should give status code 200 with client", func(t *testing.T) {
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrder should give status code 200 with request id", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", &requestId, &idempotency, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrder should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52J", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOrder should fail when order id missing", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOrder should fail when xApiVerion missing", func(t *testing.T) {
		resp, _, err := cashfree.PGFetchOrder(nil, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGFetchOrder should give status code 401", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrder(&XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ###############################
	// Fetch Order With Context
	// ###############################

	t.Run("PGFetchOrder should status code 200", func(t *testing.T) {
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrderWithContext should give status code 200 with client", func(t *testing.T) {
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrderWithContext should give status code 200 with request id", func(t *testing.T) {

		requestId := "test"
		idempotency := "test"

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", &requestId, &idempotency, nil)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGFetchOrderWithContext should give status code 404", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52J", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOrderWithContext should fail when order id missing", func(t *testing.T) {

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOrderWithContext should fail when xApiVerion missing", func(t *testing.T) {
		resp, _, err := cashfree.PGFetchOrderWithContext(ctx, nil, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)

	})

	t.Run("PGFetchOrderWithContext should give status code 401", func(t *testing.T) {

		clientId := "unauthorised"
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOrderWithContext(ctx, &XApiVersion, "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6", nil, nil, nil)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
