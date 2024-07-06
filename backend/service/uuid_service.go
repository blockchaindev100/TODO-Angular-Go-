package service

import "github.com/google/uuid"

func GenerateUUID() string {
	uuid := uuid.New()
	uuidString := uuid.String()

	return uuidString
}
