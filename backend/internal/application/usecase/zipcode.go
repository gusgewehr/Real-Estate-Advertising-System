package usecase

import (
	"real-estate-api/internal/application/port"
	"real-estate-api/internal/domain"

	"go.uber.org/zap"
)

type ZipCodeUseCase struct {
	port   port.ZipCodeOutputPort
	logger *zap.SugaredLogger
}

func NewZipCodeUseCase(port port.ZipCodeOutputPort, logger *zap.SugaredLogger) *ZipCodeUseCase {
	return &ZipCodeUseCase{port: port, logger: logger}
}

func (uc *ZipCodeUseCase) GetZipCode(zipcodeStr string) (*domain.Address, error) {

	zipcode, err := uc.port.GetZipCode(zipcodeStr)
	if err != nil {
		uc.logger.Errorw("failed to get zipcode", "error", err)
		return nil, err
	}

	address := zipcode.ToAddress()

	return address, nil

}
