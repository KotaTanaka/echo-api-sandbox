package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	ListReviews() ([]*model.Review, error)
	ListReviewsByShopID(shopID int) ([]*model.Review, error)
	FindReviewByID(reviewID int) (*model.Review, error)
	CreateReview(review *model.Review) (*model.Review, error)
	UpdateReview(review *model.Review) (*model.Review, error)
	DeleteReview(review *model.Review) error
	SelectReviewsCountAndAverageByShopID(shopID int) (*model.Aggregation, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) ListReviews() ([]*model.Review, error) {
	reviews := []*model.Review{}
	if err := r.db.
		Preload("Shop.Service").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) ListReviewsByShopID(shopID int) ([]*model.Review, error) {
	reviews := []*model.Review{}
	if err := r.db.
		Preload("Shop.Service").
		Where("shop_id = ?", shopID).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) FindReviewByID(reviewID int) (*model.Review, error) {
	var review model.Review
	if err := r.db.First(&review, reviewID).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *reviewRepository) CreateReview(review *model.Review) (*model.Review, error) {
	if err := r.db.Create(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *reviewRepository) UpdateReview(review *model.Review) (*model.Review, error) {
	if err := r.db.Save(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *reviewRepository) DeleteReview(review *model.Review) error {
	return r.db.Delete(&review).Error
}

func (r *reviewRepository) SelectReviewsCountAndAverageByShopID(shopID int) (*model.Aggregation, error) {
	reviews := r.db.Model(&model.Review{}).Where("shop_id = ?", shopID)

	var reviewCount int64
	if err := reviews.Count(&reviewCount).Error; err != nil {
		return nil, err
	}

	if reviewCount == 0 {
		return &model.Aggregation{}, nil
	}

	var average float32
	if err := reviews.Select("AVG(evaluation)").Row().Scan(&average); err != nil {
		return nil, err
	}

	return &model.Aggregation{
		Count:   reviewCount,
		Average: average,
	}, nil
}
