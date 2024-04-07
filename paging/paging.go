package paging

import (
	"fmt"
	"reflect"

	"github.com/yudhaananda/go-common/validation"
)

type Paging[T comparable] struct {
	Page     int    `json:"page" form:"page"`
	Take     int    `json:"take" form:"take"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
	IsActive bool   `json:"-"`
	Filter   T
}

func (f *Paging[T]) SetDefault() {
	f.Page = 1
	f.Take = 10
}

func (f *Paging[T]) QueryBuilder() (string, []any) {
	query := " WHERE 1=1"
	if f.IsActive {
		query += " AND status=1"
	}
	ref := reflect.ValueOf(f.Filter)
	tpe := ref.Type()
	args := []any{}

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if !validation.IsEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			query += " AND " + tpe.Field(i).Tag.Get("db") + "= ?"
			args = append(args, ref.Field(i).Interface())
		}
	}

	return query, args
}

func (f *Paging[T]) PaginationQuery(args []any) (string, []any) {
	query := ""
	// Adding OrderBy Statement
	if f.OrderBy != "" {
		query += " ORDER BY ?"
		args = append(args, f.OrderBy)
	}

	// Adding Limit and Offset statement
	if f.Take > 0 {
		query += " LIMIT ?"
		args = append(args, f.Take)
		if f.Page > 0 {
			query += " OFFSET ?"
			args = append(args, f.Take*(f.Page-1))

		}
	}
	return query, args
}
