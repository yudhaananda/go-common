package querybuilder

import (
	"fmt"
	"reflect"
	"time"
)

func BuildTableMember(model any) string {
	member := ""

	ref := reflect.ValueOf(model)
	tpe := ref.Type()

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if isQueryNeedComa {
			member += ", "
		}
		member += tpe.Field(i).Tag.Get("db")
		isQueryNeedComa = true
	}
	return member
}

func BuildCreateQuery(model any) (string, []any) {
	query := " "
	args := []any{}

	ref := reflect.ValueOf(model)
	tpe := ref.Type()

	table := "("
	values := "("

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if !isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			if isQueryNeedComa {
				table += ", "
				values += ", "
			}
			table += tpe.Field(i).Tag.Get("db")

			isQueryNeedComa = true
			values += "?"
			args = append(args, ref.Field(i).Interface())
		}
		if i == tpe.NumField()-1 {
			table += ")"
			values += ")"
		}
	}
	query += table + " VALUES " + values

	return query, args
}

func BuildUpdateQuery(id int, model any) (string, []any) {
	query := " SET "
	args := []any{}

	ref := reflect.ValueOf(model)
	tpe := ref.Type()

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if !isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			if isQueryNeedComa {
				query += ", "
			}

			isQueryNeedComa = true

			query += tpe.Field(i).Tag.Get("db") + "=?"
			args = append(args, ref.Field(i).Interface())
		}
	}
	query += " WHERE id=?"
	args = append(args, id)
	return query, args
}

func isEmpty(check string) bool {
	return check == "0" || check == "" || check == fmt.Sprint(time.Time{})
}
