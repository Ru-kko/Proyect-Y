package util

import "reflect"

func StructToMap(s interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	typo := reflect.TypeOf(s)
	valor := reflect.ValueOf(s)

	if typo.Kind() == reflect.Struct {
		for i := 0; i < typo.NumField(); i++ {
			key := typo.Field(i)
			value := valor.Field(i).Interface()
			if !reflect.ValueOf(value).IsZero() {
				res[key.Name] = value
			}
		}
	}

	return res
}
