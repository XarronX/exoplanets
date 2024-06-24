package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func NilUUID() uuid.UUID {
	return uuid.Nil
}

func ParseUUID(uuidStr string) (uuid.UUID, error) {
	u, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, err
	}

	return u, nil
}
