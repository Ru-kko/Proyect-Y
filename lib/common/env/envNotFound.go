package env

import (
	"Proyect-Y/common/errors"
	"fmt"
)

type EnviromentNotSet struct {
	errors.ProyectYError
	Name string
}

func (e EnviromentNotSet) Error() string {
	return fmt.Sprintf("Please set %s enviroment on your .env file or set PROD='true' os enviroment", e.Name)
}
