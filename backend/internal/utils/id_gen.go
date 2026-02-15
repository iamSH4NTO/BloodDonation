package utils

import (
	"fmt"
	"time"
)

// GenerateUniqueID generates a permanent string ID in the format BD-XXXXXX
func GenerateUniqueID() string {
	return fmt.Sprintf("BD-%06d", time.Now().UnixNano()%1000000)
}
