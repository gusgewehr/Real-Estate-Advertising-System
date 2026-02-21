package test

import (
	"real-state-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPaginationParams_Normal(t *testing.T) {
	params := domain.BuildPaginationParams(2, 10, 100)

	assert.Equal(t, 2, params.Page)
	assert.Equal(t, int64(10), params.PageSize)
}

func TestBuildPaginationParams_PageSizeMinusOne(t *testing.T) {
	params := domain.BuildPaginationParams(1, -1, 50)

	assert.Equal(t, 1, params.Page)
	assert.Equal(t, int64(50), params.PageSize)
}

func TestPaginationParams_Offset_FirstPage(t *testing.T) {
	params := domain.PaginationParams{Page: 1, PageSize: 10}
	assert.Equal(t, int64(0), params.Offset())
}

func TestPaginationParams_Offset_SecondPage(t *testing.T) {
	params := domain.PaginationParams{Page: 2, PageSize: 10}
	assert.Equal(t, int64(10), params.Offset())
}

func TestPaginationParams_Offset_ThirdPage(t *testing.T) {
	params := domain.PaginationParams{Page: 3, PageSize: 25}
	assert.Equal(t, int64(50), params.Offset())
}

func TestCreatePaginatedResponse_FirstPage(t *testing.T) {
	data := []string{"a", "b", "c"}
	params := domain.PaginationParams{Page: 1, PageSize: 3}
	resp := domain.CreatePaginatedResponse(data, 10, params)

	assert.Equal(t, 3, len(resp.Data))
	assert.Equal(t, 1, resp.Page)
	assert.Equal(t, int64(3), resp.PageSize)
	assert.Equal(t, int64(10), resp.TotalItems)
	assert.Equal(t, 4, resp.TotalPages)
	assert.True(t, resp.HasNextPage)
	assert.False(t, resp.HasPrevPage)
}

func TestCreatePaginatedResponse_MiddlePage(t *testing.T) {
	data := []string{"d", "e", "f"}
	params := domain.PaginationParams{Page: 2, PageSize: 3}
	resp := domain.CreatePaginatedResponse(data, 10, params)

	assert.True(t, resp.HasNextPage)
	assert.True(t, resp.HasPrevPage)
}

func TestCreatePaginatedResponse_LastPage(t *testing.T) {
	data := []string{"j"}
	params := domain.PaginationParams{Page: 4, PageSize: 3}
	resp := domain.CreatePaginatedResponse(data, 10, params)

	assert.False(t, resp.HasNextPage)
	assert.True(t, resp.HasPrevPage)
	assert.Equal(t, 4, resp.TotalPages)
}

func TestCreatePaginatedResponse_EmptyData(t *testing.T) {
	var data []string
	params := domain.PaginationParams{Page: 1, PageSize: 10}
	resp := domain.CreatePaginatedResponse(data, 0, params)

	assert.Equal(t, 0, len(resp.Data))
	assert.Equal(t, int64(0), resp.TotalItems)
	assert.Equal(t, 0, resp.TotalPages)
	assert.False(t, resp.HasNextPage)
	assert.False(t, resp.HasPrevPage)
}
