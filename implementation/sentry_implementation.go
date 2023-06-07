package cashfree_pg_sdk_go

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func SetupSentry(environment string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://a5d54ec768b044ffb1e5c948b9b929e1@o330525.ingest.sentry.io/4504762466304000",
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		EnableTracing:    true,
		Environment:      environment,
		Release:          "2.0.2",
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
