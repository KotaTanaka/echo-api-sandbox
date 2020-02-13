/*
Package data | 汎用的な構造体
*/
package data

/*
MessageResponse | メッセージを返却するレスポンス
*/
type MessageResponse struct {
	Message string `json:"message"`
}

/*
ErrorResponse | エラーレスポンス
*/
type ErrorResponse struct {
	Code          int      `json:"code"`
	Message       string   `json:"message"`
	DetailMessage []string `json:"detailMessage"`
}
