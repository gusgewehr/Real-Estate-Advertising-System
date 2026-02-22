package usecase

import (
	"real-estate-api/internal/application/port"

	"go.uber.org/zap"
)

type PaginationUseCase struct {
	port   port.PaginationOutput
	logger *zap.SugaredLogger
}

func NewPaginationUseCase(port port.PaginationOutput, logger *zap.SugaredLogger) *PaginationUseCase {
	return &PaginationUseCase{port, logger}
}

func (uc *PaginationUseCase) GetTotalItems(object interface{}) int64 {

	table := uc.port.GetTableName(object)
	if table == "" {
		return 0
	}

	totalItems := uc.port.GetTotalItems(table)

	return totalItems
}
