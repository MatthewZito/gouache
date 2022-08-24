package services

import (
	"fmt"
	"log"
	"os"
)

// Logger specifies the necessary methods to implement a LoggerClient.
type Logger interface {
	Printf(string, ...interface{})
}

// LoggerClient holds a logger object to be used across the service.
type LoggerClient struct {
	l Logger
}

// NewLogger initializes a new LoggerClient in the context of a given locale.
func NewLogger(locale string) *LoggerClient {
	return &LoggerClient{l: log.New(os.Stdout, fmt.Sprintf("[%s] ", locale), log.LstdFlags)}
}

// Logf allows a LoggerClient to log variadically.
func (c *LoggerClient) Logf(format string, a ...interface{}) {
	if c.l != nil {
		c.l.Printf(format, a...)
	}
}
