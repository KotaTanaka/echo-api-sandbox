/*
Package clientdata クライアントAPI関連構造体
*/
package clientdata

/*
CreateReviewRequestBody | レビュー投稿リクエストボディ
@type Request
*/
type CreateReviewRequestBody struct {
	ShopID     uint   `json:"shopId" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
	Evaluation int    `json:"evaluation" validate:"required"`
}
