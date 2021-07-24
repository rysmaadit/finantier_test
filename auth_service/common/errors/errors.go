package errors

import (
	"errors"
	"fmt"
	"strings"
)

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{message: err.Error()}
}

type DBError struct {
	message string
}

func (e *DBError) Error() string {
	return e.message
}

func NewDBError(err error, message string) *DBError {
	return &DBError{message: fmt.Sprintf("%s: %s", message, err.Error())}
}

func NewDBErrorf(err error, format string, args ...interface{}) *DBError {
	message := fmt.Sprintf(format, args...)
	return &DBError{message: fmt.Sprintf("%s: %s", message, err.Error())}
}

type ConflictError struct {
	message string
}

func (e *ConflictError) Error() string {
	return e.message
}

func NewConflictErrorf(format string, args ...interface{}) *ConflictError {
	return &ConflictError{message: fmt.Sprintf(format, args...)}
}

type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{message: err.Error()}
}

type InternalError struct {
	message string
}

func (e *InternalError) Error() string {
	return e.message
}

func NewInternalError(err error, message string) *InternalError {
	return &InternalError{message: fmt.Sprintf("%s: %s", message, err.Error())}
}

type ExternalError struct {
	message string
}

func (e *ExternalError) Error() string {
	return e.message
}

func NewExternalError(format string, args ...interface{}) *ExternalError {
	return &ExternalError{message: fmt.Sprintf(format, args...)}
}

type UnauthorizedError struct {
	message string
}

func (e *UnauthorizedError) Error() string {
	return e.message
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{message: message}
}

type ValidationError struct {
	ValidationErrors map[string]string `json:"validation_errors"`
}

func NewValidationError(validationErrors map[string]string) *ValidationError {
	return &ValidationError{ValidationErrors: validationErrors}
}

func (e *ValidationError) Error() string {
	message := make([]string, 0)
	for _, msg := range e.ValidationErrors {
		message = append(message, fmt.Sprintf("%s", msg))
	}
	return strings.Join(message, ", ")
}

func New(message string) error {
	return errors.New(message)
}
