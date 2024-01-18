package util

import (
	"estore/protoc/typepb"

	"gorm.io/gorm/clause"
)

// 分頁
func LimitOffset(a *typepb.Request) clause.Expression {

	offset := 0
	size := 20

	if a == nil {
		return clause.Limit{Limit: &size, Offset: offset}
	}

	// 分頁
	if a.Pagination != nil {
		if a.Pagination.Size > 0 {
			size = int(a.Pagination.Size)
		}
		if a.Pagination.Page > 0 {
			offset = int(a.Pagination.Page) * size
		}
	}

	return clause.Limit{Limit: &size, Offset: offset}

}
