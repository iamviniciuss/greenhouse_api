package infra

type ErrorHandler interface {
	CaptureError(err error)
}
