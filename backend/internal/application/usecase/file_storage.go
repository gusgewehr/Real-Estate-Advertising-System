package usecase

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"real-estate-api/internal/application/port"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type FileStorageUseCase struct {
	fileStorage port.FileStorageOutputPort
	logger      *zap.SugaredLogger
	filePath    string
	hostURL     string
	fileUrl     string
}

func NewFileStorageUseCase(fileStorage port.FileStorageOutputPort, filePath, hostUrl, fileUrl string, logger *zap.SugaredLogger) *FileStorageUseCase {
	return &FileStorageUseCase{
		fileStorage: fileStorage,
		logger:      logger,
		filePath:    filePath,
		hostURL:     hostUrl,
		fileUrl:     fileUrl,
	}
}

func (uc *FileStorageUseCase) Upload(file io.Reader, fileName string) (string, error) {

	os.MkdirAll(uc.filePath, os.ModePerm)

	fileExt := filepath.Ext(fileName)
	uuidName, _ := uuid.NewUUID()
	fileName = fmt.Sprint(uuidName.String(), fileExt)
	fullPath := filepath.Join(uc.filePath, fileName)

	res, err := uc.fileStorage.Upload(fullPath, file)
	if err != nil {
		return "", err
	}

	imageUrl := uc.hostURL + res

	return imageUrl, nil
}
