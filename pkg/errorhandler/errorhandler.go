package errorhandler

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ErrorLevel represents the severity level of an error
type ErrorLevel int

const (
	Info ErrorLevel = iota
	Warning
	Error
	Critical
)

// Define common error types
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal error")
)

// CustomError represents a custom error type
type CustomError struct {
	Code     int         // Error code
	Message  string      // Error message
	Level    ErrorLevel  // Error level
	Context  interface{} // Context information
	Original error       // Original error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Level: %d, Message: %s, Context: %v, Original: %v",
		e.Code, e.Level, e.Message, e.Context, e.Original)
}

// NewError creates a new custom error
func NewError(code int, message string, level ErrorLevel, context interface{}, original error) *CustomError {
	return &CustomError{
		Code:     code,
		Message:  message,
		Level:    level,
		Context:  context,
		Original: original,
	}
}

// Wrap wraps an error with an additional message
func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

// Unwrap unwraps an error to get the original error
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// LogError logs the error to both file and console
func LogError(err error) {
	log.Printf("ERROR: %v", err)
}

// InitLogFile initializes the log file
func InitLogFile(logFilePath string) error {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	return nil
}

// NotifyError sends an error notification (example: send email)
func NotifyError(err error, recipientEmail string) {
	// Add code here to send an email using SMTP protocol for example
	fmt.Printf("Sending error notification to %s: %v\n", recipientEmail, err)
}

// ErrorHandler is an interface for custom error handlers
type ErrorHandler interface {
	HandleError(err error)
}

// Global registry of custom error handlers
var customErrorHandlers []ErrorHandler

// RegisterErrorHandler registers a custom error handler
func RegisterErrorHandler(handler ErrorHandler) {
	customErrorHandlers = append(customErrorHandlers, handler)
}

// TriggerCustomErrorHandlers triggers all registered custom error handlers
func TriggerCustomErrorHandlers(err error) {
	for _, handler := range customErrorHandlers {
		handler.HandleError(err)
	}
}

// AggregateError aggregates multiple errors into one
type AggregateError struct {
	Errors []error
}

func (e *AggregateError) Error() string {
	var result string
	for _, err := range e.Errors {
		result += err.Error() + "\n"
	}
	return result
}

func NewAggregateError(errors []error) *AggregateError {
	return &AggregateError{Errors: errors}
}

// Multi-language support (example)
var languageMessages = map[string]map[int]string{
	"en": {
		1001: "Resource not found",
		1002: "Unauthorized access",
		1003: "Forbidden access",
		1004: "Internal server error",
	},
	"zh": {
		1001: "资源未找到",
		1002: "未经授权访问",
		1003: "禁止访问",
		1004: "内部服务器错误",
	},
}

func GetLocalizedMessage(code int, lang string) string {
	if messages, exists := languageMessages[lang]; exists {
		if message, exists := messages[code]; exists {
			return message
		}
	}
	return fmt.Sprintf("Unknown error code: %d", code)
}
