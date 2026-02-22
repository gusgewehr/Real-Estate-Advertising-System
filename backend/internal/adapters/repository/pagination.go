package repository

import (
	"real-estate-api/internal/infrastructure"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaginationRepo struct {
	db     *infrastructure.Database
	logger *zap.SugaredLogger
}

func NewPaginationRepo(db *infrastructure.Database, logger *zap.SugaredLogger) *PaginationRepo {
	return &PaginationRepo{db, logger}
}

func (r *PaginationRepo) GetTotalItems(table string) int64 {
	var total int64

	result := r.db.Orm.Table(table).Select("count(1) as total").Scan(&total)
	if result.Error != nil {
		r.logger.Errorw("Error getting total items", "error", result.Error)
	}

	return total
}

func (r *PaginationRepo) GetTableName(structTable interface{}) string {
	stmt := &gorm.Statement{DB: r.db.Orm}
	err := stmt.Parse(&structTable)
	if err != nil {
		r.logger.Errorw("Error parsing table name", "error", err)
		return ""
	}

	return stmt.Schema.Table
}
