package infra

import (
	"fmt"

	mockTestify "github.com/stretchr/testify/mock"
)

type ErrorHandlerMock struct {
	mockTestify.Mock
}

func NewErrorHandlerMock() *ErrorHandlerMock {
	return &ErrorHandlerMock{}
}

func (cn *ErrorHandlerMock) CaptureError(err error) {
	fmt.Println(err)
}
