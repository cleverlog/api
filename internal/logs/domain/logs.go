package domain

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ServiceName string
	Level       string
	SpanID      uuid.UUID
	Timestamp   time.Time
	Source      string
	Message     string
}
