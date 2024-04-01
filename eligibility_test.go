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

func Test_cashfree_pg_eligibility(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	orderId := "order_342Z7ns5LWu4x4xIFvQqmF7x52Jc6"
	// amount := float32(1000.0)
	XApiVersion := "2023-08-01"
	ctx := context.Background()

	t.Run("PGEligibilityFetchPaymentMethods should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethods(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaymentMethods should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		_, _, err := cashfree.PGEligibilityFetchPaymentMethods(nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaymentMethods should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchPaymentMethods(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaymentMethods should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethods(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaymentMethods should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethods(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #######################
	// Payment Methods With Context
	// #######################

	t.Run("PGEligibilityFetchPaymentMethodsWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethodsWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaymentMethodsWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		_, _, err := cashfree.PGEligibilityFetchPaymentMethodsWithContext(ctx, nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaymentMethodsWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchPaymentMethodsWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaymentMethodsWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethodsWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaymentMethodsWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaymentMethodsRequest{
			Queries: cashfree.PaymentMethodsQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaymentMethodsWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #######################
	// Pay Later
	// #######################

	t.Run("PGEligibilityFetchPaylater should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylater(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaylater should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		_, _, err := cashfree.PGEligibilityFetchPaylater(nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaylater should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchPaylater(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaylater should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylater(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaylater should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylater(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #######################
	// Pay Later With context
	// #######################

	t.Run("PGEligibilityFetchPaylaterWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylaterWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaylaterWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		_, _, err := cashfree.PGEligibilityFetchPaylaterWithContext(ctx, nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaylaterWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchPaylaterWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchPaylaterWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylaterWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchPaylaterWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchPaylaterRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchPaylaterWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #######################
	// Cardless emi
	// #######################

	t.Run("PGEligibilityFetchCardlessEMI should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMI(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchCardlessEMI should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		_, _, err := cashfree.PGEligibilityFetchCardlessEMI(nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchCardlessEMI should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchCardlessEMI(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchCardlessEMI should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMI(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchCardlessEMI should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMI(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #######################
	// Pay Later With context
	// #######################

	t.Run("PGEligibilityFetchCardlessEMIWithContext should give status code 200", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMIWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchCardlessEMIWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		_, _, err := cashfree.PGEligibilityFetchCardlessEMIWithContext(ctx, nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchCardlessEMIWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGEligibilityFetchCardlessEMIWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchCardlessEMIWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMIWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchCardlessEMIWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchCardlessEMIRequest{
			Queries: cashfree.CardlessEMIQueries{
				OrderId: &orderId,
				CustomerDetails: &cashfree.CustomerDetailsCardlessEMI{
					CustomerPhone: "8904216227",
				},
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchCardlessEMIWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #########################
	// Offers
	// #########################

	t.Run("PGEligibilityFetchOffers should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{
			Queries: cashfree.OfferQueries{
				OrderId: &orderId,
			},
		}

		_, _, err := cashfree.PGEligibilityFetchOffers(nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchOffers should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchOffers(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchOffers should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		_, _, err := cashfree.PGEligibilityFetchOffers(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchOffers should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{
			Queries: cashfree.OfferQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchOffers(&XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #########################
	// Offers With Context
	// #########################

	t.Run("PGEligibilityFetchOffersWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{
			Queries: cashfree.OfferQueries{
				OrderId: &orderId,
			},
		}

		_, _, err := cashfree.PGEligibilityFetchOffersWithContext(ctx, nil, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchOffersWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{}

		resp, httpRes, err := cashfree.PGEligibilityFetchOffersWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGEligibilityFetchOffersWithContext should fail when request is missing", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		_, _, err := cashfree.PGEligibilityFetchOffersWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGEligibilityFetchOffersWithContext should give status code 401", func(t *testing.T) {

		req := "TEST"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		eligibilityFetchPaymentMethodsRequest := cashfree.EligibilityFetchOffersRequest{
			Queries: cashfree.OfferQueries{
				OrderId: &orderId,
			},
		}

		resp, httpRes, err := cashfree.PGEligibilityFetchOffersWithContext(ctx, &XApiVersion, &eligibilityFetchPaymentMethodsRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
