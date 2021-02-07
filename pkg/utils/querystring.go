package utils

import (
	"reflect"
	"strconv"
	"strings"
)

//QueryString !
func QueryString(r interface{}) string {
	o := reflect.ValueOf(r)
	s := ""

	for i := 0; i < o.NumField(); i++ {
		p := LowerFirstChar(o.Type().Field(i).Name)
		s = s + p

		f := o.Field(i)
		switch f.Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			s = s + "=" + strconv.FormatInt(f.Int(), 10)
		default:
			s = s + "=" + f.String()
		}

		s = s + "&"

	}
	return strings.TrimSuffix(s, "&")
}
