package errorhandler

import "fmt"

type err struct {
	Message    string
	Code       int
	Underlying string
}

func New(message string, code int, underlying string) *err {
	return &err{
		Message:    message,
		Code:       code,
		Underlying: underlying,
	}
}

func (e *err) Error() string {
	return fmt.Sprintf("Error: %s\nCode: %d", e.Message, e.Code)
}

func (e *err) DetailedError() string {
	return fmt.Sprintf("Error: %s\nCode: %d\n\nUnderlying error: %+v", e.Message, e.Code, e.Underlying)
}
