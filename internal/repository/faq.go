package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type IFAQRepository interface {
}

type FAQRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) IFAQRepository {
	return &FAQRepository{db}
}

func (fr *FAQRepository) CreateFAQ(faq *entity.FAQ) (*entity.FAQ, error) {

	return faq, nil
}
