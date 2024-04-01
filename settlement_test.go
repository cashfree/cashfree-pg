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

func Test_cashfree_pg_settlement(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	orderId := "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6"
	// amount := float32(1000.0)
	XApiVersion := "2023-08-01"
	ctx := context.Background()

	t.Run("PGOrderFetchSettlement should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGOrderFetchSettlement(&XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchSettlement should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGOrderFetchSettlement(&XApiVersion, orderId+"test", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchSettlementWithContext should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGOrderFetchSettlementWithContext(ctx, &XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchSettlementWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGOrderFetchSettlementWithContext(ctx, &XApiVersion, orderId+"test", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
