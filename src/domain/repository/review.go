package repository

import (
	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"github.com/jinzhu/gorm"
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
	r.db.Find(&reviews)

	return reviews, nil
}

func (r *reviewRepository) ListReviewsByShopID(shopID int) ([]*model.Review, error) {
	reviews := []*model.Review{}
	r.db.Where("shop_id = ?", shopID).Find(&reviews)

	return reviews, nil
}

func (r *reviewRepository) FindReviewByID(reviewID int) (*model.Review, error) {
	var review *model.Review
	r.db.First(&review, reviewID)

	return review, nil
}

func (r *reviewRepository) CreateReview(review *model.Review) (*model.Review, error) {
	r.db.Create(&review)

	return review, nil
}

func (r *reviewRepository) UpdateReview(review *model.Review) (*model.Review, error) {
	r.db.Save(&review)

	return review, nil
}

func (r *reviewRepository) DeleteReview(review *model.Review) error {
	r.db.Delete(&review)

	return nil
}

func (r *reviewRepository) SelectReviewsCountAndAverageByShopID(shopID int) (*model.Aggregation, error) {
	reviews := r.db.Model(&model.Review{}).Where("shop_id = ?", shopID)
	var reviewCount int
	reviews.Count(&reviewCount)
	var average float32
	reviews.Select("avg(evaluation)").Row().Scan(&average)

	return &model.Aggregation{
		Count:   reviewCount,
		Average: average,
	}, nil
}
