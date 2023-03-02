package uuid

import (
	"log"

	"github.com/google/uuid"
)

// UUID generates uniq uuid.
func UUID() string {
	log.Printf("try to get uuid")
	return uuid.New().String()
}
