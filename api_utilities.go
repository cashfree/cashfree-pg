/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2025-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// Execute executes the request
//  @return PGCreatePAR200Response
func (_this *Cashfree) PGCreatePAR(pARRequest *PARRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*PGCreatePAR200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PGCreatePAR200Response
	)
	if _this.XEnableErrorAnalytics == nil {
		flag := false
		_this.XEnableErrorAnalytics = &flag
	}

	if *_this.XEnableErrorAnalytics {
		SetupSentry(*_this.XEnvironment)
		defer CaptureError("PGCreatePAR")
	}

	ctx := context.Background()

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(*_this.XEnvironment)].URL

	localVarPath := localBasePath + "/pg/utilities/pars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if pARRequest == nil {
		return localVarReturnValue, nil, reportError("pARRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	if xIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-idempotency-key", xIdempotencyKey, "")
	}
	// body params
	localVarPostBody = pARRequest

parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", XApiVersion, "")

if _this.XPartnerMerchantId != nil {
	localVarHeaderParams["x-partner-merchantid"] = *_this.XPartnerMerchantId
}

if _this.XClientId != nil {
	localVarHeaderParams["x-client-id"] = *_this.XClientId
}

if _this.XClientSignature != nil {
	localVarHeaderParams["x-client-signature"] = *_this.XClientSignature
}

if _this.XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *_this.XClientSecret
}

if _this.XPartnerApiKey != nil {
	localVarHeaderParams["x-partner-apikey"] = *_this.XPartnerApiKey
}
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v IdempotencyError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimitError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// With Context
// Execute executes the request
//  @return PGCreatePAR200Response
func (_this *Cashfree) PGCreatePARWithContext(ctx context.Context, pARRequest *PARRequest,  xRequestId *string, xIdempotencyKey *string, httpClient *http.Client) (*PGCreatePAR200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PGCreatePAR200Response
	)

	if _this.XEnableErrorAnalytics == nil {
		flag := false
		_this.XEnableErrorAnalytics = &flag
	}

	if *_this.XEnableErrorAnalytics {
		SetupSentry(*_this.XEnvironment)
		defer CaptureError("PGCreatePAR")
	}

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(*_this.XEnvironment)].URL

	localVarPath := localBasePath + "/pg/utilities/pars"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if pARRequest == nil {
		return localVarReturnValue, nil, reportError("pARRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	if xIdempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-idempotency-key", xIdempotencyKey, "")
	}
	// body params
	localVarPostBody = pARRequest

parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", XApiVersion, "")

if _this.XPartnerMerchantId != nil {
	localVarHeaderParams["x-partner-merchantid"] = *_this.XPartnerMerchantId
}

if _this.XClientId != nil {
	localVarHeaderParams["x-client-id"] = *_this.XClientId
}

if _this.XClientSignature != nil {
	localVarHeaderParams["x-client-signature"] = *_this.XClientSignature
}

if _this.XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *_this.XClientSecret
}

if _this.XPartnerApiKey != nil {
	localVarHeaderParams["x-partner-apikey"] = *_this.XPartnerApiKey
}

	localVarHeaderParams["x-sdk-platform"] = "gosdk-5.0.5"
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v IdempotencyError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimitError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
// With Context


