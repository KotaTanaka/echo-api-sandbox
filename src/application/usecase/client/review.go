package clientusecase

import (
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

	shop, err := u.shopRepository.FindShopByID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	reviews, err := u.reviewRepository.ListReviewsByShopID(shopID)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	service, err := u.serviceRepository.FindServiceByID(int(shop.ServiceID))
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

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
	shop, err := u.shopRepository.FindShopByID(int(body.ShopID))
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	review := new(model.Review)
	review.ShopID = shop.ID
	review.Comment = body.Comment
	review.Evaluation = body.Evaluation

	review, err = u.reviewRepository.CreateReview(review)
	if err != nil {
		return nil, dto.InternalServerError(err)
	}

	return &dto.ReviewIDResponse{
		ReviewID: review.ID,
	}, nil
}
