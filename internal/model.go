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

type MetricStore struct {
	MetricGroup string
	MetricName  string
	Value       float64
	Timestamp   time.Time
	Labels      []byte
}
