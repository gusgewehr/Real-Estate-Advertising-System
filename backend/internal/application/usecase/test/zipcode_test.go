package test

import (
	"real-state-api/internal/application/usecase"
	"real-state-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockZipCodeOutput struct {
	mock.Mock
}

func (m *MockZipCodeOutput) GetZipCode(zipCodeStr string) (*domain.ZipCode, error) {
	args := m.Called(zipCodeStr)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ZipCode), args.Error(1)
}

func TestZipCodeUseCase_GetZipCode_Success(t *testing.T) {
	mockPort := new(MockZipCodeOutput)
	uc := usecase.NewZipCodeUseCase(mockPort, newTestLogger())

	mockPort.On("GetZipCode", "98200000").Return(&domain.ZipCode{
		Cep:        "98200000",
		Logradouro: "flores da cunha",
		Bairro:     "Centro",
		Localidade: "ibirubá",
		Uf:         "RS",
	}, nil)

	addr, err := uc.GetZipCode("98200000")

	assert.NoError(t, err)
	assert.NotNil(t, addr)
	assert.Equal(t, "98200000", addr.ZipCode)
	assert.Equal(t, "ibirubá", addr.City)
	assert.Equal(t, "RS", addr.StateAbbr)
	mockPort.AssertExpectations(t)
}

func TestZipCodeUseCase_GetZipCode_NotFound(t *testing.T) {
	mockPort := new(MockZipCodeOutput)
	uc := usecase.NewZipCodeUseCase(mockPort, newTestLogger())

	mockPort.On("GetZipCode", "99999999").Return(nil, domain.ErrNotFound)

	addr, err := uc.GetZipCode("99999999")

	assert.Error(t, err)
	assert.Nil(t, addr)
	assert.ErrorIs(t, err, domain.ErrNotFound)
	mockPort.AssertExpectations(t)
}

func TestZipCodeUseCase_GetZipCode_ConnectionError(t *testing.T) {
	mockPort := new(MockZipCodeOutput)
	uc := usecase.NewZipCodeUseCase(mockPort, newTestLogger())

	mockPort.On("GetZipCode", "98200000").Return(nil, domain.ErrConnection)

	addr, err := uc.GetZipCode("98200000")

	assert.Error(t, err)
	assert.Nil(t, addr)
	assert.ErrorIs(t, err, domain.ErrConnection)
	mockPort.AssertExpectations(t)
}
