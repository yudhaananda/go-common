package sql

import (
	"context"
	"database/sql"
	"reflect"
)

type Tx sql.Tx

type DBSql[T comparable] struct {
	DB *sql.DB
}

func (d *DBSql[T]) ExecContext(ctx context.Context, query string, trx *sql.Tx, args ...any) (id int64, err error) {
	if trx == nil {
		trx, err = d.DB.Begin()
		if err != nil {
			return
		}

		defer trx.Commit()
	}

	sql, err := trx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	sqlRes, err := sql.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	id, err = sqlRes.LastInsertId()
	if err != nil {
		return
	}

	return
}

func (d *DBSql[T]) GetContext(ctx context.Context, query string, args ...any) (res []T, err error) {
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

func (d *DBSql[T]) CountContext(ctx context.Context, res *int, query string, args ...any) (err error) {
	sql, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rowCount, err := sql.QueryContext(ctx, args...)
	if err != nil {
		return
	}

	defer rowCount.Close()
	for rowCount.Next() {
		err = rowCount.Scan(res)
		if err != nil {
			return
		}
	}
	return
}
