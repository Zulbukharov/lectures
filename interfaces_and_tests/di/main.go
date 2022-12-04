package main

import (
	"fmt"
	"log"
)

type SuperLogger interface {
	Log(Logger)
}

type Logger interface {
	Log(string)
}

type SimpleLogger struct {
}

func (s SimpleLogger) Log(message string) {
	fmt.Printf("%v\n", message)
}

type CustomLogger struct{}

func (s CustomLogger) Log(message string) {
	log.Printf("%v\n", message)
}

type Service struct {
	logger *SimpleLogger
}

func NewService(logger *SimpleLogger) *Service {
	return &Service{
		logger: logger,
	}
}

func main() {
	customLogger := &CustomLogger{}
	service := NewService(customLogger)
	service.logger.Log("hello")

}
