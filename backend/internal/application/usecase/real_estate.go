package usecase

import (
	"errors"
	"real-state-api/internal/application/port"
	"real-state-api/internal/domain"
	"time"

	"go.uber.org/zap"
)

type RealEstateUseCase struct {
	realEstatePort port.RealEstateOutputPort
	paginationPort port.PaginationInput
	logger         *zap.SugaredLogger
}

func NewRealEstateUseCase(realEstatePort port.RealEstateOutputPort, paginationPort port.PaginationInput, logger *zap.SugaredLogger) *RealEstateUseCase {
	return &RealEstateUseCase{realEstatePort: realEstatePort, paginationPort: paginationPort, logger: logger}
}

func (uc *RealEstateUseCase) Create(property *domain.RealEstatePropertyInput) error {

	errs := property.IsValid()
	if len(errs) > 0 {
		message := ""
		for _, err := range errs {
			message += err.Error() + "\n"
		}
		return errors.New(message)
	}

	propertyOutput := property.ToDb()
	propertyOutput.CreatedAt = time.Now()
	propertyOutput.UpdatedAt = time.Now()

	err := uc.realEstatePort.Create(propertyOutput)
	if err != nil {
		return err
	}

	return nil
}

func (uc *RealEstateUseCase) List(page int, pageSize int64) (*domain.PaginatedResponse[*domain.RealEstatePropertyInput], error) {
	var propertiesInput []*domain.RealEstatePropertyInput

	totalItems := uc.paginationPort.GetTotalItems(domain.RealEstateProperty{})

	paginationParams := domain.BuildPaginationParams(page, pageSize, totalItems)

	offset := paginationParams.Offset()

	properties, err := uc.realEstatePort.List(int(offset), int(paginationParams.PageSize))
	if err != nil {
		return nil, err
	}

	for _, property := range properties {
		propertyInput := property.FromDb()
		propertiesInput = append(propertiesInput, &propertyInput)
	}

	response := domain.CreatePaginatedResponse[*domain.RealEstatePropertyInput](propertiesInput, totalItems, paginationParams)

	return response, nil
}
