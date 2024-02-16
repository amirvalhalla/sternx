package query

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"sternx/infrastructure/repository"
	"sternx/infrastructure/util"
)

func FindByID(id uuid.UUID) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where("id = ?", id)
	}
}

func FindByCustomColumn[T any](columnName string, value T) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where(fmt.Sprintf("%s = ?", columnName), value)
	}
}

func WithOffset(pageIndex, pageSize int) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Offset(util.CalcOffset(pageIndex, pageSize))
	}
}

func WithLimit(pageSize int) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Limit(pageSize)
	}
}

func WithOrderBy(columnName string, sort SortType) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Order(columnName + " " + sort.String())
	}
}

func WithoutAnySearch() repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q
	}
}

func WithNullColumnCondition(columnName string) repository.SelectQuery {
	return func(q *gorm.DB) *gorm.DB {
		return q.Where(fmt.Sprintf("%s IS NULL", columnName)).
			Or(fmt.Sprintf("%s = ?", columnName), "")
	}
}
