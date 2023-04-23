package infra

import "fmt"

type DatabaseError struct {
	Message string `json:"message,omitempty"`
}

func (de *DatabaseError) Error() string {
	return fmt.Sprintf(de.Message)
}
