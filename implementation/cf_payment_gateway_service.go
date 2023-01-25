package cashfree_pg_sdk_go

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type CFEnvironment int

const (
	SANDBOX    CFEnvironment = 0
	PRODUCTION CFEnvironment = 1
)

type CFConfig struct {
	Environment  *CFEnvironment
	ClientId     *string
	ClientSecret *string
	ApiVersion   *string
	Timeout      *time.Duration
	Proxy        *string
}

type CFHeader struct {
	RequestID      *string
	IdempotencyKey *string
}

type CFResponseHeader struct {
	ApiVersion          *string
	RateLimit           *string
	RateLimitRemaining  *string
	RateLimitRetry      *string
	RateLimitType       *string
	RequestID           *string
	IdempotencyKey      *string
	IdempotencyReplayed *string
}

// Method for internal purpose
func checkSession(session *CFConfig) *CFError {
	if session.Environment == nil {
		msg := "environment is missing"
		code := "environment_missing"
		return &CFError{
			Message: &msg,
			Code:    &code,
		}
	}
	if session.ClientId == nil {
		msg := "xClientId is missing"
		code := "xClientId_missing"
		return &CFError{
			Message: &msg,
			Code:    &code,
		}
	}
	if session.ClientSecret == nil {
		msg := "xClientSecret is missing"
		code := "xClientSecret_missing"
		return &CFError{
			Message: &msg,
			Code:    &code,
		}
	}
	if session.ApiVersion == nil {
		msg := "xApiVersion is missing"
		code := "xApiVersion_missing"
		return &CFError{
			Message: &msg,
			Code:    &code,
		}
	}
	return nil
}

// Method for internal purpose
func setEnvironmentInternal(session *CFConfig) (context.Context, *CFError) {
	var c context.Context
	if *session.Environment == SANDBOX {
		c = context.WithValue(context.Background(), ContextServerIndex, 0)
	} else if *session.Environment == PRODUCTION {
		c = context.WithValue(context.Background(), ContextServerIndex, 1)
	} else {
		msg := "value for environment has to be PRODUCTION or SANDBOX"
		code := "environment_invalid"
		return nil, &CFError{
			Message: &msg,
			Code:    &code,
		}
	}
	return c, nil
}

// Prepare CFError and send
func getError(response *http.Response) CFError {
	if response != nil {
		var cfError CFError
		errorData, e := ioutil.ReadAll(response.Body)
		if e != nil {
			msg := e.Error()
			code := "something_went_wrong"
			return CFError{
				Message: &msg,
				Code:    &code,
			}
		}
		e = json.Unmarshal(errorData, &cfError)
		if e != nil {
			msg := e.Error()
			code := "something_went_wrong"
			return CFError{
				Message: &msg,
				Code:    &code,
			}
		}
		return cfError
	}
	msg := "Empty Response Received"
	code := "something_went_wrong"
	return CFError{
		Message: &msg,
		Code:    &code,
	}
}

func getHeader(r *http.Response) CFResponseHeader {
	if r != nil {
		xApiVersion := r.Header.Get("x-api-version")
		xRateLimit := r.Header.Get("x-ratelimit-limit")
		xRateLimitRemaining := r.Header.Get("x-ratelimit-remaining")
		xRateLimitRetry := r.Header.Get("x-ratelimit-retry")
		xRateLimitType := r.Header.Get("x-ratelimit-type")
		XRequestId := r.Header.Get("x-request-id")
		xIdempotencyKey := r.Header.Get("x-idempotency-key")
		xIdempotencyReplayed := r.Header.Get("x-idempotency-replayed")
		responseHeader := CFResponseHeader{
			ApiVersion:          &xApiVersion,
			RateLimit:           &xRateLimit,
			RateLimitRemaining:  &xRateLimitRemaining,
			RateLimitRetry:      &xRateLimitRetry,
			RateLimitType:       &xRateLimitType,
			RequestID:           &XRequestId,
			IdempotencyKey:      &xIdempotencyKey,
			IdempotencyReplayed: &xIdempotencyReplayed,
		}
		return responseHeader
	}
	return CFResponseHeader{}
}

