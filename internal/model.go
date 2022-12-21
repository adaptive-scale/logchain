package internal

import (
	"time"
)

type LogStore struct {
	AppName   string
	LogLevel  string
	Timestamp time.Time
	Message   []byte
	Labels    []byte
}
