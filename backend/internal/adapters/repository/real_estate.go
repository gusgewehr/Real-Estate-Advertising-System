package repository

import (
	"real-estate-api/internal/domain"
	"real-estate-api/internal/infrastructure"

	"go.uber.org/zap"
)

type RealEstateRepository struct {
	db     *infrastructure.Database
	logger *zap.SugaredLogger
}

func NewRealEstateRepository(db *infrastructure.Database, logger *zap.SugaredLogger) *RealEstateRepository {
	return &RealEstateRepository{db, logger}
}

func (r *RealEstateRepository) Create(property *domain.RealEstateProperty) error {

	result := r.db.Orm.Create(&property)
	if result.Error != nil {
		r.logger.Errorw("Error creating property", "error", result.Error)
		return result.Error
	}

	return nil
}

func (r *RealEstateRepository) List(offset, limit int) ([]*domain.RealEstateProperty, error) {
	var properties []*domain.RealEstateProperty

	result := r.db.Orm.Offset(offset).Limit(limit).Find(&properties)
	if result.Error != nil {
		r.logger.Errorw("Error listing properties", "error", result.Error)
		return nil, result.Error
	}

	return properties, nil
}
