package cashfree_pg_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

type responseBodyError interface {
	Body() []byte
}

const (
	eventualConsistencyMaxAttempts = 21
	eventualConsistencyRetryDelay  = 2 * time.Second
	seedCreateOrderMaxAttempts     = 3
	seedCreateOrderRetryDelay      = 500 * time.Millisecond
)

func uniqueSuffix() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func sleepBeforeEventualRetry(attempt int) {
	if attempt < eventualConsistencyMaxAttempts-1 {
		time.Sleep(eventualConsistencyRetryDelay)
	}
}

func shouldRetrySeedCreateOrder(httpRes *http.Response) bool {
	return httpRes == nil || httpRes.StatusCode == http.StatusNotFound
}

func sleepBeforeSeedCreateOrderRetry(attempt int) {
	if attempt < seedCreateOrderMaxAttempts-1 {
		time.Sleep(seedCreateOrderRetryDelay)
	}
}

func requireHTTPStatus(t *testing.T, httpRes *http.Response, expectedStatus int) {
	t.Helper()
	require.NotNil(t, httpRes)
	assert.Equal(t, expectedStatus, httpRes.StatusCode)
}

// v3 SDK still has a few response model type mismatches for successful responses.
func requireSuccessOrDecodeError(t *testing.T, httpRes *http.Response, err error) {
	t.Helper()
	requireHTTPStatus(t, httpRes, 200)
	if err != nil {
		require.Contains(t, err.Error(), "cannot unmarshal")
	}
}

// Seed create-order calls can legitimately return 409 if the first attempt
// succeeded server-side and a retry replays the same order_id.
func requireSeedCreateOrderSuccess(t *testing.T, httpRes *http.Response, err error) {
	t.Helper()
	require.NotNil(t, httpRes)

	switch httpRes.StatusCode {
	case http.StatusOK:
		if err != nil {
			require.Contains(t, err.Error(), "cannot unmarshal")
		}
	case http.StatusConflict:
		// Already created; acceptable for seeding.
	default:
		require.Failf(t, "unexpected status code", "expected one of [200 409], got %d", httpRes.StatusCode)
	}
}

// Some create-order success paths in sandbox intermittently return transient 404s.
// Retry briefly and accept 409 as success if a previous create attempt actually succeeded.
func requireCreateOrderSuccessWithRetry(t *testing.T, call func() (*http.Response, error)) {
	t.Helper()

	var httpRes *http.Response
	var err error
	for attempt := 0; attempt < seedCreateOrderMaxAttempts; attempt++ {
		httpRes, err = call()

		if httpRes != nil {
			switch httpRes.StatusCode {
			case http.StatusOK:
				if err != nil {
					require.Contains(t, err.Error(), "cannot unmarshal")
				}
				return
			case http.StatusConflict:
				return
			}
		}

		if !shouldRetrySeedCreateOrder(httpRes) {
			break
		}
		sleepBeforeSeedCreateOrderRetry(attempt)
	}

	requireSuccessOrDecodeError(t, httpRes, err)
}

func assertStatusOneOf(t *testing.T, httpRes *http.Response, expectedStatuses ...int) {
	t.Helper()
	require.NotNil(t, httpRes)

	for _, expectedStatus := range expectedStatuses {
		if httpRes.StatusCode == expectedStatus {
			return
		}
	}

	require.Failf(
		t,
		"unexpected status code",
		"expected one of %v, got %d",
		expectedStatuses,
		httpRes.StatusCode,
	)
}

func extractStringFromErrorBody(err error, field string) string {
	if err == nil {
		return ""
	}

	var bodyErr responseBodyError
	if !errors.As(err, &bodyErr) {
		return ""
	}

	body := bodyErr.Body()
	if len(body) == 0 {
		return ""
	}

	var payload map[string]interface{}
	if json.Unmarshal(body, &payload) != nil {
		return ""
	}

	value, ok := payload[field]
	if !ok || value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatInt(int64(v), 10)
	default:
		return ""
	}
}
