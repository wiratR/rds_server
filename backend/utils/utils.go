package utils

import (
	"log"
	"strings"

	"github.com/gofrs/uuid"
)

func GenerateAccountToken() string {
	// Generate a new UUID
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return strings.ToUpper(uuid.String())
}
