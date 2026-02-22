package port

import "real-estate-api/internal/domain"

type ExchangeRateOutput interface {
	Create(rate *domain.ExchangeRate) error
	GetLatest() (*domain.ExchangeRate, error)
}

type ExchangeRateInput interface {
	Create(rate *domain.ExchangeRateInput) error
	GetLatest() (*domain.ExchangeRateInput, error)
}
