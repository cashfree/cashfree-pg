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

func Test_cashfree_pg_settlement(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	orderId := "invalid_order_" + uniqueSuffix()
	XApiVersion := "2023-08-01"
	ctx := context.Background()

	t.Run("PGOrderFetchSettlement should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := uniqueSuffix()

		resp, httpRes, err := cashfree.PGOrderFetchSettlement(&XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchSettlement should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := uniqueSuffix()
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGOrderFetchSettlement(&XApiVersion, orderId+"test", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		require.NotNil(t, httpRes)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	t.Run("PGOrderFetchSettlementWithContext should give status code 404", func(t *testing.T) {

		req := "TEST"
		idemp := uniqueSuffix()
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGOrderFetchSettlementWithContext(ctx, &XApiVersion, orderId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("PGOrderFetchSettlementWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := uniqueSuffix()
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGOrderFetchSettlementWithContext(ctx, &XApiVersion, orderId+"test", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		require.NotNil(t, httpRes)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
