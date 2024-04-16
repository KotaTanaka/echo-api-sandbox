package clientusecase

import (
	"github.com/jinzhu/gorm"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
)

type ReviewUsecase interface {
	GetReviewList(query *clientdto.ReviewListingQuery) (*clientdto.ReviewListingResponse, *dto.ErrorResponse)
	CreateReview(body *clientdto.CreateReviewRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse)
}

type reviewUsecase struct {
	db *gorm.DB
}

func NewReviewUsecase(db *gorm.DB) ReviewUsecase {
	return &reviewUsecase{db: db}
}

func (u *reviewUsecase) GetReviewList(query *clientdto.ReviewListingQuery) (*clientdto.ReviewListingResponse, *dto.ErrorResponse) {
	var shop model.Shop
	var service model.Service
	var reviews []model.Review

	if u.db.Find(&shop, query.ShopID).Related(&reviews).RecordNotFound() {
		return nil, dto.NotFoundError("Shop")
	}

	u.db.First(&service, shop.ID)

	res := &clientdto.ReviewListingResponse{}
	res.ShopID = shop.ID
	res.ShopName = shop.ShopName
	res.ServiceID = service.ID
	res.WifiName = service.WifiName
	res.Total = len(reviews)
	res.ReviewList = []clientdto.ReviewListingResponseElement{}

	var evaluationSum int
	for _, review := range reviews {
		evaluationSum += review.Evaluation
		res.ReviewList = append(
			res.ReviewList,
			clientdto.ReviewListingResponseElement{
				ReviewID:   review.ID,
				Comment:    review.Comment,
				Evaluation: review.Evaluation,
				Status:     review.PublishStatus,
				CreatedAt:  review.CreatedAt,
			},
		)
	}

	if res.Total > 0 {
		res.Average = float32(evaluationSum) / float32(res.Total)
	}

	return res, nil
}

func (u *reviewUsecase) CreateReview(body *clientdto.CreateReviewRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	shop := model.Shop{}
	u.db.First(&shop, body.ShopID)

	review := new(model.Review)
	review.ShopID = shop.ID
	review.Comment = body.Comment
	review.Evaluation = body.Evaluation

	u.db.Create(&review)

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}
