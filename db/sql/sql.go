package sql

import (
	"context"
	"database/sql"
	"reflect"
)

type DBSql struct {
	DB *sql.DB
}

func (d *DBSql) GetContext(ctx context.Context, res []interface{}, query string, args ...any) (err error) {
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
		var model interface{}

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
		res = append(res, &model)
	}
	return
}

func (d *DBSql) CountContext(ctx context.Context, res int, query string, args ...any) (err error) {
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
