package filter

import (
	"fmt"
	"html/template"
	"reflect"
	"time"

	htmxmodel "github.com/yudhaananda/go-common/htmx_model"
	"github.com/yudhaananda/go-common/validation"
)

func ToHTMXFilter(model any) (result []htmxmodel.HTMXFilter, dateFilter []htmxmodel.DateJQuery) {
	ref := reflect.ValueOf(model)
	tpe := ref.Type()
	for i := 0; i < tpe.NumField(); i++ {
		if ref.Field(i).CanConvert(reflect.ValueOf(time.Time{}).Type()) {
			dateFilter = append(dateFilter, htmxmodel.DateJQuery{Value: tpe.Field(i).Tag.Get("form")})
		}
		value := template.HTML(fmt.Sprint(ref.Field(i).Interface()))
		if validation.IsEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			value = ""
		}
		result = append(result, htmxmodel.HTMXFilter{
			Type:  template.HTML(tpe.Field(i).Tag.Get("type")),
			Id:    template.HTML(tpe.Field(i).Tag.Get("form")),
			Label: template.HTML(tpe.Field(i).Tag.Get("name")),
			Value: value,
		})
	}
	return
}
