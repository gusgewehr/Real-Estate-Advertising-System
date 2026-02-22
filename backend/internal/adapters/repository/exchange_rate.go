package repository

import (
	"real-estate-api/internal/domain"
	"real-estate-api/internal/infrastructure"

	"go.uber.org/zap"
)

type ExchangeRateRepository struct {
	db     *infrastructure.Database
	logger *zap.SugaredLogger
}

func NewExchangeRateRepository(db *infrastructure.Database, logger *zap.SugaredLogger) *ExchangeRateRepository {
	return &ExchangeRateRepository{
		db:     db,
		logger: logger,
	}
}

func (r *ExchangeRateRepository) Create(rate *domain.ExchangeRate) error {

	result := r.db.Orm.Create(&rate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ExchangeRateRepository) GetLatest() (*domain.ExchangeRate, error) {
	var rate domain.ExchangeRate

	result := r.db.Orm.Last(&rate).Order("date(created_at)")
	if result.Error != nil {
		return nil, result.Error
	}

	return &rate, nil
}
