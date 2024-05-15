package adminusecase

import (
	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	admindto "github.com/KotaTanaka/echo-api-sandbox/application/dto/admin"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type ReviewUsecase interface {
	GetReviewList() (*admindto.ReviewListingResponse, *dto.ErrorResponse)
	UpdateReviewStatus(reviewID int, body *admindto.UpdateReviewStatusRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse)
	DeleteReview(reviewID int) (*dto.ReviewIDResponse, *dto.ErrorResponse)
}

type reviewUsecase struct {
	serviceRepository repository.ServiceRepository
	shopRepository    repository.ShopRepository
	reviewRepository  repository.ReviewRepository
}

func NewReviewUsecase(
	serviceRepository repository.ServiceRepository,
	shopRepository repository.ShopRepository,
	reviewRepository repository.ReviewRepository,
) ReviewUsecase {
	return &reviewUsecase{
		serviceRepository: serviceRepository,
		shopRepository:    shopRepository,
		reviewRepository:  reviewRepository,
	}
}

func (u *reviewUsecase) GetReviewList() (*admindto.ReviewListingResponse, *dto.ErrorResponse) {
	reviews, err := u.reviewRepository.ListReviews()
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	res := &admindto.ReviewListingResponse{
		Total:      len(reviews),
		ReviewList: make([]admindto.ReviewListingResponseElement, len(reviews)),
	}

	for i, review := range reviews {
		shop, err := u.shopRepository.FindShopByID(int(review.ShopID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		service, err := u.serviceRepository.FindServiceByID(int(shop.ServiceID))
		if err != nil {
			return nil, dto.InternalServerError(err)
		}

		res.ReviewList[i] = admindto.ReviewListingResponseElement{
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
		}
	}

	return res, nil
}

func (u *reviewUsecase) UpdateReviewStatus(reviewID int, body *admindto.UpdateReviewStatusRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	review, err := u.reviewRepository.FindReviewByID(reviewID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	if body.Status == "public" {
		review.PublishStatus = true
	} else if body.Status == "hidden" {
		review.PublishStatus = false
	} else {
		return nil, dto.InvalidParameterError([]string{"Status is 'public' or 'hidden'"})
	}

	review, err = u.reviewRepository.UpdateReview(review)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}

func (u *reviewUsecase) DeleteReview(reviewID int) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	review, err := u.reviewRepository.FindReviewByID(reviewID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	err = u.reviewRepository.DeleteReview(review)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}