// Create Order
func CreateOrder(session *CFConfig, header *CFHeader, order CFOrderRequest) (*CFOrder, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfigurationWithTimeoutAndProxy(session.Timeout, session.Proxy))
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(OrdersApiService)
	apiCreateOrderRequest := service.CreateOrder(ctx)
	apiCreateOrderRequest = apiCreateOrderRequest.XClientId(*session.ClientId)
	apiCreateOrderRequest = apiCreateOrderRequest.XClientSecret(*session.ClientSecret)
	apiCreateOrderRequest = apiCreateOrderRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiCreateOrderRequest = apiCreateOrderRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiCreateOrderRequest = apiCreateOrderRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	apiCreateOrderRequest = apiCreateOrderRequest.CFOrderRequest(order)
	orderResponse, response, err := apiClient.OrdersApi.CreateOrderExecute(apiCreateOrderRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)

		return orderResponse, &responseHeader, &cfError
	}
	return orderResponse, &responseHeader, nil
}

// Order Pay UPI
func OrderPay(session *CFConfig, header *CFHeader, orderPay CFOrderPayRequest) (*CFOrderPayResponse, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(OrdersApiService)
	orderPayRequest := service.OrderPay(ctx)
	orderPayRequest = orderPayRequest.CFOrderPayRequest(orderPay)
	orderPayRequest = orderPayRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			orderPayRequest = orderPayRequest.XRequestId(*header.RequestID)
		}
	}
	orderPayResponse, response, err := apiClient.OrdersApi.OrderPayExecute(orderPayRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return orderPayResponse, &responseHeader, &cfError
	}
	return orderPayResponse, &responseHeader, nil
}

func OrderPaySessionId(session *CFConfig, header *CFHeader, orderPaySessionId CFOrderPaySessionIdRequest) (*CFOrderPayResponse, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(OrdersApiService)
	orderPayRequest := service.OrderPay(ctx)
	orderPayRequest = orderPayRequest.CFOrderPaySessionIdRequest(orderPaySessionId)
	orderPayRequest = orderPayRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			orderPayRequest = orderPayRequest.XRequestId(*header.RequestID)
		}
	}
	orderPayResponse, response, err := apiClient.OrdersApi.OrderPayExecuteSessionId(orderPayRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return orderPayResponse, &responseHeader, &cfError
	}
	return orderPayResponse, &responseHeader, nil
}

// Get Order API
func GetOrder(session *CFConfig, header *CFHeader, orderId string) (*CFOrder, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(OrdersApiService)
	getOrderRequest := service.GetOrder(ctx, orderId)
	getOrderRequest = getOrderRequest.XClientId(*session.ClientId)
	getOrderRequest = getOrderRequest.XClientSecret(*session.ClientSecret)
	getOrderRequest = getOrderRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			getOrderRequest = getOrderRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			getOrderRequest = getOrderRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	order, response, err := apiClient.OrdersApi.GetOrderExecute(getOrderRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return order, &responseHeader, &cfError
	}
	return order, &responseHeader, nil
}

// Create Link API
func CreateLink(session *CFConfig, header *CFHeader, cFLinkRequest CFLinkRequest) (*CFLink, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(PaymentLinksApiService)
	createPaymentLinkRequest := service.CreatePaymentLink(ctx)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientId(*session.ClientId)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientSecret(*session.ClientSecret)
	createPaymentLinkRequest = createPaymentLinkRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	createPaymentLinkRequest = createPaymentLinkRequest.CFLinkRequest(cFLinkRequest)

	link, response, err := apiClient.PaymentLinksApi.CreatePaymentLinkExecute(createPaymentLinkRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return link, &responseHeader, &cfError
	}
	return link, &responseHeader, nil
}

// Get Link by LinkId API
func GetLinkByLinkID(session *CFConfig, header *CFHeader, linkId string) (*CFLink, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(PaymentLinksApiService)
	createPaymentLinkRequest := service.GetPaymentLinkDetails(ctx, linkId)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientId(*session.ClientId)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientSecret(*session.ClientSecret)
	createPaymentLinkRequest = createPaymentLinkRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	link, response, err := apiClient.PaymentLinksApi.GetPaymentLinkDetailsExecute(createPaymentLinkRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return link, &responseHeader, &cfError
	}
	return link, &responseHeader, nil
}

