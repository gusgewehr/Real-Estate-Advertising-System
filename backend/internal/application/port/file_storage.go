package port

import (
	"io"
)

type FileStorageInputPort interface {
	Upload(file io.Reader, fileName string) (string, error)
}

type FileStorageOutputPort interface {
	Upload(fullPath string, reader io.Reader) (string, error)
}
