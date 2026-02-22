package test

import (
	"errors"
	"io"
	"real-estate-api/internal/application/usecase"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock do FileStorageOutputPort ---

type MockFileStorageOutput struct {
	mock.Mock
}

func (m *MockFileStorageOutput) Upload(fullPath string, reader io.Reader) (string, error) {
	args := m.Called(fullPath, reader)
	return args.String(0), args.Error(1)
}

// --- Testes ---

func TestFileStorageUseCase_Upload_Success(t *testing.T) {
	mockStorage := new(MockFileStorageOutput)
	uc := usecase.NewFileStorageUseCase(mockStorage, "/tmp/uploads", "http://localhost:8080", "/files", newTestLogger())

	mockStorage.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("/tmp/uploads/somefile.jpg", nil)

	reader := strings.NewReader("file content")
	url, err := uc.Upload(reader, "photo.jpg")

	assert.NoError(t, err)
	assert.Contains(t, url, "http://localhost:8080")
	assert.Contains(t, url, ".jpg")
	mockStorage.AssertExpectations(t)
}

func TestFileStorageUseCase_Upload_PreservesExtension(t *testing.T) {
	mockStorage := new(MockFileStorageOutput)
	uc := usecase.NewFileStorageUseCase(mockStorage, "/tmp/uploads", "http://localhost:8080", "/files", newTestLogger())

	mockStorage.On("Upload", mock.MatchedBy(func(path string) bool {
		return strings.HasSuffix(path, ".png")
	}), mock.Anything).Return("/tmp/uploads/uuid.png", nil)

	reader := strings.NewReader("file content")
	url, err := uc.Upload(reader, "image.png")

	assert.NoError(t, err)
	assert.NotEmpty(t, url)
	mockStorage.AssertExpectations(t)
}

func TestFileStorageUseCase_Upload_StorageError(t *testing.T) {
	mockStorage := new(MockFileStorageOutput)
	uc := usecase.NewFileStorageUseCase(mockStorage, "/tmp/uploads", "http://localhost:8080", "/files", newTestLogger())

	mockStorage.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("", errors.New("disk full"))

	reader := strings.NewReader("file content")
	url, err := uc.Upload(reader, "photo.jpg")

	assert.Error(t, err)
	assert.EqualError(t, err, "disk full")
	assert.Empty(t, url)
	mockStorage.AssertExpectations(t)
}
