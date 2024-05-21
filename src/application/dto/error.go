package dto

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ErrorResponse struct {
	Code          int      `json:"code"`
	Message       string   `json:"message"`
	DetailMessage []string `json:"detailMessage"`
}

func InvalidRequestError(errorMessages []string) *ErrorResponse {
	return &ErrorResponse{
		Code:          http.StatusBadRequest,
		Message:       "Invalid request.",
		DetailMessage: errorMessages,
	}
}

func InvalidParameterError(errorMessages []string) *ErrorResponse {
	return &ErrorResponse{
		Code:          http.StatusBadRequest,
		Message:       "Invalid parameters.",
		DetailMessage: errorMessages,
	}
}

func NotFoundError(target string) *ErrorResponse {
	return &ErrorResponse{
		Code:          http.StatusNotFound,
		Message:       "Record not found.",
		DetailMessage: []string{target + " is not found."},
	}
}

func InternalServerError(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:          http.StatusInternalServerError,
		Message:       "Internal server error occurred.",
		DetailMessage: []string{err.Error()},
	}
}

func HandleDBError(err error, target string) *ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFoundError(target)
	}

	return InternalServerError(err)
}
