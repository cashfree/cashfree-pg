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

func Test_cashfree_pg_offers(t *testing.T) {
	clientId := os.Getenv("clientid")
	XClientSecret := os.Getenv("clientsecret")
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &XClientSecret
	cashfree.XEnvironment = cashfree.SANDBOX
	XApiVersion := "2022-09-01"
	offerId := "682beb8e-6a82-4c29-8f37-c2624b1febc9"
	ctx := context.Background()

	t.Run("PGCreateOffer should give status 200", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		resp, httpRes, err := cashfree.PGCreateOffer(&XApiVersion, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOffer should fail when xApiVersion is missing", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		_, _, err := cashfree.PGCreateOffer(nil, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGCreateOffer should fail when request is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGCreateOffer(&XApiVersion, nil, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGCreateOffer should give 401 for invalid credentials", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		resp, httpRes, err := cashfree.PGCreateOffer(&XApiVersion, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// ############################
	// Create Offer with context
	// ############################

	t.Run("PGCreateOfferWithContext should give status 200", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		resp, httpRes, err := cashfree.PGCreateOfferWithContext(ctx, &XApiVersion, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("PGCreateOfferWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		_, _, err := cashfree.PGCreateOfferWithContext(ctx, nil, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGCreateOfferWithContext should fail when request is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGCreateOfferWithContext(ctx, &XApiVersion, nil, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGCreateOfferWithContext should give 401 for invalid credentials", func(t *testing.T) {

		var minAmount float32 = 200.0
		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		createOfferRequest := cashfree.CreateOfferRequest{
			OfferMeta: cashfree.OfferMeta{
				OfferTitle:       "Go-Sdk-test",
				OfferDescription: "Testing",
				OfferCode:        "Go-Sdk-50",
				OfferStartTime:   "2035-03-21T08:09:51Z",
				OfferEndTime:     "2045-03-21T08:09:51Z",
			},
			OfferTnc: cashfree.OfferTnc{
				OfferTncType:  "text",
				OfferTncValue: "cashfree",
			},
			OfferDetails: cashfree.OfferDetails{
				OfferType: "NO_COST_EMI",
				DiscountDetails: &cashfree.DiscountDetails{
					DiscountType:      "flat",
					DiscountValue:     "10",
					MaxDiscountAmount: "10",
				},
			},
			OfferValidations: cashfree.OfferValidations{
				MinAmount:  &minAmount,
				MaxAllowed: 1000.0,
				PaymentMethod: cashfree.OfferValidationsPaymentMethod{
					OfferEMI: &cashfree.OfferEMI{
						Emi: cashfree.EMIOffer{
							Type:    "cardless_emi",
							Issuer:  "hdfc bank",
							Tenures: []int32{3, 6},
						},
					},
				},
			},
		}

		resp, httpRes, err := cashfree.PGCreateOfferWithContext(ctx, &XApiVersion, &createOfferRequest, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #########################
	// Get Offer
	// #########################

	t.Run("PGFetchOffer should give status 200", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOffer(&XApiVersion, offerId, &req, &idem, http.DefaultClient)
		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGFetchOffer should fail when xApiVersion is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGFetchOffer(nil, offerId, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGFetchOffer should fail when request is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGFetchOffer(&XApiVersion, "", &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOffer should give 401 for invalid credentials", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGFetchOffer(&XApiVersion, offerId, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})

	// #########################
	// Get Offer With Context
	// #########################

	t.Run("PGFetchOfferWithContext should give status 200", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		cashfree.XClientId = &clientId

		resp, httpRes, err := cashfree.PGFetchOfferWithContext(ctx, &XApiVersion, offerId, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 400, httpRes.StatusCode)

	})

	t.Run("PGFetchOfferWithContext should fail when xApiVersion is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		_, _, err := cashfree.PGFetchOfferWithContext(ctx, nil, offerId, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)

	})

	t.Run("PGFetchOfferWithContext should fail when request is missing", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))

		resp, httpRes, err := cashfree.PGFetchOfferWithContext(ctx, &XApiVersion, "", &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 404, httpRes.StatusCode)

	})

	t.Run("PGFetchOfferWithContext should give 401 for invalid credentials", func(t *testing.T) {

		req := "test"
		idem := strconv.Itoa(int(time.Now().Unix()))
		c := "unauthorised"
		cashfree.XClientId = &c

		resp, httpRes, err := cashfree.PGFetchOfferWithContext(ctx, &XApiVersion, offerId, &req, &idem, http.DefaultClient)

		require.NotNil(t, err)
		require.Nil(t, resp)
		assert.Equal(t, 401, httpRes.StatusCode)

	})
}
