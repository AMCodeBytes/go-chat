package utilities

import (
	"errors"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return "", errors.New("failed to generate uuid")
	}

	return uuid.String(), nil
}
