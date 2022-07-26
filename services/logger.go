package services

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Printf(string, ...interface{})
}

type LoggerClient struct {
	l Logger
}

func NewLogger(locale string) *LoggerClient {
	return &LoggerClient{l: log.New(os.Stdout, fmt.Sprintf("[%s] ", locale), log.LstdFlags)}
}

func (c *LoggerClient) Logf(format string, a ...interface{}) {
	if c.l != nil {
		c.l.Printf(format, a...)
	}
}
