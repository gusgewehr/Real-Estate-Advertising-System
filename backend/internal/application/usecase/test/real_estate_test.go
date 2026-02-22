package test

import (
	"errors"
	"real-estate-api/internal/application/usecase"
	"real-estate-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRealEstateOutput struct {
	mock.Mock
}

func (m *MockRealEstateOutput) Create(property *domain.RealEstateProperty) error {
	args := m.Called(property)
	return args.Error(0)
}

func (m *MockRealEstateOutput) List(offset, limit int) ([]*domain.RealEstateProperty, error) {
	args := m.Called(offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.RealEstateProperty), args.Error(1)
}

type MockPaginationInput struct {
	mock.Mock
}

func (m *MockPaginationInput) GetTotalItems(object interface{}) int64 {
	args := m.Called(object)
	return args.Get(0).(int64)
}

func validTestAddress() domain.Address {
	return domain.Address{
		ZipCode:      "98200000",
		Street:       "flores da cunha",
		Neighborhood: "Centro",
		City:         "ibirubá",
		StateAbbr:    "RS",
	}
}

func TestRealEstateUseCase_Create_Success(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	mockRE.On("Create", mock.AnythingOfType("*domain.RealEstateProperty")).Return(nil)

	input := &domain.RealEstatePropertyInput{
		Type:    "SELL",
		Address: validTestAddress(),
		Value:   350000,
	}
	err := uc.Create(input)

	assert.NoError(t, err)
	mockRE.AssertExpectations(t)
}

func TestRealEstateUseCase_Create_InvalidInput(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	input := &domain.RealEstatePropertyInput{
		Type:    "",
		Address: domain.Address{},
		Value:   0,
	}
	err := uc.Create(input)

	assert.Error(t, err)
	mockRE.AssertNotCalled(t, "Create")
}

func TestRealEstateUseCase_Create_PortError(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	mockRE.On("Create", mock.AnythingOfType("*domain.RealEstateProperty")).Return(errors.New("db error"))

	input := &domain.RealEstatePropertyInput{
		Type:    "RENT",
		Address: validTestAddress(),
		Value:   2500,
	}
	err := uc.Create(input)

	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	mockRE.AssertExpectations(t)
}

// --- Testes List ---

func TestRealEstateUseCase_List_Success(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	addr := validTestAddress()
	properties := []*domain.RealEstateProperty{
		{Id: 1, Type: "SELL", Address: addr, Value: 350000},
		{Id: 2, Type: "RENT", Address: addr, Value: 2500},
	}

	mockPag.On("GetTotalItems", mock.Anything).Return(int64(2))
	mockRE.On("List", 0, 10).Return(properties, nil)

	result, err := uc.List(1, 10)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.Data))
	assert.Equal(t, int64(2), result.TotalItems)
	assert.Equal(t, 1, result.Page)
	mockRE.AssertExpectations(t)
	mockPag.AssertExpectations(t)
}

func TestRealEstateUseCase_List_Empty(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	mockPag.On("GetTotalItems", mock.Anything).Return(int64(0))
	mockRE.On("List", 0, 0).Return([]*domain.RealEstateProperty{}, nil)

	result, err := uc.List(1, -1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Data))
	mockRE.AssertExpectations(t)
}

func TestRealEstateUseCase_List_PortError(t *testing.T) {
	mockRE := new(MockRealEstateOutput)
	mockPag := new(MockPaginationInput)
	uc := usecase.NewRealEstateUseCase(mockRE, mockPag, newTestLogger())

	mockPag.On("GetTotalItems", mock.Anything).Return(int64(10))
	mockRE.On("List", 0, 10).Return(nil, errors.New("db error"))

	result, err := uc.List(1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRE.AssertExpectations(t)
}
