package utils

import "github.com/google/uuid"

func ParseUUIDString(uuidString string) (uuid.UUID, error) {
	uuidReal, err := uuid.Parse(uuidString)
	return uuidReal, err
}
