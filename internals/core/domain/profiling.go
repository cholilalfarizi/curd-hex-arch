// internals\core\domain\profiling.go

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profiling struct {
	ID        uuid.UUID `json:"id"`
	Method   string    `json:"method"`
	Duration  int64     `json:"duration"`
	Timestamp time.Time `json:"timestamp"`
}