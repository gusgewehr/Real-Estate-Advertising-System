package test

import (
	"real-estate-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeRateInput_IsValid_Positive(t *testing.T) {
	input := &domain.ExchangeRateInput{Value: 0.25}
	err := input.IsValid()
	assert.NoError(t, err)
}

func TestExchangeRateInput_IsValid_Zero(t *testing.T) {
	input := &domain.ExchangeRateInput{Value: 0}
	err := input.IsValid()
	assert.Error(t, err)
	assert.EqualError(t, err, "value must be greater than zero")
}

func TestExchangeRateInput_IsValid_Negative(t *testing.T) {
	input := &domain.ExchangeRateInput{Value: -1}
	err := input.IsValid()
	assert.Error(t, err)
}

func TestExchangeRateInput_ToDb(t *testing.T) {
	input := &domain.ExchangeRateInput{Value: 0.25}
	rate := input.ToDb()

	assert.NotNil(t, rate)
	assert.Equal(t, 0.25, rate.Value)
	assert.Equal(t, uint(0), rate.Id)
}

func TestExchangeRate_FromDb(t *testing.T) {
	rate := &domain.ExchangeRate{Id: 1, Value: 0.25}
	input := rate.FromDb()

	assert.NotNil(t, input)
	assert.Equal(t, 0.25, input.Value)
}
