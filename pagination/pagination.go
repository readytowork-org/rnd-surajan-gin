package pagination

import (
	"strconv"

	"gorm.io/gorm"
)

func DefaultPageVariables() (int, int) {
	defaultPageNo := 1
	defaultPageSize := 5
	return defaultPageNo, defaultPageSize
}

func Paginate(pageNo, pageSize string, defaultPageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// strconv.Atoi converts string to int
		limit, _ := strconv.Atoi(pageSize)
		pgNo, _ := strconv.Atoi(pageNo)
		// Don't need to check if pgNo is <= 0, if it is, data will be shown of pgNo=1
		// Calculate Offset
		if limit > 100 {
			limit = 100
		} else if limit <= 0 {
			limit = defaultPageSize
		}
		offset := (pgNo - 1) * limit
		// Return limit and offset logic
		return db.Offset(offset).Limit(limit)
	}
}
