package infra

import (
	"fmt"
	"os"

	sentry "github.com/getsentry/sentry-go"
)

type SentryErrorHandler struct {
}

func NewSentryErrorHandler() *SentryErrorHandler {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DNS"),
		Environment: os.Getenv("ENVIRONMENT"),
	})

	if err != nil {
		panic(err)
	}

	return &SentryErrorHandler{}
}

func (eh *SentryErrorHandler) CaptureError(err error) {
	fmt.Println("Sentry CaptureError:", err.Error())
	sentry.CaptureMessage(err.Error())
}
