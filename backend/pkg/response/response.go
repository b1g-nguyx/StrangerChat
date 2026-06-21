package response

import (
	"errors"
	"time"

	"github.com/b1g-nguyx/strangerchat-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
)

// APIResponse represents the standard structure for all REST API responses
type APIResponse struct {
	Message   string    `json:"message"`
	ErrorCode int       `json:"error_code,omitempty"` // omitempty: hidden if 0 (success)
	Data      any       `json:"data,omitempty"`       // omitempty: hidden if nil
	Time      time.Time `json:"time"`
}

// Success returns a standardized success response.
// Note: We removed the statusCode parameter because success is usually 200 or 201.
// You can pass it if you want, but sticking to fiber.StatusOK is standard for most cases.
func Success(c *fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(APIResponse{
		Message: message,
		Data:    data,
		Time:    time.Now().UTC(),
	})
}

// Error takes a standard Go error. If it is a custom AppError, it automatically
// extracts the HTTP Code, Error Code, and standardized Message.
// Example: return response.Error(c, apperror.ErrInvalidJSON)
func Error(c *fiber.Ctx, err error) error {

	// Check if the error matches our centralized AppError dictionary
	if appErr, ok := errors.AsType[*apperror.AppError](err); ok {
		return c.Status(appErr.HTTPCode).JSON(APIResponse{
			Message:   appErr.Message,
			ErrorCode: appErr.ErrorCode,
			Time:      time.Now().UTC(),
		})
	}

	// Fallback for unknown errors (e.g., standard database crash, nil pointer, etc.)
	return c.Status(fiber.StatusInternalServerError).JSON(APIResponse{
		Message:   "Internal Server Error: " + err.Error(),
		ErrorCode: 9999, // Generic unknown error code
		Time:      time.Now().UTC(),
	})
}
