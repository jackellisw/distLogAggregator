package models

import (
	"time"
)

type LogEntry struct {
	TimeStamp   time.Time
	LogLevel    string
	ServiceName string
	Message     string
}
