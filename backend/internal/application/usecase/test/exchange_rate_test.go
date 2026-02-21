package test

import (
	"errors"
	"real-state-api/internal/application/usecase"
	"real-state-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockExchangeRateOutput struct {
	mock.Mock
}

func (m *MockExchangeRateOutput) Create(rate *domain.ExchangeRate) error {
	args := m.Called(rate)
	return args.Error(0)
}

func (m *MockExchangeRateOutput) GetLatest() (*domain.ExchangeRate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ExchangeRate), args.Error(1)
}

func newTestLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	return logger.Sugar()
}

func TestExchangeRateUseCase_Create_Success(t *testing.T) {
	mockPort := new(MockExchangeRateOutput)
	uc := usecase.NewExchangeRateUseCase(mockPort, newTestLogger())

	mockPort.On("Create", mock.AnythingOfType("*domain.ExchangeRate")).Return(nil)

	input := &domain.ExchangeRateInput{Value: 0.25}
	err := uc.Create(input)

	assert.NoError(t, err)
	mockPort.AssertExpectations(t)
}

func TestExchangeRateUseCase_Create_InvalidValue(t *testing.T) {
	mockPort := new(MockExchangeRateOutput)
	uc := usecase.NewExchangeRateUseCase(mockPort, newTestLogger())

	input := &domain.ExchangeRateInput{Value: 0}
	err := uc.Create(input)

	assert.Error(t, err)
	assert.EqualError(t, err, "value must be greater than zero")
	mockPort.AssertNotCalled(t, "Create")
}

func TestExchangeRateUseCase_Create_PortError(t *testing.T) {
	mockPort := new(MockExchangeRateOutput)
	uc := usecase.NewExchangeRateUseCase(mockPort, newTestLogger())

	mockPort.On("Create", mock.AnythingOfType("*domain.ExchangeRate")).Return(errors.New("db error"))

	input := &domain.ExchangeRateInput{Value: 0.25}
	err := uc.Create(input)

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	mockPort.AssertExpectations(t)
}

func TestExchangeRateUseCase_GetLatest_Success(t *testing.T) {
	mockPort := new(MockExchangeRateOutput)
	uc := usecase.NewExchangeRateUseCase(mockPort, newTestLogger())

	mockPort.On("GetLatest").Return(&domain.ExchangeRate{Id: 1, Value: 0.25}, nil)

	result, err := uc.GetLatest()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0.25, result.Value)
	mockPort.AssertExpectations(t)
}

func TestExchangeRateUseCase_GetLatest_Error(t *testing.T) {
	mockPort := new(MockExchangeRateOutput)
	uc := usecase.NewExchangeRateUseCase(mockPort, newTestLogger())

	mockPort.On("GetLatest").Return(nil, errors.New("not found"))

	result, err := uc.GetLatest()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockPort.AssertExpectations(t)
}
