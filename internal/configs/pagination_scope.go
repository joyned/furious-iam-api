package configs

import (
	"furious/iam-api/pkg/utils"
	"math"

	"gorm.io/gorm"
)

func Paginate(value any, pagination *utils.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	if pagination.Limit == 0 {
		pagination.Limit = 10
	}

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
