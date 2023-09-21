package env

import (
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	if os.Getenv("PROD") == "true" {
		return
	}

	godotenv.Load()
}

func Parse[T interface{}](obj *T) error {
	rType := reflect.TypeOf(*obj)
	fieldNum := rType.NumField()

	values := reflect.ValueOf(obj).Elem()

	for i := 0; i < fieldNum; i++ {
		field := rType.Field(i)
		fieldTypo := field.Type.Kind()

		val := os.Getenv(field.Name)

		if val == "" {
			return EnviromentNotSet{
				Name: field.Name,
			}
		}

		switch fieldTypo {
		case reflect.String:
			values.Field(i).SetString(val)
			break
		case reflect.Int:
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			values.Field(i).SetInt(int64(intVal))
		default:
			panic("InsertEnv only accepts values of type int or string")
		}
	}

	return nil
}
