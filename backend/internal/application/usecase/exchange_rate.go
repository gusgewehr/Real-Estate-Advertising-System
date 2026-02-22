package usecase

import (
	"real-estate-api/internal/application/port"
	"real-estate-api/internal/domain"
	"time"

	"go.uber.org/zap"
)

type ExchangeRateUseCase struct {
	port   port.ExchangeRateOutput
	logger *zap.SugaredLogger
}

func NewExchangeRateUseCase(port port.ExchangeRateOutput, logger *zap.SugaredLogger) *ExchangeRateUseCase {
	return &ExchangeRateUseCase{port, logger}
}

func (uc *ExchangeRateUseCase) Create(rate *domain.ExchangeRateInput) error {

	err := rate.IsValid()
	if err != nil {
		return err
	}

	rateOutput := rate.ToDb()
	rateOutput.CreatedAt = time.Now()

	return uc.port.Create(rateOutput)
}

func (uc *ExchangeRateUseCase) GetLatest() (*domain.ExchangeRateInput, error) {

	rate, err := uc.port.GetLatest()
	if err != nil {
		uc.logger.Errorw("failed to get latest exchange rate", "error", err)
		return nil, err
	}

	rateInput := rate.FromDb()

	return rateInput, nil
}
