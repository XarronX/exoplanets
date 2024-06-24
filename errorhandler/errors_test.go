package errorhandler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New("test message", 500, "underlying error")
	assert.NotNil(t, err, "New() should not return nil")

	assert.Equal(t, "test message", err.Message, "Message field should be 'test message'")
	assert.Equal(t, 500, err.Code, "Code field should be 500")
	assert.Equal(t, "underlying error", err.Underlying, "Underlying field should be 'underlying error'")
}

func TestError(t *testing.T) {
	err := &err{
		Message: "test message",
		Code:    500,
	}

	expectedErrorString := fmt.Sprintf("Error: %s\nCode: %d", err.Message, err.Code)
	assert.Equal(t, expectedErrorString, err.Error(), "Error() output should match expected string")
}

func TestDetailedError(t *testing.T) {
	err := &err{
		Message:    "test message",
		Code:       500,
		Underlying: "underlying error",
	}

	expectedDetailedErrorString := fmt.Sprintf("Error: %s\nCode: %d\n\nUnderlying error: %+v", err.Message, err.Code, err.Underlying)
	assert.Equal(t, expectedDetailedErrorString, err.DetailedError(), "DetailedError() output should match expected string")
}
