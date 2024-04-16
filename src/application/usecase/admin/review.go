package adminusecase

import (
	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
)

type ReviewUsecase interface {
	GetReviewList() (*admindto.ReviewListingResponse, *dto.ErrorResponse)
	UpdateReviewStatus(reviewID int, body *admindto.UpdateReviewStatusRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse)
	DeleteReview(reviewID int) (*dto.ReviewIDResponse, *dto.ErrorResponse)
}

type reviewUsecase struct {
	db *gorm.DB
}

func NewReviewUsecase(db *gorm.DB) ReviewUsecase {
	return &reviewUsecase{db: db}
}

func (u *reviewUsecase) GetReviewList() (*admindto.ReviewListingResponse, *dto.ErrorResponse) {
	reviews := []model.Review{}
	u.db.Find(&reviews)

	res := &admindto.ReviewListingResponse{}
	res.Total = len(reviews)
	res.ReviewList = []admindto.ReviewListingResponseElement{}

	for _, review := range reviews {
		shop := model.Shop{}
		u.db.First(&shop, review.ShopID)

		service := model.Service{}
		u.db.First(&service, shop.ID)

		res.ReviewList = append(
			res.ReviewList, admindto.ReviewListingResponseElement{
				ReviewID:   review.ID,
				ShopID:     shop.ID,
				ShopName:   shop.ShopName,
				ServiceID:  service.ID,
				WifiName:   service.WifiName,
				Comment:    review.Comment,
				Evaluation: review.Evaluation,
				Status:     review.PublishStatus,
				CreatedAt:  review.CreatedAt,
				UpdatedAt:  review.UpdatedAt,
				DeletedAt:  review.DeletedAt,
			},
		)
	}

	return res, nil
}

func (u *reviewUsecase) UpdateReviewStatus(reviewID int, body *admindto.UpdateReviewStatusRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	var review model.Review

	if u.db.Find(&review, reviewID).RecordNotFound() {
		return nil, dto.NotFoundError("Review")
	}

	if body.Status == "public" {
		review.PublishStatus = true
	} else if body.Status == "hidden" {
		review.PublishStatus = false
	} else {
		return nil, dto.InvalidParameterError([]string{"Status is 'public' or 'hidden'"})
	}

	u.db.Save(&review)

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}

func (u *reviewUsecase) DeleteReview(reviewID int) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	var review model.Review

	if u.db.Find(&review, reviewID).RecordNotFound() {
		return nil, dto.NotFoundError("Review")
	}

	u.db.Delete(&review, reviewID)

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}
