/*
Package data 汎用的な構造体
*/
package data

import (
	"net/http"
)

/*
ErrorResponse | エラーレスポンス
@type Response
*/
type ErrorResponse struct {
	Code          int      `json:"code"`
	Message       string   `json:"message"`
	DetailMessage []string `json:"detailMessage"`
}

/*
InvalidRequestError | リクエスト不正エラーのコンストラクタ
*/
func InvalidRequestError(errorMessages []string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid request."
	errorResponse.DetailMessage = errorMessages

	return errorResponse
}

/*
InvalidParameterError | パラメータ不正エラーのコンストラクタ
*/
func InvalidParameterError(errorMessages []string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid parameters."
	errorResponse.DetailMessage = errorMessages

	return errorResponse
}

/*
NotFoundError | 対象が存在しないエラーのコンストラクタ
*/
func NotFoundError(target string) *ErrorResponse {
	errorResponse := new(ErrorResponse)

	errorResponse.Code = http.StatusBadRequest
	errorResponse.Message = "Invalid request."
	errorResponse.DetailMessage = []string{"Not Found " + target + "."}

	return errorResponse
}
