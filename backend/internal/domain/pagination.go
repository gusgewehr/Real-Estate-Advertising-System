package domain

import "math"

type PaginatedResponse[T any] struct {
	Data        []T   `json:"data"`
	Page        int   `json:"page"`
	PageSize    int64 `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	HasNextPage bool  `json:"has_next_page"`
	HasPrevPage bool  `json:"has_prev_page"`
}

type PaginationParams struct {
	Page     int
	PageSize int64
}

func BuildPaginationParams(page int, pageSize, totalItems int64) PaginationParams {
	if pageSize == -1 {
		pageSize = totalItems
	}

	return PaginationParams{Page: page, PageSize: pageSize}
}

func (p PaginationParams) Offset() int64 {

	return int64(p.Page-1) * p.PageSize
}

func CreatePaginatedResponse[T any](data []T, totalItems int64, paginationParams PaginationParams) *PaginatedResponse[T] {
	totalPages := int(math.Ceil(float64(totalItems) / float64(paginationParams.PageSize)))

	return &PaginatedResponse[T]{
		Data:        data,
		Page:        paginationParams.Page,
		PageSize:    paginationParams.PageSize,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		HasNextPage: paginationParams.Page < totalPages,
		HasPrevPage: paginationParams.Page > 1,
	}
}

type RealEstatePaginatedResponse = PaginatedResponse[*RealEstatePropertyInput] // para swagger
type NoBodyResponse struct{}                                                   // para swagger
type ErrorResponse struct {                                                    // para swagger
	Error string `json:"error"`
}
