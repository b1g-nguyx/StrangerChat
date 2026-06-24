package apperror

import "net/http"

// AppError is a custom error type for managing API errors centrally
type AppError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	HTTPCode  int    `json:"-"` // Not exposed in JSON, used by fiber
}

// Error implements the standard error interface
func (e *AppError) Error() string {
	return e.Message
}

// ==========================================
// CENTRALIZED ERROR DICTIONARY
// ==========================================

var (
	// System Errors (1000 - 1999)
	ErrInvalidJSON = &AppError{ErrorCode: 1001, Message: "Invalid JSON format", HTTPCode: http.StatusBadRequest}
	ErrValidation  = &AppError{ErrorCode: 1002, Message: "Validation failed", HTTPCode: http.StatusBadRequest}
	ErrInternal    = &AppError{ErrorCode: 1003, Message: "Internal server error", HTTPCode: http.StatusInternalServerError}

	// Auth & User Errors (2000 - 2999)
	ErrUnauthorized    = &AppError{ErrorCode: 2001, Message: "Unauthorized access", HTTPCode: http.StatusUnauthorized}
	ErrInvalidEmailPwd = &AppError{ErrorCode: 2002, Message: "Invalid email or password", HTTPCode: http.StatusUnauthorized}
	ErrEmailExists     = &AppError{ErrorCode: 2003, Message: "Email already exists", HTTPCode: http.StatusConflict}
	ErrUserNotFound    = &AppError{ErrorCode: 2004, Message: "User not found", HTTPCode: http.StatusNotFound}
)