// Get Orders by Link Id
func GetOrdersByLinkID(session *CFConfig, header *CFHeader, linkId string) ([]CFLinkOrders, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(PaymentLinksApiService)
	createPaymentLinkRequest := service.GetPaymentLinkOrders(ctx, linkId)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientId(*session.ClientId)
	createPaymentLinkRequest = createPaymentLinkRequest.XClientSecret(*session.ClientSecret)
	createPaymentLinkRequest = createPaymentLinkRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			createPaymentLinkRequest = createPaymentLinkRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	links, response, err := apiClient.PaymentLinksApi.GetPaymentLinkOrdersExecute(createPaymentLinkRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return links, &responseHeader, &cfError
	}
	return links, &responseHeader, nil
}

// Get Payments for Order API
func GetPaymentsForOrder(session *CFConfig, header *CFHeader, orderId string) ([]CFPaymentsEntity, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(PaymentsApiService)
	apiGetPaymentsForOrder := service.GetPaymentsfororder(ctx, orderId)

	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XClientId(*session.ClientId)
	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XClientSecret(*session.ClientSecret)
	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiGetPaymentsForOrder = apiGetPaymentsForOrder.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiGetPaymentsForOrder = apiGetPaymentsForOrder.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	paymentEntities, response, err := apiClient.PaymentsApi.GetPaymentsfororderExecute(apiGetPaymentsForOrder)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return paymentEntities, &responseHeader, &cfError
	}
	return paymentEntities, &responseHeader, nil
}

// Get Payment by Payment ID
func GetPaymentByPaymentID(session *CFConfig, header *CFHeader, orderId string, cfPaymentId int32) (*CFPaymentsEntity, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(PaymentsApiService)
	apiGetPaymentsForOrder := service.GetPaymentbyId(ctx, orderId, cfPaymentId)

	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XClientId(*session.ClientId)
	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XClientSecret(*session.ClientSecret)
	apiGetPaymentsForOrder = apiGetPaymentsForOrder.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiGetPaymentsForOrder = apiGetPaymentsForOrder.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiGetPaymentsForOrder = apiGetPaymentsForOrder.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	paymentEntity, response, err := apiClient.PaymentsApi.GetPaymentbyIdExecute(apiGetPaymentsForOrder)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return paymentEntity, &responseHeader, &cfError
	}
	return paymentEntity, &responseHeader, nil
}

