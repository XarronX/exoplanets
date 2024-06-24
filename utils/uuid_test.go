package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()

	assert.NotEqual(t, uuid1, uuid2)
}

func TestNilUUID(t *testing.T) {
	nilUUID := NilUUID()

	assert.Equal(t, uuid.Nil, nilUUID)
}

func TestParseUUID_Valid(t *testing.T) {
	validUUIDStr := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	expectedUUID, _ := uuid.Parse(validUUIDStr)

	parsedUUID, err := ParseUUID(validUUIDStr)

	assert.NoError(t, err)
	assert.Equal(t, expectedUUID, parsedUUID)
}

func TestParseUUID_Invalid(t *testing.T) {
	invalidUUIDStr := "invalid-uuid-string"

	_, err := ParseUUID(invalidUUIDStr)
	assert.Error(t, err)
}
