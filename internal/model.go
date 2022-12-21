package internal

import (
	"time"
)

type LogStore struct {
	AppName   string
	LogLevel  string
	Timestamp time.Time
	LogLine   []byte
	Labels    []byte
}
