package clientusecase

import (
	"fmt"
	"strconv"

	"github.com/KotaTanaka/echo-api-sandbox/application/dto"
	clientdto "github.com/KotaTanaka/echo-api-sandbox/application/dto/client"
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/KotaTanaka/echo-api-sandbox/domain/repository"
)

type ReviewUsecase interface {
	GetReviewList(query *clientdto.ReviewListingQuery) (*clientdto.ReviewListingResponse, *dto.ErrorResponse)
	CreateReview(body *clientdto.CreateReviewRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse)
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

func (u *reviewUsecase) GetReviewList(query *clientdto.ReviewListingQuery) (*clientdto.ReviewListingResponse, *dto.ErrorResponse) {
	shopID, err := strconv.Atoi(query.ShopID)
	if err != nil {
		return nil, dto.InvalidParameterError([]string{"ShopID must be number."})
	}

	reviews, err := u.reviewRepository.ListReviewsByShopID(shopID)
	if err != nil {
		return nil, dto.HandleDBError(err, fmt.Sprintf("Reviews(Shop ID:%d)", shopID))
	}

	if len(reviews) == 0 {
		return &clientdto.ReviewListingResponse{}, nil
	}

	res := &clientdto.ReviewListingResponse{
		ShopID:     reviews[0].Shop.ID,
		ShopName:   reviews[0].Shop.ShopName,
		ServiceID:  reviews[0].Shop.Service.ID,
		WifiName:   reviews[0].Shop.Service.WifiName,
		Total:      len(reviews),
		ReviewList: make([]clientdto.ReviewListingResponseElement, len(reviews)),
	}

	var evaluationSum int
	for i, review := range reviews {
		evaluationSum += review.Evaluation
		res.ReviewList[i] = clientdto.ReviewListingResponseElement{
			ReviewID:   review.ID,
			Comment:    review.Comment,
			Evaluation: review.Evaluation,
			Status:     review.PublishStatus,
			CreatedAt:  review.CreatedAt,
		}
	}

	if res.Total > 0 {
		res.Average = float32(evaluationSum) / float32(res.Total)
	}

	return res, nil
}

func (u *reviewUsecase) CreateReview(body *clientdto.CreateReviewRequest) (*dto.ReviewIDResponse, *dto.ErrorResponse) {
	shop, err := u.shopRepository.FindShopByID(int(body.ShopID))
	if err != nil {
		return nil, dto.HandleDBError(err, fmt.Sprintf("Shop(ID:%d)", body.ShopID))
	}

	review := &model.Review{
		ShopID:     shop.ID,
		Comment:    body.Comment,
		Evaluation: body.Evaluation,
	}

	review, err = u.reviewRepository.CreateReview(review)
	if err != nil {
		return nil, dto.HandleDBError(err, "Review")
	}

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}
