package utils

import (
	"fmt"
	"time"
)

// GenerateUniqueID generates a permanent string ID in the format BD-XXXXXX
func GenerateUniqueID() string {
	// Use UnixNano and a random seed or more entropy if possible.
	// For simplicity and better uniqueness than just % 1000000:
	now := time.Now()
	return fmt.Sprintf("BD-%06d", (now.UnixNano()/1000)%1000000)
}
