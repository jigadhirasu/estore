package util

import (
	"estore/protoc/typepb"

	"github.com/IBM/fp-go/array"
	"gorm.io/gorm/clause"
)

// 排序
func OrderBy(defaultOrder clause.OrderBy) func(a *typepb.Request) clause.Expression {
	return func(a *typepb.Request) clause.Expression {
		if a == nil {
			return defaultOrder
		}

		if len(a.OrderBy) < 1 {
			return defaultOrder
		}

		orderby := array.Map[*typepb.OrderBy, clause.OrderByColumn](func(a *typepb.OrderBy) clause.OrderByColumn {
			return clause.OrderByColumn{
				Column: clause.Column{Name: a.Field},
				Desc:   a.Order == typepb.Order_DESC,
			}
		})(a.OrderBy)

		return clause.OrderBy{Columns: orderby}
	}
}
