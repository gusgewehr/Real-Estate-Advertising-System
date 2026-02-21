package port

import "real-state-api/internal/domain"

type RealEstateInputPort interface {
	Create(property *domain.RealEstatePropertyInput) error
	List(page int, pageSize int64) (*domain.PaginatedResponse[*domain.RealEstatePropertyInput], error)
}

type RealEstateOutputPort interface {
	Create(property *domain.RealEstateProperty) error
	List(offset, limit int) ([]*domain.RealEstateProperty, error)
}