// Initiate Refund
func InitiateRefund(session *CFConfig, header *CFHeader, refundRquest CFRefundRequest, orderId string) (*CFRefund, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(RefundsApiService)
	initiateRefundRequest := service.Createrefund(ctx, orderId)
	initiateRefundRequest = initiateRefundRequest.XClientId(*session.ClientId)
	initiateRefundRequest = initiateRefundRequest.XClientSecret(*session.ClientSecret)
	initiateRefundRequest = initiateRefundRequest.XApiVersion(*session.ApiVersion)
	initiateRefundRequest = initiateRefundRequest.CFRefundRequest(refundRquest)
	if header != nil {
		if header.RequestID != nil {
			initiateRefundRequest = initiateRefundRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			initiateRefundRequest = initiateRefundRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	refund, response, err := apiClient.RefundsApi.CreaterefundExecute(initiateRefundRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return refund, &responseHeader, &cfError
	}
	return refund, &responseHeader, nil
}

// Fetch Refund Data
func FetchRefundData(session *CFConfig, header *CFHeader, refundId string, orderId string) (*CFRefund, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(RefundsApiService)
	fetchRefundRequest := service.GetRefund(ctx, orderId, refundId)
	fetchRefundRequest = fetchRefundRequest.XClientId(*session.ClientId)
	fetchRefundRequest = fetchRefundRequest.XClientSecret(*session.ClientSecret)
	fetchRefundRequest = fetchRefundRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			fetchRefundRequest = fetchRefundRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			fetchRefundRequest = fetchRefundRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	refund, response, err := apiClient.RefundsApi.GetRefundExecute(fetchRefundRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return refund, &responseHeader, &cfError
	}
	return refund, &responseHeader, nil
}

// Pre Authorization Request
func SetPreAuthorizationRequest(session *CFConfig, header *CFHeader, authorizationRequest CFAuthorizationRequest, orderId string) (*CFPaymentsEntity, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(OrdersApiService)
	preAuthorizationRequest := service.Preauthorization(ctx, orderId)
	preAuthorizationRequest = preAuthorizationRequest.XClientId(*session.ClientId)
	preAuthorizationRequest = preAuthorizationRequest.XClientSecret(*session.ClientSecret)
	preAuthorizationRequest = preAuthorizationRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			preAuthorizationRequest = preAuthorizationRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			preAuthorizationRequest = preAuthorizationRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	preAuthorizationRequest = preAuthorizationRequest.CFAuthorizationRequest(authorizationRequest)
	paymentEntity, response, err := apiClient.OrdersApi.PreauthorizationExecute(preAuthorizationRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return paymentEntity, &responseHeader, &cfError
	}
	return paymentEntity, &responseHeader, nil
}

// Get Settlements
func Getsettlements(session *CFConfig, header *CFHeader, orderId string) (*CFSettlementsEntity, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(SettlementsApiService)
	settlementRequest := service.Getsettlements(ctx, orderId)
	settlementRequest = settlementRequest.XClientId(*session.ClientId)
	settlementRequest = settlementRequest.XClientSecret(*session.ClientSecret)
	settlementRequest = settlementRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			settlementRequest = settlementRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			settlementRequest = settlementRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	settlements, response, err := apiClient.SettlementsApi.GetsettlementsExecute(settlementRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return settlements, &responseHeader, &cfError
	}
	return settlements, &responseHeader, nil
}

// Get Instruments By Customer Id
func GetInstrumentsByCustomerID(session *CFConfig, header *CFHeader, customerId string, instrumentType string) ([]CFFetchAllSavedInstruments, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(TokenVaultApiService)
	apiFetchSavedInstrumentRequest := service.FetchAllSavedInstruments(ctx, customerId)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XClientId(*session.ClientId)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XClientSecret(*session.ClientSecret)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.InstrumentType(instrumentType)

	instruments, response, err := apiClient.TokenVaultApi.FetchAllSavedInstrumentsExecute(apiFetchSavedInstrumentRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return instruments, &responseHeader, &cfError
	}
	return instruments, &responseHeader, nil
}

// Get Instruments By Instrument Id
func GetInstrumentsByInstrumentID(session *CFConfig, header *CFHeader, customerId string, instrument_id string) (*CFFetchAllSavedInstruments, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(TokenVaultApiService)
	apiFetchSavedInstrumentRequest := service.FetchSpecificSavedInstrument(ctx, customerId, instrument_id)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XClientId(*session.ClientId)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XClientSecret(*session.ClientSecret)
	apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiFetchSavedInstrumentRequest = apiFetchSavedInstrumentRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	instrument, response, err := apiClient.TokenVaultApi.FetchSpecificSavedInstrumentExecute(apiFetchSavedInstrumentRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return instrument, &responseHeader, &cfError
	}
	return instrument, &responseHeader, nil
}

// Get Instrument CryptoGram By Instrument Id
func GetInstrumentCryptogramByInstrumentID(session *CFConfig, header *CFHeader, customerId string, instrument_id string) (*CFCryptogram, *CFResponseHeader, *CFError) {
	apiClient := *NewAPIClient(NewConfiguration())
	e := checkSession(session)
	if e != nil {
		return nil, nil, e
	}
	ctx, invalidEnvironmentError := setEnvironmentInternal(session)
	if invalidEnvironmentError != nil {
		return nil, nil, invalidEnvironmentError
	}
	service := new(TokenVaultApiService)
	apiFetchCryptogramRequest := service.FetchCryptogram(ctx, customerId, instrument_id)
	apiFetchCryptogramRequest = apiFetchCryptogramRequest.XClientId(*session.ClientId)
	apiFetchCryptogramRequest = apiFetchCryptogramRequest.XClientSecret(*session.ClientSecret)
	apiFetchCryptogramRequest = apiFetchCryptogramRequest.XApiVersion(*session.ApiVersion)
	if header != nil {
		if header.RequestID != nil {
			apiFetchCryptogramRequest = apiFetchCryptogramRequest.XRequestId(*header.RequestID)
		}
		if header.IdempotencyKey != nil {
			apiFetchCryptogramRequest = apiFetchCryptogramRequest.XIdempotencyKey(*header.IdempotencyKey)
		}
	}
	cryptogram, response, err := apiClient.TokenVaultApi.FetchCryptogramExecute(apiFetchCryptogramRequest)
	responseHeader := getHeader(response)
	if err != nil {
		cfError := getError(response)
		return cryptogram, &responseHeader, &cfError
	}
	return cryptogram, &responseHeader, nil
}
