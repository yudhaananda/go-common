package sql

import (
	"context"
	"database/sql"
	"reflect"
)

type DBSql[T comparable] interface {
	GetContext(ctx context.Context, query string, args ...any) (res []T, err error)
	CountContext(ctx context.Context, query string, args ...any) (res int, err error)
}

type dbSql[T comparable] struct {
	DB *sql.DB
}

func New[T comparable](db *sql.DB) DBSql[T] {
	return &dbSql[T]{
		DB: db,
	}
}

func (d *dbSql[T]) GetContext(ctx context.Context, query string, args ...any) (res []T, err error) {
	sql, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	row, err := sql.QueryContext(ctx, args...)
	if err != nil {
		return
	}

	defer row.Close()
	for row.Next() {
		var model T

		s := reflect.ValueOf(&model).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = row.Scan(columns...)
		if err != nil {
			return
		}
		res = append(res, model)
	}
	return
}

func (d *dbSql[T]) CountContext(ctx context.Context, query string, args ...any) (res int, err error) {
	sqlCount, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rowCount, err := sqlCount.QueryContext(ctx, args...)
	if err != nil {
		return
	}

	defer rowCount.Close()
	for rowCount.Next() {
		err = rowCount.Scan(&res)
		if err != nil {
			return
		}
	}
	return
}
