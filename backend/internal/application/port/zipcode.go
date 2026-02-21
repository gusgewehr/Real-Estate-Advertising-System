package port

import "real-state-api/internal/domain"

type ZipCodeOutputPort interface {
	GetZipCode(zipCodeStr string) (*domain.ZipCode, error)
}

type ZipCodeInputPort interface {
	GetZipCode(zipcodeStr string) (*domain.Address, error)
}
