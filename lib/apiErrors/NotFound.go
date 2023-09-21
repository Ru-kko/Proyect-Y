package apierrors

import (
	"Proyect-Y/common/errors"
	"fmt"
)

type ServiceNotFound struct {
	*errors.ProyectYError
	Name    string `json:"name"`
	Message string `json:"message"`
	Paht    string `json:"path"`
}

func (e ServiceNotFound) Error() string {
	return fmt.Sprintf("%s - %s", e.Message, e.Paht)
}
