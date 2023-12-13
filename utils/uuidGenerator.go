package utils

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}
