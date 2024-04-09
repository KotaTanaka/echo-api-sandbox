package data

import (
	"net/http"
)

type ErrorResponse struct {
	Code          int      `json:"code"`
	Message       string   `json:"message"`
	DetailMessage []string `json:"detailMessage"`
}

func InvalidRequestError(errorMessages []string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid request."
	errorResponse.DetailMessage = errorMessages

	return errorResponse
}

func InvalidParameterError(errorMessages []string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid parameters."
	errorResponse.DetailMessage = errorMessages

	return errorResponse
}

func NotFoundError(target string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid request."
	errorResponse.DetailMessage = []string{"Not Found " + target + "."}

	return errorResponse
}