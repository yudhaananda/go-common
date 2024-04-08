package htmxmodel

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"time"

	"github.com/yudhaananda/go-common/formatter"
)

type HTMX[T comparable] struct {
	Model T
}

func (m HTMX[T]) GenerateHTML(html string) (result HTMXResult) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if ref.Field(i).CanConvert(reflect.ValueOf(time.Time{}).Type()) {
			result.DateJQuery = append(result.DateJQuery, DateJQuery{Value: tpe.Field(i).Tag.Get("form")})
		}
		form := tpe.Field(i).Tag.Get("form")
		name := tpe.Field(i).Tag.Get("name")
		memberType := tpe.Field(i).Tag.Get("type")
		result.Members = append(result.Members, MemberStruct{
			Member: template.HTML(fmt.Sprintf(html, memberType, form, form, form, name)),
		})
	}
	return
}

func (m HTMX[T]) ToColumn(name string) (result Column) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}

		if tpe.Field(i).Tag.Get("header") == "Id" {
			result.Id = template.HTML(fmt.Sprint(ref.Field(i).Interface()))
		}

		member := processValue(ref.Field(i).Type().Name(), time.DateOnly, ref.Field(i).Interface())

		result.Row = append(result.Row, MemberStruct{
			Member: member,
		})
	}
	result.Name = template.HTML(name)
	return
}

func processValue(name, timeFormat string, val interface{}) (member template.HTML) {
	if strings.Contains(name, "Null") {
		switch s := val.(type) {
		case formatter.Null[string]:
			if s.Valid {
				member = template.HTML(s.Data)
			} else {
				member = ""
			}
		case formatter.Null[int]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[int64]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[int32]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[int16]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[int8]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[float32]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[float64]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		case formatter.Null[time.Time]:
			if s.Valid {
				member = template.HTML(s.Data.Format(timeFormat))
			} else {
				member = ""
			}
		case formatter.Null[bool]:
			if s.Valid {
				member = template.HTML(fmt.Sprint(s.Data))
			} else {
				member = ""
			}
		default:
			member = template.HTML(fmt.Sprint(val))
		}

	} else {
		member = template.HTML(fmt.Sprint(val))
	}

	return
}

func (m HTMX[T]) ToHeader() (result HTMXGet) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("header") == "-" {
			continue
		}
		result.Header = append(result.Header, MemberStruct{
			Member: template.HTML(tpe.Field(i).Tag.Get("header")),
		})
	}
	return
}

func (m HTMX[T]) ToModalMember() (result []ModalMember) {
	ref := reflect.ValueOf(m.Model)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		if tpe.Field(i).Tag.Get("type") == "-" {
			continue
		}
		result = append(result,
			ModalMember{
				Id:          template.HTML(tpe.Field(i).Tag.Get("json")),
				Type:        template.HTML(tpe.Field(i).Tag.Get("type")),
				Name:        template.HTML(tpe.Field(i).Tag.Get("json")),
				Value:       processValue(ref.Field(i).Type().Name(), "2006-01-02T15:04:05Z", ref.Field(i).Interface()),
				Placeholder: template.HTML(tpe.Field(i).Tag.Get("header")),
			})
	}
	return
}
