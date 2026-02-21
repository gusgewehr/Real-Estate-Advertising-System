package repository

import (
	"fmt"
	"io"
	"os"
)

type LocalFileStorage struct {
}

func NewLocalFileStorage() *LocalFileStorage {
	return &LocalFileStorage{}
}

func (s *LocalFileStorage) Upload(fullPath string, reader io.Reader) (string, error) {

	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return "", fmt.Errorf("writing file: %w", err)
	}

	return fullPath, nil
}
