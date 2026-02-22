package test

import (
	"real-estate-api/internal/application/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaginationOutput struct {
	mock.Mock
}

func (m *MockPaginationOutput) GetTotalItems(table string) int64 {
	args := m.Called(table)
	return args.Get(0).(int64)
}

func (m *MockPaginationOutput) GetTableName(structTable interface{}) string {
	args := m.Called(structTable)
	return args.String(0)
}

func TestPaginationUseCase_GetTotalItems_Success(t *testing.T) {
	mockPort := new(MockPaginationOutput)
	uc := usecase.NewPaginationUseCase(mockPort, newTestLogger())

	type FakeModel struct{}

	mockPort.On("GetTableName", mock.Anything).Return("real_estate_properties")
	mockPort.On("GetTotalItems", "real_estate_properties").Return(int64(42))

	total := uc.GetTotalItems(FakeModel{})

	assert.Equal(t, int64(42), total)
	mockPort.AssertExpectations(t)
}

func TestPaginationUseCase_GetTotalItems_EmptyTableName(t *testing.T) {
	mockPort := new(MockPaginationOutput)
	uc := usecase.NewPaginationUseCase(mockPort, newTestLogger())

	mockPort.On("GetTableName", mock.Anything).Return("")

	total := uc.GetTotalItems(struct{}{})

	assert.Equal(t, int64(0), total)
	mockPort.AssertNotCalled(t, "GetTotalItems")
}
