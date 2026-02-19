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

func Test_cashfree_pg_links(t *testing.T) {

	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	XApiVersion := "2023-08-01"
	linkId := "go_sdk_link_" + uniqueSuffix()
	ctx := context.Background()

	flag := true
	seedCreateLinkRequest := cashfree.CreateLinkRequest{
		LinkId:       linkId,
		LinkAmount:   1,
		LinkCurrency: "INR",
		LinkPurpose:  "gosdk_testing",
		CustomerDetails: cashfree.LinkCustomerDetailsEntity{
			CustomerPhone: "9999999999",
		},
		LinkNotify: &cashfree.LinkNotifyEntity{
			SendSms: &flag,
		},
	}
	seedRequestID := "seed-" + uniqueSuffix()
	seedIdempotency := uniqueSuffix()
	_, seedLinkHTTPRes, seedLinkErr := cashfree.PGCreateLink(&XApiVersion, &seedCreateLinkRequest, &seedRequestID, &seedIdempotency, http.DefaultClient)
	requireSuccessOrDecodeError(t, seedLinkHTTPRes, seedLinkErr)

	t.Run("Test Orders PGCreateLink xapiversion missing", func(t *testing.T) {

		createLinkRequest := cashfree.CreateLinkRequest{
			LinkId:       strconv.Itoa(int(time.Now().Unix())),
			LinkAmount:   1,
			LinkCurrency: "INR",
			LinkPurpose:  "gosdk_testing",
			CustomerDetails: cashfree.LinkCustomerDetailsEntity{
				CustomerPhone: "9999999999",
			},
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGCreateLink(nil, &createLinkRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGCreateLink request missing", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGCreateLink(&XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGCreateLink 401 status code", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		flag := true

		createLinkRequest := cashfree.CreateLinkRequest{
			LinkId:       strconv.Itoa(int(time.Now().Unix())),
			LinkAmount:   1,
			LinkCurrency: "INR",
			LinkPurpose:  "gosdk_testing",
			CustomerDetails: cashfree.LinkCustomerDetailsEntity{
				CustomerPhone: "9999999999",
			},
			LinkNotify: &cashfree.LinkNotifyEntity{
				SendSms: &flag,
			},
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGCreateLink(&XApiVersion, &createLinkRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ########################
	// Craete Link WIth Context
	// ########################

	// t.Run("Test Orders PGCreateLink 200 status code", func(t *testing.T) {

	// 	flag := true
	// 	cashfree.XClientId = &clientId

	// 	createLinkRequest := cashfree.CreateLinkRequest{
	// 		LinkId:       strconv.Itoa(int(time.Now().Unix())),
	// 		LinkAmount:   1,
	// 		LinkCurrency: "INR",
	// 		LinkPurpose:  "gosdk_testing",
	// 		CustomerDetails: cashfree.LinkCustomerDetailsEntity{
	// 			CustomerPhone: "9999999999",
	// 		},
	// 		LinkNotify: &cashfree.LinkNotifyEntity{
	// 			SendSms: &flag,
	// 		},
	// 	}
	// 	req := "test"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	resp, httpRes, err := cashfree.PGCreateLinkWithContext(ctx, &XApiVersion, &createLinkRequest, &req, &idemp, http.DefaultClient)

	// 	require.Nil(t, err)
	// 	require.NotNil(t, resp)
	// 	assert.Equal(t, 200, httpRes.StatusCode)

	// })

	t.Run("Test Orders PGCreateLink xapiversion missing", func(t *testing.T) {

		createLinkRequest := cashfree.CreateLinkRequest{
			LinkId:       strconv.Itoa(int(time.Now().Unix())),
			LinkAmount:   1,
			LinkCurrency: "INR",
			LinkPurpose:  "gosdk_testing",
			CustomerDetails: cashfree.LinkCustomerDetailsEntity{
				CustomerPhone: "9999999999",
			},
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGCreateLinkWithContext(ctx, nil, &createLinkRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGCreateLink request missing", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGCreateLinkWithContext(ctx, &XApiVersion, nil, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGCreateLink 401 status code", func(t *testing.T) {

		c := "unauthorised"
		cashfree.XClientId = &c
		flag := true

		createLinkRequest := cashfree.CreateLinkRequest{
			LinkId:       strconv.Itoa(int(time.Now().Unix())),
			LinkAmount:   1,
			LinkCurrency: "INR",
			LinkPurpose:  "gosdk_testing",
			CustomerDetails: cashfree.LinkCustomerDetailsEntity{
				CustomerPhone: "9999999999",
			},
			LinkNotify: &cashfree.LinkNotifyEntity{
				SendSms: &flag,
			},
		}
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGCreateLinkWithContext(ctx, &XApiVersion, &createLinkRequest, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #####################
	// Fetch Link
	// #####################

	t.Run("Test Orders PGFetchLink 200 status code", func(t *testing.T) {

		cashfree.XClientId = &clientId

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, httpRes, err := cashfree.PGFetchLink(&XApiVersion, linkId, &req, &idemp, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGFetchLink(nil, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGFetchLink(&XApiVersion, "invalid", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("Test Orders PGFetchLink 401 status code", func(t *testing.T) {

		c := "unathorised"
		cashfree.XClientId = &c

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGFetchLink(&XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #####################
	// Fetch Link With Context
	// #####################

	t.Run("Test Orders PGFetchLink 200 status code", func(t *testing.T) {

		cashfree.XClientId = &clientId
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, httpRes, err := cashfree.PGFetchLinkWithContext(ctx, &XApiVersion, linkId, &req, &idemp, http.DefaultClient)
		requireSuccessOrDecodeError(t, httpRes, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGFetchLinkWithContext(ctx, nil, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGFetchLinkWithContext(ctx, &XApiVersion, "invalid", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("Test Orders PGFetchLink 401 status code", func(t *testing.T) {

		c := "unathorised"
		cashfree.XClientId = &c

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGFetchLinkWithContext(ctx, &XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #####################
	// Fetch Link Orders
	// #####################

	t.Run("Test Orders PGFetchLink 200 status code", func(t *testing.T) {

		cashfree.XClientId = &clientId

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrders(&XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGLinkFetchOrders(nil, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrders(&XApiVersion, "invalid", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("Test Orders PGFetchLink 401 status code", func(t *testing.T) {

		c := "unathorised"
		cashfree.XClientId = &c

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrders(&XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #####################
	// Fetch Link With Context
	// #####################

	t.Run("Test Orders PGFetchLink 200 status code", func(t *testing.T) {

		cashfree.XClientId = &clientId
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrdersWithContext(ctx, &XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGLinkFetchOrdersWithContext(ctx, nil, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("Test Orders PGFetchLink api version missing code", func(t *testing.T) {

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrdersWithContext(ctx, &XApiVersion, "invalid", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assertStatusOneOf(t, httpRes, 400, 404)

	})

	t.Run("Test Orders PGFetchLink 401 status code", func(t *testing.T) {

		c := "unathorised"
		cashfree.XClientId = &c

		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGLinkFetchOrdersWithContext(ctx, &XApiVersion, linkId, &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ##############################
	// Cancel Link
	// ##############################

	// t.Run("Test Order PGCancelLink 200 status coed", func(t *testing.T) {

	// 	flag := true
	// 	cashfree.XClientId = &clientId

	// 	createLinkRequest := cashfree.CreateLinkRequest{
	// 		LinkId:       strconv.Itoa(int(time.Now().Unix())),
	// 		LinkAmount:   1,
	// 		LinkCurrency: "INR",
	// 		LinkPurpose:  "gosdk_testing",
	// 		CustomerDetails: cashfree.LinkCustomerDetailsEntity{
	// 			CustomerPhone: "9999999999",
	// 		},
	// 		LinkNotify: &cashfree.LinkNotifyEntity{
	// 			SendSms: &flag,
	// 		},
	// 	}
	// 	req := "511f216d-9c69-49a8-9b16-9110b103c8af"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	resp, _, _ := cashfree.PGCreateLink(&XApiVersion, &createLinkRequest, &req, &idemp, http.DefaultClient)

	// 	cashfree.XClientId = &clientId
	// 	idemp = strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(time.Now().Unix()))
	// 	newResp, httpRes, err := cashfree.PGCancelLink(&XApiVersion, resp.GetLinkId(), &req, &idemp, http.DefaultClient)

	// 	require.Nil(t, err)
	// 	require.NotNil(t, newResp)
	// 	assert.Equal(t, 200, httpRes.StatusCode)
	// })

	// t.Run("Test Order PGCancelLink api version missing", func(t *testing.T) {
	// 	cashfree.XClientId = &clientId
	// 	req := "test"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	_, _, err := cashfree.PGCancelLink(nil, "cancelLinkId", &req, &idemp, http.DefaultClient)

	// 	require.NotNil(t, err)
	// })

	// t.Run("Test Order PGCancelLink api version missing", func(t *testing.T) {
	// 	cashfree.XClientId = &clientId
	// 	req := "test"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	resp, httpRes, err := cashfree.PGCancelLink(&XApiVersion, "randomlinkId", &req, &idemp, http.DefaultClient)

	// 	require.NotNil(t, err)
	// 	require.Nil(t, resp)
	// 	assert.Equal(t, 422, httpRes.StatusCode)
	// })

	// t.Run("Test Order PGCancelLink 401 status coed", func(t *testing.T) {
	// 	c := "unauthorised"
	// 	cashfree.XClientId = &c
	// 	req := "test"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	resp, httpRes, err := cashfree.PGCancelLink(&XApiVersion, "cancelLinkId", &req, &idemp, http.DefaultClient)

	// 	require.NotNil(t, err)
	// 	require.Nil(t, resp)
	// 	assert.Equal(t, 401, httpRes.StatusCode)
	// })

	// ##############################
	// Cancel Link With Context
	// ##############################

	// t.Run("Test Order PGCancelLink 200 status coed", func(t *testing.T) {

	// 	flag := true
	// 	cashfree.XClientId = &clientId

	// 	createLinkRequest := cashfree.CreateLinkRequest{
	// 		LinkId:       strconv.Itoa(int(time.Now().Unix())),
	// 		LinkAmount:   1,
	// 		LinkCurrency: "INR",
	// 		LinkPurpose:  "gosdk_testing",
	// 		CustomerDetails: cashfree.LinkCustomerDetailsEntity{
	// 			CustomerPhone: "9999999999",
	// 		},
	// 		LinkNotify: &cashfree.LinkNotifyEntity{
	// 			SendSms: &flag,
	// 		},
	// 	}
	// 	req := "511f216d-9c69-49a8-9b16-9110b103c8af"
	// 	idemp := strconv.Itoa(int(time.Now().Unix()))
	// 	resp, _, _ := cashfree.PGCreateLink(&XApiVersion, &createLinkRequest, &req, &idemp, http.DefaultClient)

	// 	cashfree.XClientId = &clientId
	// 	idemp = strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(time.Now().Unix()))
	// 	newResp, httpRes, err := cashfree.PGCancelLinkWithContext(ctx, &XApiVersion, resp.GetLinkId(), &req, &idemp, http.DefaultClient)

	// 	require.Nil(t, err)
	// 	require.NotNil(t, newResp)
	// 	assert.Equal(t, 200, httpRes.StatusCode)
	// })

	t.Run("Test Order PGCancelLink api version missing", func(t *testing.T) {
		cashfree.XClientId = &clientId
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		_, _, err := cashfree.PGCancelLinkWithContext(ctx, nil, "cancelLinkId", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
	})

	t.Run("Test Order PGCancelLink 401 status coed", func(t *testing.T) {
		c := "unauthorised"
		cashfree.XClientId = &c
		req := "test"
		idemp := strconv.Itoa(int(time.Now().Unix()))
		resp, httpRes, err := cashfree.PGCancelLinkWithContext(ctx, &XApiVersion, "cancelLinkId", &req, &idemp, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)
	})

}
