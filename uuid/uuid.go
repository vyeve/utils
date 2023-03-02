package uuid

import (
	"github.com/google/uuid"
)

// UUID generates uniq uuid.
func UUID() string {
	return uuid.New().String()
}
