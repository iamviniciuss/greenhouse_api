package infra

import "fmt"

type ValidationError struct {
	Message string `json:"message,omitempty"`
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf(ve.Message)
}
