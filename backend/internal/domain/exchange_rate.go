package domain

import (
	"errors"
	"time"
)

type ExchangeRate struct {
	Id        uint
	Value     float64
	CreatedAt time.Time
}

type ExchangeRateInput struct {
	Value float64 `json:"value"`
}

func (e *ExchangeRateInput) IsValid() error {
	if e.Value <= 0 {
		return errors.New("value must be greater than zero")
	}
	return nil
}

func (e *ExchangeRateInput) ToDb() *ExchangeRate {
	return &ExchangeRate{
		Value: e.Value,
	}

}

func (e *ExchangeRate) FromDb() *ExchangeRateInput {
	return &ExchangeRateInput{
		Value: e.Value,
	}
}
