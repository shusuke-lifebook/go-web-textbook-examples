package apperror

import "net/http"

func NewValidation(msg string, details []FieldIssue, cause error) *AppError {
	return &AppError{
		Code:     "VALIDATION_FAILED",
		Message:  msg,
		HTTPCode: http.StatusUnprocessableEntity, // 422
		Cause:    cause,
		Details:  details,
	}
}

func NewBadRequest(msg string, cause error) *AppError {
	return &AppError{
		Code:     "BAD_REQUEST",
		Message:  msg,
		HTTPCode: http.StatusBadRequest,
		Cause:    cause,
	}
}

func NewNotFound(msg string, cause error) *AppError {
	return &AppError{
		Code:     "NOT_FOUND",
		Message:  msg,
		HTTPCode: http.StatusNotFound,
		Cause:    cause,
	}
}

func NewConflict(msg string, cause error) *AppError {
	return &AppError{
		Code:     "CONFLICT",
		Message:  msg,
		HTTPCode: http.StatusConflict,
		Cause:    cause,
	}
}

func NewInternal(cause error) *AppError {
	return &AppError{
		Code:     "INTERNAL",
		Message:  "internal server error",
		HTTPCode: http.StatusInternalServerError,
		Cause:    cause,
	}
}
