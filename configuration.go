/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2025-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type CFEnvironment int

const (
	SANDBOX    CFEnvironment = 0
	PRODUCTION CFEnvironment = 1
)

var (
	XApiVersion = "2025-01-01"
)

type Cashfree struct {
	XPartnerMerchantId *string
	XClientId *string
	XClientSignature *string
	XClientSecret *string
	XPartnerApiKey *string
	XEnvironment *CFEnvironment
	XEnableErrorAnalytics *bool
}

type PGWebhookEvent struct {
	Type   string
	Raw    string
	Object interface{}
}

// Execute executes the request
// @return PGWebhookEvent
func (_this *Cashfree) PGVerifyWebhookSignature(signature string, rawBody string, timestamp string) (*PGWebhookEvent, error) {
	if _this.XClientSecret == nil {
		return nil, errors.New("set your secret key, cashfree.XClientSecret := *secretKey")
	}
	signatureString := timestamp + rawBody
	hmacInstance := hmac.New(sha256.New, []byte(*_this.XClientSecret))
	hmacInstance.Write([]byte(signatureString))
	bytesData := hmacInstance.Sum(nil)
	generatedSignature := base64.StdEncoding.EncodeToString(bytesData)
	if generatedSignature == signature {
		var object interface{}
		err := json.Unmarshal([]byte(rawBody), &object)
		if err != nil {
			return nil, errors.New("something went wrong when unmarshalling raw body")
		}
		if objectAsMapInterface, ok := object.(map[string]interface{}); ok {
			if webhookType, ok := objectAsMapInterface["type"].(string); ok {
				return &PGWebhookEvent{Type: webhookType, Raw: rawBody, Object: object}, nil
			}
		}
		return &PGWebhookEvent{Type: "", Raw: rawBody, Object: object}, nil
	}
	return nil, errors.New("generated signature and received signature did not match")
}

func SetupSentry(environment CFEnvironment) {
	env := "sandbox"
	if environment == PRODUCTION {
		env = "production"
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://a5d54ec768b044ffb1e5c948b9b929e1@o330525.ingest.sentry.io/4504762466304000",
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		EnableTracing:    true,
		Environment:      env,
		Release:          "5.0.5",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			delete(event.Contexts, "device")
			delete(event.Contexts, "os")
			if len(event.Exception) > 0 && event.Exception[0].Stacktrace != nil {
				frames := event.Exception[0].Stacktrace.Frames
				for _, v := range frames {
					if strings.Contains(v.Filename, "cashfree-pg") {
						return event
					}
				}
			}
			return nil
		},
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func CaptureError(api string) {
	r := recover()
	if r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		sentry.CaptureMessage(api + "::" + err.Error())
		sentry.Flush(time.Second * 2)
	}
}

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/5.0.5/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "https://sandbox.cashfree.com/pg",
				Description: "Sandbox server",
			},
			{
				URL: "https://api.cashfree.com/pg",
				Description: "Production server",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
